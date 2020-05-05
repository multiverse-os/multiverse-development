<p align="center"><img src="https://github.com/hackwave/multiverse-development/blob/master/multiverse-logo.png" height="300" width="300"></p>

# Multiverse OS Install Notes
Multiverse OS currently can only be installed through a manual process, the following is the final iteration of the manual installation guide that will be converted to the the first alpha installer. 


______
## Updates (to be merged into instlal notes)


  [HOST][Additions to HOST initial setup & configuration]
  * Initialize GPG key (until scrable suite key system is in place), use this GPG key to initialize pass-store (`sudo apt-get install pass` followed by `pass init {key_name}`).

  [HOST][Setup APP/SERVICE CONTROLLER VM and ROUTER CONTROLLER VM]
  * Move all VMs into repsective VMs. Currently, both are Debian9. But experiments are being conducted on using Alpine3.7 for the ROUTER CONTROLLER VM. 



## storage disks
````
    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none'/>
      <source dev='/dev/disk/by-uuid/735bcd84-3306-48c7-b5a7-774ea11d8a81'/>
      <target dev='vd[X]' bus='virtio'/>
    </disk>

    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none'/>
      <source dev='/dev/disk/by-uuid/022bd185-eab7-4faf-96ac-b3e74ecc9f65'/>
      <target dev='vd[X]' bus='virtio'/>
    </disk>

    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none'/>
      <source dev='/dev/disk/by-uuid/8efad1a2-d5a3-413c-ac27-c0b1524d7064'/>
      <target dev='vd[X]' bus='virtio'/>
    </disk>
````

