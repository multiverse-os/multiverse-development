## Multiverse OS: Application Brainstorming
Multiverse OS is not just a protocol for secure operating system design using
hardware virtualization based isolation and management of multiple internet
identities in a psuedoanonymous manner. It is also collection of secure, 
intuitive tools which make up a a complex creative workspace for engineering,
a wide variety of art, and more.

The basic design helps one securely manage
multiple internet idenities and organizes each project associated with
an identity. 

Below are brainstorming ideas for the collection of software that will
be maintained and included with Multiverse OS by default. 

   * **Text editor** A gedit type editor, it shouldnt mimic vim or text editing, it should be GUI, advanced, tec. 
    * [Functionality] pasting in a link, auto grabs meta data and foramts it into a link and the link formats into markdown/html based on predefined templates and modifyable as needed. saving a bunch of useful links is just pasting into the file and it opragnzes itself. then you can sort all x items like links

  * **Multi-layer user database** Multiverse DB - layered DB, with KV, graphDB and reletional table DB>

  * **Terminal**

  * **VM Manager**

  * **IM and IRC client**

  * **Torrent client and tracker**

  * ...


===============================================================================
## Notes On Multiverse OS Image Creation
Below are a collection of random notes relating to creating installation
media for Multiverse OS. There are a variety of methods of producing 
installation media under Debian with the official way being utilizing 
the tool `debootstrap` which greatly simplifies the procses. 

However there are other ways that are worth understanding before developing
software to automate the procses. This is important because creation of
installation media and live CDs is important and it should be so easy that
creation of installation and live images should happen in the background
make distributing the software securely and allow for the creation of
custom live cd. 

**Method: QCOW2**
`qemu-img create -f qcow2 multiverse-os/images/controller.multiverse.qcow2 60G`

**Method: mkisofs**
The `mkisofs` is one way of creating installation or live media. The
dependencies are:

`ethdetect`
`netcfg`
`pcmciautils-udeb`
`xorriso`

Then after this software is downloaded the following command can be used (and
modified) to create installation media based on an existing disk:

`xorriso -as mkisofs -r -checksum_algorithm_iso md5,sha1,sha256,sha512 -V 'Debian 9.4.0 amd64 n' -o /srv/cdbuilder.debian.org/dst/deb-cd/out/2amd64/debian-9.4.0-amd64-NETINST-1.iso -jigdo-jigdo /srv/cdbuilder.debian.org/dst/deb-cd/out/2amd64/debian-9.4.0-amd64-NETINST-1.jigdo -jigdo-template /srv/cdbuilder.debian.org/dst/deb-cd/out/2amd64/debian-9.4.0-amd64-NETINST-1.template -jigdo-map Debian=/srv/cdbuilder.debian.org/src/ftp/debian/ -jigdo-exclude boot1 -md5-list /srv/cdbuilder.debian.org/src/deb-cd/tmp/2amd64/stretch/md5-check -jigdo-min-file-size 1024 -jigdo-exclude 'README*' -jigdo-exclude /doc/ -jigdo-exclude /md5sum.txt -jigdo-exclude /.disk/ -jigdo-exclude /pics/ -jigdo-exclude 'Release*' -jigdo-exclude 'Packages*' -jigdo-exclude 'Sources*' -J -J -joliet-long -cache-inodes -isohybrid-mbr syslinux/usr/lib/ISOLINUX/isohdpfx.bin -b isolinux/isolinux.bin -c isolinux/boot.cat -boot-load-size 4 -boot-info-table -no-emul-boot -eltorito-alt-boot -e boot/grub/efi.img -no-emul-boot -isohybrid-gpt-basdat -isohybrid-apm-hfsplus boot1 CD1`


===============================================================================
## Multiverse OS: The Controller VM, the nucleus of the cluster

