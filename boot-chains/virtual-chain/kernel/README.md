# Kernel










The final component of the boot process, the production ready version of Multiverse should not support DKMS, it should have all the kernel modules needed for the system built into it and changes to the system require signging the changes with an authorized key and rebuilding the system for enhanced security (preventing all kernel module based root kits, which is not as popular currently as the initramfs but community "farmed" viruses via honeypots will provide greater knowledge about the current popular models of attack (if we are able to expect the vector at least to a degree) so we can focus development on defense of different aspects of the system based on the popularity of attacks. 

For example, attacking the initramfs has been the most popular way to attack unix systems since it gets you in before the decryption process of the hard disks, it is hard to detect if done properly, and the system around updating the initramfs and ensuring it is secure, has not been tampered with, and so on is simply non-existant for modern mainstream desktop environments, and is only available on enterprise type servers using their own BIOS implementations or cutting edge open source BIOS options. 


