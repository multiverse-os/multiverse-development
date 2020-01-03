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

sudo mkdir -p /var/multiverse/images/os-images
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


chown -R root:libvirt /usr/lib/qemu/
chmod 4750 /usr/lib/qemu/qemu-bridge-helper



## Configurations
/etc/rc.local
/etc/motd
/etc/issue



/etc/qemu/bridge.conf


/etc/security/limits.conf


/etc/sysctl.conf

/etc/sysctl.d/{TWO FILES COPY IN HERE}
