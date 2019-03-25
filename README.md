<p align="center"><img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4" height="300" width="300"></p>
This repository contains the final manual install guide, and large collection of old notes, research, designs, and more from over a year of design and development. The final guide will then be used as the basis for the first alpha version of the Multiverse OS installer image.

Actively, Multiverse OS developers are finishing up a clear containers implementation specifically for Multiverse OS called `portalgun` but earlier designs only used KVM/Qemu and the final manual install guide and alpha Multiverse OS installer built from this guide will use KVM/Qemu.
 

```
`-eq` is for comparing strings
`=` is for comparing strings
```



"Stop talking about the code, and lets run some tests."

*A very brief and limited rough-draft summary*


____________________________________________________________________________
## Notes/Development

  * We should move towards only sharing a single p9 share named:
    `multiverse-portal` or `mverse-portal` or `mv-portal`

    And link in any additional folders to the primary one. This would simplify
    the `/etc/fstab` and make it consistent enough to generate it automatically.


____________________________________________________________________________
## Multiverse OS: Security focused real-time hypervisor
#### Designed for normal people (not corporations) to cluster all their different computers from embedded to personal into a secure, modern, super-computer

Epehemeral immutable virtualized Linux bringing bleeding edge virtualization technology (advertisers double-speak refers to as "Cloud") to the "end-user" or people. Utilizing features like nested virtual machines, open-source bootloaders like `uboot` to sign kernels from the BIOS up on every machine in the cluster from HOST to (App)lication VM, these features, that are almost exclusively only available to corporations to the every-day computing experience.

	
Fundamental components that make up the _three (3) primary_ virtualization layers of Multiverse OS, as opposed to most virtualization systems are two layers:

  (1) **Hypervisor** or **Virtual Machine Monitor (VMM)** or 


  (2) Minimal host software stack, with intentional disabling of functionality closest to hardware to prevent physical attacks (such as malicious attack in public space commonly referred to as "evil maid") and flawed/backcdoor and LAN hacking.

Under the security concept known as "security-in-depth" and engineering concepts derived from "clean-room-engineering" we assume every vector is exploitable. The design assumes the hardware is backdoored, and that every input is potentially an attack on the security of the system. Currently the project is pre-alpha, it is constantly evolving, and at the time of this update (Sept. 2018) it has been in active development for three years. 

Input into the machine from secure user is verified using the **Scramble (Identity) Suite** protocol, and machines the the cluster require any input to be signed, and all output including error output to be encrypted to the user defined by the **Scramble Suite**. The shell used by this user is referred to as the **Scramble Shell**. Using these concepts we are attempting to segregate the concept of a personal computer from the hardware, for the purpose of creating a super-cluster of computers that will be able to support a secure computing environment for any user with a **Scramble ID** generated from the **Scramble (Identity) Suite**.

Using the **Scramble ID**, a user interface (UI) with the personal Multiverse OS cluster, through the *User Controller VM*. The Multiverse OS user, **NEVER** interacts directly with the HOST machine, it is not only segregated from the user interface (UI), but it is intentionally for all practical purposes, impossible to see, use or interact with anything other than the *User Controller VM* and the VMs nested within. In addition, to the *User Controller VM* the host also has a *Router (Networking) Controller VM*, containing at lesat three (3) dedicated routers, to provide very complex networking outside of the control of the user of the *User Controller VM*. This creates a computing environment in which neither the HOST or the *User Controller VM* can directly access the internet. Instead the  to send a very heavily rate-limited set of limited commands only providing the most basic commands to dedicated (App)lication and Service VMs.

(TODO: Add citiatons, diagrmas, and get more edits from peers.)


_____
## Current list of Portals (VMs) Formulas to be released on Alpha Release
Below is the evolving list of Multiverse OS portals (VMs) that are planned to be released with the Alpha version fo Multiverse.

````
	portalgun
	├── application
	│   ├── browser
	│   │   ├── chromium
	│   │   ├── firefox
	│   │   └── torbrowser
	│   ├── console
	│   ├── filebrowser
	│   ├── media
	│   ├── project
	│   │   ├── audio
	│   │   ├── development
	│   │   │   ├── hardware
	│   │   │   │   └── go
	│   │   │   └── software
	│   │   │       ├── c
	│   │   │       ├── cpp
	│   │   │       ├── ruby
	│   │   │       └── rust
	│   │   └── visual
	│   └── steam
	├── controller
	│   ├── router
	│   │   └── formula
	│   └── user
	│       └── formula
	└── service
	    ├── client
	    │   └── torrent
	    ├── crypto
	    │   └── wallet
	    │       ├── bitcoin
	    │       └── ethereum-classic
	    ├── router
	    │   ├── galaxy
	    │   ├── star
	    │   └── universe
	    └── server
		└── web
````
____________________________________________________________________________
## Secure Open Hardware 

Multiverse OS developers are working closely with open source hardware engineers to develop keyboards, mice, USB disks, tablets, phones, and laptops that will work in tandem with Multiverse OS to provide better security than off the shelf proprietary hardware.

More developments as we come closer to revelaing *.STL files, get closer to manufacturing, and start doing runs of prototypes. 


____________________________________________________________________________
## Multiverse OS: A Real-time (Hypervisor) Operating System (RTOS) 
Multiverse OS will ideally be ran on HOST machines that are using chips like STM32 which support real-time operation, instead of the current tick system. This will improve our ability to use Multiverse OS on cutting edge open-source scientific equipment.

*TODO: Need to merge in the many other articles in our community resources about RTOS.*

#### RTOS & Real-time Hypervisor resources
riot-os.org

#### RTOS is critical for Multiverse OS
Becuase Multiverse OS is being designed for both end-user, and to improve citizen science by building cutting edge open source scientific equipment. RTOS is critical, powerful robotics, scientific equipment.

timesys OS - science ,mmeasuring, roboots
LabViewZ


#### RTOS as virtual machine
EmbedOS
Lynx OS


#### Patching linux for RT functionality
rt.wiki.kernel.org - rt patches for linux kernel;
linutronix


____________________________________________________________________________
## Relevant Independent, Peer-reviewed and News Articles 

"My VM is lighter, safer, faster, more preformant, can play new FPS, isolates my tasks. your container sucks" (find the link to this article)

**Resources**

https://github.com/cznic/ql


