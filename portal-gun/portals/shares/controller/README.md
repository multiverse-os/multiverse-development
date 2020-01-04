# Multiverse OS: Controllers
The Multiverse OS controller portals (VMs) are specialized hypervisor VMs
providing all their functionality as nested VMs isolated using full-virtualized
and para-virtualized machines, with further layers of isolation using VMs and
containers. 

The current design has two (2) types of hypervisor VMs running on the Host:


  (1) **Network (Router) Controller**

  (2) **(User) Interface Controller**


*Multiple of each type can run on a host. Eventually these may be further
nested into a single hypervisor VM in order to contain the entire logic in a
single package.*

_______________________________________________________________________________
## General Controller Configuration
This is the configuration that is necessary across all types of controllers, 
both networking and user interface.

For mounting a 9p folder without erroring at the startup, modify 
`/etc/initramfs-tools/modules` and add the following modules:


```
9p
9pnet
9pnet_virtio
```

This will allow unpriviledged mounting of p9 share folders.

**TODO: Put in XML for clock, cpu, memory, etc**

_______________________________________________________________________________


# Multiverse OS: (User) Interface Controllers
Eventually, Multiverse OS Interface controllers should contain each variation
within the controller interface, that way we do not need 5 differnet Ubuntu17
servers for example, instead we just need 1 and have 5 different profiles that
can be launched. 

# Multiverse OS Default HOST Controllers



### Alpine based controller

## Setup the 

libvirt
libvirt-daemon
qemu
qemu-system-x86_64

qemu-img
#dev
vim 

rc-update add libvirtd
______
### XML 


**Attribute: <seclabel>**
[resource](red_hat_enterprise_6_20.20)


# USER Controller VM - Agent/Daemon Service
### User Interface (UI) to Multiverse OS
The following is the roadmap to design and develop the USER CONTROLLER VM
agent/daemon process that will always run and manage backgorund tasks to
facilitate the ui to Multiverse OS.


# Controller Agent 


 [1][The agent should guarantee settings are correct and if they are modified to settings that would cause failure, the setting changed should be auto revereted to the known defaults]
  [Controller Default Settings]:
    * `user` should be in groups: [`kvm`, `libvirt`, `libvirt-dbus`]

````
      sudo usermod -a -G libvirt user
      sudo usermod -a -G kvm user
````

### Notes/Research/Scratch

  [*][Get rid of virt-manager] Its super easy to add images for exampe:
     `virsh attach-disk Guest1 /var/lib/vlibvirt/imgaes/FileName.img sbd`

### Entropy / RNG

  * rngd -b -r /dev/hwrng/ -o /dev/random/


### Functional Requirements

### Key Libraries

  * [Packages within the libguestfs suite]
    [url][libguestfs.org = has a list of all sub-projects/modules that make up the libguestfs package suite. has things like extract windows registry, convert guest to kvm (v2v somehow), format, get kernel from disk,  etc] set of tools for afcessing modifying vm disk images (could be the basis for auto shrinking and growing VM hds! obivously auto back up too silly)
        [*] monitor disk usage
        [*] creating guests
        [*] **SCRIPTING CHANGES TO VMs** (this is what we would use to fucking auto shrink and grow!!!!)
     * backups
     * formating
     *reziing (!!)
     * Cloning
     * Building
     * V2V?
     * P2V?
     * __editing files inside Guests and viewing__


    [SUB PACKAGES] the ones that jumped out to me becuse we wnated to do things like save a version (auto vM image versioning/snapshots, KEEP IN MIND using a system where the user files are segregated from the standard OS files


### Roadmap


