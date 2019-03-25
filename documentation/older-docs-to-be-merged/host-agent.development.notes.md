# Host Daemon/Agent (Multiverse AGENT)

**rebuild routers using BIOS for complete boot with signed efi bios all the way to signed kernel**

HOST agent will be responsible for managing the two controller VMs (router controller VM, user controller VM). 

  The router controller VM may be Universe0 if possible, nesting inside galaxy0 and star0 (and maybe vventually planet0).

_______________

Agents already exist and come with libvirt, and use spice channel to communicate. 

__Example:__
	[qemu-ga] handles things like shutdown, reoboting vms, freezing,
	thawing filesystems (for live backups)

	[spice agent] cutting and pasting in and out of vm, smooth mouse
	movement ,etc

injects mouse position and state to the guest when using client mouse mode it neables to move freely between guest and client. aligning guest resolution when entering full screen.

**communciation is handled by a specialized virtual PCI device, called VDI. this is how messages are passed betewwen host and spice agent running on guest** this supports multple chanells [main, display, inputs, cursor, playback (audio received from the sever to be played the client, record (audio capture the guest)

spice protocl defines the messages

	[multiverse] what are we going to do?

*should we make a specialized graphics PCI device (think QXL for spice) that doesnt display but instead converts framebuffer data direclty to CBOR or other binary format, transfers via DMA to render in window via wayland*



___________________________________________________________________

___________________________________________________________________
## USE vPMU
This will be SUPER useful if running a hosting company or for MULTIVESRSE OS - letting others use your resources for crypto, or using others for crypto.
BUT ONLY AVAILABE ON INTEL
___________________________________________________________________



___________________________________________________________________
## dynamic expansion of fixed size HDs

[dynamic expansion of fixed size linux VHDs]
Because its very easy for us to snapshot while a machine is running, 
and adjust size (shirning and growing based on a velocity. we could
design a special Mutliverse OS HD style leveraging existing types with
the host agent making adjustments based on logged changes over time.
_this concept has been done by redhat too, its not impossible_







# =======================================================================
## Multiverse OS Agent will NOT REQUIRE proces on (APP/Service) VM

  * Instead Multiverse OS should hold open a channel that will allow 1-way communciation. 

# ========================================================================
## Multiverse OS HOST init (coming out of initramfs)

We will use kernel patches and modified init to PURPOSEFULLY cripple
the HOST machine, making it unable to control the processes of the controller. we can even START it BEFORE our init process, making it higher
authorization. while only accepting input that is signed by a special key
outside of a few commands. This prevents attacks in either direction.

	[CPU]
	DO NOT overcommit CPU! dont. not ever. especially not rpdouction

	[memory]
	we dont wnat host contorlling VM memory

 * dont overcommit memory in gneeral or ever. we want to preallocate where possible and enrypt. isnce our vms are mostly run within the controller. we just need to protect the controller primarily.


  __ksm__ shaed memroy using copy on write.
  **IS INSECURE, introduces side channels**
  

[ok: everything is insecure, what do we do?]
[answer: CGORUPS]

not just for containers. cgroups are "control groups" whjioch allow allocation of resources: CPU time, system memory, network bandwith.


  * cpu

  * memory

  * blkio (the most important bottle neck in modern computing)

  * device

[improve futher: HUGE PAGE TABLES]

transparent huge pages is a kernel feature that reduces TLB entries needed for an application by also allowing all free memory to be used as cache preformance is increased. 

to usetransparent huge pages, no special configuraiton qemu.conmfi si required. 

just set `/sys/klernel/mm/redhat_transparent_hugepages_enables` is set to "always"

*L3 cahce is better used with huge pages* huge page incraseing cpuy cache hits aginst the transaction looksasie buffer (TLB huge pages SIGNIFICANTLY incrases preforamnce particularly for large memroy and memory-ointexssive work laods.huge page use less memory 




___________________________________________________________________
## Don't overlog, dont log everything, think about what you are logging

  * Use syslog to push logs and journalctl to CONTROLLER VM, dont EVER use IO on any VM but the controller VM. 

  * Use tempfs or other Memfs (maybe with TTL) and feed logs to CONTROLLER VM. 


[controller vm log system]
feed all the logs to the controller VM, parse the logs, convert it to a consistent data type (ie json, xml, etc) and a database that is searchable. 

logs that are not searable are useless. use the logs to generate statistics
about users for exmaple on a web application (fuck giving google info) use
your logs dummy.

____________________________________________________________________
## security


  [kmod]
  * `allow_unsafe_assinged_interrupts 0`

  [smart]
  * disable smart, it should only be used on physical NOT virtual HDs

   `chkconfig --del smartd`

____________________________________________________________________
## PCI management with HOST AGENT

  * All PCI devices should be iterated and a XML file generated for each
device that canb e included into any machine. 
    * [KEY IDEA] we want to identify all the variables like attribute values (determine all possible options, use these to VALIDATE INPUT!!!)

    * if a XML is attached without a value defined, ask the user to fill in the value unless qieuyt and use default values.


  * additionally depending on the device, if it is a device that is passed
 through for example, it should track the kernel module, disable it, to
 free up the device. or renable it when the HOST needs it.
____________________________________________________________________
## PCI and `virsh`


  [gathering pci device info]
  `virsh noedev-dumpxml pci_0000_00_19_0`
  **this functionality should be added to protalgun but have better wyas of listing and pulling pci devices.**


  [attaching a single device by storing that device info in xml]
  `virsh attach-device guest-rehl6-64 xml-storing-pci-info.xml`
  
  [redefine pci device from virtual machine (maybe useful for controller)]
  `setsebool -P virt_use_sysfs 1`
   


# =======================================================================
## [HOST multiverse folder structure][the HOST agent should watch these folders for changes, create items in these folders and generally just be aware of this]


 **ROUTER NAMING** Need to switch back to universe0.router (inside we will have galaxy0 and star0). This will allow creation of universe1 and so on. as needed)

  `universe0.router.multiverse`



 **Create run folder to hold pid data, XML status, and our custom channels and status output files** later when we are building the agent daemon, we will create these files and put them in a memFS/tempfs




`/var/multiverse`:[will contain instance/install specific files, like the VM images for the specific install. these are unique to the install.]
	mkdir -p /var/multiverse

symbolic link /var/run/multiverse to run inside /var/multiverse

`cd /var/multiverse && ln -s /var/run/multiverse run`

`/var/run/multiverse`:[will contain instance/install specific files pertaining to this specific RUN, as in this PID, example: socket to connect to process.] **PID and UNIX SOCKETs should go here** *important its here because /var/run is delted every boot so no old pids are sitting around. Could symbolic link this to /var/multiverse/run for ease of use*

	mkdir -p /var/run/multiverse

# Create a folder for each VM (this should be done during VM creation in portalgun)

	root@host:/var/multiverse/run# mkdir universe.router
	root@host:/var/multiverse/run# mkdir galaxy.router
	root@host:/var/multiverse/run# mkdir star.router
	root@host:/var/multiverse/run# mkdir controller.router
	root@host:/var/multiverse/run# mkdir controller.user


`/etc/multiverse`:[install speciific configurations. these will be generated using template in `/usr/share/multiverse` then placed here to be edited by user/admin]

	mkdir -p /etc/multiverse


`/usr/share/multiverse`:[will contain default configs, templates]
_For example, `openssh` provides the default `sshd_config`_

	mkdir -p /usr/share/multiverse


## ======================================================================
# HOST AGENT should attach all `open as...` with special HOST AGENT generated file that just fires alarms, because nothing should EVER open on HOST. 
*(Multiverse OS doesnt use qemu-ga, its too insecure)* but it is a good reference for agents in general*

  * Modify kernel for controlled IN/OUT

## ======================================================================
## Getting rid of virt-manager by getting rid of the need for graphical CONSOLE in favor of using VirtIO console, and out-file for esy to read stauts.


_should this *host* folder be tempfs to delte on every boot? probs_
/var/run/multiverse/host/status
  ^ output time to login screen [i.e. 22 seconds to login screen]
  ^ output time to internet [i.e.e 30 seconds to obtain WAN IP: {show WAN ip}]
  ^ output dhcp lease, wan ip and other internet info needed to improve portalgun UI

  ^ put other things in here, like running processes (unique processes, rsources, etc)



____________________________________________________________________
## Timers and Clocks in QEMU

[Timer/Clock Reviews]
  * _pit_: programmable interval timer - a timer with periodic interrupts. [how is this used by athe vm?]= not usable with higher precision times.

  * _rtc_ - real-time-clock - continiously running timer with preiodic interrupts

  * _tsc_ - time stamp counter -0 counts number of ticks since reset, no interrupts. [whats this used for?]

  * _kvmclock_ - reocommend clock source for KVM guest, virtual machines KVM pvclock or kvmclock lets guest machine reads thost pshyyscial machines wall clock time. [!!]


[track atrbites]
  bool = unsupported or not

  guest = RTC always tracks guest machine time

  wall = RTC always truacks host time (best)

[tick polciy attributes]
  delay = continue to delvier normal rate (so ticks are delayed)

  catchup = deliver a hgier rate to catch up

  merge = merge into single tick
  discard = missed ticks are discarded (most likely best, like udp frames
  in a multiplayer FPS)

## ======================================================================
##
## HOST Configuration & Processes
##
## _BOTH configuration and processes that are unique to the VM, need to 
## be tracked for variety of reasons including preformance turning,
## security, automation, etc.
## ======================================================================

[CONFIGURATION] **Literally all the files that need to be edited in order to properly configure the Multiverse OS HOST environment.**

* ONLY have User account and locked down account for running user space unpriviledged service for Service VM or application for APP VM.

* Always have SELinux enabled: `setenforce 1`

* remove all uncessary services, check what ports are being listened on (AutoFS, NFS, FTP, HTTP, NIS, telnetd, sendmail and so on)


`/etc/libvirtd/libvirtd.conf`:
	L85: `unix_sock_group = "libvirt"`
	L108: `unix_sock_admin_perms = "0770"`
	L111: `unix_sock_dir = "/var/run/multiverse/"`

[PROCESSES] **Literally EVERY unique process run on the machine** (Also which ones were disabled, so we can track the "default" set and which ones are enabled, the resources used and so on.

	systemctl enable `libvirtd`
	systemctl enable `virtlogd`


## ======================================================================
##
## HOST Daemon/Agent Functional Requirements (Alpha) 
## ======================================================================
[0][Access qemu://session (unpriviledged libvirtd access) without SSH, using 9p over VirtIO]
This is still not ideal but a major step in security over providing full SSH access. Using group permisisons (using mapping or passthrough; whichever maintains permissions) assigned to the user, this could be limited exactly to the file that needs to be used to make the connection. 

Add the folder to the storage pool on host, then add the 9p share to the CONTROLLER VM. Boot the CONTROLLER VM and modify the `/etc/fstab`:

```
HostSockets /media/user/HostSockets 9p trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail 0 0 
```
Then add the following connection the `virt-manager`:

```
qemu+unix:///session?socket=/media/user/HostSockets/libvirt-sock
```

Failed to connect due to the following error: **"Verify the 'libvirtd' daemon is running."** [on HOST].

______
  [1][Automatically start the ROUTER VMs and then the default CONTROLLER VM, in the specified (default) order. Only boot the next machine after the current one is booted:


## How to determine if the VM is booted:
[First Attempt][qemu xml]
The first guess was the look in:

	`/run/user/1000/libvirt/qemu/run/universe.router.multiverse.xml`

````
<domstatus state='running' reason='booted' pid='2315'>
````

But this reports as booted INSTANTLY, as soon as the XML is created. It also does not indicate if the VM is online and receiving/sending packets.


[Second Attempt][Ping and other net tools]
Partial success, because this method does work for `universe.router.multiverse`:

````
$> ping 10.0.0.1
PING 10.0.0.1 (10.0.0.1) 56(84) bytes of data.
64 bytes from 10.0.0.1: icmp_seq=1 ttl=64 time=0.039 ms
````

But it will not work for `galaxy.router.multiverse` because it is specifically designed to segregate the network and not allow traffic to cross over this router unless its leaving for the VPN.


````
$> ping 10.1.1.1
PING 10.1.1.1 (10.1.1.1) 56(84) bytes of data.

--- 10.1.1.1 ping statistics ---
63 packets transmitted, 0 received, 100% packet loss, time 63384ms
````

[Third Attempt][VirtIO] 
There seems to be two obvious ways: 

	(1) Either have a 9p share over virtIO and write to file, indicating
	boot and internet is succssful that can be watched by agent, or 
	
	(2) even better, would be to just use the VirtIO serial connection
	and establish a tty to check. This only supports short polling so
	even better would be listneing on this, and waiting for each VM to
	send a singal/notifiaiton each event happens (but cant we just check
	the logs then?)






