# Compiled Multiverse Install Notes


## 9pfs
https://github.com/halfwit/9fs

Atomic boolean
https://github.com/foursquare/fsgo/tree/master/concurrent/atomicbool

**embeddable openvpn**
https://github.com/glacjay/govpn
**map api to fs**
With everything mapped as a Rest API, we can remap it to a FS style to model the planet 9 system by using a library like:
https://github.com/gophergala/api-fs
apifs enables users to interact with REST APIs (and any HTTP resource) as a file system.


**redis fs**
https://github.com/MatthiasWinkelmann/redis-fs
could be combined with a pure go redis server implementation. this could be a great way to replace stupid tools like dropbox.
# Linux
## Custom PAM authorization
Custom Multiverse OS authorization, using keys and ephemeral keys, and eventually login into any Multiveres OS machine to access your machine that lives in the p2p decentralized cloud. 

*pam*
https://github.com/zro/pam - PAM interaction for Go. Currently consists of functionality to implement your own module for PAM to load/communicate with for each service.

*golang-pam-auth*
https://github.com/AmandaCameron/golang-pam-auth - This is a Go implementation of a simple PAM authentication module. It returns PAM_SUCCESS for "test" and PAM_USER_UNKNOWN for everyone else.

*pam api*
https://github.com/gogap/pam

example: https://github.com/kovetskiy/eyed - eye special PAM 
example: https://github.com/otms61/pam_line-otp OTP 
example: https://github.com/albertito/remoteu2f
example: https://github.com/krockot/gopam api examples
example: oauth2 to login pam_exec
example: https://github.com/mobilizingcs/pam_ohmage

sqrl - look up the login software from the security now guy

https://github.com/mhausenblas/dnpipes - 
Distributed Named Pipes (or: dnpipes) are essentially a distributed version of Unix named pipes comparable to, for example, SQS in AWS or the Service Bus in Azure.

*streaming* https://github.com/cespare/window

 run ssh, https, and openvpn on the same port 
https://github.com/shawnl/multiplexd
Dynamic Nat-Traversing VPN using Go and OpenVPN 
https://github.com/Pursuit92/tvpn

## 9pfs
https://github.com/halfwit/9fs

Atomic boolean
https://github.com/foursquare/fsgo/tree/master/concurrent/atomicbool

** IDK WHERE THIS GOES **
Need to install OVMF: `apt-get install ovmf`
And then when making the VM, specify (in virt-manager new VM dialog, dunno virt-install steps):
New VM:
Connection: QEMU/KVM User session
Local install media

Memory: a lot, I did 8192 MiB, a third of my machine's memory, do what you feel

Enable storage
This is the VM all the other VMs basically run in, so could be good to be fairly large, but always use a disk image (not a whole physical hard drive) because otherwise there's the chance of a VM breakout by an attacker formatting the hard drive so their code can be run on next host reboot. Remember too that further disks can be mounted if the initial disk image gets cramped. Also think about if you are going to want multiple controller vms (want one for gaming? an OSX mv?). I went for 200 GB out of 1 TB


Check "Customize configuration before install"
Stuck with Usermode networking, will change later.

In Overview, set Firmware to "UEFI x86_64: /usr/share/OVMF/OVMF_CODE.fd"
Other customizations (talking about a controller aka all the hardware vm):
Boot Options: check Autostart

Remove unused hardware/input: for example, Tablet, Display Spice, Sound (will pass whole card through later), Console, Channel qemu-ga, Channel spice, Video qxl (will pass whole card), 

Add Hardware:
NVIDIA graphics card bus id
NVIDIA audio bus id
?Intel PCI Express Root Port (the one associated with the video card, see below somewhere for the command to get the IOMMU groups, search 'iommu_groups')
  * Had this as a device, but it was permissssion erroring me. 
  * All three of these devices are sent to /sys/bus/pci/drivers/vfio-pci/new-id using /etc/rc.local or modprobe options


### USB PASSTHROUGH
Put the PCI ids in the /etc/modprobe.d/vfio.conf
In virt-manager, add hardware, choose the IDs. Delete the virtual USB connections.

### SHARED FILE PLACE
virt-manager: Add hardware: Filesystem
Type: mount
Driver: Path
Mode: Mapped
Write Policy: Default
Source path: <enter path here or browse to it: Browse; Browse local>
Target path: multiverse

### SOME NETWORKS
virt-manager, as root: Edit; Connection Details: Virtual Networks
delete default, if it's still there
Add new network:
Name: virbr1
Enable IPv4
Network: 10.2.2.0/24
Enable DHCPv4
Start: 10.2.2.2
End: 10.2.2.254
Isolated

actulaly did a bunch of stuff manually for virbr0

_Everett bridge__
everett0
10.255.255.0/24
(and basically everything else the same)

`virsh net-edit everett0`
deleted the dhcp if you made it earlier

delete the dhcp and ip block from virbr1 and virbr2

this is where you'd delete all the networking frmo the host computer too.

#### IPTABLES
https://unix.stackexchange.com/a/145933
All parts of that answer that say YourIP, put in 10.255.255.100
Also some stuff to drop outgoing ssh and other outgoing connections


#### SSH KEYS
Make ssh keys and move to persistence folder
`ssh-keygen -C user@host -t ed25519 -f /home/user/.local/multiverse-os/portal-gun/persistence/id_ed25519`
Add public key to user's authorized_keys

```
mkdir /home/user/.ssh
cat /home/user/.local/multiverse-os/portal-gun/persistence/id_ed25519.pub >> /home/user/.ssh/authorized_keys
```

>>>>>>>>>> TO DO NEXT >>>>>>>>>>>
Setup router vms

Reboot

Install debian

Need to load iptables at boot (once sure the other networking works)



**Notes**

## NETWORKING
**FIXING METAPHROS**
I'm done with this uni0, gal0, sol0 bullshit. It made no sense, I was trying to hammer this metaphor in and it didn't work with Multiverse metaphor and it was harder. 

Instead your home universe, C137 (if you want to reference rick and morty), is the controller VM. This "controller" term I used turns out is a common term for this type of thing so that stays. But each VM inside of it is a virtual universe, every VM on the same level is a parallel universe. As in a different baremetal machine, each baremetal machine in the cluster is a unvierse, with pocket or virtual universes inside of it. 

So that means the main routers outside of the controller currently named virbr0, virbr1, virbr2, and virbr3 (development bridge) are better termed "Everett Bridges" so its Everett0, Everett1, Everett2, Everett3  and so on. This actually makes a lot more sense in the context of the metaphor since the Everett bridgem is a term used when speculating that wormholes may be a dimensional tear and let you travel between dimensions and since one of these everett bridges is where the interclsuter networking would occur, this makes perfect since. INFACT the interclustering networking should probably just be a router hook to everett1 just like the controller VM is and be on the same 10.2.2.* network. 

**NAMIMG**
So this makes the naming even better and more descriptive:

router.wan-firewall.universe0 (Any LAN and therefore WAN exposed device is numbered, for each device exposed this way.)
router.vpn-firewall.universe0 (Provides the VPN connection and firewall)
router.tor-firewall.universe0

**NEW FINDING**
When reworking my networking to fix some issues I found an even better way to do the networking. Instead of doing isolated networks and being FORCED to use the fucking host as a connected device. One can simply attach each network to the previous network by *Forwarding NAT to virbr1-nic*. This will give the network an ip address on the other network and conennects its network to the other one.



------------------

--- linux manage------------

=== crons
https://github.com/michaloo/go-cron - Simple golang wrapper over github.com/robfig/cron and os/exec as a cron replacement. its a cli but can be repurposed for lib

=== open
https://github.com/skratchdot/open-golang -  Open a file, directory, or URI using the OS's default application for that object type. Optionally, you can specify an application to use.

-- UI -----------------------------------------------------------------------------
https://github.com/therecipe/qt -- Qt binding for Go (Golang) which supports Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi

https://github.com/grd/go-ipc-example - Go lacks premium-grade gui functionality. The idea is to make a "buddy" gui app, written in any language of choice, and use json-rpc for communication with the app using stdin/stdout pipes.


(QUICK mockingup mockups) https://github.com/alexflint/gallium  - Write desktop applications in Go, HTML, Javascript, and CSS.

Gallium is a Go library for managing windows, menus, dock icons, and desktop notifications. Each window contains a webview component, in which you code your UI in HTML. Under the hood, the webview is running Chromium.

-- monitoring ---------------------------------------------------------------------
https://github.com/mayhewj/jtop - A process monitor for Linux, mostly processes, basic user info, could be a good starting point

-- user management --------------------------------------------------
https://github.com/joonakannisto/determin-ed - (FUCK YES)  Create deterministic ed25519 keys from seedfile and password for openssh-key-v1 format

https://github.com/max107/gouser/blob/master/user.go -- uses the os.exec command, basic, doesnt seem to have group management

-- examples/starting points --------------------------------------------------
https://github.com/fathomdb/applyd - basic controls over a linux debian server, its 3 years old and not documented but it may ahve some things to scavenge


##########################################################################################

https://github.com/tweekmonster/luser - Drop-in replacement for os/user in Go
##################
crypto manager
https://github.com/zaibon/cryptogo - very basic starter webui for different cryptocurrencies

