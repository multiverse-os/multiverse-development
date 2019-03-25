# NOTES


* The concept of putting the IP address in the mac is clever but eventually we will likely want to randomize it regularly to avoid Multiverse OS detection via mac addresses, especially since mac addresses are possible to detect over the internet. But in the meantime if we continue using the IP address in the mac scheme , we should definitely switch to full hex instead of just 254 being fe. 10 is not 10 in hex; it is 0a, and 100 is 64. So if we do this partial hex it will be very confusing once we have more networks. 

 So for example 10.100.100.1
 MAC would be  00:00:0a:64:64:01

---
* Use `lscpu` command and parse out all features supported for building a explicit as possible libvirt (and eventaully Multiverse specific) configuration. 
  * Multiverse OS should have reference lists, one should contain CPU features which have security problems. Then these features can be disabled on every machine they are found on automatically.

* **Disabling lahf_lm** was an old fix for a bug that does not exist anymore and should definitely not be done in any config. 

* **svm** might be worth looking into, it is an acronym for "secure virtual machine"; the specifics of what that actually meant would be useful for determining if this would enhance our security.
  * Enabling *EDB* (Execute Disable Bit; or arbitrary execution of non-executable code in memory to prevent some typse of exploits and malware. Additionally, other BIOS/CPU manufacturers call this feature: (1) XD (eXecute Disable), and (2) NX (No eXecute). 





---


* Should `images` and `portals` be moved to `/var/multiverse/images/*` for simplicity? Previously this folder contained both `os-images` which would contain all os images for setup of virtual machines and virtual machines disk images under `portals`. 

 Perhaps portal gun should be containing configurations, settings, provisioning, etc and the /var/multiverse/images/(portals|os-images) will contain the active/current images used by the machine in question.

* [suspension and power management] 

```
*ACPI sleep states*
S3 = suspend to memory 
S4 = suspend to disk 
```

## Updates to Libvirt, QEMU, and related software
##### [nemu]



##### [qemu]



##### [libvirt updated]
* SATA is supported now for CDROms, this is a massive improvement over IDE (or should be PATA actaully, wtf, seriously this shit is out of date). 


* Cache options are now supported in libvirt XML and do not need to be passed as kvm command-line options. 



##### [virt-manager updated]


## Preformance 

**[memory:huge pages]**

Enable in OS by preforming the following commands:
```
echo 2048 > /proc/sys/vm/nr_hugepages
mkdir -p /mnt/hugepages
mount -t hugetlbfs hugetlbfs /mnt/hugepages
mkdir -p /mnt/hugepages/libvirt/bin
systemctl restart libvirtd
```


## Recent Install Notes

* Retried installing recently after updates were submitted by wavelet, realized when setting up controller that I had accidently used `i440fx` instead of the preferred `q35`. This MUST be changed. We need to opt for XML generation by our software ASAP. 

**TODO** Write a go script to quickly modify the XML to change to Q35 AND remove addresses then load it into libvirt so that the addresses will be generated. (Eventually just generate addresses ourselves and track used. 

*The virDomainPCIAddressesSetGrow() function may be a good place to start in research on this subject.*


The `virtualization` package in fedora likely has a comparable package in debian. This apckage provides the command `virt-host-validate` which will check for haredware virtualization, /dev/kvm device, /de/net/tun, and LXC minimal kernel (2.6.26). A comparable piece of software for multiverse that could be leveraged by other software and used standalone to quickly check if the machine in question supports functionality required for Multiverse and running VMs. 

## CPU Configuration
**We want to write software to fill in the 'blanks' of our VMs and make them appear like bare-metal machines. From renaming devices so they use more realistic names, random generation of hardware UUIDs/hardware addresses/serial numbers (maybe simple YAML that has the regex or glob for generation for each device brand/device/etc.** This also potentially provides security benefits from viruses gathering details of hardware by hardware serial numbers and similar to verify a given server is a the one infected. Additionally, CPU features can be forced enabled even if they are not supported which could be used to disguise hardware further. 

KVM domains can use `host-passthrough` but name of CPU makes no difference for pure qemu VMs. 


*OPTIONS*
```
vmcoreinfo
hpt resizing=X
gic version=[0-9]
htm state=ON|OFF
pvspinlock state=ON|OFF
pae
acpi
apic
hap
privnet
pvspinlock state=ON|OFF
ioapic driver='qemu'
smm state=ON|OFF
  tseg 
hyperv
  related state=ON|OFF
  vapic state=ON|OFF
  spinlocks state=ON|OFF
  synic state=ON|OFF
  vpindex state=ON|OFF
  reset state=ON|OFF
  vendor_ID state=ON|OFF value="KVM Hv"
  frequences state=ON|OFF
  reenlightment state=ON|OFF
  tlbflush state=ON|OFF
  ipi state=ON|OFF
  evmcs state=ON|OFF
kvm
  hidden state=ON|OFF
```


* More recently it seems to be decided that host-passthrough is the preferred option for mode. Pinning for the controller should happen on the first chip to second to last (enough for routers). 


**Features**
Ideally we need a way to list all the features provided by a given CPU. Then we can make a list of disallowed for security reasons, and then disable all the CPU features with any potential CPU issue. Then we can basically use our CPUs with all the bad features subtracted for increased security. 

#### Caches (level 1,2,3)
Since 3.3.3+ we can finally define level 3 cache without kvm command-line options. 

`libvirt` now supports `<cache level='3' mode='passthrough' />` so we may not need to pass our kvm command-line options for level3 cache enabling.  Or you can emulate it or disable it




## Lines of Research


* [Disabling membaloon] 
By default libvirt keeps adding this back in and after much discussion between Multiverse OS developers and developers of other security enhanced linux, regardless of what is believed about memory balooning, we do not want this feature and want it to not only be disabled by default but to not be an option at all. 

   *From a help ticket in 2010, obivously out of date, it required modifying qemu source.*

---
* WE have NOT been setting up our VMs correctly, Q35 IS uses PCIE by default. And will make passing data between host and guest via PCIE bus much easier. This has been an important line of research when looking at the best way to move Xserver client data to the server quickly, preferably via some sort of direct memory-access (DMA). 

For example, a controller with the model `<controller type='pci' model='dmi-to-pci-bridge' />`.  This may provide the foundation for for custom PCI software. This is not DMA but it will provide a starting point in searching through the codebase for relevant examples for our software. *What this does provide is a way to add more PCI bridges or 'hubs' to increase the total number of avaialble PCI slots.*

We have been using the default for windows machines: `i440fx` which by definition is a legacy machine, **this should not even be an option in Multiverse.** This is another reason to drop support for libvirt and implement our own KVM/Qemu abstraction layer that removes legacy, deprecated, windows-specific, and insecure options altogether. This would make our systems more secure, faster, and make understanding the VM system under Multiverse OS significantly easier. 

 


---
* The `<on_*>`; for example: `<on_crash>`, `<on_reboot>`, `<on_poweroff>`... A new option has been added: `<on_lockfailure>` so one can increase the security by doing `<on_lockfailure>poweroff</on_lockfailure>`. One option is `preserve` which allows it to be saved for later deep analysis. 

For Multiverse this functionality would be great, furthermore, if we could hook onto a variety of other functionality, like NIC unplug or NIC disabled, or HD full, etc. And then instead of a simple set of fixed options, if we co-uld run any command or script. 


--- 



* Look into `<seclabel>` element again, its either deprecated or fairly new, found it in an old XML file when reviewing XML to outline a Go script to manage XML files.

 * Look into controlling thread allocation for routers, controllers and app VMs (make it dynamic preferably and easily changable from our UI). Uing the `<iothreads>` element that is put under `<domain>`. 1.2.8+. *You can even set the thread id for maximum control and monitoring*

* shares under cpu configuration is set by the OS if not defined. So we should definitely be defining this so we are not handing off control of our VMs over to OS defaults. 

## Nested Virtualization
The Multiverse OS design now has two operating modes:

  1) The original, and preferred mode due to maximum isolation, modularity and enhanced security utilizing nested virtualization for application VMs running inside the controller VM. 

  2) The new "lighter weight" operating mode combines the responsibility of the controller VM, and functions more like Qubes OS. The host+controller responsibilities remain on the host and application VMs are run direclty on the host instead of running nested inside the controller VM. This is being designed specifically for running Multiverse on a wider variety of devices, intended for use with machines with less CPUs but still providing all the security enhanced features and functionlaity that comes with Multiverse OS by subtracting features that are either not available or not supported by older or more affordable hardware. 

Currently the alpha-installer is being built to support automated installation for both operating modes, in addition simple standalone software will be provided in conjuction with the installer to easily and automatically determine the best operating mode for a machine during installation. 





## Security


* We need to check firmware on memory, cpu, keyboard etc, wherever we possibly can and take checksums. This will allow us to monitor for unexepeected changes from the most nefarious and dangerous viruses. 

* [memory options in xml] **MUST** use `<nosharepages>`  to disable shared pages. Each VM should ahve their own memoory allocations, not use memory balooning and have it encrypted to a key only the VM itself knows.  In addition, `<locked>` locks the host memory and host will NOT be allowed to swap out memory.

  **This is required from some real-time software and likely should be done for at least our controller** For services offering computer resources to OTHER suers can leverage this for denial-of-service (DOS) attacks. And so the `<hard_limit>` instead because it allows some of the memory to be locked. Hard limit is the maximum memo-ry the guest can use. If the limit is set too low the guest can bee killed by the kernel. 

  **IF you are optimizing via locking the memory ensure that there is a HARD LIMIT to prevent the guest from locking the entirity of the host memory as a DOS.** Additionally a `<min_guarantee>` option can be used to gaurantee minimum memory allocation. **It is important to acknowledge we dont want memory ballooning for a variety of reasons, we want fixed, static, encrypted, private and LOCKEd memory allocation ideally. At the very least we want that for the controller.** 

```
<hard_limit unit='G'>1</hard_limit>
<swap_hard_limit unit='G'>2</swap_hard_limit>
```

  `<source>` allows a specific file to be used, this might be nice to do because we can make a software that provides a "soft" file  to allow us to control the system at a much deeper level. 

  `<access>` should be set to private

  `<allocation>` will be ondemand by default but likely should be immediate. ESPCIALLY on the host, we run a fixed number of VMs, it should cut up the memory and allocate it and that should ideally not change. 



## Decision Reasoning And Related Notes

* Q35 is suggested for Linux VMs vs i440fx which is the default and suggested option for Windows VMs. Q35 (or sometimes referred to as emulation of ICH9 host chipset) supports PCIE bus better and therefore will work better with PCI-passthrough. For more information, the Intel ICH9 diagrams provide a great introduction to what is being emulated. *alternatives* that provide preformance boost can be found in the nemu codebase. Q35 does not have support for legacy guests (ie XP/2000) and limited support for legacy devices. 

**Mistakenly used i440fx?** Switch it using `virsh edit` to switch under `<os>` element `<type>` attribute `machine=` from `pc-i440fx-2.\*` to `pc-q35`. After making this change, it is important to delete the `<address>` element from each PCI device, deleting this will allow libvirt to automatically add an address using the pool of available options, and since changing the chipset will change the available addresses, deleting them allows each one to be reassigned based on the new address assignment. 

## Library/Dependency Updates

* Since libvirt 2.2\* virtual networking supports "open" in addition to the original "NAT" and "routed" options when configuring a virtual bridge (something we would like to avoid entirely soon as possible). The new open mode connects the guest to the virtual bridge WITHOUT modifying iptables on the host machine. No modifications are made to the host firewall, in contrast, both "NAT" and "Routed" options both modify the firewall, specifically by adding entries to the iptables of the host machine. 
