# Networking Linux

  _MULTIPLEX (IE SSH, HTTP, OPENVPN over SAME PORT)_
  * https://github.com/shawnl/multiplexd  run ssh, https, and openvpn on the same port 

  _KERNEL_MODULES_
  * https://github.com/barnex/linux-kernel-modules/tree/master/ktest - cool example of hello world device kernel module with C examples
  _MEMORY_
  * https://github.com/mephux/memory 
* **Daemonization**
  * https://github.com/farseer810/go-daemonizer - Nice daemonizer

* **OS**
  * https://github.com/insionng/gopher-os - OS concept 

  https://github.com/achilleasa/gopher-os - tickless kernel 

* **Webframework**
  * https://github.com/kataras/iris - http ready for go

* **Database**
  * Bolt DB Elastic search https://github.com/blevesearch/bleve 

  * https://github.com/graphql-go/graphql - graph ql

  * https://github.com/spaolacci/murmur3 - murmur3 native go

  __KEY_VALUE_
  * https://github.com/dgraph-io/badger - fast pure go kv
* **TOTP**
  https://github.com/yulvil/totpjs/blob/master/totpjs.go - provides web ui for 1 pre configured OTP

* **Merkle Tree**
  * https://github.com/ymmah/trillian - MASSIVE merkle trees

  * https://github.com/cbergoon/merkletree - nice implementation
  * https://github.com/porkchop/merklepatricia - with patricia radix tree

  * https://github.com/Daeinar/merkle-tree-tools/blob/master/merkletree.go - with consistent hashes to make things better

  * https://github.com/betawaffle/go-smt - sparse merkle trees

  * https://github.com/rodrigue-tech/hashtree -c omplex merkle tree implmentation

* **Pipelines/Steps/Chain**
  * https://github.com/mitchellh/multistep - individual discrete steps that can be chained together.
__GENERAL STUFF FOR PROGGS__
* **Diffie Helmen Merkle key exchange edch**
  * group key exchange https://github.com/corvuscrypto/group_ecdhm
* **NDB**
  * https://github.com/abligh/gonbdserver - pure go nbd

* **Kernel**

  _BAREBONES_
  * https://github.com/jjyr/bootgo
* **SHRED/WIPE/DESTROY DISK**
  * https://github.com/traetox/goDiskDestroy

  

* **Events**
  * https://github.com/SkynetAtack/go-events/blob/master/main.go - events
# LEDGER!
https://github.com/Prakhar0409/Distributed-Ledger
# PEERFLIX!!!!!
https://github.com/Sioro-Neoku/go-peerflix - watch movie your downloading in VLC
## Sig Handling Libs
https://github.com/bluele/gsignal/blob/master/examples/example.go - very nice
https://github.com/ianchildress/redshirt - another good simple example, similar to above

https://github.com/antwhite/trafficlight/blob/master/main.go - very nice EXAMPLE OF VANILLA sig handling usage, probablyb est route
## Events Emit
https://github.com/jonhoo/go-events/blob/master/examples_test.go - cool event system because it has channels in a tree, so you can have listening at higher levels.
## Config	
https://github.com/gopk/config - super super nice setting and getting values. 

https://github.com/Santiclause/goconfig - There are sighup signals that tell the program to reload THIS ACTUALLY RESPECTS THAT, steal that code and combine it with the desired.
https://github.com/olebedev/config - very very nice looking JSON and YAML config loader. This could be the one

https://github.com/tsuru/config - YAML only (can just covert to yaml) but has a nice read and readandwatch for live updating. Has nice GetBool to set the datatype of the values being pulled instead of just assuming strings.

https://github.com/gravitational/configure - nice, types are defined by the struct being loaded, supports env, cli flags and yaml and lets you load them in your rpefrred order
---
# Vritual FS
https://github.com/olivere/vfs - sandboxing

https://github.com/chai2010/vfs - cool examples
https://github.com/kimiazhu/vfs - cool examples
https://github.com/blang/vfs
https://github.com/bazelbuild/sandboxfs - virtual fs for sandboxing !!!!!
https://github.com/magicalbanana/vfs
https://github.com/milochristiansen/axis2 - virt fs api
https://github.com/sourcegraph/ctxvfs  -vfs with context

__VIRT MACHINE__

