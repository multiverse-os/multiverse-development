// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//
// Portions Copyright 2017 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the THIRD-PARTY file.

use epoll;
use std::cmp;
use std::fs::File;
use std::io::{self, Read, Seek, SeekFrom, Write};
use std::os::linux::fs::MetadataExt;
use std::os::unix::io::{AsRawFd, RawFd};
use std::result;
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::mpsc;
use std::sync::Arc;

use logger::{Metric, METRICS};
use memory_model::{GuestAddress, GuestMemory, GuestMemoryError};
use rate_limiter::{RateLimiter, TokenType};
use sys_util::EventFd;
use virtio_gen::virtio_blk::*;

use super::{
    ActivateError, ActivateResult, DescriptorChain, EpollConfigConstructor, Queue, VirtioDevice,
    TYPE_BLOCK, VIRTIO_MMIO_INT_VRING,
};
use crate::{DeviceEventT, EpollHandler, Error as DeviceError};

const CONFIG_SPACE_SIZE: usize = 8;
const SECTOR_SHIFT: u8 = 9;
pub const SECTOR_SIZE: u64 = (0x01 as u64) << SECTOR_SHIFT;
const QUEUE_SIZE: u16 = 256;
const NUM_QUEUES: usize = 1;
const QUEUE_SIZES: &[u16] = &[QUEUE_SIZE];

// New descriptors are pending on the virtio queue.
const QUEUE_AVAIL_EVENT: DeviceEventT = 0;
// Rate limiter budget is now available.
const RATE_LIMITER_EVENT: DeviceEventT = 1;
// Number of DeviceEventT events supported by this implementation.
pub const BLOCK_EVENTS_COUNT: usize = 2;

#[derive(Debug)]
enum Error {
    /// Guest gave us bad memory addresses.
    GuestMemory(GuestMemoryError),
    /// Guest gave us offsets that would have overflowed a usize.
    CheckedOffset(GuestAddress, usize),
    /// Guest gave us a write only descriptor that protocol says to read from.
    UnexpectedWriteOnlyDescriptor,
    /// Guest gave us a read only descriptor that protocol says to write to.
    UnexpectedReadOnlyDescriptor,
    /// Guest gave us too few descriptors in a descriptor chain.
    DescriptorChainTooShort,
    /// Guest gave us a descriptor that was too short to use.
    DescriptorLengthTooSmall,
    /// Getting a block's metadata fails for any reason.
    GetFileMetadata,
    /// The requested operation would cause a seek beyond disk end.
    InvalidOffset,
}

#[derive(Debug)]
enum ExecuteError {
    BadRequest(Error),
    Flush(io::Error),
    Read(GuestMemoryError),
    Seek(io::Error),
    Write(GuestMemoryError),
    Unsupported(u32),
}

impl ExecuteError {
    fn status(&self) -> u32 {
        match *self {
            ExecuteError::BadRequest(_) => VIRTIO_BLK_S_IOERR,
            ExecuteError::Flush(_) => VIRTIO_BLK_S_IOERR,
            ExecuteError::Read(_) => VIRTIO_BLK_S_IOERR,
            ExecuteError::Seek(_) => VIRTIO_BLK_S_IOERR,
            ExecuteError::Write(_) => VIRTIO_BLK_S_IOERR,
            ExecuteError::Unsupported(_) => VIRTIO_BLK_S_UNSUPP,
        }
    }
}

#[derive(Clone, Copy, Debug, PartialEq)]
enum RequestType {
    In,
    Out,
    Flush,
    GetDeviceID,
    Unsupported(u32),
}

fn request_type(mem: &GuestMemory, desc_addr: GuestAddress) -> result::Result<RequestType, Error> {
    let type_ = mem
        .read_obj_from_addr(desc_addr)
        .map_err(Error::GuestMemory)?;
    match type_ {
        VIRTIO_BLK_T_IN => Ok(RequestType::In),
        VIRTIO_BLK_T_OUT => Ok(RequestType::Out),
        VIRTIO_BLK_T_FLUSH => Ok(RequestType::Flush),
        VIRTIO_BLK_T_GET_ID => Ok(RequestType::GetDeviceID),
        t => Ok(RequestType::Unsupported(t)),
    }
}

fn sector(mem: &GuestMemory, desc_addr: GuestAddress) -> result::Result<u64, Error> {
    const SECTOR_OFFSET: usize = 8;
    let addr = match mem.checked_offset(desc_addr, SECTOR_OFFSET) {
        Some(v) => v,
        None => return Err(Error::CheckedOffset(desc_addr, SECTOR_OFFSET)),
    };

    mem.read_obj_from_addr(addr).map_err(Error::GuestMemory)
}

fn build_device_id(disk_image: &File) -> result::Result<String, Error> {
    let blk_metadata = match disk_image.metadata() {
        Err(_) => return Err(Error::GetFileMetadata),
        Ok(m) => m,
    };
    // This is how kvmtool does it.
    let device_id = format!(
        "{}{}{}",
        blk_metadata.st_dev(),
        blk_metadata.st_rdev(),
        blk_metadata.st_ino()
    )
    .to_owned();
    Ok(device_id)
}

fn build_disk_image_id(disk_image: &File) -> Vec<u8> {
    let mut default_disk_image_id = vec![0; VIRTIO_BLK_ID_BYTES as usize];
    match build_device_id(disk_image) {
        Err(_) => {
            warn!("Could not generate device id. We'll use a default.");
        }
        Ok(m) => {
            // The kernel only knows to read a maximum of VIRTIO_BLK_ID_BYTES.
            // This will also zero out any leftover bytes.
            let disk_id = m.as_bytes();
            let bytes_to_copy = cmp::min(disk_id.len(), VIRTIO_BLK_ID_BYTES as usize);
            default_disk_image_id[..bytes_to_copy].clone_from_slice(&disk_id[..bytes_to_copy])
        }
    }
    default_disk_image_id
}

struct Request {
    request_type: RequestType,
    sector: u64,
    data_addr: GuestAddress,
    data_len: u32,
    status_addr: GuestAddress,
}

