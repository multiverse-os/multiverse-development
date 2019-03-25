

## Basic Design
Multiverse OS intentionally prevents users from ever directly using the bare-metal host operating system. Instead, by using nested virtual machines (VMs), the user always operates within VMs. The "controller" VM is passed the bare-metal host GPU which gives the VM-near native speed, while preventing potentially dangerous GPU drivers from being installed on the bare-metal host.

VMs "above" and "below" the "controller" VM are only accessible through very limited API's exposed through virtio. Using virtio is fast, and has the additional benefit of being obfuscated since HTTP and SSH are more common transports. And, since it is already used by the existing system, it only minimally increases the attack surface.

## Identifying Modern Computer Security Problems
Multiverse OS was created in response to research into growing sophistication, automation, and capital inflow applied to malware, root-kits, worms and viruses. The rate of increasing sophistication is astounding; kids, black hats and mad scientists have been joined by state actors, private intelligence corporations and the growing ransomware industry. Meanwhile, most operating systems have failed to keep up with improving operating system defenses for "average" users.

A system can be hardended but even "advanced" users are often affected, as it is increasingly harder to escape the ubiquitiousness of the growing XSS attack surface they interact with each day when browsing the internet.

In the current context of the internet, it is best to assume every website could be exploited. The operating system should be paranoid, so the user does not have to be. A free, open-source, easy-to-use hardened OS secured with ephemeral compartmentalization using nested VMs is needed to address the emerging persistent automated malware threat facing both *"average"* and *"advanced"* users:

* A growing number of sophisticated root kits, malware and automated viruses are being written by state intelligence, private intelligence, organized crime and corrupt companies interested in access to the large inflow of capital into offensive computer security. Anytime a port is opened on a computer with internet access on a WAN IP connections will be made almost immediately from bots investigating. 
* A growing number of web applications over-utilize rush implemented Javascript, resulting in insecure applications vulnerable to cross-site-scripting (XSS) vulnerabilities. The problem is amplified by large numbers of poorly implemented server configurations and growing numbers of sites failing to provide working alternatives that do not require Javascript. CSS or other alternatives are not even available when the Javascript is being used solely for superficial visual improvement of the UI. The state of the internet forces internet users into regularly running code from anonymous untrusted parties on their machines. Even if the hosts are not malicious, poor Javascript regularly exposes users to XSS attacks.
* Limited infrastructure to validate Javascript loaded from other domains, and widespread reliance on CDNs and conglomerate hosting creates central points of weakness. Javascript is even included in advertisements, and has frequently been exploited to spread malware through ads on pages.
* Vulerable/exploitable software anywhere in the system compromises all the data on the system including: cryptographic keys, personal photos, personal documents, source code, browsing history and more. 
 The modern Internet ecosystem, regardless of private browsing mode, do-not-track options, or even claims that search engines do not track, allows for user tracking using almost anything imaginable to create unique fingerprints. Browser settings, connection patterns, DNS queries and more can be used to compromize anonymity. This data unncessarily gives away valuable private data including location, social connections, browsing habits, interests, and research or work topics.
* Clustering (creating a cloud) securely is not easy with diverse collections of hardware and not typically designed for general home use. Many tools for for home LAN networks are typically lacking even basic security.

Given this context in which we browse the internet, it is just a matter of time before you execute malicious javascript that directs you to a page with reverse shell exploits and you are exposed to a vulnerability you did not yet patch.

#### Why not just use QubesOS?

Qubes OS may be the most secure open source operating system currently available. After months of use, reading code, reading forums, I came to the conclusion that the QubesOS developers have different requirements for their operating system. While QubesOS provides a secure environment, and generally works well for development and general internet browsing, it did not provide all the features I needed for my general internet use. 

