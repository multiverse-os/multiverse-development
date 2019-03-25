## Multiverse

"...same planet, different dimension...", multiverse allows processes to be executed ub completely in parellel in different isolated operating systems. Each operating system has its own network, some their own routers and internet connections. Multiverse OS utilizes virtualized clustering to run software interacting with the internet in ephemeral virtualized environments controlled with extremely rigid access control from ephemeral virtualized environmens.

####  The Multiverse Operating System <style="font-size:16px;">is intended to be an easy-to-use, general use with reliable security provided by compartmentalized ephemeral VMs nested within isolated networks with secure microkernel VM routers. As much as possible complexity from the security model is transparently hidden within the UI of a full featured Desktop Environment based on Gnome or I3, capable of per application optional pseudoanonymous full VM sandboxing within ephemeral VMs. The network of virtual machines that support Multiverse OS can be deployed across several diverse bare-metal servers to create clusters of secure VMs.</style>

*Multiverse OS draws inspirtation from previous secure operating systems such as Tails OS, Whonix and QubesOS. An easy way to conceptualize Multiverse OS design is that it is like running multiple Debian based QubesOS dom0 hosts in virtual machines in parallel. The user interacts with Multiverse OS from within a virtual machine with PCI devices passed, resulting in absolutely zero user interaction with the bare-metal host hypervisor. Each `controller-vm` running on a bare-metal host, has its own multiple customized, isolated networks. Like QubesOS the user experience is designed to be like a traditional linux desktop environment experience, with as much of the complexity as possible hidden away from users.*

## Identifying Modern Computer Security Problems
Multiverse OS was created in response to research into the growing sophistication, automation, and capital inflow into malware, root-kits, worms and viruses. The rate of increasing sophistication is astoundingi, kids, black hats and mad scientists have been joined by state actors, private intelligence corporations and the growing ransomware industry. Meanwhile, most operating systems have failed to keep up with improving operating system defenses for "average" users.

A system can be hardended but even "advanced" users are affected, as it is increasingly harder to escape the ubiquitiousness of the growing XSS attack surface they interact with each day when browsing the internet.

In the current context of the internet, it is best to assume every website could be exploited. The operating system should be paranoid, so the user does not have to be. A free, open-source, easy-to-use hardened OS secured with ephemeral compartmentalization using nested VMs is needed to address the emerging persistent automated malware threat facing both *"average"* and *"advanced"* users:

* A growing number of sophisticated root kits, malware and automated viruses are being written by state intelligence, private intelligence, organized crime and corrupt companies interested in access to the large inflow of capital into offensive computer security. Anytime a port is opened on a computer with internet access on a WAN IP connections will be made almost immediately from bots investigating. 
* A growing number of web applications over utilize Javascript implemented in a rush resulting in insecure applications vulnerable to cross-site-scripting (XSS) vulnerabilities. The problem is amplified by large numbers of poorly implemented server configurations and growing numbers of sites failing to provide working alternatives that do not require Javascript. Alternatives are not even available when the javascript is being used solely to 'improve' the UI. The state of the internet forces internet users into regularly running code from anonymous untrusted parties on their machines. Even if the hosts are not malicious, poor Javascript regularly exposes users to XSS attacks.
* Limited infrastructure to validate Javascript loaded from other domains, and many people relying CDNs and conglomerate hosting creates central points of weakness. Javascript is even included in advertisements, which allows malware programmers to spread through ads on pages.
* Vulerable/Exploitable software anywhere in the system comprimises all the data on the system including: cryptography keys, personal photos, personal documents, source code, browsing history and more. 
 The modern Internet ecosystem, regardless of private browsing mode, do-not-track options, or even claims that search engines do not track, allow for user tracking using almost anything imaginable to create unique fingerprints. Including browser settings, connection-pattern, DNS queries and others compromize anonymity. This data unncessarily gives away valuable private data from location, social connections, browsing habits, interests, research done and more.
