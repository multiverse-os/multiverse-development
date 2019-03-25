
# Interesting software for related/auxillary ideas


https://github.com/1lann/TOHT-Proxy - Disguise TCP traffic as HTTP


# TCP/IP Network Stack

[Kernel Module]
**conntrack**
Every call to c.Connections() will return all connections active since the last call to c.Connections(). The connections can either still be established, or have been terminated since the last call. Connections which are established and teared down in between calls to c.Connections() will also be reported.
GIT https://github.com/typetypetype/conntrack

## Messaging
  **messanger**
  Nice modular simple message example. COuld be a nice starting point.
  GIT https://github.com/go-distributed/messenger

## Configuration
  **Config module**
  Simple Viper type module for key/value configuration using a basic JSON file. One level but nice nonetheless. Done in a nice module format
  GIT https://github.com/CarterTsai/my_go_module

## SNMP
  **smtp**
  GIT https://github.com/likexian/mailer-go

## ID
[ID]
  **basex** - Produce youtube like IDs
  GIT https://github.com/dineshappavoo/basex


## Physical/Virtual Device
[MAC]
  **Macouflage** - Mac address anonymous tool linux based os. Features the ability to generate macs by device, randomized, etc. Large gradient of control for better ability to hide. There is a tool and a library.
  GIT https://github.com/subgraph/macouflage

[Virtual Device]
  **Vsock**
  Package vsock provides access to Linux VM sockets (AF_VSOCK) for communication between a   hypervisor and its virtual machines.
  _NOTE_ __To use, use the following command when starting the VM__
  __-device vhost-vsock-pci,id=vhost-vsock-pci0,guest-cid=3__
  GIT https://github.com/mdlayher/vsock

  **vnet**
  Virtual connection persistent on native tcp conncetion
  GIT https://github.com/ymmuse/vnet


## Simple Service Discovery (SSDP)
https://github.com/kazyx/ssdprecv

## Network Stack
  * Very nice look tuntap, love func names https://github.com/traetox/goTunTap


## Netlink
The netlink package provides a communication protocol which allows process exchanging information no matter ther are in user or kernel space. *Most kernal's process need communicate with user's process in Linux, but traditional Unix's IPC (pipe, message queue, shared memory and singal) can not offer a strong support for the communication between user's process and kernel. Linux provides a lot other methods which allow user's process can communicate with kernel, but they are very hard to use. To make these method easier to user for user, especially for Operational Engineer is the reason why we develop netlink.* https://github.com/eleme/netlink
  * most info https://github.com/eleme/netlink
  * https://github.com/aporeto-inc/netlink-go - well organized
  * Bestest https://github.com/milosgajdos83/tenus
  * Best https://github.com/mdlayher/netlink
  * Most developed https://github.com/vishvananda/netlink
    ^ https://github.com/osrg/goplane/ good use of vishvanda netlink shit
  * simple https://github.com/eleme/netlink
  * https://github.com/mpleso/netlink - nice new netlink lib
  * https://github.com/jsimonetti/rtnetlink 
  * https://github.com/subgraph/go-nfnetlink - Netfilter is composed of several subsystems in the Linux kernel, some of which provide access from userland over a netlink socket interface. The protocol API for accessing these subsystems share a common set of protocol conventions called nfnetlink (netfilter netlink).
  * https://github.com/mpleso/vnet - lowest yet?
  * https://github.com/mpleso/netlink
  * https://github.com/mdlayher/netlink - nice looking socket

  **tenus**
  netwlink from runc
  GIT https://github.com/milosgajdos83/tenus


## PCAP Libraries
# Tap
  * https://github.com/jamescun/tuntap/blob/master/device_linux.go


## Ethernet
[USERSPACE]
  * https://github.com/jcudit/gotcp - 
[FILE DISCRIPTOR]
  * https://github.com/justincormack/go-memfd - This is a Go library for working with Linux memfd, memory file descriptors.
These provide shareable anonymous memory, which can be passed around via file descriptors, and also locked from write or resize. They are designed to let programs that do not trust each other communicate via shared memory without issues of naming, truncation, or race conditions due to modifications.

_NETSTACK_
  * https://github.com/ggaaooppeenngg/netstack https://github.com/google/netstack Netstack is a network stack written in Go.

