# Development Journal
The on-going evolution of the Multiverse OS design. Logged as journal entries over time with the newest first. 


** WE WILL NEED TO KEEP LIBVIRT (for VIRSH) or fork it!**

virsh domcontrol *guest*
virsh domstate (running)

virsh domjobinfo guest1


virsh domjobabort guest1

## LUN CAN BE SHARED BETWEEN GUESTS!

a disk to shared betwen quests instead of 9p

## MAP LAYERS IN QCOW

qemu-img map

## there are ALOT of USB device types!! REVIE THIS 16.4 red hat docs


## tftp boot

<tftp root=/varlib/tfpboot'

  <bootp file="boot file anem"

## LOTS OF PCI too!!!

pci-root

pcie-root

pci-birdge

dmi-to-pci-bridge


AND WE CAN MIX THIS WITH MMIO VIRTIO, use UDP with RCP over it??

**can delegate to external snapshots**



this is a kernel backend, the nromal one is supposedly userpsace

<interface type="network">
  <model type="virtio" />
  <driver name="qemu" />
</interface>

this makes it usespace appaerently


and this ameks ti kernel
## features
<pae>: increases physical address allows 3d bit guets to be bigger

<acpi: useful power management

apic lets us set the IRQ management or what to do with itnerupts even got eoi or end of interurpt where we can set on or off 


hap hardware assisted paging avialble in the ahrdare
hap 

## evemts configuration ##

<on_poweroff>destroy

<on_lockfailure>

<on_crash>action

how do we deinfe the scripts that make up the actions? Maybe we can modify lbivirt and customize it?


## MEMORY BACKING
**WE NEED TO SET THIS!!**

<memorybacking>
<hugepages>
  <page size=1" unit"=g" nodesetx

<nosharedpages/>
</



## vhost_net experimental_zcopytx=1
modprobe -r vhost_net

BUTT both of these rely on the kernel. WE CAN do VM-tO_VM direct transffer IT IS POSSSIBLE!!!!
________________________________________________________________
## gravity
make the FS torrent based, so it cant be fucked with

mutable torrents with signatures whitelsited to edit

## PCI
cacluate how much money a user has spent based on computer resourfces used