1. **Multiverse OS should use Debian, to benefit from scripts written for Tails OS and Whonix OS** instead of Fedora.
2. **libvirtd for interfacing with QEMU/KVM** instead of QubesOS choice of libvirtd communicating with Xen. Xen has better performance but QEMU/KVM has better support for total number of GPU's capable of passthrough.
1. **Multiverse OS should use VirtIO for inter-VM networking instead of SSH** This allows implementation of communication over virtIO and disabling all typical tcp/ip communication to limit virus mobility. 
2. **Multiverse OS should use passwords, one-time-passwords, HMAC signatures, and an encrypted password store to generate and gpg encrypt 64 random luks hard drive passwords**, validation and verification can happen over VirtIO connections. This contrasts with the QubesOS approach of no passwords once a user is logged in.
3. **Multiverse OS should use the Alpine Linux security-focused microkernel for routers and most utility vms, rather than Fedora**. Fedora Desktop has a much larger attack surface than Alpine Linux, a security-focused microkernel packaged with Grsecurity.
3. **Mutliverse OS should be built primarily with shell and compiled languages like Go, C or Rust** instead of QubesOS' choice of Python. Python does not compile into binaries and Go has very mature libraries for working with libvirtd, QEMU and KVM. 
3. **Multiverse OS requires wide support for GPU PCI-passthrough** because using GPU passthrough in combination with locking down the bare metal host makes the attack surface of the host (dom0 equivilent) almost non-existant since the user never logs into it, or directly interacts with it. 
3. **Additional support, utility-vms and work done to support general use for "average" users including features like torrenting, linux gaming, developing and scientific paralell processing programming use cases securely** 
4. **Everything should be run as an unprivileged user** rather than root.

#### A different approach to QubesOS-like security
Multiverse OS' focus is an operating system that is easy to use, with defaults that are built with general-use computing needs in mind, rich customization/development environment, and high security KVM/QEMU compartmentalization with nested ephemeral VMs.

Multiverse OS runs on `Debian8 Stretch/Main` using `libvirtd` to compartmentalize QEMU/KVM VMs, with potential support for LXC, LXD, rkt and nspawn containers.

**Input/Output Passed To Controller VM**

The physical USB ports, GPU, audio are assigned to `controller VMs` using PCI-passthrough running a nested Multiverse OS Debian8 installation.

**Application Sandboxing In Controller VM**

Applications are run within QEMU/KVM virtual machines and accessible to users by using `xpra` over TCP. This method does not provide a shell to the connecting user and the virtual machines can be run from separate unprivileged user accounts. More complete networking isolation may be possible if VirtIO could be used to provide a `xpra` tcp port.

Containers may be incorporated through libvirtd, adding the possibility for additional layers of isolation by nesting containers inside of VMs or VMs inside of containers.

*Containers in their current state appear to trade off security for increased preformance, and so with sandbox effectiveness being critical to the Multiverse OS security design, containers alone should not be used. If anything has the ability to increase effective isolation, it should be used or provided as an option.*

Efforts are being made to make this transparent to the user.

**Layered Networking**

The physical network cards are assigned to router VMs using PCI-passthrough running the `Alpine Linux` security-focused microkernel and layered to provide isolated subnets that can be isolated from the LAN, transparently proxied through VPNs, Tor or remote servers.

**Long Running Utility VMs In Segregated Networks**

Long running applications requiring no human interaction and special networking are isolated from the `controller-vms` and stored in `utility vms`. 

For example, a torrent application can be run within a `utility vm`, downloading torrents by being provisioned to mount shared hard drives or folders with minimal privileges. It could have read-only rights to watch the user's browser downloads folder for torrent files and automatically download the resulting file into a separate folder readable by a media server utility vm.

By segregating these long-running utilities, more general uses emerge, allowing for a greater number of people to securely access the internet.

Efforts are being made to make this transparent to the user.

#### Types of VMs
Multiverse OS uses four primary types of VMs to simplify management, each of these types have their own sub-types. These VMs provide the building blocks needed to provide functionality for a wide variety of use cases.

