## GRUB Based CONTROLLER VM Fallback system


 /boot/grub/x86_64-efi$ cat terminal.lst 
iat_keyboard: at_keyboard
iserial: serial
iserial_*: serial
oaudio: morse
ocbmemc: cbmemc
ogfxterm: gfxterm
oserial: serial
oserial_*: serial
ospkmodem: spkmodem



[!][CUSTOMIZE BOOT VISUALs FOR MULTIVERSE OS AESTHETIC]**(Modify /boot/grub/* and then write a script/program that will automatically update a /boot/grub to Multiverse-IZE various BOOT looks)** 
	_and base it on VM type: so differnet look for APP VM, Service VM,_
	_Controller VM, and HOST)_

	[!][Upgradable items]
	    [*] unicode font (?)
	    [*] grub.cfg to modify ENTRY NAMEs
 	        __(set video basd on graphics card/VM type)__
                [IMPORTANT][NOTE] Can define kernel modules to call
		in at this point, so we can call in AMD or NVIDIA
 		and probably get video working for the TYPING in part
		of cryptsetup.
		

	       [*][still grub.cfg]
          	[MEGA IMPORTANT][NOTE] Can add new entries, for example
	        FALL BACK CONTROLLER, or just a VM that has GRUB, and
		we generate a list of the available CONTROLLERs, and
		then we boot to selected controller in GRUB menu
 
		   _[legitiamtely VERY INTERSTING idea/solution for_
		   _the fallback CONTROLLER VM]_


