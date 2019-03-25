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
#### Battle Net
https://github.com/mitchellh/go-bnet
#### Git
[git2go](https://github.com/libgit2/git2go)
[gitfs](https://github.com/hanwen/gitfs)

bindings for libgit2
#### (DHT) Chord lookup algorithm
[go-chord](https://github.com/taylorflatt/go-chord-implementation)
## WebUI 
https://github.com/abcum/webkit/ - webkit that supports easy ability to disable javascript execution!!!
## transparent proxy
https://github.com/mbland/hmacproxy
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
#### Notes / Research
* Binding to :0 will ask the kernel for an available port within the ephemeral port range
the kernel will assign it a free port number somewhere above 1024.

	// tell pinger that it is privileged.
	// NOTE: You must run `setcap cap_net_raw=+ep pocket-loss-monitor`

	
## Pkt.Connections
[shared-conn](https://github.com/jdhenke/shared-conn)
Allows you to share a net.Conn between different processes easily
	
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

**codecs**
## Account System
    [Unique IDs]
	    [nuid][https://github.com/nats-io/nuid]

## Game Example (Decentralized MMO) 
	  [craft][https://github.com/microo8/craft]
	  an hobby minecraft "clone" in go 

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