1. **Router VM** A router VM is currently primarily being built from the Alpine Linux microkernel. This OS is incredibly small, security focused and uses minimal resources. The primary router VM receives the bare-metal NIC by utilizing the PCI-Passthrough feature.
  * WAN Router VM - Designed to receive NIC card and manage connection with other routers or provide limited access to utility VMs that need LAN internet access.
  * Firewall Router VM - The firewall and proxy server. This is where your openvpn client will be setup. Ideally you are connecting to a server you rent and put a VPN on yourself or you are using a reputable VPN seller that takes Bitcoin. The Whonix box will route its traffic through this to hide Tor access from your ISP.
  * Whonix Router VM - For transparent proxying of applications through Tor.
  * Isolating VPN Router VM (optional) - For further transparent proxying through a VPN.
2. **Controller VM** A controller VM receives the GPU and, like the host VM, does not access the internet directly. It has very limited control of the host machine through VirtIO interface, verfied by one-time-passwords, HMAC and nonces to prevent replays. Application VMs are created within Controller VMs as unprivileged user.
3. **Application VM** An application VM is by default an ephemeral VM that starts applications within containers/jailshells and exposes window GUI access using xpra TCP connections. TCP connections allow access to the appliaction without giving out full SSH shell access.
  - Example types:
  * Vault - Has either Electrum or a full Bitcoin node running in a VM to isolate it from computers that browse the internet. 
  * MultiVault - An application VM that supplies an interface to several isolated VMs with different cryptocurrencies stored.
4. **Utility VM** As explained above, these are VMs with special networking requirements, such as LAN access. They can be isolated from the Multiverse network or have an interface for communication with or control by a controller VM. Ideally, each running utility can provide access to its data without requiring /dev/input (mouse and keyboard) type input. This list is meant to provide examples of the concept and is by no means exhaustive.
  * USB Protection Utility VM - Designed to isolate unknown USB devices in a contained and monitored environment. This VM ideally has many common tools necessary for malware analysis installed and no internet connection.
  * Captive Portal Utility VM - Can turn on and authenticate with a captive portal to enable the internet connection for the Wan Router VM connected to a network with the captive portal.
  * Torrent VM - Watches a mounted folder for torrent files then downloads torrent files to a different mounted folder.
  * File Server VM - Mount network NFS or Samba drives to make them accessible to other VMs or share NFS or Samba drives on a LAN.
  * Microcontroller Utility VM - Connect and control a microcontroller like an Arduino. Provides a server and can interface (i/o) with the connected and detected microcontroller.
  * Gitlab Server VM - Provide local git repositories for development.
  * Backup and Archive VM - Configured for remote backup of specific folders.
  * Wireless Hacking Utility VM - Designed to automatically crack and log whatever wireless AP is in range.
  * IP-over-DNS Utility VM - Receive IP over DNS connections to provide an internet connection to the Multiverse OS network when a normal connection is not available.


## Clustering, assembling a *space station*
Multiverse OS is designed to work with a single bare-metal computer but it also designed to work across several diverse bare-metal CPUs and even remote servers. Specifically, Multiverse OS enables general-use clustering or what is now myoptically marketed as cloud computing, enabling end users to easily take advantage of advances in virtualization and system provisioning automation.

Currently, clustering or cloud computing is primarly used by commerical industries but Multiverse OS enables even "novice" end users to combine two or more home computers into a powerful and secure general use computer. 

To accomplish this, Multiverse OS establishes two categories of bare-metal computer to help organize the computers into a single more powerful computer. The metaphor currently being used is a *space station*, 
!!!!
  * Host Controller - A host installed on a bare-metal computer system to be directly used by the user. Requires a GPU capable of being passed to the Controller VM.
  * Host Auxiliary - A host installed on a bare-metal computer system to be clustered with a host controller computer system.

Additional network interface cards (NICs) can be added to the router VM to increase the total amount of network I/O or dedicate cards for separate isolated networks. 