* Clustering (creating a cloud) securely is not easy with diverse collections of hardware and not typically designed for general home use. Many tools for for home LAN networks are typically lacking even basic security.

Given this context in which we browse the internet, it is just a matter of time before you execute malicious javascript that directs you to a page with reverse shell exploits and you are exposed to a vulnerability you did not yet patch.

#### Why not just use QubesOS?

After using Qubes OS, from my experience, I believe it may be the most secure open source operating system currently available.

However, after months of use, reading code, reading forums, I came to the conclusion that the QubesOS developers have different requirements for their operating system than I have for mine.

While QubesOS provides a highly secure environment, generally works well for development and general internet browsing, it did not provide all the features I needed for my general internet use. Additionally, while it is a very secure choice of operating system, there is a lot of room for improvement on the current design.

1. **Multiverse OS should use Debian, to benefit from scripts written for Tails OS and Whonix OS** instead of Fedora.
2. **libvirtd for interfacing with QEMU/KVM** instead of QubesOS choice of libvirtd communicating with Xen. Xen has better preformance but QEMU/KVM has better support for total number of GPU's capable of passthrough.
1. **Multiverse OS should use VirtIO for inter-VM networking instead of SSH** it would be better to implement communication over virtIO and disable all typical tcp/ip communication to limit virus mobility. 
2. **Multiverse OS should use passwords, one-time-passwords, HMAC signatures, and pass to generate and gpg encrypt 64 random luks hard drive passwords**, validation and verification can happen over VirtIO connections.
3. **Multiverse OS should use Alpine Linux security focused microkernel for routers and most utility vms not Fedora**. Fedora Desktop has a much larger attack surface than Alpine linux, a security focused microkernel packaged with gr security.
3. **Mutliverse OS should be built primarily with shell and compiled languages like go language, c or rust** instead of QubesOS choice of python. Python does not compile into binaries and go language has very mature libraries for working with libvirtd, qemu and kvm. 
3. **Multiverse OS requires wide support for GPU PCI-passthrough** because using GPU passthrough in combination with locking down the bare metal host makes the attack surface of the host (dom0 equivilent) almost non-existant since the user never logs into it, or directly interacts with it. 
3. **Additional support, utility-vms and work done to support general use for "average" users including features like torrenting, linux gaming, developing and scientific paralell processing programming use cases securely** 
4. **Multiverse OS should by default use a different shell designed for to seemlessly send commands locally and across the Multiverse OS internal network. 

The Multiverse shell should support improved system of cooperative computer usage, building on what tmux has given a glimpse of.

Ideally it should be built with a language that is elegant, easy to read, wraps fast binaries in a consistent way, configures in a consistent way (likely through a collection of config.yaml files) and is very quickly written to speed up development.

It should also be very quick to program in. Ideally a language like Ruby would be ideal if it provides sufficient security (the shell should be capable of restricting commands to the point of jailing the shell).  Ruby is a good choice because even none programmers can typically make sense of its code, and yaml files are very novice user friendly. 

Using Ruby, Go or Rust a shell can be built a long with a reliable framework that scripts can tie into. Both simplifying and securing the routine tasks repeated over and over like input validations, ouput santization, and so on. Leaving the scripter left with just gluing together the peices needed for the scale of their application and the simple easy to read logic that reads like a a recipie. 

Ruby would be advantageous since it is also the language used by both Chef, a software for provisioning operating systems and Metasploit a framework for executing offensive computer actions. 

While designing this aspect of the project, I have come to the opinon that shell advancements have stalled because the languages that they depend on have not reached the levels of the other higher level programming languages. Bash 4 adds a lot of new features, and there are an increasing number of repositories for scripts, the infrastructure is still years behind even new higher level operating systems.

For example, being able to tie provisioning these in with a distributed shell, that can send commands based on tags or other relations, could enable users to use Multiverse OS in unexpected ways. 

One major difference already planned is bringing more types of relationships and associations between the files beyond a simple tree structure.

