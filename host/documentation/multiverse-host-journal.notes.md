## HOST Development Roadmap 
Guide to developing Multiverse OS componetns required for the HOST AGENT operation, defense *honey-pots, other traps, lockdown, etc*, VM contorl, autostart, reboots, etc. ...

	* HOST should be modified so that /home /var and /tmp are separated.
	  then these can be mounted from \*.raw or \*.qcow2 (same with
	  controller)



## Networking Updates

After learning about **netmap** a super high speed user space network model 
utilizing **pf_ring**. 

This in combination with __VSOCK__. Since _virtchannels_ the previous idea
for general VM-to-VM communication relies on a older virt-serial that has
a lot of flaws that __VSOCK__ solves it is the obvious choice. 

Along with that net-pci, or a new custom net-pcie would be another great
choice with the entire network stack put on the device. 



______
## Moving to two (2) CONTROLLER model

 (1) USER CONTROLLER, interface for user to interact with nested
     (App)lication and Service VMs. 
     [Example: Ubuntu17 specially used for running steam games, so games are]
     [segreated from rest of computing]


**CONVERT Qcow2 used on the 640GB main image to RAW for major preforamnce incresae**

 (2) ROUTER CONTROLLER
     _passthrough of network devices_

```
<network>
  <name>passthrough0</name>
  <forward mode='hostdev'  managed='yes'>
    <pf dev="eth0" />
  </forward>
</network>

```