------

!!!! Didn't read through the roadmap at all yet
## Design Roadmap 

*It works for me, I'm not just referencing the latest CCC event name, it really does* but hopefully soon it will work for you too. Multiverse OS design successully running on my primary desktop while under active development. For maximum inclusivity, support for multiple languages is one of the first planned features. Accompanied with a design structure to encourage community modification.

The current development focus is Multiverse OS is automation of installation and VM management with shell scripts and go language binaries to help others easily install and begin using Multiverse OS.

Contributors can find places to start by reading the source code, and looking for TODO: flags. TODO: Flags explain what needs to be done. One simply needs to add the source code, remoe the TODO: portion of the comment and update the comment text. Then create a pull request.

There also exists folders such as common-scripts, bash-scripts, patches and lockdowns which allow for customization and extension by using accompanying templates to extend the functionality. Please share your patches and improvements.

##### Fundamentals

  * Automated Install Of Every Basic VM Type (Host, Controller, Router, Utility)
    1. Minimal as possible bootstrap that works with as many types of linux/unix as possible
    1. Provisioning framework
    1. Template VMs
      1. Template VMs require entropy-tools package to feed in data from urandom which is not ideal. COnsider designing a very easy entropy generation device using a raspberry pi to feed VMs.
      https://github.com/wadey/cryptorand
    2. Throwaway VMs (based off any template)
    3. Static VMs
    4. Ephemeral VMs
    5. Lockdown scripts - scripts for locking down different aspects of the operating system. There are global lockdowns and per operating system lockdown scripts.
    6. long running defense, schedule vulernability checker, metadata remover
  * Timewarp - take snapshots every x amount of time. Do a live copy, shutdown and snapshot. Then add the ability to right click files and restore from previous version. Delete old snapshots, scheduling defined by yaml settings. When needed, boot snapshot to its own ephemeral vm and do a diff, 
  * Multiverse Browser - A browser where each tab is its own throwaway QEMU/KVM VM or a container 
  * Inter Virtual Machine networking should all happen over VirtIO, because of this honeypots should watch ssh, ping, http and so on to provide a warning system since these communications would be likely only initiated by misfiring malware
  * LAN tools for cooperating with other household members, media center, shared playlists, shared todo lists, shared notes, chat
  * Custom SSHD server, possibly other custom communication servers that honeypot and alert.
  * Distributed file integrity checker between machines over VirtIO
  * Ability to create, edit, update, delete compartmentalized applications
  * Automatic snapshotting 
  * optional ephemeral booting
  * gnome dropdown add-on 
  * controller > host controls - [router reset, switch controller (run controller in bg? just dont use gpu), switch between whonix out and vpn out. add X number of vpn layers for multiple hops, load several times of VPNs for different endpoints
  * GUI to visualize the network and cluster. drag and drop control
  * list of lspci and vfio info
  * Multiverse OS Account System
    https://github.com/cozancin/Bash-Shlet
    1. Consistent configuration system that abstracts below systems
      * Easy to modify list of what processes/daemons start on boot, limit ram and turn back on if crashes.
    2. General user configuration yaml
      * Per VM settings
      * Manage aliases, make it easy to update, list
      * Packages, broken up into templates categories.
      * Process management
      * Selection of prompt
      * Selection of color scheme
      * specify lockdown scripts
      * ssh configs, per server rsa keys, remote servers, hostnames, port forwarding with reverse proxying. use as endpoint
      * Associate GPG, Onion Address
      * onion port forwarding
      * port forwarding
      * crontab -e
      * backup system settings, locations and offsite locations
      * encrypted contacts
      * IM Accounts (xmpp, torchat, ricochet, etc)
      * Wallpaper settings
      * Location of configuration files
        *repository
      * Location of general script files
        *repository
      *Notes/Todo/Calendar
       
      
    3. A CLI tool and GUI tool for configuration file modification
    2. Connection history and type
proxies)
    4. Backup scripts
    5. Note system, todolist, per project, calendar
    5. register IM accounts (xmpp, and shit)
    4. Is Gamer? (Setup ubuntu server with steam)
    5  Is Artist? (Setup OSX for Adobe and music software)
    4. Is Developer? (Change prompt to support git and other features)
    6. General bash script management organized in tree with auto completion
    7. Fast easy to use file search from terminal
  * Folder for scripts of individual system lockdowns for each VM
  * Folder for scripts of individual system patches for each VM install
  * Framework for creating, sharing and deploying scripts to generate Utility VMs and Router VMs
  * Application level firewall like little snitch, possibly using subgraph OS firewall
  * package with subgraph coy IM
  * USBKill - when a special USB drive is removed, can be attached to your arm, the memory is wiped and the computer is deleted.
  * Backup system with integrity checks, backups encrypted and stored on HDs. Special drives are picked up by a VM that runs and sends the backups offiste using .git, spideroak, ssh, etc. Check time stamps, only back up relevant stuff.
  * Outside USB response - if an unregistered USB stick is inserted, wipe the memory and turn off the computer. 
  * modify cryptsetup to boot a decoy OS or Nuke the machine or decoy and Nuke
  * deal with javascript keyboard input anonymity issues
  * optionally Check HD's for activity, if they are not busy after x amount of time then lock
  * 2-factor with stenographic key hidden in picture, that is used to help unlock the cryptsetup, the usb can then be the USBKill/Deadmanswitch 