______
## Installing from Debian disk
Because Multiverse OS does not have its own ISO yet, we start by installing Debian. (Use the script below in the [Obtaining ISO Images section](#obtaining-iso-images) to download, verify the checksums and verify the signature on the checksums.) 

Use `host` for the hostname, no domain, `user` for the `user`, encrypt the HD using LUKS and use a good password for both `root` and `user`. 

*Debian desktop environment* is optional and will be removed at the final lockdown phase, so select whichever is comfortable to you (including none). Only select *System Utilities*, **DO NOT** install *SSH server*, *Print Server* or *Web Server*. 

| Installation Options     |  Value                  |
|--------------------------|-------------------------|
| **Hostname**             | `host`                  |
| **Username**             | `user`                  |
| **Packages**             | *System utilities*      |


______
## Preparing the environment

In an ideal setup, all necessary files are predownloaded and the **Host Machine** never touches the internet. Either use the packages included in this git repository (available for amd64 architecture), copy necessary packages from the `/var/cache/apt/archives/` folder of another install, or use **apt-get**. If using **apt-get**, limit the access to the internet on the **Host Machine** by only using **apt-get** and once necessary packages and os images are downloaded, immediately stopping and disabling the network service with `systemctl stop networking && systemctl disable networking`.

**During a manual Multiverse OS install, the browser should absolutely NEVER be opened on the Host Machine.** For increased security, do not add the user to `sudo` group.

The `apt` packages to avoid needing to access the internet beyond obtaining this original repository are located within `./machines/host.machine/os-install/apt-packages`. They can be installed manually using `dpkg -i {package}.deb`, or the `*.deb` files can be copied to `/var/cache/apt/archives/` on the **Host Machine** so `apt-get install` can find them automatically.

The following instructions assume the commandline text editor to be used is vim, this can be replaced with emacs, vi, or your preference.

```
su

apt-get install sudo ovmf qemu virt-manager git pass vim dirmngr
# virt-manager will be phased out very soon
# vim is optional, install preferred command line text editor
# dirmngr is needed for validation of GPG signature on images

# DEV: These libraries are useful during development
# so the controller can connect virtmanager back to
# the host and control the 3 routers from the controller
sudo apt-get install openssh-server openssh-sftp-server ssh

# Remove nano to make vim the default editor
apt-get remove nano
# nano is removed to make vim the default editor
```

*Other packages that may be required in the future: ...*


#### Obtain the 'multiverse-development' repository
Development of the alpha Multiverse OS installer is actively progressing at `https://github.com/hackwave/multiverse-development`. Obtaining this repo is not required to install Multiverse OS manually but it includes all configuration files and includes scripts to simplify installation. The repository issues can also be used to obtain support, and pull requests that satisfy requirements are accepted and appreciated. 

```
cd /home/user/.local/share
git clone https://github.com/hackwave/multiverse-development multiverse
```

**DEVELOPMENT NOTE:** For increased security, components of `multiverse` working directories should be eventually moved to `/var/multiverse`, for the same reason web servers are hosted from `/var/www`. Avoiding use of the user home folder limits obtainable data from potential directory transversal bugs. For now we will leave them in the home folder for ease of manual setup. In addition, `~/.local/share/multiverse`, `/var/lib/multiverse/`, `/usr/lib/multiverse/`, /etc/multiverse` could be used as well. 


#### Setup the Multiverse OS 'pass-store'
The **Host Machine** `pass-store` holds all VM LUKS HD passwords and both `root` and `user` password. In contrast to QubesOS which opts to remove all passwords, Multiverse OS opts to encrypt all 32 character random passwords. 

While QubesOS relies solely on ephemerality, Multiverse OS always opts for a defense-in-depth approach.

```
cd /home/user/
pass init
```

#### Git Configuration
The basic backup system and most likely the **apt-get** update system will be built using `git` and signed commits.

**NOTE** Do not put your email or name, pseudo-anonymity is acheived by using a purposefully generic default option. Custom names can be utilized at higher levels, the user wil never interact with the **Host Machine** and so it should NOT be configured to the user. 

```
git config --global user.email "you@example.com"
git config --global user.name "Your Name"
```

This is not actively being used outside of development, but it may be important for both backing up and development.


______
## Obtaining ISO Images
Multiverse OS is currently using a combination of Debian and Alpine Linux. Debian is used for a majority of systems, including the base system, and Alpine Linux is used exclusively for **Router VMs**. A Whonix Gateway VM can optionally be used for the [Star Router VM](#building-star-router-vm). Additional distros and operating systems can be used for controller and application VMs, but will not be covered here in depth.

The ISO download scripts below can also be found in the `images` folder of this repo.

#### Why Debian?
Debian has coverage across A LOT of devices, including ARM64 and many others. Changes will need to be made to the kernel eventually to support RTOS and other customizations needed to satisfy Multiverse OS design, but Debian makes a great starting point since multiple Tor-based and other security projects use it, it's committed to the mission of free and open software, it's well supported, and has wide device coverage. 

#### Why Alpine?
Alpine was chosen because it is a solid microkernel that has a much smaller footprint than vanilla Debian. In addition, using a different operating system for routers prevents the cluster from being a monoculture. The diversity of operating systems reduces the possibility of a single exploit being used on every building block of the Multiverse OS cluster. In addition, it is built with security in mind, it comes with `grsec` security patches, which Debian does not come with by default, it uses patches to C which enhance security , includes a much more limited set of software than vanilla Debian, and where there is software overlap, Alpine is frequently running different software versions than vanilla Debian. 

### Securely Obtaining the Linux Install/Live Images
One can obtain the Debian ISO using the following **wget** command without needing to open the browser. In addition, the **wget** commands for the signatures are supplied and they must be checked to have a secure Multiverse OS system.

A Debian script, along with other operating systems: [ubuntu, fedora, alpine, debian], is included to assist in securely downloading and verifying the distribution install and live images. Verification is done by using both the checksum and gpg signature to verify the validity ofthe checksum. 

These scripts are found in `images/os-images`.

**DEVELOPMENT NOTE:** Eventually an automated system that relies on P2P verification to add additional layers of security on important files like this will be built into Multiverse OS. 

**NOTE** If the output says "Good signature", then the file is valid, you will receive a WARNING is not certified with a trusted signature because you have not set the key as trusted, setting the key to trusted in GPG will remove the warning.

**DEVELOPMENT NOTE:** Ideally we should be using a minified microkernel version of Debian to replace Alpine. Optionally SEL4 looks like a good option as well. While Alpine has proven to be a good choice, it is not ideal and there are several issues, including but not limited to: no https download options for ISO images (arguably not important with checksums and signatures), and other similar trust issues. 


##### Obtaining the Alpine Linux Image
As above, a simplified download and verify script has been created because the process of verification is critical to Multiverse OS installation security and simplifying this procses will ensure that it happens, and it is all too often overlooked. MITM attacks are real and serious.

**DEVELOPMENT NOTE:** Both of these scripts need to be updated and improved by abandoning bash and using ideally Go. Multiverse OS will primarily be Rust for lower level components and Go for higher level components, similar to how Debian is made up of C and Python. 


______
## Userspace Virtual Machines
From the start, Multiverse OS was designed to be used from unprilvileged userspace, so that any potential breakouts or security vulnerabilities do not lead to exposure of the root account. 

The **Host Machine** is the bare-metal machine, the building block of the Multiverse OS cluster.

To do this virtual machines must be run in userspace. To do this the permissions and configuration of the **Host Machine** must be correct. 


#### User
This account should have been created during the installation of Debian.

Multiverse OS is run from the account `user`, so this user first needs to be added to the appropriate groups, specifically: `kvm` and `libvirt`.

**ERROR** If you receive a permission denied error relating to KVM kernel, then your user has not been added to the `kvm` group. Remember, you will need to log in and out for the group permission to take effect. [__HINT__] You can avoid logging in and out by using either  `su - user` or what is essentially the same thing `exec su -l user`. (_Generate this error and save detailed verbatim messages so it can be easily CTRL+F in this document for people who are stuck at an error_)

**NOTE** _You will need to logout and log back in for your new group based permissions to take effect._ Keep in mind that typing only `groups` will show the __active__ groups for the user, the __saved__ groups for the user can be determined by using `groups user`.


```
# Some guides/answers for adding a user to group will not add the -a flag
# -a is short --append, whcih is essential, because otherwise the user will
# be REMOVED from any groups not in the list
usermod -a -G kvm user
usermod -a -G libvirt user
groups

# An example list of groups from an install with the Gnome desktop environment:
# user cdrom floppy audio dip video plugdev netdev bluetooth scanner libvirt kvm

# Minimal list of groups:
# user plugdev netdev libvirt kvm
# (Will be reviewed further in lockdown process)

```
**DEVELOPMENT NOTE:** We should remove any groups in this list that are not absolutely necessary as part of the lockdown process of the **Host Machine**. 




#### Configuring userspace memory limits
Next the memory limit needs to be removed so user session VMs can access the whole machine's memory. Two lines need to be added to the `/etc/security/limits.conf file:

```
@kvm             soft    memlock         -1
@kvm             hard    memlock         -1
```

### Configuring User Session Storage
XML files for the storage pools discussed here can be found in `./machines/host.multiverse/xml/storage/`. The content of the files is also included below for convenience.

The `images` storage pool stores all the images both the debian and alpine linux distribution images but also the machine `qcow2` images. Create the pool with `virsh pool-define images.xml`. Then ensure it starts on boot with `virsh pool-autostart images`.

**images.xml**
```
<pool type='dir'>
  <name>images</name>
  <capacity unit='bytes'>0</capacity>
  <allocation unit='bytes'>0</allocation>
  <available unit='bytes'>0</available>
  <source>
  </source>
  <target>
    <path>/var/multiverse/images</path>
  </target>
</pool>
```

The `machines` storage pool encompasses shared storage that is mounted to each machine using `9p` and contains the VM's xml file (used by `virsh define`), a human-readable README describing the machine, configuration files, and a provisioning script to install the config files on the VM.

Create the pool with `virsh pool-define machines.xml`. Then ensure it starts on boot with `virsh pool-autostart machines`.

**machines.xml*
```
<pool type='dir'>
  <name>machines</name>
  <capacity unit='bytes'>0</capacity>
  <allocation unit='bytes'>0</allocation>
  <available unit='bytes'>0</available>
  <source>
  </source>
  <target>
    <path>/var/multiverse/machines</path>
  </target>
</pool>

```

And finally `os-images` is a folder within the images folder containing just the iso images used for installing operating systems. This is segregated to allow it to be mounted to a VM without needing to mount all the other VM images. 

**os-images.xml*
```
<pool type='dir'>
  <name>os-images</name>
  <capacity unit='bytes'>0</capacity>
  <allocation unit='bytes'>0</allocation>
  <available unit='bytes'>0</available>
  <source>
  </source>
  <target>
    <path>/var/multiverse/images/os-images</path>
  </target>
</pool>

```

And finally `os-images` is a folder within the images folder containing just the iso images used for installing operating systems. This is segregated to allow it to be mounted to a VM without needing to mount all the other VM images. 

### Networking
In the future, Multiverse OS networking will have all networking handled in a userspace network stack that is built outside of the kernel and avoids all kernelspace tools. This is important because all major breakouts from VMs have been using specially crafted packets. Currently with the libvirt virtual networks, creating a virtual bridge connects the **Host Machine** to the bridge with a virtual network interface. This will be changed to truly isolate the **Host Machine** in the userspace network stack implementation to come.

### Networking
In the future, Multiverse OS networking will have all networking handled in a userspace network stack that is built outside of the kernel and avoids all kernelspace tools. This is important because all major breakouts from VMs have been using specially crafted packets. Currently with the libvirt virtual networks, creating a virtual bridge connects the **Host Machine** to the bridge with a virtual network interface. This will be changed to truly isolate the **Host Machine** in the userspace network stack implementation to come.

In this alpha version, the older (less secure) method of operating three virtual networks (`virbrX`) will be used. 


### Setting up the virtual networks
Multiverse OS makes use of three primary virtual networks: "universe" (`virbr0` with subnet `10.0.0.0/24`), "galaxy" (`virbr1` with subnet `10.1.1.0/24`), and "star" (`virbr2` with subnet `10.2.2.0/24`).  

![alt=Multiverse OS Network Diagram, Universe Router connects to LAN and virbr1 through Physical Network Interface, Galaxy Router connects to virbr0 and virbr1, Star router connects to virbr1 and virbr2](https://github.com/hackwave/multiverse-development/blob/master/documentation/resources/images/multiverse-network-graph.png)

Since we are using virtual network bridges, the **Host Machine** must have a network device on the bridge of each network. The **Host Machine** should always have an ip address ending in `254` and the mac address should end in `fe` (254 in hex).

#### Universe Virtual Network
Find the `virbr0.xml` file in `./machines/host.multiverse/xml/networks' or copy the following to a file named `virbr0.xml`:

```
<network>
  <name>virbr0</name>
  <bridge name='virbr0' stp='on' delay='0'/>
  <mac address='00:00:0a:00:00:fe'/>
  <domain name='virbr0'/>
  <ip address='10.0.0.254' netmask='255.255.255.0'>
  </ip>
</network>
```

After saving it, it can be used with the `virsh net-define` command:

```
virsh net-define "virbr0.xml"
virsh net-autostart virbr0 
virsh net-start virbr0
```

#### Galaxy Virtual Network
Repeat the steps above with the `virbr1.xml` file:

```
<network>
  <name>virbr1</name>
  <bridge name='virbr1' stp='on' delay='0'/>
  <mac address='00:00:0a:01:01:fe'/>
  <domain name='virbr1'/>
  <ip address='10.1.1.254' netmask='255.255.255.0'>
  </ip>
</network>
```

Then define the network using the following:

```
virsh net-define "virbr1.xml"
virsh net-autostart virbr1
virsh net-start virbr1
```

#### Star Virtual Network
Repeat the previous for `virbr2.xml`

```
<network>
  <name>virbr2</name>
  <bridge name='virbr2' stp='on' delay='0'/>
  <mac address='00:00:0a:02:02:fe'/>
  <domain name='virbr2'/>
  <ip address='10.2.2.254' netmask='255.255.255.0'>
  </ip>
</network>
```

Then define the network using the following:

```
virsh net-define "virbr2.xml"
virsh net-autostart virbr2
virsh net-start virbr2
```

#### Allow userspace access to virtual network bridges
Now the newly created bridges need to be exposed to userspace by listing them in `/etc/qemu/bridge.conf`.


```
mkdir /etc/qemu/
echo -e "allow virbr0\nallow virbr1\nallow virbr2" > /etc/qemu/bridge.conf
# The -e flag allows for the interpretation of \n line breaks so that each allow
# {bridge_name} is on its own line.
```

Next the `qemu-bridge-helper` that uses the config file must have its permissions corrected so that members of the `libvirt` group can access it.

```
chown -R root:libvirt /usr/lib/qemu/
chmod 4750 /usr/lib/qemu/qemu-bridge-helper
```

**ERROR** The following error is caused by not having a properly formated or missing `bridge.conf` which exposes root bridges to userspace virtual machines.
> "libvirtError: internal error: /usr/lib/qemu/qemu-bridge-helper --use-vnet --br=virbr2 --fd=24: failed to communicate with bridge helper: Transport endpoint is not connected stderr=failed to parse default acl file `/etc/qemu/bridge.conf'"


### PCI Passthrough
Multiverse OS relies on the concept of PCI passthrough, where PCI devices are directly passed to the VM, and disabled on the **Host Machine**. This is supported on all modern devices, including ARM devices. 

#### VT BIOS Setting
The following script can determine if it is enabled in the BIOS:

```
# VT check
# Grep AMD and Intel from /proc/cpuinfo vmx for intel and svm for amd
echo "VT Check"
echo "========"
echo -e "\e[1m\e[21mChecking if required BIOS setting is enabled:"
echo -e "\e[1m\e[21mVirtualization technology (VT) enabled in BIOS."
if grep -q vmx /proc/cpuinfo; then

        current_cpu_vt="vmx"
        current_cpu_vt_enabled=0
	echo -e "Intel BIOS VM Setting:\e[37m.... [\e[32m Virtualization Technology (VT) Enabled \e[37m]"
elif grep -q svm /proc/cpuinfo; then
        current_cpu_vt="svm"
        current_cpu_vt_enabled=0
	echo -e "AMD BIOS VM Setting:  \e[37m.... [\e[32m Virtualization Technology (VT) Enabled \e[37m]"
else
        current_cpu_vt_enabled=1
fi
```

It may have to be enabled in your BIOS is if it is not already enabled by default. 


#### Enabling iommu
Modify grub by editing `/etc/default/grub` and adding `intel_iommu=on` or `amd_iommu=on` (depending on your chipset) to the `GRUB_CMDLINE_LINUX_DEFAULT` line.

If you are using AMD use the following:

```
GRUB_CMDLINE_LINUX_DEFAULT="quiet amd_iommu=on"
```
Or if you are using Intel use the following:

```
GRUB_CMDLINE_LINUX_DEFAULT="quiet intel_iommu=on"
```

The entire `/etc/default/grub` file is below. Uncomment the appropriate GRUB_CMDLINE_LINUX_DEFAULT line.

```
GRUB_DEFAULT=0                                   
GRUB_TIMEOUT=5                                   
GRUB_DISTRIBUTOR="Multiverse OS"                 
# For Intel:
#GRUB_CMDLINE_LINUX_DEFAULT="quiet intel_iommu=on"
# For AMD:
#GRUB_CMDLINE_LINUX_DEFAULT="quiet amd_iommu=on"  
GRUB_CMDLINE_LINUX=""                            
```

After modifying grub, update it and reboot the **Host Machine** for the changes to take effect:

```
update-grub
reboot
```

#### PCI Passthrough Kernel Modules: ['vfio-pci', 'vfio']
Next step in the process is enabling the `vfio-pci` kernel module.

You can check if this is loaded in your kernel using the following command:

```
lsmod | grep vfio
```

By default it will not be loaded and you will need to add `vfio-pci` and `vfio` to the `/etc/modules` file so that they are loaded at boot. 

```
echo "vfio-pci" >> /etc/modules
echo "vfio" >> /etc/modules
```

#### Unbinding NIC for passing to Universe Router VM
Then finally the NIC needs to prevented from loading into the **Host Machine** kernel so that it is free to be loaded by the **Router VM**. 

Before permanently disabling the NIC, it is best to test the process manually before adding it to settings.

Locate the PCI address using the following command:

```
lspci -nn
```

Locate the NIC that will be passed over to the router, for example:

```
05:00.0 Ethernet controller [0200]: Qualcomm Atheros QCA8171 Gigabit Ethernet [1969:10a1] (rev 10)
```

Two sets of numbers are important for the purposes of PCI Passthrough: (1) The PCI Address, which is found in the format of `xx:xx.x`. It is worth nothing, this number is not static and capable of changing in certain circumstances. For example, if a device is moved from one PCI slot to another in the motherboard the devices PCI Address will change to match the slot. The PCI Address will be necessary when unbinding and later when passing to the VM.


The other important number for PCI Passthrough is (2) the vendor ID, other number which is found at the end of the line in the form `[xxxx:xxxx]`. In contrast to how it is displayed when using `lspci -nn`, one must echo the ID in the format of `xxxx xxxx` to bind it to vfio-pci to make it accessible to unpriviledged users for PCI passthrough to VMs. The Vendor ID is static, and does not change regardless of what slot the PCI device is placed in, and is same across two devices of the same type. For example, two realtek network cards of the same make placed into a motherboard will more than likely have identical vendor IDs.  

Using the above Qualcom NIC card as an example the unbind command will look like:

```
echo "0000:05:00.0" > /sys/bus/pci/devices/0000\:05\:00.0/driver/unbind
```

#### Adding device to vfio-pci
After the device is unbound, it needs to be made available to `vfio-pci`. The command for the example device shown above would be:

```
echo "1969 10a1" > /sys/bus/pci/drivers/vfio-pci/new_id
```


#### After binding a device
After binding a device to `vfio`, it will create a iommu group file in `/dev/vfio/`, but by default the permissions will be restricted to root users. 

Changing permissions with the `/dev/` folder is done using a `udev` rules file in `/etc/udev/rules.d/10-vfio-permissions.rules`.

```
# Grant /dev/vfio ownership to root:libvirt
SUBSYSTEM=="vfio" OWNER="root" GROUP="libvirt" MODE="0660"
```


#### AMD vfio group accessibility
On AMD **Host Machines** (tested with the Opteron G5) the entire group will need to be unbound for any item of the group to be passed to a VM. 

After that you may or may not run into the issue that the device is still not accessible because the group is not assignable. If that is an issue, each device in the group will need to be unbound. First list all the devices in each group using the following command:

```
find /sys/kernel/iommu_groups/ -type l
```

The above command will output the following:

```
/sys/kernel/iommu_groups/10/devices/0000:00:15.2
/sys/kernel/iommu_groups/10/devices/0000:00:15.0
/sys/kernel/iommu_groups/10/devices/0000:05:00.0
/sys/kernel/iommu_groups/10/devices/0000:04:00.0
/sys/kernel/iommu_groups/10/devices/0000:00:15.3
```

Then unbind each device, for example:

```
echo "0000:00:15.3" > /sys/bus/pci/devices/0000\:00\:15.3/driver/unbind
echo "0000:00:15.2" > /sys/bus/pci/devices/0000\:00\:15.2/driver/unbind
echo "0000:00:15.0" > /sys/bus/pci/devices/0000\:00\:15.0/driver/unbind
echo "0000:04:00.0" > /sys/bus/pci/devices/0000\:04\:00.0/driver/unbind
```

Now the NIC card should be able to be passed to your universe **Router VM** when it is created below. 

### Automating device preparation at startup with rc.local
For now these can be put into your `/etc/rc.local` file to automate this process on boot. 

**NOTE** Debian 9 has deprecated `rc.local`, but until we can develop a better system `rc.local` is a fine solution. It's easy to reenable. Create the file `/etc/rc.local` and include the devices relevant for your case:

```
#!/bin/sh -e

# Unbind NIC
echo "0000:05:00.0" > /sys/bus/pci/devices/0000\:05\:00.0/driver/unbind
# Unbind all members of group 10 (group that contains the NIC)
echo "0000:00:15.3" > /sys/bus/pci/devices/0000\:00\:15.3/driver/unbind
echo "0000:00:15.2" > /sys/bus/pci/devices/0000\:00\:15.2/driver/unbind
echo "0000:00:15.0" > /sys/bus/pci/devices/0000\:00\:15.0/driver/unbind
echo "0000:04:00.0" > /sys/bus/pci/devices/0000\:04\:00.0/driver/unbind
# Add the NIC to vfio-pci to make it accessible to VMs
echo "1969 10a1" > /sys/bus/pci/drivers/vfio-pci/new_id

exit 0
```

Note that `/etc/rc.local` should begin with `#!/bin/sh -e`, and must exit with code 0. If any of the lines generate errors, it will fail to run.

After creating the file, you must `chmod +x /etc/rc.local`, then you can test it using `systemctl start rc.local`, then check `systemctl status rc.local` which will let you know if there was an error.



**DEVELOPMENT** The `rc.local` file will be replaced with a general Multiverse OS daemon that is configured with Ruby and does long running background jobs like backing up, updating, and other management, in addition to startup tasks. It will be integrated and its process managed by `systemctl`.

#### Unbinding by kernel module options
An alternative method for device unbinding is using a configuration file in `/etc/modprobe.d/` the following can be added to automatically unbind and bind devices to vfio:

```
# /etc/modprobe.d/multiverse.conf
options vfio-pci ids=...  # comma seperated ids in `xxxx:xxxx` format
```

However, in the past this has had selective success, for example working for a computer's GPU but failing to work for NIC cards. With some hardware, this method works for binding the device after blacklisting the device's module (see below).

#### Blacklisting devices
During the lockdown portion of the installation process many kernel modules wil be blacklisted on the **Host Machine** to limit the attack surface and simplify the functionality of the **Host Machine**. Instead of unbinding the devices in the `rc.local` file, all of the kernel modules for the devices in question can be determined using `lspci -k` and blacklisted. This will prevent them from being bound during boot, removing the requirement to unbind them later. This may prove to be the best solution in the long term.

**DEVELOPMENT NOTE:** Automation in the future will be handled by a Multiverse daemon/agent running on the **Host Machine** that will be managed by `systemctl`. It will notify the **Controller VM** of accessible devices, unbind/bind devices, perform general maintainence and more. 


______
## Building the first Router VM: Universe Router
The current design of Multiverse OS utilizes Alpine Linux to function as the **Router VMs**. 

**DEVELOPMENT** Need to use custom MAC addresses to conceal information about the **Router VMs**. Any MAC address that starts with `54:52...` will be a giveaway that it is a QEMU-based VM. This is why custom MAC addresses are important, at the very least for `universe.router.multiverse` because it is the router that is exposed to the internet. 


#### Router VM Configuration
The Universe router is the the router which has all of the physical network devices passed to it and is the gateway into the Multiverse cluster, and is first line of defense against attacks from the local LAN (or WAN attacks that penetrate into the LAN).

| **Router VM Setting Name**                       | **Setting Value**                           |
|--------------------------------------------------|---------------------------------------------|
| **Basic Details** Name                           | `universe.router.multiverse`                |
| **CPU Count**                                    | `1`                                         |
| **CPU Configuration**                            | `Copy host CPU configuration`               |
| **Memory** Current allocation                    | `256` (Can be as low as 128)|
| **Memory** Maximum allocation                    | `256` (Can be as low as 128)|
| **Boot Options** Autostart                       | `TRUE`                                      |
| **CDRom**                                        | After `alpine-setup` it should be removed   |
| **Sound Card**                                   | Delete sound card                           |
| **Virtual Disk** Disk bus                        | `VirtIO`                                    |
| **Virtual Disk** Storage format                  | `qcow2`                                     |
| **Virtual Disk** Cache mode                      | `none`                                      |
| **Virtual Disk** IO mode                         | `native`                                    |
| **Virtual Network Interface** Bridge name        | `virbr0`                                    |
| **Virtual Network Interface** Device model       | `virtio`                                    |
| **Virtual Network Interface** MAC Address        | `00:00:10:00:00:01`                         |
| **Add Hardware** PCI Host Device (PCI passthrough)   | Find by PCI address, repeat for all physical networking devices                        |
| **Add Hardware** Filesystem                          | Mount with **driver** `Path`, **mode** `Mapped` with `immediate` **write policy**, the **source path** is `~/multiverse-os/machines/universe.router.multiverse` and **target path** is `multiverse`                         |
| **Controller IDE**                                   | Delete IDE controller after install                       |
| **USB Redirector 1** | Delete this device |
| **USB Redirector 2** | Delete this device |

**DEVELOPMENT NOTE:** Eventually we want to completely abandon spice to control the VM in favor of a custom Multiverse OS protocol to make it more difficult for generic malware to target Multiverse OS. 

#### Configuring Alpine
Once the machine boots, type `root` in the login and you will not be prompted for a password. Initiate the Alpine installer with `setup-alpine`.

The installation process is a collection of questions, starting with locales and keyboard type. It is recommended to use widely used generic answers that are not unique regardless of if they are accurate. After the initial setup, you will not be interacting with the VM directly, so options such as keyboard mapping are unimportant.

Answer the first questions:

```
us
us
```

Next is the system hostname. All Multiverse OS machines use `host`, custom hostnames to make Multiverse cluster network navigation easier will be abstracted on top of the traditional system. 

```
host
```

The remaining `alpine-setup` options are (eth1 is set up first because eth0 is the virbr0 bridge NIC):

```
eth1
dhcp
eth0
none
yes
```

By selecting `yes`, the `/etc/network/interfaces` file will be opened in `vi` for you to manually modify. Input the following configuration:

```
auto lo
iface lo inet loopback

auto eth1
iface eth1 inet dhcp

auto eth0
iface eth0 inet static
	address 10.0.0.1
	netmask 255.255.255.0
	gateway 10.0.0.1
	up route del -net default gw 10.0.0.1 netmask 0.0.0.0
```

**DEVELOPMENT** A custom Multiverse configuration to automate these steps in the `alpine-setup` installation needs to be developed to simplify provisioning the **Router VMs**. `alpine-setup` allows for the dumping over an answers file that can be used to automate the answering of questions, if this does not work, directly modifying the `setup-alpine` script and supplying it in the shared storage could make complete setup of a Multiverse router configuration 

After saving, continue the `alpine-setup` process. Do not use a password, so that after confuring, it can be used as a base image that can be set up and have a password assigned later.

Do not use a `sshd` server, we will make modifications via `p9` shares, spice or a custom Multiverse OS protocol.

If it shows 'sda' at the disk setup step, it is because you forgot to switch IDE to VirtIO in the configuration. Shutdown the VM, change it and start the process again. You will need to go into Boot options and select IDE CDRom and move it before IDE Disk 1 in order to recontinue the install process.

```
(root password) empty
(root password verify) empty
UTC
none
f (detect and use fastest mirror)
none
chrony
vda
sys
y
```

**NOTE** Don't use `lvm` option, an lvm is less secure and should only be used when necessary.

After in the installation is complete, the notifications will tell you to reboot. Rather than rebooting, configure the shared storage by modifying `/etc/fstab` (below), then shutdown and remove the **CD Rom Drive** and the **IDE Controller** from the VM. 


#### Configuring Shared Storage
To simplify configuration, orchestration, and back up we set up shared storage using plan9. Configure `/etc/fstab`:

Delete the `/dev/cdrom` and `/dev/usbdisk` lines and add the following line to configure shared storage:

```
multiverse		/mnt/multiverse		9p		trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail		0	0
```

And make the mount folder: `mkdir /mnt/multiverse`

#### Install configuration files from shared storage
After the shared storage is successfully mounted, all the necessary configuration files are located in `/mnt/multiverse/config/`. 

There are two scripts currently used for provisioning **Router VMs**, the first is `packages.sh` which will install relevant packages for the specific **Router VM**. The second is `provision.sh` which creates symbol links from the shared storage configuration files to the VM system. 

```
# After editing `/etc/fstab` and rebooting the machine, use the newly mounted 
# shared storage scripts to provision the Router
cd /mnt/multiverse
./packages.sh
./provision.sh
```

Using the `/mnt/multiverse/provision.sh` script, the configuration files will be installed by creating symbolic links from the local VM filesystem to the files on the shared storage.

At this point, any config file modifications can be done either directly on the VM (install a text editor if needed: `apk add vim`, then uninstall when finished: `apk del vim`) or on the host machine.

Files that are likely to need editing are `/etc/shorewall/hosts` and `/etc/shorewall/rules`. See the manual setup instructions below for details on configuration files.

Finally, clear the logs, remove the bash history and poweroff:

```
rm -rf /var/log/*
rm .ash_history
poweroff
```

At the end of the configuration process, take a snapshot named "FreshInstall" and make a clone named `template.universe.router.multiverse`

Now continue on to [building the Galaxy **Router VM**](#building-galaxy-router-vm).

### Setting up a Multiverse OS router manually
Below steps are provided for manual configuration without the use of the shared storage configuration files found within this git repository.

#### Setting up dhcpd
One of the core functionality of a router is providing `dhcpd` for the network. Install the software by running `apk add dhcp`. 

Start by copying the `/etc/dhcp/dhcpd.conf.example` and editing it.

```
cp /etc/dhcp/dhcpd.conf.example /etc/dhcp/dhcpd.conf
```

The `/etc/dhcp/dhcpd.conf` file should includes both the subnet it is managing the DHCP leases of and any other subnet that is connected so that the DHCP system is aware. This file is also where static addresses are assigned. If you have prefered DNS servers, (for example, 8.8.8.8 or 9.9.9.9), they can be entered at the `domain-name-servers` option.

```
default-lease-time 9000;
max-lease-time 90000;

ddns-update-style interim;
authoritative;
log-facility local7;

subnet 10.0.0.0 netmask 255.255.255.0 {
  range 10.0.0.2 10.0.0.254;
  option domain-name-servers 85.214.20.141, 194.150.168.168;
  option domain-name "universe.router.multiverse";
  option routers 10.0.0.1;
  option broadcast-address 10.0.0.255;
}

subnet 192.168.0.0 netmask 255.255.255.0 {
  option routers 192.168.0.1;
  option broadcast-address 10.0.0.255;
}

host galaxy {
  hardware ethernet 00:00:01:00:00:02;
  fixed-address 10.0.0.2;
  option domain-name "galaxy.router.multiverse";
}

host voyager {
  hardware ethernet 00:00:01:00:00:10;
  fixed-address 10.0.0.10;
  option domain-name "universe.voyager.multiverse";
}
```

After the configuration is finalized, add the dhcp server to startup using the following command:

```
rc-update add dhcpd default
```


#### Configure the Shorewall router
Shorewall simplifies the management of router iptables for complex router configuration and firewall. 

**DEVELOPMENT NOTE:** Eventually it would be preferable to only use iptables, and use our own abstraction that is built around the needs of Multiverse OS routers. 

Start by configuring the zones, these provide a label for the various subnets known to the router. Zones support nesting so `lan` zone can be defined to be within the `wan`.

```
cat "wan		ipv4" >> /etc/shorewall/zones
cat "lan:wan		ipv4" >> /etc/shorewall/zones
cat "uni		ipv4" >> /etc/shorewall/zones
```

Configure `/etc/shorewall/hosts` and add the `eth1` device as `lan` zone:

!!!MISSING SOME INSTRUCTION LINES HERE!!!

```
cat "uni		eth0:10.0.0.0/24		-" >> /etc/shorewall/hosts
```

**NOTE** Don't just mindlessly copy-and-paste, ensure that `lan` zone matches your LAN subnet, which is typically `192.168.0.0/24` or `192.168.1.0/24`. 

Then add the devices to the shorewall interfaces `/etc/shorewall/interfaces`:

```
cat "wan		eth0		routefilter,tcpflags,logmartians,nosmurfs,sourceroute=0" >> /etc/shorewall/interfaces
cat "uni		eth1		dhcp,routefilter,tcpflags,logmartians,nosmurfs" >> /etc/shorewall/interfaces
```

The shorewall `masq` configuration file does not exist for some reason by default so we start by creating it:

```
touch /etc/shorewall/masq
cat "#
cat "#INTERFACE		SOURCE		ADDRESS		PROTO	PORT	IPSEC	MARK	USER	SWITCH" >> /etc/shorewall/masq
cat "eth1		0.0.0.0/0" >> /etc/shorewall/masq
```

To finalize the IP masquerading, `/etc/shorewall/snat` needs to be configured:

```
cat "MASQUERADE		0.0.0.0/0			eth0:!192.168.0.0/24
```


Next are the `/etc/shorewall/policy` and `/etc/shorewall/rules` configuration files, which are both very important files to understand because they are essentially the firewall configuration files. 

In essence, we want to allow internet traffic go from the universe network (and therefore galaxy and solar) out to the LAN and WAN networks but connections are not allowed to be made to internal networks. 

```
cat "fw		all		ACCEPT" >> /etc/shorewall/policy
cat "uni	lan		ACCEPT" >> /etc/shorewall/policy
cat "uni	wan		ACCEPT" >> /etc/shorewall/policy
cat "all	all		DROP" >> /etc/shorewall/policy
```

The next file `/etc/shorewall/rules` supplies specific firewall configuration, for example, we could have a long running service VM running a torrent client located within the Universe. We want to forward a port `51413` to this VM without exposing any other access to this VM, and not allow this VM to access the LAN.

In addition, if you wanted to expose a SAMBA share, you could do the same thing exposing only `139`, `445`, `137`, and `138`. The VM would not be able to access any other VM or LAN computers, and LAN computers can only interact with Universe **Router VM** on all ports except the 5 ports passed to the voyager service VM. 

The example cases in the file below should be replaced with any configuration necessary for your situation.

```
#
# Shorewall -- /etc/shorewall/rules
#
# For information on the settings in this file, type "man shorewall-rules"
#
# The manpage is also online at
# http://www.shorewall.net/manpages/shorewall-rules.html
#
##############################################################################################################################################################
#ACTION		SOURCE		DEST		PROTO	DPORT	SPORT	ORIGDEST	RATE	USER	MARK	CONNLIMIT	TIME	HEADERS	SWITCH	HELPER

?SECTION ALL
?SECTION ESTABLISHED
?SECTION RELATED
?SECTION INVALID
?SECTION UNTRACKED
?SECTION NEW
# Transmission RPC/WebUI
DNAT    lan     uni:10.1.1.100  tcp	9091
# Transmission Torrent Client
DNAT	lan	uni:10.1.1.100	tcp	51413
# Samba
DNAT	lan	uni:10.1.1.100	tcp	137
DNAT	lan	uni:10.1.1.100	tcp	138
DNAT	lan	uni:10.1.1.100	tcp	139
DNAT	lan	uni:10.1.1.100	tcp	445
```

Add `shorewall` firewall to startup:

```
rc-update add shorewall default
```

For it to actually work at startup `/etc/shorewall/shorewall.conf` needs to be modified:

```
# Line 12 needs to be modified to Yes
STARTUP_ENABLED=Yes
```

Don't forget to uninstall `vim` (or other editor) before saving the template.

```
apk del vim
```

Finally, clear the logs, remove the bash history and poweroff:

```
rm -rf /var/log/*
rm .ash_history
poweroff
```

At the end of the configuration process, take a snapshot named "FreshInstall" and make a clone named `template.universe.router.multiverse`

**DEVELOPMENT NOTE:** A `pass-store` needs to be initialized so once the template and first snapshot are made, a 32 character password can be generated and assigned to the production `universe.router.multiverse` VM. 

______
## Building Galaxy Router VM
The **Galaxy Router VM** is attached to the **Universe Router** and the **Star Router**, it does not have access to the LAN like the Universe network does and it transparently routes all traffic through a VPN. The **Galaxy Router VM** has two virtual network cards, the first attached to `virbr0` and the second to `virbr1`.

If you have this `git` repository cloned on the host machine, you can create the VM with `virsh define /home/user/multiverse-os/machines/galaxy.router.multiverse/xml/galaxy.router.multiverse.xml`. After creation, add the alpine image as a CD Rom and set it as the boot device and create the qcow2 virtual disk.

VM Setting details for manual setup are below:

| **Router VM Setting Name**                       | **Setting Value**                           |
|--------------------------------------------------|---------------------------------------------|
| **Basic Details** Name                           | `galaxy.router.multiverse`                  |
| **CPU Count**                                    | `1`                                         |
| **CPU Configuration**                            | `Copy host CPU configuration`               |
| **Memory** Current allocation                    | `256` (Can be less if necessary, low as 128)|
| **Memory** Maximum allocation                    | `256` (Can be less if necessary, low as 128)|
| **Boot Options** Autostart                       | `TRUE`                                      |
| **CDRom**                                        | After `alpine-setup` it should be removed   |
| **Sound Card**                                   | Delete sound card                           |
| **Virtual Disk** Disk bus                        | `VirtIO`                                    |
| **Virtual Disk** Storage format                  | `qcow2`                                     |
| **Virtual Disk** Cache mode                      | `none`                                      |
| **Virtual Disk** IO mode                         | `native`                                    |
| **Virtual Network Interface (1)** Bridge name    | `virbr0`                                    |
| **Virtual Network Interface (1)** Device model   | `virtio`                                    |
| **Virtual Network Interface (1)** MAC Address    | `00:00:10:00:00:02`                         |
| **Virtual Network Interface (2)** Bridge name    | `virbr1`                                    |
| **Virtual Network Interface (2)** Device model   | `virtio`                                    |
| **Virtual Network Interface (2)** MAC Address    | `00:00:10:01:01:01`                         |
| **Filesystem** (Add Hardware)                                  | Mount with **driver** `Path`, **mode** `Mapped` with `immediate` **write policy**, the **source path** is `~/multiverse-os/machines/galaxy.router.multiverse` and **target path** is `multiverse`                         |
| **Controller IDE**                                   | Delete IDE controller after install                       |


After saving, continue the `alpine-setup` process using the following options:

```
us
us
host
eth0
none
eth1
none
yes
```

And configure static ips for the network interfaces in `/etc/network/interfaces/`:

```
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet static
	address 10.0.0.2
	netmask 255.255.255.0
	gateway 10.0.0.1

auto eth1
iface eth1 inet static
	address 10.1.1.1
	netmask 255.255.255.0
	gateway 10.1.1.1
```

```
empty
empty
# Do not use a password, so after configuring, it can be used as a base image that can be setup then have a password assigned later
(root password) empty
(root password verify) empty
UTC
none
f (detect and use fastest mirror)
# Do not use a SSH server, we will make modifications via p9 shares, spice or a custom Multiverse OS protocol
none
chrony
# If it shows 'sda' it is because you forgot to switch IDE to VirtIO in the configuration, shutdown,
# change it and start the process again. You will need to go into Boot options and select IDE CDRom and uncheck IDE Disk 1 
# in order to recontinue the install process
vda
sys
y
```

**NOTE** Don't use `lvm` option, an lvm is less secure and should only be used when necessary.

After in the installation is complete, the notifications will tell you to reboot, instead shutdown and remove the **CD Rom Drive** and the **IDE Controller**. 


#### Configuring Shared Storage
To simplify configuration, orchestration, and back up we setup shared storage using plan9. Configure `/etc/fstab`:

Delete the `/dev/cdrom` and `/dev/usbdisk` lines and add the following line to configure shared storage:

```
cat "multiverse		/mnt/multiverse		9p		trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail		0	0" >> /etc/fstab
```

And make the mount folder: `mkdir /mnt/multiverse`

#### Configuration from shared storage
If you already have the configuration files in the shared storage folder, a `provision.sh` script is supplied to create symbolic links for configurations files to the shared storage.

```
##
## Galaxy Router Config Installer
##==========================================
# DHCPd, Shorewall and various other /etc/*
# configuration files need to be installed
# from the shared storage

echo -e "Multiverse OS: Galaxy Router Config Installer"
echo -e "============================================="
echo -e "Installing configuration files..."
echo -e ""
echo -e "dhcpd"
echo -e "Deleting existing dhcpd configuration files..."
rm -f /etc/dhcp/dhcp.conf
echo -e "Creating symbolic links from shared storage configuration files..."
ln -s /mnt/multiverse/config/etc/dhcp/dhcp.conf /etc/dhcp/
echo -e ""
echo -e "shorewall"
echo -e "Deleting existing shorewall configuration files..."
rm -f /etc/shorewall/hosts
rm -f /etc/shorewall/interfaces
rm -f /etc/shorewall/masq
rm -f /etc/shorewall/policy
rm -f /etc/shorewall/rules
rm -f /etc/shorewall/shorewall.conf
rm -f /etc/shorewall/snat
rm -f /etc/shorewall/zones
echo -e "Creating symbolic links from shared storage configuration files..."
ln -s /mnt/multiverse/config/etc/shorewall/hosts /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/interfaces /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/masq /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/policy /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/rules /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/shorewall.conf /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/snat /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/zones /etc/shorewall/
echo -e ""
echo -e "sysctl.d"
echo -e "Creating symbolic link for 05-multiverse.conf file"
rm -f /etc/sysctl.d/05-multiverse.conf
ln -s /mnt/multiverse/config/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/
echo -e ""
echo -e "other /etc/* files"
rm -f /etc/hosts
rm -f /etc/issue
rm -f /etc/motd
echo -e "Deleting existing /etc/* configuration files..."
ln -s /mnt/multiverse/config/etc/hosts /etc/
ln -s /mnt/multiverse/config/etc/issue /etc/
ln -s /mnt/multiverse/config/etc/motd /etc/

echo -e ""
echo -e "Configuration file installation completed!"
```

If the shared storage configuration files do not exist, copy the files from the `router.universe.multiverse`.


#### Post-installation configuration
After the files are migrated to the shared storage, the configuration files need to be modified and both `dhcpd` and `shorewall` need to be added to the boot process. 


```
apk update
apk add shorewall
apk add dhcp
apk add openvpn
```

#### Setting transparent OpenVPN proxy
The galaxy router has a different functional purpose from the Universe **Router VM**. The galaxy router provides transparent proxying for VPN services. 

Using your own private OpenVPN server running on VPS or a VPN server, obtain a `.ovpn` file. 

```
cp {.ovpn} /etc/openvpn/openvpn.conf
```

**DEVELOPMENT** Create a minimal API to easily upload `.ovpn` files and easily select locations. Eventually a scriptable filter system to route across several VPNs is the planned design.


Once `openvpn` is configured, it can be added to startup with the following command:

```
rc-update add openvpn default
```

At the end of the configuration process, take a snapshot named "FreshInstall" and clone the entire VM, duplicate the hard disk and name it `template.{VM Name}`.

______
## Building Star Router VM
The **Star Router VM** is attached to the internet through the **Galaxy Router**, it does not have access to the LAN like the Universe network does and it transparently routes all traffic through the Galaxy Router's vpn and `tor`. The **Star Router VM** has two virtual network cards, the first attached to `virbr1` and the second to `virbr2`.

There are two options for implementing the Star Router:
  (1) Download the `Whonix Gateway`, verify it and use it directly. With a small number of changes to the VM configuration, it works seamlessly as a drop-in solution for the Multiverse OS `star.router.multiverse`. Throughout development, this has been the route for several Multiverse OS developers. 
  (2) Use a Alpine based router provisioned to transparently route all traffic over Tor. 

There are advantages and draw-backs to each option, and detailed explanation for doing both is below. 


### Whonix Star Router
Throughout development the Whonix Gateway was used as the **Star Router VM**. It is a solid choice, well tested and serves the exact same function as the result of the Multiverse OS Star Router VM provisioning scripts: transparently routing all traffic over Tor. 

#### Obtaining Whonix Qcow2 Image, Verify Checksum, and Verify Signature
The guide for obtaining the developer GPG can be found on the [Whonix site](https://www.whonix.org/wiki/Whonix_Signing_Key), a script to download verify the checksums and signatures like with the Debian and Alpine ISO images above is included in the repository `images` folder and below:

**NOTE** The script below will also shrink the Whonix Gateway Router `*.qcow2` from `100GB` to `5GB`. 

```
#!/bin/sh

WHONIX_VERSION="13.0.0.1.4"

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"


echo -e $text"Whonix Gateway $WHONIX_VERSION"$reset
echo -e $text"========================="$reset
echo -e $text"Downloading and verifying Whonix Gateway Libvirt KVM (VM) image..."$reset

# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more sophisticated, for example it should scan the folder, find the highest version number so it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate

wget https://download.whonix.org/linux/$WHONIX_VERSION/Whonix-Gateway-$WHONIX_VERSION.libvirt.xz
wget https://download.whonix.org/linux/$WHONIX_VERSION/Whonix-Gateway-$WHONIX_VERSION.libvirt.xz.asc

wget https://download.whonix.org/linux/$WHONIX_VERSION/Whonix-Gateway-$WHONIX_VERSION.sha512sums
wget https://download.whonix.org/linux/$WHONIX_VERSION/Whonix-Gateway-$WHONIX_VERSION.sha512sums.asc

wget https://www.whonix.org/patrick.asc

echo -e $subheader"##  DOWNLOAD"$reset
echo -e $text"Successfully downloaded (1) the Libvirt KVM image, (2) the Libvirt KVM image signature,"$reset
echo -e $text"(2) checksum file, and (3) checksum signature, and (4) Whonix Developer"$reset
echo -e $text"key"$reset

echo -e $subheader"##  CHECKSUM"$reset
echo -e $text"Listing the checksums for each file downloaded..."$reset
cat Whonix-Gateway-$WHONIX_VERSION.sha512sums


# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying on the user to manually compare.
echo -e $text"Executing 'sha512sum' on each file downloaded..."$reset
sha512sum Whonix-Gateway-$WHONIX_VERSION.libvirt.xz

echo -e $text"Manually compare the values, and verify the checksums match..."$reset

echo -e $subheader"## SIGNATURE"$reset
echo -e $text"Import the Whonix Developer key with gpg --import..."$reset
gpg --import patrick.asc

echo -e $text"Verifying the signature file with the Alpine developer release key..."$reset
gpg --verify Whonix-Gateway-$WHONIX_VERSION.libvirt.xz.asc


echo -e $header"    **NOTE** This script is simple and does not actually do comparisons, it simplifies"
echo -e "    the process by automating the steps, it is up to you to actually compare"
echo -e "    the checksums and read the output of the gpg --verify command."$reset

echo -e $subheader"##  Extracting from *.xz archive"$reset
tar -xvf Whonix-Gateway-$WHONIX_VERSION.libvirt.xz

echo -e $text"Remvoing *.xz archive to save space"$reset
rm Whonix-Gateway-$WHONIX_VERSION.libvirt.xz
 
echo -e $subheader"##  Shrink Whonix Gateway"$reset
echo -e $text"Shrinking *.qcow2 file to ~5 GB from 100 GB"$reset
qemu-img convert -O qcow2 Whonix-Gateway-13.0.0.1.4.qcow2 Whonix-Gateway-13.0.0.1.4.smaller.qcow2 
rm Whonix-Gateway-13.0.0.1.4.qcow2
mv Whonix-Gateway-13.0.0.1.4.smaller.qcow2 Whonix-Gateway-13.0.0.1.4.qcow2


echo -e $success"\nComplete!"$reset
```

The `Whonix-Gateway-13.0.0.1.4.libvirt.xz` achive contains a `*.qcow` file that can be used in combination with an included `Whonix-Gateway-13.0.0.1.4.xml` file to quickly deploy the Whonix Gateway. 

```
virsh define Whonix-Gateway-13.0.0.1.4.xml
```

Using the `define` command will add the `*.xml` file, then using the below `edit` command will allow one to change the storage device path to `/var/multiverse/images/Whonix-Gateway-13.0.0.1.4.qcow2` and adding two network devices 1 connected to `virbr1` and the other `virbr2`. Alternatively, one can start using the configuration in the normal `star.router.multiverse` process, and simply use the `Whonix-Gateway-13.0.0.1.4.qcow2` image and it will work without any further configuration.

```
virsh edit Whonix-Gateway-13.0.0.1.4.xml
```


##### Downsides to using Whonix Gateway as Star Router 
One of the major downsides is that this model, along with many other container solutions and so on, require users to trust the developers and creators of the image files. Multiverse OS design always prefers to opt for a trustless model where all images are created locally from simple configuration/provisioning files for easy auditing, the eventual implementation of a android APK style permission system for images, and code analysis tools that look for obvious attempts at malware-like behavior.

Multiverse OS VM manager is called `portalgun` and it will build all images locally to remove the need to trust developers, and provide better security for general use computing.  

In addition, as development continues on Multiverse OS **Router VMs** advanced features, and simplified control from the **Controller VM** desktop environment will simplify interaction with the Star Router, giving fine grain control over the transparent Tor proxying. 


### Multiverse Star Router

| **Router VM Setting Name**                       | **Setting Value**                           |
|--------------------------------------------------|---------------------------------------------|
| **Basic Details** Name                           | `star.router.multiverse`                    |
| **CPU Count**                                    | `1`                                         |
| **CPU Configuration**                            | `Copy host CPU configuration`               |
| **Memory** Current allocation                    | `256` (Can be less if necessary, low as 128 MiB. Whonix version require 524 MiB)|
| **Memory** Maximum allocation                    | `256` (Can be less if necessary, low as 128 MiB. Whonix version require 524 MiB)|
| **Boot Options** Autostart                       | `TRUE`                                      |
| **CDRom**                                        | After `alpine-setup` it should be removed   |
| **Sound Card**                                   | Delete sound card                           |
| **Virtual Disk** Disk bus                        | `VirtIO`                                    |
| **Virtual Disk** Storage format                  | `qcow2`                                     |
| **Virtual Disk** Cache mode                      | `none`                                      |
| **Virtual Disk** IO mode                         | `native`                                    |
| **Virtual Network Interface (1)** Bridge name    | `virbr1`                                    |
| **Virtual Network Interface (1)** Device model   | `virtio`                                    |
| **Virtual Network Interface (1)** MAC Address    | `00:00:10:01:01:02`                         |
| **Virtual Network Interface (2)** Bridge name    | `virbr2`                                    |
| **Virtual Network Interface (2)** Device model   | `virtio`                                    |
| **Virtual Network Interface (2)** MAC Address    | `00:00:10:02:02:01`                         |
| **Filesystem** (Add Hardware)                                  | Mount with **driver** `Path`, **mode** `Mapped` with `immediate` **write policy**, the **source path** is `~/multiverse-os/machines/star.router.multiverse` and **target path** is `multiverse`                         |
| **Controller IDE**                                   | Delete IDE controller after install                       |


#### Alpine setup steps

After saving, continue the `alpine-setup` process using the following options:

```
# Do not use a password, so the after confuring, it can be used as a base image, that can be setup then have a password assigned later
(root password) empty
(root password verify) empty
UTC
none
f (detect and use fastest mirror)
# Do not use a SSH server, we will make modifications via p9 shares, spice or a custom Multiverse OS protocol
none
chrony
# If it shows 'sda' it is because you forgot to switch IDE to VirtIO in the configuration, shutdown,
# change it and start the process again. You will need to go into Boot options and select IDE CDRom and uncheck IDE Disk 1 
# in order to recontinue the install process
vda
sys
y
```

**NOTE** Don't use `lvm` option, an `lvm` in combination with VMs potentially increase the attack surface.

After in the installation is complete, the notifications will tell you to reboot, instead shutdown and remove the **CD Rom Drive** and the **IDE Controller**. 

Start by configuring the network interfaces in `/etc/network/interfaces/`:

```
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet static
	address 10.1.1.2
	netmask 255.255.255.0
	gateway 10.1.1.1

auto eth1
iface eth1 inet static
	address 10.2.2.1
	netmask 255.255.255.0
	gateway 10.2.2.1
```

#### Configuring Shared Storage
To simplify configuration, orchestration, and back up we setup shared storage using plan9. Configure `/etc/fstab`:

Delete the `/dev/cdrom` and `/dev/usbdisk` lines and add the following line to configure shared storage:

```
cat "multiverse		/mnt/multiverse		9p		trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail		0	0" >> /etc/fstab
```

And make the mount folder: `mkdir /mnt/multiverse`

#### Configuration files already exist in shared storage
If you already have the configuration files in the shared storage folder, a `/mnt/multiverse/provision.sh` script is supplied to create symbolic links for configurations files to the shared storage.

```
##
## Star Router Config Installer
##==========================================
# DHCPd, Shorewall and various other /etc/*
# configuration files need to be installed
# from the shared storage

echo -e "Multiverse OS: Star Router Config Installer"
echo -e "============================================="
echo -e "Installing configuration files..."
echo -e ""
echo -e "dhcpd"
echo -e "Deleting existing dhcpd configuration files..."
rm -f /etc/dhcp/dhcp.conf
echo -e "Creating symbolic links from shared storage configuration files..."
ln -s /mnt/multiverse/config/etc/dhcp/dhcp.conf /etc/dhcp/
echo -e ""
echo -e "shorewall"
echo -e "Deleting existing shorewall configuration files..."
rm -f /etc/shorewall/hosts
rm -f /etc/shorewall/interfaces
rm -f /etc/shorewall/masq
rm -f /etc/shorewall/policy
rm -f /etc/shorewall/rules
rm -f /etc/shorewall/shorewall.conf
rm -f /etc/shorewall/snat
rm -f /etc/shorewall/zones
echo -e "Creating symbolic links from shared storage configuration files..."
ln -s /mnt/multiverse/config/etc/shorewall/hosts /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/interfaces /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/masq /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/policy /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/rules /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/shorewall.conf /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/snat /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/zones /etc/shorewall/
echo -e ""
echo -e "sysctl.d"
echo -e "Creating symbolic link for 05-multiverse.conf file"
rm -f /etc/sysctl.d/05-multiverse.conf
ln -s /mnt/multiverse/config/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/
echo -e ""
echo -e "other /etc/* files"
rm -f /etc/hosts
rm -f /etc/issue
rm -f /etc/motd
echo -e "Deleting existing /etc/* configuration files..."
ln -s /mnt/multiverse/config/etc/hosts /etc/
ln -s /mnt/multiverse/config/etc/issue /etc/
ln -s /mnt/multiverse/config/etc/motd /etc/

echo -e ""
echo -e "Configuration file installation completed!"
```


#### Post-installation configuration
After the files are migrated to the shared storage, the configuration files need to be modified and both `dhcpd` and `shorewall` need to be added to the boot process. 

```
apk update
# Remember to delete vim after configuration is complete
apk add vim
apk add shorewall
apk add dhcp
```

#### Setting up transparent Tor proxy
The star router has a different functional purpose from the Universe **Router VM** and is a variant of the Galaxy **Router VM**. The star router provides transparent proxying for Tor services, similar to Whonix but much simpler design. 

Tor is not available in the main `apk` repository in Alpine but it is in the `edge` package repository.

To enable the repository needed to install Tor edit `/etc/apk/repositories`.

At the end of the configuration process, take a snapshot named "FreshInstall" and clone the entire VM, duplicate the hard disk and name it `template.{VM Name}`.


______
## Building the Controller VM
The entirety of the Multiverse OS user experience occurs within a virtual machine, these virtual machines are categorized within Multiverse OS as **Controller VMs**. The **Controller VM** has the GPU, USB and other devices passed to it. For security reasons, the USB devices are further PCI passed to nested **USB Proxy Service VMs** launched from within the **Controller VM** to protect the **Controller VM** from USB spreading malware.

When the Multiverse operating system boots, it automatically starts the three **Router VMs** and the default **Controller VM**. The **Host Machine** is never used by the user, it is compeletely inaccessible and the keyboard, mouse and monitor are virtually plugged into the **Controller VM** enabling the entire user experience to occur within. All resources not used by the three **Router VMs** are allocated to the **Controller VM**. A user can have multiple **Controller VMs** set up, but with normal hardware setups (i.e. only one grapics card), only one can be run at a time.

This enables the user to utilize save-state like functionality for their primary desktop environment, utilizing snapshots and loading from previous states, making a failed `apt-get dist upgrade` trivial to revert if desired. When tuned, the **Controller VM** is capable of near bare-metal preformance, capable of playing modern FPS without lag while having all the security of Whonix-like transparent routing and other security features provided by Multiverse OS. 

In addition, different **Controller VMs** can be configured and used for different purposes. For example, a user can isolate all gaming activity to a single dedicated **Controller VM**, or use an OSX-based **Controller VM** to isolate business activity, and so on.

#### Virtual Machine Configuration

. Select **Local install media (ISO image or CDROM)**, select **Debian**, provide it with the bare-metal machine total memory minus the memory allocated to **Router VMs** (`total - 768 MB` if using the default 256 MB for each router).

Assign all CPUs to this device, all but 1 will be pinned, and the last CPU will be shared between itself and the **Router VMs**.

Create a disk image. The size will depend on factors like whether you plan on having many **Controller VMs** or just one. For a simple setup with one **Controller VM**, give it as much of the hard drive as possible while leaving space for the host machine OS and the **Router VMs**.

| **Router VM Setting Name**                       | **Setting Value**                           |
|--------------------------------------------------|---------------------------------------------|
| **Basic Details** Name                           | `controller.multiverse`                |
| **CPU Count**                                    | `ALL`                                       |
| **CPU Configuration**                            | `Copy host CPU configuration`               |
| **Memory** Current allocation                    | `ALL - 768MB`                               |
| **Memory** Maximum allocation                    | `ALL - 768MB`                               |
| **Boot Options** Autostart                       | `TRUE`                                      |
| **CDRom**                                        | After installation should be removed        |
| **Virtual Disk** Disk bus                        | `VirtIO`                                    |
| **Virtual Disk** Storage format                  | `qcow2`                                     |
| **Virtual Disk** Storage Size                    | `ALL - 50GB`                                |
| **Virtual Disk** Cache mode                      | `none`                                      |
| **Virtual Disk** IO mode                         | `native`                                    |
| **Virtual Network Interface** Bridge name        | `virbr2`                                    |
| **Virtual Network Interface** Device model       | `virtio`                                    |
| **Virtual Network Interface** MAC Address        | `00:00:10:02:02:10`                         |
| **PCI Passthrough** (Add Hardware, PCI Host Device)                              | GPU, USB, etc                      |
| **Controller IDE**                                   | Delete IDE controller after install                       |

**NOTE** Do not use different values in the Current and Maximum memory, as this shifting memory size can create an attack surface.

Configure device before launching, remove USB redirects, remove all excess, and under CPU select **Copy host CPU configuration**.

To enable GPU passthrough, the VM Firmware (in the "Overview" tab in virt-manager) should be UEFI x86_64:/usr/share/OVMF/OVMF_CODE/fd. The corresponding XML is the `loader` section to the `os` block below:

```
  <os>
    <type ...your VM's info...>...</type>
    <loader readonly='yes' type='pflash'>/usr/share/OVMF/OVMF_CODE.fd</loader>
  </os>
```

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

**DEVELOPMENT** Develop a script/Go application that lists the available hard disks (subtracting the root disk to prevent installation/configuration failure) that are available for passthrough. The user should be able to select multiple items in the list. Then each item selected should by default passed to each Controller VM, but like the PCI devices can have a specific VM specified. The output configuration/script should similar to the PCI devices one to have continuity between different passthrough configurations. It needs to both (while we use libvirt) modify the XML, and create udev rules files that adds the `libvirt` (eventually `kvm` or even `multiverse` group). 

When the VM boots, it should say "Debian GNU/Linux UEFI Installer menu"

Install as normal. As of the Debian 9.4.0 install ISO, the installer doesn't correctly install the grub bootloader correctly the first attempt at installation. If this is the case, after removing the CDROM, the computer will boot to "UEFI Interactive Shell". Reattach the CDROM to the VM, enable it as first boot device. When it boots, select "Advanced options ..." from the installer menu, and go into "Rescue mode". Tap Enter to select the default options until the option to select the device to use as a root file system. Select `/dev/vda2`. Answer "Yes" to mount separate /boot/efi partition. At the next screen, select "Reinstall GRUB boot loader". Install GRUB to `/dev/vda`. After GRUB installation, turn off the VM and remove the CDROM and Controller IDE.

VMs cloned from this VM are likely to also need the GRUB reinstall step to boot, even if the parent machine boots properly.

**NOTE** Ideal **Controller VM** size has not been determined, ideally it should be possible to operate with a small VM size, with the majority of its space provided by networked/clustered space added later. The original space needs to satisfy the needs of the **Controller VM** operating system but also leave enough for the **Host Machine** (bare-metal) operating system. Either of these running out of space is painful to resolve and must be avoided with intelligent scripting systems that avoid it from happening. 

**DEVELOPMENT** Intelligent system to avoid **Host Machine** and **Controller VM** from ever running out of space. A fallback **Controller VM** that does not use complex drivers, that can be launched if other **Controller VMs** fail to launch in order to diagnose and solve the problem from a simplified **Controller VM** environment. 

For mounting a 9p folder without erroring at the startup, modify `/etc/initramfs-tools/modules` and add the following modules:

```
9p
9pnet
9pnet_virtio
```

Finally it requires an initramfs update:

```
sudo update-initramfs -u
```

At the end of the configuration process, take a snapshot named "FreshInstall" and clone the entire VM, duplicate the hard disk and name it `template.{VM Name}`.

______
## Building Voyager VM
The Voyager VM is the first example of a **Service VM**, a Multiverse OS VM type that isolates long-running background service. In the case of Voyager VM, it isolates and encapsualtes the long-running task of LAN/WAN file-sharing. This allows Multiverse OS to be secure while still enabling important common general user task like torrenting. 

Voyager VM torrenting VM can run either inside the Galaxy VM to pass all torrent traffic transparently over a VPN or it can be run on the Universe network to give it somewhat direct internet access. 

| **Router VM Setting Name**                       | **Setting Value**                            |
|--------------------------------------------------|----------------------------------------------|
| **Basic Details** Name                           | `universe.voyager.multiverse`                |
| **CPU Count**                                    | `1`                                          |
| **CPU Configuration**                            | `Copy host CPU configuration`                |
| **Memory** Current allocation                    | `1024` (Can be less if necessary, low as 128)|
| **Memory** Maximum allocation                    | `1024` (Can be less if necessary, low as 128)|
| **Boot Options** Autostart                       | `TRUE`                                       |
| **CDRom**                                        | After `alpine-setup` it should be removed    |
| **Sound Card**                                   | Delete sound card                            |
| **Virtual Disk** Disk bus                        | `VirtIO`                                     |
| **Virtual Disk** Storage format                  | `qcow2`                                      |
| **Virtual Disk** Cache mode                      | `none`                                       |
| **Virtual Disk** IO mode                         | `native`                                     |
| **Virtual Network Interface** Bridge name        | `virbr1`                                     |
| **Virtual Network Interface** Device model       | `virtio`                                     |
| **Virtual Network Interface** MAC Address        | `00:00:10:01:01:10`                          |
| **Filesystem** (Add Hardware)                                  | Mount with **driver** `Path`, **mode** `Mapped` with `immediate` **write policy**, the **source path** is `~/multiverse-os/machines/universe.voyager.multiverse` and **target path** is `multiverse`                         |
| **PCI Passthrough** (Add hardware, PCI Host Device)                              | All network cards                           |
| **Controller IDE**                                   | Delete IDE controller after install                       |

After saving, continue the `alpine-setup` process using the following options:

```
# Do not use a password, so the after confuring, it can be used as a base image, that can be setup then have a password assigned later
(root password) empty
(root password verify) empty
UTC
none
f (detect and use fastest mirror)
# Do not use a SSH server, we will make modifications via p9 shares, spice or a custom Multiverse OS protocol
none
chrony
# If it shows 'sda' it is because you forgot to switch IDE to VirtIO in the configuration, shutdown,
# change it and start the process again. You will need to go into Boot options and select IDE CDRom and uncheck IDE Disk 1 
# in order to recontinue the install process
vda
sys
y
```

**NOTE** Don't use `lvm` option, an lvm is less secure and should only be used when necessary.

After in the installation is complete, the notifications will tell you to reboot, instead shutdown and remove the **CD Rom Drive** and the **IDE Controller**. 

Start by configuring the network interfaces in `/etc/network/interfaces/`:

```
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet static
	address 10.1.1.10
	netmask 255.255.255.0
	gateway 10.0.0.1
```

#### Configuring Shared Storage
To simplify configuration, orchestration, and back up we setup shared storage using plan9. Configure `/etc/fstab`:

Delete the `/dev/cdrom` and `/dev/usbdisk` lines and add the following line to configure shared storage:

```
cat "multiverse		/mnt/multiverse		9p		trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail		0	0" >> /etc/fstab
```

And make the mount folder: `mkdir /mnt/multiverse`

#### Configuration files already exist in shared storage
If you already have the configuration files in the shared storage folder, a `/mnt/multiverse/provision.sh` script is supplied to create symbolic links for configurations files to the shared storage.

```
# Needs to be written
```

At the end of the configuration process, take a snapshot named "FreshInstall" and clone the entire VM, duplicate the hard disk and name it `template.{VM Name}`.

______
## Host Machine Lockdown
After the three **Router VMs**, the basic **Service VMs** are setup, and the **Controller VM** is setup, configured and tuned; it is time for the **Host Machine** to be locked down, and access to it cut-off from the user. 

The lockdown process involves uninstalling all unnecessary software, disabling of all unnecessary kernel modules, setuping up a firewall, locking down user accounts, and removing the root user.

### Uninstalling unnecessary software

#### Don't install recommends by default: depends vs recommends
With `apt-get` there is the concept of *depends* vs the concept *recommends*, and on the **Host Machine** we only want to install packages if they are absolutely necessary to operation.

By default, `apt-get` installs recommended packages by default, this is generally fine but for both the **Host Machine** and **Router VMs** specifically must not have any uncessary packages to limit the attack surface. 

```
echo 'APT::Install-Recommends "false";' > /etc/apt/apt.conf.d/05recommends
```

#### Removing uncessary folders
In addition to above, part of the process of preparing the environment is removing unnecessary folders and files:

```
rm -rf /home/user/Videos/
rm -rf /home/user/Desktop/
rm -rf /home/user/Templates/
rm -rf /home/user/Music/
rm -rf /home/user/Pictures/
rm -rf /home/user/Public/
rm -rf /home/user/Downloads/
rm -rf /home/user/Documents/
```

### Disabling unnecessary kernel modules
There are several ways of disabling kernel modules and understanding the different ways is important because they disable the kernel module at different times,  which may be important.

For example, if you are disabling the USB controllers, it is important to disable this AFTER the initramfs boot stage so the keyboard can be used to type in the LUKs password for the HD. However other devices can be disabled in the Grub kernel selection, so it is disabled immediately, like with the network device, so it can't be used by malware in the initramfs. 

```
# Remove the annoying speaker
modprobe -r pcspkr
echo "# Stop Annoying Sound" >> /etc/modprobe.d/blacklist.conf
echo "blacklist pcspkr" >> /etc/modprobe.d/blacklist.conf
```

Using `lspci -k` one can find out which kernel module is used by the various devices on the **Host Machine**. This can be a guide to how best to lockdown the **Host Machine**. 

#### Blacklisting with fake install
Another way to blacklist kernel modules is to use `fakeinstall`, after installing `fakeinstall` one can modify `/etc/modprobe.d/blacklist.conf`:

```
install <modulename> /bin/true
```

### Setting up firewall

### Limiting access for user accounts

### Removal of root account

______
## Multiverse OS glossary of terms
Multiverse OS is complex in detail, in use however, essentially all of this complexity is abstracted and the user experience is specifically designed so that a cluster, with layers of internal networks to completely isolate the Multiverse OS cluster from the LAN network. But the user experience is presented to the user in such a way that it both looks, feels and operates as if it is a single computer. This is done so that the average, or even novice user can benefit from the incredible security obtained from running Multiverse OS.  

* **Host Machine** the bare-metal host, the building block of a Multiverse OS cluster, this machine is _NEVER_ used by the user. It never directly accesses peripheral devices, USB devices, networking to protect it from viruses and other threats. 

* **Virtual Machine (VM)** full hardware virtualization using KVM and QEMU.

* **Controller VM** the VM that is used by the user, all activity occurs within controller VMs. It is ephemeral, functioning like a live CD.

* **Application VM** specialized ephemeral VMs that serve an application and relay the window the the Controller VM. 

* **Service VM** a specialized long-running VM that runs background tasks or services in an isolated environment and within an isolated network.

