# Multiverse OS Alpha Installer Build
https://github.com/aelsabbahy/goss
https://github.com/muesli/cache2go
https://github.com/dgraph-io/dgraph
https://github.com/siddontang/ledisdb
https://github.com/syndtr/goleveldb <very nice>
https://github.com/spf13/afero
## httpfs
https://github.com/cznic/httpfs
## simple cotnainer
https://github.com/converseai/simple_container
!!!!!!!!!!!!!!

USES FUCKING MRUBYT O CONFIG CONTAINERS (DOCKER BUT WHO GIVES A FUCK!)
!!!!!!!!!!!!!!!!!!!
BEST SYS INFO https://github.com/zcalusic/sysinfo
https://github.com/u-root/u-root/tree/6fd12df7e9d96fce0acb7ab723b7f129e5b07107/pkg/pci
https://github.com/AdaptiveScale/lxdui
https://github.com/hackwave/goiardi CHEF IN GO
https://github.com/jdmelo/libvirt-go-xml
https://github.com/jdmelo/libvirt-go-xml
https://github.com/twitchyliquid64/bob-the-builder
### Mk devices
if err := syscall.Mkfifo(f, 0700); err != nil && !os.IsExist(err) {
			return nil, fmt.Errorf("mkfifo: %s %v", f, err)
}

if err := syscall.Mkfifo(outPipe, 0644); err != nil {
		return nil, nil, err
}

if err := syscall.Mkfifo(path, syscall.S_IFIFO|0666); err != nil {
		logger.Log.Fatalf("%v", err)
}


const (
	nodev    = unix.MS_NODEV
	noexec   = unix.MS_NOEXEC
	nosuid   = unix.MS_NOSUID
	readonly = unix.MS_RDONLY
	rec      = unix.MS_REC
	relatime = unix.MS_RELATIME
	remount  = unix.MS_REMOUNT
	shared   = unix.MS_SHARED
)

var (
	rliminf  = unix.RLIM_INFINITY
	infinity = uint64(rliminf)
)

// set as a subreaper
func subreaper() {
	err := unix.Prctl(unix.PR_SET_CHILD_SUBREAPER, uintptr(1), 0, 0, 0)
	if err != nil {
		log.Printf("error setting as a subreaper: %v", err)
	}
}

// nothing really to error to, so just warn
func mount(source string, target string, fstype string, flags uintptr, data string) {
	err := unix.Mount(source, target, fstype, flags, data)
	if err != nil {
		log.Printf("error mounting %s to %s: %v", source, target, err)
	}
}

// in some cases, do not even log an error
func mountSilent(source string, target string, fstype string, flags uintptr, data string) {
	_ = unix.Mount(source, target, fstype, flags, data)
}

## Make dev and nodes

		// Mknod creates a filesystem node (file, device special file or named pipe) named path
		// with attributes specified by mode and dev.
		func Mknod(path string, mode uint32, dev int) error {
			return syscall.Mknod(path, mode, dev)
		}

		// Mkdev is used to build the value of linux devices (in /dev/) which specifies major
		// and minor number of the newly created device special file.
		// Linux device nodes are a bit weird due to backwards compat with 16 bit device nodes.
		// They are, from low to high: the lower 8 bits of the minor, then 12 bits of the major,
		// then the top 12 bits of the minor.
		func Mkdev(major int64, minor int64) uint32 {
			return uint32(((minor & 0xfff00) << 12) | ((major & 0xfff) << 8) | (minor & 0xff))
		}


## Make char device by chaining mknod with mkdev
		// make a character device
		func mkchar(path string, mode, major, minor uint32) {
			// unix.Mknod only supports int dev numbers; this is ok for us
			dev := int(unix.Mkdev(major, minor))
			err := unix.Mknod(path, mode, dev)
			if err != nil {
				if err.Error() == "file exists" {
					return
				}
				log.Printf("error making device %s: %v", path, err)
			}
		}

### cmd line
cat /proc/modules | grep hello

cat /proc/cmdline 
to see what things are loaded at boot
--

https://github.com/fabric8io/kansible

https://github.com/1lann/cete kvdb badger
https://github.com/bmeg/arachne badger graph

**Experiment with using btrfs as the multipath networ drive between Multiverse OS virtual machines**
https://github.com/manifoldco/promptui
https://github.com/tyler-smith/go-bip32 - KEY SYSTEM

https://github.com/NVZGUL/NetInterfaceApi
https://github.com/PouuleT/exec-in-net - conceptually important
https://github.com/samalba/buse-go

terminus for hardware stats

https://github.com/mars9/ramfs - 9p but all in memory!

https://github.com/vasi/qcow2
https://github.com/briandowns/aion - rest server cron

############ https://github.com/lxc/distrobuilder

https://github.com/pfactum/xk MACROKERNEL FROM C+KVM VMs

https://github.com/bootchk/rustDevContainers VAGGA RUST CONTAINER

https://github.com/cyphar/initrs RUST INIT FOR CONTIANERS

https://github.com/gsora/lmod KERNEL MODUELS

https://github.com/mgoltzsche/cntnr

https://github.com/huwwynnjones/oci_rs
https://github.com/oracle/railcar
https://github.com/utaal/shipc RUST CONTAINER WTIH ROOTLESS

https://github.com/lastbackend/lastbackend WEBUI
https://github.com/ttrahan/deploy-kubernetes-runcli WEBUI
## WebUI 
https://github.com/abcum/webkit/ - webkit that supports easy ability to disable javascript execution!!!

https://github.com/alice02/runc-with-network

https://github.com/vburenin/consmgr manage console sockets

https://github.com/polydawn/repeatr < timeless >

https://github.com/AkihiroSuda/runrootless
https://github.com/mgoltzsche/cntnr rootless




https://github.com/caseyr003/terraform-oci-workshop
https://github.com/opencontainers/image-tools
https://github.com/anuvu/stacker
https://github.com/openSUSE/umoci
https://github.com/projectatomic/buildah
https://github.com/clearcontainers/osbuilder
https://github.com/containers/image
https://github.com/containers/build

https://github.com/hashicorp/go-reap

