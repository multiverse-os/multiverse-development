#!/bin/sh

# ======================================
#                 __   
#      /\/\ \_\ \_ \ \ V E R S E 
#
# ======================================
#  # Global Variables
#  Variables relating to Multiverse OS
# --------------------------------------

MV_USER="user"
MV_GROUP="libvirt"
MV_CONFIG_PATH="/var/multiverse/"
USER_HOME="/home/user"

MV_PATH="$USER_HOME/multiverse/"
GIT_SRC_DIR="$USER_HOME/multiverse"
CURRENT_USER=$(whoami)
if [ $CURRENT_USER = "user" ]; then
  echo $fail"[Error] Must be logged in as root. Run 'su' and try again."$reset
  exit 0
fi
if grep -q Intel /proc/cpuinfo; then
  CPU_TYPE="Intel"
elif grep -q AMD /proc/cpuinfo; then
  CPU_TYPE="AMD"
else
  echo $fail"[Error] Failed to detect CPU manufacturer."$reset
  exit 0
fi

# ======================================
#
#   Multiverse OS Script Color Palette
# --------------------------------------
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
warning="\e[93m"
fail="\e[91m"
reset="\e[0m"
# =====================================

echo $header"Multiverse OS Host Machine Setup"$reset
echo $accent"================================"$reset
echo $text"# Installing and updating software packages"$reset
echo $text"# Setting up 'xpra' repository"$reset
wget -qO - https://xpra.org/gpg.asc | apt-key add -
cd /etc/apt/sources.list.d/
wget https://xpra.org/repos/buster/xpra.list
echo $strong"Updating apt packages"$reset
echo $text"Updating repositories"$reset
apt-get update
echo $text"Upgrading existing packages"$reset
apt-get -y upgrade
echo $text"Installing base Multiverse OS Controller VM packages: [ovmf, qemu, git, vim, sudo, virt-manager, pass, dirmngr, tor, xpra, python-pip, python-netifaces, python-cups]"$reset
apt-get install -y ovmf qemu git vim sudo virt-manager pass dirmngr tor python-netifaces python-pip python-cups

echo $strong"Installing 'python-pip' Packages"$reset
echo $text"Installing base Multiverse OS Controller VM python pip packages [numpy, pyinotify, opencv-python, pyopengl, pyopengl-accelerate]"$reset
pip install numpy
pip install pyinotify
pip install opencv-python
pip install pyopengl
pip install pyopengl-accelerate

echo $subheader"# Building /var/multiverse folder structure for machine configurations"$reset
mkdir -p /var/multiverse/images
mkdir -p /var/multiverse/machines
chown -R user:user /var/multiverse
chmod 711 /var/multiverse

## NOTE =======================================================
##  For the Controller VM we should just use the same files
##  from the host machine so changes are tracked across
##  machines. All images and sensitive files will be moved to
##  the host machine's '/var/multiverse' folder so this will
##  be viable. 

#cd $USER_HOME && git clone https://github.com/hackwave/multiverse-development multiverse
#chown -R user:user $GIT_SRC_DIR

## NOTE =======================================================
## 
##  These files should not be in '/var/multiverse' and instead
##  in the folder that is shared across host/controller shared
##  storage. These files are the ones that are generic and
##  stored in the general Multiverse OS development git.
##
##  The images themselves are not stored in the git but the
##  files will be already downloaded from the Host Machine
##  setup.
##
##
#cd $GIT_SRC_DIR/images/os-images && ./alpine-dl-and-verify.sh
#cd $GIT_SRC_DIR/images/os-images && ./debian-dl-and-verify.sh
#cd $GIT_SRC_DIR/images/os-images && ./whonix-dl-and-verify.sh
# =============================================================

echo $subheader"# Configuring Controller VM User Environment"$reset
echo $strong"Configuring git user to be generic for increased psuedoanonymity"$reset
git config --global user.email "you@example.com"
git config --global user.name "Your Name"

## NOTE In the future we should be adding to multiverse group, creating this group and moving away from libvirt
echo $strong"Adding generic user 'user' to groups: [kvm, libvirt] "$reset
usermod -a -G kvm user
usermod -a -G libvirt user

