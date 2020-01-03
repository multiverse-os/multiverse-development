# Multiverse OS: Router Controller 
The most recent design for the network controller is one that takes in a 
physical passed through network.

After trying several options, Debian9 base system with two (2) nested Alpine
Linux based routers.

  (1) Galaxy Router 

  (2) Star Router

In the previous designs we had three (3) Alpine Linux routers, but we are 
making the encapsulating Debian9 system Universe which holdes the galaxy and
star router, which makes the metaphor work better. 

_______________________________________________________________________________
## Sub-VM Communicaiton & Control
We need to get access to data in the VMs and control them immediately. This
can be built upon and improved upon but the basics are important now. 

WE need unix socket communicaiton, maybe even HTTP over socket, just JSON REST
API, in addition, regular console output so we can control the systems without
SSH or any graphics needed. THis will greatly reduce the resources used for
all the VMs.


This will be the early version of the system agent, the most generic type that
can be the basis for several types. so we will also want system profile


We should be feeding time up to VMs on every heartbeat singla


## base libs
bytearena/schnapps - assortment of kvm/qmp tools



[hash tables]
burntsushi/wendy

arriqaaq/chord
armon/go-chord
chuckha/dht - chord
fastfn/dendrite - chord
r-medina/gmaj - chord
nallg00d/chord-go
shariqazz/customchord
itsvamshiks/chord
ramusrygaard/chord
xirenhua/chord
phanitejakesha/chord
wang502/chord
soumyadeep2007/chord

tilakneha/tripletstorechord

taylorflatt/go-chord-implementation

lukaspj/go-chat - using chord

sam1323123/go_dht


wheelcomplex/gauss

zgreat/chordchain

cbocovic/chordfs





cznic/kv
lunny/nodb - kv,list,hash,zset,bitmap,set
portworx/kvdb
go-openapi/kvstore
tidwall/kvnode - redisapi leveldb storage raft support
zegl/goriak/ riak kv
cpuguy83/kvfs - kv fuse fs
chrislee87/kvstore - time series indexingt data compression
icza/kvcache 

yourbasic/graph
google/badwolf
gonum/gonum


docker/libkv
tidwall/buntdb

golang/leveldb
syndtr/goleveldb
beeker1121/goque

houzuoguo/tiedot - nosql

asdine/storm

cetrifugal/cetrifugo


guelfey/go.dbus



emitter-oi/emitter
olebedev/emitter
chuckpreslar/emission
cloudfoundry/route-emitter
desertbit/event
abcum/emitr
segmentio/events


antsmartin/kqueue-event-loop
asaskevich/eventbus
tidwall/evio -f ast event loop enetworking
docker/go-events



kljensen/golang-html5-sse-example




subchen/frep - json/toml/yaml config data and templates
golang-samples/template
tmrts/boilr
kvannotten/ricer
heavyhorst/remco
hairyhenderson/gomplate
klingnet/tampl
casba/tplm
zert/bootsmann


[shell interafces]
jroimartin/gocui - console userisnterfaces
abiosoft/ishell 

hackwave/textui

hackwave/go-os (has xdg lib inside and meant to stick other os stuff in)


markelog/list


posener/complete
firba1/complete



vtolstov/go-qcow
vasi/qcow2
zchee/go-qcow2

cep21/circuit - circuit breaker battern


sogko/slumber


go-cmd/cmd


systemboot/systemboot

[cuistom init]
driusan/dainit
miekg/dinit
mustafaakin/gonit
robinmonjo/dock
bitrise-core/bitrise-init
ddollar/init
kubic-project/kubic-init
fanus/kubeinit
aerokube/init
pablo-ruth/go-init
tormentaos/init
convox/init convox-archive/init
yonkornilov/cloudless-init
zero-os/cloud-init-server
ramr/go-reaper
kevinschoon/gaffer - runc init
lonord/docker-init
estintax/pinit
tehmoon/docker-init
enonck/go-init
goatos/dinit


plasma-umass-systemgo - near compelete systemd rewrite in go


digitalrebar/provision dhcp/pxe/tfp
coreos/matchbox network boot
  kelseyhightower/cpic - coreos pxe image customizer
ggiamarchi/pxe-pilot - cli manage pxe
dimensiondataresearch/mcp2-dhcp-server
pin/tftp
google/netboot

ortss/engine - generic transport agnstoic strema layer


longsleep/entropyd


jchv/eventgun - forward file system events across sockets


hackwave/scramblekey-go

