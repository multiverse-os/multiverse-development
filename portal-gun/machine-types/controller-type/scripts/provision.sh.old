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