impl Request {
    fn parse(avail_desc: &DescriptorChain, mem: &GuestMemory) -> result::Result<Request, Error> {
        // The head contains the request type which MUST be readable.
        if avail_desc.is_write_only() {
            return Err(Error::UnexpectedWriteOnlyDescriptor);
        }

        let mut req = Request {
            request_type: request_type(&mem, avail_desc.addr)?,
            sector: sector(&mem, avail_desc.addr)?,
            data_addr: GuestAddress(0),
            data_len: 0,
            status_addr: GuestAddress(0),
        };

        let data_desc;
        let status_desc;
        let desc = avail_desc
            .next_descriptor()
            .ok_or(Error::DescriptorChainTooShort)?;

        if !desc.has_next() {
            status_desc = desc;
            // Only flush requests are allowed to skip the data descriptor.
            if req.request_type != RequestType::Flush {
                return Err(Error::DescriptorChainTooShort);
            }
        } else {
            data_desc = desc;
            status_desc = data_desc
                .next_descriptor()
                .ok_or(Error::DescriptorChainTooShort)?;

            if data_desc.is_write_only() && req.request_type == RequestType::Out {
                return Err(Error::UnexpectedWriteOnlyDescriptor);
            }
            if !data_desc.is_write_only() && req.request_type == RequestType::In {
                return Err(Error::UnexpectedReadOnlyDescriptor);
            }
            if !data_desc.is_write_only() && req.request_type == RequestType::GetDeviceID {
                return Err(Error::UnexpectedReadOnlyDescriptor);
            }

            req.data_addr = data_desc.addr;
            req.data_len = data_desc.len;
        }

        // The status MUST always be writable.
        if !status_desc.is_write_only() {
            return Err(Error::UnexpectedReadOnlyDescriptor);
        }

        if status_desc.len < 1 {
            return Err(Error::DescriptorLengthTooSmall);
        }

        req.status_addr = status_desc.addr;

        Ok(req)
    }

    fn execute<T: Seek + Read + Write>(
        &self,
        disk: &mut T,
        disk_nsectors: u64,
        mem: &GuestMemory,
        disk_id: &[u8],
    ) -> result::Result<u32, ExecuteError> {
        let mut top: u64 = u64::from(self.data_len) / SECTOR_SIZE;
        if u64::from(self.data_len) % SECTOR_SIZE != 0 {
            top += 1;
        }
        top = top
            .checked_add(self.sector)
            .ok_or(ExecuteError::BadRequest(Error::InvalidOffset))?;
        if top > disk_nsectors {
            return Err(ExecuteError::BadRequest(Error::InvalidOffset));
        }

        disk.seek(SeekFrom::Start(self.sector << SECTOR_SHIFT))
            .map_err(ExecuteError::Seek)?;

        match self.request_type {
            RequestType::In => {
                mem.read_to_memory(self.data_addr, disk, self.data_len as usize)
                    .map_err(ExecuteError::Read)?;
                METRICS.block.read_bytes.add(self.data_len as usize);
                METRICS.block.read_count.inc();
                return Ok(self.data_len);
            }
            RequestType::Out => {
                mem.write_from_memory(self.data_addr, disk, self.data_len as usize)
                    .map_err(ExecuteError::Write)?;
                METRICS.block.write_bytes.add(self.data_len as usize);
                METRICS.block.write_count.inc();
            }
            RequestType::Flush => match disk.flush() {
                Ok(_) => {
                    METRICS.block.flush_count.inc();
                    return Ok(0);
                }
                Err(e) => return Err(ExecuteError::Flush(e)),
            },
            RequestType::GetDeviceID => {
                if (self.data_len as usize) < disk_id.len() {
                    return Err(ExecuteError::BadRequest(Error::InvalidOffset));
                }
                mem.write_slice_at_addr(disk_id, self.data_addr)
                    .map_err(ExecuteError::Write)?;
            }
            RequestType::Unsupported(t) => return Err(ExecuteError::Unsupported(t)),
        };
        Ok(0)
    }
}

/// Handler that drives the execution of the Block devices
pub struct BlockEpollHandler {
    queues: Vec<Queue>,
    mem: GuestMemory,
    disk_image: File,
    disk_nsectors: u64,
    interrupt_status: Arc<AtomicUsize>,
    interrupt_evt: EventFd,
    queue_evt: EventFd,
    rate_limiter: RateLimiter,
    disk_image_id: Vec<u8>,
}

impl BlockEpollHandler {
    fn process_queue(&mut self, queue_index: usize) -> bool {
        let queue = &mut self.queues[queue_index];
        let mut used_any = false;

        while let Some(head) = queue.pop(&self.mem) {
            let len;
            match Request::parse(&head, &self.mem) {
                Ok(request) => {
                    // If limiter.consume() fails it means there is no more TokenType::Ops
                    // budget and rate limiting is in effect.
                    if !self.rate_limiter.consume(1, TokenType::Ops) {
                        // Stop processing the queue and return this descriptor chain to the
                        // avail ring, for later processing.
                        queue.undo_pop();
                        break;
                    }
                    // Exercise the rate limiter only if this request is of data transfer type.
                    if request.request_type == RequestType::In
                        || request.request_type == RequestType::Out
                    {
                        // If limiter.consume() fails it means there is no more TokenType::Bytes
                        // budget and rate limiting is in effect.
                        if !self
                            .rate_limiter
                            .consume(u64::from(request.data_len), TokenType::Bytes)
                        {
                            // Revert the OPS consume().
                            self.rate_limiter.manual_replenish(1, TokenType::Ops);
                            // Stop processing the queue and return this descriptor chain to the
                            // avail ring, for later processing.
                            queue.undo_pop();
                            break;
                        }
                    }
                    let status = match request.execute(
                        &mut self.disk_image,
                        self.disk_nsectors,
                        &self.mem,
                        &self.disk_image_id,
                    ) {
                        Ok(l) => {
                            len = l;
                            VIRTIO_BLK_S_OK
                        }
                        Err(e) => {
                            error!("Failed to execute request: {:?}", e);
                            METRICS.block.invalid_reqs_count.inc();
                            len = 1; // We need at least 1 byte for the status.
                            e.status()
                        }
                    };
                    // We use unwrap because the request parsing process already checked that the
                    // status_addr was valid.
                    self.mem
                        .write_obj_at_addr(status, request.status_addr)
                        .unwrap();
                }
                Err(e) => {
                    error!("Failed to parse available descriptor chain: {:?}", e);
                    METRICS.block.execute_fails.inc();
                    len = 0;
                }
            }
            queue.add_used(&self.mem, head.index, len);
            used_any = true;
        }

        used_any
    }

