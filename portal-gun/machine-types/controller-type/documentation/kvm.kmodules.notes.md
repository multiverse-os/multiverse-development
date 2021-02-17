filename:       /lib/modules/5.5.0-2-amd64/kernel/arch/x86/kvm/kvm-intel.ko
name:           kvm_intel
parm:           enable_shadow_vmcs:bool
parm:           nested_early_check:bool
parm:           vpid:bool
parm:           vnmi:bool
parm:           flexpriority:bool
parm:           ept:bool
parm:           unrestricted_guest:bool
parm:           eptad:bool
parm:           emulate_invalid_guest_state:bool
parm:           fasteoi:bool
parm:           enable_apicv:bool
parm:           nested:bool
parm:           pml:bool
parm:           dump_invalid_vmcs:bool
parm:           preemption_timer:bool
parm:           ple_gap:uint
parm:           ple_window:uint
parm:           ple_window_grow:uint
parm:           ple_window_shrink:uint
parm:           ple_window_max:uint
parm:           pt_mode:int
parm:           enlightened_vmcs:bool
###############################################################################
filename:       /lib/modules/5.5.0-2-amd64/kernel/arch/x86/kvm/kvm.ko
name:           kvm
parm:           nx_huge_pages:bool
parm:           nx_huge_pages_recovery_ratio:uint
parm:           ignore_msrs:bool
parm:           report_ignored_msrs:bool
parm:           min_timer_period_us:uint
parm:           kvmclock_periodic_sync:bool
parm:           tsc_tolerance_ppm:uint
parm:           lapic_timer_advance_ns:int
parm:           vector_hashing:bool
parm:           enable_vmware_backdoor:bool
parm:           force_emulation_prefix:bool
parm:           pi_inject_timer:bint
parm:           halt_poll_ns:uint
parm:           halt_poll_ns_grow:uint
parm:           halt_poll_ns_grow_start:uint
parm:           halt_poll_ns_shrink:uint
###############################################################################
filename:       /lib/modules/5.5.0-2-amd64/kernel/drivers/vfio/vfio.ko
name:           vfio_pci
parm:           ids:Initial PCI IDs to add to the vfio driver, format is "vendor:device[:subvendor[:subdevice[:class[:class_mask]]]]" and multiple comma separated entries can be specified (string)
parm:           nointxmask:Disable support for PCI 2.3 style INTx masking.  If this resolves problems for specific devices, report lspci -vvvxxx to linux-pci@vger.kernel.org so the device can be fixed automatically via the broken_intx_masking flag. (bool)
parm:           disable_vga:Disable VGA resource access through vfio-pci (bool)
parm:           disable_idle_d3:Disable using the PCI D3 low power state for idle, unused devices (bool)
###############################################################################
filename:       /lib/modules/5.5.0-2-amd64/kernel/drivers/vfio/vfio_iommu_type1.ko
name:           vfio_iommu_type1
parm:           allow_unsafe_interrupts:Enable VFIO IOMMU support for on platforms without interrupt remapping support. (bool)
parm:           disable_hugepages:Disable VFIO IOMMU support for IOMMU hugepages. (bool)
parm:           dma_entry_limit:Maximum number of user DMA mappings per container (65535). (uint)
filename:       /lib/modules/5.5.0-2-amd64/kernel/drivers/vhost/vhost_net.ko
name:           vhost_net
parm:           experimental_zcopytx:Enable Zero Copy TX; 1 -Enable; 0 - Disable (int)
###############################################################################
filename:       /lib/modules/5.5.0-2-amd64/kernel/drivers/nvme/host/nvme.ko
name:           nvme
parm:           use_threaded_interrupts:int
parm:           use_cmb_sqes:use controller's memory buffer for I/O SQes (bool)
parm:           max_host_mem_size_mb:Maximum Host Memory Buffer (HMB) size per controller (in MiB) (uint)
parm:           sgl_threshold:Use SGLs when average request segment size is larger or equal to this size. Use 0 to disable SGLs. (uint)
parm:           io_queue_depth:set io queue depth, should >= 2
parm:           write_queues:Number of queues to use for writes. If not set, reads and writes will share a queue set. (uint)
parm:           poll_queues:Number of queues to use for polled IO. (uint)
###############################################################################
filename:       /lib/modules/5.5.0-2-amd64/kernel/drivers/nvme/host/nvme-core.ko
name:           nvme_core
parm:           multipath:turn on native support for multiple controllers per subsystem (bool)
parm:           admin_timeout:timeout in seconds for admin commands (uint)
parm:           io_timeout:timeout in seconds for I/O (uint)
parm:           shutdown_timeout:timeout in seconds for controller shutdown (byte)
parm:           max_retries:max number of retries a command may have (byte)
parm:           default_ps_max_latency_us:max power saving latency for new devices; use PM QOS to change per device (ulong)
parm:           force_apst:allow APST for newly enumerated devices even if quirked off (bool)
parm:           streams:turn on support for Streams write directives (bool)
###############################################################################
filename:       /lib/modules/5.5.0-2-amd64/kernel/net/netfilter/nf_conntrack.ko
name:           nf_conntrack
parm:           tstamp:Enable connection tracking flow timestamping. (bool)
parm:           acct:Enable connection tracking flow accounting. (bool)
parm:           nf_conntrack_helper:Enable automatic conntrack helper assignment (default 0) (bool)
parm:           expect_hashsize:uint
parm:           enable_hooks:Always enable conntrack hooks (bool)
###############################################################################