IMMUTABLE RADIX https://github.com/hashicorp/go-immutable-radix
https://github.com/hashicorp/go-memdb made with immutable radix
https://github.com/hashicorp/raft-mdb


SOCKADDR https://github.com/hashicorp/go-sockaddr

MEMBERLIST https://github.com/hashicorp/memberlist


RAFT https://github.com/hashicorp/raft
https://github.com/hashicorp/raft-boltdb


PROVISION https://github.com/hashicorp/serf

DBUS https://github.com/guelfey/go.dbus


<transparent proxy>https://github.com/elazarl/goproxy

<BIOS>https://github.com/coreos/seabios

https://github.com/coreos/baselayout

https://github.com/coreos/bcrypt-tool
<IMPORTANT>https://github.com/coreos/go-iptables
<IMPORTANT> https://raw.githubusercontent.com/coreos/baselayout/master/modprobe.d/aliases.conf
https://github.com/jessfraz/netns

func OpenTun(name string) (*os.File, string, error) {
	tun, err := os.OpenFile(tunDevice, os.O_RDWR, 0)
	if err != nil {
		return nil, "", err
	}

	var ifr ifreqFlags
	copy(ifr.IfrnName[:len(ifr.IfrnName)-1], []byte(name+"\000"))
	ifr.IfruFlags = syscall.IFF_TUN | syscall.IFF_NO_PI

	err = ioctl(int(tun.Fd()), syscall.TUNSETIFF, uintptr(unsafe.Pointer(&ifr)))
	if err != nil {
		return nil, "", err
	}

	ifname := fromZeroTerm(ifr.IfrnName[:ifnameSize])
	return tun, ifname, nil
}

https://github.com/coreos/flannel/blob/master/pkg/ip/ipnet.go
Ip marshalling 


// Mkdev returns a Linux device number generated from the given major and minor
// components.
func Mkdev(major, minor uint32) uint64 {
	dev := (uint64(major) & 0x00000fff) << 8
	dev |= (uint64(major) & 0xfffff000) << 32
	dev |= (uint64(minor) & 0x000000ff) << 0
	dev |= (uint64(minor) & 0xffffff00) << 12
	return dev
}

func Pipe(p []int) (err error) {
	if len(p) != 2 {
		return EINVAL
	}
	p[0], p[1], err = pipe()
	return
}


func Mkfifo(path string, mode uint32) error {
	return Mknod(path, mode|S_IFIFO, 0)
}

https://github.com/golang/sys/blob/master/unix/syscall_linux.go
// BindToDevice binds the socket associated with fd to device.
func BindToDevice(fd int, device string) (err error) {
	return SetsockoptString(fd, SOL_SOCKET, SO_BINDTODEVICE, device)
}

func PtracePeekText(pid int, addr uintptr, out []byte) (count int, err error) {
	return ptracePeek(PTRACE_PEEKTEXT, pid, addr, out)
}

func PtracePeekData(pid int, addr uintptr, out []byte) (count int, err error) {
	return ptracePeek(PTRACE_PEEKDATA, pid, addr, out)
}

func PtracePeekUser(pid int, addr uintptr, out []byte) (count int, err error) {
	return ptracePeek(PTRACE_PEEKUSR, pid, addr, out)
}

func PtracePokeText(pid int, addr uintptr, data []byte) (count int, err error) {
	return ptracePoke(PTRACE_POKETEXT, PTRACE_PEEKTEXT, pid, addr, data)
}

func PtracePokeData(pid int, addr uintptr, data []byte) (count int, err error) {
	return ptracePoke(PTRACE_POKEDATA, PTRACE_PEEKDATA, pid, addr, data)
}

func PtracePokeUser(pid int, addr uintptr, data []byte) (count int, err error) {
	return ptracePoke(PTRACE_POKEUSR, PTRACE_PEEKUSR, pid, addr, data)
}

func PtraceGetRegs(pid int, regsout *PtraceRegs) (err error) {
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

func PtraceSetRegs(pid int, regs *PtraceRegs) (err error) {
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}

func PtraceSetOptions(pid int, options int) (err error) {
	return ptrace(PTRACE_SETOPTIONS, pid, 0, uintptr(options))
}

func PtraceGetEventMsg(pid int) (msg uint, err error) {
	var data _C_long
	err = ptrace(PTRACE_GETEVENTMSG, pid, 0, uintptr(unsafe.Pointer(&data)))
	msg = uint(data)
	return
}

func PtraceCont(pid int, signal int) (err error) {
	return ptrace(PTRACE_CONT, pid, 0, uintptr(signal))
}

func PtraceSyscall(pid int, signal int) (err error) {
	return ptrace(PTRACE_SYSCALL, pid, 0, uintptr(signal))
}

func PtraceSingleStep(pid int) (err error) { return ptrace(PTRACE_SINGLESTEP, pid, 0, 0) }

func PtraceAttach(pid int) (err error) { return ptrace(PTRACE_ATTACH, pid, 0, 0) }

func PtraceDetach(pid int) (err error) { return ptrace(PTRACE_DETACH, pid, 0, 0) }

//sys	reboot(magic1 uint, magic2 uint, cmd int, arg string) (err error)

func Reboot(cmd int) (err error) {
	return reboot(LINUX_REBOOT_MAGIC1, LINUX_REBOOT_MAGIC2, cmd, "")
}

func ReadDirent(fd int, buf []byte) (n int, err error) {
	return Getdents(fd, buf)
}

//sys	mount(source string, target string, fstype string, flags uintptr, data *byte) (err error)

func Mount(source string, target string, fstype string, flags uintptr, data string) (err error) {
	// Certain file systems get rather angry and EINVAL if you give
	// them an empty string of data, rather than NULL.
	if data == "" {
		return mount(source, target, fstype, flags, nil)
	}
	datap, err := BytePtrFromString(data)
	if err != nil {
		return err
	}
	return mount(source, target, fstype, flags, datap)
}

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error) {
	var msg Msghdr
	var rsa RawSockaddrAny
	msg.Name = (*byte)(unsafe.Pointer(&rsa))
	msg.Namelen = uint32(SizeofSockaddrAny)
	var iov Iovec
	if len(p) > 0 {
		iov.Base = &p[0]
		iov.SetLen(len(p))
	}
	var dummy byte
	if len(oob) > 0 {
		var sockType int
		sockType, err = GetsockoptInt(fd, SOL_SOCKET, SO_TYPE)
		if err != nil {
			return
		}
		// receive at least one normal byte
		if sockType != SOCK_DGRAM && len(p) == 0 {
			iov.Base = &dummy
			iov.SetLen(1)
		}
		msg.Control = &oob[0]
		msg.SetControllen(len(oob))
	}
	msg.Iov = &iov
	msg.Iovlen = 1
	if n, err = recvmsg(fd, &msg, flags); err != nil {
		return
	}
	oobn = int(msg.Controllen)
	recvflags = int(msg.Flags)
	// source address is only specified if the socket is unconnected
	if rsa.Addr.Family != AF_UNSPEC {
		from, err = anyToSockaddr(&rsa)
	}
	return
}

