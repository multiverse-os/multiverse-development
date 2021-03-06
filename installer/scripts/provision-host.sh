

#
# Setup Script Notes
###############################################################################
# First installation notes for 2020

USER_HOME="/home/user"
GIT_SRC_PATH="/var/multiverse"
MV_CONFIG_PATH="/etc/multiverse"

MV_HOST_PATH="/development/host"
MV_CONTROLLER_PATH="/development/controller"
MV_ROUTER_PATH="/development/controller"

## Packages
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install -y ovmf qemu qemu-system-common virt-manager pass git dirmngr vim 
sudo apt-get remove -y nano minissdpd

## Default Paths
# NOTE: Not a fan of os-images under images maybe os-installation or os-media or just os
# These /images will probably just end up in portal-gun or portals
sudo mkdir -p /var/multiverse/portal-gun/ 
sudo mkdir -p /var/multiverse/portal-gun/os-images
# each portal type containing its own share socket, etc. Just copy from git
# TODO: Not a fan of "shares" these are plan9 shares, we are already modifying
# a copy of p9 server and client for a custom disk type and so shares will not
# really adequately describe the new disk type. IT will provide essentailly a 
# temporary or long term disk shared across VMs of the same type for setup or
# continued operation
sudo mkdir -p /var/multiverse/portals/shares/controller
sudo mkdir -p /var/multiverse/portals/disks/controller

sudo mkdir -p /var/multiverse/portals/shares/router
sudo mkdir -p /var/multiverse/portals/disks/router

sudo mkdir -p /var/multiverse/portals/wormholes/serial/
sudo mkdir -p /var/multiverse/portals/wormholes/channel/
sudo mkdir -p /var/multiverse/portals/wormholes/console/
sudo mkdir -p /var/multiverse/portals/wormholes/parallel/
sudo chown -R user:kvm /var/multiverse
# This is where we will store multiverse.conf or multiverse.yaml, and it will define a lot of the multiverse host configuration that will allow the user to change various multiverse settings
sudo mkdir -p /etc/multiverse

# Because for not default always gets created, so lets link it to our primary default
rm -f ~/.local/share/libvirt/images
ln -s /var/multiverse/portals/disks/ ~/.local/share/libvirt/images

## User
rm -rf ~/Desktop 
rm -rf ~/Downloads 
rm -rf ~/Documents 
rm -rf ~/Music 
rm -rf ~/Videos 
rm -rf ~/Pictures 

sudo usermod -a -G kvm user
sudo usermod -a -G libvirt user


git config --global user.email "you@example.com"
git config --global user.name "Your Name"


#cd ~/ && pass init // After generate or load GPG key

#cd /home/user && git clone https://github.com/multiverse-os/multiverse-development multiverse
# After downloading, download the new os images, and install configuration files
#cd /home/user/multiverse/ && rm -rf sh && git clone https://github.com/multiverse-os/sh

## SH Framework
# Copy over vfio-bind into binary execution path 



## VM Setup (Usermode)
# NOTE: Would be better to move this to root:kvm and avoid needing libvirt group altogether

## Network Bridges
# NOTE: To be replaced with sockets
chown -R root:kvm /usr/lib/qemu/
chmod 4750 /usr/lib/qemu/qemu-bridge-helper

$GIT_SRC_PATH/host/scripts/add-bridge.sh $GIT_SRC_PATH/host/xml/networks/net0br0.xml
#net0br1.xml  net0br2.xml  net1br0.xml  net1br1.xml  net1br2.xml

## TODO: Automate router creation and controller creation

## Enable IOMMU in grub
## TODO: Fuck we did a lot of upgrades on this but forgot to save
## If using Intel procecssor, comment out grub-amd and uncomment grub-intel
#cp $GIT_SRC_PATH/development/base-files/etc/default/grub-amd /etc/default
cp $GIT_SRC_PATH/base-files/etc/default/grub-intel /etc/default
update-grub


## Configurations
# NOTE: Track all changes needed for setting up Multiverse, this will simplify the process and all these can be kept in /etc/multiverse and symbolically linked. Then the rest of the /et/multiverse folder can be custom Multiverse OS config files which will most likely be ruby or YAML based.
#/etc/rc.local
#/etc/motd
#/etc/modules
#/etc/issue
#/etc/qemu/bridge.conf
#/etc/security/limits.conf
cp $GIT_SRC_PATH/development/host/base-files/etc/security/limits.conf /etc/security
i
#/etc/sysctl.conf
#/etc/sysctl.d/{TWO FILES COPY IN HERE}

# Exit with no errors
exit 0
