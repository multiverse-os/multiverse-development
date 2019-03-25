##
##  Multiverse OS: Portal Gun
===========================================================
https://github.com/subgraph/paxrat
https://seattle.github.com/awnumar/memguard
# Multiverse Portal Gun

[cpuid](https://github.com/klauspost/cpuid)

https://github.com/aelsabbahy/goss
!!!!!!!!!!!!!!!!!!!
BEST SYS INFO https://github.com/zcalusic/sysinfo



##### DNS
https://github.com/hlandau/madns

## net-tools in c
https://github.com/zephyrproject-rtos/net-tools
!!!!!!!!!!!!!!
https://github.com/box-builder/box
USES FUCKING MRUBYT O CONFIG CONTAINERS (DOCKER BUT WHO GIVES A FUCK!)
## webui
https://github.com/go-http-utils/cookie

## WebKit GO Bindings 
https://github.com/abcum/webkit
**THIS HAS EASY ABILITY TO DISABLE JAVASCRIPT!**


# Golang log forwarder 

#### Anomoly detector
https://github.com/sec51/goanomaly
#### General Design
1) Move log folder out of /var/log

2) Replace /var/log a RamFS or MemFS mount so all logging goes to memory instead of the disk

3) Either read every log file as they are written and forward it to Syslog on the *Controller VM* 


# Multiverse OS Cloudius-Systems OSv

[Network]
**Phylotree like graph**
https://github.com/datastorm-open/visNetwork

[STD Render]

