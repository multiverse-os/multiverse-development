# Multiverse Onion Network 

https://github.com/as/torgo - pure go tor?
https://github.com/andres-erbsen/torch - pure go tor?

## Encapsualted HTTP
https://github.com/getlantern/enhttp/
// Package enhttp provides an implementation of net.Conn that encapsulates
// traffic in one or more HTTP requests. It is conceptually similar to the older
// https://github.com/getlantern/enproxy but differs in that it supports HTTP
// servers which don't support Transfer-Encoding: Chunked on uploads.
//
**buffering**
https://github.com/abcum/bump - could help with loading those big ass bmps
https://github.com/abcum/bump/blob/master/bump_test.go
## Cahce
https://github.com/abcum/cachr - **best caching library**
## DB
https://github.com/abcum/rixxdb
## web ui
https://github.com/go-http-utils/cookie
## KCP
https://github.com/xtaci/kcp-go
KCP Tun https://github.com/xtaci/kcptun
## Smux
https://github.com/xtaci/smux 
Simple Stream Multiplexing for golang

## UDT
https://github.com/oxtoacart/go-udt
##
https://golang.org/pkg/net/?m=all
## AutoTLS for Gin
https://github.com/gin-gonic/autotls
## SSH TUnenling
https://github.com/nicksellen/gotunnel
## WebRTC
https://github.com/xhs/gortcdc

## gin 

https://github.com/itsjamie/gin-cors
https://github.com/DeanThompson/ginpprof
https://github.com/Massad/gin-boilerplate

## Chord
https://github.com/armon/go-chord
## systemd
https://github.com/coreos/go-systemd
## Ruby Config
https://github.com/k0kubun/itamae-go
## Cahce
https://github.com/bluele/gcache nice api
## gosctp
https://github.com/xhs/gosctp
## DB
https://github.com/blevesearch/bleve full text
https://github.com/cayleygraph/cayley graph
https://github.com/solher/arangolite

https://github.com/rqlite/rqlite

https://github.com/tidwall/buntdb


https://github.com/couchbase/goforestdb

https://github.com/krotik/eliasdb
https://github.com/fern4lvarez/piladb http rest api

https://github.com/couchbase/moss
https://github.com/jmhodges/levigo
https://github.com/siddontang/ledisdb - http rest api
https://github.com/dgraph-io/dgraph

https://github.com/siddontang/ledisdb

https://github.com/syndtr/goleveldb <very nice>

https://github.com/jasonmain/joltDB
## Devices 
https://github.com/mdlayher/vsock
## Crypto
https://github.com/cryptoballot/entropychecker
## Networking
[netlink]
https://github.com/subgraph/go-nfnetlink
[raw]
https://github.com/ryabuhin/golang_linux_rawsockets
Work with IP layer and ICMP/TCP/UDP packets on the linux raw socket level. Golang implementation *unix ping, tracert + is host alive 

[zero copy sockets]
https://github.com/nathanjsweet/zsocket


## kobjkect
https://github.com/mdlayher/kobject Userspace events occur whenever a kobject's state changes. As an example, events are triggered whenever a USB device is added or removed from a system, or whenever a virtual network interface is added or removed.
## Transparent proxy
https://github.com/LiamHaworth/go-tproxy
## mac address
https://github.com/subgraph/macouflage
## simple cotnainer
https://github.com/converseai/simple_container
### Onion Network Device



=====

**Important Libraries**
https://github.com/takatoh/boxmuller - Generates 2 psuedorandom that are legit 


https://github.com/mdlayher/raw - raw lets you get all packets from device (like pcap sniffing)
https://github.com/golang/go/blob/e49bc465a3acb2dd72e9afa5d40e541205c7d460/src/net/iprawsock.go

https://github.com/box-builder/box - use ruby to config, cool

# Go Networking
https://github.com/milosgajdos83/tenus tenus uses runc's implementation of netlink protocol. The package only works with newer Linux Kernels (3.10+) which are shipping reasonably new netlink protocol implementat

https://github.com/intel-go/nff-go NFF-Go -Network Function Framework for GO (former YANFF) 

https://github.com/aterlo/puregobpf