    fn signal_used_queue(&self) -> result::Result<(), DeviceError> {
        self.interrupt_status
            .fetch_or(VIRTIO_MMIO_INT_VRING as usize, Ordering::SeqCst);
        self.interrupt_evt.write(1).map_err(|e| {
            error!("Failed to signal used queue: {:?}", e);
            METRICS.block.event_fails.inc();
            DeviceError::FailedSignalingUsedQueue(e)
        })
    }

    /// Update the backing file for the Block device
    pub fn update_disk_image(&mut self, disk_image: File) -> result::Result<(), DeviceError> {
        self.disk_image = disk_image;
        self.disk_nsectors = self
            .disk_image
            .seek(SeekFrom::End(0))
            .map_err(DeviceError::IoError)?
            / SECTOR_SIZE;
        self.disk_image_id = build_disk_image_id(&self.disk_image);
        METRICS.block.update_count.inc();
        Ok(())
    }
}

impl EpollHandler for BlockEpollHandler {
    fn handle_event(
        &mut self,
        device_event: DeviceEventT,
        _evset: epoll::Events,
    ) -> result::Result<(), DeviceError> {
        match device_event {
            QUEUE_AVAIL_EVENT => {
                METRICS.block.queue_event_count.inc();
                if let Err(e) = self.queue_evt.read() {
                    error!("Failed to get queue event: {:?}", e);
                    METRICS.block.event_fails.inc();
                    Err(DeviceError::FailedReadingQueue {
                        event_type: "queue event",
                        underlying: e,
                    })
                } else if !self.rate_limiter.is_blocked() && self.process_queue(0) {
                    self.signal_used_queue()
                } else {
                    // While limiter is blocked, don't process any more requests.
                    Ok(())
                }
            }
            RATE_LIMITER_EVENT => {
                METRICS.block.rate_limiter_event_count.inc();
                // Upon rate limiter event, call the rate limiter handler
                // and restart processing the queue.
                if self.rate_limiter.event_handler().is_ok() && self.process_queue(0) {
                    self.signal_used_queue()
                } else {
                    Ok(())
                }
            }
            unknown => Err(DeviceError::UnknownEvent {
                device: "block",
                event: unknown,
            }),
        }
    }
}

pub struct EpollConfig {
    q_avail_token: u64,
    rate_limiter_token: u64,
    epoll_raw_fd: RawFd,
    sender: mpsc::Sender<Box<dyn EpollHandler>>,
}

impl EpollConfigConstructor for EpollConfig {
    fn new(
        first_token: u64,
        epoll_raw_fd: RawFd,
        sender: mpsc::Sender<Box<dyn EpollHandler>>,
    ) -> Self {
        EpollConfig {
            q_avail_token: first_token + u64::from(QUEUE_AVAIL_EVENT),
            rate_limiter_token: first_token + u64::from(RATE_LIMITER_EVENT),
            epoll_raw_fd,
            sender,
        }
    }
}

/// Virtio device for exposing block level read/write operations on a host file.
pub struct Block {
    disk_image: Option<File>,
    disk_nsectors: u64,
    avail_features: u64,
    acked_features: u64,
    config_space: Vec<u8>,
    epoll_config: EpollConfig,
    rate_limiter: Option<RateLimiter>,
}

pub fn build_config_space(disk_size: u64) -> Vec<u8> {
    // We only support disk size, which uses the first two words of the configuration space.
    // If the image is not a multiple of the sector size, the tail bits are not exposed.
    // The config space is little endian.
    let mut config = Vec::with_capacity(CONFIG_SPACE_SIZE);
    let num_sectors = disk_size >> SECTOR_SHIFT;
    for i in 0..8 {
        config.push((num_sectors >> (8 * i)) as u8);
    }
    config
}

impl Block {
    /// Create a new virtio block device that operates on the given file.
    ///
    /// The given file must be seekable and sizable.
    pub fn new(
        mut disk_image: File,
        is_disk_read_only: bool,
        epoll_config: EpollConfig,
        rate_limiter: Option<RateLimiter>,
    ) -> io::Result<Block> {
        let disk_size = disk_image.seek(SeekFrom::End(0))? as u64;
        if disk_size % SECTOR_SIZE != 0 {
            warn!(
                "Disk size {} is not a multiple of sector size {}; \
                 the remainder will not be visible to the guest.",
                disk_size, SECTOR_SIZE
            );
        }

        let mut avail_features = (1u64 << VIRTIO_F_VERSION_1) | (1u64 << VIRTIO_BLK_F_FLUSH);

        if is_disk_read_only {
            avail_features |= 1u64 << VIRTIO_BLK_F_RO;
        };

        Ok(Block {
            disk_image: Some(disk_image),
            disk_nsectors: disk_size / SECTOR_SIZE,
            avail_features,
            acked_features: 0u64,
            config_space: build_config_space(disk_size),
            epoll_config,
            rate_limiter,
        })
    }
}

impl VirtioDevice for Block {
    fn device_type(&self) -> u32 {
        TYPE_BLOCK
    }

    fn queue_max_sizes(&self) -> &[u16] {
        QUEUE_SIZES
    }

    fn avail_features(&self) -> u64 {
        self.avail_features
    }

    fn acked_features(&self) -> u64 {
        self.acked_features
    }

    fn set_acked_features(&mut self, acked_features: u64) {
        self.acked_features = acked_features;
    }

    fn read_config(&self, offset: u64, mut data: &mut [u8]) {
        let config_len = self.config_space.len() as u64;
        if offset >= config_len {
            error!("Failed to read config space");
            METRICS.block.cfg_fails.inc();
            return;
        }
        if let Some(end) = offset.checked_add(data.len() as u64) {
            // This write can't fail, offset and end are checked against config_len.
            data.write_all(&self.config_space[offset as usize..cmp::min(end, config_len) as usize])
                .unwrap();
        }
    }