**[Q] What software is on the Controller VM if apps are isolated in VMs?**


  **[A]** The primary role of the Controller VM is to provide a safe 
  computing environment, acheived by proxying dangers away from
  the computer that holds sensitive data that makes up the "memories"
  or data that gives "identity". 

  
  The primary purpose of the Controller VM is to run the intuitive
  modern and reimagined Linux desktop envirnoment, bringing 
  bleeding edge "cloud"/virtaul-machine technology, packaged
  in a friendly, and seamless way, that provides the power
  of super computers to the aver user.

===============================================================================
## Key Multiverse OS Functionality
[1][Epheraml virtual machines to isolate and segregate identities]

[2][Scramble Shell: A modern desktop environment that seamlessly]
*Connects any number of computers topgether, unifying the experience*
into a single desktop experience. For example, `scp` is removed
from Scramble Shell, and `cp` is no longer just for copying
between a local disk or between disks in a computer.

It now just as seamlessly copies between disks on different computers,
folders, and so on. In addition, a command, can be sent to a class
of computers.

A predefined group, single computer, computers on x amount of time,
and so on, providing control over massive amounts of computers
by streamlining, updating and rethinking the Linux CLI user
experience, making it both incredibly powerful, more decentralized, 
but also simpler and easie to understand.

[3][Scramble Suite: Scramble Suite is a collection of identity software]
*that supports multiple online identites segreated by virtual machines*
and isolated. Providing tools to help keep them separate and simplify
complex workflow for softwawre devleopers working under multiple pen-names
and with multiple communities.

**Mycelial Key System** growing up from a single key, a complex branching
key system that supports entire branch revocation, ephemeral session keys,
safe backup and peer based restoration.

Support for multiple key types, verification, multiple identites and so
on, in a single key system. Visualize with sunburst graph


===============================================================================
## Important development/documentation tips/research/notes
* **When saving previously executed commands use the following:**
  Use `history -a; tail -5 $HISTFILE` to print out previous lines of text in
  a way that it removes the prefixing prompt (i.e. user@host:~$).

* **ICE Authority** inter-process communication.

* **Look at /TidOS** it has as some cool concepts.

* **Explicit sync** driver for user space DRM for epxlcitly syncing video
  buffers


===============================================================================
## Keys / Secrets / Keyrings
  [*] Keys for QEMU VMs **Libvirt files need to be merged in here. Should automatically setup files in `/var/multiverse`* structure. For now just using XML files dropped in. Later generate XML or whatever else config needed from DB.** _Look into what secrets is used for, and how we can tie that into multierse*


===============================================================================
## Data Structures
  [*] Tree Data Structure (or mycelia data structure) 
      * Scramble Key Tree system
      * Tree based VM templating
      * Tree based desktop organization 
      * Tree based project organziation (actually may be more complex grraph)
      * Tree based notes organization (actually may be graph)


===============================================================================
## Real-time clock simulation
  * Build a system so Multiverse OS will be reatl-time when used with the right hardware


===============================================================================
## Gnome terminal server; what gnome-terminal connect to...
Knowing now that this is the sturcture, we can use this as the model for a Multiverse OS cluster based terminal server. 
Scramble Shell

So the controller VM runs the terminal shell server. Then each VM connects to it, provides it with the stdout, stdin, stderr, etc.
Preferably by DMA (direct memory access) or something just as cool andf ast. Maybe VirtIO.



===============================================================================
## Move all loose .dot-files into ~/.config
This is the new model for HOME folder orgnaization

and we should force it. By just organizing the data, and

doing so automatically. It can be run once at the start, and
ran again when needed or scheduled.



===============================================================================
## Multiverse `pass-store`
Compatible, use the vbasic structure and git compataible

make the CLI compatible. should be able to drop pass-store fiels
into the mutliverse pass store and jsut will be able to add
more data


