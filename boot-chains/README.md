# Multiverse OS Boot Chains (Phys/Virt)
A typical chain will look something like:

`BIOS -> Bootloader -> Initrd -> Kernel`









## Project organization
It may prove to make more sense to find a different way to package the boot chains physical/virtual.



















### What is initrd?
`initrd` is a scheme for loadaing a temporary root file system into memory, which may be used as part of the Linux startup process...[essentially at this point you have selected your kernel to load in your bootloader (grub2 commonly), and if you are using LVM and Luks which is now the normative workflow, you load ramfs based system to unlock your LVM Luks partition and complete the process of booting into your selected kernel.]

**Why is this important?**

Recently when developing Multiverse OS, I discovered a rootkit on the firmwarea of my keyboard which persisted in my initramfs. This strategy is incredibly popular for legitmately threatening root kits acheiving advanced persistence. An example has been widely availbale for a while now called `horsepill`, which shows an unsophisticated by incredibly effective way to persist on a machine due to fundamental issues with the way initramfs, and the entire boot process is handled. 

And so the urgency of implementing the solution as planned in Multiverse OS design documentation has become increasingly important. And this is the development package used to complete two standard boot chains for physical machines and virtual machines. 


**Strategies to bypass attack**

The host machine should **NEVER** have the ability to generate a legitimate network device and attach it. By the time you are booted into the final kernel, all the drivers should be blacklisted. The iptables should prevent it, and we should write further software and use other configurations to disable any ability to connect to the internet. 