_ETHERNET_
  * https://github.com/songgao/ether ether is a go package for sending and receiving ethernet frames.
  * https://github.com/kopwei/gonet - very nice go networking library
  * https://www.kernel.org/doc/Documentation/networking/tuntap.txt
  * https://github.com/traetox/goTunTap - nice tun tap example with a VPN example
  * https://github.com/lab11/go-tuntap - another decent example
  * https://github.com/jamescun/tuntap - nice implementation
  * https://github.com/songgao/water - water is a native Go library for TUN/TAP interfaces.
  * https://github.com/guilhem/tentacool - tentacool is a Go server controlled via RESTful API through a Unix Domain Socket.

_SOCKETS_
  * https://github.com/arcpop/rawsocket
  * WOW SUPER RAW https://github.com/Hansal/Linux_Accounting/blob/master/sockets.go
* **low level as fuck networking**
  * https://github.com/songgao/ether/blob/master/dev_test.go
  * https://github.com/dutchcoders/netstack/blob/master/samples/http.go - assemble packets and do HTTP server
  * https://github.com/dutchcoders/netstack - completely custom network stack, supports all kinds plus custom
  * https://github.com/unigornel/go-tcpip - pure go network stack
  * https://github.com/songgao/ether/blob/master/dev_test.go

_NET_
https://github.com/joshlf/net/blob/master/ethernet.go - this may be the ticket

[Socket]
  **zsocket** Zsocket is a zero-copy ring buffer from a memory mapped file. It also contains utility functions and types to help with a handful of layer 2, 3, and 4 types. It is like libcap/pcap, except easier for working with these layers and doing packet injection. 
  _NOTE_ __Zsocket doesn't contain C/C++, its lock-free and thread-safe and has no dependencies.__
  GIT https://github.com/nathanjsweet/zsocket

[Kernel Module]
  **lmod** (Linux Module) Linux modules handling library, works with /sys/modules and the parameters folder in the module folder. It is very minimal but it is a good start, lets you see all modules, check if a given module is loaded and what parameters are loaded for a given module. 
  _NOTE_ __Probably can contribute and build ontop of this but one must learn more about the /sys/module folder system within linux.__
  GIT https://github.com/gsora/lmod

## TCP/IP

  **xtcp** Custom TCP Server with some nice defaults and features that function at a per packet level.
  GIT https://github.com/xfxdev/xtcp

  * https://github.com/gsora/lmod - works with kernel modules

_TCP_CONNECTIONS_
  * https://github.com/typetypetype/conntrack Keep track of active TCP connections (by talking to the ip_conntrack kernel module).
  * https://github.com/ehazlett/circuit - Circuit manages networks for runc. CNI network management (define and manage CNI networks and connectivity) CNI compatible (use CNI plugins) Quality of service management for networks and container interfaces Load balancing using IPVS
  * https://github.com/petar/GoTeleport - Caches connections and keeps unreliable connections alive and send along data when it comes back up.
_AF_PACKET_
  * https://github.com/google/gopacket/blob/master/afpacket/afpacket.go - zerocopy custom afpacket implenetation attached to a interface

# Vritual switch
  * https://github.com/inercia/divs
  * https://github.com/chzyer/next
  * https://github.com/enricomariafusi/v-switch 
  * https://github.com/AudriusButkevicius/pfilter - VIRTUAL CONNECTIONS FROM SINGLE PHYS

_ROUTE_
  * https://github.com/moriyoshi/routewrapper - control routes nicely

_DEV_UPINPUT_
  * https://github.com/galaktor/gostwriter - inject key invesnts
  * Multiplexing connections through single tcp connection https://github.com/zenhotels/astranet - astranet is a package for managing highly concurrent independent network streams. Millions of independent data streams between two host machines using one tcp connection only; An embedded service discovery system; NAT traversal capabilities for connecting any two machines without direct route between them using a trusted relay;
  * https://github.com/xtaci/smux - Smux ( Simple MUltipleXing) is a multiplexing library for Golang. It relies on an underlying connection to provide reliability and ordering, such as TCP or KCP, and provides stream-oriented multiplexing. The original intention of this library is to power the connection management for kcp-go.

[PACKET FILTER]
  * https://github.com/go-freebsd/pf
  * https://github.com/dutchcoders/netstack/blob/master/samples/http.go - assemble packets and do HTTP server
  * https://github.com/dutchcoders/netstack - completely custom network stack, supports all kinds plus custom
  * https://github.com/cxfksword/httpcap
