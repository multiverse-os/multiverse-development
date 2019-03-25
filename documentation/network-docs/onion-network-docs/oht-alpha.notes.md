# Multiverse OS And Onion Networking

https://github.com/smartystreets/mafsa - tree for fuzzy lookups and such


[secp256k1-zkp](https://github.com/ElementsProject/secp256k1-zkp)
xperimental fork of libsecp256k1 with support for pedersen commitments and range proofs. 
**Features:**
		* secp256k1 ECDSA signing/verification and key generation.
		* Adding/multiplying private/public keys.
		* Serialization/parsing of private keys, public keys, signatures.
		* Constant time, constant memory access signing and pubkey generation.
		* Derandomized DSA (via RFC6979 or with a caller provided function.)
		* Very efficient implementation.
## Reed solomon
https://github.com/sec51/gf256
## uinput
https://github.com/ynsta/uinput
## tor web search engine
https://github.com/justyntemme/MooseChat
https://github.com/nogoegst/tulp
https://github.com/nogoegst/onionsink
https://github.com/nogoegst/whatonion
https://github.com/nogoegst/torsh
## DB 
https://github.com/tidwall/buntdb


https://github.com/1lann/cete kvdb badger
https://github.com/bmeg/arachne badger graph
https://github.com/bsm/raft-badger
https://github.com/nak3/badger-cli
[UI]

dialog boxes that match gnome, look very nice
https://github.com/gen2brain/dlgs


https://github.com/donomii/hashare 
This CAS filesystem merges duplicate files, can keep thousands of snapshots with tiny overhead, and can be distributed, backed up to cloud storage and works well with network sync services like dropbox.
## GitQL
https://github.com/cloudson/gitql
## Redis 
https://github.com/albrow/jobs
https://github.com/wgliang/logcool
https://github.com/xyproto/permissions2

Multiverse OS is built around onion networking, it is an idea that started with the Whonix design of using dedicated virtual machine routers to route virtual machine networking to ensure complete isolation.

The *onion hash table* (**OHT**) is a foundational piece of software in the Multiverse OS ecosystem. As research and development continued on both projects, it became clear that **OHT** would actually need to be a collection of software working together to acheive the functionality required by the **OHT** in order to implement it in a modular way that individually could find their use in a variety of related and unrelated projects.

The **OHT** is important because it is the foundation to *Decentralized Onion Applications* **(DOA)**. The first **DOA** provided as an example/demo application is a decentralized chat server.

https://github.com/whyrusleeping/go-multiplex
https://github.com/jbenet/go-peerstream
[Example Code]
https://github.com/nogoegst/avant
https://github.com/miolini/metasocks
https://github.com/goshinobi/tor_multi/blob/master/tor.go
https://github.com/libp2p/go-libp2p-net
#### Mail
https://github.com/amalfra/maildir
#### Multaddress
https://github.com/multiformats/go-multiaddr
#### Trorrent tracker
https://github.com/chihaya/chihaya
#### Virtcontainer 
**virtc**
virtc is a simple command-line tool that serves to demonstrate typical usage of the virtcontainers API. __This is example software; unlike other projects like runc, runv, or rkt, virtcontainers is not a full container runtime.__

**Get your image**
Retrieve a recent Clear Containers image to make sure it contains a recent version of hyperstart agent. To download and install the latest image:
			latest_version=$(curl -sL https://download.clearlinux.org/latest)
			curl -LO "https://download.clearlinux.org/current/clear-${latest_version}-containers.img.xz"
			unxz clear-${latest_version}-containers.img.xz
			sudo mkdir -p /usr/share/clear-containers/
			sudo install --owner root --group root --mode 0644 clear-${latest_version}-containers.img /usr/share/clear-containers/
			sudo ln -fs /usr/share/clear-containers/clear-${latest_version}-containers.img /usr/share/clear-containers/clear-containers.img

**Get virtc**
Download virtcontainers project
			go get github.com/containers/virtcontainers

Build and setup your environment
			cd $GOPATH/src/github.com/containers/virtcontainers
			go build -o virtc hack/virtc/main.go
			sudo -E bash ./utils/virtcontainers-setup.sh

virtcontainers-setup.sh setup your environment performing different tasks. Particularly, it creates a busybox bundle, and it creates CNI configuration files needed to run `virtc` with CNI plugins.

**Get cc-proxy (optional)**
If you plan to start `virtc` with the hyperstart agent, you will have to use `cc-proxy` as a proxy, meaning you have to perform extra steps to setup your environment.
			go get github.com/clearcontainers/proxy
			cd $GOPATH/src/github.com/clearcontainers/proxy
			make
			sudo make install

If you want to see the traces from the proxy when `virtc` will run, you can manually start it with appropriate debug level:
			sudo /usr/libexec/clear-containers/cc-proxy -version 3

This will generate output similar to the following:
			I0410 08:58:49.058881    5384 proxy.go:521] listening on /var/run/clearcontainers/proxy.sock
			I0410 08:58:49.059044    5384 proxy.go:566] proxy started