[exit with `:wq passthrough.xml`] then define it after setting it up (can define it with various options or better owuld be create defitnions on the fly based on the varibales that wec can fill in, then create a (dunno webform?) to input the data and create the XML then auto-define it (see extra task that should not exist really below, again why we are getting rid of libvirt).


`virsh net-define passthrough0.xml`

`virsh net-autostart passthrough0`

[QUESTION][Is name the name of the device on the VM? This could be useful is so]
  
  [__ <hostdev> (*conventional*) VS. <virtualprot>
   

  [*] There are two (*2*) ways of doing passthrough now:
    (1)[conventional method is with the <hostdev> element]
      straight forward

    (2)[the new method is using <interface type="hostdev"> element]

      add two (2) new elements:
        [1] virtual port with attribute `type` can take value like `802.1Qbh` to mimic wireless card virtually.

       [!][USECASE]**During the development of Multiverse OS, we talked about wanting to use `mpath tcp` basically let packets travel DIFFERNET paths to a server, then let the sevver reassemble the packets from different sources**
   this has several interesting consequences:
     _which may speed up, but definitely add more psuedanommity, lessen an organizations ability to have all the information, support mesh networking easier, and other complex networking structures_

     so we wnated to do TCP/IP over lasers, over HDMI, over Audio (like a modem), and many other techniques, to make the internet far more complex and essnetially impossible to track relaibily. to test this we would need to start with simple ones like HDMI's CEC, audio, laser, and so on. using this new passthrough setup, we can easily incorporate this into Multiverse OS with this. 

     **it also allows for a** `paramters` **element which as the attrbiute profileid** so we can specify maybe this specialized PCI netowrk card is specifically for hgih frequency trading, so we set "profileid='finance'"

     [IMPROTANT USECASE] Additonally Mutliveres OS from the beginning was only using libvirt and virt-manger to simplify early design and cluster construction. but due to security cocnerns with malformed packets being the most common way to breka out into the hypervsior, we NEVER wanted to use the HOST networking, or host based bridges, we wanted to use custom network stack that does VM-to-VM tcp/ip transfer of packets.
     **this has the added securityy benefit of allowiung us to abstract ontop of tcp/ip addtitional features like packet signing, packet encryption. custom packet streaming protocols.** using this new technique, we can implement our own custom __VIRTUAL__ PCIe Network Card that can be assigned to all the virtual machines and networked together over VirtIO to start with but eventually move to `/dev/shm` and DMA (direct memory access).

   In fact we can put virtual memory ON the card, store the ETNRIE network stack on the card in a read-only way, have a write only ring buffer type system and do zero-copy where possible. Then we write directly from one card to another, base don protocols defined in a read only memory. broken up in best ways to access it, and config held in memory on the chip that is read only as well or maybe r/w./

  we dont get a perment unique MAC using this SR-IOV VF(virtual function) (this new passthrough) **BUT AN AGENT CAN SET THIS ON EACH REBOOT, np!**

   [!!!][VIDEo][ANPOTHJER FUCING USE CASE!] 
    instad of using a actual video card, we just have the drivers to load binary data DIRECTLY from the /dev/card0 device and pass it in binary, and process binary to wayland protocl on the CUSTOM NIC, and load it into a window on the controller.
 
   **OTHER BENFITS OF NEW METHOD!**

     * can be PCI or PCIe (aweomse! we can use this to prototype our fucking sci-equipment and custom Mutlvierse OS hardware!)

     * SUPPORTS SUB-DRIVERS, YES YOU READ THAT RIGHT!

     * builtin support for VLAN!! (look it up on redhat doc, lots of good options!)


    [<mac>, <vlan>, <virtualport>] are children of <interface> element.






___________________________________________________________________________
## Roadmap

[O][Implement ROUTER CONTROLLER]
  [*][ROUTER CONTROLLER _should_ be Universe Router wasted resources]
     So the design is:
      __HOST__:
        ROUTER_CONTROLLER(Universe0, PCI passthroug of networking
     	  [*] Galaxy0 VM
    	  [*] Star0 VM (maybe isnie galaxy, probbaly not)

        

  [*][Nest VMs]
  [*][Re-implement each router VM, use OVMF, so we can have secure BIOS to kernel] 
     **Rember to change names to newer naming structure: router.galaxyX.multiverse, where X is the ordinal VM starting with 0.**
      


[!][HOST Agent]
  [*][VM Autostart]

  [*][Kernel Patches for HOST]
    [*][Kernel lockdown of Input/Output/Error]


  [*][Defense]
    [*][Honeypots and other traps]
    [*][file integrity checking]
  
  [*][Host OS update system (without networking)]


  [*][Making the HOST epheraml!]


  [*][Making the HOST immutable (where it counts)]



_________________________________________________________________
## Notes/Research/Scratch on Host configuration and setup

[development env tip]
#  While we are using stupid virt-manager, might aswell link our images foleder to their default one they keep creating.

````
rm -rf /home/user/.local/share/libvirt/images
ln -s /home/user/multiverse /home/user/.local/share/libvirt/images
````



    


_______________________________________________________________________________________________
## Notes

  * Make a version of pass-store that automatically syncs between LAN based git instances 
    within the Multiverse clsuter. So adding to controller, updates the host, vice versa.


  * Successfully got `encrypted data storage disks` to only mount/decrypt inside controllers
    and they are not mounted or touched in the host. This prevents issues of formatting and 
    trying to boot off the disk from the host. 

    [!] Since ALL the host machine, the controller and the data store disks are encrypted
        it would be impossible to do anything but delete other disks. So at most it would
        be a troll, and not a security breach.

_______________________________________________________________________________________________
## Disassocating with encrypted data storage disks on Host machine



**Old fstab on host, data should be moved into controller that uses the disk.**
````
/dev/mapper/host--vg-swap_1 none            swap    sw              0       0
/dev/mapper/luks-8efad1a2-d5a3-413c-ac27-c0b1524d7064 /media/user/Vault ext4 nodev,nosuid,relatime,x-gvfs-show 0 2
/dev/mapper/luks-735bcd84-3306-48c7-b5a7-774ea11d8a81 /media/user/Zone ext4 nodev,nosuid,relatime,x-gvfs-show 0 2
#/dev/disk/by-uuid/735bcd84-3306-48c7-b5a7-774ea11d8a81 /mnt/735bcd84-3306-48c7-b5a7-774ea11d8a81 auto nosuid,nodev,nofail,x-gvfs-show 0 0
#/dev/disk/by-uuid/022bd185-eab7-4faf-96ac-b3e74ecc9f65 /mnt/022bd185-eab7-4faf-96ac-b3e74ecc9f65 auto nosuid,nodev,nofail,x-gvfs-show 0 0
/dev/disk/by-uuid/ca4d01fa-214a-460d-83e3-40bde9ec3ab4 /mnt/ca4d01fa-214a-460d-83e3-40bde9ec3ab4 ext nosuid,nodev,nofail,x-gvfs-show 0 0
/dev/disk/by-uuid/50b65b19-124c-4000-888c-289c05b7fa13 /mnt/50b65b19-124c-4000-888c-289c05b7fa13 ext nosuid,nodev,nofail,x-gvfs-show 0 0
/dev/disk/by-uuid/ca4d01fa-214a-460d-83e3-40bde9ec3ab4 /mnt/ca4d01fa-214a-460d-83e3-40bde9ec3ab4 ext4 nodev,nosuid,relatime,noauto 0 2
````
_______________________________________________________________________________________________


* **Stop using /home/user/* for `multiverse` folder location.** It could eventually be found
to have directory transversal issues and being relative to the home folder is dagnerous. 
since it has keys, libvirt config, images, etc.

  `/var/multiverse`


**Until an install script exists**, that pulls down data directly to the `/var/multiverse/` and `/etc/multiverse` respectively, in addition to use folders for user specific data. Then use the following stop-gap solution: 
````
su
mkdir /var/multiverse
cd /var/multiverse
ln -s /home/user/multiverse/images/ .
ln -s /home/user/multiverse/machines/ .
ln -s /home/user/multiverse/images/os-images/ .
ln -s /home/user/multiverse/machines/host.multiverse/scripts/ .
```
--- 
## Host/QEMU Controller Configuration

Raw because its a whole disk, may be able to use other types but raw is easy, and fast. 

Cache should be none for LVM. None is the fastest but not safest. Best for clustering. Writethgouh has the best data intergity and safety features.  and delete the address, and it will autofill.

    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none'/>
      <source dev='/dev/disk/by-uuid/41f02abc-defa-4c21-b2eb-94750ccc4730'/>
      <target dev='vdb' bus='virtio'/>
    </disk>


All the QEMU XML configuration can and should be done without `virt-manager`, because it is unreliable software, that can be unpredictable and most importantly does not have access to all possible XML modifications. The most important QEMU XML configrations are not accessible from `virt-manager`.

To pass through a physical hard drive, determine the disk's UUID (for example, using `lsblk -f`, `blkid` or gnome-disks), and add the following to the VMs xml. The final letter of the `target dev` element cannot be used by more than one device (for example, if "vda" already exists, name the disk "vdb" or "vdz"). Keep in mind that you can not pass through the physical disk that the host machine's operating system is installed on.
wd
```
    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='writeback'/>
      <source dev='/dev/disk/by-uuid/41f02abc-defa-4c21-b2eb-94750ccc4730'/>
      <target dev='vdb' bus='virtio'/>
    </disk>
```

Then add it to the udev rule `/etc/udev/rules.d/61-hdd-permissions.rules`. Add a new line for every hard drive. Note: because we are matching against an environmental variable set by udev rule 60, the rule number of this file must be 61 or higher.

```
ENV{ID_FS_UUID}=="41f02abc-defa-4c21-b2eb-94750ccc4730", GROUP="libvirt"
```

To test the rule, reboot or run:

```
udevadm control --reload
udevadm trigger /dev/sdX
```

where `sdX` is the device name, and check that the device is in the libvirt group.

This will be reaplced with`vdX` if we switch to using virtual drives.


_______________________________________________________________________________
## Post Installation Controller Configration
The base configuration needed for Controller VM's of any category.



````
# On the HOST machine
sudo apt-get install ssh-askpass
````


````
# Edit the /etc/hosts file to add the hostnames for the two controllers 

````



````
# Add the connections for virt-manager (hopefully wont need to do this soon)
gsettings set org.virt-manager.virt-manager.connections uris "['qemu+ssh://root@router.controller/system?socket=/var/run/libvirt/libvirt-sock', 'qemu:///session', 'qemu+ssh://root@router.controller/system?socket=/var/run/libvirt/libvirt-admin-sock', 'qemu://system']"

````




===================================================================================
## Host Lockdown
-----------------------------------------------------------------------------------

Get rid of gnome terminal server


## Remove Excess Software

Below are applications testing removal, and doing it in chunks to make sure we avoid breaking.

Check out removing `apt-get remove gnome-bluetooth` and get rid of the games until all of gnome can go. 




````
# Disable suggestions and recommendations in apt-get settings
apt-get remove evolution remove pulseaudio wpasupplicant packagekit 
packagekit-tools ssh five-or-more telepathy-idle telepathy-logger 
telepathy-mission-control-5  # Gnome contacts gimp gimp-data

````

## Hunt down all processes that are listening on a port


## Remove excess drivers and kernel modules


## Remove gnome and get rid of all host UI



## Create kernel module, modify kernel, or libinput, or higher level

To disable input into host machine console. 

Encrypt all output, require signed input.