## NOTE =======================================================
##   
##  We should be setting up an alternate hard disk from the
##  original Host Machine OS install for either direct pass-
##  through or via folder passthrough. Preferably the first
##  option so that the USER encryption password is NOT stored
##  on the Host Machine. 
##
##  The Hard-disk that is passed either way is the 'mirror' 
##  drive in this instance and is used in such a way that
##  all the primary file storage. Since typically the OS
##  drive of our Controller and Host machine are smaller 
##  and use of symbolic links and secondary drive is used
##  to shift the bulk of the user files onto larger file-
##  systems
## 
# =============================================================

# Clearing out default directories
echo $strong"Deleting default user 'user' home directories: [Documents, Downloads, Music, Pictures, Public, Templates, Videos]"$reset
rm -rf $USER_HOME/Documents/
rm -rf $USER_HOME/Downloads/
rm -rf $USER_HOME/Music/
rm -rf $USER_HOME/Pictures/
rm -rf $USER_HOME/Public/
rm -rf $USER_HOME/Templates/
rm -rf $USER_HOME/Videos/


# Add 2 lines two /etc/security/limits.conf
echo $strong"Copy 'limits.conf' to '/etc/security/' to increase memory limits for unpriviledged users in 'kvm' group..."$reset
cp $GIT_SRC_DIR/machines/host.multiverse/config/etc/security/limits.conf /etc/security

echo $subheader"# Configuring VM environment"$reset
echo $strong"Deleting default libvirt virtual network"$reset
virsh net-undefine default
echo $strong"Creating three libvirt virtual networks for the Controller VM environment: [virbr0, virbr1, virbr2]"$reset
echo $text"These have different subnets from the Host Machine versions"$reset
echo $accent"virbr0: 10.10.10.0/24"$reset
echo $accent"virbr1: 10.11.11.0/24"$reset
echo $accent"virbr2: 10.12.12.0/24"$reset
virsh net-define $GIT_SRC_DIR/machines/controller.multiverse/xml/networks/virbr0.xml
virsh net-define $GIT_SRC_DIR/machines/controller.multiverse/xml/networks/virbr1.xml
virsh net-define $GIT_SRC_DIR/machines/controller.multiverse/xml/networks/virbr2.xml

echo $text"Setting autostart on each virtual network: [virbr0, virbr1, virbr2]"$reset
virsh net-autostart virbr0
virsh net-autostart virbr1
virsh net-autostart virbr2

echo $text"Creating '/etc/qemu'..."$reset
mkdir -p /etc/qemu

echo $text"Adding 'allow virbrX' to allow unpriviledged access to virtual bridges..."$reset
echo "allow virbr0" > /etc/qemu/bridge.conf
echo "allow virbr1" >> /etc/qemu/bridge.conf
echo "allow virbr2" >> /etc/qemu/bridge.conf

echo $text"Setting permissions on '/usr/lib/qemu/qemu-bridge-helper' for group 'libvirt'..."$reset
# NOTE In the future we want to move to multiverse group as we move away from libvirt
chown -R root:libvirt /usr/lib/qemu/
chmod 4750 /usr/lib/qemu/qemu-bridge-helper

# Add three storage pools to user session
# do I need -H here?
echo $strong"Configuring virtual machine storage pools"$reset
echo $warning"_/!\\_ WARNING: STORAGE POOLS CURRENTLY MUST BE MANUALLY CONFIGURED"$reset
#sudo -u user -H virsh pool-define $GIT_SRC_DIR/machines/host.multiverse/xml/storage
#sudo -u user virsh pool-define $GIT_SRC_DIR/machines/host.multiverse/xml/storage/images.xml
#sudo -u user virsh pool-define $GIT_SRC_DIR/machines/host.multiverse/xml/storage/os-images.xml
#sudo -u user virsh pool-define $GIT_SRC_DIR/machines/host.multiverse/xml/storage/machines.xml
#sudo -u user virsh pool-autostart images
#sudo -u user virsh pool-autostart os-images
#sudo -u user virsh pool-autostart machines

# NOTE Move to using Go and templates
echo $strong"Enabling IOMMU in grub..."$reset
if [ $CPU_TYPE = "Intel" ]; then
  cp $GIT_SRC_DIR/machines/controller.multiverse/config/etc/default/grub-intel /etc/default/grub
