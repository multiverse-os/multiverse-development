# Multiverse Virtual PCI Wayland Videocard
**Review spice-space.org/spice-user-manual.html for more details and examples**

## vhost_net experimental_zcopytx=1
modprobe -r vhost_net

_Multiverse OS portal gun video will be both the virtual pci device and the device driver_



  [resources]:
    [syntehtci frame buffer driver] provides enhanced graphics preformance and superuiror for redhad desktop users. this could be a staritng point to fork for our driver.

<video {here our drive lets us put in custom flags to set config}><video>



The idea is using a customized Multiverse OS virtaul PCI video-card/gpu, we exract the framebuffer of a single window (preferably nothing else), without rendering at all on the guest machine, we take this data and conver to a binary format like CBOR. 

  

> *checkout the PCI VirtIO driver, may be a good basis for our Video portal video GPU*

Then transfer it DIRECTLY (VM-to-VM) to the Multiverse OS (bypass the HOST entirely using a virtIO channel, or userspace network stack. We will then use custom channel via /dev/SHM or /dev/DMA (direct memory access), ideally DMA. 

**we pull DIRECLTY from the memory of /dev/card0 and make it binary or just leave it alone, convert it to wayland protocl message. those messages are fed into a window on the controller.**

_[FOR DMA TO WORK]_ [we must have pre-allocated memory] which is critical for security anyways. SO PORTAL GUN SHOULD PRE_ALLOCATE MEMROY, dont over commit, divide by usage (use over time stats to adjust over time too, lets make this shti smart). 

This data is converted to wayland protocl, and fed directly into a pure wayland window for rendering. 




*cahce image data, prevent client sending same data using pixmap/pallete caching.*

This will enable us to maximize the preformance of window transfer, and will enable us to get FULL virtaulizlation isolation, and security (even enhanced beyond many other implementations like Qubes: using unprivildged VMs, not using host kernel for passing packets, using passwords, having the HOST be completely disabled, and never used by utilizing a complex system of nested VMs).

___________________________________________________________________
## You can unlock a used Grpahics card using "suspend to memory" trick

 * certain cards can not easily be returend to be reassigned. restarts 
   work, but stopping then starting or starting differnet controller
   will fial with black screen - unable to assing.

   rebooting resets it. there are tricks to reset it WITHOUT rebooting
    w
   **THIS MUST BE condensed into a function and added to HOST agent**
   **so we can freely move around our TWO grpahics cards!**

      (YES) enables BIOS support 

      (NO) disables BIOS supprot


___________________________________________________________________
## Machine POWER MANAGEMENT for VM


[ACPI]
  s3 = suspend to disk

  s4 = suspend to memory 

````
<pm>
  <suspend-to-disk enabled='no' />
  <suspend-to-mem enabled='yes' />
</pm>
````



===================================================================
___________________________________________________________________
## Support Screen-in-Screen (picture-in-picture) and grabbing screens from tablet, phone, etc
___________________________________________________________________










===================================================================
___________________________________________________________________
## Multiverse OS virtual networking PCI device
Utilziing similar concepts we can have networking PCI devices that handle the network stack direclty, isolated from both the HOST kernel but also the application/service VM kernel. 

Encrypt every packet, expect signatures during handshakes, to prevent spoofing, internal network denial of service and unprivilegded network exploration.

This will also help support resources sharing, making it safer to run your resources securely on a strangers comptuer. The real cloud internet.


______________________________________
## Multiverse OS virtual PCI card for video streaming over tor, to have streamers watching anonymous computer sessions. 

Mixed with COOP mode, utilizng lbinput multiple seats, and multiple screens, people can collectively use a computer anonymously and stream the action to adoring fans.