**WebGL**
[webgl-wireframes](https://mattdesl.github.io/webgl-wireframes/app/)


[threejs](https://github.com/spite/THREE.MeshLine)



https://github.com/teambition/compressible-go - checks if file can be compressed better (NCIE)
https://github.com/tidwall/murmur3 hash 



<nice blog: http://blog.osv.io/>

**OSv** [https://github.com/cloudius-systems/osv/wiki/Debugging-OSv]
OSv is a new open-source operating system for virtual-machines. OSv was designed from the ground up to execute a single application on top of a hypervisor, resulting in superior performance and effortless management when compared to traditional operating systems which were designed for a vast range of physical machines.


OSv has new APIs for new applications, but also runs unmodified Linux applications (most of Linux's ABI is supported) and in particular can run an unmodified JVM, and applications built on top of one.
=====
## VM And Clocks

/**
 * This clock uses uses RTC to grab wall clock time and Partition Reference Counter MSR
 * (section 12.4 from https://github.com/Microsoft/Virtualization-Documentation/raw/master/tlfs/Hypervisor%20Top%20Level%20Functional%20Specification%20v5.0b.pdf)
 * as a TSC source.
 * TODO: The MSR is simulated so the call to read its value is quite expensive. Therefore
 * eventually we should implement more efficient clock using Hyper/V Reference TSC page
 * (see section 12.6 in the same document reference above).
 */

=====
**Setup on Debian Bare Metal**
[Building OSv on Debian stable](https://github.com/cloudius-systems/osv/wiki/Building-OSv-on-Debian-stable)

-----
**Build System**
[build-system](https://github.com/cloudius-systems/osv/wiki/build-system)
Core build system

The core part of the build system is triggered by the make command and is based on a global single Makefile.

The global Makefile will detect the host system and will default the build system to be equal to the target system. To override that, one has to provide the ARCH environment variable. Currently supported values for ARCH are "x64" and "aarch64".



The kernel_base variable is set to point at the starting address of the kernel (for x64 0x200000, for AArch64 0x40080000

The final result for both images is loader.img. This is the image that contains only a ram filesystem and which is used (executed) to create the final usr.img, which is terrible of course for cross-compilation and for debugging issues (as is always the case with build systems with self-fed build artifacts).

usr.img is a different beast, and is constructed by launching loader.img and using usr.manifest to determine which modules go into the zfs filesystem. The usr.img is created starting with bare.raw, created with qemu-img create, adding the loader.img, and then using the python script imgedit.py "setpartition" command, then bare.raw is resized to its final size, and the mkfs.py script is run to create the initial ZFS filesystem based on bootfs.manifest. The mkfs.py script launches OSv to create that initial filesystem using libzfs.so inside OSv itself. Then usr.img is finally created by running upload_manifest.py, which again runs OSV to add entries to the ZFS file system based on usr.manifest.
-----
**Provision VMs**
[OSv Cloud-Init](https://github.com/cloudius-systems/osv/wiki/Cloud-init)



-----
**Network stack**
OSv has a full-featured TCP/IP network stack (on top of the network driver like virtio-net). The code was originally imported from FreeBSD, but later underwent a major overhaul to use Van Jacobson's "network channels" design, to reduce the number of locks and lock operations (even uncontended locks are slow, and contended locks are obviously worse) and the amount of slow cache-line bounces on SMP VMs.

-----

**Beyond the kernel**
**Makefile**

OSv's makefile is a complicated beast (currently undergoing a rewrite), which builds the OSv kernel, as well as a complete disk image containing the OSv kernel and an application. An application is built and specified by simple python scripts. Explaining how to build applications for OSv is beyond the scope of this mail.
**Run.py**

A small python script for running the image created by Makefile (see above) on the local host (using KVM, qemu, or Xen).
**Apps**

"apps.git" is a collection of about 60 applications known to run on OSv. The repository does not contain the applications themselves - just small scripts to fetch the applications from their own online sources, and to build them into an image using the Makefile (see above) or Capstan (see below). Some of these apps aren't quite apps, they are run-time environments on which many apps can be run. A prime example is "java" (OpenJDK 7) and openjdk8-fedora (OpenJDK 8). At some point during our development, Java support was OSv's primary focus, so our Java support is quite complete and tested.
**Capstan**

**Capstan** is an alternative to the Makefile and Run.py described above. It is similar in purpose (but not implementation) to Docker: it composes images from OSv and an application specified by a "Capstanfile" file. Capstan also allows you to upload the images you compose to a site, to download pre-composed images, and also to run these images. Using Capstan is not necessary for using OSv, but some find it more natural than the Makefile/run.py approach. I'm not one of these people ;-)
**Httpd**

Our "httpd" is a separate application (running as a thread, of course), but it is by default compiled into images created by our Makefile. It provides a "REST API" to OSv, to do all sorts of things from getting the list of threads to rebooting the machine. On top of this REST API we also have a shell (written in Lua) which runs on the host (or guest, if you prefer) and functions by sending REST (http) requests to the guest. Another thing we have on top of the REST API is a graphical administration UI.

=================================
================================================================================
## fpga uboot
https://github.com/maximeh/mkpimage
## simple cotnainer
https://github.com/converseai/simple_container


https://github.com/dop251/buse

[css only fileicons](https://github.com/picturepan2/fileicon.css)

[encoding](https://github.com/gostores/encoding)
    ASN.1
    hcl
    ini
    json
    markdown
    properties
    toml
    xmltree
    xmlsign
    yaml
## tcmu/scsi
https://github.com/coreos/go-tcmu

## Core utils
https://github.com/as/torgo

## UI
https://github.com/genshen/sshWebConsole

## Provision
https://github.com/hackwave/goiardi - CHEF IN GO

https://github.com/linuxkit/linuxkit
[Useful lines]

https://github.com/donomii/hashare 
This CAS filesystem merges duplicate files, can keep thousands of snapshots with tiny overhead, and can be distributed, backed up to cloud storage and works well with network sync services like dropbox.
## Random number generation
https://github.com/takatoh/boxmuller
# Install nvm
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.8/install.sh | bash
source ~/.bashrc

# Install LXDUI 
https://github.com/AdaptiveScale/lxdui

-----

# Terminal Emulator

## Web UI
https://github.com/tecfu/tty-table SUPPORTAS LIVE STREAMIN!

[dart]
https://github.com/updroidinc/terminal


[JS Terminal]

https://github.com/liftoff/GateOne
https://github.com/jcubic/jquery.terminal

https://github.com/YuriyNasretdinov/WebTerm + golang so nice

https://github.com/Gottox/node-webterm

http://www.erikosterberg.com/terminaljs/

https://github.com/sdgandhi/clean-terminal
https://github.com/cloudcmd/gritty
=============================
# Portal Gun
## RTOS OS
https://github.com/pdxjohnny/zephyr
http://docs.zephyrproject.org/
##
https://github.com/coreos/go-systemd
## Reverse Tunnel
https://github.com/koding/tunnel
## Circuit breaker design apttern

https://github.com/rubyist/circuitbreaker
## Reed Solomon
https://github.com/klauspost/reedsolomon
## Compress
https://github.com/klauspost/compress
## Build tools
https://github.com/walter-cd/walter - make file like commands
## SSH
https://github.com/syxolk/ssh-keycheck - list all authorized keys, last login ip, last login time, etc
## obs linux
https://github.com/dlespiau/obs
## nbd 
https://github.com/dop251/buse
https://github.com/dop251/nbd

## chord
https://github.com/armon/go-chord

# web ui
https://github.com/bluele/vermouth

## Ruby Config
https://github.com/k0kubun/itamae-go
## kobjkect
https://github.com/mdlayher/kobject Userspace events occur whenever a kobject's state changes. As an example, events are triggered whenever a USB device is added or removed from a system, or whenever a virtual network interface is added or removed.
## devices
https://github.com/bendahl/uinput

[css only fileicons](https://github.com/picturepan2/fileicon.css)

# JS Graph

http://getspringy.com/
## Path Tree
https://github.com/robfig/pathtree
## Redis Client
https://github.com/garyburd/redigo


https://github.com/minio/xfile - GETS INFO ABOUT OBJECT
## Chunked HTTP DOwnlaod
https://github.com/snikch/go-download
## FS Notify
https://github.com/rjeczalik/notify
## CHunked FS
https://github.com/EtiennePerot/splitfs
https://github.com/fsamin/go-shredder
https://github.com/advanderveer/libchunk
## cHUNKER

https://github.com/aclements/go-rabin
https://github.com/restic/chunker
https://github.com/Timechain/chunks/blob/master/chunks.go
## Backup
https://github.com/restic/restic
## Shell mux
https://github.com/georgethomas111/gomux
https://github.com/joneskoo/colorout
## Stomp (rabbitmq)
https://github.com/jjeffery/stomp
## S3 Implementation!
https://github.com/minio/mini
https://github.com/minio/sio
## ISCSI
https://github.com/kurin/tgt/tree/master/scsi
##VM
https://github.com/libvirt/libvirt-go
https://github.com/libvirt/libvirt-go-xml
## GRPC Streaming
https://github.com/itspage/grpc-streaming
## NATS streaming
https://github.com/nats-io/go-nats-streaming
## WS Multiplex
https://github.com/ckousik/wsmux
## FS Streaming
https://github.com/hayeah/gowatch
## Workers Deterministic ticking
https://github.com/VividCortex/multitick
## File Descriptor Multiplex
https://github.com/npat-efault/poller  Package poller is a file-descriptor multiplexer. It allows concurent Read and Write operations from and to multiple file-descriptors without allocating one OS thread for every blocked operation. It operates similarly to Go's netpoller (which multiplexes network connections) without requiring special support from the Go runtime. It can be used with tty devices, character devices, pipes, FIFOs, and any file-descriptor that is poll-able (can be used with select(2), epoll(7), etc.) In addition, package poller allows the user to set timeouts (deadlines) for read and write operations, and also allows for safe cancelation of blocked read and write operations; a Close from another go-routine safely cancels ongoing (blocked) read and write operations.
# SSH Multiplex
https://github.com/crosbymichael/slex
https://github.com/cosiner/socker



=================================================
## combing portal gun documents
====================================================
## Pure Go Git
https://github.com/src-d/go-git A highly extensible Git implementation in pure Go.
https://github.com/speedata/gogit Pure Go read access of a Git repository
### Visualize Crawl
[Network]
**Phylotree like graph**
https://github.com/datastorm-open/visNetwork
http://datastorm-open.github.io/visNetwork/
# Portal Gun - Development Notes
## Buffering 
bump - https://github.com/abcum/bump
## packer
https://github.com/bitgoin/packer
## DB
https://github.com/abcum/rixxdb
https://github.com/abcum/tlist - a linked time series list 
## WebKit GO Bindings 
https://github.com/abcum/webkit
## logs
[syncer](https://github.com/abcum/syncr)
A rolling append only local and remote data stream library
Can do more than just streamline logging, this is pwoerful
https://github.com/abcum/blist - binary time series list 
==================
https://github.com/nogoegst/pktconn
https://github.com/nogoegst/pickfs - virtual file system
====================
[mini-os]
https://github.com/sysml/mini-os
https://github.com/kohler/click
https://github.com/sysml/blockmon -  VHDL block 
https://github.com/sysml/HyperNF - network backend for xen
https://github.com/sysml/chaos - Chaos is a virtualization toolstack focused on performance 

https://github.com/sysml/lightvm
https://github.com/sysml/multistack MultiStack - <Kernel Support for Multiplexing and Isolating User-space Stacks.><IMPORTANT>
MultiStack is a kernel module that enables user-level network stacks to securely run alongside the in-kernel stack on the same NIC.
To isolate multiple network stacks including the in-kernel stack, a <dst ip address, dst port and protocol> 3-tuple is used. Currently, applications that run on socket APIs are isolated such that they exclusively use this 3-tuple through a call to bind() or equivalent system calls (except for special cases like fork()); MultiStack extends this primitive to user-space stacks.

### Sel4 - Verified Microkernel (Potential AppVM)
[sel4](https://github.com/seL4/seL4)
This repository contains the source code of seL4 microkernel.

[!][libsel4utils](https://github.com/seL4/libsel4utils)
 seL4 specific OS utility library -- implements threads, processes, virtual memory, elf loading etc. 

[!][camkes-vm](https://github.com/seL4/camkes-vm)
Virtual Machine build as a CAmkES component. 
	GHC and packages MissingH, data-ordlist and split (installable from cabal) For example:
		apt-get install ghc
		apt-get install cabal-install
		cabal update
		cabal install MissingH
		cabal install data-ordlist
		cabal install split
__Building__
	To build do
		make clean
		make c162_twovm_defconfig
		make silentoldconfig
		make
Then boot images/kernel-ia32-pc99 and images/capdl-loader-experimental-image-ia32-pc99 with the multiboot boot loader of your choice
For testing the C162 was configured to PXEboot (using firmware and instructions apc) pxelinux, which then used the mboot.c32 module to load the seL4 kernel and user image

_More instructions at:_ [camkes-vm](https://github.com/seL4/camkes-vm)


[c-pruner](https://github.com/seL4/pruner)

[sel4_libs](https://github.com/seL4/seL4_libs)
A collection of libraries for working on seL4.
    libsel4allocman: an allocator for managing virtual memory, malloc memory and cspaces.
    libsel4bench: a library with utilities for benchmarking on seL4.
    libsel4debug: a library with utilities for debugging on seL4. Only useful when debugging a userlevel app; potentially hacky.
    libsel4muslcsys: a library to support muslc for the root task.
    libsel4platsupport: a wrapper around libplatsupport specificially for seL4.
    libsel4simple: an interface which abstracts over the boot environment of a seL4 application.
    libsel4simple-default: an implementation of simple for the master branch of the kernel.
    libsel4simple-experimental: an implementatoin of simple for the experimental branch of the kernel.
    libsel4sync: a synchronisation library that uses notifications to construct basic locks.
    libsel4test: a very basic test infrastructure library.
    libsel4utils: a library OS - Commonly used stuff, actively maintained: implements threads, processes, elf loading, virtual memory management etc.
    libsel4vka: an allocation interface for seL4.
    libsel4vspace: a virtual memory management interface for seL4.

[libsel4vka](https://github.com/seL4/libsel4vka)
Allocation interface for seL4 -- all allocators should implement this interface. http://sel4.systems

[util-libs](https://github.com/seL4/util_libs)
Collection of OS independent utility libs:
    libcpio - a library for parsing CPIO files.
    libelf - a library for parsing ELF files.
    libethdrivers - a library for ethernet drivers.
    libpci - a library for PCI drivers.
    libplatsupport - a library of platform support utilities, interfaces for interacting with drivers, timer drivers, serial drivers and clock drivers.
    libutils - a library of generic utilities including:
        ansi.h - utilities for formatting ansi output.
        arith.h - utilities for arithmetic, ie MAX, MIN, ROUND_UP etc.
        assume.h - provides ASSUME, which allows the user to provide hints to gcc.
        builtin.h - defines conventient macros for using builtin gcc attributes.
        compile_time.h - provides compile time asserts.
        debug.h - various debugging macros.
        formats.h - formats for printf.
        list.h - a basic, void * pointer based list implementation.
        math.h - provies complex math, ie. muldivu64.
        page.h - provides virtual memory page operations.
        sglib.h - an open source template library that provides arrays, lists, red-black trees etc.
        stringify.h - provides macros for creating even more macros.
        time.h - provides temporal constants (i.e US_IN_S)
        util.h - includes all util header files.
        verification.h - macros for verification in Isabelle.
        zf_log_config.h - provides zf_log config.
        zf_log.h - an open source logging library.




https://github.com/seL4/camkes-tool 
The main CAmkES tool https://wiki.sel4.systems/CAmkES




## VisJS
http://visjs.org/timeline_examples.html
## D3 UI
https://github.com/vlandham/d3-radial


### Provision <!>
**Modify Configs**
[put](https://github.com/n3phtys/put)
Go tool to ensure a given line is inside a text file. If it is not, the tool will insert the line (configurable in which line number). 


=============================================
### Filesystem
[Rust][tfs](https://github.com/redox-os/tfs)
TFS is a modular, fast, and feature rich next-gen file system, employing modern techniques for high performance, high space efficiency, and high scalability.
TFS was created out of the need for a modern file system for Redox OS, as a replacement for ZFS, which proved to be slow to implement because of its monolithic design.
TFS is inspired by the ideas behind ZFS, but at the same time it aims to be modular and easier to implement.
TFS is not related to the file system of the same name by terminalcloud.





============================================
[CPU]
https://github.com/templexxx/cpufeat
[cpuid](https://github.com/klauspost/cpuid)
!!!!!!!!!!!!!!
https://github.com/box-builder/box
USES FUCKING MRUBYT O CONFIG CONTAINERS (DOCKER BUT WHO GIVES A FUCK!)


https://github.com/linuxkit/linuxkit

Configuration from YAML/JSON/TOML/XML.. and from one of these files served over HTTP.

Needs to be able to load multiple baremetal servers into configuration, remote and local lan


A TAP device is a virtual ethernet adapter, while a TUN device is a virtual point-to-point IP link.
## Hidden COntainer Projects

https://github.com/sandia-minimega/minimega
----
### AutoTLS Gin
https://github.com/gin-gonic/autotls
### Localization / Internationalization / i18n
CLDR project must be paired with actual i18n translation software, this provides other details, similar to what Rails would provide.
[locales](https://github.com/go-playground/locales)
Locales is a set of locales generated from the Unicode CLDR Project which can be used independently or within an i18n package; these were built for use with, but not exclusive to, Universal Translator.

----
https://github.com/donomii/hashare 
This CAS filesystem merges duplicate files, can keep thousands of snapshots with tiny overhead, and can be distributed, backed up to cloud storage and works well with network sync services like dropbox.
## Random number generation
https://github.com/takatoh/boxmuller
### New IP command
I want a new ip command that does everything ip does but also some additional features

https://github.com/picatz/iface

### Relevant Libs
[color]
https://github.com/logrusorgru/aurora
[net uitls]
https://github.com/ArroyoNetworks/splice
A high-level and multi-os library for manipulating network interfaces, links, and routes.
https://github.com/songgao/water
A simple TUN/TAP library written in native Go (golang).



[netlink]
https://github.com/mickep76/netlink/
[internet balacning]
https://github.com/danielmorandini/booster-network

[p2p lib]
https://github.com/libp2p/go-libp2p-net
[cache]

https://github.com/patrickmn/go-cache
[charts]
https://github.com/wcharczuk/go-chart

[CLI Framework/Tools]
https://github.com/urfave/cli
https://github.com/briandowns/spinner
bash coimplete https://github.com/posener/complete
----
### Provisioning Frameworks
https://github.com/hyperblock/lvdiff
https://github.com/gocircuit/circuit
https://github.com/dnaeon/gru go based with lua confgiruation
https://github.com/fabric8io/kansible ansible in go

https://github.com/ctdk/goiardi chef in go


https://github.com/rdeusser/overseer - one of the nicer looking ones

### Libvirt Libs
https://github.com/fog/fog - fog compatibility?


https://github.com/ohadlevy/virt - simple easy to use ruby interface

https://github.com/libvirt/libvirt-go-xml - xml 
https://github.com/digitalocean/go-libvirt - pure go
https://github.com/libvirt/libvirt-go - lib virt bindings
  - lib virt REST
   https://github.com/JordanDeBeer/kvm_interface [pssioble sample code]
   https://github.com/ByteArena/schnapps [sample code]



[clear containers, lots of great sample code]
https://github.com/hyperhq/runv/tree/master/hypervisor/libvirt
### Useful Code
[ability to check for existing module in the kernel]
var hasVsock bool
	_, err = exec.Command("/sbin/modprobe", "vhost_vsock").Output()
	if err == nil {
		hasVsock = true
}

### Possible Functionality

* Provide a map of the network like this: https://github.com/qb0C80aE/pottery

* Submit a code and language over REST API and spin up a VM to run the code and give the response

----
## Configuring LXD
Ater compiling from source, start one instance of lxd with:

		lxd --verbose

Then after that is started the command is ran a second time in a seperate terminal to initiate the configuration process:

		lxd init



---

		sudo groupadd lxd
		sudo usermod --append --groups lxd user


   









===========================================================










## Libinput
https://github.com/gvalkov/golang-evdev




=================================
====================================
## FILESYSTEM
# Multiverse OS Distributed FS
Multiverse OS requires a distributed filesystem that supports specific set of features:

**Features**
* Content Addressable Chunking
* Distributed and decentralized
* Versioning
* Compression
* Encryption
* Caching
* Streaming/Downloading from multiple peers at a time (torrent-like)
* File integrity (sha3 hash or something better)
* Built in HTTP sharing (right click add to share)
* Passive Background Hooks
  * Integrity Checking
  * Virus Checking
  * Sensitive information scrubbing (https://github.com/playnet-public/fscrub)

[Example Projects]
https://github.com/dchest/hesfic
https://github.com/lateefj/shylock - very nice
https://github.com/monax/hoard

https://github.com/as/fs/blob/master/fs.go

https://github.com/google/fscrypt 
https://github.com/twitchyliquid64/nugget
[Similar Projects]
Camilstore

----
https://github.com/h2non/filetype

https://github.com/tyler-smith/go-bip32 - KEY SYSTEM

https://github.com/multiformats/go-multiaddr
## Splitting/Chunking
https://github.com/aclements/go-rabin <special rolling hash>
https://github.com/EtiennePerot/splitfs
https://github.com/restic/chunker content defined chunking
## Content addressable
https://github.com/technoweenie/go-contentaddressable
----
[k/v]
https://github.com/bazil/bolt-mount
[specialfs]
https://github.com/samthor/valuefs
[torrentfs]
https://github.com/ring00/torrentfs
[Qcow2]
https://github.com/vasi/qcow2
[Ram/Memory Filesystems]
https://github.com/bbengfort/memfs

https://github.com/zbiljic/memfs

https://github.com/taruti/ramfs
https://github.com/nak3/memdbfs
https://github.com/mars9/ramfs - 9p but all in memory!

[REST API Filesystem Interaction]
https://github.com/jaipradeesh/summer
[VersionFS]
https://github.com/FooSoft/vfs
[HTTPFS]
https://github.com/prologic/httpfs <!>
https://github.com/sdgoij/ghttpfs
https://github.com/jaipradeesh/summer
https://github.com/ungerik/go-fs
https://github.com/Ackar/CrawlerFS
https://github.com/jban332/kin-fs
https://github.com/kimiazhu/vfs/blob/master/httpfs/httpfs.go
[Virtual FS] (Limit operaations of software to an abstracted FS folder)
https://github.com/blang/vfs
https://github.com/jaipradeesh/summer - very nice


[Utils]
https://github.com/ik5/fsutils

https://github.com/frozzare/go-fs 

https://github.com/deniswernert/go-fstab
https://github.com/andreaskoch/go-fswatch

