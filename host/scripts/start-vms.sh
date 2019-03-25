
# Autostart on boot

## The long term solution is to have the Multiverse OS
## agent running on the host machine launch each VM
## checking if the router is connected before launching
## the next router to ensure the networking comes online
## correctly.
##
## In addition, if the default controller fails to load
## properly a very basic, minimal resources VM should
## launch to help the user select a different controller
## VM, or load a snapshot or do some other form of 
## recovery or maintenance.
##
##

# Autostart Method
#sudo -u user -H virsh list
#sleep 1
#sudo -u user -H virsh list

# Manual Start VMs Method
sudo -u user -H virsh start universe.router.multiverse 
sleep 8
sudo -u user -H virsh start galaxy.router.multiverse
sleep 16
sudo -u user -H virsh start star.whonix-router.multiverse
sleep 8
#sudo -u user -H virsh start debian9.controller.multiverse
#sudo -u user -H virsh start ubuntu17.controller.multiverse
