## Configure Default CONTROLLER VM with intel i9 CPU and amd vega GPU



________________________________________________________________________________________________
## Notes




* **Unix conciously chose not tuse binary formats inc ofngis, but that was in the 70's things have changed**


* **CONTROLLER GNOME SETTINGS** 
  [!][Turn off automatic suspend, by default it is 20 minutes]



* **Stop using /home/user/* for `multiverse` folder location.** It could eventually be found
to have directory transversal issues and being relative to the home folder is dagnerous. 
since it has keys, libvirt config, images, etc.

  `/var/multiverse`


````

* All controllers and likely routers should have /dev/random attached as a device.

* All controllers should ahve TPM or some way of doing TPM. 

* A special PCI card with RTC maybe even high precesion HPET clocks should be placed in a
a grid to provide RTC dedicated to each VM. Would solve a lot of problems, increase speed
and genearlly be aweomse. 
________________________________________________________________________________________________

## Add passthrough PCI devices

2x PCI devices for AMD graphics card

1x PCI device for USB hub


--- 
## Host/QEMU Controller Configuration

**The physical harddrives bieng connected to a controller VM should be MUCH simpler. Just show a list**
**of the HDs, then just in the VM config, check the checkbox for each HD you want attached or something**
**similar**. Then it would automatically be added to XML and the udev rule made so it would be secure. 
Right now the UI for this stuff is garbage and itd be very easy to make it 1200x easier to use and
just default doing it the secure way that supports unpriviledged non-root users. Unlike other "secure"
OSs out right now.

All the QEMU XML configuration can and should be done without `virt-manager`, because it is unreliable software, that can be unpredictable and most importantly does not have access to all possible XML modifications. The most important QEMU XML configrations are not accessible from `virt-manager`.

To pass through a physical hard drive, determine the disk's UUID (for example, using `lsblk -f`, `blkid` or gnome-disks), and add the following to the VMs xml. The final letter of the `target dev` element cannot be used by more than one device (for example, if "vda" already exists, name the disk "vdb" or "vdz"). Keep in mind that you can not pass through the physical disk that the host machine's operating system is installed on.
wd
```
    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none'/>
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


---

## Post Installation Controller Configration
The base configuration needed for Controller VM's of any category.