    fn write_config(&mut self, offset: u64, data: &[u8]) {
        let data_len = data.len() as u64;
        let config_len = self.config_space.len() as u64;
        if offset + data_len > config_len {
            error!("Failed to write config space");
            METRICS.block.cfg_fails.inc();
            return;
        }
        let (_, right) = self.config_space.split_at_mut(offset as usize);
        right.copy_from_slice(&data[..]);
    }

    fn activate(
        &mut self,
        mem: GuestMemory,
        interrupt_evt: EventFd,
        status: Arc<AtomicUsize>,
        queues: Vec<Queue>,
        mut queue_evts: Vec<EventFd>,
    ) -> ActivateResult {
        if queues.len() != NUM_QUEUES || queue_evts.len() != NUM_QUEUES {
            error!(
                "Cannot perform activate. Expected {} queue(s), got {}",
                NUM_QUEUES,
                queues.len()
            );
            METRICS.block.activate_fails.inc();
            return Err(ActivateError::BadActivate);
        }

        if let Some(disk_image) = self.disk_image.take() {
            let queue_evt = queue_evts.remove(0);
            let queue_evt_raw_fd = queue_evt.as_raw_fd();

            let disk_image_id = build_disk_image_id(&disk_image);
            let handler = BlockEpollHandler {
                queues,
                mem,
                disk_image,
                disk_nsectors: self.disk_nsectors,
                interrupt_status: status,
                interrupt_evt,
                queue_evt,
                rate_limiter: self.rate_limiter.take().unwrap_or_default(),
                disk_image_id,
            };
            let rate_limiter_rawfd = handler.rate_limiter.as_raw_fd();

            // The channel should be open at this point.
            self.epoll_config
                .sender
                .send(Box::new(handler))
                .expect("Failed to send through the channel");

            //TODO: barrier needed here by any chance?
            epoll::ctl(
                self.epoll_config.epoll_raw_fd,
                epoll::ControlOptions::EPOLL_CTL_ADD,
                queue_evt_raw_fd,
                epoll::Event::new(epoll::Events::EPOLLIN, self.epoll_config.q_avail_token),
            )
            .map_err(|e| {
                METRICS.block.activate_fails.inc();
                ActivateError::EpollCtl(e)
            })?;

            if rate_limiter_rawfd != -1 {
                epoll::ctl(
                    self.epoll_config.epoll_raw_fd,
                    epoll::ControlOptions::EPOLL_CTL_ADD,
                    rate_limiter_rawfd,
                    epoll::Event::new(epoll::Events::EPOLLIN, self.epoll_config.rate_limiter_token),
                )
                .map_err(|e| {
                    METRICS.block.activate_fails.inc();
                    ActivateError::EpollCtl(e)
                })?;
            }

            return Ok(());
        }
        METRICS.block.activate_fails.inc();
        Err(ActivateError::BadActivate)
    }
}

#[cfg(test)]
mod tests {
    extern crate tempfile;

    use self::tempfile::{tempfile, NamedTempFile};
    use super::*;

    use libc;
    use std::fs::{metadata, OpenOptions};
    use std::sync::mpsc::Receiver;
    use std::thread;
    use std::time::Duration;
    use std::u32;

    use crate::virtio::queue::tests::*;

    const EPOLLIN: epoll::Events = epoll::Events::EPOLLIN;

    /// Will read $metric, run the code in $block, then assert metric has increased by $delta.
    macro_rules! check_metric_after_block {
        ($metric:expr, $delta:expr, $block:expr) => {{
            let before = $metric.count();
            let _ = $block;
            assert_eq!($metric.count(), before + $delta, "unexpected metric value");
        }};
    }

    impl BlockEpollHandler {
        fn set_queue(&mut self, idx: usize, q: Queue) {
            self.queues[idx] = q;
        }

        fn get_rate_limiter(&self) -> &RateLimiter {
            &self.rate_limiter
        }

        fn set_rate_limiter(&mut self, rate_limiter: RateLimiter) {
            self.rate_limiter = rate_limiter;
        }
    }

    struct DummyBlock {
        block: Block,
        epoll_raw_fd: i32,
        _receiver: Receiver<Box<EpollHandler>>,
    }

    impl DummyBlock {
        fn new(is_disk_read_only: bool) -> Self {
            let epoll_raw_fd = epoll::create(true).unwrap();
            let (sender, _receiver) = mpsc::channel();

            let epoll_config = EpollConfig::new(0, epoll_raw_fd, sender);

            let f: File = tempfile().unwrap();
            f.set_len(0x1000).unwrap();

            // Rate limiting is enabled but with a high operation rate (10 million ops/s).
            let rate_limiter = RateLimiter::new(0, None, 0, 100_000, None, 10).unwrap();
            DummyBlock {
                block: Block::new(f, is_disk_read_only, epoll_config, Some(rate_limiter)).unwrap(),
                epoll_raw_fd,
                _receiver,
            }
        }

        fn block(&mut self) -> &mut Block {
            &mut self.block
        }
    }

    impl Drop for DummyBlock {
        fn drop(&mut self) {
            unsafe { libc::close(self.epoll_raw_fd) };
        }
    }

    fn default_test_blockepollhandler(mem: &GuestMemory) -> (BlockEpollHandler, VirtQueue) {
        let mut dummy = DummyBlock::new(false);
        let b = dummy.block();
        let vq = VirtQueue::new(GuestAddress(0), &mem, 16);

        assert!(vq.end().0 < 0x1000);

        let queues = vec![vq.create_queue()];
        let mut disk_image = b.disk_image.take().unwrap();
        let disk_nsectors = disk_image.seek(SeekFrom::End(0)).unwrap() / SECTOR_SIZE;
        let status = Arc::new(AtomicUsize::new(0));
        let interrupt_evt = EventFd::new().unwrap();
        let queue_evt = EventFd::new().unwrap();

        let disk_image_id_str = build_device_id(&disk_image).unwrap();
        let mut disk_image_id = vec![0; VIRTIO_BLK_ID_BYTES as usize];
        let disk_image_id_bytes = disk_image_id_str.as_bytes();
        let bytes_to_copy = cmp::min(disk_image_id_bytes.len(), VIRTIO_BLK_ID_BYTES as usize);
        disk_image_id[..bytes_to_copy].clone_from_slice(&disk_image_id_bytes[..bytes_to_copy]);
        (
            BlockEpollHandler {
                queues,
                mem: mem.clone(),
                disk_image,
                disk_nsectors,
                interrupt_status: status,
                interrupt_evt,
                queue_evt,
                rate_limiter: RateLimiter::default(),
                disk_image_id,
            },
            vq,
        )
    }

