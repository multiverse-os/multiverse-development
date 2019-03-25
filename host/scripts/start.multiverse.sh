#!/bin/bash

#
# Multiverse OS
#
# A basic script to outline functionality of a basic tool
# for starting Multiverse OS cluster automatically. 
# 

echo -e "Multiverse OS"
echo -e "=============\n"
echo -e "Starting Multiverse Routers And Default Controller VM...\n"



echo -e "[ROUTER] Starting Universe router,..."
virsh start universe.router.multiverse
sleep 3
# TODO: Need a way to check if this is booted, or even a script 
# dropped into the server that writes to a file, or FIFO, or
# something similar with the the status, or a status code. 


echo -e "[ROUTER] Starting Galaxy router, using openvpn service AirVPN..."
virsh start galaxy.router.multiverse
sleep 4
# TODO: Need a way to check the status of the server, and confirm the
# internet works before and after connecting to the AiirVPN service. 


echo -e "[ROUTER] Starting Star router (Whonix), using tor onion routing..."
virsh start star.whonix-router.multiverse
# TODO: Confirm succesful start, get status, get info on tor connection. 
# like country, ip address, speed, if its a known tor connection.

# Improve the tor connection features to add simple rules that will
# eventually be scriptable. So one can requset only tor exit nodes
# coming out of the UK. Or above x speed, or ones that are not
# yet known to be an active tor exit node.


echo -e "[CONTROLLER] Starting controller.gravity.multiverse, the default controller"S
virsh start controller.gravity.multiverse
# TODO: Need to confirm at least the galaxy connection connects to the internet

# TODO: Need to confirm that controller boots successfully, if not, it should prompt
# user for retry or fallback to a simplified controller that can be used to
# handle repairs or select a different 


echo -e "\nWaiting for VMs to start..."
sleep 3 

virsh list


