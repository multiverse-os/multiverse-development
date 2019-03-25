##
##  Multiverse OS: Sample Code From Other Projects
===========================================================




// Manhole connects os.Stdin, os.Stdout, and os.Stderr to an interactive shell
// session on the Machine m. Manhole blocks until the shell session has ended.
// If os.Stdin does not refer to a TTY, Manhole returns immediately with a nil
// error. Copied from github.com/coreos/mantle/platform/util.go
func (c *SSHClient) Manhole(host string) error {
	fd := int(os.Stdin.Fd())
	if !terminal.IsTerminal(fd) {
		return nil
	}

	tstate, _ := terminal.MakeRaw(fd)
	defer terminal.Restore(fd, tstate)

	client, err := ssh.Dial("tcp", host+":22", c.ClientConfig)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("SSH session failed: %v", err)
	}

	defer session.Close()

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	modes := ssh.TerminalModes{
		ssh.TTY_OP_ISPEED: 115200,
		ssh.TTY_OP_OSPEED: 115200,
	}

	cols, lines, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	if err = session.RequestPty(os.Getenv("TERM"), lines, cols, modes); err != nil {
		return fmt.Errorf("failed to request pseudo terminal: %s", err)
	}

	if err := session.Shell(); err != nil {
		return fmt.Errorf("failed to start shell: %s", err)
	}

	if err := session.Wait(); err != nil {
		return fmt.Errorf("failed to wait for session: %s", err)
	}

	return nil
}

----

