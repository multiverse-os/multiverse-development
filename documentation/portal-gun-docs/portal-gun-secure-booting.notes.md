# Portal Gun: Secure Boot Process


## We can control GRUB level actions with sendkey.mod and serial.mod

This combined with chainboot, will let us provision from the lowest level! Using just grub


  *We want to programmtatically (USING RUBY + GO) configure our 
 
  BIOS + GRUB + Initramfs + Kernel (patches) + post installation (packages services, packet handling, networking, etc)

 **we want to program super low level OS for specific applicaitons or services** we dont then use the computer, we instead use the output WebUI, or regular UI over VGA or HDMI, etc. 

  Or we access it via REST API, file based API (think sysfs), and 

````
: read  ( adr len -- actual )
nip
;
: setup-alias
" nvram" find-alias 0= IF
" nvram" get-node node>path set-alias
ELSE
drop
THEN
````
___________________________________________________________--
**CAN WE GET DETERMINSITCI LAUCNEHS? IF SO WE WON THE VM GAME**

Using TFTP, PXE, send-key, etc. This is our bread and butter of portalg-un now

````
: load  ( addr -- len )
s" load" obp-tftp-package @ $call-method 
;
: ping  ( -- )
s" ping" obp-tftp-package @ $call-method
;
: setup-alias
" net" get-next-alias ?dup IF
get-node node>path set-alias
THEN
;
setup-alias
���������u0rtas-nvram.fs." Populating " pwd cr
0 VALUE my-nvram-fetch
0 VALUE my-nvram-store
0 VALUE my-nvram-size
````


_______________________________________________________
[Multiverse OS: `portalgun` kvm/qemu boot process]:

 * Our Open Source BIOS, i.e. OVMF (supports signed boot)


 * Grub2 


 * initramfs


 * Linux Kernel

___________________________________________________________
## Unique Multiverse OS Look

  * Modify Grub UI

  * root@host:/usr/share/qemu#  Modify qemu .svg



___________________________________________________________
## BIOS Options

  * OVMF (Supports signed boots)

  * openbios

  * open hackware ppc

   


root@host:/usr/share/qemu# ls
bamboo.dtb	  petalogix-ml605.dtb	    QEMU,tcx.bin
keymaps		  petalogix-s3adsp1800.dtb  QEMU,VGA.bin
openbios-ppc	  ppc_rom.bin		    slof.bin
openbios-sparc32  QEMU,cgthree.bin	    spapr-rtas.bin
openbios-sparc64  qemu-icon.bmp		    trace-events-all
OVMF.fd		  qemu_logo_no_text.svg

___________________________________________________________
## Virtio scisi!! 


`root@host:/usr/share/qemu# cat slof.bin `


virtio-scsi-init-and-scan
virtiodev virtio-scsi-shutdown

___________________________________________________________
## Additional Features


  [Physical key system]
  * Support using a USB key to boot Multiverse OS (otherwise boots to vanilla Debian9 qcow2 saved to disk). [Can be paired with passowrd, using the password to determine how the machine is booted, ie: one boots to vanilla debian9, one wipes the drive and boots to vanilla debian, backs up to offsite, wipes, boots to vanilla debian, etc]

    Using USB automatically makes the USB key a `dead-man-switch`/`kill-switch`, that will wipe the memory and shut down the computer when removed.
