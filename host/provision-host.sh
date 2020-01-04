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
sudo mkdir -p /var/multiverse/images/os-images
sudo mkdir -p /var/multiverse/images/portal-disks
# each portal type containing its own share socket, etc. Just copy from git
sudo mkdir -p /var/multiverse/portal-gun/ 
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