https://github.com/mdlayher/netlink 
https://github.com/jsimonetti/rtnetlink

https://github.com/xaionaro-go/netTree retrieve the tree of network interfaces in linux (like bridge enslaves a vlan-interface that uses a bond-interface that enslaves a physical) 

https://github.com/syxolk/ssh-keycheck - list all authorized keys, last login ip, last login time, etc

#### Stats / Monitoring 
[GOnetstat](https://github.com/elsonwu/GOnetstat)
data from /proc/net/tcp|6 and /proc/net/udp|6

#### Globbing & Path/String Matching
[doublestar](https://github.com/bmatcuk/doublestar)

#### Protocols

**Web Protocols**
	[HTTP]

	[Websockets]

	[Long-polling]
  > https://github.com/elsonwu/goio

	[SPDY]

  [gRPC]
  > https://github.com/grpc/grpc-go

   

##==     _______  __   __  _______                                                                   ##==
##==		|       ||  | |  ||       |                                                                  ##==
##==		|   _   ||  |_|  ||_     _|                                                                  ##==
##==		|  | |  ||       |  |   |                                                                    ##==
##==		|  |_|  ||       |  |   |                                                                    ##==
##==		|       ||   _   |  |   |                                                                    ##==
##==		|_______||__| |__|  |___|  [ ONION HASH TABLE ]                                              ##==
##==                                                                                                 ##==
##===================================================================================================##==

**Standard Library: Socks5 Proxy**
https://github.com/golang/net/blob/master/proxy/socks5.go

	  SUMMARY       *Onion Hash Table* **(OHT)** is **decentralized hash table* **(DHT)** with routing
		              over the *Onion Network* **(onet)**. **OHT** is the backbone of the **OHT**
	                Decentralized Applciation Framework [TODO][FIXME: Pick a better name] and shares
		              core components that are required and maintained by the in the Multiverse OS
		              project.

	  COMPONENTS    [1] *Onion Router* **(OR)**, provided by an *Onion/Tor* device
											[1a]
											[1b]
				          [2] *Peer Lattice* **(PL)** or *Peer Mesh Networking* **(PMN)**
	                    [TODO][FIXME: Is Lattice The Best Term?]
				          [3] *Onion Account System* based on new onion address specification and the
	                    *Ephemeral Key Tree* **EKT** system
							          [3a][Key System]
								      	  [*][oniongen-go][https://github.com/lostinblue/oniongen-go]
								      	  <IMPORTANT> custom built onion generation for key generation for account
	                        system
				          [4]  
				          [5]

	_OHT RESEARCH_  

      TOPICS      **Internet protocol suite (IPS)**

		              **Virtual Network Interfaces (VNI)**
		              **OHT** will utilize different classes of userspace **VNI** that completely bypasses
		              the kernel, a psuedo network device that overlays existing networking to provide
	                additional features to provide Multiverse OS with complete *Software Defined Net-*
		              *working* **(SDN)**.

		              The kernel has its own **VNI** table, but the **OHT** interface **MUST** bypass
		              this table and route **ALL** packets within *Userland*.

		              *Userland* and bypass the kernel networking modules entirely limit the attack
		              surface and importantly specific known vectors of attack which allow for attackers
	                to escape executed **VM** environment and move into the host environment. The most
	                common type of breakout found with KVM/Qemu has been specially crafted network
	                packets passing through the Linux kernel space causing a breakout exploit.

		              **Virtaul Loopback Interface (VLI)**
		              Implementations of the **IPS** include a **VNI** through which network applications
		              can communicate when executing on the same machine. It is implementeedentirely
		              within the OS's networking software and passes no packets to any *netwwork inter-*
		              *face controller* **(NIC)**.

		              Any traffic that a computer program sends to a loopback IP address is simply and
		              immediately passed back up the network software stack as if it had been received
		              from another device.

		              POSIX (or sometimes called Unix-like) systems usually name this loopback interface
	                **lo** or **lo0**.

https://github.com/sparrc/go-ping

[firewall](https://github.com/Gouthamve/go-firewall)
A firewall using nfqueue (not iptables), this makes it pretty low level or atleast very high power.
#### Battle Net
https://github.com/mitchellh/go-bnet
#### Git
[git2go](https://github.com/libgit2/git2go)
[gitfs](https://github.com/hanwen/gitfs)

bindings for libgit2
#### (DHT) Chord lookup algorithm
[go-chord](https://github.com/taylorflatt/go-chord-implementation)


#### gRPC
https://github.com/grpc/grpc-go
[go-chat](https://github.com/taylorflatt/go-chat)
gRPC + protocol buffers based chat

[remote-shell](https://github.com/taylorflatt/remote-shell)
gRPC + protocol buffers based remote shell
#### Tunnels
[tcpovericmp](https://github.com/Maksadbek/tcpovericmp)
tunnel over icmp
[tlstun](https://github.com/jsimonetti/tlstun)
socks tunnel client/server over websockets over http/tls
#### FS
[memfs](https://github.com/zbiljic/memfs)

#### Globbing & Path/String Matching
[doublestar](https://github.com/bmatcuk/doublestar)

#### DHT
https://github.com/mh-cbon/dht
https://github.com/matrix-org/dendrite
## Multiverse OS Notes
General notes during developing, testing and researching multiverse OS. 


[libvirt-wireshark] <IMPORTANT>
Check out this debian package, could be useful and avoid having to rebuild existing functionality.


# Crypto OHT
https://github.com/GridProject/GridKernel


# Audio Anlsysi

https://github.com/simonassank/aubio-go

# Numeric Analysis

https://github.com/unixpickle/num-analysis

https://github.com/reiver/go-numeric/blob/master/numeric.go


https://github.com/akualab/dataframe !!!!


### CHAT EXAMPLE

https://github.com/esimov/gifter - play gifs in the terminal
https://github.com/tombh/texttop
https://github.com/Francesco149/sharenix
========================================================================================================
========================================================================================================
========================================================================================================
========================================================================================================
--------------------------------------------------------------------------------------------------------
# OHT Browser
*This shows how to disable javascript! This could be the winner!*
https://github.com/OneOfOne/webview/blob/master/webview.go

# Users/Authentication
[crypto11]
https://github.com/thalesignite/crypto11
https://github.com/miekg/pkcs11
This is a Go implementation of the PKCS#11 API. It wraps the library closely, but uses Go idiom were it makes sense. It has been tested with SoftHSM.
    [LIMITED SSH SERVER!!!](https://github.com/GraveRaven/scpdrop)
		could be used in OHT for limited access and on the fly SCP based file transfer but also for very limited agent/provisioning in multiverse

#### SCREEN SHARING

[!][https://github.com/dominikh/xcapture]
Xcapture is a command-line driven X11 window recorder, outputting a raw video stream on standard out for processing by other tools.

## SSE
[sse][https://github.com/JanBerktold/sse]
#### PROTOCOLS
  **FTP**
		[ftp](https://github.com/fclairamb/ftpserver)

  **TFTP**
	  [tftp][https://github.com/mdlayher/tftp]
  **MPTCP**
  	[mptcp][https://github.com/mdlayher/mptcp]
	**SSHTTP**
    [sshttp][https://github.com/mdlayher/sshttp]
	**ZSTORE**
    [zstore][https://github.com/mdlayher/zstore]

## sRPC (Generic RPC?)
    [srpc][https://github.com/grandcat/srpc]
# Auditing
    [libaudit-go][https://github.com/mozilla/libaudit-go]

# Networking

## Conntrack (total number of connections)
	  [conntrack][https://github.com/typetypetype/conntrack]


## Service Discovery
https://github.com/grandcat/zeroconf

https://github.com/mdlayher/ndp
Package ndp implements the Neighbor Discovery Protocol, as described in RFC 4861. MIT Licensed. https://tools.ietf.org/html/rfc4861

## ALG
https://github.com/mdlayher/alg

## DHCP/DHCP6
https://github.com/mdlayher/dhcp6

## ARP
https://github.com/mdlayher/arp

## Routing Table
https://github.com/miekg/rip/blob/master/rip1_test.go

## Netlink (userspace)
https://github.com/mdlayher/netlink

## VSOCK VirtIO
https://github.com/mdlayher/vsock

## WIFI
https://github.com/mdlayher/wifi

## Kernel
[kobject]
https://github.com/mdlayher/kobject Package kobject provides access to Linux kobject userspace events.
Userspace events occur whenever a kobject's state changes. As an example, events are triggered whenever a USB device is added or removed from a system, or whenever a virtual network interface is added or removed.

## DNS
[DNS Server]
https://github.com/skynetservices/skydns
https://github.com/miekg/dns
*Examples* https://github.com/miekg/exdns

## Firewall/IPTables

## Devices
https://github.com/google/netstack

https://github.com/google/netstack/blob/master/tcpip/sample/tun_tcp_echo/main.go < Example of how to setup a userspace network device >

## P2P 
https://github.com/matrix-org/dendrite

## Chat

## Video
https://github.com/webtorrent/webtorrent
https://github.com/anacrolix/torrent
https://github.com/gitsummore/nile.js
========================================================================================================
**Open Source S3**
[!][https://github.com/minio/minio]

## Anonymous Geolocation
Since OHT will be over Tor, location of peers will either be voluntary or inferred by ping locations and latency
[orca][https://github.com/bbengfort/orca]
Echolocation of device with static nodes and network latency. 

## Filesystem
[mirrorfs][https://github.com/bbengfort/mirrorfs]
mirrors whatever happens in another dir 
[memfs][https://github.com/bbengfort/memfs]

**file system changes notify**
[fsnotify][https://github.com/fsnotify/fsnotify]

## Distributed Computing
[!][https://github.com/chrislusf/gleam]
(^)[https://github.com/chrislusf/gleamold]
Gleam is a high performance and efficient distributed execution system, and also simple, generic, flexible and easy to customize.
Gleam is built in Go, and the user defined computation can be written in Go, Unix pipe tools, or any streaming programs.
[!][https://github.com/acook/blacklight]
forth like VM

## Constrained networks protocol
https://tools.ietf.org/html/rfc7252
https://github.com/dustin/go-coap
https://github.com/zubairhamed/canopus
## Group Media / Codecs
#### Group Music PLayer
[oto][https://github.com/hajimehoshi/oto]

[web auidio][https://github.com/nobonobo/webaudio]

**codecs**
[go-mp3][https://github.com/hajimehoshi/go-mp3]

#### Group Video
[youtube][https://github.com/iocat/youtube]
youtube iframe api using gopherJS to avoid JS

[gophervideo][https://github.com/csos95/gophervideo]
## Compress
https://github.com/klauspost/compress
**codecs**
## Account System
    [Unique IDs]
	    [nuid][https://github.com/nats-io/nuid]

## Game Example (Decentralized MMO) 
	  [craft][https://github.com/microo8/craft]
	  an hobby minecraft "clone" in go 

    [gonet](https://github.com/xtaci/gonet)

## Reed Solomon
https://github.com/klauspost/reedsolomon
## Distributed Web UI and API
	  [teeproxy][https://github.com/chrislusf/teeproxy]
	  teeproxy is a reverse HTTP proxy. For each incoming request, it clones the request into 2 requests,
	  forwards them to 2 servers. The results from server A are returned as usual, but the results from
	  server B are ignored.

## Limited VM's for group processing
	  [macinery][https://github.com/denkhaus/machinery]
	  Very Very nice Machinery is an asynchronous task queue/job queue based on distributed message
	  passing. So called tasks (or jobs if you like) are executed concurrently either by many workers on
	  many servers or multiple worker processes on a single server using Golang's goroutines.

## Channels / PubSub
	  [go-pubsub][https://github.com/cloudfoundry/go-pubsub]
	  PubSub publishes data to subscriptions. However, it can do so much more than just push some data to
	  a subscription. Each subscription is placed in a tree. When data is published, it traverses the tree
	  and finds each interested subscription. This allows for sophisticated filters and routing.

## Server Side Rendering & Automated Interneting
	  [ottox][https://github.com/stretchr/ottox]
	  Plugins for the Otto Go JavaScript parser and interpreter

#### Secure Access To Onions
    [signature][https://github.com/stretchr/signature]
    Signature secures web calls by generating a security hash on the client (using a private key shared
	  with the server), to ensure that the request is geniune. Only a client who knows the private key
	  will be able to generate the same security hash.

###################################################
## Database
**Embedded**
[!][https://github.com/1lann/cete]
Cete is an easy-to-use, lightweight, pure Go embedded database built on Badger for use in your Go programs. Unlike most other embedded database toolkits for Go, Cete is schemaless, yet still blazing fast. It's great for cases where you need a fast, on-disk, embedded database. Cete is licensed under the MIT License.

**K/V**
[!][https://github.com/colinmarc/cdb]
(^)[https://github.com/chrislusf/cdb64]
This is a native Go implementation of cdb, a constant key/value database with some very nice properties, but without the 4GB size limit.
[redis][https://github.com/go-redis/redis]
go redis client

**Distributed**
[!][https://github.com/cockroachdb/cockroach]

**Memory**
https://github.com/kelseyhightower/memkv

###################################################
## General Networking ##
[!][https://github.com/hsheth2/gonet]
[go-networking][https://github.com/vladimirvivien/go-networking]
	<important> Very nice library for dealing with IP, protocols, etc

## Handling Packets
https://github.com/google/gopacket

## Onion Network Device
[!][https://github.com/unigornel/go-tcpip]
https://github.com/nathanjsweet/zsocket

[Low Level NEtworking]
[!][Fake Interfaces, pure logic network devices entirely in userspace. That means you don't need root, yeah, really.]

		// BaseClient is a streaming telemetry client with minimal footprint. The
		// caller must call Subscribe to perform the actual query. BaseClient stores no
		// state. All updates must be handled by the provided handlers inside of
		// Query.

*Example:*
https://github.com/openconfig/gnmi/blob/master/client/gnmi/client_test.go
*In this example, we will see a client/connection getting initalized, then udpates getting sent down it.*

## Multiplexing
https://github.com/hashicorp/yamux Yamux (Yet another Multiplexer) is a multiplexing library for Golang. It relies on an underlying connection to provide reliability and ordering, such as TCP or Unix domain sockets, and provides stream-oriented multiplexing. It is inspired by SPDY but is not interoperable with it.
-------------------------------------
## NTP Client
https://github.com/vladimirvivien/go-ntp-client

## NTP Server
https://github.com/beevik/ntp
----------------------------------
#### Network Protocols
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

	[gRPC]
		[echo][https://github.com/bbengfort/echo]
	  gRPC echo example
		[sping][https://github.com/bbengfort/sping]
	  Simple example of secured communication with gRPC and SSL/TLS 

	[WebRTC]
		[go-webrtc-datachannel][https://github.com/coreos/go-webrtc-datachannel]
		basic and old go webrtc demo
## HTTP
**HTTP Based Networking (Websockets, gRPC, webRTC, HTTP)**

#[webrtc]
[ssh+webrtc][https://github.com/nobonobo/ssh-p2p]

#[gRPC]
[gRPC Device Example]
https://github.com/openconfig/gnmi

**web app frameworks**
[kite][https://github.com/koding/kite]
[sessions][https://github.com/rivo/sessions]
very very nice cookie library
[duplo][https://github.com/rivo/duplo]
detect duplicate (or similar) images using hashes AND other techniques

**http**
[puma-dev][https://github.com/puma/puma-dev]
lowest level go http server I have seen, its fucking awesome
**Middleware**



[negroni][Echolocation of device with static nodes and network latency]
This is the library formerly known as github.com/codegangsta/negroni -- Github will automatically redirect requests to this repository, but we recommend updating your references for clarity.

Negroni is an idiomatic approach to web middleware in Go. It is tiny, non-intrusive, and encourages use of net/http Handlers.
##################################
## P2P Data / Distributed Data / Distributed KV
[!][https://github.com/minio/minio]
[!][https://github.com/utamaro/shards]
Use Reed Solomon to encode and decode file data
[sftp][https://github.com/pkg/sftp]
[consul][https://github.com/hashicorp/consul]
overkioll but may have useful code

