# Raw Socket / Interface / Device
https://github.com/u-root/u-root/tree/6fd12df7e9d96fce0acb7ab723b7f129e5b07107/pkg/pci

https://github.com/milosgajdos83/tenus

## WebKit GO Bindings 
https://github.com/abcum/webkit
## DB
https://github.com/abcum/tlist - a linked time series list 
## Ruby Config
https://github.com/k0kubun/itamae-go
## systemd
https://github.com/coreos/go-systemd
## AutoTLS for Gin
https://github.com/gin-gonic/autotls
## Reed Solomon
https://github.com/klauspost/reedsolomon
## Chord
https://github.com/armon/go-chord
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


https://github.com/mhcvs2/go/blob/babfc17df402f3a1931fc3caafe32564011703e8/src/golang.org/x/sys/unix/zsyscall_linux_amd64.go	
https://github.com/mhcvs2/go/blob/babfc17df402f3a1931fc3caafe32564011703e8/src/golang.org/x/sys/unix/syscall_linux_amd64.go
https://github.com/mhcvs2/go/blob/babfc17df402f3a1931fc3caafe32564011703e8/src/golang.org/x/sys/unix/syscall_linux.go


https://github.com/mhcvs2/go/blob/babfc17df402f3a1931fc3caafe32564011703e8/src/golang.org/x/sys/unix/dev_linux.go

fds, err := unix.Socketpair(unix.AF_LOCAL, tt.socketType, 0)
		if err != nil {
			t.Fatalf("Socketpair: %v", err)
		}
		defer unix.Close(fds[0])
defer unix.Close(fds[1])
#### Notes / Research
* Binding to :0 will ask the kernel for an available port within the ephemeral port range
the kernel will assign it a free port number somewhere above 1024.

	// tell pinger that it is privileged.
	// NOTE: You must run `setcap cap_net_raw=+ep pocket-loss-monitor`
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


func mkfifo(path string, mode uint32) error {
	return syscall.Mkfifo(path, mode)
}

func mksocket(path string) error {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	tmp := filepath.Join(dir, "."+base)
	l, err := net.ListenUnix("unix", &net.UnixAddr{Name: tmp, Net: "unix"})
	if err != nil {
		return err
	}

	err = os.Rename(tmp, path)
	if err != nil {
		l.Close()
		os.Remove(tmp) // Ignore error
		return err
	}

	l.Close()

	return nil
}

func maxFD() (uint64, error) {
	var rlim syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim); err != nil {
		return 0, fmt.Errorf("ulimit error: %v", err)
	}
	return rlim.Cur, nil
}

https://github.com/znly/cloud-debug-go/blob/893a0ba87789b28399e1d37d9190553ee0cae031/vendor/camlistore.org/pkg/osutil/syscall_posix.go
#### Kernel Modules
https://github.com/u-root/u-root/blob/6fd12df7e9d96fce0acb7ab723b7f129e5b07107/pkg/kmodule/kmodule_linux.go


ret, _, err := syscall.Syscall(syscall.SYS_INIT_MODULE, uintptr(unsafe.Pointer(&file[0])), uintptr(len(file)), uintptr(unsafe.Pointer(&[]byte(options)[0])))
	if ret != 0 {
		log.Fatalf("insmod: error inserting '%s': %v %v\n", filename, ret, err)
}


package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

var (
	device       string = "eth0"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}


veth netlink example
https://github.com/Mirantis/virtlet/blob/d7bd74f66d574a779d77c10c4b496b8bd8242625/pkg/nettools/nettools.go
------



https://github.com/containers/virtcontainers/blob/f8d2c11dcfd5c04e529e5900a195ee055734d5df/network.go

pening a raw socket
To open a socket, you have to know three things  the socket family, socket type and protocol. For a raw socket, the socket family is AF_PACKET, the socket type is SOCK_RAW and for the protocol, see the if_ether.h header file. To receive all packets, the macro is ETH_P_ALL and to receive IP packets, the macro is ETH_P_IP for the protocol field.
int sock_r;
sock_r=socket(AF_PACKET,SOCK_RAW,htons(ETH_P_ALL));
if(sock_r<0)
{
printf(error in socket\n);
return -1;
}
-----------

https://github.com/hammer-os/os/blob/master/initrd/init/linux/linux.go


