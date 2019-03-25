## Multiverse OS: Portal Gun


#### TODO

  * Restore and rebuild universe0 router (the one that is not a hypervisor)
    so that both options are possible. Accidently merged it in without
    keeping original. (check git obs)

#### Notes / Scratch / Research
We can have console output feed into stdout instead of using a virtaul display
on the VM by specifying intird, kernel and adding `console=ttyS0` to cmdline
attribute:

````
<os>
  ...
  <kernel>/boot/vmlinuz-4.9.0-6-amd64</kernel>
  <initrd>/boot/initrd.img-4.9.0-6-amd64</initrd>
  <cmdline>console=ttys0</cmdline>
  ...
</os>
````

It does require both a kernel and initrd. For portal gun, we could build our
own custom version that ports the output to our user interface controler VM,
and allows us to send keys/commands from the lowest level of boot. (This is the
equivilent to modifying the grub on the guest). 

Obviously we do not want to use HOST `/boot/*` folder, but this example will
work and it will illustrate the point; that we can have all the output of the
VM be routed to stdout of the HOST (or probably a different VM). 

This does make the typical output black since this is an override. 


## OS Info

There exist a C library with a bunch of OS informaiton in XML that can be used like download locations, in addition below is from GNOME boxes:


```
  private string[] recommended_downloads = {
        "http://redhat.com/rhel/7.5",
        "http://fedoraproject.org/fedora/28",
        "http://fedoraproject.org/silverblue/28",
        "http://ubuntu.com/ubuntu/18.04",
        "http://opensuse.org/opensuse/42.3",
        "http://debian.org/debian/9",
};
```



________________________
**OLD**

## Multiverse OS Machines
This folder contains all the basic machine templates that will be provided with
the alpha version of Multiverse OS. Each folder contains all the files needed
to quickly provision and deploy a Mutliverse OS machine.





#### Machine Folder Structure
  * Config
  * XML
  * Scripts


### Config Folder
Each machine involved in the cluster including the bare-matel machine will have
an associated folder in the `Data` folder of the Multiverse OS root folder.
Each of these folders will contain a config folder, that is will indicate all
the files that need to be modified with its structure. Eventually full configs
in this folder will be switched out with templates that will fill in values
based on system configuration and installation configuration options selected
by the user. 


### Scripts Folder
Currently there are no scripts folder but some incarnation of this will exist to allow customizations ontop of the ephemeral templates. The scripts folder will either have sub-folders for to organize the different types of scripts, hooked to different events: [on boot, before shutdown, on disconnect, etc], or the event hook will be specified at the top of the script. 

These scripts will be loaded into the VM Agent.


### Data Folder
The Data folder will contain custom user data accessible to the machine.

______
## Development
Multiverse OS alpha will not include a full version of the clear containers provisioning and management software called `portal-gun`. But pieces of it will be included in the alpha version, and development of those pieces will be outlined and planned below.

#### Development Brainstorming
Below are a list of potential features that can be implemented to improve and extend the functionality of the VM management portion of Multiverse OS known as `portal-gun`:

  * Each part of the provisioning should be a module, for example network/interfaces. That way these parts can be standardized and make the resulting configuration file simpler and easier to review. 

#### VM Setup / Provisioning
Previous implementations of Multiverse OS relied heavily on the use of special bridges that connected VMs, such as the Controller VM -> Application VM to simplify provisioning virtual machines over SSH. This method adds a lot of attack vectors, requires software that is not necessarily provided by default, and still requires a lot of user iteraction to provision. 

Currently we are quickly setting up shared storage using modifications to `/etc/fstab/` then using the files on the shared storage to configure Multiverse OS VMs. 

##### Current solution
The current solution is to attach a device, either provide a custom kernel module written in Rust or initiate a TTY, and use this to configure `/etc/fstab/` to provide shared storage then using shared storage provision the virtaul machine. 

This limits the attack surface, allows for access wtihout internet access, and does not require almost any user interaction. This can be done easily with bash scripts to satisfy alpha installer requirements and lay foundation for the rewrite in a higher level language that will come shortly aftewards.

______
## Multiverse Networking
Currently networking is hard coded to work on a MAC address system roughly 00:00:10:X:X:X, but we want to support multiple networks with different heirarchies. So we need to modify the provison script to ask questions. 

Look at survey a go language library. THis is the same lib we should use to replace setting up VMs so we can rid ourselves of virt-manager asap.

______
## Development Notes
This folder needs scripts to prepare information about VMs in a useable format for programming. For Example, the `pci-devices` section has scripts that need to prepare PCI devices for use, but for assignment to each virtual machine, we need more scripts to list virtual machines, and specifically by type relevant to Multiverse OS (Controller, Application, Service, Router). 

Then simplified API from shell or preferably for higher level languages needs to be established to change important settings ont the VM, for example easily provide a device ID and a type of Multiverse VM and it adds the device to all of the VMs within that category that exist.
