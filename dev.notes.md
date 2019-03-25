# General Multiverse OS Development Notes & Research
===============================================================================
 * A vim hotkey to insert 80 characters of "=" or "-" would be lovely
===============================================================================

 * **Create our own shell interpreter that provides coloring using regex** in
   addition to possibly additional features or at least better errors so 
   debugging is a lot easier. This would make sh much more accessible to novice
   users



===============================================================================

 * **Should we use the short hand 'users' to describe linux 'user accounts'?** I
  have heard at least understandable arguments against this terminology. Even
  just account has less namespace crossover with other arguably negative
  meaning.

  Multiverse OS design specifications already layout a plan to introduce a
  "shim" software to provide an overlay virtual FS with the goal of simplifying
  the organizational structure of the file system and the naming of individual
  files.

  In addition, to making the filesystem easier to navigate and understand
  without diving in large tombs of sometimes esoteric and reglarly boring
  help pages, new users can not only avoid being scared away from linux 
  but hopefully learn about it, begin making customizations, and be 
  introduced to the power of linux more rapidly, so that they are likely
  to share it with their friends and family. 

  So, changing things like 'user' to 'account' visually to the user
  would be incredibly simple. 

  The shim will also have an affect on the core-utilities; because even some
  of us Linux old-timers remember being confused at the various core-utils
  having entirely different user interfaces, making it difficult to remember
  which one required a capital 'R' for recursive mode and which one required
  a lowercawse 'r'. By providing a shim that offers consistent naming across
  the core-utils (in combination with many other planned changes to improve 
  the Linux user experience by modernizing it, and being willing to step
  back and find new ways not yet attempted to make Linux better, realizing
  we are only being held back by limitations that were a result of hardware,
  software, financial, and personal limitations that simply do not exist
  anymore and can be safely abandoned in favor of an udpate that would
  make Linux more accessible, better suited to be translated, providing
  the tools to translate it from the kernel up.

  **[And don't worry, the overlays will be provided alongside the existing 
   system to for as long as reaoonably possible.]**



===============================================================================
* Huffman tree?
ii


=-=====
## relevant tools

creiht/ring - cinsistent ring hash!
gholt/devicering - similar

## VMs

scandia-minimega/minimega

u-root/u-root
  * MAKE INITRAMFS here


---
**Harvey-OS**/ninep - old project being forked and maintained by
new os project.

imnplements BOTH client and server for the 9p and 9p2000 distrubted
resource prtocol in Go.

**What about the idea of impelementing this over a VIRTUAL PCI DEVICE**?


====
## containers/clearcontaienrs


containers/virtcontainers
----
# userspace network driver: ixy (in go)

ixy-langauge/ixy.go

**mpleso/vnet**

====
# distributed object store (for nvme)
daos-stack/daos

====
## custom immutable kernels using Go and patches

linuxkit/linuxkit :#

ednaganon/linuxkittry

====
## COMPRESSION



---
## general
Obviously include archive GO library in 'os lib'

* CHECK OUT ZSTANDARD, its the new best real-time compsession!!! What if we built this into bittorrent???????????

##### embedded zerotrees ezc
lossless image compression

also exists a file foramt called enhanced compression wavelet

 
#####  Snappy is fast data compression, LZ77 

used in leveldb mariadb, hadopop, mongodb, rocksdb


good for compressed streams of data because of speeexd

PORTALBLE!!!
compression ratio 20-100% lower than gzip.

**IS there a parallel version like with ParallelGZIP in Go?**

compression speed is 250MBs compressing and 500MBs decompressing

##### Z STANDARD - THE NEW BEST LOSSLESS COMPRESSION!!!!
**Used like tar.**GZ** to compress single files!!!!**

Its in the linusx kenrel 4.14

Its used in btrfs and squashfs 
**CAN WE GET IT IN QCOW2? or maybe QCOW2.5?**

Zstandard is now the default comperssion for deb packages under Ubuntu!!!


HTTP content encoindn g can now be zstd as of 2018, and media is "application/zstd"


RocksDB uses it!

file extenmsion is "\*.zst"

===============================================================================
/var/ subdirectory is for linux and other unix-like processes and the system to write data DURING the course of operation. For example, look at:

```
	/var/run/* 
```

This folder contains all running processes PIDs, memory, etc. Super important. So originally we were going to use /var/multiverse properly by only putting active VM data in there, active agents/daemon data, session keys ETC. (THIS IS CORRECT)

Currently thought it contains EVERTYHING! WRONG!!!! S TOP IT



===============

The design documentation defines coding practices, rules
and guidelines to set a minimum requirement, ensure consistency
between an extremely large project, such as an operating system,
focusing on producing source code that can be worked on by
large group of friends, aquaintences and strangers, by producing
source code that is understandable without additional documentation,
by priortizing requirements and practices inspired by successful
past projects but customized to the needs of the Multiverse OS
developers and this specific project.

Arguably the most important document in the design documentation,
is the Multiverse specification, which outlines the entire
project on a full-view scale, in addition the specific design
of each of the major primary components, the protocols they
will implement. This documentation is incredibly technical and
each sub-project specification provides even more in-depth
design specifications that will eventually be replaced by
the specification implementation in Go or Rust depending
on the project.


In addition to the technical design and specification, the
overall Multiverse design specification incldues the minimum
design requirements that every project planned to be included
in the Multiverse OS project must meet for inclusion in the
operating system and released in an official release as a
default piece of software.

These rules and guidelines specifically require the developers
use descriptive and complete variable, and function names; 
The specifics minimum requirements, and coding practices can be
found in the design documentation; but in essence, heavy emphasis
is put on clarity, general intuitiveness, with the goal being
simple APIs, and maximizing human readability.






================
Multiverse OS development currently is focused on the target of Webframe
development for building community nuclueating point at multiverse-os.org
and using the Webframe framework in combination with webview/webcomponent
to build the GUI for the alpha installer. The alpha installer is the highest
priority goal, that requires the removal of the `libvirt` dependencies in
Multiverse OS in favor of direct QEMU/NEMU interaction via command-line using
YAML (and eventually Ruby) based configuration for both the virtual machine
and provisioning of the virtual machine to the ephemeral base for all
virtual machines using a given image. 

Images are built ontop of a system of merkle checksum verified tree of check-
points that simplify deployment of images by starting with the vanilla install
media of popular Linux distributions and provisioned into several checkpoints
that service VMs (routers, long-running background services), controller VM
(user interface VMs), and app VMs (encapsulated ephemeral application VMs) from
one of the base checkpoints in the tree of versions starting from the vanilla
installation media.

This allows for trustless images that are built locally, built in a repro-
ducible way outside the boot sector which is built using secure open source
BIOS signed with the users scramble key, and chained with signed Bootloader,
intiramfs and eventual Multiverse OS kernel for maximum security.


===============================================================================
## Development Roadmap for Multiverse OS Alpha Release
The alpha release major milestone is a GUI installer and installation media
that will enable access to Multiverse OS beyond developers capable of manual
installation and troubleshooting often required to install Multiverse OS
manually. It is an important step as it is the first time we are able to
share Multiverse OS and open up access to the general public. 

**Alpha Release Candidate: Functional Requirements**
Below are the functional requirements that define the roadmap to the Alpha RC
milestone. 

  1) **Remove 'libvirt' dependency**
     Remove 'libvirt' dependencies from Multiverse OS:
       * Implement QEMU/NEMU through command-line flags controlled by YAML
         (and eventually Ruby) configuration to replace `libvirt` XML.

       * Remove the `libvirt` group and all usage of `libvirt` for virtual
         machine related user priviledges.

       * Use VSOCK VM-to-VM networking or userspace based Go software based
         switch/bridge that DOES NOT connect to the host machine at all.  


  2) **PCI Device Passthrough via VFIO/IOMMU**
     Implement PCI device management as a Go library. Instead of providing
     a terrible UI and general interface of libvirt where you simply error
     out with obscure errors when trying to add a PCI device to a virtual
     machine; simply unbind the device, rebind it to VFIO when an attempt
     to add the device, so that it proceeds sucessfully. If iommu is not
     enabled or the `vfio` and `vfio-pci` modules are missing, `modprobe`
     the two items and add them to:
       `/etc/modules`
       `/etc/modprobe.d/multiverse.conf`
     If a reboot is required for adding {architecture}_iommu to grub and
     rebuilding it; then warn the user and ask if reboot okay, and proceed
     according to user input. 


  3) **Build a virtual keyboard as prototype of open-source hardware**
     For use with `portal-gun` Multiverse OS virtual machine (portal) 
     provisoning, management, app VM installer and controller. 


  4) 

  5)


  6)


  7)


  8)


  9)


 10) 

### Optional/Possible Alpha RC Features
If these features manage to be completed before Alpha RC they will be
included otherwise they will be pushed into the next milestone (Beta RC). 

 1) Software Router provide complex chaining without requiring several fully
    virtualized with dedicated CPUs.

 2) User database/knowledge-base (three layer database, K/V, Document
    storage, Graph database

 3) User key/value store interface that is controllable with YAML config
    to define commands, data stored, size, and so on. So for example, 
    aliases can be stored, defintion of words (foriegn words, or programming
    defintions), snippets, etc.

 4) Home folder organized and backed up with merkle tree & torrent

 5) Scramble suit keyring
    * Session generation
    * HD generation of PGP, SSH, etc...
    * Tree system 

===============================================================================
## Minor/Small Ideas
The idea is small, simple, but unique ideas that can be added to the Multiverse
OS alpha release candidate (RC). 

  * **Group Chat** chat so software, email, SMS, video, audio,
    image thread, etc. Tied together, encrytped, ephemeral to a
    determined time, and per message.
    
    Software to interact with your friend list, and all
    interactions will be referred to as F2F, which provides you
    with a different feature set than P2P, because friend implies
    a given default level of trust.




