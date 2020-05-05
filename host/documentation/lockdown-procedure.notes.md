# Host Setup 

## Review INTEL I9 CPU FEATURES and disable and add the ones whcih will make the most secure and efficient CPU

## Lockdown
Any attempt at lockdown will fail without getting rid of the default suggests and recommdsn. In Multiverse OS these should be disabled by default, they are complete and utter bullshit. Exampole, removing firefox forces installing chromium. So you can't ever not have a browser by `defualt.s`.

```
If you do not want to install recomended packages you can run apt-get with the --no-install-recommends flag or aptitude with the --without-recommends/-R flag.

If you want these flags to always be enabled (I do NOT recommend this) put the following lines in your /etc/apt/apt.conf file:

APT::Install-Recommends "0";
APT::Install-Suggests "0";
```


______
## Notes

  * Make a version of pass-store that automatically syncs between LAN based git instances 
    within the Multiverse clsuter. So adding to controller, updates the host, vice versa.


  * Successfully got `encrypted data storage disks` to only mount/decrypt inside controllers
    and they are not mounted or touched in the host. This prevents issues of formatting and 
    trying to boot off the disk from the host. 

    [!] Since ALL the host machine, the controller and the data store disks are encrypted
        it would be impossible to do anything but delete other disks. So at most it would
        be a troll, and not a security breach.



______
## Disassocating with encrypted data storage disks on Host machine
**Old fstab on host, data should be moved into controller that uses the disk.**

```
/dev/mapper/host--vg-swap_1 none            swap    sw              0       0
/dev/mapper/luks-{UUID}   /media/user/DISK0 ext4 nodev,nosuid,relatime,x-gvfs-show 0 2
/dev/mapper/luks-{UUID}   /media/user/DISK1 ext4 nodev,nosuid,relatime,x-gvfs-show 0 2
#/dev/disk/by-uuid/{UUID} /mnt/{UUID} auto nosuid,nodev,nofail,x-gvfs-show 0 0
#/dev/disk/by-uuid/{UUID} /mnt/{UUID} auto nosuid,nodev,nofail,x-gvfs-show 0 0
/dev/disk/by-uuid/{UUID}  /mnt/{UUID} ext nosuid,nodev,nofail,x-gvfs-show 0 0
/dev/disk/by-uuid/{UUID}  /mnt/{UUID} ext nosuid,nodev,nofail,x-gvfs-show 0 0
/dev/disk/by-uuid/{UUID}  /mnt/{UUID} ext4 nodev,nosuid,relatime,noauto 0 2
```


______


* **Stop using /home/user/{for `multiverse` folder location}** It could eventually be found
to have directory transversal issues and being relative to the home folder is dagnerous. 
since it has keys, libvirt config, images, etc.

  `/var/multiverse`


**Until an install script exists**, that pulls down data directly to the `/var/multiverse/` and `/etc/multiverse` respectively, in addition to use folders for user specific data. Then use the following stop-gap solution: 

```
su
mkdir /var/multiverse
cd /var/multiverse
ln -s /home/user/multiverse/images/ .
ln -s /home/user/multiverse/machines/ .
ln -s /home/user/multiverse/images/os-images/ .
ln -s /home/user/multiverse/machines/universe0.multiverse.host/scripts/ .
```


______
## Host/QEMU Controller Configuration
Raw because its a whole disk, may be able to use other types but raw is easy, and fast. 

Cache should be none for LVM. None is the fastest but not safest. Best for clustering. Writethgouh has the best data intergity and safety features.  and delete the address, and it will autofill.

```
    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none'/>
      <source dev='/dev/disk/by-uuid/{UUID}'/>
      <!-- <target dev='vdb' bus='virtio'/> Older verisons we would do v** -->
      <target dev='sdb' bus='virtio'/> <!-- Here we can switch it to sdb -->
    </disk>
```

All the QEMU XML configuration can and should be done without `virt-manager`, because it is unreliable software, that can be unpredictable and most importantly does not have access to all possible XML modifications. The most important QEMU XML configrations are not accessible from `virt-manager`.

To pass through a physical hard drive, determine the disk's UUID (for example, using `lsblk -f`, `blkid` or gnome-disks), and add the following to the VMs xml. The final letter of the `target dev` element cannot be used by more than one device (for example, if "vda" already exists, name the disk "vdb" or "vdz"). Keep in mind that you can not pass through the physical disk that the host machine's operating system is installed on.

```
    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='writeback'/>
      <source dev='/dev/disk/by-uuid/{UUID}'/>
      <target dev='sdb' bus='virtio'/>
    </disk>
```

Then add it to the udev rule `/etc/udev/rules.d/61-hdd-permissions.rules`. Add a new line for every hard drive. Note: because we are matching against an environmental variable set by udev rule 60, the rule number of this file must be 61 or higher.

```
ENV{ID_FS_UUID}=="{UUID}", GROUP="libvirt"
```

To test the rule, reboot or run:

```
udevadm control --reload
udevadm trigger /dev/sdX
```

Where `sdX` device name; check that the device is in the libvirt group.

This will be replaced with`vdX` if we switch to using virtual drives. 


> **EDIT** *The newest versions we switch it back via declaration in libvirt which will be phased out, but the idea is to attempt to start hiding the most glaring indications we are in a VM to not scare off viruses to better honeypot them and trigger alarms*


______
## Post Installation Controller Configration
The base configuration needed for Controller VM's of any category.


###############################################################################
###############################################################################
## Host Lockdown
###############################################################################
Get rid of gnome terminal server


## Remove Excess Software
Below are applications testing removal, and doing it in chunks to make sure we avoid breaking.

Check out removing `apt-get remove gnome-bluetooth` and get rid of the games until all of gnome can go. 

```
  # gnome mail client
	apt-get remove pulseaudio
	apt-get remove wpasupplicant 
	apt-get remove packagekit packagekit-tools # gnome update

	apt-get remove ssh

        apt-get remove five-or-more


apt-get remove telepathy-idle telepathy-logger telepathy-mission-control-5  # Gnome contacts
apt-get remove gimp gimp-data
 

```

## Hunt down all processes that are listening on a port


## Remove excess drivers and kernel modules


## Remove gnome and get rid of all host UI



## Create kernel module, modify kernel, or libinput, or higher level

To disable input into host machine console. 

Encrypt all output, require signed input.