===============================================================================
## Multiverse / And all future web applicaitons 
Ideally, should utilize **Shared Desktop Ontologies**. But this is just a badly parsed text file
, and this is stupid. Multiverse puts things in databases, so data is always consistent
easy to access via scripts and programming. And easy shared between software.

  What these do is supply open source describing data in a consistent and open
  source way. FOr example:

  * NCAL nepomuk calednar ontology - provide vocabulary describing calendaring data (events
    tasks, jjournal entries) 

  * NCO - Contact Ontology - inspired by Vcard. 

  * NUAO - User Action ontoology - actions availble to user on desktop

  * EXIF

  * Personal information 

  * information elemtn

  * sharing
 
  * message

  * annotation

  * file

  * multimedia
   
  * download




===============================================================================
<p align="center"><img src="https://github.com/hackwave/multiverse-development/blob/master/multiverse-logo.png" height="300" width="300"></p>


## Multiverse OS Development
Below is a list of features and functionality which would immediately improve the functionality of Multiverse OS, and is deliberately avoiding being a wishlist of complex features that would not be included in the alpha installer. Features that are more complex and would be included after the alpha release will be documented elsehwere, most liklely within the machines folder `README.md` for each component VM. 


### Alpha Release
The alpha release of Multiverse OS revolves around releasing a basic installer that can be used by novice users to test the basic features and functionality of Multiverse OS.

It will not incorporate all security enhancements and features envisioned for the first initial release of Multiverse OS but instead focus on the most basic functionality which qualifies a Multiverse OS machine:

  * Specifically operating within a **Controller VM** that has all USB PCI controllers, GPU and pinned CPUs for near bare-metal preformance that is capable ofplaying video games but maintaining the security focus, and ephemerality of an operating system like Whonix or Tales.
  * Locking down the **Host Machine** and preventing any user iteraction.
  * Segregating all applications to **Application VMs**, seamlessly passing windows to the **Controller VM** using `xpra` and if possible Multiverse OS replacement that uses shared memory (DMA). 
  * Networking the system using specialized **Router VMs** that use different security enhanced operating systems and PCI passthrough of network interfaces.
  * Support **Service VMs** that can exist deeply in the Multiverse OS networking layers, or near the edges to be accessible to the local LAN and isolate long running services like Torrents, LAN file sharing, Backups or security focused features like automatically removing meta data from all images in storage. 

Read the documentation for other planned features that will be added after the initial alpha release and be the focus of the beta release. 