[console support terminal etc]
golang/term
xcat2/goconserver

subchen/go-tableify
motemen/go-colorine
vitaliy-svinchyak/cml - console markup language


coreos/systemd
coreos/systemd-rest
julienschmidt/systemd

0xef53/qmp-shell
0xef53/go-qmp
0xef53/phoenix-guest-agent

zcalusic/sysinfo
elastic/go-sysinfo
ungerik/go-sysfs
alexreptu/sysinfo
clarketm/systemstat

jacobsa/cputime

coreos/motd-http
victorgama/howe
nemith/motd


puma/puma-dev


google/gopacket
google/stenographer
songgao/packets
aerogo/packet
justinazoff/gotm - flow cutoff rotation and compression of packets
ghedo/go.pkt - cpaturing,injecting,filtering,encoding,decoding
kung-foo/freki - maniupualte packets in usermod with nfqueue
goiiot/puppy -packlet capture with nocgo and supportinb bpf
go-freebsd/pf - packet filter
untangle/packetd - userspace packet processing
cuishark/cuishark


packet
kr/pty
containerd/console pty
google/goterm
tianon/debian-golang-pty

tarm/serial
jacobsa/go-serial **
bugst/go-serial
goburrow/serial
dustin/go-rs232
npat-efault/serial - async
ecc1/serial
ncar/agnoio - transport agnostic io (ethernet serial, etc)
jlubawy/go-slip - codec for serial line internet protocl (slip)
wybiral/hookah - swiss army knife data pipes - http serial tcp pipeline

hedhyw/go-serial-detector - determining active serial ports - sysfs linux serial

jason-mitchell/serial4go - abstraction serial layer
vaelen/vcom - small basic serial terminal


spatialcurrent/go-simple-serializer - BSON <-> JSON <-> YAML <-> TOML, etc

[time series db]
eliothedeman/serial


influxdata/go-syslog - syslog parser (super fast)
ziutek/syslog - syslog server skeleton

mholt/archiver - zip/tar/tar.gz/tar.bz2/tar.xz/tar.lz4/rar,etc
golang/snappy
pierrec/lz4
klauspost/pgzip (parallel gzip)
klauspost/compress - optimized compression


clearcontainers/runtime
kata-containers/runtime
zhujintao/kyum - install kvm from filesystem
jsimonetti/go-spice - can get our viewer over html5, this would
help get rid of virt-manager while keeping our SPICE access


rafaelmartins/simplevirt - VERY NICE show alia
kitschysynq/darity - show alia, very nice, also explains one of the fd, its the kvm
fd

[genmeral userful libs]

dropbox/godropbox


[agent]
kata-containers/agent
clearcontainers/agent

[service/daemon]
hlandau/service - HAS ROOT DROP PRIVILEDGES!

takama/daemon - cool because even has a cronjob like example that runs in
the backgorund, and does something on a timer, cna even use cron timing * * * etc


emicklei/go-restful - sweet ws/rest service





mpolden/echoip


[http service]
go-chi/chi - WOW, much better than echo, gin, etc


[distributed linux services]
reconquest/shadowd - suepr lightweight distirubted shadow files


[ruby]
goruby/goruby
goby-lang/goby - ruby LIKE lang

[mruby]
mitchellh/go-mruby

taka7646/mruby-embed

watermint/grb - go and mruby binding

[ruby scriptiung]
box-builder/box - docker building with ruby
udzura/dsl-with-go-mruby


[netstack]
google/netstack
dutchcoders/netstack
fuchsia-mirror/third_party-netstack

sheepbao/gotcpip
clmul/netstack

[raw sockets]
ccsexyz/kcp-go-raw
kdar/gorawtcpsyn

crznic/crznic

yudaishimanaka/rawdump

naoyamaguchi/rawsocket_practice

ryabuhin/golang_linux_rawsockets


razc411/gobd


kyleconroy/coiltap - capture http traffic from raw


fharding/rsh - reverse shell using raw tcp scoekts


farkaskid/chatserver - using raw tcp sockets

izanbf1803/http-go-server - http server using raw sockets :DDDD


[kcp]
xtaci/kcptun
fast and secure tunnel based on KCP with N:M multiplexing

xtaci-kcp-go
reliable udp

vzex/dog-tunnel - p2p tunnel using kcp


wxiaoguang/kcp-conn
stream udp 

kdada/tinyvpn - using kcp

mitm using kcp
yinqiwen/gsnova