##### Bash Scripts

Additional bash scripts that could be implemented into as a gnome add-on or i3 script to improve usability.
  * Save all tabs in a firefox window to a list, load all tabs from list.
  * Multiple sessions in a browser for different types of research that could be selected, each session could be divided by VM but the UI presents it as separate selectable, editable and deletable sessions. For example, one session is a collection of windows, tabs and history for magnetic levitation. Another session is a collection of windows tab and history for research on NMR spectrometry.
  * Actively check integrity of files on the system
  * Regularly check for known vulnerabilities

**See notes for more details**

### Account System
A multiverse OS account is nameless but tracks settings and it automates setting up a presence and leaving a machine.

**I3 window layout automated for host, splitting servers across desktops**

**Bash Script auto-complete and save** 

**Bash Script and Configuration Repository**

**GPG Key** can be used to create sub keys for each system, to generate consistent onion addresses and used with `pass` to store passwords. Using pass, the password is not the same across machines, and it is 32-64 random characters used as a shared secret for one-time-passwords. Ideally you would want all your VMs to be LUKs encrypted, with their passwords generated in this way too.

I believe this may provide more security than currently provided by Qubes, since Qubes just has every machine have no root password. Any access is immediately escalated to root, which is argued to be okay since the machines are ephermeral besides the /rw folder. 

It would be nice to do both though, have a /rw folder + ephemeral VMs and secure the fuck out of them. My guess is it is easier to break out of a VM with root access.

**Logout sequence** define scripts for logging out

**Delete traces on user delete** done by set of scripts to remove the user presence on a machine when deleting user.

----

## FAQ
A collection of Q/A generated from research and experimenting. Some questions may only be partially related since there have been major design changes.

##### **Q** How do you enable two monitors over Spice+QXL drivers?

**A** Start by installing the required packages on the Guest operating system.

```
sudo apt-get install spice-vagent xserver-x11-vmmouse xserver-x11-video-qxl
sudo systemctl enable spice-vagent
sudo reboot
```
Edit the xml file using virsh:

`virsh edit guest_os_name`

Change `heads=1` to `heads=2`, each head being a monitor. While you are here turn up the mem on the driver. It seems to be defaulting to 16 MB and virt-manager does not allow you to change it. 

Now you can start the spice session with `remote-viewer` or `virt-viewer` and once you hit X in the client.

