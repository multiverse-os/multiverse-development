# PCI Devices and PCI Passthrough
While development continues on Multiverse OS PCI device management tools, the interim solution will be to break up known devices by components: ["GPU", "CPU", "Motherboard", "Expansion PCI Cards", "Other Peripherals"].

If your device is not included, you can help by looking at similar files and building one for your motherboard or other component.

Then early users (pre-alpha) can concat these device files together based on their own devices to quickly or at least guide implementation of scripts to unbind devices that are meant to be bound to Multiverse OS VMs. 


### PCI Device Manager
Below are a collection of notes for the design of the Multiverse OS PCI Device manager based on the current implementation and general design ideas relating to both PCI device management, and most importantly automating the process of binding/unbinding of devices to quickly connect and disconnect devices from Multiverse OS Controller VMs, and the Host Machine as Controller VMs start, shutdown and reboot. Typically only a single Controller VM is active at any given time, but switching between different categories of Controller VMs is common under Multiverse OS, and so the PCI devices must be quickly and easily passed between controller VMs intuitively without user interaction.


#### Design Summary
Until Go software can be completed that will manage and simplify PCI passthrough, manual passthrough
using per-host bash scripts that prepare the machine for launching the Controller VM.

These bash scripts will help provide the outline or blueprint for desinging the Multiverse OS PCI
passthrough. Current designs are more complicated than just a list of avaialble PCI devices and 
appending them to a list inside the VM configuration. Instead opting for a more fluid and intuitive 
approach that categorizes PCI devices, watches for changes, binds and unbinds on-the-fly, and applying
devices to VMs based on category and VM type. For example, all Network Cards are found and by default
passed to the Universe Router VM. Customization will be possible, but Multiverse OS should be anticipating
and dealing with PCI devices. 

Its also not as simple as just applying all USB devices to the Controller VM, the motherboard needs to be
mapped to make intuitive decisions that make sense. For Example, if the motherboard does not provide
PS2 keyboard/mouse support, then at least ONE (1) USB controller MUST be left to the host machine or 
at least provide automatic rebinding of USB controller PCI devices to the host machines whenever there is
no controller VM detected to be running or starting.

So while these scripts do not provide the literal blueprint, they elucidate the best ways to approach 
development of the eventual device manager. 

#### Categorizing Available Devices
Currently basic shell scripts provided with Multiverse OS crawl the contents of `lspci -nn` in order to categorize each relevant PCI device that will be made accessible to Multiverse OS VMs through PCI passthrough. 

Currently, the user must specify which devices to use with PCI passthrough in a script (currently `/etc/rc.local`). Shell scripts have been built to simplify the delcaration needed to make PCI devices assignable. 

An example `rc.local` using the Multiverse OS shell framework module for simplifying PCI passthrough

```
#!/bin/sh

# IMPORT 'vfio-management' component of Multiverse Shell Framework
. /home/user/multiverse-os/scripts/sh-framework/modules/multiverse/vfio-management.sh

#==(PCI Device Passthrough)==#
# Network PCI Devices
passthrough net 8086:15b8
passthrough net 8086:1539
passthrough net 1ae9:0310
passthrough net 168c:003e


exit 0

``` 

However the user should not be required to track down the available PCI devices for passthrough and manually initialize them in a script. 

Instead of relying on user configuration Multiverse OS should be iterating through all available devices on the Host Machine and categorizing them. The categories will indicate which Multiverse OS virtual machine types:

```
#    ["gpu", "usb", "other"] => Controller VMs
#    ["net"]                 => Universe Router VM
```

For example, the script should iterate through each PCI device, categorize, assign each (updating the XML if necessary). Keep in mind, this would still need to run every time the Host Machine is turned on.

#### Config Space Permissions
**NOTE** Because of changes in Debian Buster (likely addition of apparmor by default), config space file is not accessible to unpriviledged users and disables essentially all PCI passthrough without `chown`ing the config space file to root:libvirt, and changing the permissions with `chmod` to `0660`. Until the underlying cause can be determined this should only be done AT BOOT OF VM, and immediately changed back to root:root for security purposes.


#### Passing USB PCI Devices
**NOTE** If no PS2 device is detected on the motherboard (sometimes called Legacy Bus or LSB), and until a daemon can be created to manage PCI device passthrough on-the-fly, a single USB PCI controller should be left to the Host for debugging and development purposes. Without doing this, the user can be stuck without any easy ability to interact with the Host Machine without doing some sort of rescure procedure which may be too difficult for novice users. 


#### vfio-bind executable script
Multiverse OS tools providing both library support `scripts/sh-framework/modules/multiverse`, and separate executable script in the form of `vfio-bind`, provide different paths to simplifying the process of making PCI devices accessible for PCI passthrough and extending this existing short-term solutions when development begins in higher level languages. 

Currently both are supplied in the form of shell scripts, a `/bin/sh` framework to simplify improvements to the proto-alpha state of the installation process and transparent executable shell script that does not require trust of a binary compiled by an unknown third party. 

### Multiple Network Cards
Currently all provisioning scripts for the Universe Router VM assume you are only passing over a single network device. Changes need to be made so that it works with whatever number of cards is being supplied.

Eventually this can be the foundation of combining multiple ISP connections into a single connection for the routers connected to the Unvierse router. 
______

### Design Notes
PCI device management tools will need to be the first tools ported from the basic shell scripts used to manually install, prototype and design Multiverse OS.

Tools at this layer of Multiverse OS will be built with Go language and configured using Ruby as a scripting language. Anything further down, such as kernel modules will be written in Rust.


#### PCI Device Manager
In addition to managing passthrough of all the PCI devices, virtual network support (until it can be phased out with a custom network stack), storage, and should be handled by a single tool, and this tool will be a component of `portal gun`, Multiverse OS virtual machine management tool.


##### Resources
included are documentation that was specifically useful because of the depth of coverage on topics such as `vfio`. Often too many guides provide only information of applied knowledge regarding this subject, and often even that information is confused, often configurations have vestigal configuration options that no longer do anything. Below is some documentation that is beyond just the applied application of vfio:

[Introduction to 'vfio'](https://insujang.github.io/2017-04-27/introduction-to-vfio) This further links to important articles from linux foundation on this topic, which letsthe reader get deeper into the vfio topic.

[VFIO - Virtual Fabric IO a.k.a. PCI pass-through](https://youtube.com/watch?v=7d6gau1jea0) A video good introduction into vfio concepts, terms and examples. 

[2014 - VFIO, OVMF, GPU, and You by Alex Williamson](https://youtube.com/watch?v=NhZ9elpg2nM) Video by a VFIO developer directly overviewing state of GPU passthrough in 2014, applied but comphrensive for the time period. Good state of GPU passthrough for that time period. 

##### Writing device drivers in user space
Then talk to the device over memory mapped IO. Dont have to use interrupts, can do polling. Everything is mapped and handled by hardware.

This is done for lower latency in several applications, but importanlty for Multiverse OS is science applications and scientific equipment.

If user space drivers can be built this way, can even do things like network cards interacting with GPU to do very fast packet analysis and routing. 

**Interrupts in userspace is done with eventfd**, these can also be used for inter process communcation. 

VFIO device has IRQs which can be programmed against. 

The way to improve Xpra is using DMA directly P2P between devices (VMS, from application VM to Controller VM)
