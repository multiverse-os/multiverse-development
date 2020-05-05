# Active Development Notes
Alpha release of Multiverse OS is being completed and the design/API is being finalized because it will be somewhat frozen, with beta mostly focusing on portalgun, desktop shell, standardization of the UI framework used by core applications, bash shell and shim, and similar higher level issues before the first release compiles everything and freezes the API. The goal is avoid major changes to the API after the initial release so that software written after this release will remain useable across more than a single major version, but preferably across ~5 major verisons. The plan to acheive this involves flexible API design providing an API that will enable us to do two dramatically different cluster designs and interconnectivity between these different designs. This should accomplish our goal and allow us to focus on details of the cluster design and virtual machine. 

Since applications will depend on the stability of the VM API for example, it si very important to get this right otherwise we risk either keeping a lot of logic around from earlier versions or breaking large portions of the application implementations. 


#### Portalgun API
This is likely the most important API to ensure major changes are avoided due to the application infrastructure in Multiverse will be built ontop of "minivm" style application VMs to encapsulate and and compartmentalize with full virtualization for maximum security (or hardware compartmentalization depending on hardware). 

The Multiverse OS system currently provides a ruby interface to provision and spin up application VMs that will feel familiar to vagrant users but a rich toolkit built into the configurations which enables actions that would be difficult and complex in vagrant natural and simple in Multiverse OS's portalgun system. This is done by exposing functionality through a module automatically imported into the configuration system, that cna be used to define a single application VM or a complex network of interconnected VM cluster stored nested within a single encapsulation VM for easy migration, management, and compartmentalized but cooperative development. 



## Development Tools
Below are tools that either we have developed or found during our research that helps develop Multiverse OS:

```
virt-host-validate
```

#### Qemu Image Tool
This tool will either be embedded if a tool is not built providing near full feature complete rebuild using our underlying framework. 

```
qemu-img
```




###############################################################################
# Add virtual PCI subsystem & slots
###############################################################################
# This concept needs to be translated to portalgun, this is the
# device on the host machine:
#
# <device>
#   <name>pci_0000_00_00_0</name>
#
#   *(This is one the virtual functions specified by path)*
#   <path>/sys/devices/pci0000:00/0000:00:01.0/0000:00:00.0</path>
#
#   <parent>computer</parent>
#   <driver>
#     <name>e1000e</name>
#   </driver>
#   <capability type='pci'>
#     <domain>0</domain>
#     <bus>0</bus>
#     <slot>25</slot>
#     <function>0</function>
#     <product id='0x1502'>82579LM Gigabit Network Connection</product>
#     <vendor id='0x8086'>Intel Corporation</vendor>
#     <iommuGroup number='7'>
#       <address domain='0x0000' bus='0x00' slot='0x00' function='0x0' />
#     </iommuGroup>
#   </capability>
# </device>
#
#
# Translates to a PCI passthrough:
#
# <hostdev>
#   <source> 
#     <address domain='0' bus='0' slot='25' function='0' />
#   </source>
#   <rom bar='off' />  (lspci -v command to check the device for expansion rom)
#   (the rom can sometimes be bypassed for speed, or other issues related can 
#   prevent booting)
# </hostdev>
#
###############################################################################
# **Virtual function without the device**
# SR-IOV PCI Device, Virtual Functionality 
# Virtual functions provide data protection between virtual machines and 
# the physical server as data managed is controlled within the hardware. 
#
#
# <interface type="hostdev">