##### **Q** How do you enable copy and pasting between host and guest?

**A** Using `xpra` or if using spice install the spice-vagent on the Guest operating system.

##### **Q** How can I directly change the VM settings?

**A** `virsh` is the libvirtd console.

## (Old) Documentation Rough Drafts
Below are a collection of random rough drafts that need to be edited and organized into a general Multiverse OS wiki.

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
Releases are based on Semantic Versioning, and use the format of MAJOR.MINOR.PATCH. In a nutshell, the version will be incremented based on the following:

```
    MAJOR: incompatible and/or major changes, upgraded OS release
    MINOR: backwards-compatible new features and functionality
    PATCH: backwards-compatible bugfixes and package updates
```

______


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

Everyone runs a apt repository from their computer over Tor. IT can be turned off but this also makes it easy to publish software. This could be built as a caddy plugin 

* Grsec needs to be setup in the host machine. 

* The login may be wise to remove. Updating, control and everything else should be handled by an API available via virtIO.

* Debian hardening needs to be set in stone by creating scripts to be run post install.

* Modify the Debian source, change the default apps (like fuck nano, use default vim).

* Replace authentication system, build in a plugin system that will support OTP, key-based login and so on.

* ZFS or Gluster should be used to tie the controller VMs together

* DNS resolution in the routers. So they can pass around the cluster.

* Switch to wayland for Controller VM

* LibreSSL v OpenSSL

* Rip out the logging system, replace it with a system where all logs are in JSON. That way automation can be done based on logs.

* Start a golang based kernel module. This can be used to do various tasks.

* Consider gnome alternatives, like rat-poison.


* Apts to package: vim, electrum, golang, 


* Build a nice default .bashrc with things like:

* An idea to perform remote graphical Debian installations using a web browser

# make less more friendly for non-text input files, see lesspipe(1)
[ -x /usr/bin/lesspipe ] && eval "$(SHELL=/bin/sh lesspipe)"


# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

# Ctrl-N keeps current directory
export PS1='\[$(__vte_ps1)\]'$PS1


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
    

* Plugins!

sudo apt install firefox-esr  mozplugger

* Home Automation: http://www.sweethome3d.com/

This would be cool to use to map out the devices throughout the house.

* Find  personal information management software. It should be able to connect different text files, drop in images, drop in videos, drop in audio

* Is firefox session and form data encrypted when not used? Seriously thats fucked if not.

* push key:

cat ~/.ssh/id_dsa.pub | ssh user1@example.com "cat - >> ~/.ssh/authorized_keys"

=====

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





**NOTE** I don't know where this goes

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





























* Media filetype, analysis, hex tools, reverse engineering tools, etc SHould be able to dive deep into any file for malware anlaysis or modding.
  * Im so fucking tired fo my audio books not being divided from my music, this is simple shit.
  * ALso the file manager in media folers should have more columsn that make sense for the context


  * Digest
* Audio recorder, video recorder

* Special crypto:
  * ecdsa
    * Using merkel tree, create a tree of epehemeral keys, starting from a session key
  * Merkel tree
  * Ring signatures (This will let us revoke faster!!)
  * Blooms, inverse blooms
  * Reed-soloman


  * Socialist Millionare
  * Homomorphic crypto (lets you do math)
  * Stenography - different techniques to hide encrytped data within diferent iamges (can we not change the file size?)

