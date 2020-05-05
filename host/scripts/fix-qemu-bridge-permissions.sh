#!/bin/sh
###############################################################################

# This was from an incorrect version
#sudo chown root:libvirt /etc/qemu
#sudo chown root:libvirt /etc/qemu/bridge.conf 
#wsudo chmod 0640 /etc/qemu/bridge.conf 


chown -R root:libvirt /usr/lib/qemu/
chmod 4750 /usr/lib/qemu/qemu-bridge-helper