elif [ $CPU_TYPE = "AMD" ]; then
  cp $GIT_SRC_DIR/machines/controller.multiverse/config/etc/default/grub-amd /etc/default/grub
else
  echo $fail"[Error] Failed to detect CPU manufacturer."$reset
  exit 0
fi
echo $strong"Updating Grub..."$reset
update-grub

# NOTE Better to use template instead of this because it allows for rerunning this script wtihout iwssue
if [ $CPU_TYPE = "Intel" ]; then
  echo $strong"Adding modules to '/etc/modules': [kvm, kvm_intel, vfio, vfio_pci, vfio_iommu_type1, 9p, 9pnet, 9pnet_virtio]"$reset
  echo "kvm" > /etc/modules
  echo "kvm_intel" >> /etc/modules
elif [ $CPU_TYPE = "AMD" ]; then
  echo $strong"Adding modules to '/etc/modules': [kvm, kvm_amd, vfio, vfio_pci, vfio_iommu_type1, 9p, 9pnet, 9pnet_virtio]"$reset
  echo "kvm" > /etc/modules
  echo "kvm_amd" >> /etc/modules
else
  echo $fail"[Error] Failed to detect CPU manufacturer."$reset
  exit 0
fi

echo "vfio_pci" >> /etc/modules
echo "vfio" >> /etc/modules
echo "vfio_iommu_type1" >> /etc/modules
echo "9p" >> /etc/modules
echo "9pnet" >> /etc/modules
echo "9pnet_virtio" >> /etc/modules

# NOTE This should be changed from echo to template
echo $strong"Adding udev rules file..."$reset
echo 'SUBSYSTEM=="vfio" OWNER="root" GROUP="libvirt" MODE="0660"' > /etc/udev/rules.d/10-vfio-permissions.rules

echo $strong"[DEV] ADDING 'virt-manager' CONNECTION TO HOST-MACHINE"$reset
echo $text"[DEV] This is a hack until the Host Machine daemon/agent"$reset
echo $text"[DEV] is completed and an API provides functionality to"$reset
echo $text"[DEV] control Router VMs from the Controller VM."$reset

gsettings set org.virt-manager.virt-manager.connections uris "['qemu+ssh://user@10.1.1.254/session?socket=/run/user/1000/libvirt/libvirt-sock', 'qemu:///session', 'qemu:///system']"

echo $success"Installation Complete! Please reboot to activate iommu."$reset
# Unbind NIC using 1 of 3 methods:
#  (1) /etc/rc.local manual echo unbind 
#  (2) /etc/modprobe.d/multiverse.conf `blacklist {module-name}` obtained from lspci -k
#  (3) /etc/modprobe.d/multiverse.conf 'options vfio-pci ids=...

# Bind the device if option 2 or 3 used, manual bind with echo to vfio-pci driver

# If rc.local is created chmod +x /etc/rc.local

# Create the first Router VM


##
## Functionality required by all Controller VMs
##


# *) Setup Tor + SSH Server

# *) Setup Shiftsuit Key tree (onion key, gpg key, ssh key (both rsa and ecdsa), etc)

# *) Setup xpra server to control application VMs

# *) Setup first Application VMs 
#    * Firefox
#    * Generic Go Development
#    * Generic Ruby Development
#    * Generic Rust Development
#    * Generic C Development

# Replace comments with complex sed/single line editing bash script

## For gaming machines
##
#dpkg --add-architecture i386
#wget https://steamcdn-a.akamaihd.net/client/installer/steam.deb
#sudo apt-get install libgl1-mesa-dri:i386, libgl1-mesa-glx:i386, libc6:i386 git vim 
sudo apt-get update

# TODO: Look through other scripts to add the following
# modify apt sources to use buster
# dist-upgrade

apt-get install -y ovmf qemu git vim sudo virt-manager pass tor

sudo usermod -a -G libvirt user
sudo usermod -a -G kvm user


# TODO: Fix the download scripts and download the ISOs after pulling down the git repo
sudo git clone https://github.com/multiverse-os/multiverse-development /var/multiverse
chown -R user:user /var/multiverse



#echo "# Building /var/multiverse folder structure for machine configurations"
#mkdir -p /var/multiverse/images
#mkdir -p /var/multiverse/machines
#chown -R user:user /var/multiverse
#chmod 711 /var/multiverse



