##
##  Multiverse OS: Core 
===========================================================

The number of relevant libraries is massive, so to limit it to what I need to build a basic version of Multiverse I'm isolating this file to just the core libraries needed for basic development.


# Multiverse Programming Notes - Data Containers
New data types are being created and introduced regularly to statisfy the evolving needs of scientists and developers. This article series will select a few data types and review them, introducing you to the data type, providing historical context, examples of use, refences in academic papers and so on. 




**Suffix Array**
An array that compares data of each item starting from the end. 

_Example:_ [rusty-suffix](https://github.com/Jarusk/rusty-suffix) - Implementations of various suffix search data structures and algorithms

**Interval Heap**

https://github.com/tyler-smith/go-bip32 - KEY SYSTEM


**Standard Library: Socks5 Proxy**
https://github.com/golang/net/blob/master/proxy/socks5.go
#### DHT
#### Review
https://github.com/Safing/safing-core
https://github.com/google/fscrypt
## WebKit GO Bindings 
https://github.com/abcum/webkit
**THIS HAS EASY ABILITY TO DISABLE JAVASCRIPT!**
=====

https://github.com/skydive-project/skydive - network analysis
#### DB 
[vtree](https://github.com/abcum/vtree)
versioned radix tree
[rtree](https://github.com/abcum/rixxdb)
versioned, strongly consistent, ACID, transaction with rollback

#### low level proxy
[syscall proxy](https://github.com/vizee/tcpproxy)

[syscall-udp](https://github.com/ld86/syscall-udp)



#### Kernel Modules / PCI Devices

[!][C Langauge]
[kernel-qemu-pci](https://github.com/levex/kernel-qemu-pci)
A custom PCI device in QEMU and a module for using it 


#### CLI / Cmd
The core functionality will need to be built around a few core command line programs:

[kvmtop](https://github.com/cha87de/kvmtop)

**Portal Gun**

#### Console
[console](https://github.com/containerd/console)

[sh](https://github.com/mvdan/sh)
shell parser, formatter, and interpretter (POSIX/Bash/mksh)


#### uinput
[uinput](https://github.com/bendahl/uinput)
This package provides pure go wrapper functions for the LINUX uinput device, which allows to create virtual input devices in userspace

#### Processes
[go-procsnitch](https://github.com/subgraph/go-procsnitch)
go proc sockets lib

#### Provisioning / Orchestrating / Deploying

[goss](https://github.com/aelsabbahy/goss)
way of testing and validating servers

[silentinstall](https://github.com/alistanis/silentinstall)
define answers to shell prompts during installs with JSON to install silently

**Compose**
[libcompose](https://github.com/docker/libcompose)
`Compose` is the config format used by docker and others. Supporting this would allow
a lot of stuff to be easily ported to Multiverse OS.
*Look at the examples of other projects that used this library for reference.*

[keybd_event](https://github.com/micmonay/keybd_event)
simualte key event, may be useful for provisioning

#### Booting

[u-root](https://github.com/u-root/u-root)
universal root initramfs containing busybox like features

[dock](https://github.com/robinmonjo/dock)
micro init system for linux containers

[ginit](https://github.com/mesanine/ginit)

[dainit](https://github.com/driusan/dainit)


#### Machine (VM, Container, System, Linux Management)
[machine](https://github.com/docker/machine)
A lot of useful code, from SSH, etc
  [ssh](https://github.com/docker/machine/tree/master/libmachine/ssh)

[libaudit](https://github.com/mozilla/libaudit-go)
Interface with netlink and other linux stuff through the audit api provided by the linux kernel

**cgroups**
[cgroups](https://github.com/containerd/cgroups)
Cgroups and namespaces are the foundation of containers
/sys/fs/cgroup/cpu/test

**Containers**
[clearcontainer: hyperhq/runv](https://github.com/hyperhq/runv)
runV is a hypervisor-based runtime for OCI.

[runc](https://github.com/opencontainers/runc)
CLI tool for spawning and running containers according to the OCI specification
  [go-runc](https://github.com/containerd/go-runc)

[rancher:os](https://github.com/rancher/os)


[containerd](https://github.com/containerd/containerd)
An open and reliable container runtime 

[clearcontainers/runtime](https://github.com/clearcontainers/runtime)
collection of runtime tools, for example: kill, start, run, utils, logger, ...
  [hyperd](https://github.com/hyperhq/hyperd)
  This repo contains two parts: the daemon of HyperContainer hyperd and the CLIhyperctl.
  [hypercli](https://github.com/hyperhq/hypercli)
  Go version of Hyper.sh client command line tools.

[rkt](https://github.com/rkt/rkt)

[garden-linux](https://github.com/cloudfoundry-attic/garden-linux)


**Agents** (Ran on the VM/Container to report statistics and run commands)
[clearcontainers/agent](https://github.com/clearcontainers/agent)

[ustat](https://github.com/penberg/ustat)
unviersal stat tool that combines several tools into 1 with low overhead

[terminus](https://github.com/kelseyhightower/terminus)
very good stats program



#### Devices
[fifo](https://github.com/containerd/fifo)

[usblockout](https://github.com/subgraph/usblockout)
monitors user session and triggers grsecurity deny new usb feature

#### Filesystem / File 
[continuity](https://github.com/containerd/continuity)
A transport-agnostic, filesystem metadata manifest system

[buse-go](https://github.com/samalba/buse-go)
block device in userspace. Supporting NBD so VERY <important>


#### Humanizing
[go-units](https://github.com/docker/go-units)
Parse and print size and time units in human-readable format


#### Networking
[alg](https://github.com/mdlayher/alg)
linux AF_ALG sockets fort communciation with the kernel

[goscan](https://github.com/timest/goscan)
fast networking scan

[raw](https://github.com/mdlayher/raw)
raw device

[ether](https://github.com/songgao/ether)

[GOnetstat](https://github.com/elsonwu/GOnetstat)
data from /proc/net/tcp|6 and /proc/net/udp|6

[rnetlink](https://github.com/jsimonetti/rtnetlink)
 Package rtnetlink provides low-level access to the Linux rtnetlink API. MIT Licensed. 

[genetlink](https://github.com/mdlayher/genetlink)
genetlink implements generic netlink interactions and datatypes

[go-tuntap](https://github.com/lab11/go-tuntap)

[netlink](https://github.com/mdlayher/netlink)

[tenus](https://github.com/milosgajdos83/tenus)
random networking in linux

[zsocket](https://github.com/nathanjsweet/zsocket)
zero-copy sockets for linux

[vsock](https://github.com/mdlayher/vsock)
package vsock provides linux VM sockets. vsock provides access to Linux VM sockets (AF_VSOCK) for communication between a hypervisor and its virtual machines.

[go-nfnetlink](https://github.com/subgraph/go-nfnetlink)
library for communicating with netfilter subsystems over netlink sockets

[tcp-shaker](https://github.com/tevino/tcp-shaker)
customized tcp for handshakes that use only parts of the tcp handshake

[go-tproxy](https://github.com/LiamHaworth/go-tproxy)
linux transparent proxying using low level sockets and iptables

[go-connections](https://github.com/docker/go-connections)
Supports memsockets, unix socket, tcp socket, and so on. Referencing this will save a lot of time

[libnetwork](https://github.com/docker/libnetwork)
General networking for VMs (containers), and will likely be the core to the networking
  [Firewall/Iptables]
	[libnetwork:iptables](https://github.com/docker/libnetwork/blob/master/iptables/iptables.go)
  [/etc/hosts]
  [libnetwork:etchosts](https://github.com/docker/libnetwork/blob/master/etchosts/etchosts.go)

**DNS**
[dnsserver](https://github.com/docker/dnsserver)
A small DNS service for communicating A and SRV records. 

**Software Defined Networking**
[ovs](https://github.com/digitalocean/go-openvswitch/tree/master/ovs)
Open switch

**Message Queue**
[libchan](https://github.com/docker/libchan)
Libchan is an ultra-lightweight networking library which lets network services communicate in the same way that goroutines communicate using channels:
  * Simple message passing
  * Synchronization for concurrent programming
  * Nesting: channels can send channels
Libchan supports the following transports out of the box:
  * In-memory Go channel
  * Unix socket
  * Raw TCP
  * TLS
  * HTTP2/SPDY
  * Websocket

[spdystream](https://github.com/docker/spdystream)
*Multiplexed* spdy stream library

[macouflage-multi](https://github.com/subgraph/macouflage-multi)

[macouflage](https://github.com/subgraph/macouflage)

#### Procfs
[sysinfo](https://github.com/zcalusic/sysinfo)
Package sysinfo is a pure Go library providing Linux OS / kernel / hardware system information. It's completely standalone, has no dependencies on the host system, doesn't execute external programs, doesn't even import other Go libraries.


#### Security
[go-seecomp](https://github.com/subgraph/go-seccomp)
Go support for parsing, compiling, and installing Chromium OS Seccomp-BPF policy files. 

#### Networked Block Devices
Network block devices that can be used to share devices across VMs

[go-p9p](https://github.com/docker/go-p9p)
A modern, performant 9P library for Go. *Plan 9* is the default way


#### Kernel
[keyutils](https://github.com/jandre/keyutils)

[conntrack](https://github.com/typetypetype/conntrack)


#### Password Store
**2fa**
[2fa](https://github.com/rsc/2fa)

[pick](https://github.com/bndw/pick)
fancy pure go pass manager


#### X / Windows Manager
[xgb](https://github.com/BurntSushi/xgb)
Foundation for a lot of tools but interfaces with X server
  [xgb-util](https://github.com/BurntSushi/xgbutil)

[go-xdgdirs](https://github.com/subgraph/go-xdgdirs)
Golang library for reading and parsing XDG User Dirs 

[x11-client](https://github.com/electricface/go-x11-client)

## Linux

**Jails**
[isowrap](https://github.com/xmc-dev/isowrap)
Isowrap is a library used to execute programs isolated from the rest of the system.


## Input
[golang-evdev](https://github.com/gvalkov/golang-evdev)

#### Management 
[overlord](https://github.com/aitjcize/Overlord)
https://github.com/ppacher/attackdb

========================
##    __  __       _ _   _                            ___  ____                                     ##
##   |  \/  |_   _| | |_(_|_   _____ _ __ ___  ___   / _ \/ ___|                                    ##
##   | |\/| | | | | | __| \ \ / / _ | '__/ __|/ _ \ | | | \___ \                                    ##
##   | |  | | |_| | | |_| |\ V |  __| |  \__ |  __/ | |_| |___) |                                   ##
##   |_|  |_|\__,_|_|\__|_| \_/ \___|_|  |___/\___|  \___/|____/                                    ##
##                                                                                                  ##
!!   A collection of notes and resources related development of each Multiverse OS component        !!
##                                                                                                  ##

[cpuid](https://github.com/klauspost/cpuid)
## chord
https://github.com/armon/go-chord
## Use Ruby to config
https://github.com/box-builder/box
## Ruby Config
https://github.com/k0kubun/itamae-go
## Multiverse (NOW)
kernel
pci
syscall
chroot
systemd
grub
system call
**security**
secure random
wipe memory
safe
## Hackwave Chat
extract (urls)
## Multiverse (VN Screen Sharing)
Xine
KVM ScreenRelayer (LookingGlass)
desktop
screen
capture
## Multiverse Virtual Machines (VMs)
containers
lxc
lxz
runv (clear container)
kvm
qemu
## Multiverse Console
libinput
## Multiverse 
menomic
ruby
compil(er)
boot
init
protocol
## Multiverse Routing
rout(er)
switch
hub
vxlan(vlan)
veth
gateway
zero(copy)
## Multiverse Distributed FS
fs(filesystem)
## Multiverse Media Management
fs(filesystem)
## Provisioning / Devops (Chef, Ansible)
chef
ansible
## Go Build Tools (Compiling/Building/Releasing)
assembly(asm)
binary
interpreter
hex
reverse
mock
## Multiverse (Desktop Environment)
gnome
lock
## Multiverse Network
frames
ether(net)
compression
xz
snappy
## Multiverse UI (Hosting / WebUI)
dashboard
x11
xdg
xserver
wayland
## Multiverse (VM Monitoring/Stats/Health)
monitor
stats
metrics
health
## Multiverse (Future Development)
vpn
pgp(gpg)
## Hackwave
experiments
protocol(s)
## Zerg
xss
scan
CVE
vul(nerability)
crack
brute
dict(ionary)
###################################################################################
## Terms to search
sys
hardware
init
bios
ramfs
syslog
syscall
partition
volume
lvm
luks
stages
rust
diagram
gram
workflow / pipeline
pure native
image img
gosh
physics
avatar
uroot
uboot
uefi
dot
pm package manager
mouse pointer
window
grpc
modem
fax
sharding
graph
json db
ram db
mem db
yaml toml json xml db
specialized IP addrress RADIX TREE for IP lookup!!!!!
^ use in combination with tool that controls HOST FILE **WITHOUT CHANGING THE TEXT FILE, SO THE RADIX TREE INSTEAD OF FILE BASED LOOKUP CAN WORK FASTER!!!!**
upload
brute scan dictionary sql smtp scan pentest
version db
camilstore
magento ecommerce catalog inventory product
sniffer
maildir
file database 
media database
video databsae
music database
comic
comic database
video
inventory bag slots
store
tex scientific articles journal publishing pdf latex  epub
benchmark
crawling scraping data mining wikidata chem sparkgraph 
======================================================================================================

[tiedot](https://github.com/HouzuoGuo/tiedot)
nosql database, embedded

DHT
https://github.com/matrix-org/dendrite 
-----

[syscall version of cat in go](https://github.com/denderello/cat/blob/master/main.go)

======================================================================================================
## Unorganized Projects
======================================================================================================
## GUI Prompts
[ask: ui prompts](https://github.com/kyoh86/ask)


## Linux User System
[libnsss-stns](https://github.com/STNS/libnss_stns)
stns is linux name resolver and user authentication library

## Memory
[memory manager](https://github.com/runningwild/memory)
Simple Memory Manager for Go 

[memfd](https://github.com/justincormack/go-memfd)

memory file descriptors
## SCREEN SHARING!!!
[xcapture]([!][https://github.com/dominikh/xcapture]
Xcapture is a command-line driven X11 window recorder, outputting a raw video stream on standard out for processing by other tools.)
### DATA ANALYSIS PIPELINE
[godap](https://github.com/rapid7/godap)
DAP was created to transform text-based data on the command-line, specializing in transforms that are annoying or difficult to do with existing tools.

DAP reads data using an input plugin, transforms it through a series of filters, and prints it out again using an output plugin. Every record is treated as a document (aka: hash/dict) and filters are used to reduce, expand, and transform these documents as they pass through. Think of DAP as a mashup between sed, awk, grep, csvtool, and jq, with map/reduce capabilities.
## childpress
https://github.com/enkessler/childprocess RUBVY
## SSE
[sse][https://github.com/JanBerktold/sse]


## Benchmarking / Profiling
[gom](https://github.com/rakyll/gom)
A visual interface to work with runtime profiling data for Go, provides a simple and easy to use WebUI


## Modem
	[go-phone][https://github.com/Grayda/go-phone]
	 Allows you to use a (USB) 56K modem to detect incoming phone calls 
## Physics
[gravity][https://github.com/hawkgs/go-gravity]
The library has a set of forces (vectors) like wind, gravity, kinetic friction, etc.
## Art
[planetary][https://github.com/szll/planetry]
Planetry is a simple app for gravitational simulations of objects in space. You can specify the objects or bodys by your own. A scripting system allows you to check for specific events.
======================================================================================================
##== Multiverse Keywords & Naming ==================================================================##
======================================================================================================
Collection of scientific words that can be used 

Gravity
Time-Dialation 


======================================================================================================
##== Operating System Notes & Resear ===============================================================##
======================================================================================================
**Wayland**
There are a few strategies to determine if wayland is being used, some provide misleading output so
I want to review the different ways to learn more about the desktop environment, operating systems, 
and the standard utilities that come with Debian (buster).

One strategy is to review the environmental variables set. Environmental variables are stored in the
shell, loaded primarily after logging in and provide a lot of useful information.

To output all the environmental variables set wihin the curent environment simply use the `env`
command with no other arguments, options or flags.

			==[DEFINITION:xdg]===========================================================
			=                                                                           =
	  		xdg is an acronym stands for 'x desktop group', but the entire group bas
	      since changed their name and project to freedesktop.org
			=                                                                           =
			=============================================================================

------------------------------------------------------------------------------------------------------
======================================================================================================
##== Useful Code ===================================================================================##
https://raw.githubusercontent.com/coreos/ignition/master/internal/distro/distro.go
Various constants for linux systems

[users / os]
https://github.com/coreos/ignition/blob/master/internal/exec/util/passwd.go
create user / manage passwd file

https://github.com/coreos/ignition/blob/master/internal/authorized_keys_d/authorized_keys_d.go
sweet authorized keysa nd ssh fodler management

[building filesystems]
https://github.com/coreos/ignition/blob/master/internal/sgdisk/sgdisk.go
disk partition management

https://github.com/coreos/ignition/blob/master/internal/exec/stages/disks/disks.go
build raids/etc

https://github.com/coreos/ignition/blob/master/internal/exec/stages/files/files.go
create files and users

https://github.com/coreos/ignition/blob/master/internal/systemd/systemd.go
// WaitOnDevices waits for the devices named in devs to be plugged before returning.

https://github.com/coreos/ignition/blob/master/internal/resource/url.go
fetch from http/ftp


#### QUESTION & ANSWSER ##############################################################################

#==[ Question 1 ]====================================================================================#
#  What defines an operating system?                                                                 #
#                                                                                                    #
#----------------------------------------------------------------------------------------------------#

 _A._ From the perspective of a program the core functionality of an operating system is the system
calls made available to the program.

Operating system runs programs isolated as processes. It handles alternation between tasks, using
ineficient _polling_ technique instead of the more modern real-time interval based method.

[TODO: Provide more information]

#==[ Question 2 ]====================================================================================#
#  Are global variables bad in Go?                                                                   #
#                                                                                                    #
#----------------------------------------------------------------------------------------------------#

_A._ Yes!

#==[ Question 3 ]====================================================================================#
#  What syscall is used to allocate memory to your process and what is it allocating?                #
#                                                                                                    #
#----------------------------------------------------------------------------------------------------#



_A._ The syscall used to allocate memory is mmap, to unallociate memory munmap syscall is used.
Memory is allocated in pages of a fixed length (larger amounts can be allocated using 


######################################################################################################
 | |   (_| |__  _ __ __ _ _ __ _   _  | \ | | ___ | |_ ___ ___ 
 | |   | | '_ \| '__/ _` | '__| | | | |  \| |/ _ \| __/ _ / __|
 | |___| | |_) | | | (_| | |  | |_| | | |\  | (_) | ||  __\__ \
 |_____|_|_.__/|_|  \__,_|_|   \__, | |_| \_|\___/ \__\___|___/
=======================================================================================================
--------------------------------\__/-------------------------------------------------------------------
##== Basics =========================================================================================##
#### Linux / Debian / Systemd
  **Tor**
  <IMPORTANT>[oniongen-go][https://github.com/lostinblue/oniongen-go]
	custom built onion generation for key generation for account system

  **OS**
		[General Linux]
		  [os extensions][https://github.com/containerd/cri-containerd/blob/master/pkg/os/os.go]
	 		[sysinfo][https://github.com/zcalusic/sysinfo]
			<important>MEGA rich sysinfo, will be critical for install	
    	[sys][https://github.com/golang/sys/tree/master/unix]
    	golang base sys package, useful for reference
		[Debian]
	  	[update-notifier][https://github.com/zcalusic/update-notifier]
	  	notifies updates in icon tray when updates are available
  **windows, drawing, events, pointers, wayland**
		[go.wde][https://github.com/skelterjohn/go.wde]
  **dbus**
		[go.dbus][https://github.com/guelfey/go.dbus]
	 		[fork][https://github.com/aulanov/go.dbus]
	  	Complete native implementation of the D-Bus message protocol
	  	Go-like API (channels for signals / asynchronous method calls, Goroutine-safe connections)
			Subpackages that help with the introspection / property interfaces
  **systemD**
		[go-systemd][https://github.com/coreos/go-systemd]
		Go bindings to systemd. The project has several packages
  **Security**
		[seccomp][https://github.com/kubernetes-incubator/cri-o/tree/master/server/seccomp]
		[apparmor][https://github.com/kubernetes-incubator/cri-o/tree/master/server/apparmor]

  **Unix Sockets**
	<important>[floodgate][https://github.com/ericychoi/floodgate]
	a test proxy for unix sockets, can attack one unix socket to another (socket is just a file!)

  **Devices**
	[fifo][https://github.com/containerd/fifo]
	go fifo pkg

	[zsocket][https://github.com/nathanjsweet/zsocket]

	[network code from kata][https://github.com/kata-containers/agent/blob/master/network.go]

  **Signals & Exiting**
	[sipid][https://github.com/cloudfoundry/sipid]
	Pid management, killing, taking over pid, etc

	[goodbye][https://github.com/thecodeteam/goodbye]
	Goodbye is a Golang library that provides a standard way to execute code when a process
	exits normally or due to a signal. The Goodbye library uses a sync.Once to ensure that the
	registered exit handlers are executed only once -- whether as a result of the process exiting
	normally or as a result of a received signal.

### Database
  **Embedded Key/Value**
		[unfinished leveldb server][https://github.com/golang/leveldb]
		  [fork][https://github.com/rayyildiz/leveldb]
		super awesome leveldb implementation, but unfinished, but by google

### Compression
  **Stream Compression**
		[snappy][https://github.com/golang/snappy]

### Validation
	[valdiate][https://github.com/markbates/validate]

### Humanize 
	[go-humanize][https://github.com/dustin/go-humanize]
	[inflect][https://github.com/martinusso/inflect]

### Bin & Binary Data Storage & Nested Bin Execution
	[go-bindata][https://github.com/gnoso/go-bindata]
  (NEWEST COMMTIS)[go-bindata](https://github.com/jteeuwen/go-bindata)

	[bin2hex][https://github.com/JamesHovious/bin2hex]

####------------------------------------------------------------------------------------------------------------------------------------- Translations / i18n / Locales
	[gt][https://github.com/melvinmt/gt]
	tiny but very powerful internationalization
	[i18n][https://github.com/qor/i18n]
	I18n provides internationalization support for your application, it supports 2 kinds of storages
	(backends), the database and file system.

#### Go Utilities / Tools / Code
	[x][https://github.com/bbengfort/x]
	[mergo][https://github.com/imdario/mergo]
	A helper to merge structs and maps in Golang. Useful for configuration default values, avoiding
	messy if-statements.

		src := Foo{
			A: "one",
			B: 2,
		}
		dest := Foo{
			A: "two",
		}
		mergo.Merge(&dest, src)

  **Date & Time Libraries**
  [when](https://github.com/olebedev/when)
   nice and humanr eadable

  [dateparse](https://github.com/araddon/dateparse)
  both date and timezones parsing

	[now][https://github.com/jinzhu/now]
		time.Now() // 2013-11-18 17:51:49.123456789 Mon
		now.BeginningOfMinute()   // 2013-11-18 17:51:00 Mon
		now.BeginningOfHour()     // 2013-11-18 17:00:00 Mon
		now.BeginningOfDay()      // 2013-11-18 00:00:00 Mon
		now.BeginningOfWeek()     // 2013-11-17 00:00:00 Sun
		...

#### Go Development
  **Go Application Logging**
	[zap][https://github.com/uber-go/zap]

  **Configuration**
	[configur][https://github.com/jinzhu/configor]
	YAML, JSON, TOML, env
	[toml][https://github.com/BurntSushi/toml]
	 

#### Design patterns
  **General**
	[go-patterns][https://github.com/tmrts/go-patterns]
	A curated collection of idiomatic design & application patterns for Go language.

  **Concurrency**
	[go-promise][https://github.com/fanliao/go-promise]
	[atomic][https://github.com/uber-go/atomic]

  **Managing Go Routines / Pools / Lifecycles / Context**
	[run][https://github.com/oklog/run]
	Provides very simple run groups, to manage go routines and their lifecycles

  **Message Queue/Pubsub**
	[emission][https://github.com/chuckpreslar/emission]

#### Data Types
  **General**
	[buffer][https://github.com/djherbis/buffer]
	composable customizable buffers in go
		// file based buffer
		// Create a File-based Buffer with max size 100MB
		file, err := ioutil.TempFile("", "buffer")
		buf := buffer.NewFile(100*1024*1024, file)
		// ring buffer
		// Create a File-based Buffer with max size 100MB
		file := buffer.NewFile(100*1024*1024, someFileObj) // you'll need to Open(), Close() and Delete someFileObj.
		// If buffered data exceeds 100MB, overwrite oldest data as new data comes in
		buf := buffer.NewRing(file) // requires BufferAt interface.
		// spill buffer	
		// Buffer 32KB to Memory, discard overflow
		buf := buffer.NewSpill(32*1024, ioutil.Discard)

	[bitset][https://github.com/tomcraven/bitset]
	A simple bitset implementation in Go with basic functions for manipulation

	[computer-science][https://github.com/dansackett/computer-science]
	A large collection of sample algorithms for a wide variety of known
	solutions in Go.

	[gotomic][https://github.com/zond/gotomic]
	nonblocking map/hash datastructures

  **State Machines**
	[fsm][https://github.com/looplab/fsm]

	[state-machine][https://github.com/felixalias/state_machine]

	[transition][https://github.com/qor/transition]
	looks very good

======================================================================================================
  / ___|___  _ __ ___  _ __   ___  _ __   ___ _ __ | |_ ___ 
 | |   / _ \| '_ ` _ \| '_ \ / _ \| '_ \ / _ | '_ \| __/ __|
 | |__| (_) | | | | | | |_) | (_) | | | |  __| | | | |_\__ \
  \____\___/|_| |_| |_| .__/ \___/|_| |_|\___|_| |_|\__|___/
=======================================================================================================
----------------------|_|------------------------------------------------------------------------------
## Sandboxes & Jails

[c, mmchm process sandbox utility] using seccomp ptrace and mruby <IMPORTANT>
[mmchm](https://github.com/udzura/mmchm)
Process sandbox utility, using seccomp, ptrace and mruby 
**CONVERT THIS TO GO**


=======================================================================================================
##== Customize Linux Kernel
* Add encryption of all stdout to the active user key of the controller VM

* Require signature of active user key of the controller VM on all stdin

* Add hooks to file changes in kernel to get inotify with little overhead.

=======================================================================================================
##== Multiverse OS Installer ========================================================================##
** syslogd should be exporting ALL logs



**System Information**
_Check out terminus, it was the most feature rich next to sysinfo_
	[sysinfo][https://github.com/zcalusic/sysinfo]
	<important>MEGA rich sysinfo, will be critical for install

======
hods/Routes]
Preferably have ways to move between the three techniques
 [*] Libcointainer Compost
 [*] Chef Recipie
 [*] Ansible playbook
**Rancher comes with compose processing:**
https://github.com/rancher/rancher-compose-executor/blob/master/resources/host.go
https://github.com/rancher/go-rancher/blob/master/vnext/server_types.go
https://github.com/rancher/rancher-compose-executor/blob/master/resources/volume.go

========================================================================================================
## Replace specific line in configurations

[line-cookbook](https://github.com/udzura/line-cookbook)
Provides a way to modify a single line in a config in RUBY. Should be converted
to Go and used in the configuration of Multiverse OS

========================================================================================================

**PCI Passthrough**
Multiverse OS relies on PCI passthrough to provide usable preformance on VMs and greater isolation
between compartmentalized computers.

````
#!/bin/bash
shopt -s nullglob
for d in /sys/kernel/iommu_groups/*/devices/*; do 
    n=${d#*/iommu_groups/*}; n=${n%%/*}
    printf 'IOMMU Group %s ' "$n"
    lspci -nns "${d##*/}"
done;
````
## Locate and prepare <ALL USB>, <ALL Network Devices>, <SOUND CARD>, <GPU> and almost all other devices
## so that they can be passed through to PROXY VMs or the CONTROLLER VM
options vfio-pci ids=10de:13c2,10de:0fbb
## General kernel modules needed for PCI passthrough
vfio vfio-pci
# vfio_iommu_type1 vfio_virqfd

````
if [ ! -z "$(ls -A /sys/class/iommu)" ]; then
	for DEV in $DEVS; do
		echo "vfio-pci" > /sys/bus/pci/devices/$GROUP/$DEV/driver_override
	done
fi
````

`modprobe -i vfio-pci`

*"Note: If you also have another driver loaded this way for early modesetting (such as nouveau, radeon, amdgpu, i915, etc.), all of the aforementioned VFIO modules must precede it."* https://wiki.archlinux.org/index.php/PCI_passthrough_via_OVMF

**CPU PINNING** (ONLY REQUIRED FOR THE CONTROLLER VM)
*In order to achieve preformance needed to play video games, or edit video, or any other task that requires even a small amount of resources then one needs to configure CPU pinning.*

	*"If you are experiencing high DPC and/or interrupt latency in your Guest VM, ensure you have loaded the needed virtio kernel modules on the host kernel. Loadable virtio kernel modules include: virtio-pci, virtio-net, virtio-blk, virtio-balloon, virtio-ring and virtio."* [PCI Passthrough Documentation](https://wiki.archlinux.org/index.php/PCI_passthrough_via_OVMF)

The default behavior for KVM guests is to run operations coming from the guest as a number of threads representing virtual processors. Those threads are managed by the Linux scheduler like any other thread and are dispatched to any available CPU cores based on niceness and priority queues. Since switching between threads adds a bit of overhead (because context switching forces the core to change its cache between operations), this can noticeably harm performance on the guest. CPU pinning aims to resolve this as it overrides process scheduling and ensures that the VM threads will always run and only run on those specific cores. Here, for instance, the guest cores 0, 1, 2 and 3 are mapped to the host cores 4, 5, 6 and 7 respectively.

$ virsh edit [CONTROLLER_VM]

````
<vcpu placement='static'>4</vcpu>
<cputune>
    <vcpupin vcpu='0' cpuset='4'/>
    <vcpupin vcpu='1' cpuset='5'/>
    <vcpupin vcpu='2' cpuset='6'/>
    <vcpupin vcpu='3' cpuset='7'/>
</cputune>
````

**CPU pinning with isolcpus**
Alternatively, make sure that you have isolated CPUs properly. In this example, let us assume you are using CPUs 4-7. Use the kernel parameters isolcpus nohz_full rcu_nocbs to completely isolate the CPUs from the kernel.

Begin by modifying the grub configuration: `sudo vim /etc/defaults/grub`

````
GRUB_CMDLINE_LINUX="..your other params.. isolcpus=4-7 nohz_full=4-7 rcu_nocbs=4-7"
````

Then, run `qemu-system-x86_64` with taskset and chrt:

# chrt -r 1 taskset -c 4-7 qemu-system-x86_64 ...
The chrt command will ensure that the task scheduler will round-robin distribute work (otherwise it will all stay on the first cpu). For taskset, the CPU numbers can be comma- and/or dash-separated, like "0,1,2,3" or "0-4" or "1,7-8,10" etc.

##===================================================================================================##
##== Build System ===================================================================================##
Multiverse is such a large and important project that a build system to enable any developer to
quickly contribute and to build reproducible builds is important.

#### Mutiverse Build System Components
  **Framework (more than one of the below categories)**
	[mango][https://github.com/vladimirvivien/mango]
	  looks like a good example, i like how it uses go to do configurations

  **Interactive Shell**
  	[gosh][https://github.com/mkouhei/gosh]

  **Dependency Management**
    [dep][https://github.com/golang/dep]
		official dependnecy management package for go


   [dep](https://github.com/deejross/dep)
  **Patch Management/Code Overlay**

  **Makefile**

  **Testing**
	[General]
	  [go-scientist][https://github.com/technoweenie/go-scientist]
	  a port of a popular ruby library for testing

	[Expect]

  **CI**
		[build][https://github.com/golang/build
		official CI lib

  **Mocking**
	[fakehttp][https://github.com/sethgrid/fakettp]

  **Linting** *(Processing, Analyzing, Transforming Source Code)*
##===================================================================================================##
##== User Interface =================================================================================##
#### Text UI                                                                                       
  **Frameworks**
	[clif][https://github.com/ukautz/clif]
	Very reach CLI framework, maybe the most 

	[cli][https://github.com/urfave/cli]
	probably the best CLI framewwork, v2 is very nice

	[gocui][https://github.com/jroimartin/gocui]

	[tui-go][https://github.com/marcusolsson/tui-go]
	a UI library for terminal applications, has great demos like a really nice text editor and
	chat

  **Console (Interactive Shell) Frameworks**
	[gosh][https://github.com/vladimirvivien/gosh]
	Gosh (or Go shell) is a framework that uses Go's plugin system to create for building interactive console-based shell programs

  **Window Simulating Text Frameworks**
	[clui][https://github.com/VladimirMarkelov/clui]

	[termui][https://github.com/gizak/termui]
	
	[tview][https://github.com/rivo/tview]
	wow, very very powerful cli framework

#### Components (Loading bars, spinners, etc)
  **General POSIX Term Utilites**
	[term][https://github.com/pkg/term]
	Package term manages POSIX terminals. As POSIX terminals are connected to, or emulate, a
	UART, this package also provides control over the various UART and serial line parameters.

	[console][https://github.com/containerd/console]

  **Coloring**
	[color][https://github.com/fatih/color]
	Basic color library, nice implementation using `color.Red("text")`

  **Terminal Dashboards**
	[gotop][https://github.com/bunbunjp/gotop]
	best one

	[cryptodash][https://github.com/miguelmota/cryptodash]

	[dockdash][https://github.com/byrnedo/dockdash]

	[myterminaldashboard][https://github.com/sd65/MyTerminalDashboard]

#### Graphic UI                                                                                     
  **General GUI Libraries**
  [pixfont](https://github.com/pbnjay/pixfont)
  A simple, lightweight Pixel Font package for Go that works with the standard image/draw package.



  **2D Game Engine**
	  [2d game engine examples][https://github.com/hajimehoshi/ebiten/tree/master/examples]


 

#### Web UI
  **Web Application**
	Web application libraries and components that would be useful in web applications
	[Frameworks]

	[Components]
	  [users][https://github.com/rivo/users]
	  collection of common user workflows (registering, recovery, logging in/out)

	  [sessions][https://github.com/rivo/sessions]
	  very very nice cookie library

	  [duplo][https://github.com/rivo/duplo]
	  detect duplicate (or similar) images using hashes AND other techniques

  **Web based Terminal / VM Control**
	[jterm][https://github.com/JamesHovious/jterm]
	gopherjs bindings for jquery.terminal 

  **Charts**
	[gopherjs-frappe-charts][https://github.com/cnguy/gopherjs-frappe-charts]
	Uses gopherJS like below

	[go-chartjs][https://github.com/brentp/go-chartjs]

  **Go based JS framework** (because fuck JS)
	[vecty][https://github.com/gopherjs/vecty]
	(^)[https://github.com/gopherjs/webgl]
	(^)[https://github.com/marwan-at-work/vecty-router]
	All the Go based JS frameworks are based on vecty, a project that simplifies
	the process of converting a JS framework to a pure Go implementation.

	_Vecty is a React-like library for GopherJS so that you can do frontend development in Go_
	_instead of writing JavaScript/HTML/CSS._

	[Angular JS]
	  [go-angularjs][https://github.com/wvell/go-angularjs/]

	[Polymer JS]
	  [golymer][https://github.com/microo8/golymer]

	[Vue JS]
	  [go-vue][https://github.com/k2wanko/go-vue]

	[Bootstrap4]
	  [bootstrap4][https://github.com/nobonobo/bootstrap4]
	  Bootstrap4 components for Vecty

	[Material]
	  [material][https://github.com/agamigo/material]
	  possibly more complete than above

	  [material + vecty][https://github.com/wizenerd/layout]
	  (^)[https://github.com/wizenerd/ui]
	  (^)[https://github.com/wizenerd/footer]
	  (^)[https://github.com/wizenerd/icons]
	  (^)[https://github.com/wizenerd/grid]
	  (^][https://github.com/wizenerd/color]
	  Seems to be the compilation of all the others

	  [material tabs][https://github.com/wizenerd/tabs]

	[Mithril JS]
	  [go-mithril][https://github.com/danverbraganza/go-mithril]
	  Mithril.js is a very small and expressive client-side MVC framework. These Go bindings are
	  intended for Developers like me who welcome the strong typing and semantics of Go, and want
	  to use it to easily build great front-end experiences.

	[React JS]
	  [gr][https://github.com/bep/gr]
	  [react][https://github.com/myitcv/react]

	[Autobahn JS]
	  [go-autobahn][https://github.com/cellofellow/go-autobahn]

	[Custom Vecty Frameworks]
	  [gu][https://github.com/gu-io/gu]
	  A web ui library for Go. 

	[cookie modification with gopherjs][https://github.com/fabioberger/cookie]
	[material components][https://github.com/gernest/cute]
	[improved vdom][https://github.com/siongui/godom]
	[detect if running in browser or server][https://github.com/go-humble/detect]
	[random js utils][https://github.com/mrmiguu/jsutil/blob/master/jsutil.go]

  **Three.js**
	[three.js in fucking gopherjs!][https://github.com/Lngramos/three]
		Proof of concept. GopherJS bindings for https://threejs.org.

  **SSE**
	[go-see-sample][https://github.com/nobonobo/go-sse-sample]

	[go-sse-example][https://github.com/jraedisch/go_sse_example]

  **JS Go Examples**
	[highlightjs][https://github.com/myitcv/highlightjs]
	code highlighting

	[gophervideo][https://github.com/csos95/gophervideo]

	[jterm][https://github.com/JamesHovious/jterm]
	Thin wrapper over the jquery.terminal) library for use with gopherjs.

	[goplayspace][https://github.com/iafan/goplayspace]
	A online compiler for Go

	[gopherize.me][https://github.com/myitcv/gopherize.me]
	great example

	[gingopherjs][https://github.com/abrander/gingopherjs]
	example with gin

  **WebRTC**
	[gopherjs-webrtc][https://github.com/chrisprobst/gopherjs-webrtc]

	[webrtc][https://github.com/nobonobo/p2pfw]

	[ssh + webrtc][https://github.com/nobonobo/ssh-p2p]

  **gRPC + Websockets**
	[jterm][https://github.com/JamesHovious/jterm]
	gopherjs bindings for jquery.terminal 

	[input suggest][https://github.com/siongui/gopherjs-input-suggest]

	[protobuf][https://github.com/johanbrandhorst/protobuf]

	[grpcweb-boilerplate][https://github.com/johanbrandhorst/grpcweb-boilerplate]
	(^)[grpcweb-example][https://github.com/johanbrandhorst/grpcweb-example]
	An example implementation of a GopherJS client and a Go server using the Improbable gRPC-Web
	implementation 
	(^)[gopherjs-grpc-websocket][https://github.com/johanbrandhorst/gopherjs-grpc-websocket]

	[wsrpc][https://github.com/utamaro/wsrpc]

  **Example Dashboards**
	[Linux / Status / Monitoring / Health / Profile]
	  [dashboard][https://github.com/JamesClonk/dashboard]

	  [status-dashboard][https://github.com/hverr/status-dashboard]

	  [personal-dashboard][https://github.com/ahmetb/personal-dashboard]

	  [profile][https://github.com/pkg/profile]
	  Simple profiling support package for Go

	  [kapacitor][https://github.com/influxdata/kapacitor]
	  Very nice "Kapacitor use a DSL named TICKscript to define tasks."

	  [go-dash][https://github.com/adelq/go-dash]
	  Nice API of sys info

	  [sysdash][https://github.com/cheilman/sysdash]

	  [chronograph][https://github.com/influxdata/chronograf]

	  [telegraph][https://github.com/influxdata/telegraf] 
  	  Works out of the box with this ReactJS graph dashboard system Telegraf is an agent written
	  in Go for collecting, processing, aggregating, and writing metrics. Design goals are to
	  have a minimal memory footprint with a plugin system so that developers in the community
	  can easily add support for collecting metrics from local or remote services. Telegraf is
	  plugin-driven and has the concept of 4 distinct plugins:
		Input Plugins collect metrics from the system, services, or 3rd party APIs
		Processor Plugins transform, decorate, and/or filter metrics
		Aggregator Plugins create aggregate metrics (e.g. mean, min, max, quantiles, etc.)
		Output Plugins write metrics to various destinations

	[Science]
	  [nviz][https://github.com/makerforceio/nviz]


##===================================================================================================##
#### Functionality

**Low Bandwidth Proxy**
*Proxy, compress all images and cut out content for nodes with weak internet access*

##== Booting ========================================================================================##

#### Open Source BIOS (With Signing)

  **SeaBIOS**
	[seabios][https://github.com/coreos/seabios]
	source code for seabios legacy bios that is built with standard GNU tools  (not best options)
		SeaBIOS is built for QEMU and tested on QEMU with:

		make
		qemu -bios out/bios.bin

#### Open Source Initramfs (With Signing)


=======================================================================================================
##== Networking =====================================================================================##
Networking within Multiverse OS is similar to the complexity of networking on a LAN, it will provide
a separate network from the LAN with very limited access through proxy VMs. The entire network will
be separated through a series of router VMs. Providing access between clusters and users primarily
through the onion hash table (OHT).


#### Data File Types
  TOML, JSON, YAML, CSON (Like JSON+YAML)

#### Clustering
  **x desktop**
    [xdg](https://github.com/kyoh86/xdg)
    Light weight helper functions in golang to get config, data and cache files according to the XDG Base Directory Specification.



  **memory**



  **usb**
    [usb](https://github.com/google/gousb)

  **9p**
    [9p9](https://github.com/docker/go-p9p)
    A modern, performant 9P library for Go. 

  **network interfaces**
    [iface](https://github.com/picatz/iface)

  **UNIX SOCKET!**
    [zero](https://github.com/9b9387/zero)
    A Lightweight Socket Service with heartbeat, Can be easily used in TCP server development.


  **Cluster Jobs**
    [agent](https://github.com/buildkite/agent)
    The buildkite-agent is a small, reliable, and cross-platform build runner that makes it easy to run automated builds on your own infrastructure. It’s main responsibilities are polling buildkite.com for work, running build jobs, reporting back the status code and output log of the job, and uploading the job's artifacts.

  **DNS / Hosts**
    [tlds](https://github.com/picatz/tlds)
    provides ALL current top level DNS. THis willb e useful for checking for valid domain names BUT ALSO
    it will enable peer and friend based DNS look up on non-existant TLDs. For example, Mutlvierse users
    will be able to claim *.cat domains until it exists.

    [dns proxy](https://github.com/gophergala/dnsp)
    dnsp is a lightweight but powerful DNS server. Queries are blocked or resolved based on a blacklist or a whitelist. Wildcard host patterns are supported (e.g. *.com) as well as hosted, community-managed hosts files. Ideal for running on mobile devices or embedded systems, given its low memory footprint and simple web interface.

  **SSH Clustering**
    [LIMITED SSH SERVER!!!](https://github.com/GraveRaven/scpdrop)
		could be used in OHT for limited access and on the fly SCP based file transfer but also for very limited agent/provisioning in multiverse

    [ssh proxy](https://github.com/appleboy/drone-scp/blob/master/vendor/github.com/appleboy/easyssh-proxy/easyssh.go)

    <IMPORTANT>[ya](https://github.com/raravena80/ya)
    mega useful, even nicer YAML configuration than below, commands AND file transfers.
    with cluster management ALREADY in mind
 				 ya ssh -c "touch /tmp/file" -m host1,host2

    <IMPORTANT>[goscpwrap](https://github.com/jrossiter/goscpwrap)
    this is the best one because tis a LIBRARY, not a command line
    making it the BEST CANDIDATE FOR SSH BASED PROVISIONING!!!

		[libscp](https://github.com/kkirsche/go-scp)
    nice concept of libscp

	  [gsh](https://github.com/danielkraic/gsh)
	  provides a YAML configuration for configuring and simplfying control over remote systems
	  this is close to what is desired for Multiverse OS. But for Multiverse OS we need to be 
	  also haver a REST API to modify configuration, add hosts on the fly, configure different
	  authorized keys file
  	_additionally access flow should be stored_ for example if 1 computer can connect to a third but
	  you can connect directly to the third.

	  [1-to-n scp](https://github.com/gophergala/scpm)

    [scpgo](https://github.com/raravena80/scpgo)
    nice, bit nicer than others

    [simplescp](https://github.com/FranGM/simplescp)
	  has some nice configuration handling, like authorized keys stuff

    [rscp](https://github.com/polezaivsani/rscp)
		interesting because it supports bandwidth limiting and I wouldnt know how to do this


#### General Networking Libraries**
  **Go Std Library**
		[go-net][https://github.com/golang/net]

		[bpf:vm][https://github.com/golang/net/blob/master/bpf/vm.go]

		[rawconn][https://github.com/golang/net/blob/master/internal/socket/rawconn.go]

		[ipv4:endpoint][https://github.com/golang/net/blob/master/ipv4/endpoint.go]

		[route:address][https://github.com/golang/net/blob/master/route/address.go]
	  interesting code because it shows that there is inet and kernel inet also reveals
	  some other interesting aspects about routes and interfaces

  **Go Networking Libraries & Examples**
		the standard net lib, very rich

	[go-networking][https://github.com/vladimirvivien/go-networking]
	<important> Very nice library for dealing with IP, protocols, etc

	[gonet][https://github.com/hsheth2/gonet]

	

  **NTP**
  	[ntp client][https://github.com/vladimirvivien/go-ntp-client]

  **Network Tools / Utilities**
  Tools like ping, finger, etc that can be modified to fulfill the needs of the Multiverse OS
  networking tools
	[goping][https://github.com/jcuga/goping]
	Reads a list of addresses, ping frequency, and ping timeout from confing json file and logs ping results.

	[openup][https://github.com/jcuga/openup]
	Utility that opens port forwards via upnp. Written in golang.

  **Iptables / Firewall**
	[gobpf][https://github.com/iovisor/gobpf]
	BCC is a toolkit for creating efficient kernel tracing and manipulation programs, and
	includes several useful tools and examples. It makes use of extended BPF (Berkeley Packet
	Filters), formally known as eBPF, a new feature that was first added to Linux 3.15. Much of
	what BCC uses requires Linux 4.1 and above.

  **Packets / Networking**
	  [gopacket][https://github.com/google/gopacket]
	  very very complete packet system

	  [packets][https://github.com/songgao/packets]
	  very simple lib used in more complete water networking lib
 
  **Linux/Unix Sockets (AF_UNIX or AF_LOCAL or Local Unix Socket)**
    Local Unix Sockets AF_UNIX or AF_LOCAL (http://man7.org/linux/man-pages/man7/unix.7.html) The
    difference is that an INET socket is bound to an IP address-port tuple, while a UNIX socket is
    "bound" to a special file on your filesystem. Generally, only processes running on the same
    machine can communicate through the latter.

    In fact, INET sockets sit at the top of a full TCP/IP stack, with traffic congestion algorithms,
    backoffs and the like to handle. A UNIX socket doesn't have to deal with any of those problems,
    since everything is designed to be local to the machine, so its code is much simpler and the
    communication is faster.

  **Software Defined Networking (SDN): virtual lans, software based routers**
	  [ovn]
	    [iovisor-ovn][https://github.com/iovisor/iovisor-ovn]

	  [vxLAN]
	   [silk][https://github.com/cloudfoundry/silk]
     Silk is an open-source, CNI-compatible container networking fabric. It was inspired by
	   the flannel VXLAN backend and designed to meet the strict operational requirements of
	   Cloud Foundry.

	[Routers]
	  [router][https://github.com/cloudfoundry/gorouter/blob/master/router/router.go]

	[Open vSwitch]
	  [go-openvswitch][https://github.com/digitalocean/go-openvswitch]
	  Go packages which enable interacting with Open vSwitch and related tools

#### Onion Network (OHT)
  **Custom Device (/dev/*)**
  **OHT** and Multiverse PS will heavily use custom devices to tie into current Linux patterns
	[go-tcpip][https://github.com/unigornel/go-tcpip]

  **Virtual network interface (VNI)**
	  **OHT** will utilize a **VNI**, a psuedo network device that abstracts and adds additional features to
	  provide Multiverse with complete software defined networking (SDN). The kernel has its own 
	  virtual network interface table, but the OHT interface MUST be userland and bypass the kernel
	  networking entirely to avoid any potential breakouts via specially formed packets.

	  "The term **VIF** has also been applied when the application virtualizes or abstracts network interfaces."
	  [VNI Wikipedia](https://en.wikipedia.org/wiki/Virtual_network_interface)

#### Network Protocols
  **FTP**
	  [goftp](https://github.com/fclairamb/ftpserver)

  **HTTP Protocol**
	[HTTP Server]
		[puma-dev][https://github.com/puma/puma-dev]
  	HTTP server with incredibly low level options, everything is customizable

	[Websockets]
		[ws-machine][https://github.com/aglyzov/ws-machine]
  	websocket state machine that is fully async, nice implementation
		[stdlib:websocket][https://github.com/golang/net/tree/master/websocket]
		[gordian][https://github.com/ianremmler/gordian]
  	specialized wframeworko for multiclient (like chats)
		[Melody](https://github.com/olahol/melody) is websocket framework based on
	  github.com/gorilla/websocket that abstracts away the tedious parts of handling websockets
		[websocket][github.com/gorilla/websocket]
	  <most-popular> the most popular weboscket lib

	[Long Polling]
    [golongpoll][https://github.com/jcuga/golongpoll]
    [goio](https://github.com/elsonwu/goio)
	[gRPC]
		[echo][https://github.com/bbengfort/echo]
	  gRPC echo example
		[sping][https://github.com/bbengfort/sping]
	  Simple example of secured communication with gRPC and SSL/TLS 

	[WebRTC]
		[go-webrtc-datachannel][https://github.com/coreos/go-webrtc-datachannel]
  	basic and old go webrtc demo

#### Messaging / Queue
  **Centralized**

  **Decentralized**
	[nsq][https://github.com/nsqio/nsq]
	  (^)[https://github.com/nsqio/go-nsq]
    	  (^)[https://github.com/nsqio/go-diskqueue]
	NSQ is a realtime distributed messaging platform designed to operate at scale, handling
	billions of messages per day.

	It promotes distributed and decentralized topologies without single points of failure,
	enabling fault tolerance and high availability coupled with a reliable message delivery
	guarantee. See features & guarantees.

	[go-nats][https://github.com/nats-io/go-nats]
	  (^)[https://github.com/nats-io/go-nats-streaming]
	  (^)[https://github.com/nats-io/gnatsd]

	[Examples]
	  [nats-on-a-log][https://github.com/nats-io/nats-on-a-log]


##===================================================================================================##
##== Portal Gun =====================================================================================##
Portalgun is the toolkit provided by Multiverse OS to spin up virtual machines from recipies, and
provides a package manager, linter and validation for the recipies. Along with networking tools to
provide port routing/forwarding around the cluster.

Portal gun provides the three type of VMs used in Multiverse OS: Controller VM, Application VM,
and Service VM.

  **Multiverse Gateway** (Portal Gun)
    A multiverse gateway is a configurable reverse proxy.
	[Port Forwarding/Routing]
	Proxy direction/ports/ips are *configured at gateway start up* or *on-demand* 

##-- Provisioning / Orchestrating -------------------------------------------------------------------##
The topic of Image creation is closely tied to the topic of provisioning both of which should be 
under the domain of Portal Gun. Cluster management will be simplified to the point that operating
and managing the cluster will feel like working with a single terminal.


#### Notes / Research Relating To `building/provisioning1
We can avoid the need for root by simply allowing the necessary priviledge and specail device creation
binary to segregate this premission from the rest of the software.
	This secondary binary with the necessary permissions may even be able to be embedded. This would solev
all the issues I have been having completely without needing to get so deep into raw socket creation
that I have to manage the files myself (which may still be interesting, like using a pure memory
ring buffer.

			# Creating device requires you a `cap_mknod` privilege
			# try: sudo setcap cap_mknod+ep ./mruby/bin/mirb 


#### Relevant Libraries
  **Examples / Ideas**
	[frenzy][https://github.com/stevedomin/frenzy]
	go vagrant clone unfinished

  **SSH**
	[SSH Utilities / Tools / Libraries]
	  [easyssh][https://github.com/hypersleep/easyssh]
	  [sftp][https://github.com/pkg/sftp]

	  [goshinx][https://github.com/s8sg/goshnix]

		goshclient, err := goshnix.Init("<host_ip>", "<port>", "<uname>", "<pass>")
		envval, err := goshclient.Getenv("<key>")

		fileinfo, _ := goshclient.Stat("<filepath>")
		// Check if its a dir (as of std lib)
		if fileinfo.IsDir() {
		// ...
		}

#### Existing Provisioning Frameworks
There are many more but these specifically are the most interesting canidates 

  **Ansible / Ansible-Clones**
	[Ansible-Clones]
	  [tachyon][https://github.com/vektra/tachyon]

  **Chef / Chef-clones**
	[Chef Utilities]
	  [puck][https://github.com/KAllan357/puck]
	  Puck is a small, HTTP library written in Go, that can be used to run chef-apply on a
	  machine. When started, Puck will listen for requests, parse the JSON into a collection
	  of Chef resources, and execute a call to the chef-apply binary shipped with Chef >= 11.

##-- Virtual Machines / Containers ------------------------------------------------------------------##
  **VM Health / Metrics / Monitoring / Stats / Agent**
	  [go-smbios][https://github.com/digitalocean/go-smbios]
	  SMBIOS is a standard mechanism for fetching BIOS and hardware information from within an
	  operating system. It shares some similarities with the older DMI standard, and the two
	  are often confused.

https://github.com/projectcalico/calicoctl
#### Git
[git2go](https://github.com/libgit2/git2go)
bindings for libgit2
  **Run / Command (Cmd) / Execution / Tasks / Jobs**
	[Inline command-line execution]
	  [exec][https://github.com/pkg/exec]

  	[Scheduled Tasks]
	  [gocron][https://github.com/jasonlvhit/gocron]
		// Do jobs with params
		gocron.Every(1).Second().Do(taskWithParams, 1, "hello")
		// Do jobs without params
		gocron.Every(1).Second().Do(task)

	  [machinery][https://github.com/denkhaus/machinery]
	  Very Very nice Machinery is an asynchronous task queue/job queue based on distributed
	  message passing.

	[Process Manager]
	  [immortal][https://github.com/immortal/immortal]
	  usually i always just try to use systemd or whatever the linux system comes with but this
	  is very ncie and uses YAML configs

	[Distributed Computing / Jobs / Code Execution]
	  [gleam][https://github.com/chrislusf/gleam]
	  (^)[https://github.com/chrislusf/gleamold]
	  Gleam is a high performance and efficient distributed execution system, and also simple,
	  generic, flexible and easy to customize. Gleam is built in Go, and the user defined
	  computation can be written in Go, Unix pipe tools, or any streaming programs.

	  [funnel][https://ohsu-comp-bio.github.io/funnel/]
	  nice applciation to specify a container to run a task

	[Simple / Limited VMs]
	  [blacklight][https://github.com/acook/blacklight]
	  forth like VM So called tasks (or jobs if you like) are executed concurrently either by
	  many workers on many servers or multiple worker processes on a single server using
	  Golang's goroutines.

  **OCI Containers**
	[General]
	  [cgroups][https://github.com/containerd/cgroups]
	  control cgroups

  	[runc]
	  [go-runc][https://github.com/containerd/go-runc]

	[nspawn]
	  [go-systemd][https://github.com/coreos/go-systemd]

	[kubes]
	  [minikube][https://github.com/kubernetes/minikube]
	  Run Kubernetes locally Minikube is a tool that makes it easy to run Kubernetes locally.
	  Minikube runs a single-node Kubernetes cluster inside a VM on your laptop for users
	  looking to try out Kubernetes or develop with it day-to-day.

	  [kubeless][https://github.com/kubeless/kubeless]
	  serverless autoscaling kubes

	  [mantle][https://github.com/coreos/mantle]

	[lxc / lxd]
  	  [lxd][https://github.com/lxc/lxd]
 
 	  [go-lxc][https://github.com/lxc/go-lxc]

	[rkt]
	  [oci container example code][https://github.com/kubernetes-incubator/cri-o/blob/master/oci/container.go]
	  (^)[https://github.com/kubernetes-incubator/cri-o/tree/master/server]

	  [rkt][https://github.com/rkt/rkt]
	  rkt is a pod-native container engine for Linux. It is composable, secure, and built
	  on standards. 

	  [rklet][https://github.com/kubernetes-incubator/rktlet]
	  rktlet is a Kubernetes Container Runtime Interface implementation using rkt as the
	  main container runtime. rkt is an ongoing CNCF effort to develop a pod-native
	  container runtime.

	[runc]
	  [runc][https://github.com/opencontainers/runc]
	  CLI tool for spawning and running containers according to the OCI specification 

	[garden]
	  [garden][https://github.com/cloudfoundry/garden]
	  A rich golang client and server for container creation and management with pluggable
	  backends for The Open Container Initiative Spec and windows.

	  [cli for garden][https://github.com/contraband/gaol]

	  [guardian][https://github.com/cloudfoundry/guardian]

#### Hardware Virtualized Containers / Clear Containers
  **Runtime** *(Actual VM runtime code)*
	[runtime][https://github.com/clearcontainers/runtime]
	OCI (Open Containers Initiative) compatible runtime using Virtual Machines 

  **Agent** *(An agent is the software that runs on the VM to provide secure and limited control)*
	[agent][https://github.com/clearcontainers/agent]
	Virtual Machine agent for hardware virtualized containers

  **Proxy** *(hypervisor container proxy)*
	[proxy][https://github.com/clearcontainers/proxy]
	cc-proxy works alongside the Clear Containers runtime and shim to provide a VM-based
	OCI runtime solution. cc-proxy is a daemon offering access to the agent to both the
	runtime and shim processes. Only a single instance of cc-proxy per host is necessary
	as it can be used for several different VMs. Since Clear Containers 3.0.10, one proxy
	instance per virtual machine is launched for improved isolation.

#### Libvirtd
  **Dashboard**
	[libvirt-console-proxy][https://github.com/libvirt/libvirt-console-proxy]
	Websockets console proxy for VNC, SPICE and serial consoles This package provides a
	general purpose websockets proxy frontend for VNC, SPICE and serial console servers.

	[go-libvirt][https://github.com/digitalocean/go-libvirt]
	Package go-libvirt provides a pure Go interface for interacting with libvirt. Rather
	than using Libvirt's C bindings, this package makes use of libvirt's RPC interface,
	as documented here. Connections to the libvirt server may be local, or remote. RPC
	packets are encoded using the XDR standard as defined by RFC 4506.

	[libvirt-go][https://github.com/libvirt/libvirt-go]

	[libvirt-go-xml][https://github.com/libvirt/libvirt-go-xml]
	XML This package provides a Go API that defines a set of structs, annotated for use
	with "encoding/xml", that can represent libvirt XML documents. There is no dependancy
	on the libvirt library itself, so this can be used regardless of the way in which the
	application talks to libvirt.

#### KVM / QEMU
  **Examples & Useful Code**
	[vm-manager][https://github.com/ZeroPage/vm-manager]
	Nice use of YAML but doesnt support like what HCL offers, inline coding

  **Qemu**
	[go-qemu][https://github.com/digitalocean/go-qemu]

  **KVM**

#### NoKVM


##-- Image Building & Live/Install ISO Creation -----------------------------------------------------##
The topic of Image creation is closely tied to the topic of provisioning both of which should be 
under the domain of Portal Gun

  **Install/Command Execution/Debian Control**

  [pty](https://github.com/kr/pty)
  also includes great example showing how to hadn off TTY of ssh 
				c := exec.Command("grep", "--color=auto", "bar")
				f, err := pty.Start(c)
				if err != nil {
					panic(err)
				}

#### General Notes/Resources
Below are general notes that will be used to develop the early version of the Multiverse OS install
scripts and tools.

	  [ENV][Non-interactive env variable]
		# Ensure that necessary variables are set to enable noninteractive
		# mode in commands.
		export DEBIAN_FRONTEND=noninteractive

		# Update packages
		apt-get --yes --force-yes update
		# http://askubuntu.com/questions/146921/how-do-i-apt-get-y-dist-upgrade-without-a-grub-config-prompt
		# Core problem: post-install scripts don't care that we told apt-get --yes/--force-yes
		DEBIAN_FRONTEND=noninteractive
		UCF_FORCE_CONFFNEW=yes
		export DEBIAN_FRONTEND UCF_FORCE_CONFFNEW
		ucf --purge /boot/grub/menu.lst
		apt-get -o Dpkg::Options::="--force-confnew" --force-yes -fuy dist-upgrade

  **Need a chroot?**
  [chroot using syscalls and os lib][https://github.com/udzura/ltcontainer/blob/master/chroot.go]


#### Building OS Images For Distribution
  **Best Candidate Libraries**
	[distrobuilder][https://github.com/lxc/distrobuilder]

	[osbuilder]
	(^)[https://github.com/clearcontainers/osbuilder]
	(^)[osbuilder scripts][https://github.com/clearcontainers/osbuilder/blob/master/scripts/osbuilder.sh]
	Some useful scripts used by clear containers

  **Building Multiverse OS Images**
    This will be done using the tools provided by Debian, one can initialize a folder with any specific
    version and then run a fixed set of modifications before turning the folder into an image. [Due to
    the nature of this, the ISO is only as secure as the machine it was generated on, so it MUST be 
    verifable, it must be reproducible, and it must be simple to do inside an ephemeral VM. 

  **Build Multiverse OS Live ISO**

  **Build Multiverse OS Install ISO**

#### Unikernels & Microkernels
  **Unikernels/microkernels/nanokernels**
	[unigornel][https://github.com/unigornel/unigornel]
	Unigornel is a library operating system written in Go. It compiles Go code to unikernels
	that run under the Xen hypervisor.

#### Packer
    Packer does exactly what we are looking for but it would need to be gutted to be lean enough
    to satisfy the need of Multiverse OS
	[packer][https://github.com/hashicorp/packer]

	[Using Packer to build the image]
	To use packer to build an image use the following line:

		packer build -var-file=vars/cloud-env.json -var-file=vars/centos.json templates/baseline.json 

#### Media Collection Management
The media library will provide special filesystems for each media type, but in a way that overlays
the original linux filesystem. The special filesystems will organize the media logically, without
disrupting torrents by moving them around and changing the names. 
  **Video** (videos, gifs?)
	[Filesystem]
	[Passive Analysis, Processing & Tools]
	  [*] Auto download summary, reviews, subtitles, art

  **Images**
	[Filesystem]
	[Passive Analysis, Processing & Tools]
	  Images should provide passive metadata removal (if configured) for security, optimization, 
	  or metadata filling based on photoanlsysi, dupe removal
	  [png-tweak][https://github.com/redcap97/png-tweak]
	  optimize pngs
	  [image gallery cleaner]
	  locate images that are similar or the same as friends, compare sizes, format, extra binary data(viruses)
	  and whoever has the best one shares it with the other.


  **Roms** (Emulate games and play with friends)
    Using OHT play roms with friends

	[Gameboy]
	[gameboyGO][https://github.com/gonccalo/gameboyGO]


  **Music** (Audio and Music Videos(?))
  [web based musc player]([go-kkblox](https://github.com/appleboy/go-kkbox))


	[Filesystem]
	[Passive Analysis, Processing & Tools]
	  [*] Auto download summary, reviews, subtitles, art

  **Audio** (Podcasts, audibooks, spoken word)

  **Documents**
	[Filesystem]
	[Passive Analysis, Processing & Tools]

  **Library** (Books, scientific articles, comics?)
	[Filesystem]
	[Passive Analysis, Processing & Tools]
	  [*] Auto download summary, reviews, subtitles, art

  **Video** (TV Shows, Movies, Snippets)

  **Comics**
	[Filesystem]
	[Passive Analysis, Processing & Tools]
	  [*] Auto download summary, reviews, subtitles, art


##===================================================================================================##
##== Multiverse OS Filesystems ======================================================================## 
The filesystem will provide on-the-fly encryption and compression, duplication across the cluster,
index searching and in-file searching, advanced filtering, and provide metadata for virus scanning,
checksums, etc
  **Content Addressable Database**
 	[solidb][https://github.com/vechain/solidb]
	    Immutability
	    Scalability
	    Fault-tolerance
	    High availability

  **block management**
	[blkidx][https://github.com/phicode/blkidx]
	3-4 years old but pure go, may be useful

  **Volume / Storage / Filesystem) Utilities**
	[General]
	  [gofsutil][https://github.com/thecodeteam/gofsutil]
	  A portable Go library for filesystem related operations such as mount, format, etc.

	  [go-lvm][https://github.com/Soulou/go-lvm]
	  pure go control of lvm pvs (logical volumes)

	  [libstorage][https://github.com/thecodeteam/libstorage]
	  libStorage provides a portable and remotable storage plugin framework

	  [block code from container software][https://github.com/thecodeteam/csi-blockdevices/blob/master/block/block.go]
	  CSI-BlockDevices is a Container Storage Interface (CSI) plugin for locally
	  attached block devices. Block devices can be exposed to the plugin by
	  symlinking them into a directory, by default /dev/disk/csi-blockdevices. See
	  sample commands for details. This project may be compiled as a stand-alone binary
	  using Golang that, when run, provides a valid CSI endpoint. This project can also
	  be vendored or built as a Golang plug-in in order to extend the functionality of
	  other programs.

	[File System Notify]
	  [fsnotify][https://github.com/fsnotify/fsnotify]

	[File Format Identification (FFI)/File Types/Magic Numbers]
	  [list of filetypes][https://github.com/datatogether/ffi/blob/master/filetypes.go]

	  [tlv][https://github.com/nleiva/tlv]
	  Custom way of handling filetype by creating a consistent scheme to store this data

  **Filesystems**
	[Virtual/Abstraction Filesystems]
	  Multiverse OS will maintain the normal filesystem so that it remains completely backwards
	  compatible and easily usable by current Linux users.

	  All modifications to structure will be done in virtual or abstracted filesystems. This will
	  increase security by isolating the User, and any User Projects using virtual filesystems
	  that only include necessary files.

	  [csi-vfs][https://github.com/thecodeteam/csi-vfs]
	  CSI-VFS is a Container Storage Interface (CSI) plug-in that provides virtual filesystem
	  (VFS) support. This project may be compiled as a stand-alone binary using Golang that,
	  when run, provides a valid CSI endpoint. This project can also be vendored or built as
	  a Golang plug-in in order to extend the functionality of other programs.

	  [vfs](https://github.com/shurcooL/vfsgen)
	  [vfs](https://github.com/srinathh/vfs)
	  gatefs, httpfs, mapfs, zipfs
	  VFS is a fork of golang.org/tools/godoc/vfs for ease of vendoring 

	[Overlay Filesystems]
	  The core filesystems of the VMs will be immutable, changes are to be stored in overlays.

	[Unique Filesystem Implementations]
	  [mirrorfs][https://github.com/bbengfort/mirrorfs]
	  mirrors whatever happens in another dir 

	  [memfs][https://github.com/bbengfort/memfs]

	  [external-storage][https://github.com/kubernetes-incubator/external-storage

	[General]
	  [afero][https://github.com/spf13/afero]
	  Afero is an filesystem framework providing a simple, uniform and universal
	  API interacting with any filesystem, as an abstraction layer providing interfaces,
	  types and methods. Afero has an exceptionally clean interface and simple design
	  without needless constructors or initialization methods.

	  [blobfs]
	    (^)[blobfs][https://github.com/tsileo/blobfs]
	    (^)[blobsnap][https://github.com/tsileo/blobsnap]
	    (^)[blobstash][https://github.com/tsileo/blobstash]
	    (^)[blobsfile][https://github.com/tsileo/blobsfile]
	  BlobSnap: a snapshot-based backup system designed to provide "time machine" like
	  features. http://blobsnap.com Content addressed (with BLAKE2b as hashing algorithm),
	  files are split into blobs, and retrieved by hash, blobs are deduplicated (incremental
	  backups by default).

	  Read-only FUSE file system to navigate backups/snapshots.
	  Take snapshot automatically every x minutes, using a separate client-side scheduler
	  (provides Arq/time machine like backup). Possibility to incrementally archive blobs
	  to AWS Glacier (see BlobStash docs). Support for backing-up multiple hosts (you can
	  force a different host to split backups into "different buckets").

	  Draws inspiration from Camlistore and bup (files are split into multiple blobs using
	  a rolling checksum).

	  [bup][https://github.com/bup/bup]
	  very efficient backup system based on the git packfile format, providing fast
	  incremental saves and global deduplication (among and within files, including
	  virtual machine images). Current release is 0.29.1, and the development branch
	  is master. Please post problems or patches to the mailing list for discussion
	  (see the end of the README below).

#### Network File Systems
Qemu/KVM sharing of data should be done by mounting the same files, this can be done via something
like httpFS but may be better using plan9, multipath volume, NFS, or other network block/volumes.
	[nfs]
	  [nsfv3driver][https://github.com/cloudfoundry/nfsv3driver]

	  [nfsdriver][https://github.com/cloudfoundry/nfsdriver]

#### Multiverse OS Cluster Files
Multiverse OS will provide a rich file management system that allows for abstracted metadata 
that provides rich features without storing the data inside the actual files to preserve 
privacy.
  **Backup System**
    A backup system is a critical part of good computer security, not only will the filesystem
    support duplication of files across your cluster and friends and peers clusters but it will
    enable stegographic hiding of critical files and wide variety of ways to backup. All backups
    are encrypted. Using metadata backups are searchable. Versioning allows for individual diffs
    to be stored without storing entire copies. 
	[Remote Backup Configuration]
	  Multiple offiste locations can be configured each with their own configuration file
	  back up configurations can be downloaded from peers and shared

    [gobackup](https://github.com/huacnlee/gobackup)
    s simple tool for backing up databases and files via ftp, scp, s3


	[Stegographic Backup System]
	  Backup of critical data, like crytographic keys, can be stored in images or other files
	  and stored in social media profile pictures, or other locations

  **Passive Analysis & Processing**
    Passive analysis of files will be a cornerstone of the security system of Multiverse OS
    while also providing a simple way to hook in automation on files
	[dupefinder][https://github.com/rubenv/dupefinder]
	Detect duplicate files across different machines, using SHA256


  **Multiverse Cluster File Search**
    A search with rich filtering primarily for files and file contents but also more. Also automatically
    generating metadata for better searches, like using machine learning to detect contents of pictures
    auto caption then provide searching against images that includes description of contents. Same could
    be done for audio files, etc

    Using NLP to extract keywords and tags, search of encrytped, compressed or even images or audio can
    be done.

	[Search friends and peer files]

	[Search and replace]

	[Search and rename]
	

#### Peer2Peer (P2P) & Friend/Family2Friend/Family (F2F)
P2P and F2F file sharing will be done using mountable file systems that
move files around like a torrent, allowing selectable downloading, metadata
fields to inform peers and friends which files you would like backed up,
metadata for recommendations to specific users, and ability to search against
names and metadata even if the files are not downloaded (maybe even keywords, tags
so contents can be searched using NLP without ahving the entire contents).
  **Open Source S3**
	[Open Source S3]
	  [minio][https://github.com/minio/minio]
	
  **Distributed FS**
	[Filesystem Options]
	  List of known options for filesystems, some will probably not work
		* rootfs
		* brtfs
		* glusterfs
		* seaweedfs
		  [seaweedfs][https://github.com/chrislusf/seaweedfs]
		  This matches a lot of functional requirements for the Multiverse FS
		  SeaweedFS is a simple and highly scalable distributed file system.
		  There are two objectives: to store billions of files! to serve the files
		  fast! Instead of supporting full POSIX file system semantics, SeaweedFS
		  choose to implement only a key~file mapping. Similar to the word "NoSQL",
		  you can call it as "NoFS".
		* zfs
	  	  [zfs][https://github.com/containerd/zfs]
	  	  only a snapshotter plugin not zfs
		* ContinuityFS
	  	  [couninuityFS][https://github.com/containerd/continuity]
		* mapfs
	          [mapfs][https://github.com/cloudfoundry/mapfs/blob/master/mapfs/mapfs.go]
	
##===================================================================================================##
##== Scramble Suit (Account System) =================================================================##

#### User Account System & Authorization
Using a the decentralized system built up with Multiverse OS, it will not only be able to provide
local authorization and access control but it will provide the framework for a decentralized online
authorization that will be compatible with the existing oauth2 system.

This system will also enable users to login to their "computer" from any computer securely. A persons
computer becomes a logical thing that can be loaded from backups anywhere.
  **Account System & Identity Management & 'Ephemeral Key Tree' Authorization System**
    (Keys & Password Management) Using a master-key generated during the Multiverse install process,
    signed messages are created with time-outs to temporarily authorize a specific key, exiting within
    a tree, and cascading heirachal priviledges granted by the ephemeral key location in the tree.
	
	
	[Password Generation]
	  [xpwd](https://github.com/feckmore/xpwd)
		Mnemic based passwords pased on XKCD comic on the topic
	[Ephemeral Key Tree Keyring]

	[Identity Management System]

	[Unique IDs]
	  [nuid][https://github.com/nats-io/nuid]

	  [ulid][https://github.com/oklog/ulid]
	  Universally Unique Lexicographically Sortable Identifier

	  [uuid][https://github.com/google/uuid]
	  The uuid package generates and inspects UUIDs based on RFC 4122 and DCE 1.1: Authentication
	  and Security Services.
	
	[Remote Access Control]

	[Authorized Keys Management]

	[SSH Host Management]

	[Password Store]
   	  [gopass][https://github.com/justwatchcom/gopass]
	  designed is similar to pass, but for teams and ahs other cool features may be a good
	  starting point

  **3rd Party Auth**
  [stdlib:oauth2][https://github.com/golang/oauth2]
	bunch of login libs in the stdlib

	[goth][https://github.com/markbates/goth]
	Unlike other similar packages, Goth, lets you write OAuth, OAuth2, or any other protocol
	providers, as long as they implement the Provider and Session interfaces.
	
	With like 50 oauth providers

  **Script/Hook system** (Super advanced cron)
    A very rich scripting and hook system that you can register scripts with, share or
    download. Set of easily changable variables by configuration. And hookable to OS
    events from changes in cryptocurrency balances, scheduling, calendar events, result
    of regular webscraping, etc

  **Finances / Bookkeeping**
    Using crawler and prebuilt software for various banks, bank data is downloaded and
    stored locally to build graphs, do analysis, and to better use financial data without
    handing it off to a third party.

  **Scheduler**

  **Calendar**

  **Article and Link Sharing**
    Built in google reads like system to share pictures, links, articles decentralized with
    firends and peers.

  **Endpoint management**
    Manage avaialble endpoints, share with friends, forward ports, host projects, etc

  **Domain Name Manager**
    AUtoamtically generate TLS, assign to endpoints, add TXT records

  **Action validation**
    Commands, etc, can be run on 2-3 different VMs running side by side to verify results
    and detect anomylous behavior.

  **Contact System**
    Rich contact system that syncs across devices, updates, validates against freinds,
    lookupable through friends @friend.friend-of-friend

  **Status System**
    Basic twitter like status system built on decentralized OHT system

  **Shell History**
    Shell history needs a masive overhaul, it should be colored, timestamped, encrytped, 
    and should be smarter. If you regularly use a command, it should be easy to save it 
    to alias or just in a snippet DB that is easy to search.

  **Complete REST API**
    To empower scripting and software development of the Multiverse OS cluster, all of the
    Multiverse OS functionality will be accessible by REST API with the output data format
    up to the user (XML, JSON, etc)

    Everything from filesystem operations to devices will be accessible by REST API. This will
    also simplify access to friends and peer computers since scritps can just be changed to 
    what is pointed at.

  **Configuration Management**
    *Defaults should be more secure!*

    Configurations should be done in whatever data format (JSON, XML, YAML) is preferred by the user,
    this can be done on the fly so that the existing configuration styles can be used for expereinced 
    users. This is important to make linux easier to learn, use and customzie. 

    Configurations should automatically follow the user around, when logging into remote computers, 
    across the cluster and so on to prevent wasting time reconfiguring the same settings over and
    over (with the expection of selected configurations, with secure defautls exempt). 
	[Configuration Management]
	  <important>A tool to change configuration values in a configuration file from CLI, REST API,
	  or other ways is needed to simplify configuration and provisioning.

		[


	[sed]
	  Additionally, find and replace, and find should be simplified and the output to the terminal
	  should be made more sensical and clear.

		# Substitute with sed
		sed -i 's/expr/newexpr/g' file

		# Comment line that satisfies 
		sed -i 's/^/#\ /g' file

		# Append to every line
		sed -i '$ a text' file

		# Append after expression
		sed -i '/expr/a text' file

  **Scramble Shell decentralized secure DNS**
    A secure DNS system that uses local and peer/friend DNS systems to add additional anonymity
    layers. Additionally DNS system to interact with peers/friends will be overlayed ontop of 
    the existing system, enabling users to interact with others through tree of relations (access
    specified by the user). 

    Additionally, peers and friends can share blacklisted, blackholed hosts as well as overlay
    a free DNS system accessible by friends and peers using the scramble shell DNS system. 
		
		[go:net stdlib][https://github.com/golang/net/blob/master/dns/dnsmessage/message.go]
		DNS code from std lib

		[Cryptocurrency Addresses DNS]
	  	A system to send cryptocurrency to users based on a random address they published to
	  	either to DNS recoreds using TXT records automatically or published via OHT. So
	  	you can send to a new unused address from a friend without asking, just using their
	  	name.

  **Scramble Shell decentralized cache**
    When requesting a website, one can use bloomfilters to check if peers already ahve the data
    this adds anonymity but also prevents tampering. THis can also be used to stream from multiple
    caches alongside the original source to create on-the-fly CDNs to speed up loading times.

	[Local Media Collections]
	Already have something in your local collection, say a picture? load from the local


  **Scramble Shell decentralized network request verification**
    REquest the same page across several peers/friends and diff the results.

  **Scramble Shell Inventory System**
    To simplify both activated and available customization in the Multiverse OS Linux shell, an
    'inventory' system will help users visualize the customizations. For example, one will be able
    to hover over the CLI Prompt 'inventory slot' (because there can only be (1) CLI prompt per
    user), they will see the code used to customize the prompt, the interpreted result and see
    other available and popular options that can be dropped in to replace the active CLI prompt
    selection. 

    This system may not be explicitly named 'inventory system' but we believe this will simplify
    use of Linux and make it easier for new Linux users to discover what customizations they can
    make, and allow them to make changes, without becomming overwhelmed, while still enabling
    advnaced users to modify their CLI prompt (or other customziations) in anyway they prefer.

    They will also find it easier to share their customizations with less technically savvy
    friends sending them inventory items. All the files will exist where an experienced user
    will expect them, the inventory system just functions as a logical overlay to simplify
    ones understanding of the system.
	[Dot-Files]
	  Dot file management is an important aspect of Linux configuration, scramble shell
	  will provide install and cleanup, for any server connected to. Along with a package
	  manager for settings to allow users to share and install selected settings from
	  other peoples dot-files based on popularity.
	  [Management]
			[adm](https://github.com/udzura/adm)
	    very nice manager
	  [Examples]
	    [dotfiles][https://github.com/fatih/dotfiles]

  **Cryptocurrency**
    Multiverse OS must provide a secure way to access cryptocurrency like running each client in
    its own Application VM and providing a consistent and unified WebUI to interact with each
    cryptocurrency. These machines can be offline and create transactions that are broadcasted
    from other VMs. HD Keys, offline generation and operations. In Addition, local blockchain
    explorers are desirable for both API and offline usage. Light or full will be options.
	[Wallets]
	  Watch address balances, right hooks based on changes.
	
  	  [Blockchain][https://github.com/Akagi201/blockchain]
	  Key generation for random blockchains: ethereum, bitcoin, hyperledger fabric and so on 


	
	[Blockchain Explorers]



  **Scramble Suit User Database / Dictionary (Key/Value + Graph DB)**
    User database k/v, and graph that can be used and queried for all kinds of data
    then profiles can be created to create commands to queiry the data easily
    for example it will come with one for storing aliases, snippets, notes, dictionary

    Configuration will be done in yaml/json like files and easy to extend
	[Generic Key/Value Store]
	  [redis][https://github.com/go-redis/redis]
	  go redis client

    [BOLT DB BASED KV STORE](https://github.com/laher/kv)
    already ahs CLI command

	[Generic Distributed Databases]
	  [consul][https://github.com/hashicorp/consul]
	  distributed, highly available, data center aware for configuration. k/v
	  store, service discovery, health checking

#### Scramble Suit Toolkit
A collection of tools that will empower the user to script, customize and enable
new ways of using their computer with automation, machine learning, and clustering
Providing tools being used by corporate world and science to every user

  **Dictionary**
	https://en.wikipedia.org/wiki/DICT

  **Publishing**

	[PDF Generation]
	[publisher][https://github.com/speedata/gogit]
	XML to PDF



  **Notes (automatic wiki)**
    Notes saved, publishable into a article, automatic linking across articles, graphs with graphviz,
    networkable with friends notes. Forming a distributed wiki of notes. Version control

  **Distributed Code Version Control**
    Git is already distributed, it would not be hard to decentralize it. 

	[Read Only Access]
	  [gogit][https://github.com/speedata/gogit]

	[Existing Frameworks]
	  [gogs][https://github.com/gogits/gogs]
	  There is also gittea, the fork made by people annoyeed with the elad developer	

  **Projects**
    Project management and segregation by identity.

  **Science**
    Tools and datatypes for scientists
	[Algorithms]
	  [Bioinfomatics]
	    [bio][https://github.com/shenwei356/bio]

	[DNA datatype]
	  Store DNA and genome data in a more logical and consistent way.

	[Chemical datatype]

  **Web Crawler**
    Easy to use web crawler that can be configured to pull specific data from specific
    areas.

  **OSINT**
	[Public records]
	[Whois history lookups]
	[DNS history lookups]
	[Public IP lookups from friends]
	  [ipd][https://github.com/mpolden/ipd]

  **Secure Random**
    Randomness generated by sharing entropy with friends and peers. 

  **Machine Learning Toolkit**
    Built in models, new models can be downloaded and shared through a package management
    system. Pretrained with seeded data (or percent of data from peers or friends).
    THese will appear as devices /dev/*

  **Sorting Algorithms**
    Devices that take in a list of items and sort or organize based on the device and
    options set when using the device.

  **Basic Computer Scanning / Pen test**
    Basic computer scanning: port scanning, ip address range, service checking, pinging
    health checks, brute force, and other common pentest tools. This is important for
    passive checking of own cluster and friends clusters for potential attack vectors
    but also checking remote servers

	[Computer Scanning]
	  [Ping check]
	  [Health check]
	  [Port scan]
	  [Service check]
	  [Banner grab on port]

	[Web Application Scanning]
	  [XSS]
	  [SQLi]
	  [Subdomain Scanning]
	  [Directory Scanning]

  **Publish/Subscribe Device**

  **Subtraction/Cutting**
    Ability to remove, cut, disable data from hosts, elements on webpages, timestamps
    can be used to remove content from media (like ads). These can be shared and 
    downloaded.

	[host block]
	  high level host blocking, based on IP and other metadata that the host leaks
	  over the internet (banners, etc). Can be used for lower level ad blocking
	  and defending against malware

	[element cut]
	  using a built in transparent proxy, specified content or elements can be
	  removed. improving web experience and 

	[media cut]
	  A system which stores timestamps and file name or other data to identify 
	  files which can be modified with cuts. These cuts can be shared to for
	  example, share locations in files that advertisemnts are in podcasts 
	  so they can be skipped or cut out completely automatically.

  **Decentralized Search / Proxy Search**
    Submit your searches through peers or other endpoints to provide additonal layers
    of automization. Additionally a decentralized search will be built into Scramble
    shell.

  **Debugging**
    Memory analysis, stats, resources

	[Memory]
	  [memviz][https://github.com/bradleyjkemp/memviz]
	  Visualize your data structures using graphviz

  **Transpiling**
    Convert Go to C, or C to Go, etc

	[Transpile: c2go]
	  [c2go][https://github.com/elliotchance/c2go]


#### Scramble Suit Resources
Scramble suit and Multiverse will provide a collection of devices, REST API and other
ways to access a collection of offline tools that empower both user software, scripts
but in general. From built in converters (temperature: f-to-c, metric-to-imperial, 
transcoding videos to other qualities, transcribing audio or video files, NLP, etc)\

These tools will enable an identity of the user can have a stylogrpahy profile, that 
can be checked against to ensure the writing style remains consistent and is differnet
from other identities maintained by the user.
  **Scramble Suit Transforming/Transcoding**
    Multiverse OS does is not opiniated, it will not force choices on the user unless
    they are explicitly related to security.

	[Send message to peer based one where they are available]

	[Encryption to peer by name or email or other identifier]

	[Humanize data]
	
	[Translate]

	[Stylographic Analysis & Conforming]
	  Analyze stylographic features of text, compare to known identiteis, and provide
	  the ability to conform text to a specific identity stylography

	[Format Text / Inflection]
	  [inflect][https://github.com/martinusso/inflect]
	  [inflection][https://github.com/jinzhu/inflection]
	  [inflect][https://github.com/markbates/inflect]
	  *Example usage*
			func AddAcronym(word string)
			func AddHuman(suffix, replacement string)
			func AddIrregular(singular, plural string)
			func AddPlural(suffix, replacement string)
			func AddSingular(suffix, replacement string)
			func AddUncountable(word string)
			func Asciify(word string) string
		...

	  []

	  [Binary Analysis & Reverse Engineering]
		  [*] General stats, top 10 most common strings, encoding arch, length, etc
		  [*] Convert to hexdecimal & ASCII tring
		  [*] Output reversed engineered software as source file (way to guess functions
		  and 

	[Data Storage Format Conversions]
	  [*] Conversionb between JSON, TOML, XML, YAML, etc
		


		TOML parser for Golang with reflection.
  	[locker


	  [*] CSV, TSV to Database (SQL, mongo, etc)
	  [csvtk][https://github.com/shenwei356/csvtk]
	  a cross-platform, efficient and practical CSV/TSV toolkit

	  [goxlsx][https://github.com/speedata/goxlsx]
	  Excel-xml reader for Go

	[URL or text shortener]

	[Image analysis]
	  [*] Find % of color, and other color information
	  [*] Find shape
	    [imageshaper][https://github.com/speedata/imageshaper]
	  [*] Facial recognition	

	[Transforming: (Un)Compression, (Un)Minifying, (Un)Obstuficating]

	[Scientific Naming: Name(common names)-to-Taxon]
	  [taxonkit][https://github.com/shenwei356/taxonkit]
	  Cross-platform and Efficient NCBI Taxonomy Toolkit
	
	[Math: average, min, max, etc]

	

	[Graph/Plot]
	  [3D Plot]

	  [2D Plot]
	[Geodata]
	  An offline (and online) toolkit for looking up information relating to geodata. 
	  From IP address to region, to zipcode to state, and so on.

	[Validators]
	  A collection of validations, that can be used as a device/rest api to provide
	  consistent and relaible validation. If we can standardize validation, we can
	  dramatically improve security. From base58 validation, bitcoin address validation,
	  etc. Configurable with configurations and sharable and downloadable with package
	  manager.

 	  [Email Validation]
	    [is disposable email address?]
	    [is government email?]
	    [is student email?]
	    [is free email?]
	    [has smtp responding on host?]


	[File chunking]
	  Chunking for torrents and more

	[Fake data generation]
	  Ability to generate realistic fake data, for example, get a real address from a city
	  randomly. This can be paired with many other things to improve functionality.

	  [MAC Addresses]
	    Generate real MAC addresses with the ability to supply specific vendors or common
	    devices, or real mac addresses pulled from WIFI (by you, peers or friends).

	[Physical Property Conversions]
	  Weights, unit measurements, and other conversions will be configured by configuration
	  sharable, and installable using a package management system.


	[Labstack (Developers of Echo Web Application Framework] provide an API that provides VERY similar functionality to this section:
			Barcode Generate
			Barcode Scan
			Chart Area
			Chart Bar
			Chart Line
			Chart Pie
			Currency Convert
			DNS Lookup
			Download
			Email Verify
			Geocode Address
			Geocode IP
			Geocode Reverse
			Image Compress
			Image Resize
			Image Watermark
			PDF Compress
			PDF Image
			Text Sentiment
			Text Spellcheck
			Text Summary
			Webpage PDF
			Word Lookup


  **Scramble Suit Analysis/Extraction/Summarizing**

  **Scramble Suit Learning Dataset/Database**
    Made up of three interconnected databases that actively grows, shards can be selected based
    on knowledge type or the full set can be installed. 

    dScramble Suit provides a knowledge database that is growing and shared with other users and
    very importantly providing an API to use the data in software easily. The toolset can also
    be used for grabbing data for your own need and storing in the user database.

	  [sentry][https://github.com/datatogether/sentry]

	  [web archive WARC][https://github.com/datatogether/warc]

	  [cdxj - open way back machine][https://github.com/datatogether/cdxj]
	  Golang package implementing the CDXJ file format used by OpenWayback 3.0.0 (and later) to
	  index web archive contents (notably in WARC and ARC files) and make them searchable via a
	  resource resolution service.

	  The format builds on the CDX file format originally developed by the Internet Archive for
	  the indexing behind the WaybackMachine. This specification builds on it by simplifying the
	  primary fields while adding a flexible JSON 'block' to each record, allowing high flexiblity
	  in the inclusion of additional data.

#### 'Matter-of-fact' Knowledge Base Data Mining
	**Wikipedia**
	Wikipedia is a major source of 'matter-of-fact' data that will be mined to populate the data-
	base. Direct use of JSON API should be used to simplify data extraction and avoid use of the
	Ruby library Mechanize. Additionally, related Wikimedia datasources like **WikiData**, and
	**Wikitionary**.

	  [go-wikiparse][https://github.com/dustin/go-wikiparse]

#### Multiverse Desktop Environment
  **Scramble Shell (Desktop Environment)**
    Either built upon Gnome or a completely custom Desktop Environment built in Go or Rust using Wayland 
    including features:
	[*] Tree-based Desktop Workspaces (to organize workspaces by category and connect them together
	in a heirachical fashion) OR Continuous Desktop Ribbon
	[*] Wallpaper / Window Color to separate user identities

	[General]
	  [wallpaper][https://github.com/reujab/wallpaper]
	  A cross-platform (Linux, Windows, and macOS) Golang library for getting and setting the desktop background.
	[Gnome]
	  [gnome notes/research]
			Gnome extensions are defined in an XML file and stored in your local/config local/share folder somewhwere.


	  [gse][https://github.com/reujab/gse]
	  tool for downloading, searching and installing gnome extensions

  **Scramble Shell (mvsh)**
    Multiverse scramble shell seeks to make Linux and the shell easier to understand, learn and improve.
    A big part of this update is upgrading every GNU core-util while keeping it completely backwards
    compatible with a shim. The changes will include (1) make all the naming consistent (2) make all
    the flags consistent (3) provide progress bars, coloring and other TUI updates.

    Multiverse shell will be built around the concept of a cluster instead of a single computer. For
    example, `cp` will copy from within a computer or to a different computer, so there will be no
    need for scp anymore (though it will still exist to be completely backwards compatible). Running
    commands across multiple computers in the cluster, filtered set of computers, etc will be natural
    as possible. 
	[ls]
	  `ls` command needs a major update, it should support much simpler argument passing and 
	  very importantly it must be colored, directories provide actual fucking size, and advanced
	  filtering

	  For example:

			list all movies
			list movies newest first
	
	  This type of natural language control is critical to empowering new users, making it simple
	  to learn and allowing more complex features. THis can be done while still remaining completely
	  backwards compatible so experienced users can still use `ls` without any issues.

	[Scripting Language]
	  Either Lua or Ruby should be replace bash to not only simplify programming bash but make 
	  it dead simple to have frameworks to reuse code and codify secure best practices.

	[Calculator]
	  Built in calculation and conversion tools to reduce the use of google to do simple math
	  or simple conversions. Including date math
	
			Today - 24.years
			4 * 44

	[Vim (maybe OR Emacs style]
	  Many people do not know the available built-in hotkeys in shell/terminal, and that they are
	  based on emacs. As a vim user, I often find myself wishing some of the vim hotkeys worked. 
	  The following will be updated and modified
		Ctrl + a: Home
		Ctrl + e: End
		Ctrl + k: Cut to $
		Ctrl + u: Cut to ^
		Ctrl + w: Cut next word
		Alt + d: Cut previous word
		Ctrl + y: Paste Bash clipboard
		Alt + u: Change word to uppercase
		Alt + u: Change word to lowercase
		Alt + t: Swap current word with previous
		Ctrl + r: Reverse search
		Ctrl + s: Freeze command line
		Ctrl + q: Unfreeze command line
		Ctrl + : Kill -3 SIGQUIT

	  To get vi commands you can add this to your `~/.bashrc`
		set -o vi


	  The brace options in bash shell needs to be expanded upon, should support tab,
	  tab allowing multiselect:
		Brace expansions

		Useful for copying or moving files with long paths.

		mkdir -p /tmp/first/second/thirs
		mv /tmp/first/second/{thirs,third}

		Or generating files

		cd /tmp/
		mkdir -p {1..4}
		touch {1..4}/testfile

	[Cluster Commands & Management]
	  Commands can be run on the entire cluster, single computer, group of computers, class of 
	  computers and so on.

	[Simplification Shim]
	  A backwards compatible shim will provide consistency between ALL console commands available,
	  which will simplify using AND learning Linux.

	[Automatic Benchmarking]
	  All applications will be benchmarked, storing time of exeuction and change in computer resources
	  over the course of usage. This will allow better debugging, locating resource hogs, and potentially
	  discovering malware via anamalous activity from applications.

	[GNU Core Utils]
	  All command output should be encrypted with the users key, it should be optional to output in
	  JSON, XML, etc

	  <idea title="obstufication of binrary paths">
			since we will be using abstracted virtual filesystems to rearrange the structure, we could have two layers, and scramble the locations and structure on boot. 1 layer would provide mapping to classic structure and names, and the second to the new Multiverse OS structure.

			what is interesting about this concept is, the file system without these overays would be hard to impposssible to use by mixing everything around randomly, renaming. this is osbstufication but it could stop a lot of bots runnig automated attaks in their tracks.
 	  </idea>

	  [GNU Core Utils]
	    [tree][https://github.com/caelifer/tree]

	  [Other Utils]
	    [cmd][https://github.com/rjeczalik/cmd]

	    [iploc][https://github.com/NothNoth/iploc]
	    A unix tool that will take in an IP and replace it with IP+location data

	[Scramble Shell Console / Terminal]
	  Scramble shell will replace `gnome-terminal` providing a modern UTF-8 encoding, live updating
	  text feilds:

	  [*] Support regex based coloring
   	  [*] Stdin requires signature from active key, and stdout is encrytped to an active key
  	  [*] Support images by converting to ASCII/Emoji and displaying

	  [console][https://github.com/containerd/console]
	  package for dealing with consoles

#### Multiverse OS Clustering
  **Multiverse Cluster Logging**
    A rich logging system that forwards all logs to the controller VM. Logging should be filterable,
    searchable, colored, easy to filter out applications/services, easy to clear on a timer,
    automatically searched for specific things to alert (x number of failed logins), and machine learning
    to find anomalies. Stored as JSON and other easy to use datatypes. Histograms for log generation
    statistics
	[Examples]
	  [oklog][https://github.com/oklog/oklog]
	  OK Log is a distributed and coördination-free log management system for big ol' clusters. It's an
	  on-prem solution that's designed to be a sort of building block: easy to understand, easy to operate,
	  and easy to extend. Supports forwarding, querying, 

  **Distributed Ledger & Distributed Computing**
    Using OHT a distributed ledger and computing system will be intergrated directly into Multiverse
    OS. It will enable collective decision making, provide a framework for the distributed open
    source development and feedback system while providing tools for distributed organizing of 
    individual users and their communities
	[Examples]
	  [fabric][https://github.com/hyperledger/fabric]
     	  (^)[fabric-sdk][https://github.com/hyperledger/fabric-sdk-go]
	  (^)[Membership/Identity functionality][https://github.com/hyperledger/fabric-sdk-go/blob/master/internal/github.com/hyperledger/fabric/msp/msp.go]

##===================================================================================================##
##== Home Automation ====================================w============================================##
Track all devices, load measurements, etc

#### Media (projectors/tvs/etc)
[go-dell](https://github.com/Grayda/go-dell)
control dell projectors, good example code

[driver-samsung-tv](https://github.com/ninjasphere/driver-samsung-tv)
another example

#### Software Defined Radio (SDR)

#### Infrared (IR) Networking / Relays / Device Control
  **JS Libraries**
    [ninja-allone][https://github.com/Grayda/ninja-allone]

##===================================================================================================##
##== Robots =========================================================================================##
Track all devices, load measurements, etc

	[*] **IR** Controller

	[*] Serial, CAN [TODO: What is the abbreivation], *universal serial bus* **(USB)**, and other common
		  protocols used to interface with components, and devices.

	[*] *Software Defined Radio* **(SDR)** interace


#### Interface
  **GPIO**
	[gpio][https://github.com/brian-armstrong/gpio]

#### Fabrication
Fabrication and development of physical objects is important to computing as traditonal ink+paper 
printers, and seamless integration is essential for any modern operating system. This failure to
supply basic features that were supplied for previous generation common peiriphas is anotherr0
	**3D Printing Manager (Modeling, Slicing, and Pritning)**
  	[octoprint][https://github.com/dustin/octoprint]
	  go client for interacting with slice+print software

	**'(Laser Cutter / Drawing / Cutting) Robot & Other 2D based robots' Manager**

	**'CNC, Lathe, Drillpress, and Other Hardware' Manager**

##===================================================================================================##
##== Mobile =========================================================================================##
Mobile integration is essential for any modern operating system to be relevant and security of your
home/desktop computing environment is related to the security of your mobile device.

#### Phone & Tablet Integration
Multiverse OS seeks to eventaully provide a complete linux based replacement for Android for all
mobile devices. In the meantime, focus will be on the following functional requirements:
	[*] Auto wiping and restoring to make devices more ephemeral
	[*] Scripts to run in the background to add security. For example,. encrypting all photos as they
	    are taken
	[*] Once in range, all new photos and such are downloaded, filetypes which are set to upload and
	    are new get uplaoded when in range ie comics downloaded automatically get uploaded
	[*] Checksums on all files, checked when in range
	[*] Highly configurable fallback system will allow for setting up how services fall back and 
		  in what conditions fallback occurs. For example, one can set their device to automatically
		  and seamlessly move from SMS messages to instant/text messages once connected to internet
		  access independent from the phone ISP connection. The fallback will occur seamlessly, and
		  will not require the user to switch applciations to continue conversation, switching
		  protocols will be fully supported withing the instant/text messaging software included
		  with Multiverse OS. The same ability to switch to VOIP from *conventional* phone calls. 
	[*] SMS messages should be received within the Controller VM, be seamlessly intergrated wtihin
		  the instant/text messaging software that comes packaged with Multiverse OS.
	[*] Allow for seamless intergration of phonecalls, which can be placed on desktop, or phone, 
	  	as well as be capable of being trasnfered between the devices without interrupting the
	  	phone call.
	[*] Mobile devices with Edge/2G/3G/4G/... internet access will provide this connection to the
		  cluster as an additional endpoint for the Multiverse OS **SDN**.

##== Documentation / Presentations ==================================================================##
Documentation is very important, it needs to be easy to access, search, full of interactive examples
graphs, and offline. 
#### Docs
	**'Wiki'**
	  Multiverse OS will provide it's own 'flavor' of Markdown that will expand on the existing set
	  of features found in both Wikipedia and Github custom flavors; in addition the expanded features
	  in these 'flavors' of Markdown the Multiverse OS 'flavor' will focus on providing functionality
	  for modern computing tasks, such as 3D Models used in 3D printing, functionality for scientific
	  publishign, such as rich peer review tools, interactive graphing, and interactive datasets 
	  intended to modernize the scientific article and make the data, math and software open.

	  Major emphasis will be put the functionality for a rich and customizable peer review, supporting
	  both pre-publication and post-publication reviews. Multiverse OS seeks to include open source
	  platform for open, transparent gray-research/citizens-science articles. This publishing platform
	  will focus on providing both the development tools but professional epub/pdf/html5 generation, 
	  and pre-print like decentralized services to publish for peer review.

	  Our hope is that Multiverse OS will be among other open-source publishing platforms that will
	  empower a new generation of gray-research, citizen-science that are intensly peer-reviewed
	  and improved by authors creating aritcles that compete with the qualtiy of academic papers
	  and whitepapers produced by corproate scientists.

	  These tools will hopefully encourage the use of peer-review in other types of articles and 
	  blogs/journals, to improve the overall qualtiy of writing of the services using this Markdown
	  'flavor'.

    [Formatting]
		  [Wikimedia Flavor Markdown](https://www.mediawiki.org/wiki/Help:Formatting)
	    Wikimedia has its own custom features, which in combination with the custom Markdown features
		  provided in Github will be the basis for the Multiverse OS flavor Markdown.
     		[Italic text]
	      Equvilent to the `<em>` HTML tag, previously the `<i>` tag in earlier versions of HTML.
			  	''italic''

		    [Bold]
	      Equivilent to the `<strong>` HTML tag, previously the `<b>` tag in earlier versions of HTML.
	      	'''bold'''

			  [Numbered list]
			  Start each line with a [[Wikipedia:Number_sign|number sign]] (#).

					## More number signs give deeper
					### and deeper
					### levels.
					# Line breaks <br />don't break levels.
					### But jumping levels creates empty space.
					# Blank lines

					# end the list and start another.
					Any other start also
					ends the list.

	  			[Indent text]
			    A way to ident text blocks at different levels with no symbols or numbers.
	 					: Single indent
						:: Double indent
						::::: Multiple indent 

			    [Definition list]
		      Definition list seems redunant since this functionality could be acheived with just using
	        the indented list above in combination with bold on the `item {1,2}`. 

						;item 1
						: definition 1
						;item 2
						: definition 2-1
						: definition 2-2

  				[Bold and italic]
					  '''''bold & italic'''''

				  [Escape Wiki Markup]
					  <nowiki>no ''markup''</nowiki>


		    [Section Headings of different levels]
				== Level 2 ==
				=== Level 3 ===
				==== Level 4 ====
				===== Level 5 =====
				====== Level 6 ======

			  [Horizontal rule]
		    To add a horizontal rule, a horizontal dividing line, typically a pixel to five pixels high.
	      Using vanilla HTML, the same element could be created with the `<hr>` tag.

				Text before
				----
				Text after

		    [Bullet list]
				* Start each line
				* with an [[Wikipedia:asterisk|asterisk]] (*).
				** More asterisks give deeper
				*** and deeper levels.
				* Line breaks <br />don't break levels.
				*** But jumping levels creates empty space.
				Any other start ends the list.

				* combine bullet list
				** with definition 
				::- definition
				** creates empty space

				* combine bullet list
				** with definition 
				*:- definition
				** without empty spaces

				*bullet list
				:- definition
				:* sublist that doesn't create empty
				:* spaces after definition


#### Presentations
Multiverse OS is a complex operating system and will need to be introduced to the public by way
of technical talks, and other mediums. Early focus on community written guides and presentations
will make introduction to using, customizing Multiverse OS easier and consistent. w

	[revealgo][https://github.com/yusukebe/revealgo]
	revealgo is a small web application for giving Markdown-driven presentations implemented in
	Go! The revealgo command starts a local web server to serve the your markdown presentation
	file with reveal.js


=================================================================================================
=================================================================================================
-------------------------------------------------------------------------------------------------
##== Extra Libraries (Collected, Unsorted from other notes) ===================================##
######################################################################################################
 | |   (_| |__  _ __ __ _ _ __ _   _  | \ | | ___ | |_ ___ ___ 
 | |   | | '_ \| '__/ _` | '__| | | | |  \| |/ _ \| __/ _ / __|
 | |___| | |_) | | | (_| | |  | |_| | | |\  | (_) | ||  __\__ \
 |_____|_|_.__/|_|  \__,_|_|   \__, | |_| \_|\___/ \__\___|___/
=======================================================================================================
--------------------------------\__/-------------------------------------------------------------------

# Multiverse

https://github.com/facebookgo/inmem

https://github.com/ring00/torrentfs

https://github.com/yydesa/torcat/blob/master/torcat.go < best way to connect to tor>
https://github.com/mmcco/mutor - tor client
https://github.com/codekoala/torotator
https://github.com/adrpino/go-requests
https://github.com/ak1t0/flame
!! https://github.com/lostinblue/oniongen-go


https://github.com/mmcloughlin/pearl Tor relay implementation in Golang.

https://github.com/anacrolix/torrent

https://github.com/nogoegst/onionutil

https://github.com/upamune/ed25519
https://github.com/nogoegst/pktconn

https://github.com/reinderien/mimic

https://github.com/upamune/bip39
https://github.com/nogoegst/token/blob/master/token.go

https://github.com/hlandau/dht
https://github.com/nogoegst/rand
## FS VirtFS
https://github.com/nogoegst/pickfs/blob/master/pickfs.go

======================
## Qemu KVM
https://github.com/clearcontainers/runtime/blob/master/docs/developers-clear-containers-install.md + 
**https://github.com/digitalocean/go-qemu/** even uses the libvirt lib
https://github.com/hyperhq/runv = this runs on the HOST, it uses QEMU to run conainters inside of. https://github.com/hyperhq/runv/blob/master/cli/main.go It supports custom INIT, 



https://github.com/opencontainers/runc/blob/master/libcontainer/standard_init_linux.go

https://github.com/driusan/dainit 

TODO: Using the tree based ephemeral key system, generate keys that expire. Then generate this key and only allow input using THIS key. Encrypt all output going through standard out through this key!.

This can be easily done by just modifying ther init (standard_ijnit_linux). This will get us much greater security inside the vms. 

ONLY THe controller VM keys should be allowed to TURN OFF, TURN ON or CONTROL the other VMs. THIS TAKES AWAY CONTROL FROM THE HYPERVIISOR!!!

_BESDIES THE INIT, THE BIOS CAN BE FUCKING SET!_[ <MEGA> ]

Hand ALL visuals to the CONTROLELR VM. 

MODIFY THE CLEAR CONTAINERS LINUX KERNEL FURTHER! REMOVE THE ABILITY TO INPUT ANYTHING INTO THE CONSOLE! MODIFY THE INIT TO START THE CONTROLLER VM AND PASS THE VIDEO AND UDB.


**PORTALGUN** 
TREE PUBSUB https://github.com/cloudfoundry/go-pubsub

[pkg manager]
https://github.com/smira/aptly 

https://github.com/gliderlabs/ssh
[protalgun][https://github.com/ctdk/goiardi] in memory chef in go, cut this down and modify it to take a simpler yaml/json file. 
other options https://github.com/rdeusser/overseer, terraform
https://github.com/KAllan357/puck - anotehr chef, but super simple
https://github.com/vektra/tachyon - 
https://github.com/sudharsh/henchman
https://github.com/cloudpac/CloudCore
https://github.com/kevinburke/ansible-go
https://github.com/ansigo/ansigo
https://github.com/smunilla/qi-engine
https://github.com/charles-d-burton/golem

https://github.com/bhcleek/packer-provisioner-ansible

>> Potentially modify this to be the cookbook for using protalgun to spin up a VM
https://github.com/hyperhq/libcompose
**https://github.com/hyperhq/hyperd/tree/master/examples**

https://github.com/18F/hmacauth
**Portalgun stores these configs in a package manager that is distributed**

https://github.com/docker/go-p9p
https://github.com/progrium/go-githubfs

https://github.com/camlistore/camlistore

*fuzzy*
https://github.com/jamesboehmer/twocents
*circular buffer*
https://github.com/maruel/circular
*memfs*
https://github.com/ranveerkunal/memfs
https://github.com/victorcwai/VDFS
*maybe useful dbs&*
https://github.com/awnumar/memguard

https://github.com/karlseguin/nabu
https://github.com/pomkac/mnemonic
https://github.com/hickeroar/gobayes
https://github.com/hamaxx/goindex
https://github.com/t3rm1n4l/nitro
https://github.com/mars9/ramfs
https://github.com/amenzhinsky/go-memexec
https://github.com/e-tothe-ipi/inmemfs
https://github.com/bbengfort/memfs In memory file system that implements as many FUSE interfaces as possible. 
https://github.com/hashicorp/go-memdb + !@! https://github.com/hinshun/txnstore
https://github.com/neo5g/neo5g

_remeber to look at multipath moujhnting and NFS network file system**

https://github.com/abcum/vtree A versioned radix tree library implementation for Go 

_[using this, https://github.com/tidwall/evio, delya can be changed INSIDE the event function. So we could write the type of async/realtime ticker code used in RTOS. Caclualte when the next thing should fire and change the delay to be until that happens.


**OHT**
https://github.com/draganm/immersadb - ImmersaDB is an embedded database written in Go.
    Transactional (ACID)
    Append-Only using persistent data structures
    Blazing fast: 3-4K transactions per second on an SSD
    Memory-Mapped - memory paging used to cache reads, leveraging the whole available ram without needing to allocate it
    Constant resident memory requirements for both reading and modifying data, independent on data size
    Stores arbitrary deep tree of hashes and arrays as nodes and Data as leaves
    Event model for listeners of changes for sub-trees
_<Could be perfect fot the distributed ledger that maintains the codebase>>

**SCRAMBLESHELL**
https://github.com/pylls/gosmt - file integrity checking
https://github.com/daniel-ziegler/merklemap

https://github.com/abcum/vtree A versioned radix tree library implementation for Go 
https://github.com/pbanos/botanic

https://github.com/achim-k/go-vebt-eval/blob/master/vebt_eval.go

https://github.com/mozu0/wltree - with wait and select
++https://github.com/arwx/splaytree++
https://github.com/seehuhn/classification
++https://github.com/ckennelly/watches++ NWAY DIFFER

https://github.com/acronis/notary-verifyhash
CLI for Acronis Notary to verify hash in Merkle Patricia Tree root/proof
## Merkel with streams upport
https://github.com/vbatts/merkle
https://github.com/tidwall/pair-rtree/tree/master/3d

https://github.com/wilfreddenton/merkle

https://github.com/arbrain/abrain
https://github.com/ktodorov/go-summarizer
___https://github.com/mosuka/indigo__

https://github.com/DataDrake/filesdb

[desktop] https://github.com/alexozer/metadesk < this but the top desktops scroll left to right in a loop

Need a .bash environment that is carried to other computers and auto deletes itself after expiry.

**Then implement the scramble shell** providing a basic Key Value store. Notes. Full HTTP API for the FS and all stats. Built it converters. Built in text analysis, that suggests how to write different.

Very simple port forwarding using OHT

**Controll all systems from 1 terminal. Support for channels, gorups, etc. ALways being connected to the other machiens as if the entire cluster is a single machine.**


https://github.com/EndFirstCorp/pdf2txt
https://github.com/EndFirstCorp/rtf2txt
https://github.com/EndFirstCorp/doc2txt

implement libinput in go and support multi user coop VMs
https://github.com/progrium/termshare [term coop with OHT][https://github.com/agupta666/raptor]

shamirs share key data with your friends, pay others to host your VMs

https://github.com/dsprenkels/sss-go best 

https://github.com/saurabhkb/Yank

https://github.com/dist-ribut-us/merkle storage

**bittorrent sync** + git with signing. + TorrentFS or GitFS. Meta data is transfered. But files must be right clicked and told to be transfered to take up space unless the settings for the folder are set to download everything (not default settting)
https://github.com/jackpal/bencode-go
https://github.com/declan94/secret-share

https://github.com/lemmi/ghfs

**DHT** https://github.com/rainer37/Rnet/blob/master/dht/dht_general.go

[scramble shell][https://github.com/seletskiy/go-android-rpc] - SCRAMBLE MUST IMPLEM<ENT SEEEMLESS> MOBILE INTERGRATION!
https://github.com/jandre/passward

**Collective knowledge base, store all books, stores all, LAB reports.**
https://github.com/stchris/books

https://github.com/ncode/pretty - yaml configured multi shell control
## KVM INSTALL FROM FS
https://github.com/zhujintao/kyum/blob/master/main.go


## FS
https://github.com/polydawn/rio
## Userspace
[netstack]
https://github.com/dutchcoders/netstack <real>
https://github.com/google/netstack
https://github.com/hsheth2/gonet listenIP: ipv4.IPAll, <- real usespace >
https://github.com/joshlf/net
https://github.com/jdewald/gotcp
https://github.com/YaoZengzeng/yustack

https://github.com/songgao/ether
https://github.com/songgao/packets

https://github.com/dop251/nbd
https://github.com/dop251/buse
[tcm]
https://github.com/coreos/go-tcmu
======================
#qmp
https://github.com/digitalocean/go-qemu
https://github.com/zchee/go-qcow2

https://github.com/ByteArena/schnapps
    To be able to use KVM, you need to have the proper unix capabilities.
    If you are using the bridge NIC, you need to use the host network mode.

Features
    DNS server (only A records are supported) (doc)
    QMP server
    Random MAC address generator (doc)
    Uses libvirt events
    Manages a KVM process, its lifecycle and its configuration (doc)
    Simple VM scheduler with cluster health monitoring (doc)
    Metadata server (doc)
    Custom DHCP server (Ipv4 only) (doc)



https://github.com/0xef53/qmp-shell
## VM UI
https://github.com/fotcorn/go-kvm-web
https://github.com/scr34m/go-kvm-ui
https://github.com/farazfazli/CyanOcean
======================
## gnome search
https://github.com/yamnikov-oleg/projektor
======================
https://github.com/xyproto/battlestar = shrinks resulting executables
======================
https://github.com/latentgenius/vcardgen

https://github.com/xtaci/sjson
**OHT**

https://github.com/18F/hmacauth

https://github.com/xtaci/transocks
**Share videos, words, videos using bittorent sync!**
https://github.com/vole/vole web ui!
** TORRENT **
https://github.com/jackpal/bencode-go

**Tree comments/categories**
https://github.com/hit9/htree

https://github.com/spouk/tree_comments_go
[categories]
https://github.com/mrsinham/catego
## Merkel with streams upport
https://github.com/vbatts/merkle
**LAB**
https://github.com/ether/etherpad-lite
https://github.com/declan94/ng.simditor


[electrical component UI]
https://github.com/qeda/qeda
[inventory system]


https://github.com/justinmandac/simple-inventory
https://github.com/Orkiv/Inventory-go-client

https://github.com/starshipfactory/starsto
[drag and drop iamge sharing]
https://github.com/avesanen/dropdead
[reverse equipment]
https://github.com/justiniso/golock
[audio]
https://github.com/stevemurr/portaudio-streamer
https://github.com/stevemurr/audio-share
**Hackwave**
[meeting times]
https://github.com/theothertomelliott/meetingtime

https://github.com/rickar/cal
https://github.com/JulienBreux/rrule-go
https://github.com/EndFirstCorp/calendar

https://github.com/Zenithar/go-calendar= supports localz

[biz days]
https://github.com/mchudgins/bizdays
[schedule compare]
https://github.com/saulhoward/calendar
## Bit Client
https://github.com/btcsuite/btcd
https://github.com/GridProject/signer

https://github.com/skip2/go-qrcode/
## Hackwave Servicse
https://github.com/syncthing/syncthing
## Trading
https://github.com/Akagi201/cryptotrader
## Target
https://github.com/NoteGio/openrelay

## Multiverse/Hackwave Blog
https://github.com/kentaro-m/md2confl
https://github.com/josephlewis42/paraphrase

https://github.com/lethain/icarus - has beleve full text search
https://github.com/keitax/textvid/
https://hackmd.io/
[latex for scientific documents]
https://www.sharelatex.com/
## Search Engnine
https://github.com/arbrain/abrain
https://github.com/josephlewis42/paraphrase
https://github.com/fengvyi/Document-Statistics
https://github.com/yafeer/excavator
https://github.com/serkas/salias
## FS stream
https://github.com/hayeah/gowatch



https://githuGone is a wiki engine written in Go. It's

    KISS,
    Convention over Configuration and
    designed with Developers and Admins in mind.

With Gone, you can

    display Markdown, HTML and Plaintext straight from the filesystem.b.com/fxnn/gone

# Email
https://github.com/go-gomail/gomail


https://github.com/mailhog/MailHog
https://github.com/mailhog/mh2
https://github.com/mailhog/storage - MongoDB storage
https://github.com/mailhog/MailHog-UI - JSON ui
https://github.com/mailhog/MailHog-Server
https://github.com/mailhog/data
https://github.com/mailhog/mhsendmail
https://github.com/mailhog/smtp
https://github.com/mailhog/backends
https://github.com/mailhog/imap
https://github.com/mailhog/MailHog-MTA

https://github.com/TheCreeper/go-pop3
https://github.com/gnsx/go-smtp
["cloud" features]
https://github.com/cozy/cozy-stack
[anon temp data upload]
https://github.com/root-gg/plik

## VOIP
https://github.com/keroserene/go-webrtc



https://github.com/stevemurr/soxy
## KCP = TCP hybrid that has best of UDP with TCP features
https://github.com/xtaci/kcp-go
https://github.com/golang/snappy
*Control messages like SYN/FIN/RST in TCP are not defined in KCP, you need some keepalive/heartbeat mechanism in the application-level. A real world example is to use some multiplexing protocol over session, such as smux(with embedded keepalive mechanism), see kcptun for example.* [FEC(Forward Error Correction) Support with Reed-Solomon Codes][Control messages like SYN/FIN/RST in TCP are not defined in KCP, you need some keepalive/heartbeat mechanism in the application-level. A real world example is to use some multiplexing protocol over session, such as smux(with embedded keepalive mechanism), see kcptun for example.]
https://github.com/btcsuite/snappy-go

## RTMP SNMP SRS
https://github.com/ossrs/go-oryx-lib
https://github.com/ossrs/go-oryx
## Hackwve Market
[review anlaysis]
https://github.com/mikeflynn/quackerjack
## OCR
https://github.com/maxim2266/go-ocr
https://github.com/ghetzel/shmtool

## RTP
https://github.com/wernerd/GoRTP
https://github.com/antongulenko/RTP
[dump packets and analize]
https://github.com/yyd01245/pcap_rtpParse
https://github.com/hdiniz/rtpdump

[zrtp]
https://github.com/wernerd/ZRTPCPP

## Reed Soloman
https://github.com/templexxx/reedsolomon
======================
## LUA VM
https://github.com/yuin/gopher-lua
https://github.com/Shopify

## JVM
https://github.com/richdyang/jago
======================
## Linux Virtual Server LVS
https://github.com/nanobox-io/golang-lvs
## Userspace Networking
https://github.com/chzyer/next
https://github.com/nathanjsweet/zsocket

https://github.com/vishvananda/netlink

https://github.com/rancher/vm-net/blob/master/tapclient/main.go

https://github.com/mdlayher/ethernet
https://github.com/mdlayher/vsock/blob/master/conn_linux.go

[kernel space]
https://github.com/typetypetype/conntrack
[netlink]
https://github.com/mdlayher/netlink


[tunneling]
https://github.com/xtaci/kcptun
[websockets]
https://github.com/orijtech/wsu
## Tree
https://github.com/timtadh/file-structures/blob/master/treeinfo/info.go
[fs]
## HTTP FS Interaction
https://github.com/dolftax/summer
## VFS
https://github.com/rainycape/vfs
## Memory Mapped FS B+ Tree
https://github.com/timtadh/fs2 
## Stats 

https://github.com/kelseyhightower/terminus
https://github.com/shirou/gopsutil
https://github.com/c9s/goprocinfo
https://github.com/zcalusic/sysinfo
## Disk Images Building
https://github.com/solus-project/USpin
https://github.com/linuxkit/linuxkit
https://github.com/fdev31/archx
## Qemu KVM
https://github.com/clearcontainers/runtime/blob/master/docs/developers-clear-containers-install.md + 
**https://github.com/digitalocean/go-qemu/** even uses the libvirt lib


https://github.com/ZeroPage/vm-manager - Setup with YAML. Would make nice start
https://github.com/google/vmregistry/

https://github.com/ganshane/zvirt *libvirt agent to get data from server*

https://github.com/bakins/libvirt-http/blob/master/main.go *more xml for controlling server*

https://github.com/cha87de/kvmtop < VERY NICE PLACE TO START >

**Portal Gun**
https://github.com/clearcontainers/runtime/blob/master/docs/developers-clear-containers-install.md + 


https://github.com/coreos/go-systemd
https://github.com/godbus/dbus

[fifo - what all /dev/ deviecs use]
https://github.com/containerd/fifo

[fs]
https://github.com/containerd/btrfs

[bootloader]
[custom kernel]
[initramfs]
[init]

https://github.com/opencontainers/runc/blob/master/libcontainer/standard_init_linux.go

https://github.com/driusan/dainit 
https://github.com/pablo-ruth/go-init - not legit, not enough there.

https://github.com/google/fscrypt/blob/master/pam/login.go

https://github.com/HouzuoGuo/cryptctl *will need to use this to decrypt the drive.*

[Access]
https://github.com/libvirt/libvirt-console-proxy
Websockets console proxy for VNC, SPICE and serial consoles
https://github.com/fromanirh/netdata-virt *view graphs of server stats*

https://github.com/csu/logrus_http *output logs from servers to mainserver via HTTP*

**TODO**
* Modify Linux Kernel or some other version of the kernel to only accept console input wtih signed paramters and encrypt all output.

## Containers

[runc]
*https://github.com/containerd/cgroups* Most intersting

https://github.com/containerd/go-runc
https://github.com/containerd/containerd

https://github.com/opencontainers/runc **key interesting because it even has mounting of FD /dev/console**

## Crypto
[new tor spec]
https://github.com/tendermint/ed25519
[postquantum]
https://github.com/orijtech/prunehorst
https://github.com/orijtech/haraka

## Clear Containers Hardware virtualized containers
https://github.com/clearcontainers/agent


https://github.com/containers/storage/blob/master/drivers/vfs/driver.go
https://github.com/containers/virtcontainers
## Hiding Hypervisor
https://github.com/0xkrishnakumar/Virtualization/blob/master/hypervisor_presence/hypervisor.go






## Compatiability

[open stack]
https://github.com/dicot-project/dicot-api REST APIs to enable OpenStack compatible clients to use KubeVirt and Kubernetes 

-------------------------------------------------------------------------------------------------
=================================================================================================

# Multiverse OS
[TODO: Clean up and merge libraries]

[NOTE: LOOK AT BELLINA FOR GUI!!]

[firewall](https://github.com/Gouthamve/go-firewall)
A firewall using nfqueue (not iptables), this makes it pretty low level or atleast very high power. 

[css only fileicons](https://github.com/picturepan2/fileicon.css)
#### FS
[memfs](https://github.com/zbiljic/memfs)
[tiedot](https://github.com/HouzuoGuo/tiedot)
nosql database, embedded

[bolt-mount](https://github.com/bazil/bolt-mount)
bolt fs

[hookfs](https://github.com/osrg/hookfs)
usermode hookable fs
## systemd
https://github.com/coreos/go-systemd
## Ruby Config
https://github.com/k0kubun/itamae-go
#### Firewall
https://github.com/Safing/safing-core
#### Diposable Email
[inbucket](https://github.com/jhillyerd/inbucket)
#### Cache
https://github.com/abcum/cachr - *best cache library*
============================================
# GUI
https://github.com/BurntSushi/xgbutil
https://github.com/golang-ui/nuklear

## WEbui rendeirng
https://github.com/gu-io/gu
https://github.com/goxjs/glfw
## Beautiful QT UI for go example
https://github.com/dzhou121/gonvim
## Bellina
https://github.com/amortaza/go-bellina
https://github.com/amortaza/go-bellina-plugins
bellina seems to be the most rich UI lib I have seen
https://github.com/amortaza/go-hal-g5
https://github.com/amortaza/go-xel
[state machine](https://github.com/looplab/fsm)
# GUI
https://github.com/BurntSushi/xgbutil
https://github.com/golang-ui/nuklear
## WEbui rendeirng
https://github.com/gu-io/gu
https://github.com/goxjs/glfw
## Beautiful QT UI for go example
https://github.com/dzhou121/gonvim
## Bellina
https://github.com/amortaza/go-bellina
https://github.com/amortaza/go-bellina-plugins
bellina seems to be the most rich UI lib I have seen
https://github.com/amortaza/go-hal-g5
https://github.com/amortaza/go-xel
====================================================================================
## Core Utils                                                                     ##
====================================================================================
**CoreUtils**
https://github.com/shreyaganguly/coreutils - only a few
------------------------------------------------------------------------------------
## Shell
# Go Shell Command Line Prompt
https://github.com/otm/gluash
https://github.com/jharshman/gosh
https://github.com/lfkeitel/lish
https://github.com/desertbit/turban
#console
https://github.com/jroimartin/gocui
https://github.com/ncode/pretty
https://github.com/elpinal/coco3
## Examples
https://github.com/elves/elvish
https://github.com/michaelmacinnis/oh/blob/master/main.go
https://github.com/mvdan/sh
https://github.com/NeowayLabs/nash
https://github.com/GetStream/vg
# term sharing
https://github.com/spolu/warp
====================================================================================
## Device Types                                                                   ##
====================================================================================
**SIMD**
https://github.com/randall2602/gensimd
====================================================================================
## Networking                                                                     ##
====================================================================================
**Websockets**
https://github.com/thehowl/claws
------------------------------------------------------------------------------------

======================================================================
## Compile/Build Install/Live Multiverse OS ISO                     ##
======================================================================
This gives a very brief way to build the ISO
https://github.com/elazarl/customize_tinycorelinux/blob/master/tinycore-repack.sh
======================================================================
## Go Libraries                                                     ##
======================================================================
## Web Application Framework
### Minify
https://github.com/nonoo/html-cruncher HTML cruncher rewrites IDs, classes and names in HTML, CSS and JavaScript files in order to save bytes and obfuscate code. 
## Documentation
https://github.com/src-d/code-annotation
## Input
### Untrusted Inputs
https://github.com/google/wuffs
### User Input
## Community
### Torrent Tracker
https://github.com/crosbymichael/tracker
### Chat
##### Parse Links/IP Address/Etc
https://github.com/DCSO/ioceew
#### SSH
https://github.com/shazow/ssh-chat
### Spatial
https://github.com/dhconnelly/rtreego
## Scramble Suit
### Keyring
#### Key Types
##### Merkle Tree (Trillian)
https://github.com/google/trillian Trillian implements a Merkle tree whose contents are served from a data storage layer, to allow scalability to extremely large trees. 
https://github.com/google/trillian-examples
##### Saltpack (NaCL)
https://github.com/keybase/saltpack
#### Key Manipulation
##### Splitting/Merging/Shamirs
https://github.com/shazow/keyxor
## Console/Terminal
### Readline
https://github.com/chzyer/readline
### Console
https://github.com/containerd/console
### Terminal
https://github.com/pkg/term
## 2D Engine
https://github.com/hajimehoshi/ebiten
## Data/Blobs
### Blob (Binary or otherwise)
https://github.com/gonutz/blob (https://github.com/gonutz/bin2go)
### In Memory FS
https://github.com/dlmc/ids
### Git
https://github.com/src-d/go-git
### Compression
https://github.com/coreos/gzran
## Text
### Index/Full Text Search
https://github.com/go-ego/riot
## FIREWALL
https://github.com/juliengk/go-netfilter (would need a bit of work)
https://github.com/subgraph/go-nfnetlink
https://github.com/subgraph/fw-daemon/blob/master/sgfw/rules.go
https://github.com/subgraph/fw-daemon - using just the firewalld and modifications to make it read a simple config file that is setup for BARE METAL HOST and different one for CONTROLLER and ROUTER, etc.
## Networking
### PROXY
## TUNNEL OVER PING/ICMP
https://github.com/Maksadbek/tcpovericmp
#### Reverse Proxy
https://github.com/fatedier/frp
#### Socks
https://github.com/xiqingping/transproxy (+ Trans)
### Internet Protocol (IP
https://github.com/elsonwu/GOnetstat)
#### Subnets
https://github.com/coreos/ipnets
### SSH
#### SSH Server
https://github.com/gliderlabs/ssh
#### SSH Server OR Client
https://github.com/shazow/go-sshkit
#### Multiplex SSH Commands
https://github.com/crosbymichael/slex
## UI
### Text
#### Progress Bars
https://github.com/mitchellh/ioprogress
## Web Server
### Browser (Window, default browser)
https://github.com/pkg/browser
https://github.com/petermbenjamin/go-open
### Lets Encrypt
https://github.com/go-ego/autotls
### Cookies
https://github.com/gorilla/securecookie
### Protocols
 ### gRPC
   #### Low-memory gRPC Protocol
https://github.com/stevvooe/ttrpc
 ### Websockets
 ### WebRTC
 ### SFTP
https://github.com/pkg/sftp The sftp package provides support for file system operations on remote ssh servers using the SFTP subsystem. It also implements an SFTP server for serving files from the filesystem.
### Embedded Files
https://github.com/elazarl/go-bindata-assetfs
https://github.com/gnoso/go-bindata (Additionally compresses and converts to RO byte slice)
### Proxy
#### Socks
#### MITM
https://github.com/google/martian Outstanding! This project got a lot more mature over the last 2 months.
https://github.com/elazarl/goproxy (https://github.com/elazarl/goproxy2, some changes, may be broken)
## Debian
### Systemd (dbus, journal, machine1, socket activation, unit files)
https://github.com/coreos/go-systemd
## Devices
https://github.com/chzyer/tunnel - Tunnel Device (Does it work without root?)
https://github.com/chzyer/next
https://github.com/google/gopacket
## TUN/TAP Devices
https://github.com/pkg/taptun
## Processes
https://github.com/mitchellh/go-ps
## Visualization
### Network Logging
https://github.com/stefanszasz/network-logger
## Booting
### Initramfs
https://github.com/miekg/dinit
https://github.com/coreos/coreos-cloudinit
### BIOS
### Netboot
https://github.com/google/netboot
## Provision/Devops
### Example (Ignition), has several good useful components
https://github.com/coreos/ignition
### Composing/Recipies/Configuration (YAML preferred)
https://github.com/coreos/coreos-cloudinit (Deprecated but good code)
## Virtual Machines (VMs)
### Hardware Virtualized
### Containers
#### Docker (engine-API)
https://github.com/docker/engine-api/tree/master/types
https://github.com/moby/moby
https://github.com/moby/moby/tree/master/clientw
#### CGroups
https://github.com/containerd/cgroups
#### OCI
##### containerd
https://github.com/containerd/containerd
##### kubes
https://github.com/kubernetes/kops
https://github.com/kubernetes/kubernetes
##### Kubes release/image buidler
https://github.com/kubernetes/release -- kubes release, some useful code for image builder
https://github.com/kubernetes/kube-deploy/blob/master/imagebuilder/templates/1.8-stretch.yml
##### runc
https://github.com/opencontainers/runc
##### rkt
https://github.com/rkt/rkt
### WebUI
https://github.com/kubernetes/dashboard
## WebKit GO Bindings 
https://github.com/abcum/webkit
**THIS HAS EASY ABILITY TO DISABLE JAVASCRIPT!**
## Users/Authentication/Authorization
### TOTP/HMAC/OTP/HOTP
https://github.com/stefanszasz/totp
### Unique IDs
#### Custom Obstuficated
https://github.com/c2h5oh/hide
https://github.com/wjanssens/flexid-go
#### Human Readable 
https://github.com/FabianWilms/GoReadableID Go-Library to generate gfycat-like human readable IDs 
https://github.com/kusubooru/tklid 
Anime-flavored, tickletastic, suitably random and reasonably unique human readable ids
https://github.com/satta/whatsmyname Map machine IDs to noun pairs 
### OpenID 
#### OpenID Server
https://github.com/coreos/dex
#### OpenID Client
### Hawk Authentication Protocol
https://github.com/coreos/hawk-go
### Diameter Protocol
## Locking Down
### Netfilter Queue
https://github.com/OneOfOne/go-nfqueue
### IP Tables
https://github.com/coreos/go-iptables - Looks good, check out the test for the examples beacuse tehre are none included
### IDS Rules
https://github.com/google/gonids
## Files
### Category/Type/Magic Numbers
### Xattr (Extra attributes)
https://github.com/pkg/xattr
## Filesystem (FS)
### GitFS 
https://github.com/hanwen/gitfs
### MTPFS (Android MTP)
https://github.com/hanwen/go-mtpfs
### Generic FUSE (go-fuse)
https://github.com/hanwen/go-fuse native bindings for the FUSE kernel module.
https://github.com/hanwen/go-fuse/tree/master/example
Examples include zipFS, multizipFS, memFS, unionFS, statFS
### BTRFS
https://github.com/containerd/btrfs 
### AUFS
https://github.com/containerd/aufs
### Transport Agnostic FS
https://github.com/containerd/continuity
### Read-only, lazy loaded, Git repo FS 
https://github.com/google/slothfs
### Distributed, Encrytped, Branching FS
https://github.com/google/mawfs
### LUKS / Full Disk Encryption
https://github.com/google/fscrypt
### Keybase filesystem (KBFS), built in shamirs
https://github.com/keybase/kbfs
  - https://github.com/keybase/go-framed-msgpack-rpc (Used in above, has useful code)
### Pure GO blkid (goblkid)
https://github.com/elazarl/goblkid
### Multiverse FS
https://github.com/osrg/hookfs
### NFS
Using NFS with Quartermaster (a yaml based volume manager)
https://github.com/coreos/quartermaster/blob/master/pkg/storage/nfs/README.md
## Database
### Generic ORM
#### MongoDB, MySQL, Postgres, SQLite, QL, MSSQL
https://github.com/upper/db
### Embedded
#### File
https://github.com/coreos/bbolt
#### Distributed/Raft Network Database
https://github.com/coreos/etcd
https://github.com/ha/doozerd SERVER A consistent distributed data store. 
https://github.com/coreos/doozer CLIENT
## Monitoring
### Profile (CPU, Memory)
https://github.com/pkg/profile
### Error Information Gathering
https://github.com/coreos/mayday
## PCIE
https://github.com/aoeldemann/gopcie
## High Latency
https://github.com/klauspost/reedsolomon
## Devices (FIFO,SHM,TUN,TAP...)
### FIFO
https://github.com/containerd/fifo
### OpenVPN
https://github.com/chzyer/openvpn-ctl
### Port Level Control
https://github.com/chzyer/tracp Traffic Control in Port level for OSX/Linux 
### Virtual Network Devices (Host<->VM connections; Vsock)
https://github.com/google/testimony Testimony is a single-machine, multi-process architecture for sharing AF_PACKET data across processes. This allows packets to be copied from NICs into memory a single time. Then, multiple processes can process this packet data in parallel without the need for additional copies.
https://github.com/osrg/goplane GoPlane is an agent for configuring linux network stack via GoBGP
### Network Routers
https://github.com/osrg/libovsdb
https://github.com/google/gnxi (gRPC Network Management/Operations Interface tools)
### Netstack
https://github.com/google/gopacket
### VPNKit
https://github.com/moby/vpnkit  A toolkit for embedding VPN capabilities in your application 
## PCAP Based Instrusion Detection
### Writing Packets
https://github.com/google/stenographer Stenographer is a full-packet-capture utility for buffering packets to disk for intrusion detection and incident response purposes. It provides a high-performance implementation of NIC-to-disk packet writing, handles deleting those files as disk fills up, and provides methods for reading back specific sets of packets quickly and easily.
### PCAP Analyzer
https://github.com/jzaeske/pcap-analyser
https://github.com/heartszhang/pcapfilter (utils)
https://github.com/aoeldemann/pcaptools
https://github.com/tgogos/gopacket_pcap  Go examples making use of the gopacket / pcap libraries... 
### NFQueue Analysis
https://github.com/tgogos/gopacket_nfqueue
https://github.com/aoeldemann/pcaptools
### ARP Poisoning Detection
https://github.com/pleycpl/yahya
### Packet Analysis WebUI
https://github.com/jasonish/evebox
## Passive Protection
### Webserver
#### SSL
https://github.com/timewasted/go-check-certs
----------------------MULTIVERSE NETWORKING------------------------------

============================================================================
# Multiverse OS
============================================================================
**NOTES** [RANCHER & Kubes & Moby Design Comparison And Notes]
 [#][1] Rancher uses short polling regularly, this is probably a massive resource
sink. Make sure to avoid ALL of that shit.
 [#][2] Rancher is not very secure, doesn't wipe memory of sensitve keys, and 
much more can be done to increase the security
[Seriously Why would you short poll this?]
pollInterval, err := strconv.Atoi(pollIntervalStr)
	if err != nil {
logrus.Fatalf("Failed to convert CERTS_POLL_INTERVAL %v", err
===================================================================================
[!][Provisioning 
===================================================================================
**FEATURES**
  * Ability to set a timer on files to delete them after x time. (https://github.com/baopham/godestroy)
  * Versioned/Branched filesystem, time dialation allowing the movement forward
and backward through time. 
  * Passive Defense: (1) Integrity checking all files, (2) virus scan files, 
(3) Automatically remove metadata from images, GPS coordinates, (4) Auto-encrypt
photos once taken on phone, (5) scan vulernabilities/CVE for own computers, 
check for updates, check open ports.
  * (Provide FS devices) Transcode (Codecs), Transform, Modify, etc: (1) Style text like an identity
with established pattern, (2) Convert to base64, checksum, etc, (3)
  * ALL functionality should be accessible by HTTP API (FS, devices, etc)
-----------------------------------------------------------------------------------
===================================================================================
**LIBS**
https://github.com/petermbenjamin/go-open - Opens files
[BUILDS:TASK RUNNER/CONTINOUS INTERGRATION]
https://github.com/tockins/realize
Enhance your workflow by automating the most common tasks and using the best performing Golang live reloading.
[BUILDS:LINTING]
[BUILDS: TEST + MOCKING + BDD]
https://github.com/onsi/ginkgo
[BUILDS:COMPILING, OBSTUFICATE, SIGN, CHECKSUM, AND RELEASE]
===================================================================================
## USER INTERFACE
-----------------------------------------------------------------------------------
**XPRA Screen Sharing**
[!][Xpra/Xephyr (Pass screen over SSH or TCP)]
https://github.com/zephyrproject-rtos/zephyr
**`Looking Glass` Screen Sharing**
[!][Looking glass; Adding IVSHMEM Device to VM]
Adding IVSHMEM Device to VM

Looking glass works by creating a shared memory buffer between a host and a guest. This is a lot faster than streaming frames via localhost, but requires additional setup.

With your VM turned off open the machine configuration
		$ virsh edit [CONTROLLER_VM]

**Building My Own KVM Screen Relayer**
The "LookingGlass" project uses directX to grab the screen, it would be better to use the
open source alternative [Vulkan](https://github.com/vulkan-go/vulkan)
### GUI FRAMEWORKS
[C++ GUI Framework][Go bindings for a mature C++ GUI framework]
https://github.com/Armored-Dragon/go-imgui
### JAVASCRIPT
#### 3D Rendering
[three.js](https://github.com/mrdoob/three.js)
#### Graphs (JS)
[sigma.js](https://github.com/jacomyal/sigma.js)
A very nice distributed/decentralized graphs that could be used to map the network
======================================================================================
````
<devices>
    ...
  <shmem name='looking-glass'>
    <model type='ivshmem-plain'/>
    <size unit='M'>32</size>
  </shmem>
</devices>
````
For example, in case of 1920x1080

	1920 x 1080 x 4 x 2 = 16,588,800 bytes
	16,588,800 / 1024 / 1024 = 15.82 MB + 2 = 17.82

The result must be rounded up to the nearest power of two, and since 17.82 is bigger than 16 we should choose 32

Next create a script to create a shared memory file.

`/usr/local/bin/looking-glass-init.sh`

	#!/bin/sh

`touch /dev/shm/looking-glass`
`chown user:kvm /dev/shm/looking-glass`
`chmod 660 /dev/shm/looking-glass`

**Hide KVM** [Because NVIDIA is a terrible company, who actively sabotages their customers for short term profit]
You also might have to use the `-cpu host,kvm=off` parameter to forward the host's CPU model info to the VM and fool the virtualization detection used by Nvidia's and possibly other manufacturers' device drivers trying to block the full hardware usage inside a virtualized system. 

Next, modify the libvirt configuration:

	$ virsh edit [vmname]
	<domain type='kvm'>

is changed to:

	$ virsh edit [vmname]
	<domain type='kvm' xmlns:qemu='http://libvirt.org/schemas/domain/qemu/1.0'>

**COMMAND LINE TOOL**
https://github.com/subchen/go-cli
go-cli is a package to build a CLI application. Support command/sub-commands.

###################################################################################
-----------------------------------------------------------------------------------
## NETWORKING
[!][IP]
  * Assign subnet, get IP
  https://github.com/ehazlett/circuit (Assigns, routes, ip, subnet, bridge)
[!][Libcamoflauge, Change MAC address, but do so realistically]
[!][Map the network, find other services: DNS-SD (DNS based service discovery)) allows for specific service discovery]
[ZeroConf](https://github.com/grandcat/zeroconf) is a pure Golang library that employs Multicast DNS-SD for
  * browsing and resolving services in your network
  * registering own services
in the local network.

[!][DNS (mDNS (multicast DNS lookups)]
https://github.com/rancher/plugin-manager/blob/master/events/start_handler.go
 **TODO**
   * Support for both IPv6 and IPv4
   * Send multiple probes (exp. back-off) if no service answers (*)
   * Timestamp entries for TTL checks
   * Optimize DNS lookup requests by rotating and measuring
   * Segregate DNS requests by IDENTITY
   * Randomize DNS request routes to avoid fingerprinting
   * Need to allow DNS over HTTPS (and other backdoor methods)
   * IP over DNS 
   * Compare new multicasts with already received services
https://github.com/grandcat/zeroconf
   * IPSEC https://github.com/rancher/ipsec/blob/master/main.go
[!][ARP, need defense against arp poisoning!]
https://github.com/rancher/ipsec/blob/master/arp/arp_proxy.go
// ARPTableWatcher checks the ARP table periodically for invalid entries
// and programs the appropriate ones if necessary based on info available
// from rancher-metadata
[!][Firewall]
Racher example of using iptables and DNS
https://github.com/rancher/plugin-manager/blob/master/hostnat/watcher.go
https://github.com/rancher/metadata/blob/master/types/service.go iptables
https://github.com/rancher/metadata/blob/master/package/start.sh
**General Firewall Libraries**
https://github.com/Fullrate/fwmerge - nice looking
https://github.com/imgurbot12/goaway
https://github.com/seletskiy/hastur - decent iptables cli
*ipv6 tables* https://github.com/Project0/whaleguard
[!][Network Manager]
https://github.com/rancher/plugin-manager/blob/master/network/manager.go
https://github.com/rancher/plugin-manager/blob/master/network/local.go
[!][Network Bridge]
https://github.com/rancher/plugin-manager/blob/master/vethsync/utils/utils.go
[!][Add routes for every device that needs it]
https://github.com/rancher/plugin-manager/blob/master/routesync/watcher.go
[!][Proxy all network traffic of VM, by IDENTITY]
https://github.com/rancher/metadata/blob/master/k8sproxy/proxy.go
--------------------------------------------------------------------------
## VIRTUAL MACHINE
[!][Build VM object]
Below are some examples of data that should be stored in the VM object
https://github.com/rancher/metadata/blob/master/content/store.go
https://github.com/rancher/metadata/blob/master/types/container.go
[!][Create/Activate VM]
https://github.com/rancher/go-machine-service/blob/master/handlers/create.go
https://github.com/rancher/plugin-manager/blob/master/network/state.go
[!][Build/Register VM]
https://github.com/rancher/go-machine-service/blob/master/handlers/config.go - some useful snippets
https://github.com/rancher/go-machine-service/blob/master/handlers/commands.go - get state, check if exists,
[!][State/Status/Health VM]
https://github.com/rancher/go-machine-service/blob/master/dynamic/driver.go
	schemaRoles = []string{"service",
		"member",
		"owner",
		"project",
		"admin",
		"user",
		"readAdmin",
		"readonly",
		"restricted"}
[!][Setup VM REST API for controlling the VM]
https://github.com/rancher/metadata/blob/master/server/server.go
--------------------------------------------------------------------------
## VM VOLUME
[!][Can we get multi layer (layer0, layer1) to work?]
--------------------------------------------------------------------------
## ACCOUNT & AUTHORIZATION
[!][Setupw or build SSH keys]
  * Make the key deterministic based on `master key`
https://github.com/rancher/rke/blob/master/hosts/tunnel.go
[!][Key/Value Database]
Using separate buckets, can store, `aliases`, `snippets`, `keys`, `notes`,
`urls`, `files scheduled to be destroyed`.
 * Must support /autocomplete/
Examples:
https://github.com/baopham/snip 
--------------------------------------------------------------------------
## SYSTEMD
[systemd]
https://github.com/coreos/go-systemd
Go bindings to systemd socket activation, journal, D-Bus, and unit files 
**TODO**
 * Add in basic control functionality: start, stop, reboot, shutdown
 * // Subscribe to signals on the logind dbus
// LockSession asks the session with the specified ID to activate the screen lock.
func (c *Conn) LockSession(id string) {
	c.object.Call(dbusInterface+".LockSession", 0, id)
}
// LockSessions asks all sessions to activate the screen locks. This may be used to lock any access to the machine in one action.
func (c *Conn) LockSessions() {
	c.object.Call(dbusInterface+".LockSessions", 0)
}
// TerminateSession forcibly terminate one specific session.
func (c *Conn) TerminateSession(id string) {
	c.object.Call(dbusInterface+".TerminateSession", 0, id)
}
// TerminateUser forcibly terminates all processes of a user.
func (c *Conn) TerminateUser(uid uint32) {
	c.object.Call(dbusInterface+".TerminateUser", 0, uid)
}
// Reboot asks logind for a reboot optionally asking for auth.
func (c *Conn) Reboot(askForAuth bool) {
	c.object.Call(dbusInterface+".Reboot", 0, askForAuth)
}
 * Add API functionality to go-systemd library so that session and user data can be 
taken from the library and add ability to login.
https://github.com/EclecticIQ/go-systemd/blob/master/login1/dbus.go
 * // To avoid false positives on systems without `pam_systemd` (which is
// responsible for creating user sessions), this function also uses a heuristic
// to detect whether it's being invoked from a session leader process. This is
// the case if the current process is executed directly from a service file
// (e.g. with `ExecStart=/this/cmd`). Note that this heuristic will fail if the
// command is instead launched in a subshell or similar so that it is not
// session leader (e.g. `ExecStart=/bin/bash -c "/this/cmd"`)
 *// IsRunningSystemd checks whether the host was booted with systemd as its init
// system. This functions similarly to systemd's `sd_booted(3)`: internally, it
// checks whether /run/systemd/system/ exists and is a directory.

// PacketConns returns a slice containing a net.PacketConn for each matching socket type
// passed to this process.
//
// The order of the file descriptors is preserved in the returned slice.
// Nil values are used to fill any gaps. For example if systemd were to return file descriptors
// corresponding with "udp, tcp, udp", then the slice would contain {net.PacketConn, nil, net.PacketConn}
func PacketConns(unsetEnv bool) ([]net.PacketConn, error) {
==========================================================================
## GO APPLICATION CONFIGURATION
--------------------------------------------------------------------------
## ENV Configuration
[envconfig](https://github.com/kelseyhightower/envconfig)
Golang library for managing configuration data from environment variables 
*Its much more, it has std errosrs and shit* [This is the same stuff I found myself adding to my configuration library]
Supported Struct Field Types
envconfig supports supports these struct field types:
    string
    int8, int16, int32, int64
    bool
    float32, float64
    slices of any supported type
    maps (keys and values of any supported type)
    encoding.TextUnmarshaler
==========================================================================
## OTHER
--------------------------------------------------------------------------
## Memory Database
https://github.com/kelseyhightower/memkv 
Simple in memory k/v store.
## WGET 
https://github.com/bleenco/go-resumable - Resumable
## Tunnel
https://github.com/bleenco/localtunnel
## Websockets
[wwscc]
https://github.com/wegel/wwscc
Wegel's WebSocket Channel Connector: tunnel tcp over websocket, even when both ends of the connection are behind strict firewall 
## Virtual Network Device
[FIFO Ring Buffer based Queue]
https://github.com/fzakaria/circularfifoqueue