https://github.com/gdm85/go-websockproxy - websockets tied to a interface
  * https://github.com/dedis/onet - ooks good
  * https://github.com/enricomariafusi/v-switch/blob/master/tap/libtap.go = works wioth no netlink api

[VETH]
  * https://github.com/kopwei/gonet/blob/master/veth.go
  * pcap packet layer examples - https://github.com/mynameiscfed/go-cp-analyzer/blob/master/main.go
  * https://github.com/mbucc/vufs

# Tor Bridge
  * https://github.com/jessfraz/onion/blob/master/tor/bridge.go
  * very nice looking vpn, tinc but without presetup - https://github.com/enricomariafusi/v-switch
  * https://github.com/chzyer/next - vpn
  * https://github.com/defgrid/openvpn-peer - coole xample, with gossip
  * https://github.com/marten-seemann/quic-conn - fancy replacement for tcp

# Packet intercept
  * https://github.com/nikofil/gopacketcache/blob/master/examples/main.go - this can be used to cache packets from an interface. after they are pulled, they can be put anywhere.
  * https://github.com/tgogos/gopacket_nfqueue/blob/master/main.go - nfqeue gopacket
  * https://github.com/zaftzaft/gopacket-training/blob/master/arp/arp.go - go packet examples

_TCP INJECTION_
  * https://github.com/david415/HoneyBadger 
    HoneyBadger is primarily a comprehensive TCP stream analysis tool for detecting and recording TCP injection attacks.
    This git repository also includes a variety of prototype TCP stream injections attacks.
  * https://github.com/ginuerzh/gost - socks5 obsf4 http ssh simple tunnel written

_DEVICES_
 * FIFO Devies https://github.com/containerd/fifo - Go package for handling fifos in a sane way.
 * Block devices https://github.com/bazil/fuse - It is a from-scratch implementation of the kernel-userspace communication protocol, and does not use the C library from the project called FUSE.
  * https://github.com/cxfksword/httpcap
  * Interception https://github.com/troyxmccall/gogospoofdns dsniff is a collection of tools for network auditing and penetration testing. dsniff, filesnarf, mailsnarf, msgsnarf, urlsnarf, and webspy passively monitor a network for interesting data (passwords, e-mail, files, etc.). arpspoof, dnsspoof, and macof facilitate the interception of network traffic normally unavailable to an attacker (e.g, due to layer-2 switching). sshmitm and webmitm implement active monkey-in-the-middle attacks against redirected SSH and HTTPS sessions by exploiting weak bindings in ad-hoc PKI. 
  * https://github.com/skycoin/net - General UDP/TCP server/client library

_TCP_handshake_
  * https://github.com/tevino/tcp-shaker - sync sync-ack rst health checking
    
_PROXY_
  * https://github.com/LiamHaworth/go-tproxy Golang TProxy provides an easy to use wrapper for the Linux Transparent Proxy functionality.

Transparent Proxy (TProxy for short) provides the ability to transparently proxy traffic through a userland program without the need for conntrack overhead caused by using NAT to force the traffic into the proxy.
Step 1 - Binding a listener socket with the IP_TRANSPARENT socket option

Preparing a socket to receive connections with TProxy is really no different than what is normally done when setting up a socket to listen for connections. The only difference in the process is before the socket is bound, the IP_TRANSPARENT socket option.

syscall.SetsockoptInt(fileDescriptor, syscall.SOL_IP, syscall.IP_TRANSPARENT, 1)

Step 2 - Setting the IP_TRANSPARENT socket option on outbound connections

Same goes for making connections to a remote host pretending to be the client, the IP_TRANSPARENT socket option is set and the Linux kernel will allow the bind so along as a connection was intercepted with those details being used for the bind

_PACKETS_
  * Example of using gopacket, the quentessential packet library in go https://github.com/mynameiscfed/go-cp-analyzer
  * Example of bidrectional stream resassembly https://github.com/orivej/tcpassembly/blob/master/bidistream/bidistream.go
  * Example of reassembly and storage in elasticsearch https://github.com/ChrisRx/mongoose
  * Example sending packet https://github.com/ebiken/go-sendpacket
  * very nice example - https://github.com/chenong/goSniffer/

_PCAP_
  * https://github.com/f47h3r/swashbuckler/ - pcap example with bpf filtering for icmp packets