File assocation should be complex, including frequency of use, proximity, usage patterns, direction of the shell user (cd ../../modules vs. modules/good-module/submodule.. - an up vs down direction in the tree should indicate what possible files the user could be looking for. It should be even capable of jumping across Multiverse controller-vms seemlessly.

And realtime asynchronous interaction with other active anonymous users could create more efficient workflow. 

The shell should not just be for browsing files, but also for browsing available executable code, organized in a tree and otherwise. Suggestions appearing live as you navigate, loading and unloading the libraries as a user navigates, uses and over time based on frequency of use and other indicators.

Scripts should be obtained through a distributed package management service, information about what is available would be passed through the onion network. And download checksums would be downlaoded across multiple peers anonymously through the onion. While this introduces a small sybil attack surface, these computers are also ephemeral and and isolated from personal persistently stored data. 

>"The Sliders are on a trail of jumping between parallel Earths using a Einstein-Rosn-Podolsky bridge (a wormhole). Objects or people that travel through the wormhole begin and end in the same location geographically and chronologically."

#### A different approach to QubesOS-like security
Multiverse OS focus is an operating system that is easy to use, developed defaults that are built with general use computing needs in mind, rich customization/development environment, high security kvm/qemu compartmentalization with nested ephemeral VMs.

Multiverse OS run on `Debian8 Stretch/Main` using `libvirtd` to compartmentalize QEMU/KVM VMs, with potential support for LXC, LXD, rkt and nspawn containers.

**Input/Output Passed To Controller VM**

The physical USB ports, GPU, audio are assigned to `controller VMs` using PCI-passthrough running a nested Mutliverse OS Debian8 installation.

**Application Sandboxing In Controller VM**

Applications run within are run within QEMU/KVM virtual machines and accessible to users by using `xpra` over TCP. This method does not provide a shell to the connecting user and the virtual machines can be run from separate unpriviledged user accounts. Almost complete networking may be possible too if VirtIO could be used to provide a `xpra` tcp port.

Containers may be incorporated through libvirtd, adding the possibility for additional layers of isolation by nesting containers inside of VMs or VMs inside of containers.

*Containers in their current state appear to trade off security for increased preformance, and so with sandbox effectiveness being critical to the Multiverse OS security design, containers alone should not be used. If anything has the ability to increase effective isolation should be used or provided as an option.*

Efforts are being made to make this transparent to the user.

**Layered Networking**

The physical network cards are assigned to router VMs using PCI-passthrough running security focused microkernel `Alpine Linu` and layered to provide isolated subnets that can be isolated from the LAN, transparently proxied through VPNs, Tor or remote servers.

**Long Running Utility VM In Segregated Networks**

Long running applications requiring no human interaction and special networking are isolated from the `controller-vms` and stored in `utility vms`. 

For example, torrent applications are run within a `utility vm`, downloading torrents by being provisioned to mount with minimal priviledges. Watch the user's browser Download folder and download the resulting file into the Torrent folder. 

By segregating these long running utilities, more general uses emerge, allowing for a greater number of people securely access the internet.

Efforts are being made to make this transparent to the user.

#### Types of VMs
Multiverse OS uses four primary types of VMs to simplify management, each of these types have their own sub-types. These VMs provide the building blocks needed to provide functionality for a wide variety of use cases.

1. **A Host VM**
  * Host Controller - A host installed on a bare-metal computer system to be directly used by the computer that has a GPU capable of being passed to the Controller VM.
  * Host Auxiliary - A host isntalled on a bare-metal computer system to be clustered with a host controller computer system.
2. **A Router VM** A router VM is currently primarily being built from the Alpine Linux microkernel. This OS is incredibly small, security focused and uses minimal resources. The primary router VM recieved the bare-metal NIC by utilizing the PCI-Passthrough feature.
  * WAN Router VM - Designed to receive NIC card and manage connection with other routers or provide limited access to utility VMs that need LAN internet access.
  * Firewall Router VM - The firewall and proxy server This is where your openvpn client will be setup. Ideally you are connecting to a server you rent and put a VPN on yourself or you are using a reputable VPN seller that takes Bitcoin. The whonix box will route its traffic through this to hide Tor access from your ISP.
  * Isolating VPN Router VM
  * Whonix Router VM

##### Notes on provisioning Alpine Linux for use as a router

>*"Still, we (along with our myriad of virtual clones and alters) are invited to embrace this wider canvas, this grander vision. Our vision of the multiverse is more vast than even a few years ago as there is evidence for many possible forms a universe can take."


Setup alpine, install vim, install shorewall, copy shorewall configs, setup dhcp, enable shorewall

Alpine linux does not just pick up the routes (depends on the settings of the virbrX), so you need to manually define them in the interfaces file. This is assuming your LAN is 192.168.1.1, which is pretty standard. 

`/etc/network/interfaces`

```
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet dhcp
	#up ip route add net 0.0.0.0 netmask 0.0.0.0 gw 10.1.1.1
	dns-nameservers 8.8.8.8 8.8.4.4
	hostname sys-firewall
	up route add -net 192.168.1.0 netmask 255.255.255.0 gw 10.1.1.1
	up route add -net 0.0.0.0 netmask 0.0.0.0 gw 10.1.1.1
```
2. **A Controller VM** A controller VM received the GPU and like the host VM does not access the internet directly. It has very limited control through VirtIO interface, verfied by one-time-passwords, HMAC and nonces to prevent replays. Application VMs are created within Controller VMs as unpriviledged user.
3. **An Application VM** An application VM is by default an emphemeral VM that starts applications within containers/jailshells and exposes window GUI access using xpra TCP connections. TCP connections allow access to the appliaction without giving out full SSH shell access.
  * Vault - Has either Electrum or Full Bitcoin node running in a VM to isolate it from computers that browse the internet. 
  * MultiVault - An application VM that supplies an interface to several isolated VMs with different cryptocurrencies stored.
4. **An Utility VM**
  * USB protection Utility VM - Designed to isolate unknown USB drives in a contained and monitored environment. This VM ideally has all the tools necessary for malware analysis installed.
  * Captive Portal Utility VM - Designed to turn on and connect to a captive portal to enable the internet connection for the wan-router connected to the network with the captive portal
  * Torrent VM - Designed to watch a mounted folder for torrent files then downloads torrent files to a different mounted folder.
  * LAN File Sharing VM - Designed to mount network NFS or Samba drives to make them accessible to other VMs or share NFS or Samba drives to the network.
  * Wireless Hacking Utility VM - Designed to automatically crack and log whatever wireless AP is in range.
  * DNS-over-IP Utility VM - Designed to receive DNS over IP connections, when setup it can be used from captive portals to access the internet through the network of the Multiverse OS system.
  * Microcontroller Utility VM - Designed to connect and control a microcontroller like an arduino. Provides a server and can interface (i/o) with the connected and detected microcontroller.
  * Gitlab Server VM - Designed to provide local git repositories for development

**The multiverse metaphor** is used to help explain to uninitiated and to help developers imagine the project using the terms used in physics to describe the concept of multiple universes living side by side. Learning a little more about multiverse theory we can incorporate terms and concepts that help illustrate parts of the design while coloring the project.


There is a hierarchy of observational levels:
  * subquantal
  * quantal
  * photonic
  * atomic
  * electromagnetic
  * chemical
  * cellular
  * organismic (multicellular)
  * (maybe community should be put here, then increasing complex social structures)
  * planetary
  * solar systems
  * galaxies
  * universes

Currently, each router that has access to a physical network card is considered a universe, within that universe there is a an area outside the galaxy where utility VMs that need LAN access reside. I have taken to naming these "voyager". 

#### Utility VMs
Utility VMs should provide an interface for the controller VM. That way for example the controller VM for the torrent VM receives enhancements like terminal commands to add magnet files by url or hash. 

Each registered Utility VM should expose a interface. Optional but ideally each running utility provides access to its data without requiring /dev/input (mouse and keyboard) type input.

It could be a HTTP page that contains scripts and possible interface.

**Lan Maker VM**

gpio (to create a utility VM for connecting to raspberry pis)
https://github.com/platinasystems/gpio  -- minimalist lib

**Captive Portal VM**

Simple VM to login into captive portal, attaches to router.universe (LAN)

Kicks up a desposable firefox (also has a chromium browser incase the portal is poorly written)

**IP-over-DNS server** Can be running so you can obtain internet access over captive portals

**Backup and Archive VM**

A VM which easily configures, using YAML files (and a UI to modify YAML files, define which folders get uploaded where) to backup via:
* rsync to remote servers (ssh)
* spideroak
* github
* ?

**Example** *.yaml configuration*

```
servers:
  server:
    # Optional Name
    name: "LA"
    host: 125.125.125.1
    type: rsync
    key: ~/.ssh/chance.pub
    directores: ["~/keys", "~/pictures"]
```

## Clustering, assembling a *space station*

Multiverse OS is designed to work with a single bare-metal computer but it also designed to work across several diverse bare-metal CPUs and even remote servers. Specifically, Multiverse OS enables
general use clustering or what is now myoptically marketed as cloud computing, enabling end users to easily take advantage of advances in virtualization and system provisioning automation.

Currently, clustering or cloud computing is primarly used by commerical industries but Multiverse OS enables even "novice" end users to combine two or more home computers into a powerful and secure general use computer. 

To accomplish this, Multiverse OS establishes two categories of bare-metal computer to help organize the computers into a single more powerful computer. The metaphor currently being used is a *space station*, 

A "controller" and "auxillary"

Additional network interface cards (NICs) can be added to the router VM to increase the total amount of network I/O or dedicate cards for separate isolated networks. 

------

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
  * Save all tabs in a filefox window to a list, load all tabs from list.
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

`virsh edit skylab`

Change `heads=1` to `heads=2`, each head being a monitor. While you are here turn up the mem on the driver. It seems to be defaulting to 16 MB and virt-manager does not allow you to change it. 

Now you can start the spice session with `remote-viewer` or `virt-viewer` and once you hit X in the client.

##### **Q** How do you enable copy and pasting between host and guest?

**A** Using `xpra` or if using spice install the spice-vagent on the Guest operating system.

##### **Q** How can I directly change the VM settings?

**A** `virsh` is the libvirtd console.


# Provisioning

## Host Deploy Flow

# --**Install Steps on bare-metal host-controller** -------------------
# == Script Environment Preparation ==
# * Check if multiverse.yaml config file exists, if true: parse yaml
# * Check if system meets minimum requirments, confirm continue Y/N
# * Check if already installed
# * Check if user wants to customize installation Y/N
# * Allow user to set each installation setting
# * Run (Before) Installation patches
# * Read multiverse.yaml file if it exists
# * Read Language Settings and load language
# == Install Dependencies ==
# * Update /etc/sources/apt, set repository to stretch/main
# * apt-get update, upgrade, dist-upgrade
# == Configure Libvirt ==
# * update grub to include iommu for pci-passthrough (amd or intel depends on settings)
# == Router Creation == 
# * Create WAN Router VM using Alipine linux and build script
# Create Isolating VM using Alpine linux and a build script

# Use a Whonix Gateway to provide the third router. Do checksums and signature checking. Or eventually just implement in Alpine. 

# Start all routers on boot

# Create storage pool for each hard drive, mapped

# Create storage pool for linux distro images 

# Create controller VM using Debian (Multiverse OS)

# run lockdown scripts
#  - turn off all processes using any port not being used
#  - install shorewall configuration files  

# Identify PCI devices

# Generate vfio and/or rc.local file to pass devices. 

# ----------- steps on controller VM -------------------
# == Install Dependencies ==
# * Update /etc/sources/apt, set repository to stretch/main
# * apt-get update, upgrade, dist-upgrade
# == Configure Libvirt ==
# * update grub to include iommu for pci-passthrough (amd or intel depends on settings)

# == Router Creation == 



# Create WAN VM Router using Alpine linux and a build script
# Use a Whonix Gateway to provide the third router. Do checksums and signature checking. Or eventually just implement in Alpine. 
# Start routers on boot

# Create Template VM for general use
# - Install xpra, install entropy-tools
# Create Template for Utility and Router VM use

# Install VM management software and xpra launcher on the controller VM

# Each proxy VM (as in proxy computer not proxy networking) should have an attached downloads
#folder which should make transfering files uncessary. Home folders can be shared if necessary
#for special projects.



base it off a template an make them ephemeral. hold passwords in the host0 machine using password-store (pass). 

remove firefox from skylab (controller VM), use firefox through a proxy vm shared of xpra

xpra start ssh:SERVERHOSTNAME:100 --start-child=xterm

### dont use sudo on these machines ###@ proably not even on skylab, just use su and store the password, this makes priviledge escalation slightly more difficult.


auto setup VM
1. start from a template, that started with xfce debian. delete xfce, install i3. static ip to one not in active use.
2. ssh into the machine, add ssh key, change the ip address.
2. write ip address to hosts file of skylab
format now is 0.proxy 1.proxy, as in 10.255.255.100, 10.255.255.101
3. Install XPRA so it can actually connect
3. generate difficult root and user passwords then store them in pass, just use the SSH key to get in and muck around

# make i3 windows being passed through xpra look less shitty
sudo apt-get install lxappearance gtk-chtheme qt4-qtconfig


Start with lxappearance and choose a theme; then choose it in gtk-chtheme. In qt4-config, there is a dropdown menu setting to make qt take the GTK+ settings. That seems to work best for me. 

install droid font, maybe others, not exactly sure what i ended up using

# want sound?

edit the xpra.conf on the client `/etc/xpra/xpra.conf`, install the newest version of xpra. probably need to come up with a simple way to build this from source using a script so we do not have to make every machine stretch (testing) debian instead of the current stable.

a flag will need to be passed to xpra to start with sound

# want clipboard?

It may be best to watch a keystroke and turn off and on clipboard sharing only when its requested. I beleive this is how qubes does it and makes it significantly more secure even if slightly more annoying to use. this could be optional and you could have just always on too as an option.

4. select which folders should be shared with the machine and if they are read only or not. load everything read only if possible (Default, Mapped (most secure), 

Write /etc/fstab file so it auto loads the downloads folder and possibly the development folder


random number generation becomes a problem on template VMs

I installed rng-tools and edited /etc/default/rng and added urandom. but we should generate some entropy with the mic or something

probably should find a good source of encropy that cna be fed to VMs, and at the very least the above has to be in the base template or shit just wont work.

you also want the newest version of xpra for a variety of reasons and its on debian testing




3. register applications to be used by this server

so assign applications that are going to be used so a menu can be built and xpra scripts can be generated for that menu. 


----- general debian provisioning notes ---------------
apt-get install vim sudo git

apt-get remove nano

--- remove gnome games --

sudo apt-get remove -y hitori gnome-chess gnome-tetravex four-in-a-row five-or-more gnome-nibbles gnome-taquin gnome-robots gnome-games

### would like to control from skylab ###
-- host0 control
1. Control the VPN being used by the sys-firewall (Add support of ssh proxying)
2. Control what is being shared on voyager
3. Control the whonix firewall, flip identity
4. trigger backup server
-- skylab VM control
(should we do a second VPN on the skylab box?)
1. switch between openVPN and openVPN+whonix connectivity
1. generate app proxy VM
2. control the whonix firewall, flip identity
3. move Vms around network 


-------------- how do we get libvirt to connect to remote shit? ----------------
virsh -c qemu+ssh://user@10.10.10.254/session?socket=/run/user/1000/libvirt/libvirt-sock


# To enable it in virt-manager do:
# gsettings set org.virt-manager.virt-manager.connections uris "['qemu+ssh://user@host0/session?socket=/run/user/1000/libvirt/libvirt-sock', 'qemu:///session', 'qemu:///system']"


---------- on setup of the main box (need a switch to indicate if the setup is for a main controller box or just an auxillary server ----- we need to setup a key for logging into the second controller server. pulling this out needs to have a usb kill effect, and shut off the server, wipe memory and even do this for the auxillary servers. maybe those servers maintain a connection and if it doesnt get an acknowledgement after a while it shuts down so a network attack cant keep auxillary servers online



##  Multiverse OS 
##
##  Isolated, compartmentalized secure general use operating system
## 
####################################################################################
=# Introduction #===================================================================

## Why 
Multiverse OS was created after realizing the need for compartmentalizing several open
source development projects and to defend against the growing number of sophisticated
vectors such as XSS frameworks combined with metasploit-like frameworks. 

After experimenting for months with Qubes OS, I realized this style of OS was the 
best available solution to the problem described above. However I decided it did 
not fit many of my requirements and could be vastly improved. So I started developing
an alternative solution. My focus would be on making an operating system that fit the
use cases of typical computer users, games, torrents, development, local file sharing
 and more. Some of the major criticisms I had of Qubes OS was the use of Fedora for 
routers, inability to pass through GPUs, passwordless setup, use of fedora on the host,
and most importantly how the host and virtual machines interact. I do not claim to be
a security expert, this is an experiment but for my use case I believe it is better.
Use this at your own risk, do not trust this with your life or freedom. 


## Basic Design
Multiverse OS intentionally prevents users from ever directly using the bare-metal
host operating system. Instead by using nested virtual machines (VMs), the user always
operates within VMs. The "controller" VM is passed the bare-metal host GPU which gives
the VM near native speed, while preventing potentially dangerous GPU drivers from 
being installed on the bare-metal host.

VMs "above" and "below" the "controller" VM are only accessible through very limited
API's exposed through the VirtIO bus (is it technically a bus?). Using the VirtIO bus
is fast, obstuficatd since HTTP and SSH are more common transports and since it is
already used by the existing system, it minimally increases the attack surface.

Multiverse OS uses four primary types of VMs to simplify management, each of these 
primary types of VM has several modular sub-types. These VMs provide functionality
for a wide variety of use cases.

1. **A Router VM** A router VM is currently primarily being built from the Alpine 
Linux microkernel. This OS is incredibly small, security focused and uses minimal
resources. The primary router VM recieved the bare-metal NIC by utilizing the
pci-passthrough feature.

2. **A Controller VM** A co

2. **A Application VM**

3. **A Utility VM**

### Clustering, assembling a *space station*

Multiverse OS is designed to work with a single bare-metal computer but it also
designed to work across several bare-metal CPUs. Specifically, Multiverse OS enables
general use clustering or what is now myoptically marketed as cloud computing, enabling
end users to easily take advantage of advances in virtualization and systems automizaiton.

Currently, clustering or cloud computing is primarly used by commerical industries
but Multiverse OS enables even novice end users to combine two or more home computers
into a powerful and secure general use computer. 

To accomplish this, Multiverse OS establishes two categories of bare-metal computer to help organize the
computers into a single more powerful computer. The metaphor currently being used
is a *space station*, 


A "controller" and "auxillary"

Additional network interface cards (NICs) can be added to


Setup passthrough by enabling ioummu, enabling vfio-pci (vfio is for 4.0+ kerneles and pci-stub is for previous versions).

Then echo the `lspci -n` 00:00.00.0 number into unbind. Then pass 0000 0000 number into the vfio-set bind, which creates a `/dev/vfio/##`. 


Before (or after) the networking is setup, qemu must be configured to allow access to the bridges to unprividedged users on the host. This is important because it prevents breakouts from immediately acheiving root access on the host. That you can edit `/etc/qemu/bridge.conf` to add `allow virbr0`.

`sudo virsh net-add multiverse`

Then paste in below, switch out the macs. The DNS is currently routing to google, it would be better to route DNS requets to Tor as it provides a cheap, easy acccess distributed DNS setup. This could be done at the top router level.

I decided on a cosmic naming scheme for my setup. Multiverse OS has 3 network levels:

1. Universe-0 - Universe prime is the computer which you operate, the one you are actively at. Others will start with Universe-1 and so on. This is the PCI passthrough level, local area network is accessible. Good for torrent machines (depending on region and stealthiness of your torrent site of choice. A simple VM with torrent client and samba shares for the media center goes here.

2. Galaxy-0 - Galaxy prime is the network layer which is protected, all traffic is isolation proxyified with an openvpn connection. This can be your normal internet traffic for searching or when something is very tedious to do with Tor. This traffic can not access the lan.

3. Sol-0 The safest place, close to Earth, all traffic passes through Tor. This traffic can not access galaxy or lan networks. It is isolated, but several workstations can be behind the same whonix box and serve stuff to galaxy or lan via onion services. 

*Galaxy and Sol could be combined, and a switch could switch between VPN or Tor but never allow direct access to Universe. Multiple second level Firewall/Proxy VMs could be deployed below Universe. Multiple Universe routers could be deployed if the server has multiple physical NIC cards to assign. The initial server has three NIC cards to distribute, but these are all assigned currently to universe0*

Modify the rc.local to automatically do this. I have seen scritps that bind everything, it would be nice to find one that would bind essentially all ethernet,wireless,etc devices.

`/etc/rc.local`

```
echo "0000:03:00.0" > /sys/bus/pci/devices/0000\:03\:00.0/driver/unbind
echo "1969 e091" > /sys/bus/pci/drivers/vfio-pci/new_id

exit 0
```

When I do `lspci -n` I get

03:00.0 0200: 1969:e091 on the same line that doing just lspci showed the NIC card. Just match the 03:00.0 part. 

One for the PCI passthrough, one for the proxy/firewall VM, and one for whonix. This will let you have fine grain control over how your VM accesses the internet.

`/etc/qemu/bridge.conf`


```
allow virbr0
allow virbr1
allow virbr2
```


 That lets the unprividedged user create devices connected to it. For alpine linux, which I used for my routers. I think it was a better choice than say ddWRT or a larger OS. 

```
apk update
apk add vim
```

Setup the alpine box, `sys-net` needs a static route set.



#### net-firewall, the firewall and proxy server

This is where your openvpn client will be setup. Ideally you are connecting to a server you rent and put a VPN on yourself or you are using a reputable VPN seller that takes Bitcoin. The whonix box will route its traffic through this to hide Tor access from your ISP.

Alpine linux does not just pick up the routes, so you need to manually define them in the interfaces file. This is assuming your LAN is 192.168.1.1, which is pretty standard. 


`/etc/network/interfaces`



```
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet dhcp
	#up ip route add net 0.0.0.0 netmask 0.0.0.0 gw 10.1.1.1
	dns-nameservers 8.8.8.8 8.8.4.4
	hostname sys-firewall
	up route add -net 192.168.1.0 netmask 255.255.255.0 gw 10.1.1.1
	up route add -net 0.0.0.0 netmask 0.0.0.0 gw 10.1.1.1
```


**Remaining steps**

* openvpn on firewall/proxy alpine linux VM

* setup whonix box, route through firewall/proxy

* firewall the host, only outbound SSH connections allowed. on sys-firewall and sys-net disable access to it.

automatically setup the DNS at universe-0's router, then all names should go there or the router below it. This way connecting like ssh user@world-2 will work 

automatically remove meta data from images in a specific folder? optional all 

