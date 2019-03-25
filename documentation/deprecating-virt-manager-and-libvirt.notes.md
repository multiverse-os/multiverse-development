# Deprecating `virt-manager`
The most important and immediate goal for Multiverse OS development is
getting rid of `virt-manager` soon as possible. It is full of old features 
(i.e. parallel ports) and it does not have nearly enough of the actual features,
many of which are VERY important for running VMs efficiently, correctly, and
securely. 

**What needs to be implemented to abandon `virt-manager?**
There are a few key issues that need to be addressed for us to abandon the use 
of `virt-manager` entirely.

  (1) **Access to VMs console without internet and early enough to initiate**
      **installation** without `spice viewer` (or anything relying on internet,
      like ssh). We need to be able to do maintaince when internet fails, so we
      need access from the start of boot or at the very least at bootloader
      (grub2). 
  
        We can acheive this goal a few ways:

     	(A) Direct kernel boot allows addition of <cmdline> attribute in XML 
      	    under the <os> attribute. In this we can add flags to the kernel
 	    and this is exactly like modifying the `CMDLINE` in grub, (where
 	    `quiet` is by default in grub2). 

	    Specifically, `console=ttyS0` and adding a `Serial` device with 
   	    `pty` set.

    	(B)


  (2) **Auto-start VMs in order, waiting for boot and internet before starting**
      **next VM in cluster boot list**

	(A)

	(B)


  (3) **CLI Wizard for configuring VM XML for libvirt** (or) preparing a 
      `qemu-system-X` command line (bypassing libvirt XML).

	(A) Look at survey or some other CLI wizard.

	(B)


  (4) **Simple VM-to-VM Communication**

	(A)

	(B)

	(C)


____________
## Next Steps
The next steps after this would be provisioning VMs automatically using the special QCOW2s that hold all the things portal-gun will need to turn a VM into a specific portal *(like `firefox.app.portal`)*.




