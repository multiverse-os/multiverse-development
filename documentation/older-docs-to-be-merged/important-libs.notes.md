# Important libraries


#== C ==============================================================

libosinfo - provides all the info we need for download links, location of installation trees, iso images, llive cds

recomended CPU/memory/disk for an OS
idnetify what OS an ISO is
query hardware supported by OS
query what hardware is supported by hypervisor
determine optimual hardware for running a os/hypervisor

generate scripts for deskotp installation autoamting

has ruby bindings 

[can we just do C to go?] yeah be4cause its mainly just relying on XMl files, we could conver tthose to JSON, serve them up via an API, then ad dthe extra data we need because LOL it doesnt even include the sigs and checksums with the DL linsk :(

#== Rust ==============================================================


#== GO ==============================================================
## best sys info
zcalusic/sysinfo - still need something to parse bios shit i think



## libvirt

libvirt/libvirt-go

## programming tek

go-playground/locales

packtpublishing/go-design-patterns

cloudfoundry/go-diodes

joshuarubin/zb


## Security
awnumar/memguard


## net


newtools/zsocket

[mesh networking]
weaveworks/mesh

[stream multiplexing]
hashicorp/yamux

kcp

inconshreveable/muxado

xtaci/smux

[udp]
kcp
spance/suft

[xmpp]
agl/xmpp-client

[webrtc]
keroserene/go-webrtc


[hash tables]
secondbit/wendy
sioro-neoku/go-peerflix

stefankopieczek/gossip

wjh/dfi

armon/go-chord

## compression

[streaming]
golang/snappy


## video

vulkan-go/vulkan

## CLI/TUI

vladimirmarkelov/clui

[interact/input]
goput

deiwin/interact

## os / operating system

[systemd]
coreos/systemd

iguanesolutions/go-systemd

[gsettings]
arl/gsettings-upd

[libinput]
gvalkov/golang-evdev

[dbus and ipc]
godbus/dbus




[lvm]
google/embiggen

[cmd execution]

go-cmd/cmd

fd0/machma - paralell execution of commands


shenwei356/rush - jobs in parallel

[pty]
kr/pty



[syslog and other logs]

bsycorp/log-forwarder

papertrail/remote-syslog2


[provision/orchestration]

dnaeon/gru

vektra/tachyon

ctdk/gioardi - chef in go

[init]

jjyr/bootgo


driusan/dainit


[files]
omeid/gonzo - file processing framework

[cgroups]
containerd/cgroups

[socks]
h12w/socks

## debian

go-debos/debos

## web

[ui]
gowut


## gpu

barnex/cuda5

## db / database


djherbis/stow object persistance

hashicorp/raft-boltdb


kentik/patricia - ipcidr tagging patrictia tree


## crypto

flynn/noise - great way to make cyrpto protocols

## devices

[net]
mdlayher/*
mdlayher/netlink
mdlayher/raw
        /vsock

google/netstack
google/packet