func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error) {
	_, err = SendmsgN(fd, p, oob, to, flags)
	return
}

func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error) {
	var ptr unsafe.Pointer
	var salen _Socklen
	if to != nil {
		var err error
		ptr, salen, err = to.sockaddr()
		if err != nil {
			return 0, err
		}
	}
	var msg Msghdr
	msg.Name = (*byte)(ptr)
	msg.Namelen = uint32(salen)
	var iov Iovec
	if len(p) > 0 {
		iov.Base = &p[0]
		iov.SetLen(len(p))
	}
	var dummy byte
	if len(oob) > 0 {
		var sockType int
		sockType, err = GetsockoptInt(fd, SOL_SOCKET, SO_TYPE)
		if err != nil {
			return 0, err
		}
		// send at least one normal byte
		if sockType != SOCK_DGRAM && len(p) == 0 {
			iov.Base = &dummy
			iov.SetLen(1)
		}
		msg.Control = &oob[0]
		msg.SetControllen(len(oob))
	}
	msg.Iov = &iov
	msg.Iovlen = 1
	if n, err = sendmsg(fd, &msg, flags); err != nil {
		return 0, err
	}
	if len(oob) > 0 && len(p) == 0 {
		n = 0
	}
	return n, nil
}

----
[Debian image builder]
http://people.linaro.org/~riku.voipio/debian-images/preseed.cfg


Use nspawn containers + debootsrpa to generate Multivere ISo images

https://linux.die.net/man/8/debootstrap
--debian-installer
    Used for internal purposes by the debian-installer
--second-stage
    Complete the bootstrapping process. Other arguments are generally not needed. 


debian-installer


Enter the rootfs chroot and run the second-stage:

    sudo chroot rootfs /bin/bash
    /debootstrap/debootstrap --second-stage

    sudo apt-get install qemu-kvm-extras


  qemu-system-arm \
        -M versatilepb \
        -cpu cortex-a8 \
        -hda rootfs.img \
        -m 256 \
        -kernel vmlinuz \
        -append 'rootwait root=/dev/sda init=/bin/sh rw'

  apt-get install linux-image-versatile


debian-builder 

debian-cd


How:
Create partitions to a loopback file
Format filesystems
Run debootstrap
Install extra packages
Add users and set credentials
Do Hardcoded/Default customizations
Run user specified customizations
Install kernel
Install bootloader


