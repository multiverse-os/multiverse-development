# General Setup Notes & Tasks

### Tasks

  * The Debian-Controller without a LUKs password on the primary drive can be used for gaming or various other things. It is a fresh setup. BUT it will need to have CPU pinning at the least for better preformance.


### HOST

#### Host Setup Notes
*All of the local machine data will be installed from the multiverse-development repository and this will be the basis for installation and setting up the machine in such a way that the local data ends up in both:*

`/etc/multiverse` and `/var/multiverse`

sudo mkdir /etc/multiverse
sudo chown user:user /etc/multiverse


sudo mkdir /var/multiverse
sudo chown user:user /var/multiverse
sudo mkdir -p /var/multiverse/images/os-images
sudo mkdir -p /var/multiverse/portals/disks
sudo mkdir -p /var/multiverse/portals/shares
# Disks and shares are temporary, really portals folder will be divided up by the virtual machine class type like universe-router and galaxy router; each will contain their own share folder, disk image, and other specific files for that machine type. 


### ROUTER


### CONTROLLER