* **Virtual Machines** Develop a solid virtual machine setup
https://github.com/wwaites/nopfs
  _ PROVISION_
https://github.com/domainr/epp

 * super simple pure go container with json config https://github.com/converseai/simple_container

 * https://github.com/docker/libcontainer - proto-runc

 * https://github.com/docker/libnetwork - Networking for containers

 * https://github.com/felipejfc/go-container - pure GO container, comes with alpine, nice ntworking, new, super light


 **RKT**
 * https://github.com/rkt/rkt

 **RUNC**
  https://github.com/containerd/go-runc
  https://github.com/jessfraz/netns - Runc hook for setting up default bridge networking.

 * https://github.com/docker/machine - Machine lets you create Docker hosts on your computer, on cloud providers, and inside your own data center. It creates servers, installs Docker on them, then configures the Docker client to talk to them.

 * https://github.com/containerd/containerd - containerd is an industry-standard container runtime with an emphasis on simplicity, robustness and portability. It is available as a daemon for Linux and Windows, which can manage the complete container lifecycle of its host system: image transfer and storage, container execution and supervision, low-level storage and network attachments, etc.

containerd is designed to be embedded into a larger system, rather than being used directly by developers or end-users.

  _CLAER_CONTAINERS_

  https://github.com/containers/virtcontainers

  https://github.com/clearcontainers/runtime - intels clear container virtcointers - virtual hardware but containers combo

  __

  * razor thin pure go kvm https://github.com/kitschysynq/darity

  * https://github.com/0xef53/phoenix-guest-agent - qemu-kvm virtual machines guest side agent
* **COW**
  * https://github.com/krishnasrinivas/constor - cow userspace filesystem layers
* **Web Virtual machine panle**
  * https://github.com/abates/vpanel - angulur js

  _libvirt_

  * https://github.com/farazfazli/CyanOcean - libvirt kvm vms management

 _container_

  * https://github.com/lastbackend/lastbackend - slick as fuck CI CD CLI amazing UI containers


__TOR__


  * https://github.com/rdkr/oniongen-go/blob/master/main.go v3

  * https://github.com/gosharplite/onion

  * https://github.com/Yawning/onionwrap/blob/master/main.go

  * https://github.com/bfix/Tor-DNS - Tor Socks5 proxy resolve names

  * https://github.com/BrassHornCommunications/OnionWatch - stats

  * https://github.com/pgerber/sandboxed-tor-browser 

  * Tor partial implementaton - https://github.com/phoenix1342/TOR-Browser

   * https://github.com/equk/torjail - runs inside docker rpbobly uses xepher to acces sit

   _TOR_DEVICE_
   * https://github.com/liudanking/tuntap/blob/master/exmple/tunclient.go - some examples
   * https://github.com/lixin9311/simplevpn  -exmaples


   _WHAT?_
   * https://github.com/insionng/tor 

   _ONION_
   * https://github.com/cyphar/mkonion - tor onion service for docker containers

   * https://github.com/jessfraz/onion - route all traffic through tor in container

# Tor 

* DW scanning https://github.com/s-rah/onionscan

* https://github.com/Crypt-iQ/hidden-service

(One idea is using OHT, add an extra hop that is done across several connections and sent backa nd reassmbled)

* tor control https://github.com/gwitmond/textproto

* https://github.com/pombredanne/vidalia - get exit node list, sort them by threat on cloudflare, this can make a better client

* NEWEST https://github.com/frozenkp/gotor

* super simple tiny tor client, works? https://github.com/mmcco/mutor/blob/master/main.go

* https://github.com/bfix/gpg2hs

* Tor DNS forwarder - https://github.com/rolandshoemaker/tdf

* nother special tor client in go https://github.com/andres-erbsen/torch

* https://github.com/sycamoreone/orc -best?>

* https://github.com/yydesa/torcat

* https://github.com/tswindell/go-torc

* https://github.com/codekoala/torotator 

* https://github.com/goshinobi/tor_multi

* https://github.com/TvdW/gotor -> newest https://github.com/tmc/gotor -> https://github.com/mmcloughlin/pearl (newest)

* https://github.com/postfix/goControlTor - control tor instance
  https://github.com/Yawning/or-ctl-filter

