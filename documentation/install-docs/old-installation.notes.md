## (Old) Design Road map
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

```
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
```

Then run the update command to rebuild the grub menu.

```
update-grub /etc/default/grub
```

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

