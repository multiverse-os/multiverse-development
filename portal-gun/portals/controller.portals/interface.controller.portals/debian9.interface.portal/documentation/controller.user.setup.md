
sudo apt-get update
sudo apt-get upgrade

# These are useful for Ubuntu or 
# other deeper uses of a high end
# GPU in Multiverse (CUDA, etc)
#sudo apt-get install libc6:i386
#sudo dpkg --add-architecture i386

sudo apt-get install vim virt-manager libvirt-dev ovmf git tor pass qemu- libvirt0 virt-manager libvirt-dev libvirt-clients cryptsetup

#[**make sure cryptab is installed**]

sudo touch /etc/apt/apt.conf.d/01multiverse

# TODO: THIS WONT WORK BECAUSE YOU CAN NOT ECHO INTO FILES OWNED BY ROOT
touch /etc/apt/apt.conf/01-recommends
echo 'APT::Install-Recommends "0";' > /etc/apt/apt.conf/01-multiverse

# DONT REMOVE SOFTWARE UNTIL RECOMMENDATIONS IS TURNED OFF!
# WITHOUT IT, REMOVING firefox-esr INSTALLS chormium automaticlaly :9
#sudo apt-get remove firefox-esr  iagno gnome-klotski hitori gnome-chess quadrapassel four-in-a-row gnome-robots gnome-mines tali gnome-taquin swell-foop gnome-tetravex  gnome-sudoku gnome-nibbles gnome-mahjongg lightsoff  five-or-more  chromium* 

sudo mkdir /etc/qemu
echo "allow virbr0
allow virbr1
allow virbr2
allow virbr4
allow virbr5
allow virbr6
allow virbr7" >> /etc/qemu/bridge.conf

# Need to add 'user' to libvirt group, probably kvm and qemu too 

sudo chown -R root:libvirt /usr/lib/qemu/
sudo chmod 4750 /usr/lib/qemu/qemu-bridge-helper


# (user folder stuff like this probably does not need to be apart of controller setup)
# libreoffice* gimp gimp-data inkscape gnome-maps transmission-*Rind 14
git config --global user.email "you@example.com"
git config --global user.name "Your Name"
rm -rf Music/ Public/ Videos/ Downloads/ Pictures/ Templates/ Documents/



# Here we need to add the line to fstab for Multiverse

echo "Multiverse /media/user/Multiverse 9p trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail 0 0" >> /etc/fstab


echo "vfio-pci
vfio
9p
9pnet
9pnet_virtio" >> /etc/modules


# lol this cant be secure, need to research this further
echo "@kvm		soft	memlock		-1
@kvm		hard	memlock		-1" >> /etc/security/limits.conf




# TODO: This is old, we have a new folder structure being implemetned and will liekly be handled by the daemon/agent
# Setup the new folder structure for Multiverse
su
mkdir /var/multiverse
cd /var/multiverse
ln -s /home/user/multiverse/images/ .
ln -s /home/user/multiverse/machines/ .
ln -s /home/user/multiverse/images/os-images/ .
ln -s /home/user/multiverse/machines/host.multiverse/scripts/ .


## Buster

# edit /etc/apt/sources.list and switch first 2 stretch entries to buster and non-free packages

echo "deb http://ftp.ch.debian.org/debian/ buster main non-free 
deb-src http://ftp.ch.debian.org/debian/ buster main non-free

deb http://security.debian.org/debian-security stretch/updates main non-free
deb-src http://security.debian.org/debian-security stretch/updates main non-free

# stretch-updates, previously known as 'volatile'
# should all be made buster even if not yet supported for ease
deb http://ftp.ch.debian.org/debian/ stretch-updates main non-free
deb-src http://ftp.ch.debian.org/debian/ stretch-updates main non-free
" > /etc/apt/sources.list

sudo apt-get update
sudo apt-get upgrade
sudo apt-get dist-upgrade


## NVIDIA GPU (see other docs for now)

## AMD GPU 
# Need to upgrade to buster for the newest vega to work

#sudo apt-get install amd64-microcode
sudo apt-get install firmware-amd-graphics
sudo apt-get install xserver-xorg-video-amdgpu* 



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

````
    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none'/>
      <source dev='/dev/disk/by-uuid/41f02abc-defa-4c21-b2eb-94750ccc4730'/>
      <target dev='vdb' bus='virtio'/>
    </disk>
````


````
    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none'/>
      <source dev='/dev/disk/by-uuid/[uuid]'/>
      <target dev='vd[X]' bus='virtio'/>
    </disk>
````


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

____________
## notes from some install sometime ago


sudo apt-get update
sudo apt-get upgrade