#################################
https://github.com/malnick/cryptorious - encrypt, decrypt passwords or random strings with ssh keys, CLI
(would be cool to add git, make it more like pass, and register different backends for backing up. so it replicates offsite.
##################
https://github.com/kiasaki/hotomata - (interesting) guess its like an ansible but all in golang, very interesting
####################
https://github.com/marksheahan/nssh - (VERY FUCKING INTERESTING) does ssh over multiple hops, could be used to obstuficate or even NAT break.
Golang command line ssh utility for running programs remotely over multiple hops
###########################

### KVM/QEMU management
* guest-agent -- https://github.com/0xef53/phoenix-guest-agent
— starting point for vm management software
-- Tools for secure (host <-> virtual machine) connections -------------------------
------------------------------------------------------------------------------------
* nush -- https://github.com/songgao/nush
— Nush can be used to built a custom POSIX shell, this shell can be used to limit
shell access when SSHing between hosts, VMs within the Multiverse OS.
-- KVM/QEMU management UI ----------------------------------------------------------
* Icemenu -- https://github.com/martinlebeda/IceMenu
- This need to be investigated further, but this may be useful for constructing an application
menu of registered applications from each VM.
* kvmtop -- https://github.com/cha87de/kvmtop
- This may be useful to look at the code or provide a very simple view, but this functionality
may already be covered by the primary tool that is being built. but there also does not seem to
be any harm in providing additional small tools.
* vm-manager -- https://github.com/ZeroPage/vm-manager
A very basic management tool that allows VMs to be defined as *.yaml configuration files and
start the machine. The *.yaml code may prove to be useful in simplyifying the process while
providing advanced users an alternative to the eventual GUI tool.
https://github.com/zchee/go-qcow2/blob/master/qcow2.go - best qcow2 lib
— qmeu
https://github.com/digitalocean/go-qemu
https://github.com/quadrifoglio/go-qmp
https://github.com/0xef53/qmp-shell
https://github.com/0xef53/go-qmp

#### File systems
https://github.com/taterbase/git-mount - git mount

https://github.com/EthanG78/go_gui — no js seems cool , just html css
https://github.com/x3ro/websocket-tty

#### VMS

https://github.com/google/novm


— wayland
https://github.com/stanluk/wayland-client
https://github.com/sternix/wl


— full kits TUI
https://github.com/VladimirMarkelov/clui - wow serious shit
https://github.com/wboeke/tgui
https://github.com/hpcloud/termui / https://github.com/boombuler/termui - same?
https://github.com/shiyanhui/TextUI
https://github.com/go-gooi/gooi
https://github.com/mcastilho/terminal
https://github.com/lcaballero/archon
https://github.com/nboughton/sklui-go
https://github.com/enr/clui

— widgets TUI
## https://github.com/gosuri/uiutil — imrpoves the look of things
https://github.com/gosuri/uiprogress
https://github.com/gosuri/uilive
https://github.com/julienmoumne/hotshell — fucking sweet


— menus
https://github.com/sabey/menublock
https://github.com/4ydx/glmenu


https://github.com/xilp/systray

==========================================================================================
https://github.com/marthjod/scripts
https://github.com/WiseTrem/scripts - discovery of things
https://github.com/vineetdaniel/scripts - lots of daily use script

https://github.com/rsdoiel/shelltools - several golang shell tools that will be useful to pick apart

https://github.com/as/torgo - pure go tor?
https://github.com/andres-erbsen/torch - pure go tor?
———Multiverse OS——
https://github.com/alaa/gotor — [Example] Smuggle some HTTP on Tor using Go for fun
https://github.com/BrassHornCommunications/OnionWatch
https://github.com/alkasir/ptc
https://github.com/andreas-jonsson/torrent-crawler — torrent crawler

MOUNTING TORRENTS!!! https://github.com/k4rr1m/torrent-cache
https://github.com/e-asphyx/svhrp = Simple virtual hosting reverse proxy, goes with hrp
https://github.com/rolandshoemaker/tdf - distributed DNS on the cheap

https://github.com/bfix/gpg2hs - hidden service from GPG
https://github.com/wybiral/onions

tor router -https://github.com/Pholey/distribuTor, a controller, just for idea use
https://github.com/yydesa/torcat

https://github.com/codekoala/torotator - more tor circuits used

newest torrent? https://github.com/jaz1997/Torrent

https://github.com/santoshhh000/torr < has a tracker

pure go torrent? https://github.com/liugenping/torrent

https://github.com/as/torgo - pure go tor?
https://github.com/andres-erbsen/torch - pure go tor?

=== apt repository builder —
https://github.com/smartystreets/raptr - A simple, shell-based APT Repository builder and updater.

=============================================================================

# Multiverse OS
*Isolated, compartmentalized secure general use clustering operating system*

*You are from Earth-C137, using every version of your computer in parallel dimensions you can always switch to a version without a virus, ensuring maximum secruity.* (Not really, but the metaphor is similar to what is actually happening.)

*If the Universe has no center, then any where you look can be the center.*

## Introduction

*Multiverse is a Debian based Linux operating system (OS), utilizing virtIO based KVM/Qemu para-virtualization to provide reliable security through ephemeral compartmentalization of each individual application. Virtualized containment is apart of a defense-in-depth strategy to protect multiple isolated identities running in parallel.**

*Multiverse is an operating system that fundamentally re-approaches the Linux command line and graphical user experience to enhance and simplify with primarily introducing consistency to make learning and using Linux and additional Multiverse default security settings easier. Multiverse defaults focus on enhancing security and provide greater privacy protections using ephemeral compartmentalization to quarantine any infections.Multiverse re-approaches both the Linux command-line and graphical experience, which has been slowly built up over it's history, to make it more consistent down to the command-line flags of each application included with Linux. Increased consistency across applications makes learning and using Linux and additional Multiverse OS security features.

In addition to active application or proxy virtual machines, long running background jobs running in isolation inside *"utility VMs"* and with minimal access to necessary components. For example, a file management *"utility VM* preforms ongoing integrity analysis on all files associated with a given identity.

# A new Secure general use Linux operating system

Multiverse OS embraces the idea that average users torrent, play games, program, interact with LAN computers and for secure operating systems to be successful should be able to provide the same functionality previous operating system, but presented in a way that is secure by default and privacy protected from mass surveillance by default.

Beyond just the user experience changes, Multiverse OS emphasizes a e-evaluation of what should the operating systems focus, and default collection of tools for general use computing. Multiverse OS offers features built around modern computing needs. For example, extending default printer support further to by default just as easily support 3D printers, CNCs, laser cutters and other similar devices.

Multiverse focuses on enhancing Linux console and graphical usability by providing shim interfaces for universal consistency across applications making more accessible an OS with greater security by default. Using ephemeral compartmentalization to quarantine any infections improve security by default, hide the complexities using intuitive user interfaces and improve the default privacy.

Each piece of software started is isolated by both nested virtual machines (VMs) but also virtual networking, so if a piece of software does infect the VM, once closed, the entire system infected is deleted soon as the running software closes. Running software can persist file changes before closing; files are compared with other stored copies, scanned for viruses, eventually encrypted, duplicated/backed-up and added to a temporal change history.

Re-evaluation of fundamental aspects of Linux being explored with Multiverse goes beyond just the effort to bring consistency to the command and graphical interface and enhanced security and privacy defaults. Using the tools to provide compartmentalization, these same tools can also be used to simplify the process of combining and sharing physical computers into complex virtual machine "cloud" networks expanding across several integrated servers, all done with a simple graphical interface enabling users to move servers around the complex nested virtual network in their system by dragging and dropping servers from one router to another.

### Why?

After experimenting for months with *QubesOS*, it became clear that the solution with the most potential to defend against defend against the modern threats of evolving malware and state sponsored mass surveillance. However I decided it did not fit many of my requirements which led to the development of Multiverse OS, an alternative implementation that would combine the best aspects of *TailsOS*, *QubesOS*, *Whonix*, *Kali*, and even *Plan9*.

While QubesOS serves a a very specific role in the security and even penetration testing Linux community, Multiverse OS focuses on a secure operating system specifically designed around the use cases of typical computer users; enabling the use of ephemeral compartmentalization with: games, torrents, software+hardware development, local+remote file sharing and many other everyday tasks.

Multiverse OS also takes many of the best features of QubesOS while attempting to avoid mistakes identified within the QubesOS mailing list, such as the use of Fedora with a desktop environment for routers, inability to pass through GPUs, password-less virtual-machine setup, the use of fedora instead of the more common Debian as the host, and most importantly how the host and virtual machines interact. Multiverse is an experimental operating system, use TailsOS or Whonix for reliable, *tested*, secure computing.

*After over six months now developing this OS QubesOS likely received many improvements making much of criticism invalid. People interested in using ephemeral compartmentalization in Linux should learn about QubesOS before deciding to use Multiverse OS since they provide very similar feature sets.*

### Depth in defense security model by always assuming the worst will happen
Multiverse was created to combine ephemeral compartmentalization using full virtual machines, attempts to avoid tracking through fingerprinting and an attempt to defend against the growing number of sophisticated infectious bot networks, pooling resources of the infected typically for nefarious purposes. Even without visiting "sketchy" websites, these advanced bot networks now easily built from modern open source malware tools can leverage an XSS exploit to extend their massive bot networks. The "I don't use windows argument" makes even less sense and with all of this automated, you don't have to be special to be targeted, every connected computer is automatically targeted because spreading malware creates a valuable pool of computer resources and valuable Internet connections.

Multiverse OS re-evaluates a lot of assumptions about operating systems; at it's core, the fundamental being *rejection* of the viewpoint that viruses are uncommon, average users are unlikely to be a victim, average users are not special enough to warrant attack, or some operating systems, ones that are not windows, are outside the scope of the targeted. This viewpoint is why the responsibility of virus protection is delegated to closed-source third-party often ineffective anti-virus software and the focus of an operating system is to provide other essential tools like printer management. *Multiverse OS rejects this viewpoint and instead is built around the idea that viruses are common*, so common that protection can not be done by virus definitions based on viruses found in the wild, but instead by assuming every web-page you run JavaScript on is capable of infecting your computer, and regular virus transmission through web-site advertisements and XSS allows transmission of viruses on trusted websites. To defend against the growing number of fast evolving open-source malware/penetration testing frameworks that facilitate the construction of increasingly advanced malware networks, one must assume infection will happen, and plan around mitigation in a technological landscape which infection is a regular occurrence.

When these facts are considered with the amount of wealth invested in global mass surveillance, most computer scientists would come to the conclusion that most operating systems are not equipped to deal with the modern technological landscape.

After experimenting for months with QubesOS, I realized this style of OS was the best available solution to the problem described above. However I decided it did not fit many of my requirements, immediately I felt there were many ways this concept could be improved upon. So I started developing an alternative solution, wanting to be able to draw directly from the work of Tails, Whonix and Kali the foundational operating system changed from Fedora to Debian.

The focus would be making an operating system that fit the use cases of typical computer users, a secure operating system that supports playing games, torrenting, software development, local and remote file sharing, home automation, media management and more. One major criticism seen often in the mailing list QubesOS was the use of Fedora with a desktop environment for their routers, the inability to pass through Nvidia GPUs without hardware hacking, password-less templates that are not built locally, and the use of fedora on the host.

**Multiverse OS has not been properly peer reviewed by a quorum of security experts; Multiverse OS is an experimental proof-of-concept to seed important ideas that if successful may be slowly adopted by stable operating systems. If successful hopefully this test will help accelerate implementation of these concepts by better programmers in better operating systems. As the primary developer, I actively use Multiverse OS on all my machines, but I'm not a university trained security cryptographer. Use this at your own risk, do not trust this with your life or freedom, use better tested operating systems like Tails, Whonix or QubesOS.**

### Important Ideas
Multiverse OS is more than a secure operating system, possibly more importantly it is a re-evaluation of the entire Linux user experience. Multiverse OS makes fundamental changes based around important concepts and realizations after stepping back and imagining what major improvements could be made if we could rebuild the major components of the Linux user experience.

*A secure general use operating system is important for human freedom in the modern world, it must be able to torrent, game, do home work, manage projects and so on easily but within a secure and compartmentalized environment that is capable of protecting ones home automation system.*


* Privacy by default. Security by default.

* Ephemeral compartmentalization using full virtualization and PCI pass-through needs to be the default, and provide the foundation for the entire Multiverse system. As many of the complexities of this system should be hidden from the user by default, including logging in users directly to a *"controller VM"* operating inside the bare-metal host, with the GPU pass-through for near native performance, ensuring the user *never* has direct access to the bare-metal host or any of the top level routers during normal operation dramatically limiting the available attack surface. Interaction with these machines is limited to a very rigid API and user interface built into the desktop environment.

* Seamless windows can be achieved with Qemu/KVM with per identity window decorations using `xpra` without console access by establishing TCP connections over VirtIO.

* Multiverse should have built in support for robots in the same way every operating system provides default support for printers. Robots do not need to be human shaped, they can also be house shaped or 3D printer shaped.

* Simple, intuitive functionality to provide clustering a mixture of different hardware to create a secure network of Multiverse computers is an important step back towards secure decentralization and receiving many boons currently mostly confined to industry that came out of the latest advancements to virtualization.

* Cluster log management, configuration management, local previsioning+building should be confined to a single desktop experience. For example, all logs are forwarded to the active *"controller VM"*, the virtual machine that aggregates the cluster features and presents a single desktop experience that is familiar to novice users.

* Management of media is an important task for most suers and recognizing the importance enables Multiverse to provide these features so that these features are secure and respect privacy by default; for example, a broad media management system enables automated tracking and removal of all undesirable leakage of meta-data in any media format. This includes checking for viruses, sharing and usage.

* Media management should also include providing a central secure media center, a community multi-user computing experience, with a shared central database, media collection and consistent console and API.

* Containers are not sufficient for any security focused operating system; when weighing the risk. vs reward, containers do not provide enough performance increase to give up the massive loss of security, primarily because a significant portion of the performance gains achieved by containers is boot time.

* Powerful open source home and business building automation and robotics needs to be established before a for-profit closed source company tries to dominate this industry, effectively eliminating privacy for large portions of the population.

* Measuring is the first step required to make logical changes for desired improvements. This is true for minimizing energy usage or managing ones nutrition. By providing a secure system and consistent API for home and business automation, vast improvements from lowered energy to automated improvements with statistical learning.

* Initramfs as currently implemented in most Linux distributions is probably the most important attack surface in modern Linux. A consistent and reliable way to check integrity needs to be built in a decentralized way that allows for complexity beyond everyone using identical builds created by a central authority.

* USB management, with limited authorized input devices and integrity checking and proxy-ing must be a default feature.

* The Linux termina/console is long overdue for major update to make the bash language share some features with lua+ruby, providing human readable aliasing, clearly defined components (each piece of the user interface can be thought of as an inventory slot; each console user experience would have: a snippet manager, an alias manager, a command prompt, coloring, and so on. By laying out the user experience components, new users can very clearly see what they have access too immediately, what alternatives are available which quickly familiarizes them.

* Built in console to manage, access, sync, update, factory reset and data restore all mobile devices.

The most ambitious change is implementing an interface/shim system for all console applications, providing application naming, configuration management, consistent naming, flags, and eventually consistent API across every command line and graphical application.

Taking on this goal can vastly improve the Linux user experience, most importantly making the learning process faster and easier for new Linux users. The interface/shim allows for backwards compatibility while providing the ground work for eventual replacement with a more consistent default tool set built with Rust and Go.

This change is one of many planned, including changes that to provide better default support for use-cases like clustering. For example, `scp` becomes obsolete because the operating system storing clustering information is capable of using `cp` to copy across network file system. Similar to plan9, every connected computer component becomes available to the cluster in the updated file system.

## Multiverse Feature List

### Additional Security Components

* *Continuous integrity checking* - along with automated system backups, continuous integrity checks should regularly check files. The integrity process can be sped up using Merkle trees derived from backups as they are encrypted, and duplicates created of files can be used to guarantee integrity of critical files. Automated backup VMs run long running tasks to send the backups off-site using .git, spider oak, `ssh`/`rsync`, stenography use with social media images and videos to backup critical data like keys long term.

* *All logging from all VMs* are combined, no logs are saved on any VM. This allows for easier analysis, provides a consistent API, a consistent search mechanism and hides any information found in logs from successful intrusions to ephemeral VMs.

* *USB Kill/Dead Man Switch* built into the multi-factor authenticator used to login to the operating system. If this special USB drive is removed wipes the memory of the computer and immediately shuts down the system. To prevent interception of your system while it is on-line, the USB can attached by string (or similar) to your wrist to act as a *dead man switch*.

* *USB 2nd-Factor Login* modify `cryptsetup` to require a 2nd-factor, in the form preferably of a stenographic image stored on an encrypted or unencrypted drive. This in addition to a password will enable the system to boot and launch the controller VM. This USB drive can also then function as the dead man switch.

* *USB management* system to register USB devices, route USB connections, allow specific devices, avoid fingerprinting, defined input from single mouse or keyboard. Unexpected USB activity, such as an unregistered USB stick is inserted, or input from an unregistered input device can trigger a defined script including a default script to wipe the memory and turn off the computer.

* *Decoy and nuke OS* by extending `cryptsetup` by starting with the customized `cryptsetup` from Kali which adds support for system nuking, which can be done at boot if being forced to enter the system under duress using an alternative 'nuke password'. An alternative 'decoy OS password' will be available,to boot a decoy Debian OS install available from within Multiverse OS. Booting Multiverse OS will appear to be a typical Debian OS install until the `cryptsetup` password is typed to initiate Multiverse OS booting.

### Ephemeral VMs, Template VMs

* *Snapshot based time dilation* by using automated snapshots actively archived for each active VM, allowing any file to be right clicked to expose a menu that allows one to retrieve historical versions of any file. Obtained by running a VM in parallel and copying the file into the live system in a way that is transparent to the user.

* *Sustain maximum encryption* by monitoring the *"controller VM"* or *"application VM"* using a cloned system running in parallel requiring only the active file system. The portions needed can be decrypted on the fly and provided to the active VM and removed after optional periods of inactivity. In addition, every VMs will have hard drive encryption, these features limit decrypted data availability.

* *Password Management* passwords are not stored on the *"controller VM"*, but either in the "bare-metal host" or a galaxy level server. By segregating the password-store, the damage of a compromised *"controller VM"*. The *"controller VM*" should alone not provide privileged access to any application, utility or *"router VMs"* without additional authentication factors. Important data is made available to the user o the *"controller VM*" through consistent and very limited APIs. APIs are validated by HMAC, GPG and using standard secure REST API features to prevent replay attacks.

* Graphical UI to visualize and manage all the interconnected virtual networks that make up your Multiverse OS cluster. Allows you to drag and drop servers between the layers. Making it very easy to expose or secure a server. One should be able to drag a server from one network and drop it into another network, having all the complexities of such an action completely hidden.

* Templates are built on first boot and rebuild-able at any point - this is an important feature because it removes required trust and allows advanced Multiverse OS users to vet and modify the template code so it can be easily updated and distributed.

* Any template should have the ability to spin up a throw-away version at any moment to quickly access an environment.

* Seamless windows - One advantage of Xen is that it supports seamless windows and KVM/Qemu does not support this critical feature. However this feature can be obtained by using the popular software `xpra` which provides every feature required to provide fast seamless windows with extremely limited attack surface. It can provide access to a specific program over TCP so that the *"controller VM"* does not even have console access to its *"application VMs"* which compartmentalizes intrusions and dramatically limits attack surfaces.

Using this feature, one can even pull seamless windows through multiple hops via clever use of `ssh` proxy or OpenVPN style proxies.

* Multiverse DNS - An important protection is to capture all *.onion requests that attempt to be made over clear net so the failure is never registered. Beyond this it  also important to be able to provide easy methods of enabling custom DNS for future anti-censorship projects and allow for quick navigation of the Multiverse OS cluster.

`ssh user@router.uni0`

* Meta-data removal - Along with long running integrity checking tasks, all files should be collected for search-ability and programmatic modification but also to automatically remove identifying meta-data and add randomization before sharing of certain files to conceal the origin.

## Important Features

### Seamless mobile clustering

Any modern operating system rebuilt from the ground up needs to consider mobile interconnection, management, security, and so on in a seamless way that is consistent with the rest of the operating system. This allows a mobile device to securely access and use the full power of your `multiverse-os` cluster.

*Console* access to mobile devices to allow for simple syncing, factory resetting and re-syncing to wipe potential viruses and so on.

### Database Management

As complexity in our lives grow the need for databases increase, and `multiverse-os` store all your data in an easily search-able and customizable database that features a consistent API. For your projects you can use whatever you would like but this provides a consistent way to store, use, share and analyze your data from contacts, files, emails and so on. This allows for more complex use of available resources and empowers anyone with basic programming skills to manage things in their life more intelligently.

The database will be a simple key, value store using any number of back-ends including BoltDB, LevelDB or MongoDB. The key, value storage will have an inbuilt graph structure that is available but not required to enable easier data analysis and extrapolating.

Simple scripts to handle common functions can then be shared to empower people.

[Example Uses]

* *Store all contacts* then program simple website to for everyone in your university working group to add themselves to your contacts under a group.

* *Store lab data* then against it in R or JavaScript to make beautiful graphs for your journal articles.

* *Track word usage across emails* with a simple program to histogram in different historical time intervals. Further analysis could be done to guess your mood using statistical learning algorithms.

* *File search across your cluster* and share filenames and meta-data with friends simplify media file-sharing seamlessly even behind NAT or firewall transported by transient torrents.

* *Banking information* crawl your bank data and store it in a local database for financial analysis and integration into project planning.

* *Calendar tracking* using all your available data, predict financial status at given points in the future, map projects, store important dates, predict problems.

* *Store cluster information* to fall-back when primary servers fail.

* *Fan management* for bands or stand ups to manage contacts in geo-location that are willing to see them live for live tour generation based on selected venues, given availability.

### 3D Printer, Laser Cutter, CNC integration

Queue management for multiple 3D printer, laser cutter, and CNC integration with shareable drop boxes for models to print, laser cut and so on. This allows for maximized use, sharing of resources (even geolocal available available prints for cost of materials plus an additional percentage to rent resources).

### Home Automation

Using a *controller* typically customized or for example a steam controller for house or business automation. Connections to devices are done through IR instead of 2.4 Ghz wireless to limit the attack surface. Access is controlled by a key on a carried RFID or optionally a person guessed on available data. All data is encrypted and sent via IR using a key on the controller/computer/mobile device and the RFID to manage home automation.

This will provide easy integration into many existing devices and provide a much more secure platform for home and business automation. This API will be consistent with the rest of the `multiverse-os` interfaces making it easy to interconnect and record data from home automation, experiments, projects and so on.

[Example Uses]

* *Automatic lighting* is a simple way to lower energy costs, turning lights on determined by individuals in the room, time of day, and season.

* *Energy usage tracking by plug* to track how much energy is being used, to incorporate energy costs into projects and to learn about your energy usage is a fundamental step to lowering energy usage. This simple feature could already exists in every house and it will enable people to learn and minimize energy usage.

* *Combine music with lights* with a simple script, have music data connected with lighting to provide interesting visual effects.

* *Connect lights to media center* to automatically dim lights when VLC is playing at night. Remove blue lights after a certain hour to help with sleeping disorders.

* *Have music track people in the house, only activation the speakers necessary for one to hear their music in their bathroom, bedroom, kitchen and living room.

* *Intercom system* to allow people talk with whomever they would like throughout the house without yelling making it easier to multi-task and stay connected, this feature could even work remotely to provide a method of hands free calling that works with either ones phone or computer.

### QubesOS and Multiverse, compare and contrast
QubesOS and Multiverse OS have a lot in common, they both utilize similar ephemeral compartmentalization techniques to provide reliable security, both options utilize full virtual machine based isolation instead of containers.

QubesOS uses Xen for virtualization while Multiverse uses KVM/Qemu because despite the 2%-5% performance increase with Xen it is far easier to get GPU pass-through to work with KVM/Qemu. Pass-through is a very powerful tool that helps given even gamers access to ephemeral compartmentalization.

|                        | QubesOS                      | Multiverse OS                                   |
|:---------------------- |:----------------------------:|:-----------------------------------------------:|
| Open source            | **Yes**                      | **Yes**                                         |
| Host (dom0) OS         | **Fedora**                   | **Debian**                                      |
| Virtualization         | **Xen**                      | **KVM/Qemu**                                    |
| Full Virtualization    | **Yes**                      | **Yes**                                         |
| Ran as Root            | *Yes*                        | **No**                                          |
| Router (net-sys) OS    | **Fedora**                   | **Alpine Linux**                                |
| Seamless Windows       | **Yes, with Xen**            | **Yes, with `xpra`**                            |
| Pre-compiled templates | **Yes**                      | **No, build templates on demand using scripts** |
| GPU Pass-through       | *Yes, with modern AMD*       | **Yes, all modern Nvidia & AMD**                |
| Built in Tor support   | **Yes**                      | **Yes**                                         |
| User space in VM       | *No, in dom0*                | **Yes, bare-metal hidden with minimal API**    |

**Table 1.** Summary of differences between the existing OS offering ephemeral compartmentalization QubesOS and Multiverse OS.

Multiverse OS makes intuitive usable user experience as important as security, this focus on simplicity and intuitive user experience is long overdue for Linux and great strides have been made in the last five years, but a lot of improvements can be made simply by bringing consistency to flags used in console programs and data storage. Multiverse OS prioritizing user experience specifically intended to increase access to secure open source computing, in some ways providing security features similar to QubesOS but presented in a different way to make these important concepts of security by ephemeral compartmentalization easier to grasp using metaphors found in popular culture.*

**Multiverse OS uses full virtualization**, any operating systems offering ephemeral compartmentalization based security must use full virtualization, containers do not provide adequate security. In addition, the performance gain obtained using containers largely experienced in the speed of boot time and the difference in run time is worth sacrificing for the very significant security advantage.

**Contrasting QubesOS virtual network organization in QubesOS**, the primary organization difference is that the user operates within the dom0 bare-metal machine, security is maintained by locking down the dom0 and directing the user to use software by launching VMs from templates. This user has administrator access to every virtual machine, and no machines have passwords. Multiverse OS moves away from this design, instead storing complex randomly generated 64 character passwords for root, user and hard-drives in a pass-store encrypted with their GPG key and optionally requiring a second or third factor. Password prompts are not removed but passed upward to better compartmentalize any unauthorized access instead of just letting the ephemerality of the system to be the only defense against active intrusion.

The primary difference is that with Multiverse OS the user operates inside of a very limited virtual machine inside the bare-metal host (dom0) instead of running a desktop from the dom0. The graphics card (GPU) is passed through this *"controller VM"*, network cards (NICs) are passed through to dedicated *"router VM"* able to secure and be highly customizable through nested routers and a simplified cluster port forwarding making multiple hops to any outside. This configuration let's the user snapshot their environment, isolate their game sessions to a VM, and so on. USB Devices are passed through to dedicated to USB proxies to further reduce the *"controller VM"* attack surface. In Multiverse OS, the bare-metal host (dom0) is *never* directly used by the user, when turning on Multiverse OS the system automatically boots into the *"controller VM"* without the user may not need to even know they are doing everything within a virtual machine. By avoiding the dom0 completely, confusion related to the networking restrictions of the dom0, proxified updating is taken away and combined with a more general cluster wide update system.

## Multiverse Core Components

**Portal gun** Use `portal-gun` to open wormholes to transverse the network of networks.

[Functional Requirements]

[Models]

[Reverse Networking Libraries]
https://github.com/JamesDunne/helpme: Self-contained SSH tunneling application for both sides of the connection 
very simple,json configs, 3 years, self contained both sides


https://github.com/Freeaqingme/SshReverseProxy: Proxy SSH Connections on a layer-7 (e.g. per user name) basis. 
a month ago, proxy sssh connections per uses on layer 7

https://github.com/Wolnosciowiec/reverse-networking
a month networking in base based on reverse proxying

https://github.com/elentok/gesheft - SSH Tunnel manager gesheft list    - lists all of the tunnels
gesheft active  - lists the active tunnels (removes zombie tunnels)

**Multiverse OS networking, reverse-(SSH|Ping|DNS)-proxy 'Everett bridges'**, the goal for the graphical representation will start with two options: (1) look like a router port forwarding management. This will allow you to setup servers, remote, local, tor connections. (2) Drag and drop VM icons around and drop them onto router VMs to attach them.

Then you can pick which servers are served over which available port or onion address. This will allow servers hidden deep within the `multiverse-os` cluster to be exposed over rented server IP addresses to give the Cloudflare protection without giving up all SSL security and serve pages over onion to completely conceal the local server.

This will function as the backbone networking for `multiverse-os`. Ideally it will pass to a router and the router will pass to servers within its local network. This will allow each server to focus on its own purpose and not have to worry about how its networked and that can all exists within the Alpine Linux routers which allows for better security.

This will let `multiverse-os` function as a reliable way to manage networking for several projects in a unified manner or manage several computers for personal use seamlessly.

[Description]

`portal-gun` provides all the tools necessary to setup and manage virtual machines using `libvirt` to provide compartmentalization of identity, application session and task. Everything from template building, launching, and `xpra` management for seamless windows. 

The same `portal-gun` tool also provides a nested system of routers built on top of Alpine linux. 

In a Multiverse OS cluster, for example with off-site Multiverse OS servers, including virtual private servers (VPSs), rented dedicated servers can be added to a local cluster. In addition, cluster Internet access can be routed through `openvpn` accounts, SSH connections, DNS Tunnel servers, ICMP Tunnel servers, and onion services to allow any inaccessible deeply nested VM or virtual network (VN) exposure using built-in *advanced multi-hope tunneling*, *"reverse-networking"*, and *reverse-proxying*. These tunneling techniques described are referred to by a variety of names, but for the sake of metaphor, Multiverse OS uses the term *worm-holes* or *Everett bridges* which connect universes into connected clusters of VMs. 

These networking techniques to expose deeply nested networks or virtual machines provides substantial security enhancements to development and production services by exposing access to the Internet with surgical precision, and more importantly hides the actual IP address of the host server. It is worth mentioning, these networking techniques effectively provide the same anti-DDOS and security benefits of proxified security features provided by Cloudflare but vastly improved because you are not required to give your SSL certificate to a for-profit company whose primarily writes closed source software.

These features make Multiverse OS useful for software engineering, hosting, and managing diverse projects. You can have a server deep within your Multiverse cluster and expose it via an ssh connection to a rented VPS somewhere else in the world. Even if this service is attacked, it would need to attack the server providing outside connection, which would be an isolated with a virtual router and rebuilds itself on reset.

Allowing the use of distinctly different connections to the Internet is possible in QubesOS but it may be possible to to provide an easier way to navigate these complexities in a way that much of the complexity is hidden from the user. Despite the similarities, QubesOS and Multiverse utilize different approaches and have different goals. Multiverse OS is designed from the ground up to be used by an average user, who torrents, plays games, is programming, hosting servers, general 'Internet-ing', home automation, or makes things with 3D printers/CNCs/Laser cutters. Specifically, in the same way printers are implemented in every operating system, we are going to implement these basic features so the most reliable way of securing them becomes default.





**Scramble suit** Complete account management and data management with `scramble-suit`. *Dissociate to have distance from all your socially recognized identities with `scramble-suit`.*


[Functional Requirements]

[Models]

[Description]
For example, the Multiverse OS `scramble-suit` account system is not a simple collection of settings, distributed across the home directory in an inconsistent manner. Instead, all Multiverse OS account data is stored in a consistent key/value database with the ability to abstract graph relations. Storing all settings data, account data, file data and so on in a in a secure way that by default has built-in in encrypted off-site backup. By providing a consistent account database API (Application Programming Interface) data from home automation sensors can be stored, emails, project details and so on; it enables simple scripts to make use and much easier to use statistical learning and automation. `scramble-suit` also automates the process of setting up local settings on remote machines and automatic cleanup, the local settings carried over are defined by the server category defined within the cluster. Each server uses its own SSH key, generated per server, regularly regenerated to enhance security and always password protected with a 64 character password. Entire identities, including files, email accounts, emails, contacts, projects and Multiverse OS cluster portions are stored within the `scramble-suit` software. Everything is easily accessible using a account console, or a consistent API to store any type of collection for later use in combination with all the data stored in a given identity.

Multiverse OS re-approaches the fundamental question: *"what should an operating system do"*, a question that should have been asked regularly by all operating systems, not just Linux. By asking this question, we realize that if we are going to continue to have built in support for printers, then it must be expanded to just as easily support, share, and queue laser cutters, 3D printers, CNCs, Lathes and more. Another example is that re-evaluating the relationship between the file system and the operating system, we discover that operating system can do more than just list the names and truncated details when requested and instead fully manage duplication to prevent potential corruption, preform integrity checking, using identity and keys from the `scramble-suit` provide encryption to secure and share files, as well as configurable off-site backup, and so on.


-----

## Multiverse OS Components
Multiverse OS seeks to provide more than just re-skinning Debian, alternating the default installed apt packages and providing configuration scripts that enhance security. Multiverse OS seeks to rebuild most tools from the apt package manager to basic networking tools and write them in compilable system languages like Go language and Rust instead of the more common QubesOS seen in Debian and QubesOS.

Multiverse OS will introduce a new console interface for Linux, one that by default manages multiple machines and simplifies cluster management. For example, the new Multiverse cp will cover the features provided by `scp`, allowing copying across machines in a single Multiverse OS cluster as intuitive as copying between folders.

Multiverse OS introduces functionality to make all user operations compartmentalized and ephemeral by default. It also provides tools to manage projects, manage media collections, validate & encrypt & backup files and key/value store databases tied to to the active user identity.

Re-envision existing functionality to provide a truly updated Linux experience that addresses existing problems and shifts priorities to modern problems away from a paradigm that was conceptualized before security was a concern.

**portal-gun** the virtual-machine, provisioner, manager, destroyer and the system that facilitates the clustering of multiple bare-metal hosts into a cluster.

[Virtual Machine Management]
Provisioning, configuration management, roles, networking and port forwarding presented as simple as the port forwarding page in the webUI router.

[Controller VM Management]
A system to have at least one *"controller VM"* running at a time, with a simple fall back system that can be launched if the current system is failing. This would prevent any need for the user to ever use the bare-metal (dom0) host machine, and provide a way to interact with host over the minimal API using minimal *"controller VMs"* that allow fall back and relaunch of various *"controller VM"* options installed.

[Xpra Window Forwarding]
Beyond just managing VMs, `portal-gun` will also manage *seamless window-forwarding*, and window decorations to help indicate which windows belong to which `scramble-suit` identity. This feature is not provided by KVM/Qemu, so the third-party software `xpra` is used to provide this functionality with window decorations. `xpra` also provides very limited access and can even provide windows without any console access over a specified TCP port. 

[Simplified configuration files]
`portal-gun` supports sharing tested configuration files and provisioning scripts for a variety of virtual machines. In the same way the apt package manager makes available pre-built and tested binaries, `portal-gun` exposes the user to a variety of community built and tested virtual machines to isolate and preform a variety of background tasks or provide an application or a collection of inter-working applications in a isolated network.


#### Multiverse OS networking, reverse-(SSH|Ping|DNS)-proxy 'Everett bridges'

Ideally this will eventually look like a router port forwarding management. This will allow you to setup servers, remote, local, tor connections.

Then you can pick which servers are served over which available port or onion address. This will allow servers hidden deep within the `multiverse-os` cluster to be exposed over rented server IP addresses to give the Cloudflare protection without giving up all SSL security and serve pages over onion to completely conceal the local server.

This will function as the backbone networking for `multiverse-os`. Ideally it will pass to a router and the router will pass to servers within its local network. This will allow each server to focus on its own purpose and not have to worry about how its networked and that can all exists within the Alpine Linux routers which allows for better security.

This will let `multiverse-os` function as a reliable way to manage networking for several projects in a unified manner or manage several computers for personal use seamlessly.

-----

## Multiverse Data Structures
Configuration files should be consistent but support any preferred type (*.yaml, *.toml, *.xml, *.json) and the global conversion between types.

[Models]

**Server:**
  **network_devices:**
    *ip_address*:string
    *mac_address*:string
  *onion_network*:boolean
  *open_ports*:array
  *forwarded_ports*:array
  **keys:**
    *key_type*:int
    *public_key*:string
    *private_key*:string
**Tunnels:**
  *status*:int
  *active_connections*:int

**scramble-suit** the account, database, identity, key and project management software that includes a consistent console, REST API and reliable, automatable encrypted off-site backup to unique locations for maximum security.

[Inventory System]
After development started, the concept of the inventory system was designed to help introduce new users to the variety of available experience components. Allowing them to quickly conceptualize what they have control over, support easily switching to alternatives, and providing structure because simply put the freedom of linux can be overwhelming at first, adding some structure helps news users get familiarized quickly and allows older users to drop in their existing systems. 

The concept is still an experiment but is currently stretched to even support broad default application support required for customization in the media management system. In this way you can equip a program to a slot, such as a music player slot.

**`scramble-suit` Account Management, this is where the entire invetory would be defined and would be available per identity while supporting cloning of identities.**

* Language
* Keys
  * Default file encryption key
  * Default text encryption key
  

* Graphical Interface
  * Browser
  * File Manager

* Media Management
  * Photo Viewer
  * Photo Editor
  * Audio Editor
  * Audio Player
  * Video Editor
  * Video Player

* Text Editor


* Calendar/Todo


* Terminal Console Interface
  * CLI Prompt
  * CLI Browser (for example, Lynx)
* Bash Script Library and package management
  * Alias Editor (For console aliases)
  * Snippet editor


[Using pass to GPG encrypt VM passwords]
Since all VMs are unprivileged, at least its worth going onto the alpine boxes and making their password is generated with `pass` and encrypted to your PGP key. That the password is not the same across machines, and it is 32+ characters. Ideally you would want all your VMs to be LUKs encrypted, with their passwords generated in this way.

I believe this may provide more security than currently provided by QubesOS, since QubesOS just has every machine have no root password. Any access is immediately escalated to root, which is argued to be okay since the machines are ephemeral besides the /rw folder.

It would be nice to do both though, have a /rw folder + ephemeral VMs and secure the fuck out of them. My guess is it is easier to break out of a VM with root access.

-----

**Available metaphors** magnetosphere, protects the whole solar system

### Multiverse OS infrastructure building blocks, and the fundamentals of Multiverse cluster design
Multiverse OS is designed to work with a single bare-metal computer but it also designed to work across several bare-metal CPUs. Specifically, Multiverse OS enables general use clustering or what is now myopically marketed as cloud computing, enabling end users to easily take advantage of advances in virtualization and systems automation.

Cluster management, networking management and virtual machine management is handled by the Multiverse OS tool `portal-gun`.

Multiverse OS uses four primary types of VMs to simplify management, each of these primary types of VM has several modular sub-types. These VMs provide functionality for a wide variety of use cases. Currently, clustering or cloud computing is primarily used by commercial industries but Multiverse OS enables even novice end users to combine two or more home computers into a powerful and secure general use computer.

To accomplish this, Multiverse OS establishes two categories of **"a bare-metal computer"** to help organize the computers into a single more powerful computer *'primary host'* and *'additive host'*. Within these Multiverse OS has four basic categories of VMs that make up the OS: ["router", "controller", "application", "utility"]

Multiverse OS intentionally prevents users from ever directly using the bare-metal host operating system. Instead by using nested virtual machines (VMs), the user always operates within VMs. The "controller" VM is passed the bare-metal host GPU which gives the VM near native speed, while preventing potentially dangerous GPU drivers from being installed on the bare-metal host. It was a conscious choice to not use containers because they do not provide sufficient security and much of the performance boost comes at startup.

VMs "above" and "below" the "controller" VM are only accessible through very limited APIs exposed through the VirtIO bus (is it technically a bus?). Using the VirtIO bus is fast, obfuscated since HTTP and SSH are more common transports and since it is already used by the existing system, it minimally increases the attack surface.

Multiverse OS uses four primary types of VMs to simplify management, each of these  primary types of VM has several modular sub-types. These VMs provide functionality for a wide variety of use cases.

1. **A Bare-Metal Host** a bare-metal host is either in active use as a controller or is an auxiliary host giving its resources to the cluster. This allows your laptop and desktop to essentially function as a single computer or one of many profiles available in your cluster.

2. **A Router VM** A router VM is currently primarily being built from the Alpine Linux micro-kernel. This OS is incredibly small, security focused and uses minimal resources. The primary router VM received the bare-metal NIC by utilizing the PCI pass-through feature.

A router VM is currently primarily being built from the Alpine Linux micro-kernel. This OS is incredibly small, security focused and uses minimal resources. The primary router VM received the bare-metal NIC by utilizing the PCI pass-through feature.


*0.router.universe* Universe prime is the computer which you operate, the one you are actively at. Others will start with Universe-1 and so on. This is the PCI pass-through level, local area network is accessible. Good for torrent machines (depending on region and stealthiness of your torrent site of choice. A simple VM with torrent client and samba shares for the media center goes here.

*0.router.galaxy* Galaxy prime is the network layer which is protected, all traffic is isolation proxified with an OpenVPN connection. This can be your normal Internet traffic for searching or when something is very tedious to do with Tor. This traffic can not access the LAN.

*0.router.solar* The local solar system, this network potentially exists within the controller VM. This traffic can not access galaxy or LAN networks. It is isolated, and can be routed through a Whonix box, can serve files or services through an onion or reverse proxy.

The number sub-domain allows for the numbering, for computers or servers with more than one NIC, to support multiple router.universe VMs. Or Several route.galaxy, allowing multiple bridge connections to proxy outside networks. Multiple router.solar are the most common configuration.

Instead of 0, the best configuration would be a randomized id, but either should be supported.

[router.uni0/net-sys] Every network device receives its own router.uni0, starting with router.uni0... In QubesOS this is analogous to net-sys router.

[router.gal/net-firewall]

[Captive Portal Bypass Router]
 (DNS Route, IMCP Route, or Both) -- (when traveling may be useful)
a simple DNS tunnel - https://gihub.com/radaiming/DNS_Tunnel - A simple DNS tunnel written in Go language. Would be really fucking cool if a router could be setup that has DNS routing. Then can be turned on and used if a behind a pay wall thing or login. Maybe even ICMP could be used, a router could be used to route traffic over ICMP. maybe the same server can flip between DNS, ICMP or both (random). https://github.com/Makesadbek/tcpovericmp

[Multi-hop proxying, reverse proxying, reverse networking]

gesheft info    - shows information about a tunnel

gesheft start   <tunnel_name>
gesheft stop    <tunnel_name>
gesheft restart <tunnel_name>

A CLI tool, but clearly would be good for a lib

3. **A Controller VM**
  * Run a desktop environment, preferably gnome to make it easier for new users. This virtual machine will be pieced together by machines across the Multiverse OS cluster. It will feel like a single computer but it will be comprised of a complex network of local and remote virtual machines combined into a cluster of bare-metal servers.

is where the user accesses the cluster, a controller can be active or inactive and remote or local. The remote is an auxiliary type, and just makes its own available resources.

Additional network interface cards (NICs) are attributed to their own router or can be merged into a single connection under a universe router. This allows easy combining of two Internet connections.

A virtual machine that receives the GPU and USB pass-through. USB may be further passed down to provided an USB proxy nested. Or perhaps a universe router level to further segregate USB proxy.



4. **A Application/Proxy VM** Run single applications, like a browser preferably in alpine linux
  * *Browser* allow the user to open a browser on top of alpine linux with the downloads folder mapped to the downloads folder of the controller VM.

A thin virtual machine defined by a simple *.yaml configuration file to expose exactly what is needed by an application and re-created from a template every time it is started. For example, a downloads folder can be made write only and mounted to ~/Downloads.

Each proxy VM (as in proxy computer not proxy networking) should have an attached downloads folder which should make transferring files unnecessary. Home folders can be shared if necessary for special projects.

Base it off a template an make them ephemeral. hold passwords in the host0 machine using password-store (pass).

Remove Firefox from Skylab (controller VM), use Firefox through a *"proxy VM"* shared of `xpra`.

xpra start ssh:SERVERHOSTNAME:100 --start-child=xterm (This is done by the go language software called `portal-gun`)

**Example *.yaml configuration*

````
servers:
  server:
    # Optional Name
    type: "application"
    applications: ["/usr/bin/firefox"]
    name: "LA"
    host: 125.125.125.1
    type: rsync
    key: ~/.ssh/chance.pub
    # mounting schema [{:local => "~/.ssh/keys", :remote => "~/keys"}, "~/pictures"]
    # if only a single directory is defined, then it assumed to be the same for both
    mount_directores: ["~/keys", "~/pictures"]

````

**_/!\_** *Warning!* Do not use `sudo` on any of the *"application VM"*. Just provide functionality inside `portal-gun` to use `su` and `pass` store to retrieve and use the password. Make this privilege escalation much more difficult from the perspective of an intruder by compartmentalizing each required piece but hiding these complexities from the user.

5. **A Utility VM** is a collection of scripts to be run on a virtual machine to preform long-running background tasks on a specific set of data. By compartmentalizing this, it limits the exposure of identity data to only what is needed by the operation.
  * Long running background jobs
    * Remove meta data from images
    * Virus scan files
    * Backup files across the cluster in a torrent like way in encrypted blocks

[intellectual property is not real]

Automatically submit patents for every project put on the site and remix existing patents and resubmit. Do this constantly and encourage others to do it

Patens should not just be limited, they should be ignored, laughed at and any claims to intellectual property fully and completely rejected. Idea men are worthless, anyone creative quickly realizes that ideas are never the limiting factor.

Regularly patent patents, regularly patent automated patenting, realgetlay patent suging people or patents automatically.

"But what if people dont share their work then!" They will because it will be required to be open source because like almost everything important in computer science requires the source to remain open. This solution is modern, works and promotes far better environment for makers, creators, scientists and researchers.

Reject the techno-elite class of supposed super humans establishing churches to their supposed genius. 

Common misconceptions of patents: they mean something

Fact: They don't, use pseudonyms, publish with tor and work without restriction releasing open soruce code that benefits humanity instead of creating intangible property that keeps corporations profitable indefinitely ensuring eventual human enslavement by stupid AI (corproations). Fuck the greedy money men trying to build a new church.

[Metaverse] 

Use the biggest populated websites to define a mainstreet oe internet and create a virtual world for them.

[OSInt Toolkit]

A collection of tools to allow for a long running search and google alert for a specific user, gatering all the data.

[Identity bomb]
Autoamte the creation of account with information similar enough to yours to trigger intelligence gathering bots to compile it but with incorrect information to make the overall data set completely unreliable.

[Grab data and dumps from pastebin and co]
Interesting stuff there

http://www.leakedin.com/tag/e-mail-headers/

[Wifi hacking VM (Zerg Rush?]
A VM that automatically scans on all available conenctions, looks for openc onnections, then hacks them if possible and saves the passwords. Autoamtes hacking all enarby machines to the server.

[Maker VM]
gpio (to create a utility VM for connecting to raspberry pis)
https://github.com/platinasystems/gpio  -- minimalist lib

[Captive Portal Login VM]
Simple VM to login into captive portal, attaches to router.universe (LAN)

Kicks up a disposable Firefox (also has a chromium browser in-case the portal is poorly written)

[Archive/Backup VM]
Archiving and backup should probably be more built into the scramble-suit software, because it should already be keeping all your files in a collection. Within those files other collections will be defined by the user and those collections, some of them available by default, will help the user determine which is duplicated, how often, and so on.

A VM which easily configures, using YAML files (and a UI to modify YAML files, define which folders get uploaded where) to backup via:
* rsync to remote servers (ssh)
* spideroak
* github
* stegographic data in images saved to social media
* ?

**Example *.yaml configuration*

````
servers:
  server:
    # Optional Name
    name: "LA"
    host: 125.125.125.1
    type: rsync
    key: ~/.ssh/chance.pub
    directores: ["~/keys", "~/pictures"]
````


### Fundamental User Interface Building Blocks

1. **Every primary feature provides a consistent low level console interface** direct interaction to empower super users, make testing and programming easier. Ideally featuring a *Lua* or *Ruby* style language to both interact directly and program with.

2. **Every primary feature provides a consistent HTTP REST API** to make scripting and interconnection between features incredibly easy to interact with using existing tools and most importantly easy to automate.

3. **Every primary feature provides a graphical user interface** to allow novice users to easily migrate to `multiverse-os` because a consistent and intuitive graphic interface is provided for every feature. As a stop gap this can be anything from QT to webUI, long as the structure and interface is consistent the actual back-end used to build it is not as important.



-----



-----

### Design Road map
The scope of the project is very large so there must be a basic structure that satisfies a essential design while auxiliary pieces that would improve the design but are not critical.

**`portal-gun`**
* Automate snapshot making
  - Time dilation features - Using snapshots, one can boot a VM and pull a file from x amount of time ago. Using this one should be able to right click any file and pull a previous version.

* Start turning the proxy-multiverse go language script into portal gun.
  - Create server, define roles, pull config from
    * Create config set for routers, each type ["uni", "gal", "sol"] with useful but very limited interfaces to reset, switch OpenVPN connections, obtain new identity on tor. This all should have built in access right in gnome, in the Internet drop-down.
  - Modify local /etc/hosts file to add in the new servers
  - Modify remote server `/etc/xpra/xpra.conf` (For example, configuration is slightly different for Ubuntu vs Debian)

* All configuration is stored in key/value but can be exported or modified from the file system in the form of *.yaml or *.toml files.

[router VM]

[controller VM]

[application VM] 
  * Ideally run on Alpine Linux but support both Alpine, Debian and Ubuntu by default.
  * Use a configuration file to define: 

**Example *.yaml configuration*

````
servers:
  server:
    # Optional Name
    type: "application"
    applications: ["/usr/bin/firefox"]
    name: "LA"
    host: 125.125.125.1
    type: rsync
    key: ~/.ssh/chance.pub
    # mounting schema [{:local => "~/.ssh/keys", :remote => "~/keys"}, "~/pictures"]
    # if only a single directory is defined, then it assumed to be the same for both
    mount_directories: ["~/keys", "~/pictures"]

````

After defining the mount directories in the *.xml, then the client `/etc/fstab`.

[utility VM]


[VM manager system] 
Like the Debian apt package management but for the files needed to locally build template virtual machines.

**`scramble-suit`**
* Key/Value Store & Graph Database
  - Store all files, in this structure, the file system does not need to be represented in the traditional way, a system where a file can be in two folders (not linking). It can be organized in separate collections with their own unique hierarchies.
  - Provide a simple REST API to create a collection, add to it, edit it, delete it in a simple key/value structure. All configurations, file system data and everything else OS related will be stored in this database. This allows the `scramble-suit` identity to be moved around and reused in different Multiverse environments.

----

## Manual Multiverse OS Setup
For the time being while a install script is being developed to run on Debian OS and eventually possibly supporting both Debian and Alpine.

**Install scripts are being built using bash, basic setup and installing go language to complete the process by configuring `portal-gun` and using it to initialize and setup Multiverse OS. The below process is being implemented on several different hardware configurations to finalize the install script requirements.**

**Notes from setting up remote additive bare-metal Multiverse OS instance on a Dell PowerEdge rack server**

[Preparing the BIOS configuration for Multiverse]
If possible, update your BIOS, preferably to an open source BIOS like Coreboot or Libreboot.

In the BIOS, one must enable virtualization technology, enable nested virtual machines and any other related virtualization options. 

Enable a BIOS password to make some attacks more difficult.

Disable features related to remote administration, Multiverse relies on its own system and these features could potentially be vectors for back door type vulnerabilities. 


[Install Debian 8+, upgrade to stretch]
Do not use the Internet, do not open a browser, do not install GPU drivers.

Install from disk, use `xfce`:

`sudo apt-get install remove xfce*`
`sudo apt-get install install i3*`

Install Debian from the most recent net install, for the Desktop Environment select `xfce` and only select system utilities. Set the host name to "host", and make the user name "user". Do not install `sudo` or provide the user account with any privileges.

`su`

`apt-get remove update && apt-get upgrade && apt-get remove xfce && apt-get install lightdm i3*`

At a bare minimum, set `exec i3` in the `~/.xinit` file under user home folder. 

`cat "exec i3" > ~/.xinit`

Update to stretch, update and upgrade again. Maybe restart, your choice.


[Important Options]
These are important configuration options needed for the install

**Hardware Options**
* AMD/Intel CPU
* Nvidia/AMD GPU
* Available NICs
* Available Hard-drives

In the OS all user names are always 'user' but the identity name helps organize identity name within `scramble-suit`. The domain is always named 'host'. Actual names are abstracted and hidden; only shown to the active user.

**Software Options**
* GPG Key, from file or generated

[Prepare apt packages]
Install the necessary additional apt packages:

`apt-get install virt-manager uass vim`

Remove `nano` to force `vim` to be the default editor easily.

`apt-get remove nano`


[Prepare the user account to start a user session VM]
Add user to the necessary libvirt groups to enable the account 'user' to start virtual machines within the user session domain.

> **Error** Unable to connect to libvirt authentication unavailable no polkit agent available...

`usermod -aG libvirt user`
`usermod -aG kvm user`
`usermod -aG libvirt-qemu user`

((NEXT STEP))

# Setup the local multiverse folder (Or maybe .local/share/multiverse)

  mkdir /home/user/.multiverse-os/
  mkdir /home/user/.multiverse-os/portal-gun/persistence/
  mkdir /home/user/.multiverse-os/portal-gun/storage-pool/
  mkdir /home/user/.multiverse-os/portal-gun/os-images/
  mkdir /home/user/.multiverse-os/portal-gun/conf/
  mkdir /home/user/.multiverse-os/scramble-suit/

# During setup the default os-images should be copied
# During setup the default confs should be copied

* There should be a step here to wget the debian images and have a checksum locally (or the developer keys?) to check integrity
Keys [here](https://sks-keyservers.net/pks/lookup?op=vindex&search=debian-cd%40lists.debian.org&fingerprint=on) and [here](https://www.debian.org/CD/verify)

**Each drive should have a configs folder mounted read only.**
**Proxy VM (controller VM, auxillary VM) should get host-storage-pool read only, and os-images shared to the same directory in the host as in the proxy VM.** *This will let os-images be shared, and templates to be used from the higher levels.*

*Then extra Hds should be passed through and all USB devices, Audio devices, video devices.*

[Build the multiverse-os folder structure]
Make the `multiverse-os` folder structure, copy over (or download and validate) *.iso images for Debian and Alpine Linux. In development this is just stored in `/home/user/multiverse-os` but eventually will be stored in /etc/multiverse and ~/.multiverse-os or ~/.local/`multiverse-os`. 


[Update grub, to define kernel configuration and kernel modules required]
Define the LINUX_DEFAULT line in grub and copy over relevant kernel module configuration. Preform necessary changes to fit the NIC and GPU hardware requirements.

> **Error** Unsupported configuration: host doesn't support pass-through of host PCI devices...

Modify the configuration file depending on CPU type, in `/etc/default/grub`.

````
echo "GRUB_DEFAULT=0"                                    >  /etc/default/grub
echo "GRUB_TIMEOUT=5"                                    >> /etc/default/grub
echo "GRUB_DISTRIBUTOR=Multiverse"                       >> /etc/default/grub
# if cpu.is_a?(:intel)
echo 'GRUB_CMDLINE_LINUX_DEFAULT="quiet intel_iommu=on"' >> /etc/default/grub
# else if cpu.is_a?(:amd)
# echo 'GRUB_CMDLINE_LINUX_DEFAULT="quiet amd_iommu=on"' >> /etc/default/grub
echo 'GRUB_CMDLINE_LINUX=""'                             >> /etc/default/grub
# else
# end
````

Then run the update command to rebuild the grub menu.

`update-grub /etc/default/grub`

> **Error** Unable to complete install: 'internal error: Failed to load PCI stub module vfio-pci'...

If you encounter this error, the *vfio-pci* kernel module has not been properly loaded. There are three common methods to solve this, but the method should be whatever works for your specific system and the requirements you have based on when GPU, USB or NIC need to be loaded/enabled or unloaded/disabled.


*Commonly Suggested* One can modify `initramfs` config or grub config to load this module during startup. The ugly way, and may be loading before it is desirable.

*Sometimes Suggested* One can modify /etc/rc.local, then modify it here or load a script. This file should preferably not be used, but it can be a good fall-back/stopgap until better solutions are found. It also provides different loading timing. It comes after both other methods. This makes it useful for some tasks if video issues are experienced for example.

*A better way* is to modify `/etc/modules-load.d/`

````
  1 # /etc/modules: kernel modules to load at boot time.
  2 #
  3 # This file contains the names of kernel modules that should be loaded
  4 # at boot time, one per line. Lines beginning with "#" are ignored.
  5
  6 # Load the necessary modules for PCI pass-through
  7 vfio
  8 vfio_iommu_type1
  9 vfio_pci
````

Adding this does solve the problem. However I did not check if each module used was necessary to specify, options are often dependent upon processor/motherboard combination.

**_/!\_** Check if all three are required.


[Conditionally required module option: options iommu=pt]
With older motherboards and processors, one will have to sacrifice some security to enable all the features required for Multiverse OS.

> **Error** Unable to complete install: 'internal error: Failed to load PCI stub module vfio-pci'

**Conditionally Required** Older processors motherboard combinations want to add kernel module option: 

`option iommu=pt`

Specifying this kernel option sets the IOMMU into pass-through mode for host devices may sacrifice some I/O. 

> "When an operating system is running inside a virtual machine, including systems that use para-virtualization, such as Xen, it does not usually know the host-physical addresses of memory that it accesses. This makes providing direct access to the computer hardware difficult, because if the guest OS tried to instruct the hardware to perform a direct-memory-access (DMA) using guest-physical addresses, it would likely corrupt the memory, as the hardware does not know about the mapping between the guest-physical and host-physical addresses for the given virtual machine. The corruption is avoided because the hypervisor or host OS intervenes in the I/O operation to apply the translations, incurring a delay in the I/O operation.
> An IOMMU can solve this problem by re-mapping the addresses accessed by the hardware according to the same (or a compatible) translation table that is used to map guest-physical address to host-physical addresses." - [https://en.wikipedia.org/wiki/Input%E2%80%93output_memory_management_unit](Wikipedia Article on Input/Output memory management unit)

*The option reduces the overhead of the IOMMU for host owned devices, but also removes any protection the IOMMU may have provided against errant direct-memory-access (DMA) from devices.  If you weren't using the IOMMU before, there's nothing lost.  Regardless of pass-through mode, the IOMMU will provide the same degree of isolation for assigned devices.

*Enabling specific devices dependent on scanned variables*

Using Go language to scan PCI devices, one can build the necessary rc.local file to prepare pass-through for all the relevant files. *It will be necessary to find all network cards, USB and so on, to disable and prepare for usable by routers and controller VMs.*

If you're using Linux Kernel version 4.1=< or newer, the vfio-pci driver supports the same ids option so you can directly attach devices to vfio-pci and skip pci-stub.  vfio-pci is not generally built statically into the kernel, so we need to force it to be loaded early.To do this on Fedora we need to setup the module options we want to use with `/etc/modprobe.d`.  I typically use a file named `/etc/modprobe.d/local.conf` for local, for example, system specific, configuration. In this case, that file would include:

The default config starts from moving `blacklist.conf`, `dkms.conf`, `kvm.conf`, `vfio.conf` to `/etc/modprobe.d/`.

Then configure the kernel module using a `vfio.conf` placed in `/etc/modprobe.d/` containing the following

`options vfio-pci ids=1969:e091`

Or, for multiple devices:

`options vfio-pci ids=1969:e091,1969:e091`

To make sure this worked, `dmesg | grep vfio`. If you get `unknown parameter 'ids' ignored`, try binding the ids in /etc/rc.local instead (with the appropriate id):
`echo "1969 e091" > /sys/bus/pci/drivers/vfio-pci/new_id`

If you get some error along the lines of: 
> vfio: error, group 11 is not viable, please ensure all devices within the iommu_group are bound to their vfio bus driver

Check what's in what group and make sure you're getting all the devices in the group of interest
`find /sys/kernel/iommu_groups/ -type l`

This section should probably be lower, set up the basic vms before setting the passthrough permanently.

*Adding the `options vfio-pci ids=...` above will create the other necessary udev rules.* Defining the PCI device IDs, for some systems this only worked for the GPU and other devices such as USB had to be unbound after the hard-drive password was typed in. After the primary drive is open, the USB drives need to be unbound and bound to the *"controller VM"* using the `/etc/rc.local` and the `vfio-bind` script.

Using this combination, any PCI device can be bound and unbound at any step of the boot process allowing for highly customizable Multiverse OS setup.




[Modify the 'kvm' group to permit memory limits]
The 'user' user previously added to the group kvm, now needs to have their memlock expanded to allow virtual machines to be created in the user session instead of the root session.

> **Error** Error starting domain: internal error: Process exited prior to exec: libvirt: error : cannot limit locked memory to 1342177280: Operation not permitted


 24 # error with memory limit issues and not permitted to change it
 25
 26
 27 /etc/security/limits.conf
 28
 29 # One method of solving the issue is providing just the user upgraded memlock limits.
 30
 31 ````
 32 user             hard    memlock         -1
 33 ````
 34
 35 # The preferred way is to modify the kvm group, this will allow for more customizable configuration with multiple users adding virtual machines to the user session.
 36
 37 ````
 38 @kvm             soft    memlock         8650752
 39 @kvm             hard    memlock         -1
 40 ````
 41

verify the value of memlock:

$ ulimit -l

[vfio group permissions]

> **Error** Error starting domain: internal error: process exited while connecting to monitor: (Timestamp) qemu-system-x86_64: -device vfio-pci,host=01:00.0,id=hostdev0,bus=pci.0,addr=0x8:vfio error: 0000:01:00.0 group 11 is not viable

> **Error** Error starting domain: internal error: process exited while connecting to monitor: PM.disable_s4=1 -boot strict=on -device ich9-usb-ehci1,id=usb,bus=pci.0,addr=0x5.0x7 -device ich9-usb-uchi1,masterbus=usb.0,firstport...

If you experience the above error, you can confirm the error is related to Linux file system by attempting the following commands under the `root` user

(refer to section with udev rules for vfio)

Then under the account being used to launch the virtual machine check if the `libvirt` group was included.

`groups | grep libvirt`

[Blacklist any extra kernel modules] 

Blacklist absolutely any that are not absolutely needed to limit the attack surface of the bare-metal computer. A file `/etc/modprobe.d/blacklist.conf` is already supplied with the default Multiverse OS install.

````
ipmi_si
ipmi_devintf
````

[Disable all networking inside the bare-metal machine]

Disable all the networking scripts to avoid lengthy waits for devices or scripts to run unncessarily

`systemctl disable NetworkManager`
`systemctl disable networking`
`systemctl stop NetworkManager`
`systemctl stop networking`


[Connecting to the xpra server]

If the xpra server is not working, maybe no xpra server found

`xpra  start :13 --no-clipboard --no-pulseaudio --no-daemon`

When debugging the `xpra` server, running on the *"application VM"* to pass windows to the *"controller VM"*, logging into the *"application VM"* with `ssh` and running

`xpra  start :13 --no-randr --disable-mmap --no-clipboard --no-pulseaudio --no-daemon`

Often `stdout` errors will typically reveal any connection issues that are not present when launching `xpra` on the client. Server issues often involve (1) specific server configuration options or more often (2) missing python libraries that need to be installed with the appropriate python pip library manager.

Report any changes in requirements so install scripts can be updated to fit any changes with `xpra` until this tool can be replaced with a Go language or Rust seamless window software.


[Conditionally required kernel module: vfio_iommu_type1]

**Error** Error starting domain: internal error: process exited while connecting to monitor: _PM.disable_s4=1 -boot strict=on -device..

**Error** internal error: process exited while connecting monitor...failed to setup container for group 12: failed to set iommu for container: Operation not permitted


By checking the dmesg with *sudo dmesg | grep iommu* and with the Dell PowerEdge reports:

*vfio_iommu_type1_attach_group: No interrupt remapping IO IOMMU support on this platform.

**It is important to note that this is not required on all machines and should *not* be enabled if it is not necessary.**


((Step Next))

**Portal Gun** This portion of the software should be utilizing a golang application named portal-gun.

*One feature of portal gun should be to grab checksums from the Multiverse DHT. Then it should use those to cross check third party *.iso's like Alpine linux and check third party GPG keys.* It should also be able to hit up the alpine servers and get the newest iso. (since the baremetal should never access the internet. This can only be done after the controller is established.

Setup router.uni VM, pass the NIC over. Setup the router.


----- DIDNT DO THESE STEPS YET ------
((Step Next))

**Scramble Suit** *This needs to be done by leveraging the shifter suit portion of multiverse os. The portions of the install script that don't use [portal-gun] or [scramble-suit] can be written in bash or whatever. But these portions should be using golang applications that will be used actively.*

Generate a GPG key or load a GPG key.

Install pass, use the GPG key above to encrypt the keys.

Use this pass-store to hold all the user, root and HD passwords for the virtual machines in the multiverse-os install.

Use this GPG key to backup the HD key of the bare-metal host.

Use the GPG key to generate an onion address so its consistent.


((Step Next))

**Error** Error starting domain: internal error: /usr/lib/qemu-bridge-helper --user-vnet --br=virbr0 --fd=24: failed to communicate with bridge helper: Transport endpoint is not connected stderr=libvirt: error: cannot execute binary /usr/lib/qemu/qemu-bridge-helper: Permission denied

Somtimes this is an issue with the qemu-bridge-helper permission problems:

  mkdir /etc/qemu
  chown root:libvirt /etc/qemu
  cat "allow virbr0" > /etc/qemu/bridge.conf
  cat "allow virbr1" >> /etc/qemu/bridge.conf
  cat "allow virbr2" >> /etc/qemu/bridge.conf
  chown root:libvirt /etc/qemu/
  chown root:libvirt /etc/qemu/bridge.conf

  chown root:libvirt /usr/lib/qemu
  chown root:libvirt-qemu /usr/lib/qemu/qemu-bridge-helper
  chmod 4750 /usr/lib/qemu/qemu-bridge-helper


((Step Next))

Setup router.gal, connect to router.uni. This is the router that the controller VM connects to. This lets you define a tor whonix like setup, or persistent VPN.

A third router.sol is added ontop to provide a tor whonix like connection that is routed through the router.gal. This can let you have vpn over tor transparent to the controller or server.

* Setup alpine
* Turn off, add eth1 (virbr0)




-----

# Setup for Gaming Rig

To really be general use, we need to support a secure gaming environment. To do this one must be able to game inside a VM. There are several techniques required to provide the preformance necessary to do this.

cpufreq

manage CPU frequency..  /usr/bin/cpufreq-set, which allows you to set minimum and maximum frequencies for all cores or individually, as well as changing governers..  /usr/bin/cpufreq-info gives the current settings and /usr/bin/cpufreq-aperf seems to be a performance monitor tool. Much easier than catting and echoing!

 cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_available_governors returns "performance" and "powersave".. 

cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor is set to "powersave" by default.

echo performance > /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor

cd /sys/devices/system/cpu/cpu0/cpufreq

 

That's config info for cpu0.  You can monkey with your cpu in here.  "cat scaling_max_freq" resulted in 4300000.  So I thought I'd give this a try.

 

echo 4300000 > scaling_min_freq


<clock offset='localtime'>
<timer name='rtc' tickpolicy='catchup' track='guest'>
<catchup threshold='123' slew='120' limit='10000'/>
</timer>
<timer name='hpet' present='no'/>
   <timer name='hypervclock' present='no'/>
 </clock>


 50     <timer name='rtc' tickpolicy='discard' track='guest' />

<clock offset='localtime'>
      <timer name='rtc' track='guest'/>
    </clock>

If you have track='catchup' QEMU seems to try to send millions of interrupts so the guest clock will catch up to the current time, which can take minutes during which the guest is unusable.


for i in {0..7}; do
    echo performance > /sys/devices/system/cpu/cpu${i}/cpufreq/scaling_governor
done

echo 5120 > /proc/sys/vm/nr_hugepages

   <cputune>
     <vcpupin vcpu='0' cpuset='2'/>
     <vcpupin vcpu='1' cpuset='3'/>
     <vcpupin vcpu='2' cpuset='4'/>
     <vcpupin vcpu='3' cpuset='5'/>
     <vcpupin vcpu='5' cpuset='6'/>
   </cputune>

"I googled for ways to make the L3 cache visible in the guest and came across this post on serverfault.com, wherein the author asks whether it matters if the guest can see the details of the L3 cache or not. Here's the reply,

    I don't think it should matter.

    The host makes this data available to the guest, via a virtual CPU/Core. I can imagine that the host can provide the guest with arbitrary values without really affecting performance that much, since it's the host that ultimately determines performance anyway.

    On the other hand, if KVM does bare metal virtualisation, maybe the cache levels reported by the guest represents a direct correlation with the real CPU, since the guest has direct access to the hardware CPU. Thus installing a better CPU will give better performance in the guest.

Sorry internet, but you're wrong."


----

# Qemu XML options

<disksnapshot default='on' toggle='no'/>

-----

# Fucking timers

Xen is the best, KVM kinda sucks and is even worse when idle. https://www.ncbi.nlm.nih.gov/pmc/articles/PMC4503740/

**Why?**
Timers in the hypervisor

The Xen hypervisor uses local APIC for scheduling the timer events. The virtualized event device can work in two modes: singleshot and periodic. In order to schedule a new event, the VM has to call the VCPUOP_set_singleshot_timer or VCPUOP_set_periodic_timer hypercall accordingly. All scheduled timer events are stored in two data structures: (faster) heap and (slower) linked list. Each time the lapic interrupt occurs, the handler executes the expired timer events and, at the end, schedules next lapic interrupt looking for the earliest deadline timer. To improve the performance and lower the number of the timer interrupts, a timer_slop parameter is used. It denotes the amount of the time (in nanoseconds) that the timer can be late. All timers from such interval will be executed in a single lapic interrupt. The timer_slop is set by default to 50000ns.

Virtual timer interrupt

All interrupts in Xen are handled by the hypervisor. It then delivers them to guest OSs using so called event channels [12]. There are three types of the events: Physical IRQ, Virtual IRQ (VIRQ) and inter-domain events. Different hardware interrupts are delivered using Physical IRQs event channels. This is done to make the standard device drivers work correctly and map real IRQs inside the guest kernel. In the case of virtual devices, such as the aforementioned timer event device, the Virtual IRQs are used. There is a special virtual timer interrupt called VIRQ_timer.

A simplified process of handling hardware interrupt and delivering it to the VM using an event channel is presented in Fig 15.

**Custom solution for high precision timer**

New High Resolution Timer

To overcome the precision problems we decided to create our own timer implementation, which might be used only in cases when it is really required. For example in applications highlighted in the introduction other system processes might still use standard timer implementation built in the virtualization platform.

# Okay how do you configure the qemu xml?

*Time keeping*

The guest clock is typically initialized from the host clock. Most operating systems expect the hardware clock to be kept in UTC, and this is the default. Windows, however, expects it to be in so called *'localtime'*. 

-----

# Virtualizing OSX for Adobe, Final cut and such

-----

# Ephemeral

transient
    If present, this indicates that changes to the device contents should be reverted automatically when the guest exits. With some hypervisors, marking a disk transient prevents the domain from participating in migration or snapshots. Since 0.9.5 

 <transient/>

 The source element may contain the following sub elements:

host

snapshot
    The name attribute of snapshot element can optionally specify an internal snapshot name to be used as the source for storage protocols. Supported for 'rbd' since 1.2.11 (QEMU only). 

-----

# VirtIO Internet connectiosn!

Features/VirtioVsock

virtio-vsock is a host/guest communications device. It allows applications in the guest and host to communicate. This can be used to implement hypervisor services and guest agents (like qemu-guest-agent or SPICE vdagent).

Unlike virtio-serial, virtio-vsock supports the POSIX Sockets API so existing networking applications require minimal modification. The Sockets API allows N:1 connections so multiple clients can connect to a server simultaneously.

The device has an address assigned automatically so no configuration is required inside the guest.

Sockets are created with the AF_VSOCK address family. The SOCK_STREAM socket type is currently implemented. 

-----

# VM Optimization

**Power Management**

Intel processors have a power management feature where the system goes in power savings mode when the system is being under utilized. This feature should be turned off to avoid variance in vpp application performance. The system should be configured for maximum performance (bios configuration). The downside of this is that even when the host system is idle, the power consumption is not down.

For maximum performance, low-power processor states (C6, C1 enhanced) should be disabled. 

**Disable Interrupt Balancing (irqbalance)**

The Irqbalance daemon is enabled by default. It is designed to distribute hardware interrupts across CPUs in a multi-core system in order to increase performance. However, it can/will cause the cpu running the vpp VM to be stalled, causing dropped Rx packets. When irqbalance is disabled, all interrupts will be handled by cpu0, so the vpp VM (or any other service VMs) should NOT run on cpu0.

Disable irqbalance by setting ENABLED="0" in the default configuration file (/etc/default/irqbalance):

#Configuration for the irqbalance daemon

#Should irqbalance be enabled?
ENABLED="0"
#Balance the IRQs only once?
ONESHOT="0"

**In a VM: Disable Kernel Samepage Merging (KSM)**

KSM is a memory-saving de-duplication feature, that merges anonymous (private) pages (not pagecache ones).

While diagnosing the vpp Rx zero packet drop issue, we noticed a correlation between the /sys/kernel/debug/kvm/pf_fixed counter being incremented and the periodic Rx packet drops. We observed that disabling KSM eliminated the incrementing of these counters. KSM is enabled in Ubuntu 14.04 server on the host OS only. It is disabled when Ubuntu 14.04 server is run in a VM.

Disable KSM by writing "0" to /sys/kernel/mm/ksm/run in the host OS:

sudo bash
echo 0 > /sys/kernel/mm/ksm/run
exit

For more information, see: http://www.linux-kvm.org/page/KSM  

**In a VM: Remove VirtIO Balloon Driver**

Use of the VirtIO Balloon driver in the vpp VM causes Rx packet drops when the balloon driver calls mmap().

Remove the VirtIO Balloon Driver from the VM configuration:

If editing the xml configuration, remove the memballoon driver by setting the model='none':

  <memballoon model='none'/>

or delete the device definition from the command line parameter list:

 -device virtio-balloon-pci,id=balloon0,bus=pci.0,addr=0x6

**Huge pages**

Edit /etc/sysctl.conf and add this text to specify the number of pages you want to reserve (see pages-size)

# Allocate 256*2MiB for HugePageTables (YMMV)
vm.nr_hugepages = 256

# Members of group my-hugetlbfs(2021) can allocate "huge" Shared memory segment 
vm.hugetlb_shm_group = 2021

**Lower the halt polling interval**
You can pick a different polling interval with the kvm module option halt_poll_ns.  The default is 500000.  It seemed that setting this to 400000 or lower resolves the issue.  You can do this via a modprobe entry, 'options kvm halt_poll_ns=400000' or on the kernel command line with kvm.halt_poll_ns=400000.  You can also change it dynamically via 'echo 400000 > /sys/module/kvm/parameters/halt_poll_ns'


-----

## Built in Qemu tunneling features

 Multicast tunnel

A multicast group is setup to represent a virtual network. Any VMs whose network devices are in the same multicast group can talk to each other even across hosts. This mode is also available to unprivileged users. There is no default DNS or DHCP support and no outgoing network access. To provide outgoing network access, one of the VMs should have a 2nd NIC which is connected to one of the first 4 network types and do the appropriate routing. The multicast protocol is compatible with that used by user mode linux guests too. The source address used must be from the multicast address block.

...
<devices>
  <interface type='mcast'>
    <mac address='52:54:00:6d:90:01'/>
    <source address='230.0.0.1' port='5558'/>
  </interface>
</devices>
...

TCP tunnel

A TCP client/server architecture provides a virtual network. One VM provides the server end of the network, all other VMS are configured as clients. All network traffic is routed between the VMs via the server. This mode is also available to unprivileged users. There is no default DNS or DHCP support and no outgoing network access. To provide outgoing network access, one of the VMs should have a 2nd NIC which is connected to one of the first 4 network types and do the appropriate routing.

...
<devices>
  <interface type='server'>
    <mac address='52:54:00:22:c9:42'/>
    <source address='192.168.0.1' port='5558'/>
  </interface>
  ...
  <interface type='client'>
    <mac address='52:54:00:8b:c9:51'/>
    <source address='192.168.0.1' port='5558'/>
  </interface>
</devices>
...

UDP unicast tunnel

A UDP unicast architecture provides a virtual network which enables connections between QEMU instances using QEMU's UDP infrastructure. The xml "source" address is the endpoint address to which the UDP socket packets will be sent from the host running QEMU. The xml "local" address is the address of the interface from which the UDP socket packets will originate from the QEMU host. Since 1.2.20

...
<devices>
  <interface type='udp'>
    <mac address='52:54:00:22:c9:42'/>
    <source address='127.0.0.1' port='11115'>
      <local address='127.0.0.1' port='11116'/>
    </source>
  </interface>
</devices>
...

 Consoles, serial, parallel & channel devices

A character device provides a way to interact with the virtual machine. Paravirtualized consoles, serial ports, parallel ports and channels are all classed as character devices and so represented using the same syntax.

...
<devices>
  <parallel type='pty'>
    <source path='/dev/pts/2'/>
    <target port='0'/>
  </parallel>
  <serial type='pty'>
    <source path='/dev/pts/3'/>
    <target port='0'/>
  </serial>
  <serial type='file'>
    <source path='/tmp/file' append='on'>
      <seclabel model='dac' relabel='no'/>
    </source>
    <target port='0'/>
  </serial>
  <console type='pty'>
    <source path='/dev/pts/4'/>
    <target port='0'/>
  </console>
  <channel type='unix'>
    <source mode='bind' path='/tmp/guestfwd'/>
    <target type='guestfwd' address='10.0.2.1' port='4600'/>
  </channel>
</devices>
...

In each of these directives, the top-level element name (parallel, serial, console, channel) describes how the device is presented to the guest. The guest interface is configured by the target element.

The interface presented to the host is given in the type attribute of the top-level element. The host interface is configured by the source element.

The source element may contain an optional seclabel to override the way that labelling is done on the socket path. If this element is not present, the security label is inherited from the per-domain setting.

If the interface type presented to the host is "file", then the source element may contain an optional attribute append that specifies whether or not the information in the file should be preserved on domain restart. Allowed values are "on" and "off" (default). Since 1.3.1.

Regardless of the type, character devices can have an optional log file associated with them. This is expressed via a log sub-element, with a file attribute. There can also be an append attribute which takes the same values described above. Since 1.3.3.

...
<log file="/var/log/libvirt/qemu/guestname-serial0.log" append="off"/>
...

Each character device element has an optional sub-element <address> which can tie the device to a particular controller or PCI slot. 
-----

### net-sys, router.universe - router with passthrough NIC

**NOTE** *This is very handy to know about, /sys/class/net/ provides all the information in a filesystem style. This is much easier than chaining sed grep and so on. This is worth mentioning in guides. Multiverse should take advantage of this. For example, mac address:

``cat /sys/class/net/eth0/address``

# Set custom hardware address in the /etc/network/interfaces

### Router VM, alpine linux setup script prototypting
Router VM should support combining multiple internet connections

setup-alpine
us
us
router
eth1
dhcp
done
password
utc
none
f
openssh
chrony
vda
lvm
sys

Then go back and add eth0 to the /etc/network/interfaces

# [ROUTER] Setup openssl
  apk update
  apk add ca-certificates
  apk add openssl
# The following didn't work on the newest install, once more information is found, take more detailed notes on what this is for and why.
  update-ca-certificates

# [ROUTER] setup system to make it easy to use
  apk update
  apk add vim

# [ROUTER] install the necessary router tools
  apk add shorewall
  apk add iptables
  apk add ip6tables
  apk add dhcp

  apk add iputils
  apk add iproute2

# [ROUTER] Remove SSH, this should only be managed other ways
  apk del ssh
  !!! didn't do this yet

# [ROUTER] Configure autostart of dependent software. If just iptables is eventually used, which may be desirable. 'rc-update add iptables' is used instead.

  rc-update add shorewall


# _/!\_ TODO: Modify STARTUP_ENABLED=No to "Yes". Required to enable shorewall at startup. 

# Need to do find and replace
in /etc/shorewall/shorewall.conf
set `STARTUP_ENABLED=Yes`


# [HOST] setup bridging information
  mkdir /etc/qemu
  chown root:libvirt /etc/qemu
  cat "allow virbr0" > /etc/qemu/bridge.conf
  cat "allow virbr1" >> /etc/qemu/bridge.conf
  cat "allow virbr2" >> /etc/qemu/bridge.conf
  chown root:libvirt /etc/qemu/
  chown root:libvirt /etc/qemu/bridge.conf

  chown root:libvirt /usr/lib/qemu
  chown root:libvirt-qemu /usr/lib/qemu/qemu-bridge-helper
  chmod 4750 /usr/lib/qemu/qemu-bridge-helper


# [ROUTER] Setup DHCP, Initialize the configuration files
  apk install dhcp

  cp /etc/dhcp/dhcpd.conf.example /etc/dhcp/dhcpd.conf

  cat "auto lo" > /etc/network/interfaces
  cat "iface lo inet loopback" >> /etc/network/interfaces
  cat "" >> /etc/network/interfaces
  cat "auto eth0" >> /etc/network/interfaces
  cat "iface eth0 inet static" >> /etc/network/interfaces
  cat "      address 10.1.1.1" >> /etc/network/interfaces
  cat "      netmask 255.255.255.0" >> /etc/network/interfaces
  cat "      broadcast 10.1.1.255" >> /etc/network/interfaces
  cat "      network 10.1.1.0" >> /etc/network/interfaces
  cat "      gateway 10.1.1.1" >> /etc/network/interfaces
  cat "      up route del -net default gw 10.1.1.1 netmask 0.0.0.0" >> /etc/network/interfaces
  cat "" >> /etc/network/interfaces
  cat "auto eth1" >> /etc/network/interfaces
  cat "iface eth1 inet dhcp" >> /etc/network/interfaces


# [ROUTER] Create the symbolic to relevant config files to make maintenance easier, eventually these should be done from the Multiverse OS persistent share. This way config files can just be quickly deployed and easily edited.

  ln -s /etc/dhcp/dhcpd.conf ~/dhcp-dhcpd.conf
  ln -s /etc/shorewall/hosts ~/shorewall-hosts
  ln -s /etc/shorewall/interfaces ~/shorewall-interfaces
  ln -s /etc/shorewall/masq ~/shorewall-masq
  ln -s /etc/shorewall/policy ~/shorewall-policy
  ln -s /etc/shorewall/rules ~/shorewall-rules
  ln -s /etc/shorewall/zones ~/shorewall-zones
  ln -s /etc/shorewall/shorewall.conf ~/shorewall-shorewall.conf
  ln -s /etc/hosts ~/hosts
  ln -s /etc/motd ~/motd
  ln -s /etc/issue ~/issue
  ln -s /etc/network/interfaces ~/network-interfaces
  ln -s /etc/sysctl.d/00-alpine.conf

# _/!\_ TODO: Gather hardware addresses- better, make them custom, then setup the DNS config and configure static ip addresses for each router

# [ROUTER] Configure DHCP server. Define the DNS server #{DNS_SERVER}, for router.galaxy - it should be a good trusted European privacy non-profit DNS.

# DNS_SERVER=["85.214.20.141","194.150.168.168","213.73.91.35"]

#This ddns-update-style line is required for the routing to work, do not overlook it.
dd-update-style interim;

subnet 192.168.1.0 netmask 255.255.255.0 {
}

subnet 10.1.1.0 netmask 255.255.255.0 {
  option routers 10.1.1.1;
  option subnet-mask 255.255.255.0;
  option broadcast-address 10.1.1.255;
  option domain-name-servers #{DNS_SERVER};
  option domain-name-servers #{DNS_SERVER};
  range 10.1.1.2 10.1.1.254;
}

subnet 10.2.2.0 netmask 255.255.255.0 {
}

host firewall.universe0.mv {
  option host-name "firewall";
  hardware ethernet {MAC_ADDRESS};
  fixed-address 10.2.2.2;
}

cat "127.0.0.1         router.wan.universe0 universe0 localhost.localdomain localhost" > /etc/hosts

# [ROUTER] Update the motd
  cat "Multiverse OS Router  [ router.wan.universe0 ]" > /etc/motd

# [ROUTER] Update the issue
  cat "Multiverse OS Router  [ router.wan.universe0 ]" > /etc/issue
  cat "[ Built using Alpine Linux 3.4 ]" > /etc/issue




# _/!\_ TODO: Configure shorewall, likley best handled with a deploy shared folder that is taken out when complete.


# [HOST] For router.universe the network virsh net-edit files need to be updated to switch the ip address to *.*.*.254 for every network. Eventually it may be best to just unplug entirely. But for now we just rely on not routing that direction, then shorewall disabling all incoming and only allow ssh out. Eventually all talking needs to be moved to VirtIO custom networking.


# [ROUTER] Rebuild the /etc/sysctl.d/00-alpine.conf file for routing and disable ipv6 because right now that is a security hazard until there is better Tor support

cat "net.ipv4.ip_forward = 1" > /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.tcp_syncookies = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.conf.default.rp_filter = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.conf.all.send_redirects = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.conf.all.accept_source_route = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.conf.all.rp_filter = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.ping_group_range=999 59999" >> /etc/sysctl.d/00-alpine.conf
cat "kernel.panic = 120" >> /etc/sysctl.d/00-alpine.conf


rc-update add dhcpd default 

cat "RESOLV_CONF=\"NO\"" >> /etc/udhcpc.conf

# _/!\_ TODO: Use customized hardware addresses to obfuscate the use of QEMU, currently all the 54:52... hardware addresses are dead giveaway to anyone snooping that its a QEMU/kvm setup.


# [ROUTER] Configure shorewall to route traffic
# _/!\_ TODO: Configure shorewall, likley best handled with a deploy shared folder that is taken out when complete.

cat "wan               ipv4" >> /etc/shorewall/zones
cat "lan:wan           ipv4" >> /etc/shorewall/zones
cat "uni               ipv4" >> /etc/shorewall/zones

cat "wan             ipv4" >> /etc/shorewall/zones
cat "lan:wan         ipv4" >> /etc/shorewall/zones
cat "uni:wan         ipv4" >> /etc/shorewall/zones
cat "gal             ipv4" >> /etc/shorewall/zones
cat "vpn             ipv4" >> /etc/shorewall/zones
cat "air:vpn         ipv4" >> /etc/shorewall/zones

cat "uni               eth0:10.1.1.0/24                      -" >> /etc/shorewall/hosts
cat "lan               eth0:192.168.1.0/24                   -" >> /etc/shorewall/hosts

# TODO: _/!\_ eth1 should be whatever is the pci-passthrough. In the case of one of the machines we have 4 network devices. So determine which ones are working then use those as the wan.

cat "wan eth1          routefilter,tcpflags,logmartians,nosmurfs,sourceroute=0" >> /etc/shorewall/interfaces
cat "uni eth0          dhcp,routefilter,tcpflags,logmartians,nosmurfs" >> /etc/shorewall/interfaces


# NOTE: Oddly /etc/shorewall/masq was not included so this needs to be initialized with the prefix comments to match the rest of the files.
touch /etc/shorewall/masq
cat "eth1 0.0.0.0/0" >> /etc/shorewall/masq


cat "fw  all ACCEPT" >> /etc/shorewall/policy
cat "uni lan ACCEPT" >> /etc/shorewall/policy
cat "uni wan ACCEPT" >> /etc/shorewall/policy
cat "all all DROP"   >> /etc/shorewall/policy
#cat "gal vpn ACCEPT" >> /etc/shorewall/policy
#cat "gal lan DROP"   >> /etc/shorewall/policy
#cat "gal uni DROP"   >> /etc/shorewall/policy
#cat "air gal DROP"   >> /etc/shorewall/policy
#cat "all all DROP"   >> /etc/shorewall/policy



cat "##" >> /etc/shorewall/rules
cat "##" >> /etc/shorewall/rules
cat "## Voyager Port Forwarding for LAN Services" >> /etc/shorewall/rules
cat "" >> /etc/shorewall/rules
cat "# Transmission RPC/WebUI" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100    tcp 9091" >> /etc/shorewall/rules
cat "" >> /etc/shorewall/rules
cat "# SSH" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 22" >> /etc/shorewall/rules
cat "" >> /etc/shorewall/rules
cat "# Bittorrent" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 51413" >> /etc/shorewall/rules
cat "" >> /etc/shorewall/rules
cat "# Samba LAN" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 139" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 445" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 137" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 138" >> /etc/shorewall/rules




# [ROUTER] Add a MultiverseOS persistence share. This can later be used to supply config files to simplify the maintence and setup process.

cat "Multiverse /mnt/mv         9p   trans=virtio,9p2000.L,rw,posixacl,cache=none  0 0" >> /etc/fstab
# actually probably won't even do this, gonna find a different way of thing

# TODO: Remove the /dev/cdrom, /dev/usb, and the last drive. 

# TODO: Should disable USB and other kernel modules to limit the attack surface.

# _/!\_ TODO: Setup ephemeral status

# _/!\_ TODO: The template doesn't require a password, but after setup is done, before the ephemeral status is set, then you just need to set the password and store it in the pass-store

# Before shutdown and templating, clear the history, remove the logs, clean the machine for general purpose use. 

history -c 

# Shutdown the machines, take snapshots of each machine to enable ephemeral state, clone each machine for use as a template. Always keep a copy beyond the snapshots, since snapshots are stored inside the qcow2, if the file becomes corrupt, they are also corrupted and do little to help.



### net-firewall
**net-firewall, the firewall**
# It may be much easier to use the finished template-router.unvierse0.mv template to build the net-firewall system. It will remove much of the setup and streamline the process.

**Steps to convert existing router.system.universe0 template.**




----

**Steps to build firewall router from scratch**

clone above drive, then start customizations neccessary for net-firewall

  cat "auto lo" > /etc/network/interfaces
  cat "iface lo inet loopback" >> /etc/network/interfaces
  cat "" >> /etc/network/interfaces
  cat "auto eth0" >> /etc/network/interfaces
  cat "iface eth0 inet static" >> /etc/network/interfaces
  cat "      address 10.2.2.1" >> /etc/network/interfaces
  cat "      netmask 255.255.255.0" >> /etc/network/interfaces
  cat "      broadcast 10.2.2.255" >> /etc/network/interfaces
  cat "      network 10.2.2.0" >> /etc/network/interfaces
  cat "      gateway 10.2.2.1" >> /etc/network/interfaces
  cat "      up route del -net default gw 10.2.2.1 netmask 0.0.0.0" >> /etc/network/interfaces
  cat "" >> /etc/network/interfaces
  cat "auto eth1" >> /etc/network/interfaces
  cat "iface eth1 inet dhcp" >> /etc/network/interfaces



# Define the DNS server #{DNS_SERVER}, for router.galaxy - it should be a good trusted european privacy non-profit DNS
# DNS_SERVER=["85.214.20.141","194.150.168.168","213.73.91.35"]

subnet 10.2.2.0 netmask 255.255.255.0 {
  option routers 10.2.2.1;
  option subnet-mask 255.255.255.0;
  option broadcast-address 10.2.2.255;
  option domain-name-servers #{DNS_SERVER};
  option domain-name-servers #{DNS_SERVER};
  range 10.2.2.2 10.2.2.254;
}

# Prepare openvpn support
apk add openvpn

# _/!\_ TODO: In the special mounted folder for configs, include a folder to drop in openvpn configs. !! Also this server should have 3 options defined by a global VM file, perhaps a class or category or type. This would be OPENVPN, TOR, or SSH. Preferably OPENVPN because it is better to route Tor through some good bitcoin paid VPN to conceal the Tor traffic from your ISP.

cat "127.0.0.1         router.galaxy0 localhost.localdomain localhost" > /etc/hosts


# _/!\_ TODO: Configure shorewall, likley best handled with a deploy shared folder that is taken out when complete.
cat "uni               eth0:10.1.1.0/24                      -" >> /etc/shorewall/hosts
cat "lan               eth0:192.168.1.0/24                   -" >> /etc/shorewall/hosts
cat "air               tun0:10.4.0.0/24                      -" >> /etc/shorewall/hosts

# NOTE: Oddly /etc/shorewall/masq was not included so this needs to be initialized with the prefix comments to match the rest of the files.
cat "wan eth0          routefilter,tcpflags,logmartians,nosmurfs"      >> /etc/shorewall/interfaces
cat "gal eth1          dhcp,routefilter,tcpflags,logmartians,nosmurfs" >> /etc/shorewall/interfaces
cat "vpn tun0          nosmurfs,tcpflags"                              >> /etc/shorewall/interfaces


cat "tun0 0.0.0.0/0" >> /etc/shorewall/masq
cat "eth1 0.0.0.0/0" >> /etc/shorewall/masq


cat "gal vpn ACCEPT" >> /etc/shorewall/policy
cat "gal lan DROP"   >> /etc/shorewall/policy
cat "gal uni DROP"   >> /etc/shorewall/policy
cat "air gal DROP"   >> /etc/shorewall/policy
cat "all all DROP"   >> /etc/shorewall/policy


cat "wan             ipv4" >> /etc/shorewall/zones
cat "lan:wan         ipv4" >> /etc/shorewall/zones
cat "uni:wan         ipv4" >> /etc/shorewall/zones
cat "gal             ipv4" >> /etc/shorewall/zones
cat "vpn             ipv4" >> /etc/shorewall/zones
cat "air:vpn         ipv4" >> /etc/shorewall/zones


# Setup the VPN (This is the most unique part, need to find consitent way to feed in vpn conf file.


# Setup openVPN
# _/!\_ TODO: Allow the user to specify their AirVPN account, then have it log in generate configs and download them in the form of AirVPN.tar.bz2. 
# But in general it should just take a OpenVPN file. 

mkdir ~/AirVPN
tar -xvf ~/AirVPN.tar.bz2
cd ~/AirVPN

cp {OpenVPN.conf} /etc/openvpn/openvpn.conf


rc-update add openvpn default 

# copy configs over

sudo systemctl enable shorewall


**net-tor, router.onion**

Solar is just a whonix box, in the future we can setup a alpine linux box that routes through tor for lighter weight solution but for now whonix works just by dropping it in.

-----




((NEXT STEP))

# _/!\_ TODO: Setup the shorewall configs on the host







-----



#####  proxy server / auxillary VM / sattalite VM / controller VM
**In the future, if you have a existing cluster. You will get a onion address with a code. Or even just like a 15 word phrase. Entering this phrase on new machines would let you download your gpg key and other relevant information for your cluster. Then you can use that to generate onions, and add the new server cluster data to the old to merge.** *This is where we sould generate keys, all derived from a master key but setup in a way that they can be lost without screwing the cluster.*

*Once the portal-gun is working, we can then start adding features for, multi-hop, putting tor. in. Maybe even just like slots you drop in routers between you and the LAN.ls

**Ability to restart router, launch new ones, change VPN, change tor identity, etc should all be built into the gnome GUI.** *It should be provided by rest API, so command line interface would be nice too. This would let i3 interfaces and so on to be easily constructed too.

**proxy VM, controller VM, auxillary VM**
*When completed, the USB drives (rc.local to allow typing in hd pass without needing to dropbear), and the video card passed to the controller/sattalite VM.* **_/!\_** This is very important, Multiverse by default automatically boots into a VM, transparently. Work is all done inside of this VM. The host is hidden and updated by proxy. **In the end remove i3 but for development, leave it in**


Liste frei verwendbarer DNS-Server

Die folgenden Nameserver können als Ersatz der Provider-eigenen Nameserver verwendet werden (zum Zeitpunkt der Erstellung dieser Liste, keine Garantie, wie lange sie frei verfügbar sind)

85.214.20.141 (FoeBud)
204.152.184.76 (f.6to4-servers.net, ISC, USA)
2001:4f8:0:2::14 (f.6to4-servers.net, IPv6, ISC)
194.150.168.168 (dns.as250.net; Berlin/Frankfurt)
213.73.91.35 (dnscache.berlin.ccc.de)


This is where your OpenVPN client will be setup. Ideally you are connecting to a server you rent and put a VPN on yourself or you are using a reputable VPN seller that takes Bitcoin. The Whonix box will route its traffic through this to hide Tor access from your ISP.

Alpine linux does not just pick up the routes, so you need to manually define them in the interfaces file. This is assuming your LAN is 192.168.1.1, which is pretty standard.


# Setup as 10.2.2.100 on router.galaxy0 using the {Hardware_address}

  su
  apt-get update
  apt-get install vim
  apt-get remove nano
  apt-get install sudo


  apt-get install pass


# This is required for setting up a router on this machine to route all the nested Vms.



#TODO: Script: Modify sudoers file

#TODO: Script: update apt/sources.list to stretch


IF AUX SERVER
  rm -rf ~/Music
  rm -rf ~/Pictures
  rm -rf ~/Public
  rm -rf ~/Templates
  rm -rf ~/Videos

[Build the multiverse-os folder structure]
Make the `multiverse-os` folder structure, copy over (or download and validate) *.iso images for Debian and Alpine Linux. In development this is just stored in `/home/user/multiverse-os` but eventually will be stored in /etc/multiverse and ~/.multiverse-os or ~/.local/`multiverse-os`. 

# Setup the local multiverse folder

  mkdir /home/user/.multiverse-os/
  mkdir /home/user/.multiverse-os/portal-gun/persistence/
  mkdir /home/user/.multiverse-os/portal-gun/storage-pool/
  mkdir /home/user/.multiverse-os/portal-gun/os-images/
  mkdir /home/user/.multiverse-os/portal-gun/confs/
  mkdir /home/user/.multiverse-os/scramble-suit/

# During setup the default os-images should be copied
# During setup the default confs should be copied







[Update grub, to define kernel configuration and kernel modules required]
Define the LINUX_DEFAULT line in grub and copy over relevant kernel module configuration. Preform necessary changes to fit the NIC and GPU hardware requirements.

> **Error** Unsupported configuration: host doesn't support pass-through of host PCI devices...

Modify the configuration file depending on CPU type, in `/etc/default/grub`.

````
echo "GRUB_DEFAULT=0"                                    >  /etc/default/grub
echo "GRUB_TIMEOUT=5"                                    >> /etc/default/grub
echo "GRUB_DISTRIBUTOR=Multiverse"                       >> /etc/default/grub
# if cpu.is_a?(:intel)
echo 'GRUB_CMDLINE_LINUX_DEFAULT="quiet intel_iommu=on"' >> /etc/default/grub
# else if cpu.is_a?(:amd)
# echo 'GRUB_CMDLINE_LINUX_DEFAULT="quiet amd_iommu=on"' >> /etc/default/grub
echo 'GRUB_CMDLINE_LINUX=""'                             >> /etc/default/grub
# else
# end
````

Then run the update command to rebuild the grub menu.

`update-grub /etc/default/grub`

> **Error** Unable to complete install: 'internal error: Failed to load PCI stub module vfio-pci'...

If you encounter this error, the *vfio-pci* kernel module has not been properly loaded. There are three common methods to solve this, but the method should be whatever works for your specific system and the requirements you have based on when GPU, USB or NIC need to be loaded/enabled or unloaded/disabled.


*Commonly Suggested* One can modify `initramfs` config or grub config to load this module during startup. The ugly way, and may be loading before it is desirable.

*Sometimes Suggested* One can modify /etc/rc.local, then modify it here or load a script. This file should preferably not be used, but it can be a good fall-back/stopgap until better solutions are found. It also provides different loading timing. It comes after both other methods. This makes it useful for some tasks if video issues are experienced for example.

*A better way* is to modify `/etc/modules-load.d/modules.conf`

````
# /etc/modules: kernel modules to load at boot time.
#
# This file contains the names of kernel modules that should be loaded
# at boot time, one per line. Lines beginning with "#" are ignored.

# Load the necessary modules for PCI pass-through
vfio
vfio_iommu_type1
vfio_pci
````

Adding this does solve the problem. However I did not check if each module used was necessary to specify, options are often dependent upon processor/motherboard combination.

**_/!\_** Check if all three are required.

On hive, `modprobe vfio` turned on vfio_iommu_type1 automatically, only need vfio and vfio_pci in modules.conf for all three to load


[Conditionally required module option: options iommu=pt]
With older motherboards and processors, one will have to sacrifice some security to enable all the features required for Multiverse OS.

> **Error** Unable to complete install: 'internal error: Failed to load PCI stub module vfio-pci'

**Conditionally Required** Older processors motherboard combinations want to add kernel module option: 

`options iommu=pt`

Specifying this kernel option sets the IOMMU into pass-through mode for host devices may sacrifice some I/O. 

> "When an operating system is running inside a virtual machine, including systems that use para-virtualization, such as Xen, it does not usually know the host-physical addresses of memory that it accesses. This makes providing direct access to the computer hardware difficult, because if the guest OS tried to instruct the hardware to perform a direct-memory-access (DMA) using guest-physical addresses, it would likely corrupt the memory, as the hardware does not know about the mapping between the guest-physical and host-physical addresses for the given virtual machine. The corruption is avoided because the hypervisor or host OS intervenes in the I/O operation to apply the translations, incurring a delay in the I/O operation.
> An IOMMU can solve this problem by re-mapping the addresses accessed by the hardware according to the same (or a compatible) translation table that is used to map guest-physical address to host-physical addresses." - [https://en.wikipedia.org/wiki/Input%E2%80%93output_memory_management_unit](Wikipedia Article on Input/Output memory management unit)

*The option reduces the overhead of the IOMMU for host owned devices, but also removes any protection the IOMMU may have provided against errant direct-memory-access (DMA) from devices.  If you weren't using the IOMMU before, there's nothing lost.  Regardless of pass-through mode, the IOMMU will provide the same degree of isolation for assigned devices.

*Enabling specific devices dependent on scanned variables*

Using Go language to scan PCI devices, one can build the necessary rc.local file to prepare pass-through for all the relevant files. *It will be necessary to find all network cards, USB and so on, to disable and prepare for usable by routers and controller VMs.*

If you're using Linux Kernel version 4.1=< or newer, the vfio-pci driver supports the same ids option so you can directly attach devices to vfio-pci and skip pci-stub.  vfio-pci is not generally built statically into the kernel, so we need to force it to be loaded early.To do this on Fedora we need to setup the module options we want to use with `/etc/modprobe.d`.  I typically use a file named `/etc/modprobe.d/local.conf` for local, for example, system specific, configuration. In this case, that file would include:

The default config starts from moving `blacklist.conf`, `dkms.conf`, `kvm.conf`, `vfio.conf` to `/etc/modprobe.d/`.

Then configure the kernel module using a `vfio.conf` placed in `/etc/modprobe.d/` containing the following

`options vfio-pci ids=...`

*Adding the `options vfio-pci ids=...` above will create the other necessary udev rules.* Defining the PCI device IDs, for some systems this only worked for the GPU and other devices such as USB had to be unbound after the hard-drive password was typed in. After the primary drive is open, the USB drives need to be unbound and bound to the *"controller VM"* using the `/etc/rc.local` and the `vfio-bind` script.

Using this combination, any PCI device can be bound and unbound at any step of the boot process allowing for highly customizable Multiverse OS setup.




[Modify the 'kvm' group to permit memory limits]
The 'user' user previously added to the group kvm, now needs to have their memlock expanded to allow virtual machines to be created in the user session instead of the root session.

> **Error** Error starting domain: internal error: Process exited prior to exec: libvirt: error : cannot limit locked memory to 1342177280: Operation not permitted


 24 # error with memory limit issues and not permitted to change it
 25
 26
 27 /etc/security/limits.conf
 28
 29 # One method of solving the issue is providing just the user upgraded memlock limits.
 30
 31 ````
 32 user             hard    memlock         -1
 33 ````
 34
 35 # The preferred way is to modify the kvm group, this will allow for more customizable configuration with multiple users adding virtual machines to the user session.
 36
 37 ````
 38 @kvm             soft    memlock         8650752
 39 @kvm             hard    memlock         -1
 40 ````
 41


[vfio group permissions]

> **Error** Error starting domain: internal error: process exited while connecting to monitor: (Timestamp) qemu-system-x86_64: -device vfio-pci,host=01:00.0,id=hostdev0,bus=pci.0,addr=0x8:vfio error: 0000:01:00.0 group 11 is not viable

> **Error** Error starting domain: internal error: process exited while connecting to monitor: PM.disable_s4=1 -boot strict=on -device ich9-usb-ehci1,id=usb,bus=pci.0,addr=0x5.0x7 -device ich9-usb-uchi1,masterbus=usb.0,firstport...

If you experience the above error, you can confirm the error is related to Linux file system by attempting the following commands under the `root` user
(refer to udev rules for vfio section)

# setup routerb on proxy VM

sudo apt-get install iputils-clockdiff iputils-arping iputils-tracepath iputils-ping

sudo apt-get install shorewall

sudo touch /etc/sysctl.d/00-multiverse.conf

cat "net.ipv4.ip_forward = 1" >> /etc/sysctl.d/00-multiverse.conf
cat "net.ipv4.tcp_syncookies = 1" >> /etc/sysctl.d/00-multiverse.conf
cat "net.ipv4.conf.default.rp_filter = 1" >> /etc/sysctl.d/00-multiverse.conf
cat "net.ipv4.conf.all.send_redirects = 1" >> /etc/sysctl.d/00-multiverse.conf
cat "net.ipv4.conf.all.accept_source_route = 1" >> /etc/sysctl.d/00-multiverse.conf
cat "net.ipv4.conf.all.rp_filter = 1" >> /etc/sysctl.d/00-multiverse.conf
cat "net.ipv4.ping_group_range=999 59999" >> /etc/sysctl.d/00-multiverse.conf
cat "kernel.panic = 120" >> /etc/sysctl.d/00-multiverse.conf

























-----






[Blacklist any extra kernel modules] 

Blacklist absolutely any that are not absolutely needed to limit the attack surface of the bare-metal computer. A file `/etc/modprobe.d/blacklist.conf` is already supplied with the default Multiverse OS install.

````
ipmi_si
ipmi_devintf
````

[Connecting to the xpra server]

If the xpra server is not working, maybe no xpra server found

`xpra  start :13 --no-clipboard --no-pulseaudio --no-daemon`

When debugging the `xpra` server, running on the *"application VM"* to pass windows to the *"controller VM"*, logging into the *"application VM"* with `ssh` and running

`xpra  start :13 --no-randr --disable-mmap --no-clipboard --no-pulseaudio --no-daemon`

Often `stdout` errors will typically reveal any connection issues that are not present when launching `xpra` on the client. Server issues often involve (1) specific server configuration options or more often (2) missing python libraries that need to be installed with the appropriate python pip library manager.

Report any changes in requirements so install scripts can be updated to fit any changes with `xpra` until this tool can be replaced with a Go language or Rust seamless window software.


[Conditionally required kernel module: vfio_iommu_type1]

**Error** Error starting domain: internal error: process exited while connecting to monitor: _PM.disable_s4=1 -boot strict=on -device..

**Error** internal error: process exited while connecting monitor...failed to setup container for group 12: failed to set iommu for container: Operation not permitted


By checking the dmesg with *sudo dmesg | grep iommu* and with the Dell PowerEdge reports:

*vfio_iommu_type1_attach_group: No interrupt remapping IO IOMMU support on this platform.

**It is important to note that this is not required on all machines and should *not* be enabled if it is not necessary.**


((Step Next))

**Portal Gun** This portion of the software should be utilizing a golang application named portal-gun.

*One feature of portal gun should be to grab checksums from the Multiverse DHT. Then it should use those to cross check third party *.iso's like Alpine linux and check third party GPG keys.* It should also be able to hit up the alpine servers and get the newest iso. (since the baremetal should never access the internet. This can only be done after the controller is established.

Setup router.uni VM, pass the NIC over. Setup the router.


----- DIDNT DO THESE STEPS YET ------
((Step Next))

**Scramble Suit** *This needs to be done by leveraging the shifter suit portion of multiverse os. The portions of the install script that don't use [portal-gun] or [scramble-suit] can be written in bash or whatever. But these portions should be using golang applications that will be used actively.*

Generate a GPG key or load a GPG key.

Install pass, use the GPG key above to encrypt the keys.

Use this pass-store to hold all the user, root and HD passwords for the virtual machines in the multiverse-os install.

Use this GPG key to backup the HD key of the bare-metal host.

Use the GPG key to generate an onion address so its consistent.


((Step Next))

**Error** Error starting domain: internal error: /usr/lib/qemu-bridge-helper --user-vnet --br=virbr0 --fd=24: failed to communicate with bridge helper: Transport endpoint is not connected stderr=libvirt: error: cannot execute binary /usr/lib/qemu/qemu-bridge-helper: Permission denied

Somtimes this is an issue with the qemu-bridge-helper permission problems:

  mkdir /etc/qemu
  chown root:libvirt /etc/qemu
  cat "allow virbr0" > /etc/qemu/bridge.conf
  cat "allow virbr1" >> /etc/qemu/bridge.conf
  cat "allow virbr2" >> /etc/qemu/bridge.conf
  chown root:libvirt /etc/qemu/
  chown root:libvirt /etc/qemu/bridge.conf

  chown root:libvirt /usr/lib/qemu
  chown root:libvirt-qemu /usr/lib/qemu/qemu-bridge-helper
  chmod 4750 /usr/lib/qemu/qemu-bridge-helper

 Added to /etc/initramfs-tools/modules:

9p
9pnet
9pnet_virtio





Add os-iamges and storage-pool from ~/.multiverse/portal-gun on the host.

This will allow reuse of images and templates during development. 

In production, it should only be os-images and segregate the host. Maybe read only the templates.Acid





In the Debian package version of shorewall, /etc/shorewall is rather empty. Default configuration files in /usr/share/doc/shorewall/default-config 


sudo systemctl enable isc-dhcp-server

# Need to define a static route using the rfc3442-classless-static-rotues option

The answer is Classless Static Routes (RFC3442). In the isc dhcp server, you have to specify the option manually.

This website states concisely how to do it in a way that works for both windows and linux clients.

Here's the abbreviated version:

Add the following to dhcpd.conf at the top

option rfc3442-classless-static-routes code 121 = array of integer 8;
option ms-classless-static-routes code 249 = array of integer 8;

In the appropriate subnet block add the following two option lines

option rfc3442-classless-static-routes 32, 111, 111, 111, 254, 0, 0, 0, 0, 111, 111, 111, 254;
option ms-classless-static-routes 32, 111, 111, 111, 254, 0, 0, 0, 0, 111, 111, 111, 254;

That should create a static route for 111.111.111.254 on-link with the dhcp assigned address and keep the default router of 111.111.111.254. The special router value 0.0.0.0 means on-link. The rfc states that clients are not required to implement classless static routes, but windows does via their ms option, linux's dhclient (tested debian7, rhel6.4) does and all my IPMI and PXE clients happen to as well. You should test to make sure it works with your clients, but I'm fairly confident it will work. dhclient can be made to interpret option 121 with an exit hook, if it doesn't already support it out of the box.

Worst case, on the host node, you can add an IP that is in the lan range (in your example, 123.123.123.254) and tell the clients to use that as the default gateway.


**None of this worked for *router.solar0* so I ended up just setting the static route in the XML in the VM configuration.**




cat "Multiverse OS [Solar Router]" > /etc/motd

# Setup this with the router.solar0 < the number here is important"
cat "Multiverse router.solar [Alpine Linux 3.5]" > /etc/issue
cat "Kernel \r on an \m (\l)" >> /etc/issue

Finish setting up the router.solar0 router setup so we can drop in servers below it and clone it to grow horizontally


------



((Step Next)

Setup controller VM, this will provide user access or provide a machine to access to control cluster resources. Deep inside the virtual networking the tunnel is opened limiting the visiblity to the local network and to the connecting machine.

((Step Next))

Remove some of the attack surface by turning off some of the internet services that don't need to be turned on for Multiverse OS to work.

*sudo apt-get remove avahi-daemon bind9-host rpcbind exim4*

((Step Next))

Add a firewall, preferably one that can be easily configured, like Shorewall.

((Step 2(b)))

Setup the three isolated networks, remove default network, remove ethernet card if any.

Create user session
sdeb
Create Default Storage for User session

-----

Setup passthrough by enabling ioummu, enabling vfio-pci (vfio is for 4.0+ kerneles and pci-stub is for previous versions).

Then echo the lspci -nn 00:00.00.0 number into unbind. Then pass 0000 0000 number into the vfio-set bind, which creates a /dev/vfio/##.

Before (or after) the networking is setup, qemu must be configured to allow access to the bridges to unprividedged users on the host. This is important because it prevents breakouts from immediately acheiving root access on the host. That you can edit /etc/qemu/bridge.conf to add allow virbr0.

`sudo virsh net-add multiverse` ## this command is broken, and ``virsh net-define`` wants a xml file

Then paste in below, switch out the macs. The DNS is currently routing to google, it would be better to route DNS requets to Tor as it provides a cheap, easy acccess distributed DNS setup. This could be done at the top router level.

I decided on a cosmic naming scheme for my setup. Multiverse OS has 3 network levels:

*Galaxy and Sol could be combined, and a switch could switch between VPN or Tor but never allow direct access to Universe. Multiple second level Firewall/Proxy VMs could be deployed below Universe. Multiple Universe routers could be deployed if the server has multiple physical NIC cards to assign. The initial server has three NIC cards to distribute, but these are all assigned currently to universe0*

Modify the rc.local to automatically do this. I have seen scripts that bind everything, it would be nice to find one that would bind essentially all ethernet,wireless,etc devices.

`/etc/rc.local`

````
echo "0000:03:00.0" > /sys/bus/pci/devices/0000\:03\:00.0/driver/unbind
echo "1969 e091" > /sys/bus/pci/drivers/vfio-pci/new_id

exit 0
````

When I do `lspci -n` I get

03:00.0 0200: 1969:e091 on the same line that doing just `lspci` showed the NIC card. Just match the 03:00.0 part.

One for the PCI pass-through, one for the proxy/firewall VM, and one for Whonix. This will let you have fine grain control over how your VM accesses the Internet.

`/etc/qemu/bridge.conf`


````
allow virbr0
allow virbr1
allow virbr2
````

-----

## Frequently Encountered Errors (FEE)

**Error** Error: Unable to connect to libvirt authentication unavailable no polkit agent available...

**Solution**
This is resolved by correcting the groups of the user trying to start virtual machines in the user session.

`usermod -aG libvirt user`
`usermod -aG kvm user`
`usermod -aG libvirt-qemu user`

**Error** Error: unsupported configuration: host doesn't support pass-through of host PCI devices...

**Solution**
````
echo "GRUB_DEFAULT=0"                                    >  /etc/default/grub
echo "GRUB_TIMEOUT=5"                                    >> /etc/default/grub
echo "GRUB_DISTRIBUTOR=Multiverse"                       >> /etc/default/grub
# if intel processor
echo 'GRUB_CMDLINE_LINUX_DEFAULT="quiet intel_iommu=on"' >> /etc/default/grub
# else if amd processor
# echo 'GRUB_CMDLINE_LINUX_DEFAULT="quiet amd_iommu=on"' >> /etc/default/grub
echo 'GRUB_CMDLINE_LINUX=""'                             >> /etc/default/grub
````

Then run the update command to rebuild the grub menu

`update-grub /etc/default/grub`


> **Error** Error starting domain: internal error: process exited while connecting to monitor: (Timestamp) qemu-system-x86_64: -device vfio-pci,host=01:00.0,id=hostdev0,bus=pci.0,addr=0x8:vfio error: 0000:01:00.0 group 11 is not viable

**Solution**
Doing chmod 660 to /dev/vfio* changed the error, this is properly resolved by fixing the `udev` rules.


> **Error** Error: Unable to complete install: 'internal error: Failed to load PCI stub module vfio-pci'...

**Solution**
So the error message is telling us that the *vfio-pci* kernel module has not been properly loaded. There are three common methods to solve this:

*Commonly Suggested* One can modify `initramfs` config or grub config to load this module during startup. The ugly way, and may be loading before it is desirable.

*Sometimes Suggested* One can modify /etc/rc.local, then modify it here or load a script. This file should preferably not be used, but it can be a good fall-back/stopgap until better solutions are found. It also provides different loading timing. It comes after both other methods. This makes it useful for some tasks if video issues are experienced for example.

*A better way* is to modify `/etc/modules-load.d/`

````
  1 # /etc/modules: kernel modules to load at boot time.
  2 #
  3 # This file contains the names of kernel modules that should be loaded
  4 # at boot time, one per line. Lines beginning with "#" are ignored.
  5
  6 # Load the necessary modules for PCI pass-through
  7 vfio
  8 vfio_iommu_type1
  9 vfio_pci
````

Adding this does solve the problem. However I did not check if each module used was necessary to specify, options are often dependent upon processor/motherboard combination.

**_/!\_** Check if all three are required.



> **Error** Error starting domain: internal error: Process exited prior to exec: libvirt: error : cannot limit locked memory to 1342177280: Operation not permitted

**Solution**
````
 24 # error with memory limit issues and not permitted to change it
 25
 26
 27 /etc/security/limits.conf
 28
 29 Add the following line
 30
 31 ````
 32 user             hard    memlock         -1
 33 ````
 34
 35 I noticed another guide that suggested
 36
 37 ````
 38 @kvm             soft    memlock         8650752
 39 @kvm             hard    memlock         -1
 40 ````
 41
````


> **Error** Error starting domain: internal error: process exited while connecting to monitor: _PM.disable_s4=1 -boot strict=on -device..
> **Error** internal error: process exited while connecting monitor...failed to setup container for group 12: failed to set iommu for container: Operation not permitted

**Solution**
By checking the dmesg with *sudo dmesg | grep iommu* and with the Dell PowerEdge reports:

*vfio_iommu_type1_attach_group: No interrupt remapping IO IOMMU support on this platform.

*It is important to note that this is not required on all machines and should *not* be enabled if it is not necessary.*


> **Error** Error starting domain: internal error: process exited while connecting to monitor: PM.disable_s4=1 -boot strict=on -device ich9-usb-ehci1,id=usb,bus=pci.0,addr=0x5.0x7 -device ich9-usb-uchi1,masterbus=usb.0,firstport...

**Solution**
When managing any file permission within the `/dev/*` folder structure, one should always use the built-in `udevadm`, the `udev` config is found in `/dev/udev/*`, within the `/dev/udev/rules.d/*` folder you can find user specified rules, and in Linux you typically prefix any additions with a number which provides a weight, since folders are organized in alphabetical order by default, lower values are ran first.

For example, if there was a permission issue with a vfio group, then drop in the following rules `/etc/udev/rules.d/10-qemu-hw-users.rules`, the 10 indicating an early `udev` rules load order.


````
# Grant access to /dev/vfio
SUBSYSTEM=="vfio" OWNER="root" GROUP="libvirt"
# Grant access to /dev/usb
#SUBSYSTEM=="usb" OWNER="root" GROUP="libvirt"

````

When a new rule has been added the machine must be either restarted or udev rules reloaded with the following commands

`udevadm control --reload-rules && udevadm trigger`


-----

## Frequently Asked Questions (FAQs)
A collection of frequently asked questions to help new users, developers and anyone interested in piecing together the complexities that make up the Multiverse operating system. If this FAQ section proves to be helpful, it will be broken off into its own wiki with other helpful information about the Multiverse OS project.

**Q. How do you make windows being passed through xpra look better?**
**A.** Almost any window passed with `xpra` is likely to be built with one of two possible GUI frameworks: GTK or QT. Improving default window appearance can be easily done for either of these frameworks by installing additional appearance configuration tools and improving the default settings.

On the client machine, the `xpra` server or in other words the *"application VM"*. 

**TODO: This must be added to `portal-gun`.**

`apt-get install lxappearance gtk-chtheme qt4-qtconfig`

Start with `lxappearance` and choose a theme; then choose it in `gtk-chtheme`. In `qt4-config, there is a drop down menu setting to make qt take the GTK+ settings. That seems to work best for me.

Install Droid Font, maybe others, not exactly sure what i ended up using but this is a nice font. Also ensure to install UTF-8 for broader language support.

# want sound?

Edit the `xpra.conf` on the client `/etc/xpra/xpra.conf`, install the newest version of xpra. probably need to come up with a simple way to build this from source using a script so we do not have to make every machine stretch (testing) Debian instead of the current stable.

a flag will need to be passed to xpra to start with sound

# want clipboard?

It may be best to watch a keystroke and turn off and on clipboard sharing only when its requested. I believe this is how QubesOS does it and makes it significantly more secure even if slightly more annoying to use. this could be optional and you could have just always on too as an option.

4. select which folders should be shared with the machine and if they are read only or not. load everything read only if possible (Default, Mapped (most secure),

Write to `/etc/fstab` file so it auto loads the downloads folder and possibly the development folder


random number generation becomes a problem on template VMs

I installed `rng-tools` and edited `/etc/default/rng` and added `urandom`, but we should generate some entropy with the microphone or some other input.

probably should find a good source of entropy that can be fed to VMs, and at the very least the above has to be in the base template or shit just wont work.

you also want the newest version of xpra for a variety of reasons and its on Debian testing.


3. register applications to be used by this server

so assign applications that are going to be used so a menu can be built and xpra scripts can be generated for that menu.


### General Debian provisioning notes
`apt-get install vim sudo git`

`apt-get remove nano`

#### Remove Gnome games
`apt-get remove -y hitori gnome-chess gnome-tetravex four-in-a-row five-or-more gnome-nibbles gnome-taquin gnome-robots gnome-games`

### would like to control from skylab
-- host0 control
1. Control the VPN being used by the sys-firewall (Add support of `ssh` reverse proxy)
2. Control what is being shared on voyager
3. Control the Whonix firewall, flip identity
4. trigger backup server
#### skylab VM control
(should we do a second VPN on the Skylab box?)
1. switch between OpenVPN and OpenVPN+Whonix connectivity
1. generate app proxy VM
2. control the Whonix firewall, flip identity
3. move Vms around network


-------------- how do we get libvirt to connect to remote shit? ----------------
virsh -c qemu+ssh://user@10.10.10.254/session?socket=/run/user/1000/libvirt/libvirt-sock


# To enable it in virt-manager do:
# gsettings set org.virt-manager.virt-manager.connections uris "['qemu+ssh://user@host0/session?socket=/run/user/1000/libvirt/libvirt-sock', 'qemu:///session', 'qemu:///system']"





## Auxilary/satallite server in the cluster

sudo apt-get install tor

# TODO: Modify the torrc, base the host key on your GPG key! Then that shit is consistent and reliable and regeneraablel

sudo systemctl enable tor

# get the onion hostname

cat /var/lib/tor/hidden_service/hostname

# and backup the private key


##############################################################################################################
==============================================================================================================
# Multiverse Potential Library Notes #########################################################################
##############################################################################################################
### Idea! - Browser where every tab is run in its own virtual machine!!! HAHAHA :D ###########################
##############################################################################################################



**broad general libaries**
https://github.com/Cergoo/gol
https://github.com/jeffizhungry/stack-of-books
**Universal log system**
https://github.com/mhuisi/logg




**networking stack**
https://github.com/joshlf/net

https://github.com/unigornel/go-tcpip



**Db/api**

https://github.com/Fersca/natyla 0 Natyla is a Full-Stack REST-API/Cache/Key-Value-Store application to configure and run simple APIs in minutes. Written in Golang, it provides the same functionality as a multithreaded application running with Memcached and MongoDB.



**interesting languages**
https://github.com/hawx/vodka - A functional, stack-based, concatenative programming language of little practical use. It is based on my earlier Catcon, but this is written in Go and features a few syntax changes.


 **web frame**

https://github.com/goyy/goyy/








 


Go libraries which make work as stop gap to provide basic functionality or provide a starting point to fork and build from.

https://github.com/zwily/turbotunnel - TurboTunnel creates on-demand ssh tunnels. It listens on local ports, starts up ssh connections when something connects to those ports, and proxies data through the remote tunnel.

https://github.com/glycerine/sshego - golang/Go library for ssh tunneling (secure port forwarding)

https://github.com/mcuadros/passage - SSH tunnels on steroids

https://github.com/mgutz/sshtunnel - has an example.

https://github.com/icexin/sshtun - multiple ssh tunnel at once

https://github.com/Kane-Sendgrid/wormhole - SSH proxy tunnel for remote port forwarding indirectly between 2 machines

https://github.com/elentok/gesheft - SSH Tunnel manager, An SSH tunnel manager in GO (based on Shaft by n0nick)

https://github.com/nicksellen/gotunnel - creates ssh tunnels using simple *.yaml configuration

https://github.com/1lann/tport

======= general ssh stuffz =========================================

https://github.com/elentok/gesheft - SSH Tunnel manager gesheft list    - lists all of the tunnels
gesheft active  - lists the active tunnels (removes zombie tunnels)
gesheft info    - shows information about a tunnel

gesheft start   <tunnel_name>
gesheft stop    <tunnel_name>
gesheft restart <tunnel_name>

a clit tool, but clearly would be good for a lib

https://github.com/icexin/sshtun - multiple ssh tunnel at once. has a nice config very basic
https://github.com/segmentio/go-forward-ssh - tunnel stuff can be converted into lib but is cli

https://github.com/srri/easyssh - basic ssh lib, for client stuffs.

https://github.com/BrianBland/warden - gives each ssh a jail shell (based on containers? didnt have time to look)

[KVM/QEMU management]

* *guest-agent* -- https://github.com/0xef53/phoenix-guest-agent
— starting point for vm management software

------------------------------------------------------------------------------------
-- Tools for secure (host <-> virtual machine) connections -------------------------
------------------------------------------------------------------------------------

* nush -- https://github.com/songgao/nush
— Nush can be used to built a custom POSIX shell, this shell can be used to limit
shell access when SSHing between hosts, VMs within the Multiverse OS.

------------------------------------------------------------------------------------
-- KVM/QEMU management UI ----------------------------------------------------------
------------------------------------------------------------------------------------

* Icemenu -- https://github.com/martinlebeda/IceMenu
- This need to be investigated further, but this may be useful for constructing an application
menu of registered applications from each VM.

* kvmtop -- https://github.com/cha87de/kvmtop
- This may be useful to look at the code or provide a very simple view, but this functionality
may already be covered by the primary tool that is being built. but there also does not seem to
be any harm in providing additional small tools.

* vm-manager -- https://github.com/ZeroPage/vm-manager
A very basic management tool that allows VMs to be defined as *.yaml configuration files and
start the machine. The *.yaml code may prove to be useful in simplyifying the process while
providing advanced users an alternative to the eventual GUI tool.

https://github.com/zchee/go-qcow2/blob/master/qcow2.go - best qcow2 lib

— QEMU
https://github.com/digitalocean/go-qemu
https://github.com/quadrifoglio/go-qmp
https://github.com/0xef53/qmp-shell
https://github.com/0xef53/go-qmp





global hosting, financial automation, server-less web components http://dawn-project.com
https://github.com/dawn-network/glogchain


**Improved UI**
https://github.com/ellisonleao/howdoi

**prompt/input for install**
https://github.com/rayrutjes/ask/blob/master/ask.go
**
**


!!!!= remote control for VMs
https://github.com/wingedpig/loom - Loom is a Python Fabric-like library to interact with remote servers using SSH, written in Go.

!!!!!= EPHEMERAL VMS
-- ephemeral vms - going to need some good code for doing snapshots and cloning then starting then deleting or some shit -
|-https://github.com/dguerri/LibVirtKvm-scripts
Some basic scripts that do backing up, snaptshotting, etc
|- https://github.com/ryran/valine - VERY NICE snap shot, reverrt, start shutdown, change media, ,etc

!!!!!= PROVISINING AND SETTING UP VMs to ensure they are secure
#---- Provising/Setting up TEMPLATES, VMS, etc (maybe even host for install)
https://github.com/ares/chef-libvirt.git

!!!!!= GOLANG LIBRARIES FOR INTERACTING WITH LIBVIRT
https://github.com/rgbkrk/libvirt-go -- most mature, using c bindings
|-- https://github.com/qhsong/golang-vm-ci/blob/master/local/virt.go - some example code

https://github.com/digitalocean/go-libvirt -- pure golang

!!!!!= TOR ROUTERS
# Tor is handled by settting up a whonix gateway and using shorewall with a config to route all traffic incomming from the connected nodes getting internet from the router through Tor or a VPN
|- https://github.com/Whonix/whonix-libvirt/tree/master/usr/share/whonix-libvirt/xml -- xml files for whonix, didnt have this before, can improve the xml files when we start writing the xml files with golang instead of virt-manager


https://github.com/NullHypothesis/exitmap

https://github.com/haad/proxychains

https://github.com/moxie0/tortunnel -- single hop onion connections, for like bots or just ip masking and not anonymity


https://github.com/paradoxxxzero/butterfly -- tornado+ws terminal, looks more cooler but less secure than noVPN


https://github.com/hicknhack-software/ansible-libvirt - ansible controleld lib vert
https://github.com/digitalocean/go-libvirt -- well maintained libvert
# containers in containers using alpine
https://www.flockport.com/apps/overlays/

!! recently updated whonix box https://github.com/Whonix/whonix-libvirt
https://github.com/fog/fog-libvirt -- libvirt ruby fog api provider

# containers
  #con tools
   https://github.com/coreos/torus
   https://github.com/anacrolix/torrent
https://github.com/jackpal/Taipei-Torrent

# container libs

  lxd looks good
  lkvm is need  / Clear Containers
  runc is the more secure lxd kind of lxc
  https://github.com/coreos/rkt < secure lkvm kinda thing

  tool to make app armor for dockers on the fly

# runc
   https://github.com/hicknhack-software/ansible-libvirt anisible roles for runc app armor profiles https://github.com/mk-fg/apparmor-profiles/tree/master/profiles

  runc containers on everything (hey also use app armor :)
  https://github.com/jfrazelle/magneto
  https://github.com/vishvananda/netlink
   https://github.com/jfrazelle/netns
  https://github.com/zazab/runcmd -- program runc with golang
  http://kubernetes.io/docs/hellonode/

# sec
  ## jammer
  https://github.com/DanMcInerney/wifijammer
  ## router atker
  https://github.com/jh00nbr/Routerhunter-2.0 - atks routers over the net
  https://github.com/j91321/rext - router exploit kit
  https://github.com/reverse-shell/routersploit
https://github.com/eurialo/lightaidra -- ircd scanner and exploiter

# great selection fo scripts for a router

https://github.com/isislovecruft/scripts

# actual routers
  # https://github.com/grugq/portal -- openwrt a dedicated hardware device (a router) which forces all internet traffic to be sent over the Tor network
  # https://github.com/mojombo/proxymachine - Ruby event machine level 7 tcp router
  # https://github.com/ericpaulbishop/gargoyle - hardcore c router
  # hardcore c++ router os
  # https://github.com/bcoe/smtproutes - ruby email router
  # https://github.com/Juniper/contrail-vrouter - 97% C router

  # simple bsd firewall for the usecase I need ip4 only https://github.com/Netgate/netmap-fwd

# open source router https://github.com/opensourcerouting/quagga
  https://github.com/BIRD/bird -- maintained c router
  https://github.com/shemminger/iproute2

# dot files
  https://github.com/vishvananda/.dotfiles'

## micro kernels
https://os.inf.tu-dresden.de/L4/LinuxOnL4/overview.shtml
L4Linux is a port of the Linux kernel to the L4 microkernel API. It is a para-virtualized Linux kernel running on top of a hypervisor, completely without privileges.

L4Linux runs in user-mode on top of the µ-kernel, side-by-side with other µ-kernel applications such as real-time components.

#bitcoin
https://github.com/lightningnetwork/lightning-onion

# Janus
curl -L https://bit.ly/janus-bootstrap | bash

# Tmux + Janus
git clone https://github.com/adnichols/tmux-and-vim.git ~/.janus
sh ~/.janus/setup/setup.sh

# Plugin Minimalist Plugin Manager Installer
curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
      https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

## VIM .vimrc using Plug
# :PlugInstall to install
call plug#begin()

# VIM Training
# takac/vim-hardtime

# Auto-complete
Plug 'Valloric/YouCompleteMe', { 'do': './install.py' }
Plug 'rdnetto/YCM-Generator', { 'branch': 'stable'}
autocmd! User YouCompleteMe if !has('vim_starting') | call youcompleteme#Enable() | endif

# Buffer UI
Plug 'bling/vim-bufferline'

# Go lang
Plug 'vim-jp/vim-go-extra'
Plug 'benmills/vimux-golang' # Tests

# File Tree
Plug 'scrooloose/nerdtree', { 'on': 'NERDTreeToggle' }
Plug 'Xuyuanp/nerdtree-git-plugin'

# Shell Execution
Plug 'https://github.com/JarrodCTaylor/vim-shell-executor'

# ColorSchemes
Plug 'nanotech/jellybeans.vim'
Plug 'altercation/vim-colors-solarized'
Plug 'sickill/vim-monokai'
Plug 'chriskempson/base16-vim'

# Function/Tag Outline
Plug 'majutsushi/tagbar'

# Git Diffs
Plug 'airblade/vim-gitgutter'
Plug 'mhinz/vim-signify'

# Tabs
Plug 'gcmt/taboo.vim'

# Tmux Integeration
Plug 'benmills/vimux'
Plug 'christoomey/vim-tmux-navigator'
Plug 'ervandew/supertab'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'bling/vim-bufferline'

# Undo UI
Plug 'sjl/gundo.vim'

# Search/Fuzzy UI
Plug 'Shougo/unite.vim'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'vim-ctrlspace/vim-ctrlspace'
Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
call plug#end()


### neat #####
# https://github.com/kisom/cryptutils - cryptutils is a set of common Go packages for doing encryption using NaCl and Ed25519 keys. It also includes a set of command line tools:
    secrets: command-line secrets manager
    otpc: command-line two-factor authentication token manager
    journal: password-backed journal
    passcrypt: password-based file encryption
################
check out siv, its a deterministic encryption, may be useful for tree of keys
#################
deterministic keys for creating nested servers with keys based on parent keys to make regeneration of keys easy
https://github.com/codahale/rfc6979 - A Go implementation of RFC 6979's deterministic DSA/ECDSA signature scheme.
https://github.com/wemeetagain/go-hdwallet - Go utilities for generating hierarchical deterministic Bitcoin wallets based on BIP 32
https://github.com/NebulousLabs/hdkey -
Hierarchical Deterministic Key Derivation for arbitrary secp256k1 key pairs
https://github.com/runeaune/hdkeys -  Hierarchical deterministic key derivation library, as described in BIP-0032
#################
handle where the random source is from
https://github.com/wadey/cryptorand
also consider building an entropy machine

### Cryptocurrency manager
https://github.com/zaibon/cryptogo - very basic starter webui for different cryptocurrencies

#################################
https://github.com/malnick/cryptorious - encrypt, decrypt passwords or random strings with ssh keys, CLI
(would be cool to add git, make it more like pass, and register different backends for backing up. so it replicates offsite.
##################
https://github.com/kiasaki/hotomata - (interesting) guess its like an ansible but all in golang, very interesting
####################
https://github.com/marksheahan/nssh - (VERY FUCKING INTERESTING) does ssh over multiple hops, could be used to obstuficate or even NAT break.
Golang command line ssh utility for running programs remotely over multiple hops
###########################

snow - watcher of files - OSX LINUX andWindows - abstracts on a directory basis, doesnt do any hacky shit

### file systems
ghttps - it will let you mount (read only) any website - https://github.com/sdgoii/ghttpsfs

gorofs - https://github.com/tarm/gorofs - a utility to use emedded zip file (IN THE BINARY :) ) as a backup filesystem. Would make detection of internal files a little more difficult

same thing but with tars is tarstream - https://github.com/benmcclelland/tarstream - golang library to stream files from a filesystem in tar format

aferoboltdb - a filesystem using boltdb as the backend

gistfs - https://github.com/andrewstuart/gistfs - exposes a FS based on a users gist on github. may be useful as an example to make FS out of random web resources.

dear - https://github.com/filwisher/dear - a content addressed filesystem "entaglement" filesystem. its used in a p2p chat system. may be worth looking attached

gitlab-fuse - https://github.com/jonathanreinhart/gitlab-fuse - very simple tool to mount a gitlab using a gitlab API key. may be useful in mounting onion gitlabs.

minfs - https://github.com/zilog8/minfs - a minimal filesystem interface, has several transports, osfs, sftpfs, shimfs, zipfs. this may be the best starting point as it can be easily modified to be a modular interface.

inmemfs - https://github.com/e-tothe-ipi/inmemfs - in memory filesystem running on fuse.

### corruption detection
seal https://github.com/crasm/seal
automated corruption detection without checksum files (or filesystem support)

### backing up
backup-rsync - https://github.com/r0bj/backup-rsync - its a wrapperw ith a nice yaml config file to define rsync backups. it may be a good starting point for a VM backup software and general backing up using a ARCHIVE utility VM

### queue
go-pqueue - https://github.com/ljosa/go-pqueue - go primitive for processing a simple persistent queue backed by the local file syste. it works very similart o using maildir as a queue. it can be fed by any language.

### remote control
https://github.com/cypro666/cmdserver - a http server that accepts commands over http, but has a white list and black list to limit control. seems like a good starting point

### crypto libs
https://github.com/evantbyrne/crypt -- Utility for encrypting and decrypting files with AES-256 GCM and Scrypt. mostly cli tool, but can be modded to a lib

https://github.com/gtank/cryptopasta -- robust, hash (passwords and check), sign (hmac and such), encrypt EDSA and shit, decrypt, no cli, just a lib

https://github.com/dedis/crypto --  Advanced crypto library for the Go language (for serious, lots of specific shit for mimicing shit done like in mining crypto.)

https://github.com/liut/TeaCrypt - A Go library for the Tiny Encryption Algorithm family of ciphers.

https://github.com/phylake/go-crypto -- it can generate rsa (ssh) keys

https://github.com/postfix/cryptobox-ng -- seems like a legit solution for assemtric encryption with nacl in an application, built around making messages that can be transfered between things

### authentication/authorization
https://github.com/N4SJAMK/teamboard-crypt - Hash and compare passwords. and example server

### messaging
https://github.com/sec51/cryptoengine - This Golang package simplifies even further the usage of NaCl crypto primitives/ This simplifies even further the usage of the NaCl crypto primitives, by taking care of the nonce part. It uses a KDF, specifically HKDF to compute the nonces.

Message -> Encrypt -> EncryptedMessage -> ToBytes() -> < = NETWORK = >  <- FromBytes() -> EncryptedMessage -> Decrypt -> Message

### general OS exec stuff 
https://github.com/njones/xcp
Simple Cross Network, Cross OS, Cross Application Copy-Paste

-----------------------------------------------------------
frp - a fast reverse proxy to expose local server behind nat or firewall to internet through an outside server. I had this exxact idea to expose servers deep inside. This would let hosting happen locally but access would happen by proxy over outside cheap 5 dollar a month servers who would forward the specific traffic over. This would make DOS attacks nearly impossible, acheiving essentially the same thing accomplished with CloudFlare (fuck cloudflare and idiots who use them without thinking).

--------------------------
metrics
https://github.com/arodier/graal - this seems like the best option, gives a really nice robust rest API to give out structured information about the vms. Can be modified to add or remove. needs an auth but that can be easily implemented.

https://github.com/owainlewis/gatherd -- Golang Linux system metrics. Library for extracting useful system metrics from Linux systems

https://github.com/fcavani/monitor - Software to monitor other services in a linux server environment, something like monit, but simple and functional.


(may not be needed) subgraph / fw-daemon - https://github.com/subgraph/fw-daemon - sub graph application firewall similar to little snitch on OSX

-- HUGE -- https://github.com/0xef53/phoenix-guest-agent - this will be the basis of host -> VM (included nested) communication. It talks over virtio instead of TCP/IP which means it does not use an http server. It returns JSON responses to a white listed set of commands. There is a fair amount of command started. This will be the best way for clients to communicate and remove SOME security issues found withi the current design.


https://github.com/BenNicholls/bcon - seems itneresinting, lets you register config files under a tag or namespace. Then you can track changes, and easily search against the configs. May be easy to help users mangae things

https://github.com/mariten/json-cat - json output prettified, might be usful to pull out peices for use as lib.

https://github.com/AppliedTrust/mytotp - MyTOTP is a simple Go client for the Time-Based One Time Password (TOTP) protocol. Client software. This could be a great way to implement security between servers. It would also be nice to combine this with pass to create a very feature rich pass/otp manager. If it included accounts, or even identity packs it would be my dream program

https://github.com/yamnikov-oleg/projektor - Fast application launcher for Gnome written in Go. It looks like the original spotlight in OSX. Could be a great way to launch the various programs registered on the VMs. It seems like it may be a good starting point for a UI too for the rest of the VM management software. gtk+3


    Search and launch applications installed on your system
    Navigate through file system, open directories and files
    Execute custom command lines in background
    Open urls in the default web browser

https://github.com/jmoiron/terminal-schemer -  portable style scheme application for mate-terminal and gnome-terminal. This could be very useful in making each gnome terminal for each VM different.

https://github.com/iapazmino/go-notify - a gnome notify lib, so you can hook into the notificaitons


https://github.com/kbinani/screenshot - Go library to capture desktop to image

https://github.com/l3pp4rd/statusbar - linux window manager status bar in pure golang

https://github.com/yaronsumel/grapes - grapes is lightweight tool designed to distribute commands over ssh with ease.

pretty cool, it uses a yaml file to specify servers, allows giving them roles, then defining allowed commands. This could be useful for controlling access.


https://github.com/rapidloop/rtop-vis - rtop-vis can monitor load and memory usage of all the specified servers and visualize the data as a graph with a bit (10 minutes' worth) of history. It connects to servers via SSH and does not need anything to be installed on the servers. The collected data is not persisted. It is lost when rtop-vis is stopped.

https://github.com/tombh/texttop - (FUCKING AWESOME) - A fully interactive X Linux desktop rendered to TTY and streamed over SSH. so like a fucking ASCII remote desktop, its fucking genius.

https://github.com/gravitational/teleport - Gravitational Teleport is a modern SSH server for remotely accessing clusters of Linux servers via SSH or HTTPS. It is intended to be used instead of sshd. Teleport enables teams to easily adopt the best SSH practices like:

    No need to distribute keys: Teleport uses certificate-based access with automatic expiration time.
    Enforcement of 2nd factor authentication.
    Cluster introspection: every Teleport node becomes a part of a cluster and is visible on the Web UI.
    Record and replay SSH sessions for knowledge sharing and auditing purposes.
    Collaboratively troubleshoot issues through session sharing.
    Connect to clusters located behind firewalls without direct Internet access via SSH bastions.
    Ability to integrate SSH credentials with your organization identities via OAuth (Google Apps, Github).


https://github.com/fudanchii/edssh - (VERY INTERSTING) - ed25519 signature support for golang.org/x/crypto/ssh. this is a big improvement over the current rsa system which is likely to already be fucked.



https://github.com/s8sg/goshnix - interesting project. ssh control over multiple hosts through transprent ssh.

https://github.com/samuelngs/universe - Distributed SSH authentication system


https://github.com/peter-edge/osutils-go - OS utilities for Go.


https://github.com/codeskyblue/kexec 0-- This is a golang lib, add a Terminate command to exec.

Tested on windows, linux, darwin.

https://github.com/kless/osutil (HOLY FUCK THIS IS AWESOME) Access to operating system functionality dependent of every platform.

Access to operating system functionality dependent of every platform and utility packages for the Shell.

    config/env: set persistent environment variables
    config/shconf: parser and scanner for the configuration in format shell-variable
    distro: detects the Linux distribution
    file: common operations in files
    pkg: basic operations for the management of packages in operating systems
    sh: interprets a command line like it is done in the Bash shell
    user: provides access to UNIX users database in local files
    user/crypt: password hashing used in UNIX


https://github.com/idcos/osinstall-server/blob/master/src/utils/utils.go - similar but smaller, may be useful to combine the two
https://github.com/kelseyhightower/terminus - very similar, might be nice to pull parts from it
https://github.com/zcalusic/sysinfo - a pure go library providing linux OS / kernel /hardware system information. very similar may also have peices that are useful
----- virus / vuls  scanner -------------------------------------------------------
https://github.com/future-architect/vuls - vulnerabilit scanner. this can be used to scan the machines regularly and notify about security updates. possibility auto update based on config settings.


-- management --------------------

https://github.com/sandyskies/SATools/tree/master/server - it could be a useful starting point for limited commands over a REST API. doesnt look super great but not a bad starting point.

-- cluster ssh ----

https://github.com/elechak/fleet -  Go distributed computing with SSH
    All communication with the interpreters is encrypted using ssh.
    Use a multitude of users across all machines to help manage permissions
    Define machine groups; then easily create a pool of interpreters from those groups.
    Resource manager that determines the best machines on which to start the interpreters.


--- multiple ssh shit management -----------

https://github.com/matthunter/mussh - Multi-SSH is the open source server capable of executing commands via ssh on a group of servers.
    Written in Go
    Rest API to get/post/delete servers, groups, and commands
    Websocket to execute commands. Outputs are sent via the socket as they are available (tail -f works)
    Can specify a tunnel address and base directory for each server
    RethinkDB database

https://github.com/zenhack/spiderproxy -- makes proxy through spiderweb of ssh connections

----- supervisor/monit/bluepill kinda thing ------------------
https://github.com/immortal/immortal -  A *nix cross-platform (OS agnostic) supervisor https://immortal.run



-- ssh manager - https://github.com/trashcan/rain - Rain is a command line tool to store and categorize SSH hosts.
https://github.com/bo0rsh201/borssh -- Simple ssh wrapper that transfers your dot files via ssh according to configuratio

https://github.com/tg123/sshpiper - SSH Piper works as a proxy-like ware, and route connections by username, src ip , etc.

https://github.com/laher/sshutils-go - ssh utilities for go

'Standard-ish' implementations for:

    'Known hosts' file checker, and TODO: adding keys
    Typical client-auth implementations for password, keyring, and ssh-agent

See also:

    scp-go uses sshutils-go for authentication and known-hosts checking.
    someutils bundles scp-go as scp or some scp.


-- https://github.com/yosida95/golang-sshkey - just handles public key stuff


https://github.com/sshconf/sshconf -- config ssh manage functions ,its cli but cna be modifeid into a lib

-- remote connectiosn ------------------------------------------------
https://github.com/natefinch/deputy - deputy is a go package that adds smarts on top of os/exec .

Deputy is a type that runs Commands with advanced options not available from os/exec. See the comments on field values for details.
Run starts the specified command and waits for it to complete. Its behavior conforms to the Options passed to it at construction time.

Note that, like cmd.Run, Deputy.Run should not be used with StdoutPipe or StderrPipe.

https://github.com/toukii/oschat -- protobuf server&client

https://github.com/gliderlabs/sshfront -  Programmable SSH frontend (fairly popular, up to date, minimal but goot starting opoint)

https://github.com/shazow/ssh-chat -  ssh chat

https://github.com/1lann/tport - tport is short for tunnel port. A CLI utility for port forwarding over SSH. Sign in once to your server, forward as many ports in either direction as you need.


https://github.com/dsnet/sshtunnel - This repository contains a simple implementation of a SSH proxy daemon used to securely tunnel TCP connections in forward and reverse proxy mode. This tool provides equivalent functionality to using the ssh command's -L and -R flags.

very nice configs


https://github.com/vdemeester/praetorian - bring security to the ssh server, this lets you configure per key whitelisted commands

https://github.com/artyom/rex == rex executes given command(s) on multiple remote hosts, ssh-connecting to them in parallel.

You're expected to have passwordless access to hosts, rex authenticates itself with the help of ssh-agent that is expected to be running.


--- linux manage------------

=== crons
https://github.com/michaloo/go-cron - Simple golang wrapper over github.com/robfig/cron and os/exec as a cron replacement. its a cli but can be repurposed for lib

=== open
https://github.com/skratchdot/open-golang -  Open a file, directory, or URI using the OS's default application for that object type. Optionally, you can specify an application to use.

-- UI -----------------------------------------------------------------------------
https://github.com/therecipe/qt -- Qt binding for Go (Golang) which supports Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi

https://github.com/grd/go-ipc-example - Go lacks premium-grade gui functionality. The idea is to make a "buddy" gui app, written in any language of choice, and use json-rpc for communication with the app using stdin/stdout pipes.


(QUICK mockingup mockups) https://github.com/alexflint/gallium  - Write desktop applications in Go, HTML, Javascript, and CSS.

Gallium is a Go library for managing windows, menus, dock icons, and desktop notifications. Each window contains a webview component, in which you code your UI in HTML. Under the hood, the webview is running Chromium.

-- monitoring ---------------------------------------------------------------------
https://github.com/mayhewj/jtop - A process monitor for Linux, mostly processes, basic user info, could be a good starting point

-- user management --------------------------------------------------
https://github.com/joonakannisto/determin-ed - (FUCK YES)  Create deterministic ed25519 keys from seedfile and password for openssh-key-v1 format

https://github.com/max107/gouser/blob/master/user.go -- uses the os.exec command, basic, doesnt seem to have group management

-- examples/starting points --------------------------------------------------
https://github.com/fathomdb/applyd - basic controls over a linux debian server, its 3 years old and not documented but it may ahve some things to scavenge



https://github.com/emulbreh/sshub - Start the sshub server:

$ sshub

Configure a tunnel:

$ curl -XPOST http://${SSHUB_HOST}:4080/links/ -d '{
    "port": 12345,
    "from": {
            "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDRuinxi4hANygNImiWn6Jjhn5Wyo1tFzmz+x51wvYUNDIHUIdFeX/51yN27+kMv1yUcLvLcbUio925OVan1kFD4VzCfTJ+TqTS4cT8ZnwbrJFZeewFct1aUZeHBB9ttC1WMsXIAA9ZFyFskyN850axiKyvY8Jy4oDedb08OeWRTi+jPjEolD5e33H4JJygujwJxjpdOlbYN+Ah56CcILJXE4O+m5bxy5Krt/hR84+uqOk2aI+8pPVMQxbABPJjaNJZblK9RHGUGuOVAhhA1dW+0rKWoH2bOt6ODW7vggDG0d0G4VwkPvAEWZpkyDroIkk8tHK/jqf9qDi9UsMibVOd",
            "user": "alice",
    },
    "to": {
            "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCvBNa7e6dJGehmd8KZzgxfrmU/Cyayfd08NpWosT6Je8QNAct+xoU54cT1zYkKnxjME27BG3uF1XGNMW+jZasrh3QJAb8drX2qo65rxhlC5vA7JTQklHkCDiQyOIPtfLGIQCvQQJS3/yjQA59SbFZG4wKS8av8MCS7bW5VP75of9u1T8B8CZAUt3lA+TD6EtYWQFkKJszSOjHbrSLV5PF0QBC+X9kYIXI98ycgOXcXzInssNM7847AtobKNwRqfF83iGkq1C7lMj7dFSpXpUmnvmW41O2cCA/caz1eV1gL/B6JjNBC2FnZC+QtxkMJpi9cPgbqjvLzGEFiQiUNdSf1",
            "user": "bob",
    }
}'

seems pretty sweet, can setup proxies remotely. with an authentication layer this could be useful for on the fly connections
=============================================================================


* push key:

cat ~/.ssh/id_dsa.pub | ssh user1@example.com "cat - >> ~/.ssh/authorized_keys"

# Manpages with colors
export MANPAGER="/usr/bin/most -s"

* Increase preformance

    sudo vim /etc/fstab

Add the noatime and discard options for every SSD partition.

* Put tmp in ram?

tmpfs /tmp tmpfs defaults,size=1g


* Use only SWAP when 99% of RAM is used:

sudo vim /etc/sysctl.conf

# SWAP after 99% RAM used 
vm.swappiness = 1
# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

=============================================================================


# Multiverse OS, image iso with installer

This is a slimed down Debian kernel.

https://github.com/sip-li/debian/tree/master/base-repo

**Options**

*tklx/base* - https://github.com/tklx/base

A super slim root filesystem specifically designed to be used as a base image for application container images.

Based on Debian GNU/Linux, the rootfs.tar.xz weighs in at only 12MB, and 

**What needs to be changed?**


CODENAME=$(cat /etc/debian_codename)

/etc/debian_codename
This will need to be changed and have the versioning updated.

*/etc/apt/sources.list.d/sources.list*
This will need to be updated, we can leverage debian for some packages but primarily should be using a decentralized version.

cat >/etc/apt/sources.list.d/sources.list<<EOF
deb http://httpredir.debian.org/debian $CODENAME main
# deb http://httpredir.debian.org/debian $CODENAME contrib
# deb http://httpredir.debian.org/debian $CODENAME non-free
EOF

if [ "$CODENAME" != "sid" ]; then
cat >/etc/apt/sources.list.d/security.sources.list<<EOF
deb http://security.debian.org/ $CODENAME/updates main
# deb http://security.debian.org/ $CODENAME/updates contrib
# deb http://security.debian.org/ $CODENAME/updates non-free
EOF
fi

*why*?
# prevent initscript from running during install/update
echo "exit 101" > /usr/sbin/policy-rc.d
chmod +x /usr/sbin/policy-rc.d

dpkg-divert --local --rename --add /sbin/initctl
ln -sf /bin/true /sbin/initctl

# aggresively clean apt related files
/usr/local/sbin/apt-clean --aggressive

[NOTE]One of the strengths of the Debian system is the way in which it deals with having multiple similar packages installed, and allowing you to choose which one you use.

We already discussed one way in which Debian allows you to choose a specific package from a number of available options in the piece using the Debian alternatices system.

Another possibility is to replace a binary installed from a package with one of your own choosing, Debian allows this via the use of the dpkg-divert command.

The dpkg-divert command allows you to replace a binary installed upon the system, and have this replacement persist even if you upgrade packages.[/NOTE]






# Some stuff for docs

**Why Debian?**
Tor, Whonix and other established security based operating systems are based on Debian. Allowing work to be shared with other developers working towards the same goals.

Apt Selection: has access to largest GNU/Linux software repository with over 56,800 packages.

Security: all packages are supported with carefully backported security updates that can be safely installed automatically.

Free: Debian is 100% free software. Free from hidden backdoors, free to use, learn from, modify and redistribute.

There are over 1000 passionately committed Debian Developers worldwide. Debian also has the largest ecosystem of derivative distributions. The focus is on development not marketing, which has created a paradoxical situation in which the branding of Debian based distributions such as Ubuntu is better known in certain circles (e.g., commercial industry) than Debian itself.

Seriously, fuck just sucking corporate dick to get money. 

Community: Debian is powered by the world's oldest, largest and most vibrant free software non-profit organization. Debian democratically governs itself using the Debian Social Contract as its constitution.

No central point of failure: many other commercially sponsored Linux distributions have a central point of failure. They are largely dependent on the success and continued independence of their commercial sponsor in a competitive marketplace. As a non-profit organization Debian can not fail in the marketplace or get bought out. Debian has been around for more than 20 years and it will be around in another 20. The same can not be said for other Linux distributions.

Ian hated cops, and thats a correct response to the current state of the completely corrupt US police forces.

Debian is fucking solid, its user friendly and its a great general use operating system. If you don't like it, its probably because you don't understand how customizable linux is. Just switch the package manager out or alias whatever commands annoy you. You don't like systemd because you couldn't be bothered to learn a different style, remove it.

Developers related to this project are inspired by a belief in the power of free software, like science, to promote the progress of a free & humane society.

**Versioning**


Versioning

Releases are based on Semantic Versioning, and use the format of MAJOR.MINOR.PATCH. In a nutshell, the version will be incremented based on the following:

    MAJOR: incompatible and/or major changes, upgraded OS release
    MINOR: backwards-compatible new features and functionality
    PATCH: backwards-compatible bugfixes and package updates


----

# Don't install recommends by default

echo 'APT::Install-Recommends "false";' > /etc/apt/apt.conf.d/05recommends

----

**Lets save space!**

# This is important to decrease the total size of the image and general footprint.
apt-get purge -y --auto-remove

# By default, APT will actually _keep_ packages installed via Recommends or
# Depends if another package Suggests them, even and including if the package
# that originally caused them to be installed is removed.  Setting this to
# "false" ensures that APT is appropriately aggressive about removing the
# packages it added.

# https://aptitude.alioth.debian.org/doc/en/ch02s05s05.html#configApt-AutoRemove-SuggestsImportant

Apt::AutoRemove::SuggestsImportant "false";

# to be as small as possible on-disk, so we explicitly request "gz" versions
# and tell Apt to keep them gzipped on-disk.

# For comparison, an "apt-get update" layer without this on a pristine
# "debian:wheezy" base image was "29.88 MB", where with this it was only
# "8.273 MB".

Acquire::GzipIndexes "true";
Acquire::CompressionTypes::Order:: "gz";


# space by downloading them, and this inhibits that.  For
# users that do need them, it's a simple matter to delete this file and
# "apt-get update". :)

Acquire::Languages "none";


# aggresively clean apt related files
/usr/local/sbin/apt-clean --aggressive

----

The design of the package manager is should be similar this project: http://0install.net/
=============================================================================
##
##
*ANY* VM using QXL, to provide visuals, needs the following xorg config.

Section "Device"
  Identifier "qxl"
  Driver "qxl"
  Option "ENABLE_SURFACES" "False"
EndSection

This will vastly improve the 

**_/!\**  **Checking gpg exit codes only is insufficient. Quote Werner Koch (gnupg lead developer)** https://github.com/Whonix/gpg-bash-lib

*Sets vm.swappiness to the lowest possible value, which is 0. Experience has shown, that in virtual machines have better performance when using this setting.* This only uses swap if its absolutely needed. **NEEDS TO BE DONE ON THE TEST BOX**

*Multiverse OS is an anonymity OS, and it does not attempt to reinvent everything, instead it tries to coexist and utilize the work done by others. 
> "Anonymized operating system user name "user", /etc/hostname, /etc/hosts, /var/lib/dbus/machine-id, which should be shared among all anonymity distributions."

**~/.local/lib/ can be used to store custom udev rules or is used. ~/.local/share/ should be used to match the existing specifications. ~/.cache/**

**SSH connection fail because VM is not on? Well an improved ssh client, console thing could be better if it saw it was a VM in the cluster, then turned it on before connecting if its off, or at least prompt the user to do so.**

**Torrent Utility VM** The torrent utility VM should do auto-tuning. Where it would optimize and test by shifting number of peers, active downloads for HD speed + bandwidth. See: https://quickbox.io/

**For p2p package manager, for files that can't be cross confirmed with the official debian repo. Source code can be downlaoded and compiled. maybe eventually automated analysis. definitely reviews.

**consider looking for a UI for containers, just adapt it to save a ton of time.**

*interest note* Minimal kamailio image with enhancements including running frequently accessed files off of a tmpfs volume. This image uses a custom version of Debian Linux (Jessie) that I designed weighing in at ~22MB compressed.

**!!!** Every Multiverse OS install and installer, live cd etc should use very generic names. user@host. Ideally the home folder should appear empty in a normal linux sense. However when the user logs in, a modified version of ls and other commands, using techniques malware use, to reveal a hidden home folder, map a different name to user@host and so on. However the install still sees its host as hsot, the user could use in their prompt a real name, but abstracted on in the SSH connection? even mapping the hidden file system over the home folder. in this way an attacker gains access gets an empty computer with a single user: user. The accounting system maps its information over the client side connection, presenting the expected view: a home folder with files and username and unique host. This makes it easier for users to know what computer they are using, while still keeping the privacy related to using generic names and hiding files/configs/etc.


**Versioning Idea** Name all the major versions after hackers, politicos, journalists and so on who are in jail, or dead to honor their memories. Better than that toy story shit from vanilla debian.

### file system
brts has sub systems

### Container inside App VM?

Maybe using unprivledged containers like PFlask could separate different applications inside a isoalted App VM. would it be useful at all?
https://github.com/xor-gate/pflask - one cotnainer lib in c

### list of backup software for linux; this is important for the utility VM for backing up
https://github.com/xor-gate/restic-others

https://github.com/restic/restic - niiceeeeee golang

## Controller VM preformance tweaking
Geting some jittering with the new setup, its capable of running high end games but there may be some final tweaking required.

- JS is ridiculous, it blocks the UI of its own websites, its getting used more than ever, its super dangerous because of XSS, it seemingly has no limit on resources usage and people who program it are often novices who are not concerned about resource usage, or at least finish their project and need to be complained to before it even crosses their mind thats a problem to consider. They may never discover it relates to power usage or more importantly: battery drain. This is obvious considering the trend has been to push logic off the server and into the JS, meaning mobile devices use their battery faster than they need to if the design was more concious. That said: javascript usage can be confined and controlled in their amount of total resources allowed creating a hard limit on their ability to drag down a whole client computer. 

* Huge pages could still be used to squeeze out more power

* In the same way the CPUs are isolated, if we can isolate the memory used by the Host machine from the memory of the Controller and two router VMs. Then you remove another major vector for breakout attacks.

* A UI to select CPU features, name the CPU, and control it would be wundervoll, maybe even a full UI showing the virtual computer. Letting the user remove features like power button so for example the host cant turn it off or on. Clicking on the parts then being able to switch out the features. Letting people understand they are basically able to program the physical aspects of their computer.

* Control over cpufreq, maybe like when certain apps are turned on known to have profiles of larger usage, change the settings. When they are turned off, fallback. THis can save energy.

* Reflash bios on every boot, keep stored bios in a hidden drive that has read only access.

* The point is to extend the concept of multiple hop proxies, like with onion routing in Tor. Except with computers in a network, forcing a proxy between, each USB drive, between each Ethernet card, between any direct access to the Controller VM. Allowing ephemeral machines to function as proxies, so if they become infected, the infection is deleted when they are reset.

* The controller should have custom BIOS, custom CPU. Exposing only secure features needed, leaving the rest. Then a custom kernel, only providing features needed.

* Take steps (including testing with hypervisor detection tools) to make the VM undecteable and the HOST not detectable as a hypervisor. Includes (1) change MAC on network cards, (2) Change information on virtual components

* I found a scientific article explaining why Xen clocks are better than other VM software. Then they explained how to make an even better one than Xen. Implement this clock in Multiverse.

* ufw like tool for port forwding, opening ports on VMs, doing tunneling, multihop connecting

* The controller VM should not have any applications, anything that has not been established as a permentant Application VM, should open up in a disposable VM. The disposable VM would be configured similarly to the controller VM, so one could modify files and such, but the system files would all be thrown away when the X was hit. So any program like terminal or anything else.

* UI to visualize the networks. 

* auto wifi hacking- use this: https://github.com/xtr4nge/FruityWifi as a reference, kinda lame its php, make it better 

* [cool unique idea] Arm, and other micro controller VMs, built to mimic standard devices. Possible to insert sensors and so on, capable of feeding it random data. Simplyfing the process of programming robots and make the process cheaper. Once the code is done, you can buy all the pieces you need and have the code ready to run on the assembled bot.
  Also androids and other mobiesl

### Controller VM

Controller VM is setup to get the PCI graphics card the usb slots and so on. It also gets isolated CPUs. 

This not only vastly improves the preformance but it also lets you receive massive security benefits not seen in other VM/Container based compartmentalization operating systems. By telling the Host machine kernel not to schedule the Cores dedicated to the controller and by ensuring you correctly paired the hyperthreads when doing pinning. You ensure that the host and the controller use separate CPUs removing one of the biggest vectors of VM breakout. Exploits that take advantage of running code on the same CPU as the host and leveraging their shared access will not work shrinking the attack surface massively. 

One may ask why even use VMs, the primary reason is because we create a logical shim between the hardware and the virtual machine. Allowing programmatic control over BIOS flashing/updating/configuration, over CPU (you can just the CPUID features you want, leave out dangerous ones (watching CVE you can remove these making your CPU secure despite running a chip with known problems. like leave out management features, but keep all preformance features.) This promatic control over a baremetal machine would require robotic control to do everything one can with Multieverwse controller VM. This fine grain control provides a completely different user experience, now one does not need to wait for Intel, and gets more power than Intel in deciding what features their CPU has and is not forced to use the features sold. This is a powerful concept and fundamentally changes the security of general use computing.

Attacks against the machine would most likely infect the VM BIOS/firmware and not the host. And the BIOS can literally be reflashed on every boot, providing unpresented security. By walling off the HOST machine, ripping out everything extra and giving it no control over the Controller VM, malware in the host BIOS should have difficulty in attacking Multiverse OS.



## Utility VM Ideas
* more general but wget for whole folder?

* Backup server

* Interal DNS, ensure an onion request never leaves the network so it can be accidently sent over the internet and reveal dns information in the failed lookup

* USB proxy vms

* Electronic disobedience, easy set and forget DOS tools. Should apply attacks and check effect, use variety of tools, spit out all the recon it can do automatically and suggest attacks. Routes traffic through multihops to hide origin.

* Tor Non-Exit/Exit Relay

* Automatically convert pdfs to open formats
  - https://github.com/Debian/gpdftext


* Home web hosting, linking up with friends and even freinds of friends from contacts DB. If any site gets a ton of usage, people can share their resources and provide additional hosts, scaling up when use is larger and sclaing down when its not needed. This way people host form home but they also get the benefits of sclaing up in cloud centers. Lets actually ahve the internet not just be hosted by 1 provider creating a giant central point of weakness and essentially breaking the entire concept that makes the internet good

* [important to have early] setup to provide packages for OS upgrade of the host machine

https://github.com/strothj/alpine-debtool 


* Reproducible build VM SERVER not fucking container. Continious integration testing and so on.

## Key Ideas

* No logs on any computer in the cluster BUT the controller, all logs should be output to this server so its all centralized.
https://github.com/clustree/journalbeat-deb
 - https://github.com/mbenkmann/garcon
 - maybe a middleware for caddy!

# https://github.com/systemd/mkosi
**This is critical, in Multiverse, we want to build all the isos for the App and Utility VMs. This would allow for custom building of images. If this is not used it should be the guide. It even supports adding luks and other stuff to the iamge and multiple images.**

Unlike most bullshit like docker or even Qubes, we want the Apps to be distributed WITHOUT fucking prebuilt images that you have to trust. Removal of all trust is critical to security. So the images should be built from auditable buildscripts

*This may be even just useful for building images for VMs inside of Multiverse because it can make a wide variety of custom iso images.*

A fancy wrapper around dnf --installroot, debootstrap, pacstrap and zypper that may generate disk images with a number of bells and whistles.

In theory, any distribution may be used on the host for building images containing any other distribution, as long as the necessary tools are available. Specifically, any distro that packages debootstrap may be used to build Debian or Ubuntu images. Any distro that packages dnf may be used to build Fedora images. Any distro that packages pacstrap may be used to build Arch Linux images. Any distro that packages zypper may be used to build openSUSE images.

Additionally, bootable GPT disk images (as created with the --bootable flag) work when booted directly by EFI systems, for example in KVM via:

*Generated images are legacy-free. This means only GPT disk labels (and no MBR disk labels) are supported, and only systemd based images may be generated. Moreover, for bootable images only EFI systems are supported (not plain MBR/BIOS).*

qemu-kvm -m 512 -smp 2 -bios /usr/share/edk2/ovmf/OVMF_CODE.fd -drive format=raw,file=image.raw



=======================================================================================

**Elastic IPs** a pool of available IPs to assign to computers turning on in a specific subnet.

https://github.com/Whonix/onion-grater [whitelisting filter for dangerous tor control potocol commands] Whitelisting filter for dangerous Tor control protocol commands - https://www.whonix.org/wiki/Dev/CPFP - For example it allows using Tor Browser's New Identity feature on Anonymity Distribution Workstations, fixes Tor Browser's about:tor default homepage and Tor Button status indicator without exposing commands that are dangerous for anonymity. https://www.whonix.org/wiki/Impressum *important because MultiverseOS and oht make use of this, need to be aware of which are dangerous*

**whonix generic make file** template for making make files https://github.com/Whonix/genmkfile

**https://github.com/Whonix/ipv6-disable**

**Modify the Tor Local Browser Page To Add Pezaz**https://github.com/Whonix/whonix-welcome-page

*ANY* VM using QXL, to provide visuals, needs the following xorg config.

Section "Device"
  Identifier "qxl"
  Driver "qxl"
  Option "ENABLE_SURFACES" "False"
EndSection

=======================================================================================
##
** IDK WHERE THIS GOES **
Need to install OVMF: `apt-get install ovmf`
And then when making the VM, specify (in virt-manager new VM dialog, dunno virt-install steps):
New VM:
Connection: QEMU/KVM User session
Local install media

Memory: a lot, I did 8192 MiB, a third of my machine's memory, do what you feel

Enable storage
This is the VM all the other VMs basically run in, so could be good to be fairly large, but always use a disk image (not a whole physical hard drive) because otherwise there's the chance of a VM breakout by an attacker formatting the hard drive so their code can be run on next host reboot. Remember too that further disks can be mounted if the initial disk image gets cramped. Also think about if you are going to want multiple controller vms (want one for gaming? an OSX mv?). I went for 200 GB out of 1 TB


Check "Customize configuration before install"
Stuck with Usermode networking, will change later.

In Overview, set Firmware to "UEFI x86_64: /usr/share/OVMF/OVMF_CODE.fd"
Other customizations (talking about a controller aka all the hardware vm):
Boot Options: check Autostart

Remove unused hardware/input: for example, Tablet, Display Spice, Sound (will pass whole card through later), Console, Channel qemu-ga, Channel spice, Video qxl (will pass whole card), 

Add Hardware:
NVIDIA graphics card bus id
NVIDIA audio bus id
?Intel PCI Express Root Port (the one associated with the video card, see below somewhere for the command to get the IOMMU groups, search 'iommu_groups')
  * Had this as a device, but it was permissssion erroring me. 
  * All three of these devices are sent to /sys/bus/pci/drivers/vfio-pci/new-id using /etc/rc.local or modprobe options


### USB PASSTHROUGH
Put the PCI ids in the /etc/modprobe.d/vfio.conf
In virt-manager, add hardware, choose the IDs. Delete the virtual USB connections.

### SHARED FILE PLACE
virt-manager: Add hardware: Filesystem
Type: mount
Driver: Path
Mode: Mapped
Write Policy: Default
Source path: <enter path here or browse to it: Browse; Browse local>
Target path: multiverse

### SOME NETWORKS
virt-manager, as root: Edit; Connection Details: Virtual Networks
delete default, if it's still there
Add new network:
Name: virbr1
Enable IPv4
Network: 10.2.2.0/24
Enable DHCPv4
Start: 10.2.2.2
End: 10.2.2.254
Isolated

actulaly did a bunch of stuff manually for virbr0

_Everett bridge__
everett0
10.255.255.0/24
(and basically everything else the same)

`virsh net-edit everett0`
deleted the dhcp if you made it earlier

delete the dhcp and ip block from virbr1 and virbr2

this is where you'd delete all the networking frmo the host computer too.

#### IPTABLES
https://unix.stackexchange.com/a/145933
All parts of that answer that say YourIP, put in 10.255.255.100
Also some stuff to drop outgoing ssh and other outgoing connections


#### SSH KEYS
Make ssh keys and move to persistence folder
`ssh-keygen -C user@host -t ed25519 -f /home/user/.local/multiverse-os/portal-gun/persistence/id_ed25519`
Add public key to user's authorized_keys

```
mkdir /home/user/.ssh
cat /home/user/.local/multiverse-os/portal-gun/persistence/id_ed25519.pub >> /home/user/.ssh/authorized_keys
```

>>>>>>>>>> TO DO NEXT >>>>>>>>>>>
Setup router vms

Reboot

Install debian

Need to load iptables at boot (once sure the other networking works)



**Notes**

## NETWORKING
**FIXING METAPHROS**
I'm done with this uni0, gal0, sol0 bullshit. It made no sense, I was trying to hammer this metaphor in and it didn't work with Multiverse metaphor and it was harder. 

Instead your home universe, C137 (if you want to reference rick and morty), is the controller VM. This "controller" term I used turns out is a common term for this type of thing so that stays. But each VM inside of it is a virtual universe, every VM on the same level is a parallel universe. As in a different baremetal machine, each baremetal machine in the cluster is a unvierse, with pocket or virtual universes inside of it. 

So that means the main routers outside of the controller currently named virbr0, virbr1, virbr2, and virbr3 (development bridge) are better termed "Everett Bridges" so its Everett0, Everett1, Everett2, Everett3  and so on. This actually makes a lot more sense in the context of the metaphor since the Everett bridgem is a term used when speculating that wormholes may be a dimensional tear and let you travel between dimensions and since one of these everett bridges is where the interclsuter networking would occur, this makes perfect since. INFACT the interclustering networking should probably just be a router hook to everett1 just like the controller VM is and be on the same 10.2.2.* network. 

**NAMIMG**
So this makes the naming even better and more descriptive:

router.wan-firewall.universe0 (Any LAN and therefore WAN exposed device is numbered, for each device exposed this way.)
router.vpn-firewall.universe0 (Provides the VPN connection and firewall)
router.tor-firewall.universe0

**NEW FINDING**
When reworking my networking to fix some issues I found an even better way to do the networking. Instead of doing isolated networks and being FORCED to use the fucking host as a connected device. One can simply attach each network to the previous network by *Forwarding NAT to virbr1-nic*. This will give the network an ip address on the other network and conennects its network to the other one.




##
=======================================================================================


[Debian image builder]
http://people.linaro.org/~riku.voipio/debian-images/preseed.cfg


Use nspawn containers + debootsrpa to generate Multivere ISo images

https://linux.die.net/man/8/debootstrap
--debian-installer
    Used for internal purposes by the debian-installer
--second-stage
    Complete the bootstrapping process. Other arguments are generally not needed. 


debian-installer


Enter the rootfs chroot and run the second-stage:

    sudo chroot rootfs /bin/bash
    /debootstrap/debootstrap --second-stage

    sudo apt-get install qemu-kvm-extras


  qemu-system-arm \
        -M versatilepb \
        -cpu cortex-a8 \
        -hda rootfs.img \
        -m 256 \
        -kernel vmlinuz \
        -append 'rootwait root=/dev/sda init=/bin/sh rw'

  apt-get install linux-image-versatile


debian-builder 

debian-cd


How:
Create partitions to a loopback file
Format filesystems
Run debootstrap
Install extra packages
Add users and set credentials
Do Hardcoded/Default customizations
Run user specified customizations
Install kernel
Install bootloader


----
   logind.conf, logind.conf.d - Login manager configuration files
      /etc/systemd/logind.conf
       /etc/systemd/logind.conf.d/*.conf
       /run/systemd/logind.conf.d/*.conf
       /usr/lib/systemd/logind.conf.d/*.conf


http://man7.org/linux/man-pages/man1/login.1.html
http://man7.org/linux/man-pages/man7/environ.7.html
http://man7.org/linux/man-pages/man7/locale.7.html
       /etc/security/pam_env.conf
           Default configuration file
       /etc/environment
           Default environment file
       $HOME/.pam_environment
           User specific environment file
       /etc/pam.conf
           the configuration file
       /etc/pam.d
           the Linux-PAM configuration directory. Generally, if this
           directory is present, the /etc/pam.conf file is ignored.
       /usr/lib/pam.d
           the Linux-PAM vendor configuration directory. Files in /etc/pam.d
           override files with the same name in this directory.

       pam_unix - Module for traditional password authentication
       pam_unix.so
  unix_chkpwd is a helper program for the pam_unix module that verifies
       the password of the current user.
       pam_systemd - Register user sessions in the systemd login manager

SYNOPSIS
       pam_systemd.so
DESCRIPTION
       pam_systemd registers user sessions with the systemd login manager
       systemd-logind.service(8), and hence the systemd control group
       hierarchy.

systemd-logind.service **HANDLES MULTISEAT MANAGEMNT**
       /usr/lib/systemd/systemd-logind

=======================================================================================
# Multiverse Host Kernel Networking bypass



## template-router.wan.universe0
##
<domain type='kvm'>
  <name>template-router.universe0.mv</name>
  <uuid>48d4f548-eef1-4818-ba58-6388f29168dd</uuid>
  <memory unit='KiB'>524288</memory>
  <currentMemory unit='KiB'>524288</currentMemory>
  <vcpu placement='static'>1</vcpu>
  <os>
    <type arch='x86_64' machine='pc-i440fx-2.6'>hvm</type>
  </os>
  <features>
    <acpi/>
    <apic/>
    <vmport state='off'/>
  </features>
  <cpu mode='custom' match='exact'>
    <model fallback='allow'>Haswell</model>
  </cpu>
  <clock offset='utc'>
    <timer name='rtc' tickpolicy='catchup'/>
    <timer name='pit' tickpolicy='delay'/>
    <timer name='hpet' present='no'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>restart</on_crash>
  <pm>
    <suspend-to-mem enabled='no'/>
    <suspend-to-disk enabled='no'/>
  </pm>
  <devices>
    <emulator>/usr/bin/kvm</emulator>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2'/>
      <source file='/home/user/.local/share/libvirt/images/universe.router.mv-clone.qcow2'/>
      <target dev='hda' bus='ide'/>
      <boot order='2'/>
      <address type='drive' controller='0' bus='0' target='0' unit='0'/>
    </disk>
    <controller type='usb' index='0' model='ich9-ehci1'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x7'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci1'>
      <master startport='0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0' multifunction='on'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci2'>
      <master startport='2'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x1'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci3'>
      <master startport='4'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x2'/>
    </controller>
    <controller type='virtio-serial' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'/>
    <controller type='ide' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
    </controller>
    <interface type='bridge'>
      <mac address='52:54:00:92:ba:73'/>
      <source bridge='virbr0'/>
      <model type='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <serial type='pty'>
      <target port='0'/>
    </serial>
    <console type='pty'>
      <target type='serial' port='0'/>
    </console>
    <channel type='spicevmc'>
      <target type='virtio' name='com.redhat.spice.0'/>
      <address type='virtio-serial' controller='0' bus='0' port='1'/>
    </channel>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='spice' autoport='yes'>
      <listen type='address'/>
      <image compression='off'/>
    </graphics>
    <video>
      <model type='qxl' ram='65536' vram='65536' vgamem='16384' heads='1' primary='yes'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <hostdev mode='subsystem' type='pci' managed='yes'>
      <source>
        <address domain='0x0000' bus='0x03' slot='0x00' function='0x0'/>
      </source>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x07' function='0x0'/>
    </hostdev>
    <hostdev mode='subsystem' type='pci' managed='yes'>
      <source>
        <address domain='0x0000' bus='0x04' slot='0x00' function='0x0'/>
      </source>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x08' function='0x0'/>
    </hostdev>
    <memballoon model='virtio'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </memballoon>
  </devices>
</domain>

## template-router.firewall.universe0
##
<domain type='kvm'>
  <name>template-firewall.universe0.mv</name>
  <uuid>56e15fbf-2221-4134-ade2-216bc621531c</uuid>
  <memory unit='KiB'>524288</memory>
  <currentMemory unit='KiB'>524288</currentMemory>
  <vcpu placement='static'>1</vcpu>
  <os>
    <type arch='x86_64' machine='pc-i440fx-2.6'>hvm</type>
  </os>
  <features>
    <acpi/>
    <apic/>
    <vmport state='off'/>
  </features>
  <cpu mode='custom' match='exact'>
    <model fallback='allow'>Haswell</model>
  </cpu>
  <clock offset='utc'>
    <timer name='rtc' tickpolicy='catchup'/>
    <timer name='pit' tickpolicy='delay'/>
    <timer name='hpet' present='no'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>restart</on_crash>
  <pm>
    <suspend-to-mem enabled='no'/>
    <suspend-to-disk enabled='no'/>
  </pm>
  <devices>
    <emulator>/usr/bin/kvm</emulator>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2'/>
      <source file='/home/user/.local/share/libvirt/images/universe.router.mv-clone-clone-clone.qcow2'/>
      <target dev='hda' bus='ide'/>
      <boot order='2'/>
      <address type='drive' controller='0' bus='0' target='0' unit='0'/>
    </disk>
    <controller type='usb' index='0' model='ich9-ehci1'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x7'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci1'>
      <master startport='0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0' multifunction='on'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci2'>
      <master startport='2'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x1'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci3'>
      <master startport='4'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x2'/>
    </controller>
    <controller type='virtio-serial' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'/>
    <controller type='ide' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
    </controller>
    <filesystem type='mount' accessmode='mapped'>
      <source dir='/home/user/multiverse-os/persistent/firewall.universe0'/>
      <target dir='MultiverseOS'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x08' function='0x0'/>
    </filesystem>
    <interface type='bridge'>
      <mac address='52:54:00:33:7d:c3'/>
      <source bridge='virbr0'/>
      <model type='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <interface type='bridge'>
      <mac address='52:54:00:fa:09:e2'/>
      <source bridge='virbr1'/>
      <model type='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x07' function='0x0'/>
    </interface>
    <serial type='pty'>
      <target port='0'/>
    </serial>
    <console type='pty'>
      <target type='serial' port='0'/>
    </console>
    <channel type='spicevmc'>
      <target type='virtio' name='com.redhat.spice.0'/>
      <address type='virtio-serial' controller='0' bus='0' port='1'/>
    </channel>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='spice' autoport='yes'>
      <listen type='address'/>
      <image compression='off'/>
    </graphics>
    <video>
      <model type='qxl' ram='65536' vram='65536' vgamem='16384' heads='1' primary='yes'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <memballoon model='virtio'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </memballoon>
  </devices>
</domain>

##
##
=======================================================================================
=======================================================================================
# Multiverse OS Alpha Installer Build
https://github.com/aelsabbahy/goss
https://github.com/muesli/cache2go
https://github.com/dgraph-io/dgraph
https://github.com/siddontang/ledisdb
https://github.com/syndtr/goleveldb <very nice>
https://github.com/spf13/afero
## httpfs
https://github.com/cznic/httpfs
## simple cotnainer
https://github.com/converseai/simple_container
!!!!!!!!!!!!!!

USES FUCKING MRUBYT O CONFIG CONTAINERS (DOCKER BUT WHO GIVES A FUCK!)
!!!!!!!!!!!!!!!!!!!
BEST SYS INFO https://github.com/zcalusic/sysinfo
https://github.com/u-root/u-root/tree/6fd12df7e9d96fce0acb7ab723b7f129e5b07107/pkg/pci
https://github.com/AdaptiveScale/lxdui
https://github.com/hackwave/goiardi CHEF IN GO
https://github.com/jdmelo/libvirt-go-xml
https://github.com/jdmelo/libvirt-go-xml
https://github.com/twitchyliquid64/bob-the-builder
### Mk devices
if err := syscall.Mkfifo(f, 0700); err != nil && !os.IsExist(err) {
			return nil, fmt.Errorf("mkfifo: %s %v", f, err)
}

if err := syscall.Mkfifo(outPipe, 0644); err != nil {
		return nil, nil, err
}

if err := syscall.Mkfifo(path, syscall.S_IFIFO|0666); err != nil {
		logger.Log.Fatalf("%v", err)
}


const (
	nodev    = unix.MS_NODEV
	noexec   = unix.MS_NOEXEC
	nosuid   = unix.MS_NOSUID
	readonly = unix.MS_RDONLY
	rec      = unix.MS_REC
	relatime = unix.MS_RELATIME
	remount  = unix.MS_REMOUNT
	shared   = unix.MS_SHARED
)

var (
	rliminf  = unix.RLIM_INFINITY
	infinity = uint64(rliminf)
)

// set as a subreaper
func subreaper() {
	err := unix.Prctl(unix.PR_SET_CHILD_SUBREAPER, uintptr(1), 0, 0, 0)
	if err != nil {
		log.Printf("error setting as a subreaper: %v", err)
	}
}

// nothing really to error to, so just warn
func mount(source string, target string, fstype string, flags uintptr, data string) {
	err := unix.Mount(source, target, fstype, flags, data)
	if err != nil {
		log.Printf("error mounting %s to %s: %v", source, target, err)
	}
}

// in some cases, do not even log an error
func mountSilent(source string, target string, fstype string, flags uintptr, data string) {
	_ = unix.Mount(source, target, fstype, flags, data)
}

## Make dev and nodes

		// Mknod creates a filesystem node (file, device special file or named pipe) named path
		// with attributes specified by mode and dev.
		func Mknod(path string, mode uint32, dev int) error {
			return syscall.Mknod(path, mode, dev)
		}

		// Mkdev is used to build the value of linux devices (in /dev/) which specifies major
		// and minor number of the newly created device special file.
		// Linux device nodes are a bit weird due to backwards compat with 16 bit device nodes.
		// They are, from low to high: the lower 8 bits of the minor, then 12 bits of the major,
		// then the top 12 bits of the minor.
		func Mkdev(major int64, minor int64) uint32 {
			return uint32(((minor & 0xfff00) << 12) | ((major & 0xfff) << 8) | (minor & 0xff))
		}


## Make char device by chaining mknod with mkdev
		// make a character device
		func mkchar(path string, mode, major, minor uint32) {
			// unix.Mknod only supports int dev numbers; this is ok for us
			dev := int(unix.Mkdev(major, minor))
			err := unix.Mknod(path, mode, dev)
			if err != nil {
				if err.Error() == "file exists" {
					return
				}
				log.Printf("error making device %s: %v", path, err)
			}
		}

### cmd line
cat /proc/modules | grep hello

cat /proc/cmdline 
to see what things are loaded at boot
--

https://github.com/fabric8io/kansible

https://github.com/1lann/cete kvdb badger
https://github.com/bmeg/arachne badger graph

**Experiment with using btrfs as the multipath networ drive between Multiverse OS virtual machines**
https://github.com/manifoldco/promptui
https://github.com/tyler-smith/go-bip32 - KEY SYSTEM

https://github.com/NVZGUL/NetInterfaceApi
https://github.com/PouuleT/exec-in-net - conceptually important
https://github.com/samalba/buse-go

terminus for hardware stats

https://github.com/mars9/ramfs - 9p but all in memory!

https://github.com/vasi/qcow2
https://github.com/briandowns/aion - rest server cron

############ https://github.com/lxc/distrobuilder

https://github.com/pfactum/xk MACROKERNEL FROM C+KVM VMs

https://github.com/bootchk/rustDevContainers VAGGA RUST CONTAINER

https://github.com/cyphar/initrs RUST INIT FOR CONTIANERS

https://github.com/gsora/lmod KERNEL MODUELS

https://github.com/mgoltzsche/cntnr

https://github.com/huwwynnjones/oci_rs
https://github.com/oracle/railcar
https://github.com/utaal/shipc RUST CONTAINER WTIH ROOTLESS

https://github.com/lastbackend/lastbackend WEBUI
https://github.com/ttrahan/deploy-kubernetes-runcli WEBUI
## WebUI 
https://github.com/abcum/webkit/ - webkit that supports easy ability to disable javascript execution!!!

https://github.com/alice02/runc-with-network

https://github.com/vburenin/consmgr manage console sockets

https://github.com/polydawn/repeatr < timeless >

https://github.com/AkihiroSuda/runrootless
https://github.com/mgoltzsche/cntnr rootless




https://github.com/caseyr003/terraform-oci-workshop
https://github.com/opencontainers/image-tools
https://github.com/anuvu/stacker
https://github.com/openSUSE/umoci
https://github.com/projectatomic/buildah
https://github.com/clearcontainers/osbuilder
https://github.com/containers/image
https://github.com/containers/build

https://github.com/hashicorp/go-reap

IMMUTABLE RADIX https://github.com/hashicorp/go-immutable-radix
https://github.com/hashicorp/go-memdb made with immutable radix
https://github.com/hashicorp/raft-mdb


SOCKADDR https://github.com/hashicorp/go-sockaddr

MEMBERLIST https://github.com/hashicorp/memberlist


RAFT https://github.com/hashicorp/raft
https://github.com/hashicorp/raft-boltdb


PROVISION https://github.com/hashicorp/serf

DBUS https://github.com/guelfey/go.dbus


<transparent proxy>https://github.com/elazarl/goproxy

<BIOS>https://github.com/coreos/seabios

https://github.com/coreos/baselayout

https://github.com/coreos/bcrypt-tool
<IMPORTANT>https://github.com/coreos/go-iptables
<IMPORTANT> https://raw.githubusercontent.com/coreos/baselayout/master/modprobe.d/aliases.conf
https://github.com/jessfraz/netns

func OpenTun(name string) (*os.File, string, error) {
	tun, err := os.OpenFile(tunDevice, os.O_RDWR, 0)
	if err != nil {
		return nil, "", err
	}

	var ifr ifreqFlags
	copy(ifr.IfrnName[:len(ifr.IfrnName)-1], []byte(name+"\000"))
	ifr.IfruFlags = syscall.IFF_TUN | syscall.IFF_NO_PI

	err = ioctl(int(tun.Fd()), syscall.TUNSETIFF, uintptr(unsafe.Pointer(&ifr)))
	if err != nil {
		return nil, "", err
	}

	ifname := fromZeroTerm(ifr.IfrnName[:ifnameSize])
	return tun, ifname, nil
}

https://github.com/coreos/flannel/blob/master/pkg/ip/ipnet.go
Ip marshalling 


// Mkdev returns a Linux device number generated from the given major and minor
// components.
func Mkdev(major, minor uint32) uint64 {
	dev := (uint64(major) & 0x00000fff) << 8
	dev |= (uint64(major) & 0xfffff000) << 32
	dev |= (uint64(minor) & 0x000000ff) << 0
	dev |= (uint64(minor) & 0xffffff00) << 12
	return dev
}

func Pipe(p []int) (err error) {
	if len(p) != 2 {
		return EINVAL
	}
	p[0], p[1], err = pipe()
	return
}


func Mkfifo(path string, mode uint32) error {
	return Mknod(path, mode|S_IFIFO, 0)
}

https://github.com/golang/sys/blob/master/unix/syscall_linux.go
// BindToDevice binds the socket associated with fd to device.
func BindToDevice(fd int, device string) (err error) {
	return SetsockoptString(fd, SOL_SOCKET, SO_BINDTODEVICE, device)
}

func PtracePeekText(pid int, addr uintptr, out []byte) (count int, err error) {
	return ptracePeek(PTRACE_PEEKTEXT, pid, addr, out)
}

func PtracePeekData(pid int, addr uintptr, out []byte) (count int, err error) {
	return ptracePeek(PTRACE_PEEKDATA, pid, addr, out)
}

func PtracePeekUser(pid int, addr uintptr, out []byte) (count int, err error) {
	return ptracePeek(PTRACE_PEEKUSR, pid, addr, out)
}

func PtracePokeText(pid int, addr uintptr, data []byte) (count int, err error) {
	return ptracePoke(PTRACE_POKETEXT, PTRACE_PEEKTEXT, pid, addr, data)
}

func PtracePokeData(pid int, addr uintptr, data []byte) (count int, err error) {
	return ptracePoke(PTRACE_POKEDATA, PTRACE_PEEKDATA, pid, addr, data)
}

func PtracePokeUser(pid int, addr uintptr, data []byte) (count int, err error) {
	return ptracePoke(PTRACE_POKEUSR, PTRACE_PEEKUSR, pid, addr, data)
}

func PtraceGetRegs(pid int, regsout *PtraceRegs) (err error) {
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

func PtraceSetRegs(pid int, regs *PtraceRegs) (err error) {
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}

func PtraceSetOptions(pid int, options int) (err error) {
	return ptrace(PTRACE_SETOPTIONS, pid, 0, uintptr(options))
}

func PtraceGetEventMsg(pid int) (msg uint, err error) {
	var data _C_long
	err = ptrace(PTRACE_GETEVENTMSG, pid, 0, uintptr(unsafe.Pointer(&data)))
	msg = uint(data)
	return
}

func PtraceCont(pid int, signal int) (err error) {
	return ptrace(PTRACE_CONT, pid, 0, uintptr(signal))
}

func PtraceSyscall(pid int, signal int) (err error) {
	return ptrace(PTRACE_SYSCALL, pid, 0, uintptr(signal))
}

func PtraceSingleStep(pid int) (err error) { return ptrace(PTRACE_SINGLESTEP, pid, 0, 0) }

func PtraceAttach(pid int) (err error) { return ptrace(PTRACE_ATTACH, pid, 0, 0) }

func PtraceDetach(pid int) (err error) { return ptrace(PTRACE_DETACH, pid, 0, 0) }

//sys	reboot(magic1 uint, magic2 uint, cmd int, arg string) (err error)

func Reboot(cmd int) (err error) {
	return reboot(LINUX_REBOOT_MAGIC1, LINUX_REBOOT_MAGIC2, cmd, "")
}

func ReadDirent(fd int, buf []byte) (n int, err error) {
	return Getdents(fd, buf)
}

//sys	mount(source string, target string, fstype string, flags uintptr, data *byte) (err error)

func Mount(source string, target string, fstype string, flags uintptr, data string) (err error) {
	// Certain file systems get rather angry and EINVAL if you give
	// them an empty string of data, rather than NULL.
	if data == "" {
		return mount(source, target, fstype, flags, nil)
	}
	datap, err := BytePtrFromString(data)
	if err != nil {
		return err
	}
	return mount(source, target, fstype, flags, datap)
}

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error) {
	var msg Msghdr
	var rsa RawSockaddrAny
	msg.Name = (*byte)(unsafe.Pointer(&rsa))
	msg.Namelen = uint32(SizeofSockaddrAny)
	var iov Iovec
	if len(p) > 0 {
		iov.Base = &p[0]
		iov.SetLen(len(p))
	}
	var dummy byte
	if len(oob) > 0 {
		var sockType int
		sockType, err = GetsockoptInt(fd, SOL_SOCKET, SO_TYPE)
		if err != nil {
			return
		}
		// receive at least one normal byte
		if sockType != SOCK_DGRAM && len(p) == 0 {
			iov.Base = &dummy
			iov.SetLen(1)
		}
		msg.Control = &oob[0]
		msg.SetControllen(len(oob))
	}
	msg.Iov = &iov
	msg.Iovlen = 1
	if n, err = recvmsg(fd, &msg, flags); err != nil {
		return
	}
	oobn = int(msg.Controllen)
	recvflags = int(msg.Flags)
	// source address is only specified if the socket is unconnected
	if rsa.Addr.Family != AF_UNSPEC {
		from, err = anyToSockaddr(&rsa)
	}
	return
}

func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error) {
	_, err = SendmsgN(fd, p, oob, to, flags)
	return
}

func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error) {
	var ptr unsafe.Pointer
	var salen _Socklen
	if to != nil {
		var err error
		ptr, salen, err = to.sockaddr()
		if err != nil {
			return 0, err
		}
	}
	var msg Msghdr
	msg.Name = (*byte)(ptr)
	msg.Namelen = uint32(salen)
	var iov Iovec
	if len(p) > 0 {
		iov.Base = &p[0]
		iov.SetLen(len(p))
	}
	var dummy byte
	if len(oob) > 0 {
		var sockType int
		sockType, err = GetsockoptInt(fd, SOL_SOCKET, SO_TYPE)
		if err != nil {
			return 0, err
		}
		// send at least one normal byte
		if sockType != SOCK_DGRAM && len(p) == 0 {
			iov.Base = &dummy
			iov.SetLen(1)
		}
		msg.Control = &oob[0]
		msg.SetControllen(len(oob))
	}
	msg.Iov = &iov
	msg.Iovlen = 1
	if n, err = sendmsg(fd, &msg, flags); err != nil {
		return 0, err
	}
	if len(oob) > 0 && len(p) == 0 {
		n = 0
	}
	return n, nil
}

----
[Debian image builder]
http://people.linaro.org/~riku.voipio/debian-images/preseed.cfg


Use nspawn containers + debootsrpa to generate Multivere ISo images

https://linux.die.net/man/8/debootstrap
--debian-installer
    Used for internal purposes by the debian-installer
--second-stage
    Complete the bootstrapping process. Other arguments are generally not needed. 


debian-installer


Enter the rootfs chroot and run the second-stage:

    sudo chroot rootfs /bin/bash
    /debootstrap/debootstrap --second-stage

    sudo apt-get install qemu-kvm-extras


  qemu-system-arm \
        -M versatilepb \
        -cpu cortex-a8 \
        -hda rootfs.img \
        -m 256 \
        -kernel vmlinuz \
        -append 'rootwait root=/dev/sda init=/bin/sh rw'

  apt-get install linux-image-versatile


debian-builder 

debian-cd


How:
Create partitions to a loopback file
Format filesystems
Run debootstrap
Install extra packages
Add users and set credentials
Do Hardcoded/Default customizations
Run user specified customizations
Install kernel
Install bootloader


----
   logind.conf, logind.conf.d - Login manager configuration files
      /etc/systemd/logind.conf
       /etc/systemd/logind.conf.d/*.conf
       /run/systemd/logind.conf.d/*.conf
       /usr/lib/systemd/logind.conf.d/*.conf


http://man7.org/linux/man-pages/man1/login.1.html
http://man7.org/linux/man-pages/man7/environ.7.html
http://man7.org/linux/man-pages/man7/locale.7.html
       /etc/security/pam_env.conf
           Default configuration file
       /etc/environment
           Default environment file
       $HOME/.pam_environment
           User specific environment file
       /etc/pam.conf
           the configuration file
       /etc/pam.d
           the Linux-PAM configuration directory. Generally, if this
           directory is present, the /etc/pam.conf file is ignored.
       /usr/lib/pam.d
           the Linux-PAM vendor configuration directory. Files in /etc/pam.d
           override files with the same name in this directory.

       pam_unix - Module for traditional password authentication
       pam_unix.so
  unix_chkpwd is a helper program for the pam_unix module that verifies
       the password of the current user.
       pam_systemd - Register user sessions in the systemd login manager

SYNOPSIS
       pam_systemd.so
DESCRIPTION
       pam_systemd registers user sessions with the systemd login manager
       systemd-logind.service(8), and hence the systemd control group
       hierarchy.

systemd-logind.service **HANDLES MULTISEAT MANAGEMNT**
       /usr/lib/systemd/systemd-logind


#### Multiverse Installer
**Host Machine**

[Host LUKS Encryption]
  [Encryption: Needed for Install]
  (options, both rely on C. So maybe just use C version?)
  https://github.com/martinjungblut/cryptsetup
  https://github.com/kcolford/go-cryptsetup
  [Decryption]
  https://github.com/jamesrr39/cryptsetup-helper/blob/master/src/dmcrypthelper/cryptdevice.go

#### Multipath NFS
Special mount options in newer Linux Kernels allow for multipath mounting
     -R, --rbind
              Remount a subtree and all possible submounts somewhere else
              (so that its contents are available in both places).  See
              above, the subsection Bind mounts.
    -F, --fork
              (Used in conjunction with -a.)  Fork off a new incarnation of
              mount for each device.  This will do the mounts on different
              devices or different NFS servers in parallel.  This has the
              advantage that it is faster; also NFS timeouts go in parallel.
              A disadvantage is that the mounts are done in undefined order.
              Thus, you cannot use this option if you want to mount both
              /usr and /usr/spool.

#### Linux Device

https://github.com/golang/sys/blob/master/unix/dev_linux.go
		// Mkdev returns a Linux device number generated from the given major and minor
		// components.
		func Mkdev(major, minor uint32) uint64 {
			dev := (uint64(major) & 0x00000fff) << 8
			dev |= (uint64(major) & 0xfffff000) << 32
			dev |= (uint64(minor) & 0x000000ff) << 0
			dev |= (uint64(minor) & 0xffffff00) << 12
			return dev
		}


#### ram(4) syscall
ram - ram disk device. The ram device is a block device to access the ram disk in raw mode.
It is typically created by:
           mknod -m 660 /dev/ram b 1 1
           chown root:disk /dev/ram

#### Systemd.link
systemd.link - Network device configuration

 			 Network link configuration is performed by the net_setup_link udev
       builtin.

       The link files are read from the files located in the system network
       directory /usr/lib/systemd/network, the volatile runtime network
       directory /run/systemd/network, and the local administration network
       directory /etc/systemd/network. Link files must have the extension
       .link; other extensions are ignored. All link files are collectively
       sorted and processed in lexical order, regardless of the directories
       in which they live. However, files with identical filenames replace
       each other. Files in /etc have the highest priority, files in /run
       take precedence over files with the same name in /usr/lib. This can
       be used to override a system-supplied link file with a local file if
       needed.

[MATCH] SECTION OPTIONS

       A link file is said to match a device if each of the entries in the
       "[Match]" section matches, or if the section is empty. The following
       keys are accepted:

       MACAddress=
           The hardware address.

       OriginalName=
           A whitespace-separated list of shell-style globs matching the
           device name, as exposed by the udev property "INTERFACE". This
           cannot be used to match on names that have already been changed
           from userspace. Caution is advised when matching on
           kernel-assigned names, as they are known to be unstable between
           reboots.

       Path=
           A whitespace-separated list of shell-style globs matching the
           persistent path, as exposed by the udev property "ID_PATH".

       Driver=
           A whitespace-separated list of shell-style globs matching the
           driver currently bound to the device, as exposed by the udev
           property "DRIVER" of its parent device, or if that is not set,
           the driver as exposed by "ethtool -i" of the device itself.

       Type=
           A whitespace-separated list of shell-style globs matching the
           device type, as exposed by the udev property "DEVTYPE".

       Host=
           Matches against the hostname or machine ID of the host. See
           "ConditionHost=" in systemd.unit(5) for details.

       Virtualization=
           Checks whether the system is executed in a virtualized
           environment and optionally test whether it is a specific
           implementation. See "ConditionVirtualization=" in systemd.unit(5)
           for details.

       KernelCommandLine=
           Checks whether a specific kernel command line option is set (or
           if prefixed with the exclamation mark unset). See
           "ConditionKernelCommandLine=" in systemd.unit(5) for details.

       Architecture=
           Checks whether the system is running on a specific architecture.
           See "ConditionArchitecture=" in systemd.unit(5) for details.

[LINK] SECTION OPTIONS

       The "[Link]" section accepts the following keys:

       Description=
           A description of the device.

       Alias=
           The "ifalias" is set to this value.

       MACAddressPolicy=
           The policy by which the MAC address should be set. The available
           policies are:

           "persistent"
               If the hardware has a persistent MAC address, as most
               hardware should, and if it is used by the kernel, nothing is
               done. Otherwise, a new MAC address is generated which is
               guaranteed to be the same on every boot for the given machine
               and the given device, but which is otherwise random. This
               feature depends on ID_NET_NAME_* properties to exist for the
               link. On hardware where these properties are not set, the
               generation of a persistent MAC address will fail.

           "random"
               If the kernel is using a random MAC address, nothing is done.
               Otherwise, a new address is randomly generated each time the
               device appears, typically at boot. Either way, the random
               address will have the "unicast" and "locally administered"
               bits set.

           "none"
               Keeps the MAC address assigned by the kernel.

       MACAddress=
           The MAC address to use, if no "MACAddressPolicy=" is specified.

       NamePolicy=
           An ordered, space-separated list of policies by which the
           interface name should be set.  "NamePolicy" may be disabled by
           specifying "net.ifnames=0" on the kernel command line. Each of
           the policies may fail, and the first successful one is used. The
           name is not set directly, but is exported to udev as the property
           "ID_NET_NAME", which is, by default, used by a udev rule to set
           "NAME". If the name has already been set by userspace, no
           renaming is performed. The available policies are:

           "kernel"
               If the kernel claims that the name it has set for a device is
               predictable, then no renaming is performed.

           "database"
               The name is set based on entries in the udev's Hardware
               Database with the key "ID_NET_NAME_FROM_DATABASE".

           "onboard"
               The name is set based on information given by the firmware
               for on-board devices, as exported by the udev property
               "ID_NET_NAME_ONBOARD".

           "slot"
               The name is set based on information given by the firmware
               for hot-plug devices, as exported by the udev property
               "ID_NET_NAME_SLOT".

           "path"
               The name is set based on the device's physical location, as
               exported by the udev property "ID_NET_NAME_PATH".

           "mac"
               The name is set based on the device's persistent MAC address,
               as exported by the udev property "ID_NET_NAME_MAC".

       Name=
           The interface name to use in case all the policies specified in
           NamePolicy= fail, or in case NamePolicy= is missing or disabled.

       MTUBytes=
           The maximum transmission unit in bytes to set for the device. The
           usual suffixes K, M, G, are supported and are understood to the
           base of 1024.

       BitsPerSecond=
           The speed to set for the device, the value is rounded down to the
           nearest Mbps. The usual suffixes K, M, G, are supported and are
           understood to the base of 1000.

       Duplex=
           The duplex mode to set for the device. The accepted values are
           "half" and "full".

       AutoNegotiation=
           Enables or disables automatic negotiation of transmission
           parameters. Autonegotiation is a procedure by which two connected
           ethernet devices choose common transmission parameters, such as
           speed, duplex mode, and flow control. Takes a boolean value.
           Unset by default, which means that the kernel default will be
           used.

           Note that if autonegotiation is enabled, speed and duplex
           settings are read-only. If autonegotation is disabled, speed and
           duplex settings are writable if the driver supports multiple link
           modes.

       WakeOnLan=
           The Wake-on-LAN policy to set for the device. The supported
           values are:

           "phy"
               Wake on PHY activity.

           "magic"
               Wake on receipt of a magic packet.

           "off"
               Never wake.

       Port=
           The port option is used to select the device port. The supported
           values are:

           "tp"
               An Ethernet interface using Twisted-Pair cable as the medium.

           "aui"
               Attachment Unit Interface (AUI). Normally used with hubs.

           "bnc"
               An Ethernet interface using BNC connectors and co-axial
               cable.

           "mii"
               An Ethernet interface using a Media Independent Interface
               (MII).

           "fibre"
               An Ethernet interface using Optical Fibre as the medium.

       TCPSegmentationOffload=
           The TCP Segmentation Offload (TSO) when true enables TCP
           segmentation offload. Takes a boolean value. Defaults to "unset".

       GenericSegmentationOffload=
           The Generic Segmentation Offload (GSO) when true enables generic
           segmentation offload. Takes a boolean value. Defaults to "unset".

       UDPSegmentationOffload=
           The UDP Segmentation Offload (USO) when true enables UDP
           segmentation offload. Takes a boolean value. Defaults to "unset".

       GenericReceiveOffload=
           The Generic Receive Offload (GRO) when true enables generic
           receive offload. Takes a boolean value. Defaults to "unset".

       LargeReceiveOffload=
           The Large Receive Offload (LRO) when true enables large receive
           offload. Takes a boolean value. Defaults to "unset".

EXAMPLES

       Example 1. /usr/lib/systemd/network/99-default.link

       The link file 99-default.link that is shipped with systemd defines
       the default naming policy for links.

           [Link]
           NamePolicy=kernel database onboard slot path
           MACAddressPolicy=persistent

       Example 2. /etc/systemd/network/10-dmz.link

       This example assigns the fixed name "dmz0" to the interface with the
       MAC address 00:a0:de:63:7a:e6:

           [Match]
           MACAddress=00:a0:de:63:7a:e6

           [Link]
           Name=dmz0

       Example 3. /etc/systemd/network/10-internet.link

       This example assigns the fixed name "internet0" to the interface with
       the device path "pci-0000:00:1a.0-*":

           [Match]
           Path=pci-0000:00:1a.0-*

           [Link]
           Name=internet0

       Example 4. /etc/systemd/network/25-wireless.link

       Here's an overly complex example that shows the use of a large number
       of [Match] and [Link] settings.

           [Match]
           MACAddress=12:34:56:78:9a:bc
           Driver=brcmsmac
           Path=pci-0000:02:00.0-*
           Type=wlan
           Virtualization=no
           Host=my-laptop
           Architecture=x86-64

           [Link]
           Name=wireless0
           MTUBytes=1450
           BitsPerSecond=10M
           WakeOnLan=magic
           MACAddress=cb:a9:87:65:43:21


















#### initrd Linux kernel syscall
initrd - boot loader initialized RAM disk
CONFIGURATION
       /dev/initrd is a read-only block device assigned major number 1 and
       minor number 250.  Typically /dev/initrd is owned by root.disk with
       mode 0400 (read access by root only).  If the Linux system does not
       have /dev/initrd already created, it can be created with the
       following commands:
           mknod -m 400 /dev/initrd b 1 250
           chown root:disk /dev/initrd
       Also, support for both "RAM disk" and "Initial RAM disk" (e.g., CON‐
       FIG_BLK_DEV_RAM=y and CONFIG_BLK_DEV_INITRD=y) must be compiled
       directly into the Linux kernel to use /dev/initrd.  When using
       /dev/initrd, the RAM disk driver cannot be loaded as a module.
DESCRIPTION
       The special file /dev/initrd is a read-only block device.  This
       device is a RAM disk that is initialized (e.g., loaded) by the boot
       loader before the kernel is started.  The kernel then can use
       /dev/initrd's contents for a two-phase system boot-up.

       In the first boot-up phase, the kernel starts up and mounts an
       initial root filesystem from the contents of /dev/initrd (e.g., RAM
       disk initialized by the boot loader).  In the second phase,
       additional drivers or other modules are loaded from the initial root
       device's contents.  After loading the additional modules, a new root
       filesystem (i.e., the normal root filesystem) is mounted from a
       different device.
   Boot-up operation
       When booting up with initrd, the system boots as follows:
       1. The boot loader loads the kernel program and /dev/initrd's
          contents into memory.
       2. On kernel startup, the kernel uncompresses and copies the contents
          of the device /dev/initrd onto device /dev/ram0 and then frees the
          memory used by /dev/initrd.
       3. The kernel then read-write mounts the device /dev/ram0 as the
          initial root filesystem.
       4. If the indicated normal root filesystem is also the initial root
          filesystem (e.g., /dev/ram0) then the kernel skips to the last
          step for the usual boot sequence.
       5. If the executable file /linuxrc is present in the initial root
          filesystem, /linuxrc is executed with UID 0.  (The file /linuxrc
          must have executable permission.  The file /linuxrc can be any
          valid executable, including a shell script.)
       6. If /linuxrc is not executed or when /linuxrc terminates, the
          normal root filesystem is mounted.  (If /linuxrc exits with any
          filesystems mounted on the initial root filesystem, then the
          behavior of the kernel is UNSPECIFIED.  See the NOTES section for
          the current kernel behavior.)
       7. If the normal root filesystem has a directory /initrd, the device
          /dev/ram0 is moved from / to /initrd.  Otherwise, if the directory
          /initrd does not exist, the device /dev/ram0 is unmounted.  (When
          moved from / to /initrd, /dev/ram0 is not unmounted and therefore
          processes can remain running from /dev/ram0.  If directory /initrd
          does not exist on the normal root filesystem and any processes
          remain running from /dev/ram0 when /linuxrc exits, the behavior of
          the kernel is UNSPECIFIED.  See the NOTES section for the current
          kernel behavior.)
       8. The usual boot sequence (e.g., invocation of /sbin/init) is
          performed on the normal root filesystem.
   Options
       The following boot loader options, when used with initrd, affect the
       kernel's boot-up operation:

       initrd=filename
              Specifies the file to load as the contents of /dev/initrd.
              For LOADLIN this is a command-line option.  For LILO you have
              to use this command in the LILO configuration file
              /etc/lilo.config.  The filename specified with this option
              will typically be a gzipped filesystem image.

       noinitrd
              This boot option disables the two-phase boot-up operation.
              The kernel performs the usual boot sequence as if /dev/initrd
              was not initialized.  With this option, any contents of
              /dev/initrd loaded into memory by the boot loader contents are
              preserved.  This option permits the contents of /dev/initrd to
              be any data and need not be limited to a filesystem image.
              However, device /dev/initrd is read-only and can be read only
              one time after system startup.

       root=device-name
              Specifies the device to be used as the normal root filesystem.
              For LOADLIN this is a command-line option.  For LILO this is a
              boot time option or can be used as an option line in the LILO
              configuration file /etc/lilo.config.  The device specified by
              the this option must be a mountable device having a suitable
              root filesystem.

   Changing the normal root filesystem
       By default, the kernel's settings (e.g., set in the kernel file with
       rdev(8) or compiled into the kernel file), or the boot loader option
       setting is used for the normal root filesystems.  For an NFS-mounted
       normal root filesystem, one has to use the nfs_root_name and
       nfs_root_addrs boot options to give the NFS settings.  For more
       information on NFS-mounted root see the kernel documentation file
       Documentation/filesystems/nfs/nfsroot.txt (or
       Documentation/filesystems/nfsroot.txt before Linux 2.6.33).  For more
       information on setting the root filesystem see also the LILO and
       LOADLIN documentation.

       It is also possible for the /linuxrc executable to change the normal
       root device.  For /linuxrc to change the normal root device, /proc
       must be mounted.  After mounting /proc, /linuxrc changes the normal
       root device by writing into the proc files /proc/sys/kernel/real-
       root-dev, /proc/sys/kernel/nfs-root-name, and /proc/sys/kernel/nfs-
       root-addrs.  For a physical root device, the root device is changed
       by having /linuxrc write the new root filesystem device number into
       /proc/sys/kernel/real-root-dev.  For an NFS root filesystem, the root
       device is changed by having /linuxrc write the NFS setting into files
       /proc/sys/kernel/nfs-root-name and /proc/sys/kernel/nfs-root-addrs
       and then writing 0xff (e.g., the pseudo-NFS-device number) into file
       /proc/sys/kernel/real-root-dev.  For example, the following shell
       command line would change the normal root device to /dev/hdb1:

           echo 0x365 >/proc/sys/kernel/real-root-dev

       For an NFS example, the following shell command lines would change
       the normal root device to the NFS directory /var/nfsroot on a local
       networked NFS server with IP number 193.8.232.7 for a system with IP
       number 193.8.232.2 and named "idefix":

           echo /var/nfsroot >/proc/sys/kernel/nfs-root-name
           echo 193.8.232.2:193.8.232.7::255.255.255.0:idefix \
               >/proc/sys/kernel/nfs-root-addrs
           echo 255 >/proc/sys/kernel/real-root-dev

       Note: The use of /proc/sys/kernel/real-root-dev to change the root
       filesystem is obsolete.  See the Linux kernel source file Documenta‐
       tion/admin-guide/initrd.rst (or Documentation/initrd.txt before Linux
       4.10) as well as pivot_root(2) and pivot_root(8) for information on
       the modern method of changing the root filesystem.

   Usage
       The main motivation for implementing initrd was to allow for modular
       kernel configuration at system installation.

       A possible system installation scenario is as follows:

       1. The loader program boots from floppy or other media with a minimal
          kernel (e.g., support for /dev/ram, /dev/initrd, and the ext2
          filesystem) and loads /dev/initrd with a gzipped version of the
          initial filesystem.

       2. The executable /linuxrc determines what is needed to (1) mount the
          normal root filesystem (i.e., device type, device drivers,
          filesystem) and (2) the distribution media (e.g., CD-ROM, network,
          tape, ...).  This can be done by asking the user, by auto-probing,
          or by using a hybrid approach.

       3. The executable /linuxrc loads the necessary modules from the ini‐
          tial root filesystem.

       4. The executable /linuxrc creates and populates the root filesystem.
          (At this stage the normal root filesystem does not have to be a
          completed system yet.)

       5. The executable /linuxrc sets /proc/sys/kernel/real-root-dev,
          unmount /proc, the normal root filesystem and any other filesys‐
          tems it has mounted, and then terminates.

       6. The kernel then mounts the normal root filesystem.

       7. Now that the filesystem is accessible and intact, the boot loader
          can be installed.

       8. The boot loader is configured to load into /dev/initrd a filesys‐
          tem with the set of modules that was used to bring up the system.
          (e.g., Device /dev/ram0 can be modified, then unmounted, and
          finally, the image is written from /dev/ram0 to a file.)

       9. The system is now bootable and additional installation tasks can
          be performed.

       The key role of /dev/initrd in the above is to reuse the configura‐
       tion data during normal system operation without requiring initial
       kernel selection, a large generic kernel or, recompiling the kernel.

       A second scenario is for installations where Linux runs on systems
       with different hardware configurations in a single administrative
       network.  In such cases, it may be desirable to use only a small set
       of kernels (ideally only one) and to keep the system-specific part of
       configuration information as small as possible.  In this case, create
       a common file with all needed modules.  Then, only the /linuxrc file
       or a file executed by /linuxrc would be different.

       A third scenario is more convenient recovery disks.  Because informa‐
       tion like the location of the root filesystem partition is not needed
       at boot time, the system loaded from /dev/initrd can use a dialog
       and/or auto-detection followed by a possible sanity check.

       Last but not least, Linux distributions on CD-ROM may use initrd for
       easy installation from the CD-ROM.  The distribution can use LOADLIN
       to directly load /dev/initrd from CD-ROM without the need of any
       floppies.  The distribution could also use a LILO boot floppy and
       then bootstrap a bigger RAM disk via /dev/initrd from the CD-ROM.

FILES
       /dev/initrd
       /dev/ram0
       /linuxrc
       /initrd
NOTES
       1. With the current kernel, any filesystems that remain mounted when
          /dev/ram0 is moved from / to /initrd continue to be accessible.
          However, the /proc/mounts entries are not updated.
       2. With the current kernel, if directory /initrd does not exist, then
          /dev/ram0 will not be fully unmounted if /dev/ram0 is used by any
          process or has any filesystem mounted on it.  If /dev/ram0 is not
          fully unmounted, then /dev/ram0 will remain in memory.
       3. Users of /dev/initrd should not depend on the behavior give in the
          above notes.  The behavior may change in future versions of the
          Linux kernel.

==============
#### Rust Container
https://github.com/vishvananda/railcar

#### Microcontroller
https://github.com/oracle/smith

#### Networking
[stdlib:net](https://github.com/golang/net)
  (^)[net/ipv4/endpoint](https://github.com/golang/net/blob/master/ipv4/endpoint.go)
     Potentially way to setup a ipv4 endpoint
  (^)[net/ipv4/packet](https://github.com/golang/net/blob/master/ipv4/packet.go)
                      (https://github.com/golang/net/blob/master/ipv4/packet_go1_9.go) 
                      (https://github.com/golang/net/blob/master/ipv4/readwrite_go1_9_test.go)
     Potentially a way to reroute packets
   
**potential starting points**
* could be what we source from  : https://github.com/wrigby/flowdump/blob/master/main.go
https://github.com/ld86/syscall-udp/blob/master/main.go
https://github.com/bradleyfalzon/go-syscall-sockets/blob/master/main.go
https://github.com/prayerslayer/go-tcp/blob/master/main.go - tcp http server with workers


####
https://github.com/u-root/u-root
initramfs containing busybox like system
----
## Go Mod Handling (Kmod)
[golang-kmod](https://github.com/ElyKar/golang-kmod)


====
## Mod Handling

https://github.com/wdv4758h/rust_kernel_module/blob/master/src/stubs/linux.c
https://github.com/saschagrunert/kmod
#include <linux/module.h>
#include <linux/slab.h>

MODULE_AUTHOR("Sascha Grunert <mail@saschagrunert.de>");
MODULE_DESCRIPTION("A simple kernel module");
MODULE_LICENSE("MIT");
MODULE_VERSION("0.1.0");

// The entry and exit function
extern int init_module(void);
extern void cleanup_module(void);
----
#include <linux/module.h>

extern int rust_init(void);
extern void rust_exit(void);

static int hello_init(void) {
    return rust_init();
}

static void hello_exit(void) {
    return rust_exit();
}

module_init(hello_init);
module_exit(hello_exit);

MODULE_LICENSE("Dual BSD/GPL");
#### Container WebUI
https://github.com/coreos/etcdlabs


#### Databases
[etcd](https://github.com/coreos/etcd)


#### Kernel Modules
[conntrack](https://github.com/typetypetype/conntrack)

#### Storage
[go-tcmu](https://github.com/coreos/go-tcmu)
Go SCSI emulation via Linux TCM Userspace module


#### Configurations
There are existing configs and Multiverse OS should at least review them if not use one.

**A Container Linux Configuration, to be processed by ct, is a YAML document conforming to the following specification:**


#### Database
[riak like kv embeddable](https://github.com/tadvi/rkv)
Rkv - embeddable KV database in Go (golang).
		  Based on Riak bitcast format
		  Minimalistic design
		  Embeddable and self-contained (no C dependencies)
		  Contains both direct and goroutine friendly interfaces
		  Use rkv.NewSafe("test.kv") if you want to use with goroutines
		  Basic KV admin tool is included in /rkv subfolder, build it and install in your bin folder
		  Ability to save records with expiration
		  Use Rkv for databases under 50K records


#### WebUI TTY
https://github.com/yudai/gotty
#### Ruby
[go-mruby](https://github.com/mitchellh/go-mruby)
Ruby would be great to implement a console/terminal
that used RUBY instead of BASH 

#### Multistep / Pipeline / Workflow
[multistep](https://github.com/mitchellh/multistep)
multistep is a Go library for building up complex actions using discrete steps.

#### Filesystem
[go-fs](https://github.com/mitchellh/go-fs)
implements the ability to create, read, and write FAT filesystems using pure Go.

[go-tcmu](https://github.com/coreos/go-tcmu)
Go SCSI emulation via the Linux TCM in Userpsace module
 A daemon that handles the userspace side of the LIO TCM-User backstore. 
#### CLI UI
**IO PROGRESS**
[ioprogress](https://github.com/mitchellh/ioprogress)

**CLI Framework**
[cli](https://github.com/mitchellh/cli)
simple lightweight cli with command + subcommands
====
----
**Homedir**
https://github.com/mitchellh/go-homedir
**Packets**

**PCAP**
https://github.com/davecheney/pcap

**Standard Library: Socks5 Proxy**
https://github.com/golang/net/blob/master/proxy/socks5.go

## Potentially Interesting Image Building Tools

#### Provisioning

[routinator](https://github.com/abedra/routinator)
Good example of using templates for configs that can have variables inserted inside then moving and copying them into folders

#### Image Building 
[uspin](https://github.com/solus-project/USpin)

[go-debian](https://github.com/paultag/go-debian)
[go-bin-deb](https://github.com/mh-cbon/go-bin-deb)

[go-systemd](https://github.com/coreos/go-systemd)


[debos](https://github.com/go-debos/debos)
Debian OS Builder
  (^)[fakemachine](https://github.com/go-debos/fakemachine)
  (^)[debos-recipies](https://github.com/go-debos/debos-recipes)

#### FS
[memfs](https://github.com/zbiljic/memfs)
**nbd**
[buse-go](https://github.com/samalba/buse-go)
[go-nbd](https://github.com/frostschutz/go-nbd)
[gonbdserver](https://github.com/abligh/gonbdserver)
#### Media
**Music**
[id3](https://github.com/beevik/id3)

#### Database
[tiedot](https://github.com/HouzuoGuo/tiedot)
nosql database, embedded
#### Software defined networking
[go-openswitch/ovs](https://github.com/digitalocean/go-openvswitch/tree/master/ovs)
Package ovs is a client library for Open vSwitch which enables programmatic control of the virtual switch.
[grpc](https://github.com/grpc/grpc-go)

**distance vector routering (routing optimization)**
[go-distance-vector-routing](https://github.com/taylorflatt/go-distance-vector-routing)
DVR and a fastest path algorithm implemented using Go channels
#### (DHT) Chord lookup algorithm
[go-chord](https://github.com/taylorflatt/go-chord-implementation)

#### Packets / Networking
[gopacket](https://github.com/google/gopacket)
Very complete networking stuff
#### Tun/Tap/Device
[goTunTap](https://github.com/traetox/goTunTap)
Interesting, comes with vpn
[go-tuntap](https://github.com/0xef53/go-tuntap)
#### KVM/Qemu/Libvirt
[go-qemu](https://github.com/digitalocean/go-qemu)
[go-libvirt](https://github.com/digitalocean/go-libvirt)

[pheonix-guest-agent](https://github.com/0xef53/phoenix-guest-agent)
A guest-side agent for qemu-kvm virtual machines
[qmp-shell](https://github.com/0xef53/qmp-shell/blob/master/qmp-shell.go)


#### Iptables / Firewall
[firewall](https://github.com/Gouthamve/go-firewall)
A firewall using nfqueue (not iptables), this makes it pretty low level or atleast very high power. 


[go-iptables](https://github.com/coreos/go-iptables)
Go wrapper around iptables utility In-kernel netfilter does not have a good userspace API. The tables are manipulated via setsockopt that sets/replaces the entire table. Changes to existing table need to be resolved by userspace code which is difficult and error-prone. Netfilter developers heavily advocate using iptables utlity for programmatic manipulation.
  (^)[iptables-api](https://github.com/Oxalide/iptables-api/blob/master/main.go)


#### Globbing & Path/String Matching
[doublestar](https://github.com/bmatcuk/doublestar)


=========
#### Notes / Research
* Binding to :0 will ask the kernel for an available port within the ephemeral port range
the kernel will assign it a free port number somewhere above 1024.

	// tell pinger that it is privileged.
	// NOTE: You must run `setcap cap_net_raw=+ep pocket-loss-monitor`
  pinger.SetPrivileged(true)
=========
#### Device
https://github.com/vizee/tcpproxy splice syscall?
=========
#### Tools For Building Debian Images

xorriso - Turns folder into iso image


=========
Notes from a manual Multiverse OS setup on Intel I9.
The process of installing Multiverse OS is very detailed and ideally is done using fresh hardware that has never been used. One starts with BIOS settings and slowly builds upon secure settings.
## BIOS Settings
Execute Disable Bit = ENABLED
"It is highly recommended that you enable this BIOS feature for increased protection against buffer overflow attacks."
## Networking Devices
The Asus motherboard I'm using for the Intel I9 processor is an Atheros NIC, these drivers are not included with the Debian 9 net install. Complexities like this have been anticipated so a Realtek Gigabit network card was purchased with the rest of the components for eventualalties such as this but also to provide a wider variety of network cards and drivers for the Multiverse VM networking.

Realtek drivers are included with the Debian 9 network install and the installation can continue by using this card for *Host machine* initialization.

**NOTE** [FUTURE DEVELOPMENT]
*For maximum compatibility this eventuality should be considered and the Multiverse OS install medium should incldue the Atheros driver if it is open source.*
## Gnome
https://github.com/reujab/gse
## Debian Installation
Multiverse OS follows the same conventions as other psuedo-anonymous security oriented linux operating systems:
	[Accounts]
	Hostname: "host"
	Domain name: ""
	Full name: "user"
	Username: "user"
**NOTE** [FUTURE DEVELOPMENT]
*In the future this will all be automated and the root and user passwords will be automatically generated deterministically from a master key based on time and other factors. This will allow for OTP, automatic renewal of passwords, offline recovery, automatic input for secure booting of the host machine and eventual linux kernel modifications that will massively increase the security of the system (require user input to be signed, and encrypt user output).*
[Disks]
[x] Guided - use entire disk and setup encrypted LVM
**NOTE** Debian 9 install had issues with the install process when using M.2 drives because
**NOTE** [FUTURE DEVELOPMENT]
	[Packages]
	No desktop envirnonment (No Gnome, Xfce, KDE, Cinnamon, MATE, LXDWE)
	No web server
	No print server
	No SSH server
*ONLY standard system utilities*

[Post-Install]
**NOTE** Key Generation should happen here, PGP, SSH, etc (Basically setting up Shifter/Scramble Suit)
**NOTE** WGET urls for (1) Debian, (2) Alpine Linux, (3) Kali
+ Untar/Unzip ISOs, check if each one required exists using checksums 
  [+] Check against checksum, signature (preferred, should do it with multiverse OS dev key if not provided by official distribution).
  [+] Add a new "Storage" in libvirt for the ISOs 

sudo apt-get remove firefox-esr nano
sudo apt-get install libvirt-daemon-system vim
## DEV
## sudo apt-get install virt-manager gnome
## 
## I used the UI when developing, but ideally this should not be 

**NOTE** After the packages are settled, it is time to server connection to the internet. Do this by forcing the kernel not to load ANY networking modules. This will free them up so they can be used with the VMs too (PCI Passthrough).

Modify /etc/default/grub, add `intel_iommu=on` to the options line.

(Here we need to disable any graphics cards too eventually, but if you do this too early you will not have a monitor to setup the controller VM.)

Additionally, locate each network device and disable it: `pci-stub.ids=10de:13c2,10de:0fbb`

Start this process by using the command

````
	lspci -nn
````

The `-nn` provides a truncated output with the pci address information needed to disable the device by preventing the kernel from loading it by blacklisting it.

sudo touch /etc/modprobe.d/blacklist.conf

sudo echo "#Blacklist Firewire\nblacklist firewire_core\n\n#Blacklist PC Speaker (Annoying)\nblacklist pcspkr" > /etc/modprobe.d/blacklist.conf

# echo "0000:00:1a.0" > /sys/bus/pci/drivers/ehci_hcd/unbind

## Reverse SSHD
Here is how you do the reverse SSH if you decide that route

ssh -R NEW_SERVER:80:127.0.0.1:80 unprivuser@NEW_SERVER -v

Basically then you can ssh into like port 2222 on the chunkhost server and it goes into the PI without needing any ports open on the modem.