    // Helper function for varying the parameters of the function activating a block device.
    fn activate_block_with_modifiers(
        b: &mut Block,
        bad_qlen: bool,
        bad_evtlen: bool,
    ) -> ActivateResult {
        let m = GuestMemory::new(&[(GuestAddress(0), 0x1000)]).unwrap();
        let ievt = EventFd::new().unwrap();
        let stat = Arc::new(AtomicUsize::new(0));

        let vq = VirtQueue::new(GuestAddress(0), &m, 16);
        let mut queues = vec![vq.create_queue()];
        let mut queue_evts = vec![EventFd::new().unwrap()];

        // Invalidate queues list to test this failure case.
        if bad_qlen {
            queues.pop();
        }

        // Invalidate queue-events list to test this failure case.
        if bad_evtlen {
            queue_evts.pop();
        }

        b.activate(m.clone(), ievt, stat, queues, queue_evts)
    }

    fn invoke_handler_for_queue_event(h: &mut BlockEpollHandler) {
        // leave at least one event here so that reading it later won't block
        h.interrupt_evt.write(1).unwrap();
        // trigger the queue event
        h.queue_evt.write(1).unwrap();
        // handle event
        h.handle_event(QUEUE_AVAIL_EVENT, EPOLLIN).unwrap();
        // validate the queue operation finished successfully
        assert_eq!(h.interrupt_evt.read().unwrap(), 2);
    }

    #[test]
    fn test_request_type() {
        let m = &GuestMemory::new(&[(GuestAddress(0), 0x1000)]).unwrap();
        let a = GuestAddress(0);

        // We write values associated with different request type at an address in memory,
        // and verify the request type is parsed correctly.

        m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_IN, a).unwrap();
        assert_eq!(request_type(m, a).unwrap(), RequestType::In);