// DefaultProfile defines the whitelist for the default seccomp profile.
func DefaultProfile(sp *specs.Spec) *specs.LinuxSeccomp {
	syscalls := []specs.LinuxSyscall{
		{
			Names: []string{
				"accept",
				"accept4",
				"access",
				"alarm",
				"alarm",
				"bind",
				"brk",
				"capget",
				"capset",
				"chdir",
				"chmod",
				"chown",
				"chown32",
				"clock_getres",
				"clock_gettime",
				"clock_nanosleep",
				"close",
				"connect",
				"copy_file_range",
				"creat",
				"dup",
				"dup2",
				"dup3",
				"epoll_create",
				"epoll_create1",
				"epoll_ctl",
				"epoll_ctl_old",
				"epoll_pwait",
				"epoll_wait",
				"epoll_wait_old",
				"eventfd",
				"eventfd2",
				"execve",
				"execveat",
				"exit",
				"exit_group",
				"faccessat",
				"fadvise64",
				"fadvise64_64",
				"fallocate",
				"fanotify_mark",
				"fchdir",
				"fchmod",
				"fchmodat",
				"fchown",
				"fchown32",
				"fchownat",
				"fcntl",
				"fcntl64",
				"fdatasync",
				"fgetxattr",
				"flistxattr",
				"flock",
				"fork",
				"fremovexattr",
				"fsetxattr",
				"fstat",
				"fstat64",
				"fstatat64",
				"fstatfs",
				"fstatfs64",
				"fsync",
				"ftruncate",
				"ftruncate64",
				"futex",
				"futimesat",
				"getcpu",
				"getcwd",
				"getdents",
				"getdents64",
				"getegid",
				"getegid32",
				"geteuid",
				"geteuid32",
				"getgid",
				"getgid32",
				"getgroups",
				"getgroups32",
				"getitimer",
				"getpeername",
				"getpgid",
				"getpgrp",
				"getpid",
				"getppid",
				"getpriority",
				"getrandom",
				"getresgid",
				"getresgid32",
				"getresuid",
				"getresuid32",
				"getrlimit",
				"get_robust_list",
				"getrusage",
				"getsid",
				"getsockname",
				"getsockopt",
				"get_thread_area",
				"gettid",
				"gettimeofday",
				"getuid",
				"getuid32",
				"getxattr",
				"inotify_add_watch",
				"inotify_init",
				"inotify_init1",
				"inotify_rm_watch",
				"io_cancel",
				"ioctl",
				"io_destroy",
				"io_getevents",
				"ioprio_get",
				"ioprio_set",
				"io_setup",
				"io_submit",
				"ipc",
				"kill",
				"lchown",
				"lchown32",
				"lgetxattr",
				"link",
				"linkat",
				"listen",
				"listxattr",
				"llistxattr",
				"_llseek",
				"lremovexattr",
				"lseek",
				"lsetxattr",
				"lstat",
				"lstat64",
				"madvise",
				"memfd_create",
				"mincore",
				"mkdir",
				"mkdirat",
				"mknod",
				"mknodat",
				"mlock",
				"mlock2",
				"mlockall",
				"mmap",
				"mmap2",
				"mprotect",
				"mq_getsetattr",
				"mq_notify",
				"mq_open",
				"mq_timedreceive",
				"mq_timedsend",
				"mq_unlink",
				"mremap",
				"msgctl",
				"msgget",
				"msgrcv",
				"msgsnd",
				"msync",
				"munlock",
				"munlockall",
				"munmap",
				"nanosleep",
				"newfstatat",
				"_newselect",
				"open",
				"openat",
				"pause",
				"pipe",
				"pipe2",
				"poll",
				"ppoll",
				"prctl",
				"pread64",
				"preadv",
				"prlimit64",
				"pselect6",
				"pwrite64",
				"pwritev",
				"read",
				"readahead",
				"readlink",
				"readlinkat",
				"readv",
				"recv",
				"recvfrom",
				"recvmmsg",
				"recvmsg",
				"remap_file_pages",
				"removexattr",
				"rename",
				"renameat",
				"renameat2",
				"restart_syscall",
				"rmdir",
				"rt_sigaction",
				"rt_sigpending",
				"rt_sigprocmask",
				"rt_sigqueueinfo",
				"rt_sigreturn",
				"rt_sigsuspend",
				"rt_sigtimedwait",
				"rt_tgsigqueueinfo",
				"sched_getaffinity",
				"sched_getattr",
				"sched_getparam",
				"sched_get_priority_max",
				"sched_get_priority_min",
				"sched_getscheduler",
				"sched_rr_get_interval",
				"sched_setaffinity",
				"sched_setattr",
				"sched_setparam",
				"sched_setscheduler",
				"sched_yield",
				"seccomp",
				"select",
				"semctl",
				"semget",
				"semop",
				"semtimedop",
				"send",
				"sendfile",
				"sendfile64",
				"sendmmsg",
				"sendmsg",
				"sendto",
				"setfsgid",
				"setfsgid32",
				"setfsuid",
				"setfsuid32",
				"setgid",
				"setgid32",
				"setgroups",
				"setgroups32",
				"setitimer",
				"setpgid",
				"setpriority",
				"setregid",
				"setregid32",
				"setresgid",
				"setresgid32",
				"setresuid",
				"setresuid32",
				"setreuid",
				"setreuid32",
				"setrlimit",
				"set_robust_list",
				"setsid",
				"setsockopt",
				"set_thread_area",
				"set_tid_address",
				"setuid",
				"setuid32",
				"setxattr",
				"shmat",
				"shmctl",
				"shmdt",
				"shmget",
				"shutdown",
				"sigaltstack",
				"signalfd",
				"signalfd4",
				"sigreturn",
				"socket",
				"socketcall",
				"socketpair",
				"splice",
				"stat",
				"stat64",
				"statfs",
				"statfs64",
				"symlink",
				"symlinkat",
				"sync",
				"sync_file_range",
				"syncfs",
				"sysinfo",
				"syslog",
				"tee",
				"tgkill",
				"time",
				"timer_create",
				"timer_delete",
				"timerfd_create",
				"timerfd_gettime",
				"timerfd_settime",
				"timer_getoverrun",
				"timer_gettime",
				"timer_settime",
				"times",
				"tkill",
				"truncate",
				"truncate64",
				"ugetrlimit",
				"umask",
				"uname",
				"unlink",
				"unlinkat",
				"utime",
				"utimensat",
				"utimes",
				"vfork",
				"vmsplice",
				"wait4",
				"waitid",
				"waitpid",
				"write",
"writev",


type Link uint32

const (
	NULL                       Link = 0
	ETHERNET                   Link = 1
	AX25                       Link = 3
	IEEE802_5                  Link = 6
	ARCNET_BSD                 Link = 7
	SLIP                       Link = 8
	PPP                        Link = 9
	FDDI                       Link = 10
	PPP_HDLC                   Link = 50
	PPP_ETHER                  Link = 51
	ATM_RFC1483                Link = 100
	RAW                        Link = 101
	C_HDLC                     Link = 104
	IEEE802_11                 Link = 105
	FRELAY                     Link = 107
	LOOP                       Link = 108
	LINUX_SLL                  Link = 113
	LTALK                      Link = 114
	PFLOG                      Link = 117
	IEEE802_11_PRISM           Link = 119
	IP_OVER_FC                 Link = 122
	SUNATM                     Link = 123
	IEEE802_11_RADIOTAP        Link = 127
	ARCNET_LINUX               Link = 129
	APPLE_IP_OVER_IEEE1394     Link = 138
	MTP2_WITH_PHDR             Link = 139
	MTP2                       Link = 140
	MTP3                       Link = 141
	SCCP                       Link = 142
	DOCSIS                     Link = 143
	LINUX_IRDA                 Link = 144
	IEEE802_11_AVS             Link = 163
	BACNET_MS_TP               Link = 165
	PPP_PPPD                   Link = 166
	GPRS_LLC                   Link = 169
	LINUX_LAPD                 Link = 177
	BLUETOOTH_HCI_H4           Link = 187
	USB_LINUX                  Link = 189
	PPI                        Link = 192
	IEEE802_15_4               Link = 195
	SITA                       Link = 196
	ERF                        Link = 197
	BLUETOOTH_HCI_H4_WITH_PHDR Link = 201
	AX25_KISS                  Link = 202
	LAPD                       Link = 203
	PPP_WITH_DIR               Link = 204
	C_HDLC_WITH_DIR            Link = 205
	FRELAY_WITH_DIR            Link = 206
	IPMB_LINUX                 Link = 209
	IEEE802_15_4_NONASK_PHY    Link = 215
	USB_LINUX_MMAPPED          Link = 220
	FC_2                       Link = 224
	FC_2_WITH_FRAME_DELIMS     Link = 225
	IPNET                      Link = 226
	CAN_SOCKETCAN              Link = 227
	IPV4                       Link = 228
	IPV6                       Link = 229
	IEEE802_15_4_NOFCS         Link = 230
	DBUS                       Link = 231
	DVB_CI                     Link = 235
	MUX27010                   Link = 236
	STANAG_5066_D_PDU          Link = 237
	NFLOG                      Link = 239
	NETANALYZER                Link = 240
	NETANALYZER_TRANSPARENT    Link = 241
	IPOIB                      Link = 242
	MPEG_2_TS                  Link = 243
	NG40                       Link = 244
	NFC_LLCP                   Link = 245
	INFINIBAND                 Link = 247
	SCTP                       Link = 248
	USBPCAP                    Link = 249
	RTAC_SERIAL                Link = 250
	BLUETOOTH_LE_LL            Link = 251
)

// Define the EtherType type, for ethernet frames. Additionally define some known ethertypes.
type EtherType uint16

const (
	ETHERTYPE_IPV4    EtherType = 0x0800
	ARP               EtherType = 0x0806
	WAKE_ON_LAN       EtherType = 0x0842
	TRILL             EtherType = 0x22F3
	DECNET_PHASE_4    EtherType = 0x6003
	REVERSE_ARP       EtherType = 0x8035
	APPLETALK         EtherType = 0x809B
	APPLETALK_ARP     EtherType = 0x80F3
	IPX1              EtherType = 0x8137
	IPX2              EtherType = 0x8138
	QNET              EtherType = 0x8204
	ETHERTYPE_IPV6    EtherType = 0x86DD
	FLOWCONTROL       EtherType = 0x8808
	SLOW              EtherType = 0x8809
	COBRANET          EtherType = 0x8819
	MPLS_UNICAST      EtherType = 0x8847
	MPLS_MULTICAST    EtherType = 0x8848
	PPPOE_DISCOVERY   EtherType = 0x8863
	PPPOE_SESSION     EtherType = 0x8864
	JUMBO_FRAMES      EtherType = 0x8870
	HOMEPLUG          EtherType = 0x887B
	EAP_OVER_LAN      EtherType = 0x888E
	PROFINET          EtherType = 0x8892
	HYPERSCSI         EtherType = 0x889A
	ATA_OVER_ETHERNET EtherType = 0x88A2
	ETHERCAT          EtherType = 0x88A4
	POWERLINK         EtherType = 0x88AB
	LLDP              EtherType = 0x88CC
	SERCOS3           EtherType = 0x88CD
	MRP               EtherType = 0x88E3
	MAC_SECURITY      EtherType = 0x88E5
	IEEE1588          EtherType = 0x88F7
	FCOE              EtherType = 0x8906
	FCOE_INIT         EtherType = 0x8914
	ROCE              EtherType = 0x8915
	HSR               EtherType = 0x892F
)

// IPProtocol defines the potential protocols enclosed by an IP packet. Some representative
// symbolic constants are defined in this file, but many more exist.
type IPProtocol uint8

const (
	IPP_ICMP      IPProtocol = 0x01
	IPP_TCP       IPProtocol = 0x06
	IPP_UDP       IPProtocol = 0x11
	IPP_TLSP      IPProtocol = 0x38
	IPP_IPV6_ICMP IPProtocol = 0x3A
	IPP_SCTP      IPProtocol = 0x84
)

// PcapFile represents the parsed form of a single .pcap file. The structure
// contains some details about the file itself, but is mostly a container for
// the parsed Packets.
type PcapFile struct {
	MajorVersion uint16
	MinorVersion uint16
	TZCorrection int32 // In seconds east of UTC
	SigFigs      uint32
	MaxLen       uint32
	LinkType     Link
	Packets      []Packet
}

// Packet is a representation of a single network packet. The structure
// contains the timestamp on the packet, some information about packet size,
// and the recorded bytes from the packet.
type Packet struct {
	Timestamp   time.Duration
	IncludedLen uint32
	ActualLen   uint32
	Data        LinkLayer
}


const (
	TCP4 uint8 = iota
	UDP4
	TCP6
	UDP6
	ICMP4
	ICMP6

	TCP4Data  = "/proc/net/tcp"
	UDP4Data  = "/proc/net/udp"
	TCP6Data  = "/proc/net/tcp6"
	UDP6Data  = "/proc/net/udp6"
	ICMP4Data = "/proc/net/icmp"
	ICMP6Data = "/proc/net/icmp6"

	TCP_ESTABLISHED = iota + 1
	TCP_SYN_SENT
	TCP_SYN_RECV
	TCP_FIN_WAIT1
	TCP_FIN_WAIT2
	TCP_TIME_WAIT
	TCP_CLOSE
	TCP_CLOSE_WAIT
	TCP_LAST_ACK
	TCP_LISTEN
	TCP_CLOSING
	TCP_NEW_SYN_RECV
)

var (
	// connectionSocketsLock sync.Mutex
	// connectionTCP4 = make(map[string][]int)
	// connectionUDP4 = make(map[string][]int)
	// connectionTCP6 = make(map[string][]int)
	// connectionUDP6 = make(map[string][]int)

	listeningSocketsLock sync.Mutex
	addressListeningTCP4 = make(map[string][]int)
	addressListeningUDP4 = make(map[string][]int)
	addressListeningTCP6 = make(map[string][]int)
	addressListeningUDP6 = make(map[string][]int)
	globalListeningTCP4  = make(map[uint16][]int)
	globalListeningUDP4  = make(map[uint16][]int)
	globalListeningTCP6  = make(map[uint16][]int)
	globalListeningUDP6  = make(map[uint16][]int)
)

https://github.com/Safing/safing-core/tree/master/network/packet
	IGMP   = IPProtocol(syscall.IPPROTO_IGMP)
	RAW    = IPProtocol(syscall.IPPROTO_RAW)
	TCP    = IPProtocol(syscall.IPPROTO_TCP)
	UDP    = IPProtocol(syscall.IPPROTO_UDP)
	ICMP   = IPProtocol(syscall.IPPROTO_ICMP)
ICMPv6 = IPProtocol(syscall.IPPROTO_ICMPV6)

func createUDPSocket(ipv6Enabled bool, sa syscall.Sockaddr) (int, error) {
	// Grab the correct socket family.
	family := syscall.AF_INET
	if ipv6Enabled {
		family = syscall.AF_INET6
	}

	// Create the socket.
	fd, err := syscall.Socket(family, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return -1, errors.New("error setting the UDP socket parameters: " + err.Error())
	}

	// Set the various socket options required.
	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		return -1, errors.New("error setting the UDP socket parameters: " + err.Error())
	}

	// Bind the newly created and configured socket.
	err = syscall.Bind(fd, sa)
	if err != nil {
		return -1, errors.New("error binding the UDP socket to the configured listen address: " + err.Error())
	}

	return fd, nil
}
----
// GetHomeDir returns the home directory
// TODO: Having this here just strikes me as dangerous, but some of the drivers
// depend on it ;_;
func GetHomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}

func GetUsername() string {
	u := "unknown"
	osUser := ""

	switch runtime.GOOS {
	case "darwin", "linux":
		osUser = os.Getenv("USER")
	case "windows":
		osUser = os.Getenv("USERNAME")
	}

	if osUser != "" {
		u = osUser
	}

	return u
}


func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, fi.Mode())
}