# These are useful for Ubuntu or 
# other deeper uses of a high end
# GPU in Multiverse (CUDA, etc)
#sudo apt-get install libc6:i386
#sudo dpkg --add-architecture i386

sudo apt-get install vim virt-manager libvirt-dev ovmf git tor pass qemu- libvirt0 virt-manager




sudo mkdir /etc/qemu
echo "allow virbr0
allow virbr1
allow virbr2
allow universe0
allow galaxy0
allow star0" >> /etc/qemu/bridge.conf

# Need to add 'user' to libvirt group, probably kvm and qemu too 

sudo chown -R root:libvirt /usr/lib/qemu/
sudo chmod 4750 /usr/lib/qemu/qemu-bridge-helper

# DONT REMOVE SOFTWARE UNTIL RECOMMENDATIONS IS TURNED OFF!
# WITHOUT IT, REMOVING firefox-esr INSTALLS chormium automaticlaly :9
#sudo apt-get remove firefox-esr  iagno gnome-klotski hitori gnome-chess quadrapassel four-in-a-row gnome-robots gnome-mines tali gnome-taquin swell-foop gnome-tetravex  gnome-sudoku gnome-nibbles gnome-mahjongg lightsoff  five-or-more  chromium* 

# libreoffice* gimp gimp-data inkscape gnome-maps transmission-*Rind 14
git config --global user.email "you@example.com"
git config --global user.name "Your Name"

rm -rf Music/ Public/ Videos/ Downloads/ Pictures/ Templates/ Documents/

# Here we need to add the line to fstab for Multiverse

echo "Multiverse /media/user/Multiverse 9p trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail 0 0" >> /etc/fstab


echo "vfio-pci
vfio
9p
9pnet
9pnet_virtio" >> /etc/modules



echo "@kvm		soft	memlock		-1
@kvm		hard	memlock		-1" >> /etc/security/limits.conf





# Setup the new folder structure for Multiverse
su
mkdir /var/multiverse
cd /var/multiverse
ln -s /home/user/multiverse/images/ .
ln -s /home/user/multiverse/machines/ .
ln -s /home/user/multiverse/images/os-images/ .
ln -s /home/user/multiverse/machines/host.multiverse/scripts/ .


## Buster

# edit /etc/apt/sources.list and switch first 2 stretch entries to buster and non-free packages

echo "deb http://ftp.ch.debian.org/debian/ buster main non-free 
deb-src http://ftp.ch.debian.org/debian/ buster main non-free

deb http://security.debian.org/debian-security stretch/updates main non-free
deb-src http://security.debian.org/debian-security stretch/updates main non-free

# stretch-updates, previously known as 'volatile'
deb http://ftp.ch.debian.org/debian/ stretch-updates main non-free
deb-src http://ftp.ch.debian.org/debian/ stretch-updates main non-free
" > /etc/apt/sources.list

sudo apt-get update
sudo apt-get upgrade
sudo apt-get dist-upgrade

#sudo apt-get install amd64-microcode
sudo apt-get install firmware-amd-graphics

## AMD GPU 
# Need to upgrade to buster for the newest vega to work


sudo apt-get install xserver-xorg-video-amdgpu* 



___________________
## Lockdown notes

##
##  Controller VM lockdown
##
## ========================================================================================================================


The controller VM is the isolated container for the total number of dataf/files belonging to the user. 

Lockdown of the Controller is similar to the Host lockdown, but it is not as severe. The similarities primarily involve
the fact that below do not directly access the internet, but instead rely on a VM ran from within your active cluster.

The process does involve deleting most of the software, consolidating/"processing" `~/` config files [`.cache`, `.local`, `.config`], 
into a more efficient, binary (i.e. CBOR) storage of all configuration.

**Legacy Support** Multiverse OS always opts to remain backwards compatible so that both new and old users of linux
can collaborate, learn and build together. This is done by the using the following the  "process" to assimilate config data=///

___________

# Scratch paper for setup of controller




   vim tor git pass virt-manager libvirt0 libvirt-dev libvirt-clients qemu-utils qemu libvirt-daemon




  * xserver-xorg-video-amdgpu

  * firmware-amd-graphics

  * amd-microcode64


  * Link multiverse folder, add it to fstab
cat "Multiverse /media/user/Multiverse         9p   trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail  0 0" >> /etc/fstab




  mkdir /etc/qemu
  chown root:libvirt /etc/qemu
  cat "allow virbr0" > /etc/qemu/bridge.conf
  cat "allow virbr1" >> /etc/qemu/bridge.conf
  cat "allow virbr2" >> /etc/qemu/bridge.conf
  chown root:libvirt /etc/qemu/
  chown root:libvirt /etc/qemu/bridge.conf