ginuerzh/gost - kcp tunnel

jjzhang166/kcpraw

paralin/go-libp2p-kcp

skyglance/p2p-live-client

skyglance/kcp-tunnel

yves-yuan/net_kcp



lengzhao/libp2p - with kcp


[proxy]

snail007/goproxy - high preformance https, websockets, tcp udp dns, socks5 proxy server


[trie]
derekparker/trie - fuzzy fast fprefix
vonng/ac - double array trie


tchap/go-patricia - prefix radix tree

adamzy/cedar-go - double asrray tree with good updating


kentik/patricia - IP/CIDR specialized trie

justasitsounds/trieconfig - data graphs in tree

cheesedosa/trie - autocomplete


autprojects/flashtrie.go - ip specialized trie


[bloomfilters]

dgryski/go-bloomindex

[full text search]

neowaylabs/neosearch

bradleypeabody/fulltext

blainsmith/goreds - redis based full text


mosuka/blast - full text indexing with bleve
bishudark/bleve-api - rest api for bleve


[linux server search]

marshyski/plural -very cool

[dht searhc]
felix/dhtsearch

[search engine]

vosmtrek/violet- lightweight search engine

vsouza/google-search - search engine named google in golang


dwayhs/go-search-engine

gosearch/gosearch

c-data/gosearch

dgryski/go-postings - supports compressin

mateusduboli-searchzin - newest 


kunbetter/grid-search

[search]

google/codesearch
google/zoekt - code search using trigaram


lithammer/fuzzysearch - fast tiny fuzzy

sajari/fuzzy with suggsetion

petar/gollrb - left leaning red black binary search tree





averagesecurityguy/searchscan
search nmap and metasploit scanning