=====


// Ulimit is a human friendly version of Rlimit.
type Ulimit struct {
	Name string
	Hard int64
	Soft int64
}

// Rlimit specifies the resource limits, such as max open files.
type Rlimit struct {
	Type int    `json:"type,omitempty"`
	Hard uint64 `json:"hard,omitempty"`
	Soft uint64 `json:"soft,omitempty"`
}

const (
	// magic numbers for making the syscall
	// some of these are defined in the syscall package, but not all.
	// Also since Windows client doesn't get access to the syscall package, need to
	//	define these here
	rlimitAs         = 9
	rlimitCore       = 4
	rlimitCPU        = 0
	rlimitData       = 2
	rlimitFsize      = 1
	rlimitLocks      = 10
	rlimitMemlock    = 8
	rlimitMsgqueue   = 12
	rlimitNice       = 13
	rlimitNofile     = 7
	rlimitNproc      = 6
	rlimitRss        = 5
	rlimitRtprio     = 14
	rlimitRttime     = 15
	rlimitSigpending = 11
	rlimitStack      = 3
)

var ulimitNameMapping = map[string]int{
	//"as":         rlimitAs, // Disabled since this doesn't seem usable with the way Docker inits a container.
	"core":       rlimitCore,
	"cpu":        rlimitCPU,
	"data":       rlimitData,
	"fsize":      rlimitFsize,
	"locks":      rlimitLocks,
	"memlock":    rlimitMemlock,
	"msgqueue":   rlimitMsgqueue,
	"nice":       rlimitNice,
	"nofile":     rlimitNofile,
	"nproc":      rlimitNproc,
	"rss":        rlimitRss,
	"rtprio":     rlimitRtprio,
	"rttime":     rlimitRttime,
	"sigpending": rlimitSigpending,
	"stack":      rlimitStack,
}

