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
sudo mkdir -p /var/multiverse/images/os-images
sudo mkdir -p /var/multiverse/images/portal-disks
sudo mkdir -p /var/multiverse/portals/ # each portal type containing its own share socket, etc
sudo mkdir -p /var/multiverse/portal/shares
sudo mkdir -p /var/multiverse/portal/sockets
sudo mkdir -p /var/multiverse/portal/channels
sudo mkdir -p /etc/multiverse

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