The proxy socket specified in the example log output has to be used as `virtc's --proxy-url` option.

**Build cc-shim (optional)**


#### Containers

[cgroups]
https://github.com/containerd/cgroups
https://github.com/opencontainers/runtime-tools/blob/master/cgroups/cgroups_v1.go
[proxy]
https://github.com/clearcontainers/proxy
[init]
https://github.com/clearcontainers/hyperstart
[storage]
https://github.com/containers/storage
https://github.com/containerd/continuity A transport-agnostic, filesystem metadata manifest system
[network]
https://github.com/containernetworking/cni
https://github.com/containernetworking/plugins
[cli]

https://github.com/containerd/go-runc
[runtimes]
https://github.com/containerd/containerd
https://github.com/containerd/cri-containerd
https://github.com/containers/virtcontainers
https://github.com/clearcontainers/runtime
[image]
https://github.com/containers/image

https://github.com/containers/oci-fetch - image fetching
[build]
https://github.com/containers/build
https://github.com/clearcontainers/osbuilder

[agent]
https://github.com/clearcontainers/agent
[extra]
https://github.com/containerd/console
https://github.com/containerd/fifo
#### Log
[zerolog](https://github.com/rs/zerolog)
zero-allocation JSON logger, JSON, timestamps
Output: {"level":"info","time":1494567715,"foo":"bar","n":123,"message":"hello world"}

#### Currency
[big decimal:math](https://github.com/golang-plus/math/tree/master/big)

#### I18n

#### Go-qrcode
[go-qrcode](https://github.com/skip2/go-qrcode)
has reedsolomon and dense data packing
#### GPG
[gpg-agent](https://github.com/prep/gpg)
This is an experimental repository of a client to the GPG agent. It was built out of a desire to have a somewhat friendly interface to GPG keys stored on a smart card by way of GPG.
#### Socket Pair / Socket / Unix Socket / AF_*

**socketpair**
[socketpair](https://github.com/hackwave/socketpair)


#### Design Patterns

**assortment of container queues (fifo, lru, ..)**
[concurrent](https://github.com/golang-plus/caching)

[event-store](https://github.com/Ladicle/event-store)
http accessible fifo store

**pipe**
[pipein](https://github.com/chrhlnd/pipein)
Makes it easy to take input on pipe, and shut it down.

**queue/fifo**
[funnelqueue](https://github.com/ssgreg/funnelqueue)
FIFO, lockfree, multiproducer, single-consumer, link-list-based queue

**emitter**
https://github.com/abcum/emitr
#### Filesystem
[go-fuse](https://github.com/hanwen/go-fuse)
[gitfs](https://github.com/hanwen/gitfs)
#### Containers
[microcontroller](https://github.com/oracle/smith)
#### Common/Shared Components & Libraries
As much as possible components should be shared and reused, by making the components modular and support reuse they will hopefully find wider support and adoption by being incorporated in more projects than just Multiverse OS.

[*] **Localization**

[*] **K/V Store**

_Support for the variety of tree types: ()_
The shared key/value store requires a specific set of features:

  > _Multiple Collections/Buckets & Multidimensional collections_
 
  > _Storage Encryption and Session Encryption_

  > _Support for local-only & distributed support_

  > _Large values are chunked and evenly distributed_
    > _Consistent hashing_ splitting the hashed content evenly, enabling easy sharding of data
    > _File storage by chunking and evenly distributing files using consistent hashing_
    > Support for streaming out of the database and into the database
    > _Indexing support with full-text search_
 
  > _Versioning support_

  > _Hierarchial support_ (removing a node in the tree will remove all child branches that are rooted in the deleted node)

  > _Transactional_ ACID compliant, rollback support, one writer with multiple readers

  > _Basic tree algorithms: Nearest Neighbor Algorithms_

  > Geospatial data support for real or fake locations with basic algorithms

  > _Support for Currency Datatypes_
  > _Graph DB overlay_
  > _Caching DB overlay_
  > _Multiple compression options_ **(Snappy, Lz, and ...)**
  > _Hooks_ before_save, after_save, before_delete, after_delete, ...
  > _Permissioned_ requiring m-of-n signatures depending on collection ruleset
    > _Collection rulesets_ to define access rules, and roles
  > _In-memory_ or _File-backed_ (if file-backed, perhaps filesystem backed)
  > _Storage in binary_ and retrival available in a variety of common econding formats (JSON)

#### OHT Core Components
Three core components, with many sub-components, make up the OHT system. Additionally, the *scramble suit identity protocol* (**SSIP**) is a core component of the Multiverse operating system.

[*] __Onion Network Interface__ *(ONI/ORI)* (or Onion Router Interface)
  > Provide /proc/net/{DEVICE} (possibly /proc/net/onion)
    [Sample Code]
    > https://github.com/elsonwu/GOnetstat
  > Provide /dev/net/{DEVICE} (possibly /dev/net/onet0)
  > Provide tools to list all available onion network address endpoints (formely called hidden service endpoints) 
  > Provide routing UI (for example, to route all network traffic through the onion device, likely using iptables, or to route onion incomming traffic to specific local or remote ip:ports) [Multiverse router? Perhaps this should be its own component, onion router controller (ORC)]

[*] Onion protocol suite
  > *Link Layer* - Onion Network Interface Tunnel
    > *Internet Layer* - *IPv4* (Tor only supports IPv4)
      > *Transport Layer* - *TCP* (Tor only supports TCP) - UDP (Can be tunneled over TCP or perhaps multiple TCP connections) - SCTP (Interesting but not supported) 
      	> *Application Layer* (**Indicates Support Planned**, *Support Needs Research*, Unsupported)
        	(>) **HTTP**
        	(>) **Websockets**
        	(>) SPDY
        	(>) **SSH (SCP)**
        	(>) NTP (Network Time Protocol, current protversion 4; NTPv4)
        	(>) SMTP (Simple Mail Transfer Protocol)
          (>) Network News Transfer Protocol (NNTP / Usenet)
        	(>) DNS
        	(>) DHCP
        	(>) *XMPP* (Extensible Messaging and Presence Protocol)
          (>) *Matrix* (
        	(>) *SIP* (Session Initiation Protocol) for VOIP and associated text-messaging
              SDP (Session Description Protocol) is an associated protocol 
        	(>) FTP
        	(>) SMTP
          (>) Post office Protocol (POP)
        	(>) RIP (Routing Information Protocol)
        	(>) RTP (Real-time Transport Protocol) for delivering audio and video
              RTCP is an associated protocol for statistics and monitoring
        	(>) WebRTC *(Possible with modified client, experimented with this previously)* -
        	(>) 

[*] Scramble Routing (Using avaialble endpoints, support high bandwidth protocols using scramble hops *(perhaps with just 1 hop of Tor?)*
  [*] Additional Hop Layers *(Using additional endpoints available, configure 1 or more additional hops before entering or after exiting the onion network)* [Able to function as an onion network without Tor? But heavily relies on Tor still because they are solid and well tested?]

[*] __Scramble Suit Identity Protocol__ *(SSIP)* [Core part of Multiverse OS] *(Should camilstore be used?)*
  > Provide Ephemeral Key Hierachical Tree
  > Identities (vCard for each) *Scramble Suit is built on the concept that people maintain more than one psuedonym, therefore it is designed to help segregate identities, manging associated accounts (email adresses, social media accounts, ...), passwords, cryptographic keys, documents, images, ...)  
    [*] _Cryptographic keys_
      > PGP Keys, SSH Keys, Cryptocurrency Keys
    [*] _Password store_ + Account Details + OTP 
    [*] _Projects_ (Projects range from source code (programming), audio production (podcast, music production), video production, ...) *(Each project has its own Virtual File System that allows all associated files to be mounted, saved and archived together)* *(Projects CAN HAVE MULTIPLE OWNERS)*
      > Ruleset is defined upon creation defining who can create, edit, view, delete files, within the project
      > Virtual File System Metadata file *(Includes a list of all files, structure of the files, types of files, public metadata associated with the files for searching)*
      > Changes to filesystem are signed by owners
    [*] Dictionaries (Key/Value Buckets) (notes, snippets, vCards, hosts, ...)
      > Contacts vCard(s)
        > Contact vCard Grouping (friends, family, work, etc), forms basic channels to broadcast
    [*] Media Libraries *(Media each has their own filesystem built to the needs of the file types it contains)*
      > Documents
      > Images
      > Videos
      > Books
      > Comics
		  > Calendar
		    > Cache of group calendars (compiled from public calendar data of members in grouping), organizational calendar (compiled from public calendar data of members in organization)
      > Music
      > Versioned Source Code (Git Repositories, Mercurial, etc)
      > Email
    [*] (Social) Compiled together makes a profile page
      > (Social) Status Updates
      > (Social) Messages (text and/or media) *(Images link from documents, and is not unnecessary repeated)*
      > (Social) Articles *(Can be shared, ruleset defined on creation)*
    	> (Social) Geospatial Message (Available to people looking up nearby radius, can be encrypted to people in that radius at that time to ensure it is only available to the people 'who were there' because you had to be there)
      > (Social) Comments, "Likes" or the better equivilent which are essentially a form of comment

    [*] Computing Resources / Personal Cluster / Cloud Computer *(Real `cloud` not bullshit corporate cloud)
      > Routing Table
        > Available Endpoints (Domain names, onion addresses, WAN IP Addresses) (Can be shared with others and they can route through them)
      > Cluster Infrastructure (Details (name, server class, etc), Configuration, Provisioning Recipies, Overlay Virtual filesystem) & Networking [Is Active?] *(Can be shared, ruleset defined on creation)*
        > Router
        > Server
   
  > _Multiverse K/V Database_ (bubbleverse or blackhole) an encrypted Key/Value Database, items needed and decrypted and encrypted to a session key *(allows items to stay encrypted in memory and usage/access to master key limited to one decryption per session key)*
    > Cache For Actived Decentralized WebUI(s)


[*] __Wormhole Messaging__ (?)
    *A fault tolerant (able to receive messages offline), multi-endpoint (supports mesh networking and pathing of packet subcomponents over multiple connections), and ...*


[*] __Onion Hash Table__ *(OHT)* 
  [*] *(Wormhole network?)* Onion Hash Messaging Queue [Interchangable]  
    * Very basic implementation that can easily be expanded upon using modules/plugins or switching it out completely
    * Key/Value Store
      * Default
  [*] Key / Value Tree Database [Interchangeable]
        > In memory storage, that actively purges based on TTL, and securely wipes memory, a complete wipe is done upon closing
        > Multi dimensional, versioned binary tree
    * Very basic implementation that can easily be expanded upon using modules/plugins or switching it out completely
    * [?]

[*] __Decentralized Onion-Routed Application Framework__ or _Decentralized Onion-routed Application_ *(DOA)*
  > Special K/V field provided to describe *DOA* 
  > Store server data inside OHT database that is cached locally  
	[Examples] *(Code launched and stored in an onion address by peers participating, more peers = more duplicaiton, if it gets large, peers can shard data, onion address but can be routed by peer based DNS to easy to remember name)*
    > Decentralized Forums
    > Decentralized Link Aggregation (Reddit)
    > Journal/Magazine with collectively written articles
    > Marketplace
    > Git Repositories
====
#### <IMPORTANT> INSTALL MULTIVERSE
[cpuid](https://github.com/klauspost/cpuid)
Package cpuid provides information about the CPU running the current program.

[shutdownv2](https://github.com/klauspost/shutdown2)


#### Users
**passwords**
[dictionary password check](https://github.com/klauspost/password)

#### Assembly / Reverse Engineering
[asmfmt](https://github.com/klauspost/asmfmt)



#### Data Models
[ffjson](https://github.com/pquerna/ffjson)
drop in replacement that is much faster than stdlib

## MAC Address Lookup
[oui](https://github.com/klauspost/oui)

#### Stats
[g2s](https://github.com/peterbourgon/g2s) forward simple stats to statsd server

#### Centrifugo
[real time message](https://github.com/centrifugal/centrifugo)
websockets or sockjs server, webrtc
[web interface](https://github.com/centrifugal/web)
[go http client](https://github.com/centrifugal/gocent)
[go websockets client](https://github.com/centrifugal/centrifuge-go)
[js websockets client](https://github.com/centrifugal/centrifuge-js)

#### Standard Library
[syscall](https://golang.org/pkg/syscall/)
The large list of available syscalls in Golang from Reset, to Setenv, etc

#### ThreeJS IN TERMINAL!
https://github.com/zz85/threejs-term

#### Deduplication
[dedup](https://github.com/klauspost/dedup)
#### Compression
[snappy](https://github.com/golang/snappy)
**assortment**
[compress](https://github.com/klauspost/compress)
snappy, gzip, fse, zip, zlib, flate
#### Reed-Solomon
[reedsolomon](https://github.com/klauspost/reedsolomon)
#### Steganography
**images**
[jpeg]
https://github.com/lukechampine/jsteg
https://github.com/jbochi/strogonoff
https://github.com/henkman/outguess
[png]
https://github.com/evantbyrne/pngsecret
https://github.com/hypermassiveblackhole/steganography

[bitmap]
**network delay**
https://github.com/alex-kostirin/chanstego

#### Testing
**BDD**
[ginkgo](https://github.com/onsi/ginkgo)
BDD Testing Framework for Go
#### (gRPC + Protocol Buffers) Examples
[go-chat](https://github.com/taylorflatt/go-chat)
gRPC + protocol buffers based chat
[remote-shell](https://github.com/taylorflatt/remote-shell)
gRPC + protocol buffers based remote shell
#### (DHT) Chord lookup algorithm
[go-chord](https://github.com/taylorflatt/go-chord-implementation)
#### Relevant Libraries & Sample Code
**Networking**
[goplane](https://github.com/osrg/goplane)
[netlink](https://github.com/vishvananda/netlink)
[ntp server](https://github.com/beevik/ntp)
#### Notes / Research
* Binding to :0 will ask the kernel for an available port within the ephemeral port range
the kernel will assign it a free port number somewhere above 1024.

	// tell pinger that it is privileged.
	// NOTE: You must run `setcap cap_net_raw=+ep pocket-loss-monitor`
# Low Level NEtworking
[!][Fake Interfaces, pure logic network devices entirely in userspace. That means you don't need root, yeah, really.]

// BaseClient is a streaming telemetry client with minimal footprint. The
// caller must call Subscribe to perform the actual query. BaseClient stores no
// state. All updates must be handled by the provided handlers inside of
// Query.

*Example:*
https://github.com/openconfig/gnmi/blob/master/client/gnmi/client_test.go
*In this example, we will see a client/connection getting initalized, then udpates getting sent down it.*

[!][Coalesce & Queue]
[Queue /w Coalesce][https://github.com/openconfig/gnmi/blob/master/coalesce/coalesce.go]

````
wget https://raw.githubusercontent.com/openconfig/gnmi/master/coalesce/coalesce.go
````

**Ruby**
[go-mruby](https://github.com/mitchellh/go-mruby)
#### Globbing & Path/String Matching
[doublestar](https://github.com/bmatcuk/doublestar)
#### Example Projects Built With OHT