// ParseUlimit parses and returns a Ulimit from the specified string.
func ParseUlimit(val string) (*Ulimit, error) {
	parts := strings.SplitN(val, "=", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid ulimit argument: %s", val)
	}

	if _, exists := ulimitNameMapping[parts[0]]; !exists {
		return nil, fmt.Errorf("invalid ulimit type: %s", parts[0])
	}

	var (
		soft int64
		hard = &soft // default to soft in case no hard was set
		temp int64
		err  error
	)
	switch limitVals := strings.Split(parts[1], ":"); len(limitVals) {
	case 2:
		temp, err = strconv.ParseInt(limitVals[1], 10, 64)
		if err != nil {
			return nil, err
		}
		hard = &temp
		fallthrough
	case 1:
		soft, err = strconv.ParseInt(limitVals[0], 10, 64)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("too many limit value arguments - %s, can only have up to two, `soft[:hard]`", parts[1])
	}

	if soft > *hard {
		return nil, fmt.Errorf("ulimit soft limit must be less than or equal to hard limit: %d > %d", soft, *hard)
	}

	return &Ulimit{Name: parts[0], Soft: soft, Hard: *hard}, nil
}

// GetRlimit returns the RLimit corresponding to Ulimit.
func (u *Ulimit) GetRlimit() (*Rlimit, error) {
	t, exists := ulimitNameMapping[u.Name]
	if !exists {
		return nil, fmt.Errorf("invalid ulimit name %s", u.Name)
	}

	return &Rlimit{Type: t, Soft: uint64(u.Soft), Hard: uint64(u.Hard)}, nil
}

func (u *Ulimit) String() string {
	return fmt.Sprintf("%s=%d:%d", u.Name, u.Soft, u.Hard)
}
