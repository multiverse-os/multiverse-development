# Portal Gun Libraies
**we are NOT getting rid of libvirt, we are NOT getting rid of virsh** maybe we mod it, maybe we control the storage pools with a program that abstracts that bullshit away but **VIRSH IS FUCKING POWERFUL**!!

  * other people thought about virsh having too much shit: lindorffs/virshyt


## mpTCP
multipath tcp

## KOBJECT by mdlahyr
this is neede dto see usb events in userspace or similar shit.
----

## go-conntrack
can use this to packet isnepct and route
__


Maybe we just rip features from it, because it may be TOO powerful.



you can snashot more than disks, but also settings


virsh iface-bgin creates a snapshot of the current host interface settings, then later virsh iface-commit saves it, virsh iface-rollback reverts.


mdlayher/vsock

mdlayher/netboot

provides acces to linux VM sockets AF vsock for communication between a hypervisor and its machines


___________________________________________

## DIRECT KERNEL (HOW TO)

when installing a new guest machine system it is often useful to boot direclty from a kernel and initrd stored in the host physical machine operating system 


this allows arugments to be pased DIRECTLY to the installer 

this capabiltiy is aviable for both fully vrit and paravirt ((CHECK OUT REDHAD DOCS ON THIS< ITS GOOD)

## GuestFS 

modify disks VIA C library!!!


## LIMIT BLOCK IO:

virsh blkdeviotune domain device --total-bytes-sec --read-bytes-sec --write-bytes-sec  --read-iops-sec --write-iops-sec


## DIF A DISK IMAGE!!!

virt-df

virt-diff can be safely used on live guests because it only needs read access!!! dsoes not check boot loader unused space between partiiosn or within file systems or hidden sectors

so not a great foresenics tools

## grpahical frame buffer

<devices>
  <graphics type=sdl display="0:0">
vnc
rdp
desktop 
spice - can even dfein network

spice has complex options

chanel names

with multiple channel support like one to record and one to send


image compression

streaming mode

clipbooard
mousemode


(We should modiufy spice graphics not QXl!)


## video elemnt desribes video devices


can have ram and vram (this is how we could pull directly from the ram but then we would need to render it twice like a dummy.


## Consoles


<serial type='pty'>
  source path=/dev/pts/3
  target port=0

<channel type=unix>
  <source mode=bind  path=/tmp/guestfwd
  target type=guestfwd address=x port=x


**PARALLEL AND ISA PORTS ARE NO LONGER SUPPORTEED FUCK YOU VIRTMANGER!**

usb-serial

0 1 2 serial ports usually exist


address with type usb can tie deviec to a particular controller documented above


KVM VIRTIO CONSOLE


<console type=pty
  source path=dev/pts/5
   target type=virtio port=0


guestfwd - send content to TCP prort

virtio is paravirtualized virtio chanenel

is expossed in the guest virtual machine under /dev/vport*

and if the name is specified /dev/virtio-ports/name


type virtio serial on address element
<address type=virtio-serial

this would let us tie it to a specific device controller virtIO seiral




**look at 23.18 devices**
REVIEW THIS!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!11

NB special case is itneresting, 
pty console, ETC

MOST INTERSTING IS __**HOST PHYSICAL MACHINE DEVICE PROXY!!!!**____

OR ___**NAMED PIPE**__ OR EVEN __**UDP NETWORK CONSOLE**__ (combine with kcp which would make it super nice)


CAn use named pipe for example on our VIRTUAL SOUND CARD!!!!



## MEMORY BACKING
**WE NEED TO SET THIS!!**

<memorybacking>
<hugepages>
  <page size=1" unit"=g" nodesetx

<nosharedpages/>
</

## BIOS

<os>
<smbios mode='sysinfo' />
modes:
  * emulate: allows hypervisor to generate ALL values!

  * host copies all of block 0 and block 1 except for UUID from host physical bios

virtConnectGetSysinfo
can be used to se what values are copied

or sysinfo values in the systeinfo elemetn



<bios useserial='yes' rebootTimeout='0' />
  use serial allows bios messages to be seen on serial display


  <initarg> --unit</initarg>
<initarg> emergy.service</initarg>

specifay init here too

<init>/bin/systemd</init>
**can we just use this for signed booting?**

<sysinfo type="smbios">

can fill in own values. ew can use these values like UUId to store cryptographic info to validate and make determinstic building

</os>
_______________
## list devies

virsh nodedev-list --tress


virsh ndoedev-list --cap (usb|scsi|and so on)

## SEND KEY STROKES TO GUEST!!!!!!!!!!!!!!!!!!!!

`virsh send-key *domain* --codeset --holdtime keycode`

options for `codeset` is: [linux, xt, atset1, atset2, atset3, xt_kdb, win32, usb,rfb]

rfb- numberic values are those defined rfb extension for sending RAW KEYCODES.

usb- numberic values are defind by the USB HID specificaTION  for keyboard input. 

**USING these we could defintely prototype a custom keyboard that sends all keys ENCRYTPTED through the machine and decrypts in the kernel of the arrivng machine. OR send keys to multiple machines at once.** *could have a server that recives the keys and decides which machines to send the keys to.*



## virt-sparsify: reclaiming empty disk space
## virt-dff differences betwen virt machine files

## stats

virsh nodecpustats

virsh nodecpustats 2 --percent

lsusb -v

## INETFACE devies (support PCI passthrough now too)

has really cool flags, like 

--script - use this option to specify to a pscript file handling a bridge intead of the default one. 

--inbound bandwith
--outbound badnwith

### virsh nodesuspend (disk|mem|hybrid) 90 (rtc; in seconds we wake up the VM from suspend 

THIS CAN GET OUR VIDEO CARD BACK!!!!!!

virsh node-memory-tune [shm-pages-to-scan] [shm-sleep-miliseconds] [shm-merg-across-nodes]


### virh rescue, on the fly rscue cd for guests!

### Scripting guest configuration (se Augeas and libguestfs scripting for examples)


[see]: *readhat customer protal 21.4 the guestfish shell*

examples:
  * get locale of host and apply to guest

  * scan subnets, track open and used addresses and ports open/closed on NATed subnets and considedr when applying ntework settings.

  * every VM should have an optional router attached. stick a DNS server, DHCP, etc basically full single binary router outside each VM. Then script how this is setup._use this to route packets to different networks based on packet details_

  * script connections to tor; spin up 32 instances, discard based on: speed, too many exits from same country, if it is a known tor node, if its considered malcious, etc.

  * route packet traffic

  * list packages to add, to delte from vanilla

  * select operating system, it will autogenerate the URL, checksum, check sigture, etc.

  * assign code to kernel module using an empty kernel module which can run ruby code, or go code, or rust code, or etc. **(Think about how there is that gnome extension that lets you run some code or a script, it lets you create a huge variety fo complex gnom extensions then with rust or go instead of just js/css)**

  * do the same thing as above but for devices /dev/*

  * select kernel modules, remove them

  * define kernel flags, initramfs flags

  * initramfs, and bootloader both shold be sub-config-scirpts, basically like this one but nested.

  * define default application and whitelisted applications if APP Vm. if jailed. number of instaces ran. etc.

  * Define services that MUST be running if service VM. define if applications are jailed. number of isntances ran etc.

  


### REQUIRE ALL KERNEL MODULES TO BE SIGNED

### qemu-img 

`qemu-img` should be used driectly. **virt-manager** is super  dumb for trying to snapshot the WHOEL MACHINE only. when you can snapshot individual HDs using qemu-img. So if you have pflash or even raw disks you jhust snapshot the qcow2 image or btrfs. Then do backing up for other disks different wwyas.

## always specify virtio as bus on device drivers

<target dev="vdb" bus="virtio" />

### Applicaiton/Service/Controller VMs (Alpha)

Below is a tree list of the available portals (VMs) that Multiverse OS alpha release is expected to be released with (this list is still subject to change). 

````
portalgun
├── application
│   ├── browser
│   │   ├── chromium
│   │   ├── firefox
│   │   └── torbrowser
│   ├── development
│   │   ├── c
│   │   ├── cpp
│   │   ├── go
│   │   ├── ruby
│   │   └── rust
│   ├── filebrowser
│   ├── media
│   └── terminal
├── controller
│   ├── router
│   └── user
└── service
    ├── crypto
    │   └── wallet
    └── router
        ├── galaxy
        ├── star
        └── universe
````


**Getting rid of libvirt and virt-manager**

 * In order to get rid of virt-manager, we need a way to interact with VMs without requiring the graphics card attached and spice viewer in virt-manager. Virtmanager sucks at making XML, it does not have more than 70% of the total actual capabilities mapped to the GUI, and it just funcitons terribly.

We should begin before any VM is created, running SMBIOS analysis (bios analysis for info on motherboard etc, do sysfs+procfs analysis to get all pci cards info, etc, get kernel info (kenrl moduels etc), packages installed, and `virsh capabilites`. Then combine all of this into an API accessible profile of the HOST (and controllers that function as nested HSOT). This will allow us to submit anonymous reports and slowly improve and optimize options to fit differnet profiles. + `virsh domcapabiliteis` gives us info on iommu, (which should automatically be fixed via perment grub addition of `amd_iommu=on` and `intel_iommu=on`.



 We could definitely get rid of it, if we had a CLI wizard, that would go through each question, like VM Class [`controller`, `service`, `app`]. The next defaults would be based on previous choice. Then go through each item: [Name, BIOS, CPU, Memory, Disks/Boot options, Network Devices, PIC cards, channels (serial, parallel, console, channel, etc)](channels should be standarfdized more, like output status file, using serial file out and using a standard folder structure (but can be customized obviously using optional attributes) 

Ideally this will generate a YAML file or something similar, maybe a ruby config. We could add a 

Controllers (USB, PCI, VirtIO, SCSI, etc should just be automatic, no reason to ask. In addition, decisionsm mde like PCI passthrough, should automaticaly generate script to fire on VM start and VM stop, so it grabs the PCI card automatically and releases it back automatically, in a way its able to run unpriviledged.



 * To get rid of libvirt


````
 <mode name='host-passthrough' supported='yes'/>
    <mode name='host-model' supported='yes'>
      <model fallback='allow'>Skylake-Client</model>
      <vendor>Intel</vendor>
      <feature policy='require' name='ds'/>
      <feature policy='require' name='acpi'/>
      <feature policy='require' name='ss'/>
      <feature policy='require' name='ht'/>
      <feature policy='require' name='tm'/>
      <feature policy='require' name='pbe'/>
      <feature policy='require' name='dtes64'/>
      <feature policy='require' name='monitor'/>
      <feature policy='require' name='ds_cpl'/>
      <feature policy='require' name='vmx'/>
      <feature policy='require' name='est'/>
      <feature policy='require' name='tm2'/>
      <feature policy='require' name='xtpr'/>
      <feature policy='require' name='pdcm'/>
      <feature policy='require' name='dca'/>
      <feature policy='require' name='osxsave'/>
      <feature policy='require' name='tsc_adjust'/>
      <feature policy='require' name='avx512f'/>
      <feature policy='require' name='avx512dq'/>
      <feature policy='require' name='clflushopt'/>
      <feature policy='require' name='avx512cd'/>
      <feature policy='require' name='avx512bw'/>
      <feature policy='require' name='avx512vl'/>
      <feature policy='require' name='xsaves'/>
      <feature policy='require' name='pdpe1gb'/>
      <feature policy='require' name='invtsc'/>
    </mode>

````
____________________________________________________________________

get the alpine single binary+ jpmorgan distrubted trust ledger+
torrent tracker + torrent client (for FS with fuse)+



SSE

---


## time to boot server

jponge/time-to-boot-server

## shm

## device-sdk-go
a simpel virt device system

## Qemu 

 * schanpps - kvm tooling (dhcp client, dns server, libvirt, qmp) 

 * go-spice - spice proxy library in go

## video
gocast ? very similar to what we want at least fopr streaming 1 way


kaey/framebuffer

blackspace/gofb

## netboot
google/netboot
## Storage management
buse go

  /lostinblue/storage
  
  github.com/google/embiggen


teemow/loopback - loopback without lvm

mdhayler/block programmatic access to block devs

## PCI
cacluate how much money a user has spent based on computer resourfces used

## pty

jayshwa/go-pty

## OHT
gholt/devicering

## spi

eec1/spi

## sound
check out sonos or opus transfer

tv42/sinus sonos upnp audio devices


musicsync
## Networking

Insert packets raw into the router NICs or VM nics on a PCIE card. 
