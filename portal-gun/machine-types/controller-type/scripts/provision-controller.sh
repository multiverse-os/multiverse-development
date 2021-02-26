#
# Setup Script Notes
###############################################################################
# First installation notes for 2020

USER_HOME="/home/user"
MV_PATH="/var/multiverse"
GIT_SRC_PATH="$MV_PATH/development"
MV_CONFIG_PATH="/etc/multiverse"

## Packages
#
#  Autoremove automatically installed packages that are likely to be unnecssary
#  First mark these as manually installed to avoid autoremoval
#  Review which of these are definitely necessary
sudo apt install baobab fonts-cantarell gnome-disk-utility gnome-calculator \
	eog gnome-themes-extra libblockdev-crypto2 sudo
sudo apt autoremove -o=APT::AutoRemove::RecommendsImportant=0

sudo apt update
sudo apt upgrade -y -o=APT::Install-Recommends=0
sudo apt install -y -o=APT::Install-Recommends=0 ovmf qemu qemu-utils \
	qemu-system-common qemu-system-x86 qemu-user-static \
	libvirt0 libvirt-clients libvirt-daemon libvirt-daemon-system \
	virt-manager gir1.2-spiceclientglib-2.0 gir1.2-spiceclientgtk-3.0 \
	libfile-fcntllock-perl iproute2 dmidecode dnsmasq-base \
	pass git gnupg2 ssh ssh-askpass patch rsync dirmngr vim 
sudo apt remove -y nano

# !
#  packages included in previous version of script:
#  netcat-openbsd ebtables bridge-utils pkg-config tor tree sudo libvirt-dev  libxml2-utils libpam-cap


## Default Paths
# NOTE: Not a fan of os-images under images maybe os-installation or os-media or just os
# These /images will probably just end up in portal-gun or portals
sudo mkdir -p $MV_PATH/
# each portal type containing its own share socket, etc. Just copy from git
sudo mkdir -p $MV_PATH/portal-gun/os-images
# TODO: Not a fan of "shares" these are plan9 shares, we are already modifying
# a copy of p9 server and client for a custom disk type and so shares will not
# really adequately describe the new disk type. IT will provide essentailly a 
# temporary or long term disk shared across VMs of the same type for setup or
# continued operation
sudo mkdir -p $MV_PATH/portals/shares
sudo mkdir -p $MV_PATH/portals/disks
sudo mkdir -p $MV_PATH/portals/sockets/serial/
sudo mkdir -p $MV_PATH/portals/sockets/channel/
sudo mkdir -p $MV_PATH/portals/sockets/console/
sudo mkdir -p $MV_PATH/portals/sockets/parallel/
sudo chown -R user:kvm /var/multiverse
cd $MV_PATH
git clone https://github.com/multiverse-os/multiverse-development development
# After downloading, download the new os images, and install configuration files
cd $GIT_SRC_PATH && rm -rf sh && git clone https://github.com/multiverse-os/sh

# This is where we will store multiverse.conf or multiverse.yaml, and it will define a lot of the multiverse host configuration that will allow the user to change various multiverse settings
#sudo mkdir -p $MV_CONFIG_PATH

## This folder automatically gets created by libvirt, but in case it doesn't exist yet...
if [ ! -d "$USER_HOME/.local/share/libvirt" ]; then
	mkdir -p $USER_HOME/.local/share/libvirt
fi
## Make a symlink to the multiverse image folder because libvirt really wants its 'libvirt/images' folder
if [ -d ".local/share/libvirt/images" ]; then
	rmdir $USER_HOME/.local/share/libvirt/images
fi

ln -s $MV_PATH/images $USER_HOME/.local/share/libvirt/images
chmod 755 $USER_HOME/.local/share/libvirt

## User
cd /home/user && rmdir Desktop Downloads Documents Music Videos \
       	Pictures Public Templates

sudo usermod -a -G kvm user
sudo usermod -a -G libvirt user


git config --global user.email "you@example.com"
git config --global user.name "Your Name"


sudo cp $GIT_SRC_PATH/portal-gun/machine-types/controller-type/base-files/etc/apt/apt.conf /etc/apt
sudo cp $GIT_SRC_PATH/portal-gun/machine-types/controller-type/base-files/etc/apt/apt.conf.d/* /etc/apt/apt.conf.d/
# SH Framework
# Copy over vfio-bind into binary execution path 
sudo cp $GIT_SRC_PATH/sh/scripts/vfio-bind /usr/sbin/


## VM Setup (Usermode)
# NOTE: Would be better to move this to root:kvm and avoid needing libvirt group altogether

## Network Bridges
# NOTE: To be replaced with sockets
sudo chown -R root:kvm /usr/lib/qemu/
sudo chmod 4750 /usr/lib/qemu/qemu-bridge-helper

#$GIT_SRC_PATH/host/scripts/add-bridge.sh $GIT_SRC_PATH/host/xml/networks/net0br0.xml
#net0br1.xml  net0br2.xml  net1br0.xml  net1br1.xml  net1br2.xml


## Configurations
# NOTE: Track all changes needed for setting up Multiverse, this will simplify the process and all these can be kept in /etc/multiverse and symbolically linked. Then the rest of the /et/multiverse folder can be custom Multiverse OS config files which will most likely be ruby or YAML based.
#/etc/rc.local
#/etc/motd
#/etc/modules
#/etc/issue
#/etc/qemu/bridge.conf
sudo cp -a $GIT_SRC_PATH/portal-gun/machine-types/controller-type/base-files/etc/qemu/bridge.conf /etc/qemu
#/etc/security/limits.conf
sudo cp $GIT_SRC_PATH/portal-gun/machine-types/controller-type/base-files/etc/security/limits.conf /etc/security
#/etc/sysctl.conf
#/etc/sysctl.d/{TWO FILES COPY IN HERE}

echo 'ADDING 'virt-manager' CONNECTION TO HOST-MACHINE'
echo 'This is a hack until the Host Machine daemon/agent'
echo 'is completed and an API provides functionality to'
echo 'control Router VMs from the Controller VM.'

gsettings set org.virt-manager.virt-manager.connections uris "['qemu+ssh://user@10.1.1.254/session?socket=/run/user/1000/libvirt/libvirt-sock', 'qemu:///session', 'qemu:///system']"