* MEDIA
  *Each over the following 2 functions for (AUDIO, MUSIC, VIDEO, COMIC, BOOK, JOURNAL/MAGAZINE, GAME, 3d-MODEL
  * *Transform*: Upcase, downcase, color, transpile, Lint
  * *Evaluate/Analyze*: Grade, word count, graph, historgram, and other analysis
      >> Dyslexia typo fixer: historgram words, words with fairly large amounts, check against wrods that are nearly the same but dont match the highest percent to find typos

* Transpiler C -> Go, Rust -> Go, C++ -> Go

* Custom SSH server that falls back

* Multiplexing ports, and streams

* Pheonix virtio and libvirt from digital ocean

* Look for comparable REST code for the router, OHT.

* Piece together a Rust or Go Wayland setup

* Look at available options for Rust image builder

* Pick a ficking scripting lanuage for default. Find transpilers from Go to Script. And transpilers between scripts.
  - https://github.com/cmars/glop
  - glop is a "Glue Language for OPerations", a DSL for building autonomous, intelligent agents that operate small, localized software systems.

glop's immediate focus is on managing my own development environment and a densely-packed LXD server.


* Customize xpra or rebuild it in go or rust

* ABD console control (Android)

* Check out options for Android Multiverse OS

* Check out rust options for shell, go options, use the scripting language

* CHeck out rust and go options for malware, get that metasploit shit automated

* Rust of GO FS base
  * Media Collection FS (dedicated DB with dedicated features just for media collection
  * Project FS (dedicated DB/FS just for each prjoect, so each project is completely contextual, this could mean mounting across 3 clustered drives) 



**Files shoudl ALWAYS show size of files inside!** Everything should be compressed and encrypted, with metadata stored with it and map and reductions of data

**App VM that passes over XPRA a clean pdf viewer. no networking et c**

**by default make mkdir use the form or flag that will build directory if missing subdirecotr**

**block great firewall requests** look at toors data to figure out how to drop their probe requets

Checkout vpn gate

Hiding TLS traffic, make a fake DNS request, then look at the server who you made the fake DNS request then make all your traffic match the highest priority crypto algorithm, then wrap all your future requests with this algo, and use a proxy server rougly in the location of the server you are pretending to connect to.


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

**computers like routers that should never be touched and the host, should have a special kernel module to block any access to the /dev/uinput. This way even if there was access, it wouldn't even be able to type, it should also inject randomness to the typing so it would never do anything. alerts then are triggered to alert intrustion**

**benchmark** bnechmark kuberetes with nvidia docker image vs multiverse controller

https://github.com/tokuhirom/jailingo - jailing with chroot in go

https://github.com/appc/spec - rebuild this spec to be secure and work for multiverse



https://github.com/kavu/go_reuseport
GO_REUSEPORT is a little expirement to create a net.Listener that supports SO_REUSEPORT socket option.

For now, Darwin and Linux (from 3.9) systems are supported. I'll be pleased if you'll test other systems and tell me the results. documentation on godoc.org.
One of the features merged in the 3.9 development cycle was TCP and UDP support for the SO_REUSEPORT socket option; that support was implemented in a series of patches by Tom Herbert. The new socket option allows multiple sockets on the same host to bind to the same port, and is intended to improve the performance of multithreaded network server applications running on top of multicore systems.

The basic concept of SO_REUSEPORT is simple enough. Multiple servers (processes or threads) can bind to the same port if they each set the option as follows:

    int sfd = socket(domain, socktype, 0);

    int optval = 1;
    setsockopt(sfd, SOL_SOCKET, SO_REUSEPORT, &optval, sizeof(optval));

    bind(sfd, (struct sockaddr *) &addr, addrlen);


* Read aristotles poetics

* https://github.com/bieber/drones  A drone combat game based on a programmable virtual machine. 

> "Always be mindful of implicit trust in binary packages and files. In other words, it is essential to be aware of how the initial bootstrap process verifies digital signatures. Never skip package or metadata verification simply because it's more convenient to do so."

This software is not meant to fill the same niche as enterprise, industrial, commercial clustering software like kubernetes, and others. 

Multiverse is a general use operating system designed to present a modern and easy-to-use and learn interface for novice to advanced users.

----

 relativistic time dilation would result in the accelerated wormhole mouth aging less than the stationary one as seen by an external observer, similar to what is seen in the twin paradox. However, time connects differently through the wormhole than outside it, so that synchronized clocks at each mouth will remain synchronized to someone traveling through the wormhole itself, no matter how the mouths move around.[32] This means that anything which entered the accelerated wormhole mouth would exit the stationary one at a point in time prior to its entry.

Lorentzian wormholes (named after Hendrik Lorentz), also known as Schwarzschild wormholes or Einstein–Rosen bridges, are connections between areas of space that can be modeled as vacuum solutions to the Einstein field equations, and that are now understood to be intrinsic parts of the maximally extended version of the Schwarzschild metric describing an eternal[clarification needed] black hole with no charge and no rotation. Here, "maximally extended" refers to the idea that the space-time should not have any "edges": it should be possible to continue this path arbitrarily far into the particle's future or past for any possible trajectory of a free-falling particle (following a geodesic in the spacetime), unless the trajectory hits a gravitational singularity like the one at the center of the black hole's interior.

>An intra-universe wormhole (a wormhole between two points in the same universe) is a compact region of spacetime whose boundary is topologically trivial, but whose interior is not simply connected. Formalizing this idea leads to definitions such as the following, taken from Matt Visser's Lorentzian Wormholes (1996).

It makes me kinda sick when I see projects that justify their existence based around their role in helping business or importance to business as if capitalism or corporations were the protaganist of history.

> "Lot's of people wonder why I don't use ssh, the reason is that ssh requires both the host and the guest computer to spend effort on encrypting the data stream that is only private between them anyway. The poor things are working hard enough compiling the horendously huge viewer, the less overhead the better. So the only thing using ssh gets you is to slow things down." - Interesting

*Check if its feasible or useful to interpret Dockerfile Terraformfile or other OCI stuff. So people can migrate over easier.*

*Built in folder watching + exuection* So many times now, and helping so many friends accomplish the same task over and over. The ability to watch a folder and preform a task. This should be built into the operating system. One should be able to watch a folder and add torrents, wait for downloads and upload them to your mobile device, preform organization of media and setup symbolic links back to the downloads folder and so on.

**AI assistant?** https://github.com/itsabot/itsabot

Github finally supports mirroring from other repos

**https://github.com/tcnksm/gcli** gcli generates a skeleton (codes and its directory structure) you need to start building Command Line Interface (CLI) tool by Golang right out of the box. You can use your favorite CLI framework. [These generators can be modified to build a Multiverse OS system tool that meets all the minimum design specifications, providing a skeleton that will guide the developer in the process of staying within convention so the design spec can be satisifed] *what is sweet about this software is that it supports a framework, meaning it can support a multiverse-os framework without having to throwing away all the updates they would add to the software*

Multiverse OS developers reject the current internet paradigm, reject the idea that there is a single path for technology to progress, and reject the centralized services like AWS hosting too much of the internet, monocultures datacenters, the move towards less anonymity, move towards containers instead of secure VMs, the move away from self hosting, the move away from hosting from home, reject the concept of rebuilding everything in Javascript instead of finding a way to use existing tested, secure and better preforming software. Multiverse OS is more than just a secure operating system, Multiverse OS is using new technology to present a new way to approach internet computing and more importantly a new way to approach applications and services on the internet, a new way for the the browser to present these applications. This will change the way in which we network, host, and use internet services. It will untether us from only having HTML/CSS/JS and suffering from the limitations those have. Decentralization without needing a blockchain specifically, but thats not to say we won't build in Bitcoin support to provide the economic backbone to this new vision of internet applications. *Design for the current internet is controlled by corporations. Often ignored concept is that the way a software is funded in the economic predominante economci system has a direct and functional impact on the software, this is true for everything. The result is that much of the modern internet revolves around getting people to look at ads, this is more important than progress. Entretched powers stopping progress is not new, its the status quo.*

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

```
qemu-kvm -m 512 -smp 2 -bios /usr/share/edk2/ovmf/OVMF_CODE.fd -drive format=raw,file=image.raw
```