================================================================
## NOW
  * Build alpine base image for rebuild routers. 
  * Make router formula images
  * use virtio output file to indicate time-till(TT)-boot in ms. Then TT-internet and if the internet is required to reconnect. display WAN deets (ip) lan deets (ip). and a list of processes and details. all in CBOR or some other binary based format readable by computers. this until we can hook it directly to the controller.
  * Build a DB system to manage all the VM images, and even the scramble-id images. Make a qcow2 client that is capable of combining and disassembling qcow2 files into heriarchal trees using merkle trees (then transferalbe by torrent in individual chunks. so you can request just 1 sub identity on the fly from a remote machine.

  * VM-to-VM interface: special custom virtio pcie cards that enable connections between each other directly over /var/multiverse/interface on host 

  * take our single binary alpine see if we can insert in a bios or some other way to generate an environment that is deterministic and ecnrypted. 
  
  * build out scramble-id image (maybe raw or qcow2), inside store qcow2 of home folder for controller. but then also be able to generate on the fly home folders for specific sub-ids.


  * Build a kernel module that will provide agent support in all Multiverse OS machines. it will launch the process and keep it alive or reboot. this will be the foundation for security features and true lockdown of our various systems. (to avoid tainted messages, you have to include open source license in code)


  * start portal gun with libvirt-go[+ libvirt-go-xml] (or libvirt-rust)
    [look at libvirt sandbox for individual binary creating in c]
    + incorporate libosinfo - operating system info db 


  * start docs and making multiverse md, use chroma for blog,

================================================================
## EVENTUALLY
  * Build an example formula with a manifest file of some sort
    dictating preferably the ruby script (maybe include ruby, python
    crystal or other binaries in the build VM to give options)
    
  * A build VM that is built in a determinsitc trustless way
    to make determinstici trustless builds

  * experiment with building our own console using wayland+vulkan. It will interact with the linux input and output, (support multiverse os signed and encrytped input/output evnetually), print text crisp in 4k, support updatable selctions of text, ruby as a scripting instead of bash, -- you would start with a just CLI verison probably, then build the UI for it. have a sever so it supports coop mode out of the box, with multiple cursors (maybe you can take over, merge or split cursors). 

  * scmralbe id system server supports DNS like system, combine or merge qcow2 images or generate virtaul images exposing home var and tmp on the fly.

  * scrmable identity key will hold key/value databases to hold data crawled from net, etc
    * libguestfs = agent for modifying images of VM (like on-the-fly changing of HD size) 
________________________________________________________________


________________________________________________________________
________________________________________________________________
# Disk Images, Multiverse Data Layers, etc

Multiverse OS data layers refer to the concept of dividng the virtual machine data up into three (3) categories so that individual pieces can be immutable, ephemeral or both. 

* [layer 0][immutable][ephemeral]
  layer 0 is the OS layer, without: [`/var/`, `/home/`, `/tmp/`]

  the disk should only be r/w when `portalgun` is intializing the VM or when `portalgun` is updating the VM.

  integrity between VMs is maintained using a merkle tree contained within the disk. the merkle tree is compared against other instances to ensure no tampering has been done to the disk.

* [layer 1]
  layer1 is the `portal` (or VM) specific layer, containing all the files: source, documentation, etc relating to the Multiverse OS VM (or `portal`) needed to produce a trustless version of the VM, with the boot record signed by the **scramble suite ID** that initialized the original installation of Multiverse OS**

  this disk is attached during `portalgun` intiailziation and updates. otherwise it remains r/o and a merkle tree is maintained within it. 

  (the merk trees in each of these disk will allow simple transfer between peers via a bittorrent like system. `portalgun` will access this system and present a `apt-get` like package manager overlay abstraction to simplify downloading, installing, and maintaining secure, trustless updates for your MultiverseOS portals.)


* [layer 2]
  the last layer is where the [`/var/`, `/home/`, `/tmp/`] are located. in addition any other user data. it is encrypted and organized by **scramble suite identity**. This allows your user-data related to an identity and your `portal` to be separate, meaning any ID can be attached to any `portal` quickly and easily. this means that with a single `firefox.portal` you can spin up multiple versions tied to specific identities.

   __note__:[in the case of the `controller.user.portal` the generic high level **scramble suite identity** is loaded, which contains nested inside all of the sub-idenities that one can use when interacting with the internet. each identity includes stylographic analysis, times of day ones online, online accounts, and so on. evetually these identies will be sharable with other peers or entire groups] (does not have to be nested if we just encrypt them to keys inside the original disk, in addition, we could also shrink and grow the disk as needed. if size in an issue) but it would be nice to back up without needing to back up every associated identity but also in some cases it would be nice to pack them all together.)

________________________________________________________________
# Hiding the fact that we are in vms notes

  * lspci needs to report realistic names (could be kernel patch for this)
________________________________________________________________
# Paravirtual graphics card (custom multiverse os virt gpu) notes

[Terms]
DMI (Direct Memory Interface)

DRI (Direct Rendering Interface)

DMA (Direct Memory Access)

SHM (Shared Memory)

DRM (Direct rendering manager) - a linux kernel sub-module that interfaces with the GPU (video card). DRM exposes an API that user space programs can use to send commands and data to the GPU.

https://upload.wikimedia.org/wikipedia.org/wikipedia/commons/6/62/High_level_Overview_of_DRM.svg

The primary library to interact with DRM is `libdrm`. 

 _for low level access, check out the DMI Mapping API_

VRAM (Video Memory) - Memory located on the video card dedicated to use for the GPU.

#################################################################
## Vitual GPU (like QXL) that sends accesses memory directly and
sends it in compressed format (snappy?) to the Multiverse OS
user controller, where it will be converted to wayland protocl
message format and injected directly into a window.

**Understanding this will be super important to support building our own KVM to have all our computers in our server cabinenet and just have mouse, keyboard and mouse on our desks**

In the end, we are trying to create interprocess communicaiton,
specifically VM-to-VM (nested, but sometimes not), getting data
from the applicaiton surface data (window) to the GPU memory
of the USER CONTROLLER VM.

[IPC][The many different ways for processes to talk]



[Important][How does X over SSH work]

Our current method to solve this problem before we implement
a Multiverse OS solution is to just pull windows by using
X over SSH. This solution is great because unlike options like
`xpra` the (App)licaiton VM does NOT need a desktop environment.
Meaning they can be much more lean.