        m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_OUT, a).unwrap();
        assert_eq!(request_type(m, a).unwrap(), RequestType::Out);

        m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_FLUSH, a).unwrap();
        assert_eq!(request_type(m, a).unwrap(), RequestType::Flush);

        m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_GET_ID, a).unwrap();
        assert_eq!(request_type(m, a).unwrap(), RequestType::GetDeviceID);

        // The value written here should be invalid.
        m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_FLUSH + 10, a)
            .unwrap();
        assert_eq!(
            request_type(m, a).unwrap(),
            RequestType::Unsupported(VIRTIO_BLK_T_FLUSH + 10)
        );

        // The provided address cannot be read, as it's outside the memory space.
        let a = GuestAddress(0x1000);
        assert!(request_type(m, a).is_err())
    }

    #[test]
    fn test_sector() {
        let m = &GuestMemory::new(&[(GuestAddress(0), 0x1000)]).unwrap();
        let a = GuestAddress(0);

        // Here we test that a sector number is parsed correctly from memory. The actual sector
        // number is expected to be found 8 bytes after the address provided as parameter to the
        // sector() function.

        m.write_obj_at_addr::<u64>(123_454_321, a.checked_add(8).unwrap())
            .unwrap();
        assert_eq!(sector(m, a).unwrap(), 123_454_321);

        // Reading from a slightly different address should not lead a correct result in this case.
        assert_ne!(sector(m, a.checked_add(1).unwrap()).unwrap(), 123_454_321);

        // The provided address is outside the valid memory range.
        assert!(sector(m, a.checked_add(0x1000).unwrap()).is_err());
    }

    #[test]
    fn test_parse() {
        let m = &GuestMemory::new(&[(GuestAddress(0), 0x10000)]).unwrap();
        let vq = VirtQueue::new(GuestAddress(0), &m, 16);

        assert!(vq.end().0 < 0x1000);

        vq.avail.ring[0].set(0);
        vq.avail.idx.set(1);

        {
            let mut q = vq.create_queue();
            // write only request type descriptor
            vq.dtable[0].set(0x1000, 0x1000, VIRTQ_DESC_F_WRITE, 1);
            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_OUT, GuestAddress(0x1000))
                .unwrap();
            m.write_obj_at_addr::<u64>(114, GuestAddress(0x1000 + 8))
                .unwrap();
            assert!(match Request::parse(&q.pop(m).unwrap(), m) {
                Err(Error::UnexpectedWriteOnlyDescriptor) => true,
                _ => false,
            });
        }

        {
            let mut q = vq.create_queue();
            // chain too short; no data_desc
            vq.dtable[0].flags.set(0);
            assert!(match Request::parse(&q.pop(m).unwrap(), m) {
                Err(Error::DescriptorChainTooShort) => true,
                _ => false,
            });
        }

        {
            let mut q = vq.create_queue();
            // chain too short; no status desc
            vq.dtable[0].flags.set(VIRTQ_DESC_F_NEXT);
            vq.dtable[1].set(0x2000, 0x1000, 0, 2);
            assert!(match Request::parse(&q.pop(m).unwrap(), m) {
                Err(Error::DescriptorChainTooShort) => true,
                _ => false,
            });
        }

        {
            let mut q = vq.create_queue();
            // write only data for OUT
            vq.dtable[1]
                .flags
                .set(VIRTQ_DESC_F_NEXT | VIRTQ_DESC_F_WRITE);
            vq.dtable[2].set(0x3000, 0, 0, 0);
            assert!(match Request::parse(&q.pop(m).unwrap(), m) {
                Err(Error::UnexpectedWriteOnlyDescriptor) => true,
                _ => false,
            });
        }

        {
            let mut q = vq.create_queue();
            // read only data for IN
            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_IN, GuestAddress(0x1000))
                .unwrap();
            vq.dtable[1].flags.set(VIRTQ_DESC_F_NEXT);
            assert!(match Request::parse(&q.pop(m).unwrap(), m) {
                Err(Error::UnexpectedReadOnlyDescriptor) => true,
                _ => false,
            });
        }

        {
            let mut q = vq.create_queue();
            // status desc not writable
            vq.dtable[1]
                .flags
                .set(VIRTQ_DESC_F_NEXT | VIRTQ_DESC_F_WRITE);
            assert!(match Request::parse(&q.pop(m).unwrap(), m) {
                Err(Error::UnexpectedReadOnlyDescriptor) => true,
                _ => false,
            });
        }

        {
            let mut q = vq.create_queue();
            // status desc too small
            vq.dtable[2].flags.set(VIRTQ_DESC_F_WRITE);
            assert!(match Request::parse(&q.pop(m).unwrap(), m) {
                Err(Error::DescriptorLengthTooSmall) => true,
                _ => false,
            });
        }

        {
            let mut q = vq.create_queue();
            // should be OK now
            vq.dtable[2].len.set(0x1000);
            let r = Request::parse(&q.pop(m).unwrap(), m).unwrap();

            assert_eq!(r.request_type, RequestType::In);
            assert_eq!(r.sector, 114);
            assert_eq!(r.data_addr, GuestAddress(0x2000));
            assert_eq!(r.data_len, 0x1000);
            assert_eq!(r.status_addr, GuestAddress(0x3000));
        }
    }

    #[test]
    #[allow(clippy::cognitive_complexity)]
    fn test_virtio_device() {
        let mut dummy = DummyBlock::new(true);
        let b = dummy.block();

        // Test `device_type()`.
        {
            assert_eq!(b.device_type(), TYPE_BLOCK);
        }

        // Test `queue_max_sizes()`.
        {
            let x = b.queue_max_sizes();
            assert_eq!(x, QUEUE_SIZES);

            // power of 2?
            for &y in x {
                assert!(y > 0 && y & (y - 1) == 0);
            }
        }

        // Test `read_config()`.
        {
            let mut num_sectors = [0u8; 4];
            b.read_config(0, &mut num_sectors);
            // size is 0x1000, so num_sectors is 8 (4096/512).
            assert_eq!([0x08, 0x00, 0x00, 0x00], num_sectors);
            let mut msw_sectors = [0u8; 4];
            b.read_config(4, &mut msw_sectors);
            // size is 0x1000, so msw_sectors is 0.
            assert_eq!([0x00, 0x00, 0x00, 0x00], msw_sectors);

            // Invalid read.
            num_sectors = [0xd, 0xe, 0xa, 0xd];
            check_metric_after_block!(
                &METRICS.block.cfg_fails,
                1,
                b.read_config(CONFIG_SPACE_SIZE as u64 + 1, &mut num_sectors)
            );
            // Validate read failed.
            assert_eq!(num_sectors, [0xd, 0xe, 0xa, 0xd]);
        }

        // Test `features()` and `ack_features()`.
        {
            let features: u64 = (1u64 << VIRTIO_BLK_F_RO)
                | (1u64 << VIRTIO_F_VERSION_1)
                | (1u64 << VIRTIO_BLK_F_FLUSH);

            assert_eq!(b.avail_features_by_page(0), features as u32);
            assert_eq!(b.avail_features_by_page(1), (features >> 32) as u32);
            for i in 2..10 {
                assert_eq!(b.avail_features_by_page(i), 0u32);
            }

            for i in 0..10 {
                b.ack_features_by_page(i, u32::MAX);
            }
            assert_eq!(b.acked_features, features);
        }

        // Test `activate()`.
        {
            // It should fail when not enough queues and/or evts are provided.
            check_metric_after_block!(
                &METRICS.block.activate_fails,
                1,
                assert!(match activate_block_with_modifiers(b, true, false) {
                    Err(ActivateError::BadActivate) => true,
                    _ => false,
                })
            );
            check_metric_after_block!(
                &METRICS.block.activate_fails,
                1,
                assert!(match activate_block_with_modifiers(b, false, true) {
                    Err(ActivateError::BadActivate) => true,
                    _ => false,
                })
            );
            check_metric_after_block!(
                &METRICS.block.activate_fails,
                1,
                assert!(match activate_block_with_modifiers(b, true, true) {
                    Err(ActivateError::BadActivate) => true,
                    _ => false,
                })
            );
            // Otherwise, it should be ok.
            assert!(activate_block_with_modifiers(b, false, false).is_ok());

            // Second activate shouldn't be ok anymore.
            check_metric_after_block!(
                &METRICS.block.activate_fails,
                1,
                assert!(match activate_block_with_modifiers(b, false, false) {
                    Err(ActivateError::BadActivate) => true,
                    _ => false,
                })
            );
        }

        // Test `write_config()`.
        {
            let new_config: [u8; 8] = [0x00, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00];
            b.write_config(0, &new_config);
            let mut new_config_read = [0u8; 8];
            b.read_config(0, &mut new_config_read);
            assert_eq!(new_config, new_config_read);
            // Invalid write.
            check_metric_after_block!(&METRICS.block.cfg_fails, 1, b.write_config(5, &new_config));
            // Make sure nothing got written.
            new_config_read = [0u8; 8];
            b.read_config(0, &mut new_config_read);
            assert_eq!(new_config, new_config_read);
        }
    }

    #[test]
    fn test_invalid_event_handler() {
        let m = GuestMemory::new(&[(GuestAddress(0), 0x10000)]).unwrap();
        let (mut h, _vq) = default_test_blockepollhandler(&m);
        let r = h.handle_event(BLOCK_EVENTS_COUNT as DeviceEventT, EPOLLIN);
        match r {
            Err(DeviceError::UnknownEvent { event, device }) => {
                assert_eq!(event, BLOCK_EVENTS_COUNT as DeviceEventT);
                assert_eq!(device, "block");
            }
            _ => panic!("invalid"),
        }
    }

    // Cannot easily test failures for
    //  * queue_evt.read
    //  * interrupt_evt.write

    #[test]
    #[allow(clippy::cognitive_complexity)]
    fn test_handler() {
        let m = GuestMemory::new(&[(GuestAddress(0), 0x10000)]).unwrap();
        let (mut h, vq) = default_test_blockepollhandler(&m);

        let blk_metadata = h.disk_image.metadata();

        for i in 0..3 {
            vq.avail.ring[i].set(i as u16);
            vq.dtable[i].set(
                (0x1000 * (i + 1)) as u64,
                0x1000,
                VIRTQ_DESC_F_NEXT,
                (i + 1) as u16,
            );
        }

        vq.dtable[1]
            .flags
            .set(VIRTQ_DESC_F_NEXT | VIRTQ_DESC_F_WRITE);
        vq.dtable[2].flags.set(VIRTQ_DESC_F_WRITE);
        vq.avail.idx.set(1);

        // dtable[1] is the data descriptor
        let data_addr = GuestAddress(vq.dtable[1].addr.get() as usize);
        // dtable[2] is the status descriptor
        let status_addr = GuestAddress(vq.dtable[2].addr.get() as usize);

        {
            // let's start with a request that does not parse
            // request won't be valid bc the first desc is write-only
            vq.dtable[0]
                .flags
                .set(VIRTQ_DESC_F_NEXT | VIRTQ_DESC_F_WRITE);
            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_IN, GuestAddress(0x1000))
                .unwrap();

            invoke_handler_for_queue_event(&mut h);

            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 0);
        }

        // now we generate some request execute failures

        {
            // reset the queue to reuse descriptors & memory
            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());

            // first desc no longer writable
            vq.dtable[0].flags.set(VIRTQ_DESC_F_NEXT);
            vq.dtable[1].flags.set(VIRTQ_DESC_F_NEXT);
            // let's generate a seek execute error caused by a very large sector number
            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_OUT, GuestAddress(0x1000))
                .unwrap();
            m.write_obj_at_addr::<u64>(0x000f_ffff_ffff, GuestAddress(0x1000 + 8))
                .unwrap();

            invoke_handler_for_queue_event(&mut h);

            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 1);
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_IOERR
            );
        }

        {
            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());

            vq.dtable[1]
                .flags
                .set(VIRTQ_DESC_F_NEXT | VIRTQ_DESC_F_WRITE);
            // set sector to a valid number but large enough that the full 0x1000 read will fail
            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_IN, GuestAddress(0x1000))
                .unwrap();
            m.write_obj_at_addr::<u64>(10, GuestAddress(0x1000 + 8))
                .unwrap();

            invoke_handler_for_queue_event(&mut h);

            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 1);
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_IOERR
            );
        }

        // test unsupported block commands
        // currently 0, 1, 4, 8 are supported

        {
            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());

            // set sector to 0
            m.write_obj_at_addr::<u64>(0, GuestAddress(0x1000 + 8))
                .unwrap();
            // ... but generate an unsupported request
            m.write_obj_at_addr::<u32>(16, GuestAddress(0x1000))
                .unwrap();

            invoke_handler_for_queue_event(&mut h);

            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 1);
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_UNSUPP
            );
        }

        // now let's write something and read it back

        {
            // write

            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());

            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_OUT, GuestAddress(0x1000))
                .unwrap();
            // make data read only, 8 bytes in len, and set the actual value to be written
            vq.dtable[1].flags.set(VIRTQ_DESC_F_NEXT);
            vq.dtable[1].len.set(8);
            m.write_obj_at_addr::<u64>(123_456_789, data_addr).unwrap();

            check_metric_after_block!(
                &METRICS.block.write_count,
                1,
                invoke_handler_for_queue_event(&mut h)
            );

            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 0);
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_OK
            );
        }

        {
            // read

            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());

            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_IN, GuestAddress(0x1000))
                .unwrap();
            vq.dtable[1]
                .flags
                .set(VIRTQ_DESC_F_NEXT | VIRTQ_DESC_F_WRITE);

            check_metric_after_block!(
                &METRICS.block.read_count,
                1,
                invoke_handler_for_queue_event(&mut h)
            );

            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, vq.dtable[1].len.get());
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_OK
            );
            assert_eq!(m.read_obj_from_addr::<u64>(data_addr).unwrap(), 123_456_789);
        }

        {
            // testing that the flush request completes successfully,
            // when a data descriptor is provided

            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());

            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_FLUSH, GuestAddress(0x1000))
                .unwrap();

            invoke_handler_for_queue_event(&mut h);
            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 0);
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_OK
            );
        }

        {
            // testing that the flush request completes successfully,
            // without a data descriptor

            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());
            vq.dtable[0].next.set(2);

            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_FLUSH, GuestAddress(0x1000))
                .unwrap();

            invoke_handler_for_queue_event(&mut h);
            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 0);
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_OK
            );

            vq.dtable[0].next.set(1);
        }

        {
            // testing that the driver receives the correct device id

            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());
            vq.dtable[1].len.set(VIRTIO_BLK_ID_BYTES);

            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_GET_ID, GuestAddress(0x1000))
                .unwrap();

            invoke_handler_for_queue_event(&mut h);
            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 0);
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_OK
            );

            assert!(blk_metadata.is_ok());
            let blk_meta = blk_metadata.unwrap();
            let expected_device_id = format!(
                "{}{}{}",
                blk_meta.st_dev(),
                blk_meta.st_rdev(),
                blk_meta.st_ino()
            );

            let mut buf = [0; VIRTIO_BLK_ID_BYTES as usize];
            assert_eq!(
                m.read_slice_at_addr(&mut buf, data_addr).unwrap(),
                VIRTIO_BLK_ID_BYTES as usize
            );
            let chars_to_trim: &[char] = &['\u{0}'];
            let received_device_id = String::from_utf8(buf.to_ascii_lowercase())
                .unwrap()
                .trim_matches(chars_to_trim)
                .to_string();
            assert_eq!(received_device_id, expected_device_id);
        }

        {
            // test that a device ID request will fail, if it fails to provide enough buffer space

            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());
            vq.dtable[1].len.set(VIRTIO_BLK_ID_BYTES - 1);

            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_GET_ID, GuestAddress(0x1000))
                .unwrap();

            invoke_handler_for_queue_event(&mut h);
            assert_eq!(vq.used.idx.get(), 1);
            assert_eq!(vq.used.ring[0].get().id, 0);
            assert_eq!(vq.used.ring[0].get().len, 1);
            assert_eq!(
                m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                VIRTIO_BLK_S_IOERR
            );
        }

        // test the bandwidth rate limiter
        {
            // create bandwidth rate limiter that allows only 80 bytes/s with bucket size of 8 bytes
            let mut rl = RateLimiter::new(8, None, 100, 0, None, 0).unwrap();
            // use up the budget
            assert!(rl.consume(8, TokenType::Bytes));

            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());
            h.set_rate_limiter(rl);

            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_OUT, GuestAddress(0x1000))
                .unwrap();
            // make data read only, 8 bytes in len, and set the actual value to be written
            vq.dtable[1].flags.set(VIRTQ_DESC_F_NEXT);
            vq.dtable[1].len.set(8);
            m.write_obj_at_addr::<u64>(123_456_789, data_addr).unwrap();

            // following write procedure should fail because of bandwidth rate limiting
            {
                // leave at least one event here so that reading it later won't block
                h.interrupt_evt.write(1).unwrap();
                // trigger the attempt to write
                h.queue_evt.write(1).unwrap();
                h.handle_event(QUEUE_AVAIL_EVENT, EPOLLIN).unwrap();

                // assert that limiter is blocked
                assert!(h.get_rate_limiter().is_blocked());
                // assert that no operation actually completed (limiter blocked it)
                assert_eq!(h.interrupt_evt.read().unwrap(), 1);
                // make sure the data is still queued for processing
                assert_eq!(vq.used.idx.get(), 0);
            }

            // wait for 100ms to give the rate-limiter timer a chance to replenish
            // wait for an extra 50ms to make sure the timerfd event makes its way from the kernel
            thread::sleep(Duration::from_millis(150));

            // following write procedure should succeed because bandwidth should now be available
            {
                // leave at least one event here so that reading it later won't block
                h.interrupt_evt.write(1).unwrap();
                h.handle_event(RATE_LIMITER_EVENT, EPOLLIN).unwrap();
                // validate the rate_limiter is no longer blocked
                assert!(!h.get_rate_limiter().is_blocked());
                // make sure the virtio queue operation completed this time
                assert_eq!(h.interrupt_evt.read().unwrap(), 2);

                // make sure the data queue advanced
                assert_eq!(vq.used.idx.get(), 1);
                assert_eq!(vq.used.ring[0].get().id, 0);
                assert_eq!(vq.used.ring[0].get().len, 0);
                assert_eq!(
                    m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                    VIRTIO_BLK_S_OK
                );
            }
        }

        // test the ops/s rate limiter
        {
            // create ops rate limiter that allows only 10 ops/s with bucket size of 1 ops
            let mut rl = RateLimiter::new(0, None, 0, 1, None, 100).unwrap();
            // use up the budget
            assert!(rl.consume(1, TokenType::Ops));

            vq.used.idx.set(0);
            h.set_queue(0, vq.create_queue());
            h.set_rate_limiter(rl);

            m.write_obj_at_addr::<u32>(VIRTIO_BLK_T_OUT, GuestAddress(0x1000))
                .unwrap();
            // make data read only, 8 bytes in len, and set the actual value to be written
            vq.dtable[1].flags.set(VIRTQ_DESC_F_NEXT);
            vq.dtable[1].len.set(8);
            m.write_obj_at_addr::<u64>(123_456_789, data_addr).unwrap();

            // following write procedure should fail because of ops rate limiting
            {
                // leave at least one event here so that reading it later won't block
                h.interrupt_evt.write(1).unwrap();
                // trigger the attempt to write
                h.queue_evt.write(1).unwrap();
                h.handle_event(QUEUE_AVAIL_EVENT, EPOLLIN).unwrap();

                // assert that limiter is blocked
                assert!(h.get_rate_limiter().is_blocked());
                // assert that no operation actually completed (limiter blocked it)
                assert_eq!(h.interrupt_evt.read().unwrap(), 1);
                // make sure the data is still queued for processing
                assert_eq!(vq.used.idx.get(), 0);
            }

            // do a second write that still fails but this time on the fast path
            {
                // leave at least one event here so that reading it later won't block
                h.interrupt_evt.write(1).unwrap();
                // trigger the attempt to write
                h.queue_evt.write(1).unwrap();
                h.handle_event(QUEUE_AVAIL_EVENT, EPOLLIN).unwrap();

                // assert that limiter is blocked
                assert!(h.get_rate_limiter().is_blocked());
                // assert that no operation actually completed (limiter blocked it)
                assert_eq!(h.interrupt_evt.read().unwrap(), 1);
                // make sure the data is still queued for processing
                assert_eq!(vq.used.idx.get(), 0);
            }

            // wait for 100ms to give the rate-limiter timer a chance to replenish
            // wait for an extra 50ms to make sure the timerfd event makes its way from the kernel
            thread::sleep(Duration::from_millis(150));

            // following write procedure should succeed because ops budget should now be available
            {
                // leave at least one event here so that reading it later won't block
                h.interrupt_evt.write(1).unwrap();
                h.handle_event(RATE_LIMITER_EVENT, EPOLLIN).unwrap();
                // validate the rate_limiter is no longer blocked
                assert!(!h.get_rate_limiter().is_blocked());
                // make sure the virtio queue operation completed this time
                assert_eq!(h.interrupt_evt.read().unwrap(), 2);

                // make sure the data queue advanced
                assert_eq!(vq.used.idx.get(), 1);
                assert_eq!(vq.used.ring[0].get().id, 0);
                assert_eq!(vq.used.ring[0].get().len, 0);
                assert_eq!(
                    m.read_obj_from_addr::<u32>(status_addr).unwrap(),
                    VIRTIO_BLK_S_OK
                );
            }
        }

        // test block device update handler
        {
            let f = NamedTempFile::new().unwrap();
            let path = f.path().to_path_buf();
            let mdata = metadata(&path).unwrap();
            let mut id = vec![0; VIRTIO_BLK_ID_BYTES as usize];
            let str_id = format!("{}{}{}", mdata.st_dev(), mdata.st_rdev(), mdata.st_ino());
            let part_id = str_id.as_bytes();
            id[..cmp::min(part_id.len(), VIRTIO_BLK_ID_BYTES as usize)].clone_from_slice(
                &part_id[..cmp::min(part_id.len(), VIRTIO_BLK_ID_BYTES as usize)],
            );

            let file = OpenOptions::new()
                .read(true)
                .write(true)
                .open(path)
                .unwrap();
            h.update_disk_image(file).unwrap();

            assert_eq!(h.disk_image.metadata().unwrap().st_ino(), mdata.st_ino());
            assert_eq!(h.disk_image_id, id);
        }
    }
}
