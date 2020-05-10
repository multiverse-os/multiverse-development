# Host packages
**Boot-chains & network-chains** are the going to be the name used for our networking and booting concept.

**multiversed** A daemon to control host, provide extremely limited API to control it. Including pass packages for installation. Each command requires signature from admin key. Controls VM startup, switching vms, enabling and disabling PCI devices. Simplifed control over VM resources. Control over network-chains and boot-chain.  [Maybe some of this will be migrated over to a special boot-chain for administation on the host, to prevent access to this aspect of the machine without having physical access]. 

**usbguard** Use this to until we build a similar program to lock down the host 
usb. And rely on services to autoboot the controller and networking, if the 
controller fails to boot fall back to emergency mode. 

*MultiverseLogin* (We can simplify the boot process by only encrypting the VM images. Or we can do this with luks. This would also allow us to only checkout and decrypt the disk images in use while leaving the other controllers securely encrypted. This screen should probably check if there are other nodes in this users cluster, and to do that they will need to "login" to the p2p cluster discovery system.] Switch default login screen with status screen for VMs booting. Show the 
following details:
	* Networking topology
	* Endpoints (IP Address, region, etc) 
	* Controller status 
	* A menu with the following options:
		* Available controllers	
	* Install new controller which lets the user select from a collection 
	  of registered operating systems which on alpha release will include:				* Debian
			* Ubuntu
			* MultiverseOS or possibly named GravityOS 
			  (which is based heavily on Debian)
			* SteamOS (which is based heavily on Ubuntu)
			* Fedora
			* OSX
			* (Maybe: Fuschia)
			* (Maybe: One of the Linux distributions with windows 
			  executable support and themed to present itself in a 
			  similar way to windows)
*OR* The original idea was to have this inside a maximally simplifed controller
so that the user never directly itneracts with the host machine, reducing its
attack surface as much as possible. The goal is to make the host difficult to
near impossible to interfere or intercept data on the VMs. 


**Install/remove package script** built a script to take the output of apt `list --installed` and generate a
install/remove package script to put the current system, or preferably the default packages that come with 
Debian inline with the output. This will simplify after setting a machine up,
generating scripts to provision another machine. 

**Create an image that is mounted and takes in log data** from the host machine
so it can be sent seamlessly zero copy to the controller. 

**Start 