func Mkchar(path string, mode, major, minor uint32) error {
	_, err := os.Lstat(path) // character device already exists
	if err == nil {
		return nil
	}

	dev := int(unix.Mkdev(major, minor))
	if err = unix.Mknod(path, mode, dev); err != nil {
		return &os.PathError{"mknod", path, err}
	}
	return nil
}









----


https://github.com/vishvananda/netlink

https://github.com/mdlayher/vsock
https://github.com/golang/go/blob/master/src/net/rawconn.go

https://github.com/golang/go/blob/master/src/net/protoconn_test.go
c1, err := DialUnix("unixgram", a1, nil)
c, err := ListenIP("ip4:icmp", la)

https://github.com/golang/go/blob/master/src/net/iprawsock_posix.go

INTEFACE
// Addrs returns a list of unicast interface addresses for a specific
// interface.
func (ifi *Interface) Addrs() ([]Addr, error) {
	if ifi == nil {
		return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterface}
	}
	ifat, err := interfaceAddrTable(ifi)
	if err != nil {
		err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
	}
	return ifat, err
}



https://github.com/golang/go/blob/master/src/net/iprawsock.go
// SyscallConn returns a raw network connection.
// This implements the syscall.Conn interface.
func (c *IPConn) SyscallConn() (syscall.RawConn, error) {
	if !c.ok() {
		return nil, syscall.EINVAL
	}
	return newRawConn(c.fd)
}


https://github.com/manifoldco/promptui
 w
https://github.com/mdlayher/raw
https://github.com/prep/socketpair
https://github.com/Azure-Samples/network-go-manage-network-interface/blob/master/example.go

https://github.com/eleme/netlink - great readme

https://github.com/mdlayher/netlink
https://github.com/subgraph/go-nfnetlink
https://github.com/mickep76/netlink/blob/master/netlink.go
https://github.com/tehnerd/gnl2go/blob/master/ipvs.go
https://github.com/rirenner/netlink
https://github.com/milosgajdos83/tenus
https://github.com/aporeto-inc/netlink-go
https://github.com/yetu/upstart-netlink

https://github.com/lambdasoup/go-netlink/blob/master/netlink/netlink.go

https://github.com/jjh2kiss/netlinkconnector

https://github.com/apuigsech/netlink/blob/master/netlink.go 
https://github.com/remyoudompheng/go-netlink/blob/master/socket.go

https://github.com/mqliang/libipvs


https://github.com/google/gopacket
https://github.com/golang/go/blob/master/src/net/file_unix.go
https://github.com/golang/go/blob/master/src/net/interface.go

https://github.com/utamaro/shards


[netstack]
https://github.com/dutchcoders/netstack
[raw packet]
https://github.com/mdlayher/raw


https://github.com/arcpop/rawsocket

https://github.com/wheelcomplex/gorawpacket
[pcap]
https://github.com/Micheloss/Go_Raw/blob/master/raw.go
https://github.com/kyleconroy/coiltap/blob/master/pcap.go
[tun/tap examples]
https://github.com/jamescun/tuntap




https://github.com/songgao/water
https://github.com/pkg/taptun


https://github.com/virtmonitor/virNetTap

https://github.com/vsergeev/tinytaptunnel

https://github.com/jaracil/tuntap

https://github.com/hwhw/gotun/blob/master/gotun.go
https://github.com/enricomariafusi/v-switch
[sockets]
https://github.com/hashicorp/go-sockaddr
https://github.com/ryabuhin/golang_linux_rawsockets/blob/master/icmpservice/icmpservice.go
----
Basic dev folder interaction
https://github.com/opencontainers/runc/blob/master/libcontainer/devices/devices.go

https://github.com/FarmRadioHangar/fdevices - stream realtime events about devices

// IsDevice validate if the path is a type of device file
func (s SystemPath) IsDevice() bool {
	return s.IsStat(os.ModeDevice)
}

// IsNamedPipe validate if the path is a pipe file
func (s SystemPath) IsNamedPipe() bool {
	return s.IsStat(os.ModeNamedPipe)
}

// IsSocket validate if the path is a socket file
func (s SystemPath) IsSocket() bool {
	return s.IsStat(os.ModeSocket)
}

// IsCharDevice validate if the path is a char file
func (s SystemPath) IsCharDevice() bool {
	return s.IsStat(os.ModeCharDevice)
}







https://github.com/frozzare/go-fs