#### Alpha Development Features
Below are list of features that could be implemented immediately to improve functionality and would not requrie a lot of developpment resources so that it is possible to implement them before the alpha installer is complete. 

  * **Divide this 'multiverse-development' into [`/var/multiverse`, `/usr/lib/multiverse/`, `/etc/skel`, `~/.local/share/multiverse`]**

  * **Generate Scramble Shell Key Tree and all associated keys**

  * **Launch a basic version of OHT** to network between **Host Machines** to form clusters

  * **Router VMs** run their own DNS servers to block ad-servers and known malware and map hostnames to Multiverse VMs (following the naming structure setup and connecting satallite installs to connect multiple **Host Machines** together to create super computers controllable from a single Multiverse OS Controller VM)

  * **Add default templates for all Multiverse OS configurations to /etc/skel/**

  * **Host Machine Agent/Daemon** Using the daemon skeleton included in this repository as the base structure, a daemon will be created to manage the following:
    * PCI Device Preparation (look at vfio-bind and other related scripts as outlines for required functionality)
    * VM Startup, checking if the **Router VMs** connect before starting the next **Router VM** in the sequence, then finally turning on the default **Controller VM**
    * Provide a limited API for control over Routers (On, Off, Reset, Status [network, cpu, i/o, memory]), and controllers
    * Provide a limited API for reset, shutdown of **Host Machine**
    * Provide API needed to create a menu to select and start **Controller VMs**, switch **Controller VMs**, reset, power off 
    * Failsafe system to load a very basic/minimalist **Controller VM** that can be used to turn on other **Controller VMs** or diagnosing issues/debugging if the default **Controller VM** fails to launch after `x` amount of time.
      * Key sequence to return/rebind USB devices to the host machine for maintainence [Require a maintainece password set at installation] 
    * Generate and maintain:
      * `/var/multiverse`
      * `/usr/lib/multiverse`
      * `/etc/multiverse`
      * `~/.local/share/multiverse`
      * `~/.config/multiverse`
    * Lockdown functionality
      * Watch for new logins and alert/shutdown
      * Watch for new connections alert/shutdown
      * ...similar notification systems for other similar systems...

  * **Router VM Daemon/Agent** (Add details when possible)

  * **Controller VM Daemon/Agent** (Add details when possible)

  * **Application VM Daemon/Agent** (Add details when possible)

  * **Service VM Daemon/Agent** (Add details when possible)

  * **Host Machine Update System** Setup an `aptly` (included in the go-libs folder in this repository) based repository **Service VM** that can provide seamless but segregated/isolated access to limited updates and applications to the **Host Machine**. This would be connected in an isolated network that is not accessibleto the internet when it is acessible to the **Host Machine**. This **Service VM** would be created when updates are required, functioning as an epehemeral Live CD and provide repositories to the **Host Machine** in an airgapped way. This would occur without much or any user interaction preferably beyond accepting/denying update notifications. 
    * If we decide to cut off internet access to **Controller VM** which is very likley once `portal-gun` matures and we shrink down the resources and size of the kernel wwe can use this same system to upgradethe **Controller VM**


  * **CLI Virtual Machine tools** with the goal of replacement `virsh` with tools that are focused around Multiverse OS VM functionlaity; including but not limited to listing Multiverse VMs by type, focusing/providing ONLY unprividleged VM access and removing all root VM functionality, adding PCI devices to VMs or Multiverse OS VM types [Controller, Router, Service, Application] including the steps that involve unbinding the PCI devices and not just assignment in XML like libvirt, starting VMs, stopping VMs, checking network connectivity using APIs supplied by daemon/agent, automatic CPU pinning for controllers, automatic L3 cache enabling, incorporating suspend to memory hacks to bypass GPU reset limitations, 


===============================================================================
## Potential Alpha Features
The following features are likely too ambitious to be in the alpha release and will likely be the core features that define the beta release.

  * **Use `u-boot` and custom BIOS built using scripts included Multiverse OS to sign the boot process and ensure that there is no tampering from the BIOS up to the OS for each VM.**

  * **Build out functionality to ensure initramfs can not access the internet, add checksums/validity checking, to ensure there is no tampering with Multiverse OS initramfs. Do not make it easy after initial setup to rebuild the initramfs without requiring signature from the initial installer of Multiverse OS. Initramfs malware is very effective and there is currently very much to protect against it. Multiverse OS needs to focus on this attack vector and do whatever we can to limit this attack surface on the Host Machine and every VM in Multiverse cluster.**

  * **Custom Network TCP/IP Stack** to bypass all kernelspace networking since this has been the attack surface almostall hypervisor breakouts have occured in. Using shared memory (DMA), custom devices that provide direct device to device routing to **Router VMs**

  * **Custom xpra replacement for seamless window sharing from Application VMs to Controller VMs** using shared memory (DMA), to avoid providing the control xpra provides, limiting the attack surface and speeding up the entire system and avoiding using python and instead opting for Rust langauge implementation using Vulkan and Wayland.

  * **A GUI interface and Gnome intergration (or replacement)** to control various aspects of the Multiverse OS system, abstracting the features into a single interface making the cluster of machines (and possible cluster of several **Host Machines** forming a super computer) function as a single computer to the user. 

  * **Start switching out GNU core-utils** For replacements that match the Multiverse OS design philosophies, and make all the commands support output in data formats that are easier to program against (JSON, YAML, XML, ...) and *most importantly make all the flags consistent, so it is easier for novice users to learn to use command-line. but also ensure EVERY command is backwards compatible to existing commands so that advanced users can continue using the commands and patterns they already learned while making it easier for new linux users to learn the new consistent patterns.**

    * Support output in whatever format, REST API and file based API for easier scripting and programming

    * Configuration that supports regex based coloring output to allow simple theming 