#cd $GIT_SRC_DIR/images/os-images && ./alpine-dl-and-verify.sh
#cd $GIT_SRC_DIR/images/os-images && ./debian-dl-and-verify.sh
#cd $GIT_SRC_DIR/images/os-images && ./whonix-dl-and-verify.sh


## TODO: Edit fstab to include p9 fs share for development on host via p9 share to controller
## Multiverse P9FS Passthrough
##=============================
#multiverse /media/user/Multiverse 9p trans=virtio,9p2000.L,rw,posixacl,cache=none,nofail 0 0

## TODO: For the plan9 share to work without root access and mount automatically at boot, 
## the following modules must be added to `/etc/modules` so that these kernel modules are
## loaded at boot time of the controller

# at boot time, one per line. Lines beginning with "#" are ignored.

Edit `/etc/modules`:
````
9p
9pnet
9pnet_virtio
````

sudo usermod -aG libvirt user # Phase this out asap
sudo usermod -aG kvm user



## TODO: Add default storage pools using `/var/multiverse/portalgun/*`; `os-images` under portalgun/images and `portals` for storage of vm disk images. Inside portals, storage pool may need to be added for 'controller-portals' and `app-portals' or possibly these can just be added AS portals. 


# TODO: Add bridges via script, change permissions on /var/lib/qemu/qemu-bridge-controller 
# Add bridges

# install vim golang ovmf virt-manager ... 
#
# Setup Script Notes
###############################################################################
# First installation notes for 2020


## Packages
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install ovmf qemu pass git dirmngr vim 
sudo apt-get remove nano minissdpd

## Default Paths
# NOTE: Not a fan of os-images under images maybe os-installation or os-media or just os
# These /images will probably just end up in portal-gun or portals
sudo mkdir -p /var/multiverse/
# each portal type containing its own share socket, etc. Just copy from git
sudo mkdir -p /var/multiverse/portal-gun/os-images
# TODO: Not a fan of "shares" these are plan9 shares, we are already modifying
# a copy of p9 server and client for a custom disk type and so shares will not
# really adequately describe the new disk type. IT will provide essentailly a 
# temporary or long term disk shared across VMs of the same type for setup or
# continued operation
sudo mkdir -p /var/multiverse/portals/shares
sudo mkdir -p /var/multiverse/portals/disks
sudo mkdir -p /var/multiverse/portals/sockets/serial/
sudo mkdir -p /var/multiverse/portals/sockets/channel/
sudo mkdir -p /var/multiverse/portals/sockets/console/
sudo mkdir -p /var/multiverse/portals/sockets/parallel/
# This is where we will store multiverse.conf or multiverse.yaml, and it will define a lot of the multiverse host configuration that will allow the user to change various multiverse settings
sudo mkdir -p /etc/multiverse

# Because for not default always gets created, so lets link it to our primary default
rm ~/.local/share/libvirt/images
ln -s /var/multiverse/portals/disks/ ~/.local/share/libvirt/images

## User
cd /home/user && rm -rf Desktop Downloads Documents Music Videos Pictures 

sudo usermod -a -G kvm user
sudo usermod -a -G libvirt user


git config --global user.email "you@example.com"
git config --global user.name "Your Name"


cd /home/user && git clone https://github.com/multiverse-os/multiverse-development multiverse
# After downloading, download the new os images, and install configuration files
cd /home/user/multiverse/ && rm -rf sh && git clone https://github.com/multiverse-os/sh

## SH Framework
# Copy over vfio-bind into binary execution path 



## VM Setup (Usermode)
# NOTE: Would be better to move this to root:kvm and avoid needing libvirt group altogether

chown -R root:kvm /usr/lib/qemu/
chmod 4750 /usr/lib/qemu/qemu-bridge-helper



## Configurations
# NOTE: Track all changes needed for setting up Multiverse, this will simplify the process and all these can be kept in /etc/multiverse and symbolically linked. Then the rest of the /et/multiverse folder can be custom Multiverse OS config files which will most likely be ruby or YAML based.
/etc/rc.local
/etc/motd
/etc/issue
/etc/qemu/bridge.conf
/etc/security/limits.conf
/etc/sysctl.conf
/etc/sysctl.d/{TWO FILES COPY IN HERE}
