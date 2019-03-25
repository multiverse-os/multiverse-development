# Linux Bootloaders and us (multiverse os developers)

[Resources]
wiki.xenproject.org.wiki/hvmloader
hypervisor loader

____________________
## Tricking everyone into thinking we are not VMs (we totally are!)

qemu is relevant since it provies the ability to run unmodifeid guests, by fooling the qemu guest to think its on a physical machine, so qemu is in charge of:

  * physical to psudeo physical address transaction
  * emulation device communction with host hard
  * syscall to hypercall translation
  * more

________________
## Linux Kernel is a UEFI binary

  * there is a sys call that the linux kernel tells everything to kill itself so it can take over.

## uboot

  * uboot has rtc support

  * graphical boot menu with logos

  * bootmenu 
________________
## Removable Boot
The purpose is to not know what is bing booted, just boot off a drive
_______________
## Programming Multiverse Developer OS Guidelines

[Getting started]
__I dont care if google does it this way or that way, but we will not be using global variables, we will not be inconsistent with our naming, we will be consistent above all else. We will not have half our code be comments for generating docs that are less useful than reading the code (which would be less itimidating without all the useless comments)__
______


Code should explain itself, dont use comments if you dont need them. Use restraint.

*Column 80* (based on terminals), people read even less (keep this in mind when designing).






*style, convention, etc needs to be scalable acorss a team of d2d (developer-to-developer) decentralized development.*


[__AVOID THESE__]

 * no global vairalbes, no singletons

 * no recnetly used list

 * Avoid Getters and Setters, not great for non-programmers. get like set means too many things. we wnat to use defintions that mean ideally 1 meaning. they are not even opposites, unset, or reset is opposite of set for example NOT get.

 * lego naming: linguistic morphology of identifier names, DONT USE CLASS, EXCEPTION, ERROR, etc. in names wherever possible. Use actual names not descriptions, be concrete, omit needlessly words. use good metaphors.

dont OVER ABSTRACT, and dont UNDER ABSTRACT

_______________
## XML defining Bootloaders in Libvirt (we are moving away from libvirt asap)

# Type: element specifies type OPERATING SYSTEM (OS)
#  hvm: indicates os is designed to run on bare metal so requires full virtualization (debian, fedora, etc)
#  linux(bad name): refers to an OS that supports Xen 3 hypervisor geust ABI (dumb, we are not all use xen). 
#  arch specifying CPU arch to virtualization machine to the machine type. the capabilties XML provides alloed values (we should just use this then to VALIDATE lol) or just not use libvirt.
<type>hvm</type>
# Loader: refers to the FIRMWARE BLOB. master NVMRAM store file defined in qemu.conf (where?) template atrbiutes can be used per domain overaideing map of master. NVRAM stores the config file. ntrasnient domains if the NVRAM file has been created by libvbit is left behind and responbility lies on applicationss to use and sav.
# TYPE: allows "type" pr "pflash" 
#  pflash = UEFI image
<loader readonly='yes' secure='no' type='rom'>
/usr/lib/xen/boot/hvmloader
</loader>
<nvram template='/usr/share/OVMF/OVMF_VARS.fd'>
/var/lib/libvirtnvram/guest_VARS.fd
</nvram>
# This is neat, we can say NEVER allow USB to boot. AHaha fuck you evil --m-a-i-d--, no scratch that fuck you evil undercover fascists.
<boot dev='hd'/>
<boot dev='cdrom'/>
<bootmenu enable='yes' timeout="3000" />
<smbios mode='sysinfo />
<bios useserial="yes" rebootTimeout="0"
</os>

_______________
**Keep in mind that direct kernel booting would include uboot(bootloaders in general) since they use a kernel dummy**

[Grub2]



[Uboot][Supports widest variety of things]
Best for offense and defense because of this!

  **simple example from source**
  sudo apt-get install qemu 

  make versatilepb_config ARCH=arm CROSS_COMPILE=arm-none-eabi-

  ^ Creates the config, so lets now compile!

  make all ARCH=arm CROSS_COMPILE=arm-none-eabi-

  qemu-system -M versatilepb -m 128M -nographic -kernel u-boot.bin

  which outputs

   DRAM: 0kb
   Flash: 0kb
   ---
    In:Serial
    Out:Serial
    Err:Serial
    $~>

  **UBOOT** is super cool and can load itself from a SPECIFIC address if needed.





[Syslinux]
