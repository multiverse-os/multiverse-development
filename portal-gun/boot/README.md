# Multiverse OS: Secure and trustless boot system
The Multiverse OS system is designed around the concept of automating the local builds of self-signed open-soruce BIOS, bootloaders, and additionally cryptographically secured, functionality limited initramfs that track intergrity, changes, and require authorization to update.

*There are still a lot of work to do to improve the security of the boot process of the Host machine, associated devices, virtual machines (controller, service, and application machines).*


### BIOS
Currently we are utilizing OVMF (sometimes referred to as tiano core), an open soruce BIOS that supports secure booting. But we have yet to leverage this functionality. 


### Bootloader
Multiverse OS will support and simplify building the `uboot` bootloader and simplify deployment to cluster connected mobile devices, other hardware connected to the LAN, and virtual machines where use of `uboot` makes sense. 
 

### Initramfs
Currently this is one of the best Linux attack vectors for persistence and maintaing a virus long term. All of these issues must be addressed, using a variety of techniques to ensure every time the initramfs is updated, it is done so in such a way that is cryptographically secure, provides bare minimum of functionality needed by the machine using it, provides file intergrity checking, logging of jupdates to various initramfs systems, tracking all the various versions needed for the different virtual machines, host and other devices (for example, mobile devices)

===============================================================================
## Multiverse OS Secure Boot Software Suite
Below are will be the beginning of breaking up the functional requirements of
the Multiverse OS cluster secure boot system into a collection of individual
pieces of softare to begin the next phase of design and prepare for the early
development and implementation of pre-alpha versions. 