( https://github.com/nogoegst/tor-getinfo/blob/master/tor-getinfo.go
  	 https://github.com/yawning/bulb

* https://github.com/shw700/tortime/blob/master/tortime.go

* https://github.com/nogoegst/torsh

* https://github.com/prsolucoes/go-tor-crawler
  https://github.com/ak1t0/flame

* https://github.com/phoenix1342/TOR-Browser

* https://github.com/nogoegst/onionutil

* https://github.com/freedumbhost/torhost-control

( https://github.com/Pholey/distribuTor

* https://github.com/Yawning/tor-fw-helper

* https://github.com/willscott/onionproxy

* onion balance https://github.com/nogoegst/avant

* https://github.com/nogoegst/onionize

* https://github.com/OpenBazaar/go-onion-transport

* https://github.com/rdkr/oniongen-go

* https://github.com/ViGrey/deadmann 

* https://github.com/subgraph/roflcoptor

__NETWORKING__

**OHT**
  * Packet level software load balance - https://github.com/sipb/spike

  * F2F Friend-to-Friend Networking - https://github.com/brendoncarroll/go-f2f
  * UDP Networking
    * https://github.com/fhaynes/bifrost - A UDP networking library meant for games in Go (Game networking is interesting, it uses UDP but then implements some of the TCP protocol, would be useful for high latency, high loss networks, like OHT)

    * https://github.com/aurelien-rainone/udpnet - High quality UDP based network protocol for games. udpnet aims to provide some building blocks to create a UDP-based network game procotol: virtual connection reliability packet ordering congestion avoidance

  * _CONSISTENT_HASHING_
    - https://github.com/dgryski/go-maglev/ - consistent hashing library

* **Database**
---

* **HTTP** Raw HTTP Clients/Servers

* https://github.com/jbussdieker/go-rawhttp 
* https://github.com/aglyzov/ws-machine/blob/master/examples/blockchain/blockchain.go - nice websockets client/server with statemachine

  - Key/Value
    * https://github.com/cockroachdb/cockroach - cockroach SQL server, highly consistent acid	

    * etcd https://github.com/coreos/etcd
  - Graph

* **Pub/Sub**
  - https://github.com/secondbit/peter - A distributed pub/sub network written in Go. Built on Pastry, modeled after SCRIBE. 

* **Entropy**
  * https://github.com/longsleep/entropyd - Ramdom entropy provider and consumer via network for Linux 

  * https://github.com/cryptoballot/entropychecker 

---

* **Websockets TUnneL**
https://github.com/beefsack/go-under-cover
https://github.com/namedwebsockets/networkwebsockets
* **http client with custom DNS**
https://github.com/codequest-eu/dnsdialer

* **9p fs**
https://github.com/sirnewton01/plan9adapter
https://github.com/akmistry/go-nbd
* **device to dialer**
https://github.com/FTwOoO/go-tun2io
https://github.com/kevin-cantwell/logio
* **Pentesting**


    - DNS Proxy - https://github.com/seedifferently/nogo - nogo blocks access to various sites (ads, tracking, porn, gambling, etc) by acting as a DNS proxy server with host blacklist support.

It requires minimal setup, and includes a simple web control panel for managing the host blacklist.

* https://github.com/lair-framework/lair - Collaborative attack framework


* **BitTorrent**



  * Magnet URL https://github.com/gitchs/torrent2magnet/blob/master/main.go

  * very nice unique with react based webui https://github.com/stratospark/torro

  * torrent fs - https://github.com/ring00/torrentfs == doesnt seem to work, but it does mount a drive, so close

  * Webui client https://github.com/bbucko/torrot

  * https://github.com/anacrolix/torrent - The Go torrent library, regularly updated, well miantianed, full featured

  * https://github.com/jackpal/Taipei-Torrent - doesnt compile

  * https://github.com/nicolov/torrent-playground - doesnt compile

  * https://github.com/BamboV/torrent_center - tracker and simple cleint

  * Decently complex, but unifinsiehd torrent client https://github.com/alexmarchant/torgo-legacy

  * https://github.com/fantomius/GoTorrent - simple CLI client - doesnt compile

  * torrent client with hookers and booze, webbased UI https://github.com/ecdsa521/gourmet - doesnt compile

  * https://github.com/hongjunChoi/bittorrent_client_go - simple recent lib - doesnt compile



  _LIBRARY_
  * https://github.com/Vigneshsam/libtorrent - very nice library 

  _Tracker_
  * https://github.com/drbawb/babou - Nice torrent tracker

  * https://github.com/leighmacdonald/mika -redis backend torrent trakcer for private trackers

  * Good example, very simple - https://github.com/sergeyignatov/simpletracker

  * Private tracker, seeder, supportive, etc

  _Self_hosted_torrents_
  * https://github.com/tricklecloud/trickle - Nice UI, manages collection of friends and the torrents they are sharing
  
  _TorrentFS_
  * https://github.com/uwedeportivo/torrentzip - Special because apparently regardless of platform (windows/linux/mac) and regardless of the metadata (created at, etc) if the same files are put into a zip, it will be exactly the same bits.

  * https://github.com/ring00/torrentfs - TorrentFS with FUSE

  _UTILITIES_
  * manipulating torrent feeds - https://github.com/laplaceliu/gotorrent

  * https://github.com/rmmmmpl/folivora - rss feeds

  * **js video playing** https://github.com/brion/ogv.js/

  * https://github.com/bmatsuo/bt.exp - mktorrent, bencode and such

  * https://github.com/middelink/go-parse-torrent-name - Prases the torrent name, able to pick out the quality, the scene group, the name of the show or whatever and everything else. Nice job

  * https://github.com/billyninja/parse_torrent parse torrent

  * https://github.com/gitchs/torrent2magnet - generate magnet

  * https://github.com/polvi/mktorrent - library for creating torrent files with webseeds

  * https://github.com/bbpcr/Yomato - tracker 

  * Torrent search engine https://github.com/DistributedSolutions/DIMWIT and client

  _HTTP_
  * Bittroent -> Streaming HTTP server https://github.com/dveselov/torrentino

  * https://github.com/steeve/torrent2http/blob/master/torrent2http.go - run a single torrent to host a website

  _DHT_
  * https://github.com/anacrolix/dht - bittrorrent dht implementation, i think its extracted from another lib

  _SIMPLE_
  * Simple torrent client https://github.com/jronak/Torrent

  * https://github.com/richardwilkes/torrent - simple, and recent torrent implementation
    BEP-3 The BitTorrent Protocol Specification
    BEP-23 Tracker Returns Compact Peer Lists
    BEP-27 Private Torrents
  
  * Simple incomplete client - https://github.com/Kelfitas/go-torrent

  * https://github.com/regisb/slivers - simple, its peerid generator code

* **Universal tools/GNU tools/Core**
* https://github.com/as/torgo - target windows plan9 osx linux, accessible homogeneous interopuable GNU coreutils

--- 


* **udev**
  * https://github.com/eikenb/udev-notify - udev notification

* **BACKUP**
  * https://github.com/marcopaganini/netbackup - backup with rsync rdiff and rclone


---
* **DEV ENV**
  * https://github.com/tockins/realize - run build watch file changes with custom paths with webui

---
* **Uroot* universal root mostly go except 4 bins https://github.com/u-root/u-root


---
* **Spin Up Containers VM servicse**
  * https://github.com/twa16/userspace - spin up small workspaces throw way for dev and dploy test
https://github.com/twa16/userspace-cli

upsin

---
**SCISI BLock**
* https://github.com/coreos/go-tcmu Go bindings to attach Go Readers and Writers to the Linux kernel via SCSI. http://linux-iscsi.org/wiki/Main_Page

It connects to the TCM Userspace kernel API, and provides a loopback device that responds to SCSI commands. This project is based on open-iscsi/tcmu-runner, but in pure Go. Native support for LIO in QEMU/KVM, libvirt, and OpenStack™ (setup, code) makes it an attractive storage option for cloud deployments. 

 
---

  - Network Devices (Hardrives)
  - Analysis/Security
    * Pentesting
    * Monitoring
      - https://github.com/mehrdadrad/mylg - myLG is an open source software utility which combines the functions of the different network probes in one network diagnostic tool.

  - P2P Libraries
    * DNS Seeding
      - https://github.com/gombadi/dnsseeder - Bitcoin DNS seeder


  - DNS
    * DNS Servers
      - https://github.com/longsleep/xudnsd - Very simple DNS server implementation
  - SSH Proxying/Reverse 
    * VPN
    * Reverse Proxy (Expose services with only outbound accessibility)
    * 

  - Local network media (music sync, movies)
https://github.com/GAumala/MediaServer



    * https://github.com/beck917/pillX - a simple & powerful network Library written in Go. UDP and TCP and WEbsockets, more general library 



* **P2p**
https://github.com/subutai-io/p2p

* **Networking** at a packet level, create devices, use raw packat and device access, for analysis using ring buffers, and low level filtering, multi-services on 1 port

  * https://github.com/arcpop/network - General network example. decent device/socket creation. 

  * https://github.com/ehazlett/libdiscover - p2p lib
  * https://github.com/ehazlett/junction - Example use of libdiscover


  _evdev_
  * https://github.com/jteeuwen/evdev - evdev is a pure Go implementation of the Linux evdev API. It allows a Go application to track events from any devices mapped to /dev/input/event[X].

  * https://github.com/gvalkov/golang-evdev - This package provides Go language bindings to the generic input event interface in Linux. The evdev interface serves the purpose of passing events generated in the kernel directly to userspace through character devices that are typically located in /dev/input/.

  _COMPRESSION_
  * https://github.com/dsnet/compress


* **tranports**
  * HTTP proxy that uses like facebook, and shit to pass data

* **NSQ**
  * https://github.com/nsqio/nsq NSQ realtime distributed messaging

* **awesome http server, full control**
https://github.com/puma/puma-dev


* **Queue**
  * Priority queue - https://github.com/insionng/prior

* **Websockets**
  * https://github.com/x3ro/websocket-tty - shell commands from websockets. 

# OHT
  * https://github.com/onionNet/oht

## HTTP pcap parse
__htp__
https://github.com/cxfksword/httpcap
https://github.com/ghedo/go.pkt

https://github.com/clearthesky/httpparse/blob/master/main.go
https://github.com/fdns/capture/blob/master/capture.go - has code for DNS
https://github.com/Acey9/apacket/blob/master/sniffer/sniffer.go
https://github.com/ga0/netgraph/tree/master/ngnet

https://github.com/buger/goreplay

https://github.com/olahol/capreq/blob/master/capture.go

https://github.com/ObjectIsAdvantag/smartproxy/blob/master/main.go

__dns__
https://github.com/bmorton/dnswatch


# HTTP management of fs
https://github.com/dolftax/summer

* **Filesystem encryption** 
 _Drive and fodler_
 * https://github.com/google/fscrypt - very nice

 _LUKS_
 *  A disk encryption utility that helps setting up LUKS-based disk encryption using randomly generated keys, and keep all keys on a dedicated key server. https://github.com/HouzuoGuo/cryptctl
  * https://github.com/FooSoft/vfs - VERSIONING file system userpsace 
# network blcok device
https://github.com/p0rtalgun/nbdd


# HTTPFS
https://github.com/prologic/httpfs

# HTTP 2 stream pub sub
https://github.com/technosophos/drift 
# Chat example
https://github.com/marcusolsson/tui-go - perfect example for terminal chat UI
* **Chat**
  * FULL FREAUTRED CHAT https://github.com/blamarche/assemble-web-chat

  * Messaging from browser to mobile - https://github.com/titan-x/titan

  * Stomp over websockets  - https://github.com/inhies/stompsocket

  * https://github.com/jkurambhatti/go-chat-app - chat

  * https://github.com/matter123/chat


__X11 UI__
https://github.com/BurntSushi/xgb - niiice
https://github.com/BurntSushi/xgbutil - another go binding for x
_session_
https://github.com/jouyouyun/simple-session

_lock_
https://github.com/chlunde/trlock

_x11cleint_
https://github.com/electricface/go-x11-client - pure go?
__P2P NETWORKING__

* **File Browser**
  _WEBBASED_
  * https://github.com/alphapeter/filecommander - more a JOSN 2.0 API 
  * https://github.com/kernel164/gofe - very clean new google UI 

  * https://github.com/hekar/pallet - works with caddy

* **SYNC** File syncing
  * https://github.com/zeisss/mediasyncer - gossip network based filesyncing
** PERS
https://github.com/libercv/peerbackup
* **Peerfliex**
https://github.com/marten-seemann/quic-conn - fancy replacement for tcp

https://github.com/mafintosh/torrent-stream
https://github.com/netCommonsEU/PeerStreamer-peerviewer - go + peerstream

https://github.com/matiasinsaurralde/ng-chunked-audio	
https://github.com/ipkg/difuse
https://github.com/libp2p/go-libp2p-swarm
https://github.com/tmthrgd/apt-p2p

* **Torrent Sync**
https://github.com/syncthing/syncthing

https://github.com/NebulousLabs/merkletree

* **DHT**
https://github.com/nathanpotter/go-chord	
* **Highlevel networking**


https://github.com/aniket-gore/distributed-peer-to-peer-system

https://github.com/evanphx/mesh

https://github.com/matiasinsaurralde/loudp2p - music streaming example
https://github.com/libp2p/go-libp2p-swarm
https://github.com/rsms/gotalk
https://github.com/jbenet/go-peerstream
* **API Examples**

* **Network Block Device**
  * https://github.com/demonicblue/netdriver - Userspace driver in Linux utilizing NBD (Network Block Device).
  

* **DHT**
  * https://github.com/fastfn/dendrite - DHT chord protocol with vnodes for testing

* **Gossip**
  * https://github.com/hashicorp/memberlist
# HTTPFS
*https://github.com/elegios/httpfs

# Automation / Provisoning

* https://github.com/ZoidbergConspiracy/telepath - uses teleport to manage servers in a cluster
# XMPP

* https://github.com/mattn/go-xmpp

* https://github.com/agl/xmpp-client - pure go

# processes
https://github.com/jandre/procfs
# hardware virt clear containers
* https://github.com/containers/virtcontainers
https://github.com/clearcontainers/agent
# WEb framework
* https://github.com/primefour/xserver - wow lots done
# Portable Web APP

* https://github.com/UnnoTed/fileb0x

# Distributed FS

* https://github.com/RobinUS2/xyzfs  -looks good, gossip, shards, replica sets, GETPUT, file represattion
# WEb App

* https://github.com/rs/xid - uses different factors to make a unique ID
# Hash

* https://github.com/pierrec/xxHash pure go

# Compression
* https://github.com/ulikunitz/xz

# TCP

* https://github.com/xfxdev/xtcp

# CrptFS
https://github.com/rfjakob/gocryptfs
# Secure tokens, fuck jwt
https://github.com/nogoegst/token 
# Compression
* https://github.com/ulikunitz/xz


# Secure tokens, fuck jwt
https://github.com/nogoegst/token 

# Diffie Helmans
* https://github.com/riobard/go-x25519

* example with diffie helmans doing vpn https://github.com/RoPe93/govpn

# websockets
* https://github.com/zreigz/ws-vpn - websocket vpn
# weboscket p2p
https://github.com/noxyal/websocket-p2p
# Balancer syncer
https://github.com/zeisss/mediasyncer

# Ring sync
https://github.com/ZiroKyl/RingSync


# Users
_Identicon_
https://github.com/nowshad-sust/awesome-identicon
__CONSISTENT RING__

# Thead safe redis consistent hash pub sub
https://github.com/mixer/redutil


# Consistent Ring Node Mnaager
https://github.com/go-trellis/node
https://github.com/billhathaway/consistentHash
# Consistent Ring P2p
https://github.com/wayne666/consistent-hash
https://github.com/Nitro/ringman
This is a consistent hash ring implementation backed by our fork of Hashicorp's Memberlist library and the hashring library.

# Jump Algo COnsistent Ring
https://github.com/renstrom/go-jump-consistent-hash
https://github.com/chenjm1217/jump_hash
__DATABASE__

# Doozer
https://github.com/ha/doozerd - real time updates infrquently updated data 
https://github.com/ha/doozer - pretty sweet, can literally store full on files with their own /root/hierarchy thing

These processes communicate with each other using a standard fully-consistent distributed consensus algorithm.

# Consistent DB
https://github.com/abcum/rixxdb
    In-memory database
    Built-in encryption
    Built-in compression
    Built-in item versioning
    Multi-version concurrency control
    Rich transaction support with rollbacks
    Multiple concurrent readers without locking
    Atomicity, Consistency and Isolation from ACID
    Durable, configurable append-only file format for data persistence
    Flexible iteration of data; ascending, descending, ranges, and hierarchical ranges