----
   logind.conf, logind.conf.d - Login manager configuration files
      /etc/systemd/logind.conf
       /etc/systemd/logind.conf.d/*.conf
       /run/systemd/logind.conf.d/*.conf
       /usr/lib/systemd/logind.conf.d/*.conf


http://man7.org/linux/man-pages/man1/login.1.html
http://man7.org/linux/man-pages/man7/environ.7.html
http://man7.org/linux/man-pages/man7/locale.7.html
       /etc/security/pam_env.conf
           Default configuration file
       /etc/environment
           Default environment file
       $HOME/.pam_environment
           User specific environment file
       /etc/pam.conf
           the configuration file
       /etc/pam.d
           the Linux-PAM configuration directory. Generally, if this
           directory is present, the /etc/pam.conf file is ignored.
       /usr/lib/pam.d
           the Linux-PAM vendor configuration directory. Files in /etc/pam.d
           override files with the same name in this directory.

       pam_unix - Module for traditional password authentication
       pam_unix.so
  unix_chkpwd is a helper program for the pam_unix module that verifies
       the password of the current user.
       pam_systemd - Register user sessions in the systemd login manager

SYNOPSIS
       pam_systemd.so
DESCRIPTION
       pam_systemd registers user sessions with the systemd login manager
       systemd-logind.service(8), and hence the systemd control group
       hierarchy.

systemd-logind.service **HANDLES MULTISEAT MANAGEMNT**
       /usr/lib/systemd/systemd-logind


#### Multiverse Installer
**Host Machine**

[Host LUKS Encryption]
  [Encryption: Needed for Install]
  (options, both rely on C. So maybe just use C version?)
  https://github.com/martinjungblut/cryptsetup
  https://github.com/kcolford/go-cryptsetup
  [Decryption]
  https://github.com/jamesrr39/cryptsetup-helper/blob/master/src/dmcrypthelper/cryptdevice.go

#### Multipath NFS
Special mount options in newer Linux Kernels allow for multipath mounting
     -R, --rbind
              Remount a subtree and all possible submounts somewhere else
              (so that its contents are available in both places).  See
              above, the subsection Bind mounts.
    -F, --fork
              (Used in conjunction with -a.)  Fork off a new incarnation of
              mount for each device.  This will do the mounts on different
              devices or different NFS servers in parallel.  This has the
              advantage that it is faster; also NFS timeouts go in parallel.
              A disadvantage is that the mounts are done in undefined order.
              Thus, you cannot use this option if you want to mount both
              /usr and /usr/spool.

#### Linux Device

https://github.com/golang/sys/blob/master/unix/dev_linux.go
		// Mkdev returns a Linux device number generated from the given major and minor
		// components.
		func Mkdev(major, minor uint32) uint64 {
			dev := (uint64(major) & 0x00000fff) << 8
			dev |= (uint64(major) & 0xfffff000) << 32
			dev |= (uint64(minor) & 0x000000ff) << 0
			dev |= (uint64(minor) & 0xffffff00) << 12
			return dev
		}


#### ram(4) syscall
ram - ram disk device. The ram device is a block device to access the ram disk in raw mode.
It is typically created by:
           mknod -m 660 /dev/ram b 1 1
           chown root:disk /dev/ram

#### Systemd.link
systemd.link - Network device configuration

 			 Network link configuration is performed by the net_setup_link udev
       builtin.

       The link files are read from the files located in the system network
       directory /usr/lib/systemd/network, the volatile runtime network
       directory /run/systemd/network, and the local administration network
       directory /etc/systemd/network. Link files must have the extension
       .link; other extensions are ignored. All link files are collectively
       sorted and processed in lexical order, regardless of the directories
       in which they live. However, files with identical filenames replace
       each other. Files in /etc have the highest priority, files in /run
       take precedence over files with the same name in /usr/lib. This can
       be used to override a system-supplied link file with a local file if
       needed.

[MATCH] SECTION OPTIONS

       A link file is said to match a device if each of the entries in the
       "[Match]" section matches, or if the section is empty. The following
       keys are accepted:

       MACAddress=
           The hardware address.

       OriginalName=
           A whitespace-separated list of shell-style globs matching the
           device name, as exposed by the udev property "INTERFACE". This
           cannot be used to match on names that have already been changed
           from userspace. Caution is advised when matching on
           kernel-assigned names, as they are known to be unstable between
           reboots.

       Path=
           A whitespace-separated list of shell-style globs matching the
           persistent path, as exposed by the udev property "ID_PATH".

       Driver=
           A whitespace-separated list of shell-style globs matching the
           driver currently bound to the device, as exposed by the udev
           property "DRIVER" of its parent device, or if that is not set,
           the driver as exposed by "ethtool -i" of the device itself.

       Type=
           A whitespace-separated list of shell-style globs matching the
           device type, as exposed by the udev property "DEVTYPE".

       Host=
           Matches against the hostname or machine ID of the host. See
           "ConditionHost=" in systemd.unit(5) for details.

       Virtualization=
           Checks whether the system is executed in a virtualized
           environment and optionally test whether it is a specific
           implementation. See "ConditionVirtualization=" in systemd.unit(5)
           for details.

       KernelCommandLine=
           Checks whether a specific kernel command line option is set (or
           if prefixed with the exclamation mark unset). See
           "ConditionKernelCommandLine=" in systemd.unit(5) for details.

       Architecture=
           Checks whether the system is running on a specific architecture.
           See "ConditionArchitecture=" in systemd.unit(5) for details.

[LINK] SECTION OPTIONS

       The "[Link]" section accepts the following keys:

       Description=
           A description of the device.

       Alias=
           The "ifalias" is set to this value.

       MACAddressPolicy=
           The policy by which the MAC address should be set. The available
           policies are:

           "persistent"
               If the hardware has a persistent MAC address, as most
               hardware should, and if it is used by the kernel, nothing is
               done. Otherwise, a new MAC address is generated which is
               guaranteed to be the same on every boot for the given machine
               and the given device, but which is otherwise random. This
               feature depends on ID_NET_NAME_* properties to exist for the
               link. On hardware where these properties are not set, the
               generation of a persistent MAC address will fail.

           "random"
               If the kernel is using a random MAC address, nothing is done.
               Otherwise, a new address is randomly generated each time the
               device appears, typically at boot. Either way, the random
               address will have the "unicast" and "locally administered"
               bits set.

           "none"
               Keeps the MAC address assigned by the kernel.

       MACAddress=
           The MAC address to use, if no "MACAddressPolicy=" is specified.

       NamePolicy=
           An ordered, space-separated list of policies by which the
           interface name should be set.  "NamePolicy" may be disabled by
           specifying "net.ifnames=0" on the kernel command line. Each of
           the policies may fail, and the first successful one is used. The
           name is not set directly, but is exported to udev as the property
           "ID_NET_NAME", which is, by default, used by a udev rule to set
           "NAME". If the name has already been set by userspace, no
           renaming is performed. The available policies are:

           "kernel"
               If the kernel claims that the name it has set for a device is
               predictable, then no renaming is performed.

           "database"
               The name is set based on entries in the udev's Hardware
               Database with the key "ID_NET_NAME_FROM_DATABASE".

           "onboard"
               The name is set based on information given by the firmware
               for on-board devices, as exported by the udev property
               "ID_NET_NAME_ONBOARD".

           "slot"
               The name is set based on information given by the firmware
               for hot-plug devices, as exported by the udev property
               "ID_NET_NAME_SLOT".

           "path"
               The name is set based on the device's physical location, as
               exported by the udev property "ID_NET_NAME_PATH".

           "mac"
               The name is set based on the device's persistent MAC address,
               as exported by the udev property "ID_NET_NAME_MAC".

       Name=
           The interface name to use in case all the policies specified in
           NamePolicy= fail, or in case NamePolicy= is missing or disabled.

       MTUBytes=
           The maximum transmission unit in bytes to set for the device. The
           usual suffixes K, M, G, are supported and are understood to the
           base of 1024.

       BitsPerSecond=
           The speed to set for the device, the value is rounded down to the
           nearest Mbps. The usual suffixes K, M, G, are supported and are
           understood to the base of 1000.

       Duplex=
           The duplex mode to set for the device. The accepted values are
           "half" and "full".

       AutoNegotiation=
           Enables or disables automatic negotiation of transmission
           parameters. Autonegotiation is a procedure by which two connected
           ethernet devices choose common transmission parameters, such as
           speed, duplex mode, and flow control. Takes a boolean value.
           Unset by default, which means that the kernel default will be
           used.

           Note that if autonegotiation is enabled, speed and duplex
           settings are read-only. If autonegotation is disabled, speed and
           duplex settings are writable if the driver supports multiple link
           modes.

       WakeOnLan=
           The Wake-on-LAN policy to set for the device. The supported
           values are:

           "phy"
               Wake on PHY activity.

           "magic"
               Wake on receipt of a magic packet.

           "off"
               Never wake.

       Port=
           The port option is used to select the device port. The supported
           values are:

           "tp"
               An Ethernet interface using Twisted-Pair cable as the medium.

           "aui"
               Attachment Unit Interface (AUI). Normally used with hubs.

           "bnc"
               An Ethernet interface using BNC connectors and co-axial
               cable.

           "mii"
               An Ethernet interface using a Media Independent Interface
               (MII).

           "fibre"
               An Ethernet interface using Optical Fibre as the medium.

       TCPSegmentationOffload=
           The TCP Segmentation Offload (TSO) when true enables TCP
           segmentation offload. Takes a boolean value. Defaults to "unset".

       GenericSegmentationOffload=
           The Generic Segmentation Offload (GSO) when true enables generic
           segmentation offload. Takes a boolean value. Defaults to "unset".

       UDPSegmentationOffload=
           The UDP Segmentation Offload (USO) when true enables UDP
           segmentation offload. Takes a boolean value. Defaults to "unset".

       GenericReceiveOffload=
           The Generic Receive Offload (GRO) when true enables generic
           receive offload. Takes a boolean value. Defaults to "unset".

       LargeReceiveOffload=
           The Large Receive Offload (LRO) when true enables large receive
           offload. Takes a boolean value. Defaults to "unset".

EXAMPLES

       Example 1. /usr/lib/systemd/network/99-default.link

       The link file 99-default.link that is shipped with systemd defines
       the default naming policy for links.

           [Link]
           NamePolicy=kernel database onboard slot path
           MACAddressPolicy=persistent

       Example 2. /etc/systemd/network/10-dmz.link

       This example assigns the fixed name "dmz0" to the interface with the
       MAC address 00:a0:de:63:7a:e6:

           [Match]
           MACAddress=00:a0:de:63:7a:e6

           [Link]
           Name=dmz0

       Example 3. /etc/systemd/network/10-internet.link

       This example assigns the fixed name "internet0" to the interface with
       the device path "pci-0000:00:1a.0-*":

           [Match]
           Path=pci-0000:00:1a.0-*

           [Link]
           Name=internet0

       Example 4. /etc/systemd/network/25-wireless.link

       Here's an overly complex example that shows the use of a large number
       of [Match] and [Link] settings.

           [Match]
           MACAddress=12:34:56:78:9a:bc
           Driver=brcmsmac
           Path=pci-0000:02:00.0-*
           Type=wlan
           Virtualization=no
           Host=my-laptop
           Architecture=x86-64

           [Link]
           Name=wireless0
           MTUBytes=1450
           BitsPerSecond=10M
           WakeOnLan=magic
           MACAddress=cb:a9:87:65:43:21


















#### initrd Linux kernel syscall
initrd - boot loader initialized RAM disk
CONFIGURATION
       /dev/initrd is a read-only block device assigned major number 1 and
       minor number 250.  Typically /dev/initrd is owned by root.disk with
       mode 0400 (read access by root only).  If the Linux system does not
       have /dev/initrd already created, it can be created with the
       following commands:
           mknod -m 400 /dev/initrd b 1 250
           chown root:disk /dev/initrd
       Also, support for both "RAM disk" and "Initial RAM disk" (e.g., CON‐
       FIG_BLK_DEV_RAM=y and CONFIG_BLK_DEV_INITRD=y) must be compiled
       directly into the Linux kernel to use /dev/initrd.  When using
       /dev/initrd, the RAM disk driver cannot be loaded as a module.
DESCRIPTION
       The special file /dev/initrd is a read-only block device.  This
       device is a RAM disk that is initialized (e.g., loaded) by the boot
       loader before the kernel is started.  The kernel then can use
       /dev/initrd's contents for a two-phase system boot-up.

       In the first boot-up phase, the kernel starts up and mounts an
       initial root filesystem from the contents of /dev/initrd (e.g., RAM
       disk initialized by the boot loader).  In the second phase,
       additional drivers or other modules are loaded from the initial root
       device's contents.  After loading the additional modules, a new root
       filesystem (i.e., the normal root filesystem) is mounted from a
       different device.
   Boot-up operation
       When booting up with initrd, the system boots as follows:
       1. The boot loader loads the kernel program and /dev/initrd's
          contents into memory.
       2. On kernel startup, the kernel uncompresses and copies the contents
          of the device /dev/initrd onto device /dev/ram0 and then frees the
          memory used by /dev/initrd.
       3. The kernel then read-write mounts the device /dev/ram0 as the
          initial root filesystem.
       4. If the indicated normal root filesystem is also the initial root
          filesystem (e.g., /dev/ram0) then the kernel skips to the last
          step for the usual boot sequence.
       5. If the executable file /linuxrc is present in the initial root
          filesystem, /linuxrc is executed with UID 0.  (The file /linuxrc
          must have executable permission.  The file /linuxrc can be any
          valid executable, including a shell script.)
       6. If /linuxrc is not executed or when /linuxrc terminates, the
          normal root filesystem is mounted.  (If /linuxrc exits with any
          filesystems mounted on the initial root filesystem, then the
          behavior of the kernel is UNSPECIFIED.  See the NOTES section for
          the current kernel behavior.)
       7. If the normal root filesystem has a directory /initrd, the device
          /dev/ram0 is moved from / to /initrd.  Otherwise, if the directory
          /initrd does not exist, the device /dev/ram0 is unmounted.  (When
          moved from / to /initrd, /dev/ram0 is not unmounted and therefore
          processes can remain running from /dev/ram0.  If directory /initrd
          does not exist on the normal root filesystem and any processes
          remain running from /dev/ram0 when /linuxrc exits, the behavior of
          the kernel is UNSPECIFIED.  See the NOTES section for the current
          kernel behavior.)
       8. The usual boot sequence (e.g., invocation of /sbin/init) is
          performed on the normal root filesystem.
   Options
       The following boot loader options, when used with initrd, affect the
       kernel's boot-up operation:

       initrd=filename
              Specifies the file to load as the contents of /dev/initrd.
              For LOADLIN this is a command-line option.  For LILO you have
              to use this command in the LILO configuration file
              /etc/lilo.config.  The filename specified with this option
              will typically be a gzipped filesystem image.

       noinitrd
              This boot option disables the two-phase boot-up operation.
              The kernel performs the usual boot sequence as if /dev/initrd
              was not initialized.  With this option, any contents of
              /dev/initrd loaded into memory by the boot loader contents are
              preserved.  This option permits the contents of /dev/initrd to
              be any data and need not be limited to a filesystem image.
              However, device /dev/initrd is read-only and can be read only
              one time after system startup.

       root=device-name
              Specifies the device to be used as the normal root filesystem.
              For LOADLIN this is a command-line option.  For LILO this is a
              boot time option or can be used as an option line in the LILO
              configuration file /etc/lilo.config.  The device specified by
              the this option must be a mountable device having a suitable
              root filesystem.

   Changing the normal root filesystem
       By default, the kernel's settings (e.g., set in the kernel file with
       rdev(8) or compiled into the kernel file), or the boot loader option
       setting is used for the normal root filesystems.  For an NFS-mounted
       normal root filesystem, one has to use the nfs_root_name and
       nfs_root_addrs boot options to give the NFS settings.  For more
       information on NFS-mounted root see the kernel documentation file
       Documentation/filesystems/nfs/nfsroot.txt (or
       Documentation/filesystems/nfsroot.txt before Linux 2.6.33).  For more
       information on setting the root filesystem see also the LILO and
       LOADLIN documentation.

       It is also possible for the /linuxrc executable to change the normal
       root device.  For /linuxrc to change the normal root device, /proc
       must be mounted.  After mounting /proc, /linuxrc changes the normal
       root device by writing into the proc files /proc/sys/kernel/real-
       root-dev, /proc/sys/kernel/nfs-root-name, and /proc/sys/kernel/nfs-
       root-addrs.  For a physical root device, the root device is changed
       by having /linuxrc write the new root filesystem device number into
       /proc/sys/kernel/real-root-dev.  For an NFS root filesystem, the root
       device is changed by having /linuxrc write the NFS setting into files
       /proc/sys/kernel/nfs-root-name and /proc/sys/kernel/nfs-root-addrs
       and then writing 0xff (e.g., the pseudo-NFS-device number) into file
       /proc/sys/kernel/real-root-dev.  For example, the following shell
       command line would change the normal root device to /dev/hdb1:

           echo 0x365 >/proc/sys/kernel/real-root-dev

       For an NFS example, the following shell command lines would change
       the normal root device to the NFS directory /var/nfsroot on a local
       networked NFS server with IP number 193.8.232.7 for a system with IP
       number 193.8.232.2 and named "idefix":

           echo /var/nfsroot >/proc/sys/kernel/nfs-root-name
           echo 193.8.232.2:193.8.232.7::255.255.255.0:idefix \
               >/proc/sys/kernel/nfs-root-addrs
           echo 255 >/proc/sys/kernel/real-root-dev

       Note: The use of /proc/sys/kernel/real-root-dev to change the root
       filesystem is obsolete.  See the Linux kernel source file Documenta‐
       tion/admin-guide/initrd.rst (or Documentation/initrd.txt before Linux
       4.10) as well as pivot_root(2) and pivot_root(8) for information on
       the modern method of changing the root filesystem.

   Usage
       The main motivation for implementing initrd was to allow for modular
       kernel configuration at system installation.

       A possible system installation scenario is as follows:

       1. The loader program boots from floppy or other media with a minimal
          kernel (e.g., support for /dev/ram, /dev/initrd, and the ext2
          filesystem) and loads /dev/initrd with a gzipped version of the
          initial filesystem.

       2. The executable /linuxrc determines what is needed to (1) mount the
          normal root filesystem (i.e., device type, device drivers,
          filesystem) and (2) the distribution media (e.g., CD-ROM, network,
          tape, ...).  This can be done by asking the user, by auto-probing,
          or by using a hybrid approach.

       3. The executable /linuxrc loads the necessary modules from the ini‐
          tial root filesystem.

       4. The executable /linuxrc creates and populates the root filesystem.
          (At this stage the normal root filesystem does not have to be a
          completed system yet.)

       5. The executable /linuxrc sets /proc/sys/kernel/real-root-dev,
          unmount /proc, the normal root filesystem and any other filesys‐
          tems it has mounted, and then terminates.

       6. The kernel then mounts the normal root filesystem.

       7. Now that the filesystem is accessible and intact, the boot loader
          can be installed.

       8. The boot loader is configured to load into /dev/initrd a filesys‐
          tem with the set of modules that was used to bring up the system.
          (e.g., Device /dev/ram0 can be modified, then unmounted, and
          finally, the image is written from /dev/ram0 to a file.)

       9. The system is now bootable and additional installation tasks can
          be performed.

       The key role of /dev/initrd in the above is to reuse the configura‐
       tion data during normal system operation without requiring initial
       kernel selection, a large generic kernel or, recompiling the kernel.

       A second scenario is for installations where Linux runs on systems
       with different hardware configurations in a single administrative
       network.  In such cases, it may be desirable to use only a small set
       of kernels (ideally only one) and to keep the system-specific part of
       configuration information as small as possible.  In this case, create
       a common file with all needed modules.  Then, only the /linuxrc file
       or a file executed by /linuxrc would be different.

       A third scenario is more convenient recovery disks.  Because informa‐
       tion like the location of the root filesystem partition is not needed
       at boot time, the system loaded from /dev/initrd can use a dialog
       and/or auto-detection followed by a possible sanity check.

       Last but not least, Linux distributions on CD-ROM may use initrd for
       easy installation from the CD-ROM.  The distribution can use LOADLIN
       to directly load /dev/initrd from CD-ROM without the need of any
       floppies.  The distribution could also use a LILO boot floppy and
       then bootstrap a bigger RAM disk via /dev/initrd from the CD-ROM.

FILES
       /dev/initrd
       /dev/ram0
       /linuxrc
       /initrd
NOTES
       1. With the current kernel, any filesystems that remain mounted when
          /dev/ram0 is moved from / to /initrd continue to be accessible.
          However, the /proc/mounts entries are not updated.
       2. With the current kernel, if directory /initrd does not exist, then
          /dev/ram0 will not be fully unmounted if /dev/ram0 is used by any
          process or has any filesystem mounted on it.  If /dev/ram0 is not
          fully unmounted, then /dev/ram0 will remain in memory.
       3. Users of /dev/initrd should not depend on the behavior give in the
          above notes.  The behavior may change in future versions of the
          Linux kernel.

==============
#### Rust Container
https://github.com/vishvananda/railcar

#### Microcontroller
https://github.com/oracle/smith

#### Networking
[stdlib:net](https://github.com/golang/net)
  (^)[net/ipv4/endpoint](https://github.com/golang/net/blob/master/ipv4/endpoint.go)
     Potentially way to setup a ipv4 endpoint
  (^)[net/ipv4/packet](https://github.com/golang/net/blob/master/ipv4/packet.go)
                      (https://github.com/golang/net/blob/master/ipv4/packet_go1_9.go) 
                      (https://github.com/golang/net/blob/master/ipv4/readwrite_go1_9_test.go)
     Potentially a way to reroute packets
   
**potential starting points**
* could be what we source from  : https://github.com/wrigby/flowdump/blob/master/main.go
https://github.com/ld86/syscall-udp/blob/master/main.go
https://github.com/bradleyfalzon/go-syscall-sockets/blob/master/main.go
https://github.com/prayerslayer/go-tcp/blob/master/main.go - tcp http server with workers


####
https://github.com/u-root/u-root
initramfs containing busybox like system
----
## Go Mod Handling (Kmod)
[golang-kmod](https://github.com/ElyKar/golang-kmod)


====
## Mod Handling

https://github.com/wdv4758h/rust_kernel_module/blob/master/src/stubs/linux.c
https://github.com/saschagrunert/kmod
#include <linux/module.h>
#include <linux/slab.h>

MODULE_AUTHOR("Sascha Grunert <mail@saschagrunert.de>");
MODULE_DESCRIPTION("A simple kernel module");
MODULE_LICENSE("MIT");
MODULE_VERSION("0.1.0");

// The entry and exit function
extern int init_module(void);
extern void cleanup_module(void);
----
#include <linux/module.h>

extern int rust_init(void);
extern void rust_exit(void);

static int hello_init(void) {
    return rust_init();
}

static void hello_exit(void) {
    return rust_exit();
}

module_init(hello_init);
module_exit(hello_exit);

MODULE_LICENSE("Dual BSD/GPL");
#### Container WebUI
https://github.com/coreos/etcdlabs


#### Databases
[etcd](https://github.com/coreos/etcd)


#### Kernel Modules
[conntrack](https://github.com/typetypetype/conntrack)

#### Storage
[go-tcmu](https://github.com/coreos/go-tcmu)
Go SCSI emulation via Linux TCM Userspace module


#### Configurations
There are existing configs and Multiverse OS should at least review them if not use one.

**A Container Linux Configuration, to be processed by ct, is a YAML document conforming to the following specification:**


#### Database
[riak like kv embeddable](https://github.com/tadvi/rkv)
Rkv - embeddable KV database in Go (golang).
		  Based on Riak bitcast format
		  Minimalistic design
		  Embeddable and self-contained (no C dependencies)
		  Contains both direct and goroutine friendly interfaces
		  Use rkv.NewSafe("test.kv") if you want to use with goroutines
		  Basic KV admin tool is included in /rkv subfolder, build it and install in your bin folder
		  Ability to save records with expiration
		  Use Rkv for databases under 50K records


#### WebUI TTY
https://github.com/yudai/gotty
#### Ruby
[go-mruby](https://github.com/mitchellh/go-mruby)
Ruby would be great to implement a console/terminal
that used RUBY instead of BASH 

#### Multistep / Pipeline / Workflow
[multistep](https://github.com/mitchellh/multistep)
multistep is a Go library for building up complex actions using discrete steps.

#### Filesystem
[go-fs](https://github.com/mitchellh/go-fs)
implements the ability to create, read, and write FAT filesystems using pure Go.

[go-tcmu](https://github.com/coreos/go-tcmu)
Go SCSI emulation via the Linux TCM in Userpsace module
 A daemon that handles the userspace side of the LIO TCM-User backstore. 
#### CLI UI
**IO PROGRESS**
[ioprogress](https://github.com/mitchellh/ioprogress)

**CLI Framework**
[cli](https://github.com/mitchellh/cli)
simple lightweight cli with command + subcommands
====
----
**Homedir**
https://github.com/mitchellh/go-homedir
**Packets**

**PCAP**
https://github.com/davecheney/pcap

**Standard Library: Socks5 Proxy**
https://github.com/golang/net/blob/master/proxy/socks5.go

## Potentially Interesting Image Building Tools

#### Provisioning

[routinator](https://github.com/abedra/routinator)
Good example of using templates for configs that can have variables inserted inside then moving and copying them into folders

#### Image Building 
[uspin](https://github.com/solus-project/USpin)

[go-debian](https://github.com/paultag/go-debian)
[go-bin-deb](https://github.com/mh-cbon/go-bin-deb)

[go-systemd](https://github.com/coreos/go-systemd)


[debos](https://github.com/go-debos/debos)
Debian OS Builder
  (^)[fakemachine](https://github.com/go-debos/fakemachine)
  (^)[debos-recipies](https://github.com/go-debos/debos-recipes)

#### FS
[memfs](https://github.com/zbiljic/memfs)
**nbd**
[buse-go](https://github.com/samalba/buse-go)
[go-nbd](https://github.com/frostschutz/go-nbd)
[gonbdserver](https://github.com/abligh/gonbdserver)
#### Media
**Music**
[id3](https://github.com/beevik/id3)

#### Database
[tiedot](https://github.com/HouzuoGuo/tiedot)
nosql database, embedded
#### Software defined networking
[go-openswitch/ovs](https://github.com/digitalocean/go-openvswitch/tree/master/ovs)
Package ovs is a client library for Open vSwitch which enables programmatic control of the virtual switch.
[grpc](https://github.com/grpc/grpc-go)

**distance vector routering (routing optimization)**
[go-distance-vector-routing](https://github.com/taylorflatt/go-distance-vector-routing)
DVR and a fastest path algorithm implemented using Go channels
#### (DHT) Chord lookup algorithm
[go-chord](https://github.com/taylorflatt/go-chord-implementation)

#### Packets / Networking
[gopacket](https://github.com/google/gopacket)
Very complete networking stuff
#### Tun/Tap/Device
[goTunTap](https://github.com/traetox/goTunTap)
Interesting, comes with vpn
[go-tuntap](https://github.com/0xef53/go-tuntap)
#### KVM/Qemu/Libvirt
[go-qemu](https://github.com/digitalocean/go-qemu)
[go-libvirt](https://github.com/digitalocean/go-libvirt)

[pheonix-guest-agent](https://github.com/0xef53/phoenix-guest-agent)
A guest-side agent for qemu-kvm virtual machines
[qmp-shell](https://github.com/0xef53/qmp-shell/blob/master/qmp-shell.go)


#### Iptables / Firewall
[firewall](https://github.com/Gouthamve/go-firewall)
A firewall using nfqueue (not iptables), this makes it pretty low level or atleast very high power. 


[go-iptables](https://github.com/coreos/go-iptables)
Go wrapper around iptables utility In-kernel netfilter does not have a good userspace API. The tables are manipulated via setsockopt that sets/replaces the entire table. Changes to existing table need to be resolved by userspace code which is difficult and error-prone. Netfilter developers heavily advocate using iptables utlity for programmatic manipulation.
  (^)[iptables-api](https://github.com/Oxalide/iptables-api/blob/master/main.go)


#### Globbing & Path/String Matching
[doublestar](https://github.com/bmatcuk/doublestar)


=========
#### Notes / Research
* Binding to :0 will ask the kernel for an available port within the ephemeral port range
the kernel will assign it a free port number somewhere above 1024.

	// tell pinger that it is privileged.
	// NOTE: You must run `setcap cap_net_raw=+ep pocket-loss-monitor`
  pinger.SetPrivileged(true)
=========
#### Device
https://github.com/vizee/tcpproxy splice syscall?
=========
#### Tools For Building Debian Images

xorriso - Turns folder into iso image


=========
Notes from a manual Multiverse OS setup on Intel I9.
The process of installing Multiverse OS is very detailed and ideally is done using fresh hardware that has never been used. One starts with BIOS settings and slowly builds upon secure settings.
## BIOS Settings
Execute Disable Bit = ENABLED
"It is highly recommended that you enable this BIOS feature for increased protection against buffer overflow attacks."
## Networking Devices
The Asus motherboard I'm using for the Intel I9 processor is an Atheros NIC, these drivers are not included with the Debian 9 net install. Complexities like this have been anticipated so a Realtek Gigabit network card was purchased with the rest of the components for eventualalties such as this but also to provide a wider variety of network cards and drivers for the Multiverse VM networking.

Realtek drivers are included with the Debian 9 network install and the installation can continue by using this card for *Host machine* initialization.

**NOTE** [FUTURE DEVELOPMENT]
*For maximum compatibility this eventuality should be considered and the Multiverse OS install medium should incldue the Atheros driver if it is open source.*
## Gnome
https://github.com/reujab/gse
## Debian Installation
Multiverse OS follows the same conventions as other psuedo-anonymous security oriented linux operating systems:
	[Accounts]
	Hostname: "host"
	Domain name: ""
	Full name: "user"
	Username: "user"
**NOTE** [FUTURE DEVELOPMENT]
*In the future this will all be automated and the root and user passwords will be automatically generated deterministically from a master key based on time and other factors. This will allow for OTP, automatic renewal of passwords, offline recovery, automatic input for secure booting of the host machine and eventual linux kernel modifications that will massively increase the security of the system (require user input to be signed, and encrypt user output).*
[Disks]
[x] Guided - use entire disk and setup encrypted LVM
**NOTE** Debian 9 install had issues with the install process when using M.2 drives because
**NOTE** [FUTURE DEVELOPMENT]
	[Packages]
	No desktop envirnonment (No Gnome, Xfce, KDE, Cinnamon, MATE, LXDWE)
	No web server
	No print server
	No SSH server
*ONLY standard system utilities*

[Post-Install]
**NOTE** Key Generation should happen here, PGP, SSH, etc (Basically setting up Shifter/Scramble Suit)
**NOTE** WGET urls for (1) Debian, (2) Alpine Linux, (3) Kali
+ Untar/Unzip ISOs, check if each one required exists using checksums 
  [+] Check against checksum, signature (preferred, should do it with multiverse OS dev key if not provided by official distribution).
  [+] Add a new "Storage" in libvirt for the ISOs 

sudo apt-get remove firefox-esr nano
sudo apt-get install libvirt-daemon-system vim
## DEV
## sudo apt-get install virt-manager gnome
## 
## I used the UI when developing, but ideally this should not be 

**NOTE** After the packages are settled, it is time to server connection to the internet. Do this by forcing the kernel not to load ANY networking modules. This will free them up so they can be used with the VMs too (PCI Passthrough).

Modify /etc/default/grub, add `intel_iommu=on` to the options line.

(Here we need to disable any graphics cards too eventually, but if you do this too early you will not have a monitor to setup the controller VM.)

Additionally, locate each network device and disable it: `pci-stub.ids=10de:13c2,10de:0fbb`

Start this process by using the command

````
	lspci -nn
````

The `-nn` provides a truncated output with the pci address information needed to disable the device by preventing the kernel from loading it by blacklisting it.

sudo touch /etc/modprobe.d/blacklist.conf

sudo echo "#Blacklist Firewire\nblacklist firewire_core\n\n#Blacklist PC Speaker (Annoying)\nblacklist pcspkr" > /etc/modprobe.d/blacklist.conf

# echo "0000:00:1a.0" > /sys/bus/pci/drivers/ehci_hcd/unbind

## Reverse SSHD
Here is how you do the reverse SSH if you decide that route

ssh -R NEW_SERVER:80:127.0.0.1:80 unprivuser@NEW_SERVER -v

Basically then you can ssh into like port 2222 on the chunkhost server and it goes into the PI without needing any ports open on the modem.



