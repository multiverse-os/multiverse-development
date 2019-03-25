#!/bin/sh
#==============================================================================
echo "Multiverse OS: Hardware isolated networking"
echo "============================================"
echo "Starting Multiverse networking..."
echo "Booting up Universe0 and Universe1 networks."
#==============================================================================
# UNIVERSE ROUTERS
#==============================================================================
echo "\nUNIVERSE ROUTER(s)"
echo "==================="
#------------------------------------------------------------------------------
# Network 0
#------------------------------------------------------------------------------
echo "Starting universe0.router [0|10.0.0.0/24]"
virsh start universe0.router
#------------------------------------------------------------------------------
# Network 1
#------------------------------------------------------------------------------
echo "Starting universe1.router [1|10.1.1.0/24]"
virsh start universe1.router
# Sleeping to wait for VM(s) boot up
# TODO: Should be replaced by an event loop tracking boot progress of VMs. 
sleep 4


#==============================================================================
# GALAXY ROUTERS
#==============================================================================
echo "\nGALAXY ROUTER(s)"
echo "================="
#------------------------------------------------------------------------------
# Network 0
#------------------------------------------------------------------------------
echo "Starting galaxy0.router [1|10.0.0.0/24]"
virsh start galaxy0.router
#------------------------------------------------------------------------------
# Network 1
#------------------------------------------------------------------------------
virsh start galaxy1.router
echo "Starting galaxy1.router [1|10 .1.1.0/24]"
sleep 4


#==============================================================================
# Completed
#==============================================================================
# TODO: Check if successfully started, have internet access
# and if so then possibly change sysfs type file to indicate
# controller is ready to boot. Or somerthing similar
echo "All UNIVERSE and GALAXY routers booted..."
echo "Ready to start the user interface controller portal..."
virsh list
echo "=============================================================================="
echo "# Multiverse OS CONTROLLER(s)"
echo "Available User Interface Controllers:"
virsh list --all | grep controller