So we should learn how this works, so we can leverage this process
and send it over VirtIO via a virtual PCIe card instead of the
compartively clunky SSH protocol (its not even UDP). 

	__Since we are moving essentially stremaing video, UDP is for obvious reasons a better choice.__ So even if we just moved the X
over SSH method to a UDP protocl, we could get huge gains.

**X over SSH: How it works**
A basic explanation is that a program like firefox, implements inside it a `x client`, and the SSH server runs an `x server` that have mouse, keyboard and monitor plugged in. 

So SSH intercepts the connection to the local X server and sends it over the SSH connection. IN the same way that SSH can rewire other aspects of the linux infrastructure (see port forward, reverse port forward, and so on in SSH).

In Multiverse OS, X is being abandoned for Wayland, for a variety of security related reasons are alone enough to force the switch. But we have unique requrements with Multiverse OS. So we will be using a combination of virtual linux devices (*virtual PCIe video card*) and wayland to implement this without needing a desktop manager on each (APP)licaiton VM.


    [X over SSH][Resources]
	(url|http://dustwell.com/how-x-over-ssh-really-works.html)


And so we will be looking into wayland, and we have already begun
this work, and it is named: Milky(wayland) and currently is
written in Go. (But we will most likely be moving this to Rust.

[o-][An overview of `Xwayland`]
But first lets look at Xwayland, this is an `x server` running as 
a `wl_client`. 

It is capable of displaying native X11 client applications within the wayland compositor environment. This is similar to the way XQuartz runs X applicaiton in macOSs native windowing system. **this means we may be able to do the reverse and using *windows*, *linux*, and *osx* windows all available and workingly seamlessly together in our USER CONTROLLER VM.**


[O_O][How do we protect the /dev/card0 memory and other aspects that we want ideally isolated from the HOST (in case it is comprimised) and only accessible from the USER CONTROLLER VM?]

We could use encryption, where the key is only inside the VM. This is what we are doing for many other aspects of the system (RAM, HD, NET, and so on). But encrypting data in the card would undoubtly add delays and lag. What else can we do? What other components of linux and our VM do we need to worry about?

[USER SPACE][not always a good idea]
its faster, it prevents breakouts by avodiing sending packets through the kernel BUT it means we have to control who can send packets to the cards oursvelves, othwerise it can be super dangerous with DMA moving data around.

## DPDK 
data plane development kit - network interface controllers for fast packet processing

userspace networking with dpdk

#################################################################
## Multiverse OS User Controller Window Manager: Milkywayland

[o-][Wayland Compositor]
A wayland display server, a display server that implements the wayland protocol is referred to as a `wayland compositor`. 

They preform the task of compositing all windows (surfaces and associated objects) into a single coherent image to render.

[o-][libinput: one input to control them all]

`libinput` handles input devices for MULTIPLE wayland compistors and also rpvodes generic xorg server input as well. A common way to handle input events while minimzing amount of cusotm code needed for compositors to include custom devices.

[o-][wayland toolkits]

*clutter* supports ruby, and vala (bunch others too) to quickly assemble window/gui interface for your software. in this model of software desgin, you create a daemon, then give it a UI using clutter.

*efl*

*gtk3.20+* are wayland complete.

*qt5* is wayland complete*

*sdl* since 2.0.4 has good wayland support

*glfw 3.2* has wayland support


[Wayland Components]

*wl_display* - core global object, a special object to encapsulate the wayland protocol

*wl_registry* - compositor registers all the global objects it wants available to all clients

*wl_compositor* - handles combining the difference surfaces into a single output

*wl_surface* - an object representing a rectangular area of the screen defined by: [`location`, `size`, `pixel content`]

*wl_buffer* - attached to wl_surface boject, providing displayable content

*wl_output* - an object representing the displayable area of the screen

*wl_pointer*, *wl_keyboard*, *wl_touch* objects represent the differnent input devices like pointers

*wl_seat* an object representing a seat (a set of input devices: [`pointer`, `keyboard`, `touch`]

*wl_shell* [todo: explain this] not apart of the wayland protocol but is a stopgap solution to interface with existing software during development of wayland.

*wl_mempool*

[o-][Compositor: A closer look]
The wayland compositor cna define and export its own addtional interfaces. This feature is used to extend the protocol beyond the basic functionality provided by the core interfaces.

[o-][XDG-Shell protocol: Clear up some confusion]
XDG-Shell protocol is an extended way to manage surfaces under wayland compositors. The traditional way is to manipulate (using predefined actions: [`maximize`, `minimize`, `fullscreen`, etc] surfaces is the way the *wl_shell-*() functions that held in libwayland client.

xdg-shell protocl on the contrary is supposed to be provided by the wyaland compositor. 


[Example Software][How should Multiverse OS look?]

KWIN (KDE4) looks great, they mastered the box but they ahve other 3d window effects that look excellent. Many epople missed the entire point of the material desgin protocl but this applied using vulkan and hardware accelerating all the UI components could be incredibly powerful and beaitufl.

For example it would be great to "flip" a window and see all the details aobut the **scramble suite ID** associated with it, including its internet visibility timeline, and estimated location based oon that information, etc). 

**SCRAMBLE ID** should be assumed to be stored on a USB, can be a virtual usb but USB should be the place it is put!

Or flipping the notes program and seeing it in the hierachy of interconnections between notes. 

Or transversing the filesystem in a graph mode and hierachical mode at the same time.

#################################################################
## Direct Memory Access (DMA)

DMA or rDMA (remote DMA) is likely our most efficient path to 
moving window data between the controller and applicaition VM.

To start: DMA is not DMA, as in there are many types of DMA:

  * DMAe (engine)

  * DMAm (mapping)

#################################################################
## 8b and 10b encoding

PCIe uses 8b and 10b encoding, this is what the format should be to
move data around.

Supports easy multiplexing

#################################################################
## Interfaces: I2C, SPI, RI232 or how computers talk

[I2C]

[Userspace `spidev` userspace driver (SPI)]



#################################################################
## Shared Memory [Buffer] (SHM)

This was our first thought when approaching this problem, using 
zero-copy SHM inside of PCIe devices accessible from userspace 
to acheive highly efficient VM-to-VM transfer of data without
interacting with any of the kernels.

Wayland provides:

*wl_shm*

*wl_shm_pool*

The drawback to this method is that the compositor may need to do additional work (uslaly copy the shared data to the GPU) to display it, which leads to slower grpahics performance.

#################################################################
## Vulkan (Hardware (GPU) accelerated API, i.e. OpenGL, DirectX)

Client and compositor can share this GPU-space buffer using a special
handler refernece to it. This method allows the compositor to
avoid additonla copies of the data to the GPU, resutling in *faster graphics preformance over shm buffers*.

#################################################################
## The Framebuffer
The Linux framebuffer is lower level than the window manager (like X Server). You won't have any information regarding windows, locations,
etc becuase that is the whole point of the window manager. But you
can determine that manually, or run a super basic window manager, 
and use that to get window information and then switch to pulling info
from the fb based on that information. 

The framebuffer can be interacted with direclty very easily in linux:

For example, make snow by sending random to the framebuffer, like making a random file, sending random to a file.

`cat /dev/urandom -> /dev/fb0`

*Putting this on a timer would give us the snow style from an old TV which could be a cool effect.*

**Why use the framebuffer**
Because unless we are getting the window information from a window manager and then going even lower than the framebuffer to pull data directly from the memory of the video card (maybe this should be the idea) then the framebuffer is the layer we want to interact with.

#### Remote Framebuffer for linux
This topic is dealt wtih in p9, and it is in reality what we are looking for for wormholes to send windows from portals to `milkywayland` running on our controller.




[peer reviewed journal articles]
remote framebuffer review article: really good
(url|https://www.cs.indiana.edu/~adkulkar/papers/v9fb.pdf)




#_############################################################_#

________________________________________________________________


  * [idea][lets move away from p9 in favor of using some form of HD/IMAGE, maybe **nbd** etc

  [check out rust mentat db]


 [RESERACH][ for new CLI tools that fall back gracefully, one way is implemented by /usr/bin/sensible-editor] review this when working on this part


 [SHOW A] **create RAW storage volume, it supports max capacity and allocation in virt-amagner**


 [CUSTOM SNAPSHOT SYSTEM](auto-snapshot, basically versioned auto-backing up of VM image.)

  [__ALL PARTS ARE BUILT LOCALLY, USING EASY TO REVIEW/READ `portalgun`
     CONFIGURATION. ISOs are downlaoded, checksums are
     checked. then everything is built locally. 
       THIS MAKES MULTIVERSE OS portalgun VMs ACTUALLY SECURE
        because they are built in a trustless way.
     (atuomatically) and shown to the user. (later we will
     ask peers to verify results to add more security)__]
  **Unlike existing poorly desgined, and woefully insecure systems,**
  **like many of the container systems (that were thrown into production**
   *before v1.0, which was still running as root! goes to show its not*
   *just the reckles devs overhyping their product, but reckless users*
   *recklessly risking other peopels resouces]*

    [CONTEXT][Current VM image design is at least three (3) DISKS]
      [1]> OS Image (minimum space shurnk down maybe after install so it fits prfectly. (or MAYBE just swithc to LIVE ISO images for this, sitll decising) [!][READ ONLY][IMMUTABLE]
      [2]> Packages, configurations and other changes, making up the changes to the BASE image (just the installed standard OS) that defines the class/subclass of VM.
	[i.e. EXAMPLE: Base VM is: `app.debian`, which is combined with the CLASS(or type)_DISK_IMAGE which adds, removes packages, runs scripts
to implement `firefox.app` (App)licatin VM.] *(in the future we will 
 want to generate ICONs on the USER CONTROLLER VM to streamline usage, 
 mkaing the user epxerience nealry identical to other operating systems.
 **THis is designed to make highly secure operating system easy to use,**
 **intuitive, and slick design ideally eventuall**

 **Also we wnat to have USER CONTROLLER AGENT intercept ALL "open as.."** **so the request can be executed in a isoalted VM, like PDFs should alway** be in the special PDF/media (App)lication VM.

 
  **Nautilus is a massive attack surface, so even interaction with the FS**
 using that software is isoalted and connected back to the COntroller to 
  utilize a cluster of interconnected VMs to form a super computer, that
  is capable of HIGHLY complex routing, to provide online identity segregation, project segregation (for open source developers segregated
communciation with OSS communities) and so on.


  project
b


   



 [vm images][libguestfs][url: libguestfs.org = has a list of all sub-projects/modules that make up the libguestfs package suite. has things like extract windows registry, convert guest to kvm (v2v somehow), format, get kernel from disk,  etc] set of tools for afcessing modifying vm disk images (could be the basis for auto shrinking and growing VM hds! obivously auto back up too silly)
     * monitor disk usage
     * creating guests
     * **SCRIPTING CHANGES TO VMs** (this is what we would use to fucking auto shrink and grow!!!!)
     * backups
     * formating
     *reziing (!!)
     * Cloning
     * Building
     * V2V?
     * P2V?
     * __editing files inside Guests and viewing__


    [SUB PACKAGES] the ones that jumped out to me becuse we wnated to do things like save a version (auto vM image versioning/snapshots, KEEP IN MIND using a system where the user files are segregated from the standard OS files

    *


 [!][Show A: USB rediction over TCP: this would be great for USB proxy VMs, we use special alpine lick VM to attach ANY USB. Then we make the contents accessible via TCP to a different VM that opens files (like a filesystem] 


 [FIX CLOCKS XML]
  * kvmclock - reocommend clock source for KVM guest, virtual machines KVM pvclock or kvmclock lets guest machine reads thost pshyyscial machines wall clock time. [!!]

 [!][I said we shoudlnt combine HOST and CONTROLLERs because of potential tampering. but we should be taking and checking HASHES and storing in read-only locations. This should prevent any tampering and allow us to centralize images across ALL the clsuter (maybe even torrent to spread them out over all users of multiverse]]




[!] should portal-gun JUST be the name of the video system that handles video movements between VMs. Then possibly call the VM management `portals`, or `singularity`, or `wormholes`. 

`wormhole graphics`?

_____________
## Look into linked and chained files for higher level layers of a multiverse file system
____________
## Use chardev or parallel socket or similar to create SysFS-LIKE API for each VM.

Exampe: 
echo "1" > /vm/network/ens3/state (turns state to on/up)
/vm/network/ens3/dns

/vm/vda1/(all files)

/vm/network/hosts (add things to hosts file directly, or list host file, or replace it)

/vm/state/reboot
/vm/state/snapshot



## ADD entries to VM Hosts from from CONTROLLERS and HOST!
Sep 07 19:41:41 host dnsmasq[4364]: read /var/lib/libvirt/dnsmasqvirbr0.addnhost




### /var/multiverse
  `/var/multiverse/sockets/` - contain the sockets for connecting to libvirtd until further notice

		done via hard link

  `/var/multiverse/images/` ['os', 'templates', 'service', 'application', 'controller']
                         

  `/var/multiverse/machines` (eventually portals, or portal codes or something thematically correct) will hold VM configurations in RUBY (so they can be programtical (or yaml)

  
__

````
nobody    1225  0.0  0.0  52228   256 ?        S    01:10   0:00 /usr/sbin/dnsmasq --conf-file=/var/lib/libvirt/dnsmasq/virbr0.conf --leasefile-ro --dhcp-script=/usr/lib/libvirt/libvirt_leaseshelper
````


### multiverse files

root@host:/usr/share/qemu# ls
bamboo.dtb	  petalogix-ml605.dtb	    QEMU,tcx.bin
keymaps		  petalogix-s3adsp1800.dtb  QEMU,VGA.bin
openbios-ppc	  ppc_rom.bin		    slof.bin
openbios-sparc32  QEMU,cgthree.bin	    spapr-rtas.bin
openbios-sparc64  qemu-icon.bmp		    trace-events-all
OVMF.fd		  qemu_logo_no_text.svg


/usr/share/doc/multiverse 
   copyright (lol)
   TODO.Debian
   multiverse-doc.html
   changelog.Debian.gz




### portal gun pxe

root@host:/usr/lib/ipxe/qemu# ls
efi-e1000.rom	  efi-pcnet.rom    pxe-e1000.rom     pxe-pcnet.rom
efi-eepro100.rom  efi-rtl8139.rom  pxe-eepro100.rom  pxe-rtl8139.rom
efi-ne2k_pci.rom  efi-virtio.rom   pxe-ne2k_pci.rom  pxe-virtio.rom


### portal gun notes




  * WE MUST USE encrypted huge pages for ram. no balooning!



  * watch processes, do essentially ps aux grep looking for qemu

  * watch `/var/run/user/1000/libvirt/qemu/run` *.xml
    [make this available over sysfs like files]
    [make this info avialble over REST API] 
    [ add more info]

  <channelTargetDir path='/home/user/.config/libvirt/qemu/channel/target/domain-2-galaxy.router.multiv'/>

  **THIS HAS IF IT BOOTED! WE CAN CHECK IF IT BOOTED< SO WE CAN DO FALL BACK SO EASILY. We short poll this (eventually we setup async event driven shit but until then we short poll and check if its booted, if after X amount of time its not booted then guess what? WE FUCKING LAUNCH OUR GRUB2 backed fALlback VM**
<domstatus state='running' reason='booted' pid='2367'>
  <monitor path='/home/user/.config/libvirt/qemu/lib/domain-2-galaxy.router.multiv/monitor.sock' json='1' type='unix'/>
  <vcpus>
    <vcpu id='0' pid='2379'/>
  </vcpus>
 



### Notes/Scratch Pad

  * [scramble shell] Support notifcaitons in UI using echo "message" /user/notificaitons, and have a REST API too



  [Open source BIOS options]
    [open hackware] simple ppc bios for qemu
    [ovmf] our current BIOS of choice, supports signed boots



  * Portal gun should utilize dd and grub heavily
/boot/grub/x86_64-efi$ cat terminal.lst 
iat_keyboard: at_keyboard
iserial: serial
iserial_*: serial
oaudio: morse
ocbmemc: cbmemc
ogfxterm: gfxterm
oserial: serial
oserial_*: serial
ospkmodem: spkmodem






  [*][Configuration of libvirtd] 

````
# Start the virtlogd to gather logs from the HOST
systemctl start virtlogd.socket

# In order to connect via 'virt-manager' start the following:
systemctl status libvirtd.service


#  ├─ 810 /usr/sbin/libvirtd
#  ├─1134 /usr/sbin/dnsmasq --conf-file=/var/lib/libvirt/dnsmasq/virbr1.conf --leasefile-ro --dhcp-script=/usr/lib/libvirt/libvirt_leaseshelper
#  ├─1225 /usr/sbin/dnsmasq --conf-file=/var/lib/libvirt/dnsmasq/virbr0.conf --leasefile-ro --dhcp-script=/usr/lib/libvirt/libvirt_leaseshelper
#  └─1342 /usr/sbin/dnsmasq --conf-file=/var/lib/libvirt/dnsmasq/virbr2.conf --leasefile-ro --dhcp-script=/usr/lib/libvirt/libvirt_leaseshelper

````

Using the above, we can more quickly side-step `virt-manager` and hopefully `libvirt` as well.

### Alpha Functional Requirements
The alpha version of Multiverse OS will be released when the following functional requirements have been met, in the meantime, interested developers can manually setup Multiverse OS using the provided instructions, feel free to contact the developers or submit an issue if you run into any issues with the manual install process.

  [Scramble Shell][The DESKTOP SHELL that powers the Multiverse OS desktop environment]
  Below are the components that make up the basic structure of the **Scramble Shell** desktop environment. 

    [*][Scramble (Identity) Suite]

    [*][Scramble Terminal] A terminal environment that allows for seamless control of all the virtual machines that make up the Multiverse OS cluster. The terminal environment provides consistent interface to the entire cluster, but in a way that functionally presented as a single "super" computer. (i.e. We remove scp, and improve cp, so that copying across VMs in the cluster is as simple as copying files between hard-disks.)



  [Portal Gun][The VM MANAGEMENT system that is the foundation of the Multiverse OS cluster] 

    [*][Image/Disk Management] Multi-layer disk management, with built in integrity checking using merkle-trees, syncing using modified Bittorrent protocol, (tree based customization system).

   [*][Custom Userspace Network Stack] To avoid any use of the HOST kernel for VM-to-VM networking. This avoids the most common breakout method: special malformed packets that breakout when passing through the kernel.
  
   In addition, provides greater speed, prevents HOST level tampering with Multiverse cluster, which may eventually enable sharing of resources. 

   [*][Onion Network] 


  

### Entries


  [SIMPLE][NEWBS CAN HELP]
  [!][CUSTOMIZE BOOT VISUALs FOR MULTIVERSE OS AESTHETIC]**(Modify /boot/grub/* and then write a script/program that will automatically update a /boot/grub to Multiverse-IZE various BOOT looks)** 
	_and base it on VM type: so differnet look for APP VM, Service VM,_
	_Controller VM, and HOST)_

	[!][Upgradable items]
	    [*] unicode font (?)
	    [*] grub.cfg to modify ENTRY NAMEs
 	        __(set video basd on graphics card/VM type)__
                [IMPORTANT][NOTE] Can define kernel modules to call
		in at this point, so we can call in AMD or NVIDIA
 		and probably get video working for the TYPING in part
		of cryptsetup.
		

	       [*][still grub.cfg]
          	[MEGA IMPORTANT][NOTE] Can add new entries, for example
	        FALL BACK CONTROLLER, or just a VM that has GRUB, and
		we generate a list of the available CONTROLLERs, and
		then we boot to selected controller in GRUB menu
 
		   _[legitiamtely VERY INTERSTING idea/solution for_
		   _the fallback CONTROLLER VM]_









  [SIMPLE][NEWBS CAN HELP]
  [SECURITY ENHANCEMENT](Use the "CONTROLLER AGENT" or `portalgun` to track `/boot` integrity using merkle-tree)
  Preform cleanup
  **and sign key files so that secure boot from BIOS can be enabled**


  [SCRAMBLE (IDENTITY) SUITE] Foundation Components that are required for an ALPHA build from SCRAMBLE SUITE:

	* Base Scramble KEY used to generate derivative keys and deterministic keys.

	* PAM authentication

	* Pass-store system that supports storing account details, tags (or categories, so sorting by identites for example), OTP/TOTP support, git saving, backwards compatbile with .pass-store,...

	* Support SSH-Keys, GPG-Keys, Onion-keys, Bitcoin Keys, Ethereum Classic Keys, x509 keys, TLS keys,...

	


  [NOTE][Need to review and improve controller names] Currently the names for the two CONTROLLER VMs are: ['controller.app.multiverse', 'controller.router.multiverse']

	'controller.user.multiverse' may be a better name since it is the `user` interface to control the multiverse cluster. Additionaly, the 'controller.router' may be better described as the 'controller.network.multiverse'.

        


  [Filesystem structure][Configuration and application data] **Define locations for general configuration, specific configuration, and user specific application data**
  
    [HOST]

    [USER INTERFACE CONTROLLER]

    [ROUTER CONTROLLER]

[Host Multiverse OS Directories][Ideally we want to phase out the use of `/home/user/multiverse/` because if eventual directory transversals were possible, we would not want to doing this from the scope of the USER folder since this folder has much of the configuration and keyfiles.]


    [!][`/etc/multiverse`]__location for MULTIVERSE OS configuration files__
	(YAML files to configure things on a GLOBAL level)
	(Example):
		* default editor (vim, nano, neovim, etc)
		* default datatype output (json, xml, CBOR, etc)
		* ... think of more
		
		


    [!]['/var/multiverse/'][unique files to the user/multiverse OS install, like the images!]
	[`/usr.`]

	[`/var/multiverse/endpoints/`]__specify the avialable endpoints that can be used to provide exit/serving to the internet and what PORTS__
		[`/var/multiverse/endpoints/remote-ips`] 
	        _list of remote servers that are accessible by SSH for example,. allowing reverse SSH and taking over a specific port and forwarding to an internal cluster VM (very important)]_
		[`/var/multiverse/endpoints/onion-addresses`]
     	        _generated onion addresses for random things_
		[`/var/multiverse/endpoints/domain-names`]
	        _registerded domain names, preferably accessed via API_
	

	[`/var/multiverse/images/`]
		[`/var/multiverse/images/os/`]
	        |->[new ones can be added by specifying the ISO, the developer GPG key file, and the SHA256SUMS (or some other similar file)`][MAKE A GO LANG SCRIPT THAT WE CAN PUNCH IN THE VALUES, gnerate simple YAML configs and then we dont need 50 unique dumb bash scripts]
		[`/var/multiverse/images/templates`]
		[`/var/multiverse/images/controller`]
		[`/var/mutliverse/images/application`]
		[`/var/multiverse/images/service`]

    [!]['/usr/share/multiverse]_general multiverse os configuration files_
	['/usr/share/multiverse/DEFAULT_VM_YAML`]
	[`/usr/share/multiverse/


	[INFO]
	['usr/share` usually contain default files, or base files] 
	[EXAMPLES OF `/usr/share/*` FOLDERS]:
	  [`/usr/share/openbios/`]
	  _contains openbios-ppc, QEMu.VGA.bin, QEMu.tcx.bin, etc_
	  [`/usr/share/openssh/`]
	  _contains default 'sshd_config' to source from when initializing files during a new installation._
	  [`/user/share/git-core/`]
	  _literally as a "templates" folder_


  [Concept][Function/Feature] **New type of link, that copies data across drives, and keeps them synced**
  Any changes in either updates the other, this enables simple duplicate management for sensitive files
  that require multiple copies to be kept across several HDs. Like keyspairs.


  [Core][Design/Structure] **Migrate away from using virt-manager**
  
  This can be successfully acocmplished by providing basic `portal-gun` functionality. 

    [*] List VMs (track processes, be capable of checking equivilent of `ps aux | grep qemu` so qemu virtual machines are not started twice (which is currently possible due to virt-manager bugs). 

    [*] Create `libvirtd` VMs (QEMU XML for network, storage, VMs). 





  [Core][Design/Structure] **Encapsulate all the ROUTER VMs in ROUTER CONTROLLER VM**

  This new design encapsualtes all Multiverse OS virtual machines in TWO CONTROLLER VMs, that provide all the VMs in a nested environment. Each CONTROLLER VM is encrypted, the RAM is encrypted, (eventually the networking will bypass the HOST kernel), and other efforts will be made to limit any possible damage to CONTROLLER VMs from the HOST.
  
  HOST should not even be able to shut down the CONTROLLER VM, requests must be signed with the CONTROLLER VM session-key. This is a security-in-depth approach that is intended to defend primarily against hardware backdoors, physical capture, and similar attacks.  

  Additionally, a ROUTER VM, could be just a special SERVICE VMs to simplify Multiverse VM classes. 
  

  [Development][Experiment] **Access HOST from CONTROLLER VM `libvirtd` without requiring full SSH access or custom software** 

    By using p9 share, and symbolic linking, added the unix sockets ot access the HOST `libivrtd`. This should allow slightly more secure access to assist in development. 


   