[syslog][
ekanite/ekanite - syslog server with built in serach




[routers]
github.com/rtr7/router7

[OS]
fuschia-mirror


[SINGLE BINARY CONTAINERS]

genuinetools/binctr

rootless-containers/runrootless


gustavosbarreto/go-microconainter - ocol


[cli ui]
jroimartin/gocui - tui

[serial term emu]
ishuah/bifrost
[serial]
jacobsa/go-serial

[binary analysis]
grantseltzer/prism

[c transpiler]
elliotchance/c2go

konstantin8105/c4go



[qemu]
intel/qemu-lite


[gopherjs ui]

gopherjs/vecty - frontend toolkit

norunners/vue

huckridgesw/hvue


lngramos/three - threejs gopherjs


gmlewis/go-threejs 

dave/dropper -d rag and drop


dave/console


agamigo/material - with gopherjs


nobonobo/vecty-chatapp
gernest/cute - matierla ui wiuth gopherjs and vecty


cnguy/gopherjs-frappe-charts - simpel charts for frontend, including github heatmap


gui-io/gu - 

lifeng1335/gopherjs-vue-examples

[webrtc]
nobonobo/p2pfw



[motd]

nemith/motd - utilities to display upon login


[time using systemd]

trstringer/go-systemd-time


func Booted() bool {
  fi, err := os.Lstat("/run/systemd/system")
  return err == nil && fi.IsDir()
}


_______________________________________________________________________________
## General

[CLI]
[^][TUI]
sparkymat/grip - holy shit, produeces amazing looking tuis, almost regular gui
using CSS style formating

[^][^][Chat examples - useful for cluster UI too]
koolay/console-chat

[^][Command-line tools]

[^][Autocomplete]
c-bata/go-prompt - drop down auto complete

posener/complete - generate zsh/bash/other autocomplete files


[GUI]

[^][WebUI]

[^][^][SSE]
prasannavl/estream

[^][X Server]

[^][^][Xgb Utils]
burntsushi/xgbutil - by far the most complete version of this

[^][Wayland]



[FS & Memory]


mohanson/acdb - go objects in both fs and mem

[os]

[^][singls]
mohanson/exit - super nice and simple exit but no control over singals



[bash scripting in go!]
tcnksm/bash-init

_______________________________________________________________________________
## OHT

**tylertreat/bloomfilters**

hkparker/tlb


google/gopacket

hkparker/imux

proxyd - proxy between tcp tls and unix socket

pure go-i2p router

[Hash Table Examples]

[^][Pastry]
burntsushi/wendy - some cool examples

[^][Kademlia (bittorrent]
prettymuchbryce/kademlia


[message protocols]


[^][nats]
nats-io/go-nats-streaming - cool example, cool distributed msg system for clsuters



[^][es: easy session/stream protocol]
ooclab/es - supports multipathing, p2p, udp, tcp, tunnel, inverse, reverse, tec

[networking - low levle]
[^][knx protocol] 
vapourismo/knx-go - provides groupgateway, grouprouter, bridge, etc


[!!][Ideas]

 * Combine kcp with tidwall/evio (event loop networking)
   to provide async connections between oht members
_______________________________________________________________________________
## Scramble Suite
[merkle]
google/trillian - maybe bloated


chasestarr/merkle - awesome, small and does the job!



[luks]

[gpg]


[ssh]


[tree/mycelia data structures for key heiarchy]

lleo/go-hamt-functional - hash array mapped trie - 

[mega simple vm for bitcoin to handle multisgs]
gravity-exchange/vm


github.com/patrickmn/go-cache - very nice and super basic cache kv


[PRIVATE TMP]
0xef53/libpam-privtmp

_______________________________________________________________________________
## VM Interesting Features

````
   <keywrap>
     <cipher name='aes' state='off' />
   </keywrap>

Determines if guest can preform the S390 key management operations a clear key can be proted by encrypting it under a unqiue wrapper key that is generated for each guest VM running on the host.
````


_______________________________________________________________________________
## Building Linux Kernel and PCI devices without C deps


[novm]
this project has a lot of interesting code evne though its abandoned like `mmio` device code

  also device.go appears to have everything needed to make a device in pure go
  including driver  etc seriously good looking

  `pci.go` is seirously good looking, whcih is def code for making virtual PCI
   because it goes with `pcihost.go` which is the host bridge that it is attached to

   `virtio.go` deps but we should be able to correct that or at least use it in combo with other virtio go libs

    in this code there is virtioDevice which includes PCI device
   `virtio_block.go` has virito block deivces *which includes virtioMMIOblock device which IU have not seen. and viritoPCIblock

    `virtio_console.go` looks very interseting too

    `virtio_net.go` doesnt have any obvious deps in c. but seriously aweomse. this is how we can really get some intersesting shit going

   includes `virtio_stream.go` which looks good



     -- look at facebookincubator/oculus-go-kernel drivers/vhost for some interesting files

_______________________________________________________________________________
## VM Shares [ 9p vs nfs ]

 * NFS or p9, should be over vsock! nfs can be faster in many benchmarks
   but go versions of each can increase the speed by concurrency++


  * vsock doesnt need network interface and doesnt setup net connect

  * it may subsitute virtserial


  Use NFS with vsock for sharing foldedrs across deviecs!


hyperhq/hyperstart - c init initramfs but it has VSOCK!

_______________________________________________________________________________
## VM Communication Methods

[virtserial]
slow old not ideal, only really 1:1 beause 1:N sucks

**virt channels are a TYPE of virt serial!**

[memory modules]

````
<devices>
  <memory model='dimm' access='private' discard='yes'>
    <target>
      <size unit='KiB'>524287</size>
      <node>0</node>
    </target>
  </memory>
</devices>


<devices>
  <memory model='nvdimm' access='private' discard='yes'>
    <source>
       <pagesize unit='KiB'>4096</pagesize>
       <nodemask>1-3</nodemask>
    </source>
    <target>
      <size unit='KiB'>524287</size>
      <node>0</node>
    </target>
  </memory>
</devices>



<devices>
  <memory model='nvdimm' access='private' discard='yes'>
    <source>
        <!-- could be a memfs/ramfs or super fast nvme hd -->
        <path>/tmp/nvdimm</path> 
    </source>
    <target>
      <size unit='KiB'>524287</size>
      <node>0</node>
      <label>
       <size unit='KiB'>128</size>
      <label>
    </target>
  </memory>
</devices>
````

dimm = dimm module
nvdimm = non-volitile


access = `private` or `shared` (same as memory backing)
dsicard = can fine tune discard of data per module only dimm



[shared memory]
share a /dev/shm on each computer, probalby super fast, but limited scope

````
<devices>
  <shmem name='shmem_server'>
    <model type='ivshmem-doorbell'/>
    <size unit='M'>2</size>
    <server path='/tmp/socket-shmem'/>
    <msi vectors='32' ioeventfd='on' />
  </shmem>
<devices>
````

server = provides unix socket on host defaults to `/var/lib/libvirt/shmem/name-sock

msi = off/on msi interrupts. 

model = ivshmem-plain (severless) vs ivshmem-doorbell (shmem server)

[vsock]

````
<devices>
  <vsock model='virtio'>
    <cid auto='no' address='3' />
  </vsock>
<devices>
````

_______________________________________________________________________________
## Multiverse OS General
[unikernels/microkerns][==========================================================]
ilackarms/unik - go based unikernel compilation

ovs - sepcial single binary OS


[Custom FS]

[^][fuse]
jacobsa/fuse - comes with great examples, and is modern fuse. errofs, flusfs,
forgetfs, hellofs, memfs, cachingfs, interrupt,fs mount sample


[apt repo server]

aptly-dev/aptly - big

ayufan/debian-repository - small

[make deb packages]
xor-gate/debpkg


[INSTALLER]

go-debos/debos - debian os builder
uses yaml and deboostrap and such to build custom debian. **best looking**


solus-project/uspin


animuspexus/aipsetup - system for creating and maintaing own GNU+Linux distro
ported from python


[systemd in go]
lostinblue/systemgo


[gsettings-upd][gnome]

arl/gsettings-upd


[MOTD!!!!]

victorgama/howe


[journald]

ssgreg/journald
send msgs to journald

_______________________________________________________
## Unix sockets


vaitekunas/unixsock

hodgesds/sockopts-go

navinds25/unixsockets_expts


[custom protocol]
aki237/up - very cool, lets you setup a very simple http like protocol
to send data over sockets

_______________________________________________________________________________
## Wayland and Xserver


mikkeloscar/flis - most complete wayland thingy


zenhack/go.wayland

elliotmr/wl

yuhang/xserver


linuxdeepin/go-x11-client



sworne/golock


[Multiverse OS/Portal-Gun Installer/Login/Effect]

peterhellberg/pixel-experiments
  * tunnel could be modified to look like a wormhole or sliders tunnel, and can be
used as a transition with loading bar.

  * same with starfield


_______________________________________________________________________________
## OHT CHat 

[emoji]
peterhellberg/emojilib



_______________________________________________________________________________
## Cbor (and other datatypes)

cbor is a binary format. c like json but better for sending data

bnclabs/gson - json, cbor

mozilla-services/go-cose - cbor object signing and ecnryption


segmentio/objconv - yaml json cbor convert

toravir/csd - cbor stream decorer



2tvenom/cbor 

_______________________________________________________________________________
## Inverse PRoxy


hkparker/imux -inverse mux


_______________________________________________________________________________
## KCP


ccsexyz/kcp-go-raw


_______________________________________________________________________________
##  Netlink


subgraph/go-nfnetlink

mdlayher/netlink

_______________________________________________________________________________
## VSOCK


mdlayher/vsock

_______________________________________________________________________________
## p9 file sharing


docker/go-p9p


deedlefake/p9
_______________________________________________________________________________
## SHM 

jeppeter/go-shm

_______________________________________________________________________________
## Transpile C programs to GO


elliotchance/c2go


rsc/c2go


_______________________________________________________________________________
## RTMP / Screen Sharing / SRS


[rtmp]

zhangpeihao/gortmp

c-bata/rtmp


[multi-protocol]

ossrs/go-oryx
srs, rtmp

_______________________________________________________________________________
## bloom

**tylertreat/bloomfilters**
tylertreat/inversebloomfilter

_______________________________________________________________________________
## socket frameworks

tld extract
joeguo/tldextract
_______________________________________________________________________________
## socket frameworks


henrylee2cn/teleport

p2p, revers,e rpc, micro, game, etc huge


wpajqz/linker - like every socket type and network type and ways to linik them
_______________________________________________________________________________
## Raw Sockets
newtools/zsocket - zero copy socket **BEST**



hashicorp/go-sockaddr



crznic/crznic - cooooooool




naoyamaguichi/rawsocket_practice

monochromegane/smux - socket multiplexer



bisreal8191/sniffer


ryabuhin/golang_linux_rawsockets


kdar/gorawtcpsyn
_______________________________________________________________________________
## mTCP


aclarembeau/go-mptcp-api

_______________________________________________________________________________
## SystemD + SysV and creating daemons/processes


[dev][==========================================================]
[process][==========================================================]
[service][daemon][agent][==========================================================]
jvehent/service-go - good service/agent/daemon example that works for systemd
and systemV


[^][system process][system V]
teepark/gosysvipc
[^][system process][systemd]
coreos/systemd


[BOTH]

oneumyvakin/initme


miros/init-exporter
_______________________________________________________________________________
## Pacckage manager


ciscocloud/distributive checks/packages - has [alpine] `apk` and [debian] `apt-get`
systemctl code  users and groups
  * also has checksum for filesystem


_______________________________________________________________________________
## NETOWRKING


[firewall]
[^][packet filter : pf]
go-freebsd/pf
_______________________________________________________________________________
## BOOTING


[NETBOOT]
google/netboot - in go
  dhcp4, dhcp6, pcap, pxieicore, tftp, thirdparty, etc


  pixiecore looks cool, debuan buntu is ready for it. pixiecore binary is built statifcally and syhould work for all distros but is built on ddebian

  has nice ui

  command line all-in-one-tool for easy netbooting


_______________________________________________________________________________
## Portal Gun



[image]

google/embiggen-disk - shrink and grow disk on-the-fly automatically





[virt networking][virtio][virtserial][virtsock][virtchannel][==========================================================]
stefanha.github.io/virtio - spec docs!

**virtual io**

[^][vsock]
*xpra uses vsock (NEW!) || wiki.qemu.org/Features/virtiovsock*
can accept connections from multiple clients (unliek virtserial)

**We can use vsock to do connections between Vms to and not have tcp/ip addresses which reduces attack surface**

deos not require networking
posix sockets api so minimal mods (unlike virtserial)

mdlayher/vsock - includes cool vcp or scp like tool to moving files across sock


portalgun-io/echo-vsock - simple server/client examples

clownix/cloonix_vsock - provides BASH/TERMINAL over virtsock (c program) maybe
transpile?




[^][virtserial][ist shit compared to virtsock!]



[^][virtchannel]
0xef53/phoenix-guest-agent - uses virtio-serial port to do commands
says virtserial, but its clearly using socket connection specifically virthcannel style

should drop the envent loop it uses in favor of *evio*

*uses virtserial but can take concepts from it*


[vm][==========================================================]

[^][spice]
jsimonetti/go-spice - can use html5 to show and interact with screen :)

[^][qemu]
**UP TO DATE DOCS**:
   qemu.weilnetz.de/doc/qemu-qmp-ref.html
   qemu.weilnetz.de/doc/qemu-ga-ref.html - guest agent protocol
   


intel/qemu-lite - slimmed down qemu C 


[^][^][qemu ineteraction via commandline]
quadrifoglio/go-qemu **VERY NICE programming API, one of bet**

zeropage/vm-manager - has a nice yaml based config we can start with good starting point

mirantis/virtlet - REALLY GOOD yaml configs, just look at those for examples at least

rafaelmartins/simplevirt


virtmonitor/kvm - bsic ineteraction




vodik/qerty 


giantswarm/onsho


github.com/mgibula/qemu-mgr


[^][^][TUN TAP]

go-tuntap - simple tuntap driver for qemu/kvm


[^][^][qmp and qemu in C]
usqcd-software/qmp

[^][^][qmp and qemu in Rust]

arcnmx/qapi-rs

[^][^][qemu via QMP!]
wiki.qemu.org/qmp
linux-kvm.org/index.php?title=monitorprotocol&oldid=3100


hyperhq/runv **STARTING POINT** BEST QEMU qmp!
  * sock2py
  * great bridge code
  * excellent tap code
  [^] hyperhq/hyperd image/tarexport/save-oci has code for layers
  [^] hyperhq/hypercli



intel/govmmmmm - **STARTING POINT!@!!** BEST qemu command based 

p0rtalgun/qmp - WOW **STARTING POINT!**


0xef53/go-qmp **STARTING POINT!**
0xef53/qmp-shell - great work **STARTING POINT!**


quadrifoglio/go-qmp 

h00s/goqemu - basic but maybe helpful, obvs the one above is the starting point

sgsullivan/qmpmeddler

BASH LIB - arcnmx/qemucomm

RUBY xanclic/qemu.rb

RUBY SEND KEYS VIA QMP  mvidner/sendkeys


[images][==========================================================]
[^][raw]
[^][qcowX][2,3]
vasi/qcow2 - fuse is using a stupid osx thing
zchee/qcow2


[^][multi layer file systems]


aosc-dev/ciel




[booting][==========================================================]

**vboot** - in u-root cmds is the way to sign kernel, initramfs ETC LOOK INTO THIS!!!

systemboot/systemboot
golang u-boot adaptation to verify booting 

**We should be USING THIS with a any kernel (debian, alpine, etc) to boot**
u-root/u-root - root file system (initramfs) written in go

This will let us carry singatures up from BIOS, it will enable all of our 
custom features. Including our own luks crypt options. It is aweomse. 

it even has qemu exmaples

**so instead of learning a garbage bootloader code that grub forces, we can
just write it in go, which is way easier. doing all the same things. can even
hide our presence by mimiccing go.**

        has a bunch of cmds gnu-cmds like lsmod, hexdump ,grep, ls, etc 
        we want to update these and make them consistnet while being
        bakcwards compatible
      

**infact a great zerg thign would be macosx looking boot, and linux grub looking boot**
       [^]exmaple u-root
            u-root/u-bmc

[^][dhcp - used for netboot]
  u-root/dhcp4

[^][init]

[^][^][examples]

lostinblue/go-init

[^][^][systemd init]
[^][^][sysV init]
peterbourgon/runsvinit



[^][netboot][tftp][pxe][ipxe]


[^][^][netboot examples]
giantswarm/mayu


coreos/matchbox - yaml configed






[cluster management][==========================================================]
jocko - kafka/distirbuted commit log service, single binary, no deps

xcat2/goconserver

[CLUSTER UI]

murdinc/awsviz - a center out branching tree (more like mycelia)

_______________________________________________________________________________
## Wayland

scrambleshell/external-weston
scrambleshell/external-wayland



lostinblue/small-gpu-rendering-examples

_______________________________________________________________________________
## CONTAINER containers containers
**the good containers we will be using, to implem,ent routers for example**

* google/gvisor

opencontainers/runc

opencontainers/runtime-tools


opencontainers/image-tools


opencontainers/runc
  * **REALLY GOOD PROCESS CODE** has resore process code
  * *Good user code* and group code
  * has init code


opencontainers/selinux - **SELINUX IN GO!**


vishvananda/wormhole - good tunneling code, can probalby grab a bunch of stuff


vishvanananda/gocapability


[KATA RUNTIME]

inside virtcontainers/device/drivers/*

we can find VFIO device and block device 

  has `vhost_block.go`
  has `vhost_net_user.go`
  has `vhost_scsi.go`

  all very interesting


   nice qemu launching based on command and supprots LOTS of differnet yteps!!!

_______________________________________________________________________________
## libguestfs (BEST IMAGE MANAGEMENT!!!!)


libguestfs/libguestfs/golang

_______________________________________________________________________________
## LINUX KIT


pkg/host-timesync-daemon **THIS SYNCS RTCTIME OVER VSOCK!!!**


pkg/init/cmd/init and rc.init ARE GREAT INITS!!!! has mkchar !!! (the complex one that is mkdev + mknod)


memlogd - built to grab logs and stick them in memfs and put em in ciricular buffer

_______________________________________________________________________________
## VM UI

novnc - novnc/nonvc 

gospice html5 - go-spice
_______________________________________________________________________________
## LAYER
 IMAGES



[rootfs]
cloudfoundry/grootfs


[ruby]
lostinblue/rootfslan 
_______________________________________________________________________________
## VM Devices


vishvananda/open-isci - interesting high pref transprot independent multi platform
partitioned into user and kernnel parts



[virtio-forwarder]
netronome/virtio-forwarder
 is a supersapce networking app that forwards bidirectional traffic between SR-IOV (uses zero mq)

flaws: why not direct? why put control under host machine? **BUT IT HAS A LOT OF GOOD RESORUCES**

````
<devices>
  <interface type='vhostuser'>
    <source type='unix' path='/tmp/virtio-forwarder' mode='client' />
    <model type='virtio' />
    <alias name='net1' />
    <address type='pci' ... />
  </interface>
</devices>
````

````
<devices>
  <interface type='vhostuser'>
    <mac address='00:00:00:00:00' />
    <source type='unix' path='/tmp/virtio-forwarder-sopck' mode='client' />
    <model type='virtio' />
    <alias name='net1' />
    <address type='pci' ... />
  </interface>
</devices>
````

````
<qemu:commandline>
  <qemu:arg value='-global' />
  <qemu:arg value='virtio-pci.disable-modern=off'/>
</qemu:commandline>
````
  


[interesting exmaples]


[firmware]

intel/fwupd - updating firmware on linux automatic safe and relaible, linux vendor firmware service **INVESTIGATE THIS!!**


intel/ixpdimm_sw - firmware and other code related to intel dimm memory. MEMORY FIRMWARE!


intel/ipmctl - persistent memory module code - discover PMMs, montior health, view and update firmware, configure data-at-rest-security debu and troubleshoot.  **holy crap**

_______________________________________________________________________________
## Bonding Bond NEtwork Device
**Bonding provides a method for aggregating multiple interfaces into a single logical onbded device.** This can provide network redudandncy, i.e. have two ISP connections? bond the devices.

This is specified in 802.3ad spec Linux bonding drivers provides various falvors of bonding interface depending on mode or bonding polices such as round robin or active aggregation (probably mtcp comes in handy here too).


intel/bond-cni


vishvananda/netlink - super rich userspace program in linux to communicate with kernel (netliink) and this talks to that **CAN USE THIS TO ADD BRIDGES!**

vishvananda/netns - namespaces, like what containers use


_______________________________________________________________________________
## CLUSTERING

 * [cloud integregated advanced orchestrator] **wow actually very very nice 
and hits a lot of the multieverse os feature requirements**

ciao-project/ciao

**ui** intel/ciao-webui
_______________________________________________________________________________
## QEMU vs NEMU vs Qemu-lite


intel/nemu - fork of qemu 




intel/ccloudvm - controls qemuy to do workloads creating developmpent or demo enviroments for complex projects
        check out /ccvm/vm.go


[resources]

intel/rmd - resource managemnet daemon provide uniform interface portal for hardware resource manatgement on x86

[related tools]


sriov-cni - code for virtio-pci is in here

_______________________________________________________________________________
## Router Controller Setup
Below is the guide to configuring the network controller on Debian9.

#### Package Configuration

**TODO: Disable apt-get suggestions and recommendations** so we can delete 
things like firefox without chrome being installed automatically.

````
sudo apt-get update
sudo apt-get upgrade

sudo apt-get install virt-manager ovmf libvirt libvirt-daemon qemu qemu-img

sudo apt-get install vim

sudo apt-get remove nano firefox 
````
#### Storage Configuration
Use the user session libvirt configuration, add the following storage
directories:


  (1) base-images
  (2) os-images
  (3) portal-images
  (4) network-controller (the 'portal' folder, which contains this guide and
      the files and config details needed for setup and configuration)

#### Galaxy Configuration
Add a single plan9 share (Driver: Path, Mode: Mapped, Cache: None). 


Path will be to the `galaxy.portal.multiverse` named `galaxy-portal`. Then finalize the 
setup by adding the following line to the newly created VM's `/etc/fstab`:



````
galaxy-portal /mnt/galaxy-portal 9p trans=virtio,9p2000.L,posixacl,rw,cache=none,nofail 0 0
````

Reboot, and if the mount is showing up, run the two scripts that handle
setting up the router.

````
cd /mnt/galaxy-portal
./packages.sh
./provision.sh
````

Next install their open VPN configuration

#### Virtual Bridges On Each Network Controller
Using three bridges, the first one implementing DHCP, and other features which
enable it to function as a router is the `universe0` router. Then the other
two (2) will only establish the subnet. This model enables simple creation f
several galaxy networks isolated with their own transparent proxy routing. 

  (1) universe0
      **subnet** 10.0.0.0/24
      **dhcp** 10.0.0.2-254
      **mac** 00:00:10:00:00:01
      **ip** 10.0.0.1

  (2) galaxy0
      **subnet** 10.1.1.0/24
      **dhcp** disabled
      **mac** 00:00:10:01:01:fe
      **ip** 10.1.1.254

      *note* ideally we disable access to this bridge altogether from the host

  (3) star0
      **subnet** 10.2.2.0/24
      **dhcp** disabled
      **mac** 00:00:10:02:02:fe
      **ip** 10.2.2.25

      *note* ideally we disable access to this bridge altogether from the host
      and if we opt to use whonix it will be 10.152.152.0/24

#### Connecting the routers to the higher level virbr0's
Using `session` in virt-manager (for now), we can experiment with the 
"Network Interaces" tab, with `brX`, `bondX`, and `vlanX` for connecting
the galaxy router and star router to their respective etherenet slots.





````
2: ens12: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 00:00:10:10:10:01 brd ff:ff:ff:ff:ff:ff
3: ens13: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master br1 state UP group default qlen 1000
    link/ether 00:00:10:11:11:01 brd ff:ff:ff:ff:ff:ff
4: ens14: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 00:00:10:12:12:01 brd ff:ff:ff:ff:ff:ff
````








_______________________________________________________________________________
## Notes









