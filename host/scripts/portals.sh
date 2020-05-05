#!/bin/bash
# =============================================================================
# 
#   Multiverse OS Bash Framework Features
# =============================================================================
#   GLOBALs: Validation Globals
# -----------------------------------------------------------------------------
ALPHA="abcdefghijklmnopqrstuvqrstuvwxyz"
NUMERIC="0123456789"
ALPHANUMERIC=$ALPHA+$NUMERIC
## Multiverse OS Shell Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
warning="\e[93m"
fail="\e[91m"
reset="\e[0m"
green=$success
red=$fail

## OS Helpers
CURRENT_USER=$(whoami)

# -----------------------------------------------------------------------------
# CONFIGURATION
LIST_ALL=1
# Controllers
UBUNTU_PORTAL="ubuntu18-04.ui.portal"
DEBIAN_PORTAL="debian8.interface.portal"
# Network VMs
## TODO: These seem like they can be generated via loop over # of NIC cards
NET0BR0="net0br0"
NET0BR0_ROUTER="network0.bridge0.router"
NET0BR1="net0br1"
NET0BR1_ROUTER="network0.bridge1.router"
# Default Controller (Mount when nothing is specific)
PORTAL=""
DEFAULT_CONTROLLER=$UBUNTU_PORTAL
DEFAULT_ALIAS="default"
UBUNTU_ALIAS="ubuntu"
DEBIAN_ALIAS="debian"
FEDORA_ALIAS="fedora"
ALPINE_ALIAS="alpine"
# CONTROLLER ACTION
START="start"
STOP="stop"
REBOOT="reboot"
RESTART=$REBOOT
EDIT="edit"
DESTROY="destroy"

#==============================================================================
# Functions
# -----------------------------------------------------------------------------
list_vms() {
	echo -e "$success"
	sleep 3
	if [ $LIST_ALL ] ; then
		virsh list
	else
		virsh list --all
	fi
	sleep 2
	echo -e "$reset"
}

start_router_vms() {
	## 
	## TODO: Must iterate over loop of # of NICs setting up network 
	##       per network. 
	## 


	#======================================================================
	# Network 0
	echo "$subheaderStarting [ NETWORK 0 ]$reset"
	echo -e "$textBooting up [ NETWORK 0 ]: net0br0 and net0br1...$reset"
	virsh start network0.bridge0.router
	virsh start network0.bridge1.router
	sleep 10
	echo -e "$textCompleted booting [ NETWORK 0 ] bridges.$reset"
	#======================================================================
	# Network 1
	echo "$subheaderStarting [ NETWORK 1 ]$reset"
	echo -e "$textBooting up [ NETWORK 1 ]: net0br0 and net0br1...$reset"
	virsh start network1.bridge0.router
	virsh start network1.bridge1.router
	sleep 5
	echo -e "$textCompleted booting [ NETWORK 1 ] bridges.$reset"
	#======================================================================
	# Completed, Checking status...
	#======================================================================
	echo -e "$text[ Multiverse OS: Networking ]"
	echo -e "$textChecking status of routers and networks...$reset$green"
	list_vms
	echo -e "$reset$subheader=[^] VMs Successfully Luanched [^]=$reset"
	echo -e "Available Controllers, select one to boot:"
	echo -e "$successCONTROLLER VMs$reset"
	echo -e "$success"
	virsh list --all | grep controller
	echo -e "$reset"
}


default_controller_action() {
	2=$DEFAULT_CONTROLLER
	if [ -z $1 ] ; then
		echo -e "$fail[Error]$reset No action provided, defaulting to action: 'start'$reset"
		1=$START
	elif [ $1 = $START || $1 = $REBOOT || $1 = $RESTART || $1 = $STOP || $1 = $EDIT || $1 = $DESTROY ] ; then
		echo -e "$successRunning action $strong $1 $reset$success DEFAULT_CONTROLLER named $DEFAULT_CONTROLLER$reset"
	default:
		echo -e "$fail[Error]$reset Unkown action: $2 $reset"
		echo -e "$textSetting action to 'start' and continuing..."
		1=$START
	fi
	exec_action
}

exec_action() {
	if [ -z $2 || $2 = "default" ] ; then
		default_controller_action
	elif [ $2 = $UBUNTU_PORTAL || $2 = "ubuntu" ] ; then
		echo -e "$subheaderRunning action $1 on UBUNTU_PORTAL... $reset"
		PORTAL=$UBUNTU_PORTAL
	elif [ $2 = $DEBIAN_PORTAL || $2 = "debian" ] ; then
		echo -e "$subheaderRunning action $1 on DEBIAN_PORTAL... $reset"
		PORTAL=$DEBIAN_PORTAL
	elif [ $2 = $FEDORA_PORTAL || $2 = "fedora" ] ; then
		echo -e "$subheaderRunning action $1 on FEDORA_PORTAL... $reset"
		PORTAL=$FEDORA_PORTAL
	elif [ $2 = $ALPINE_PORTAL || $2 = "alpine" ] ; then
		echo -e "$successRunning action $1 on ALPINE_PORTAL... $reset"
		PORTAL=$ALPINE_PORTAL
	else
		echo -e "$successNo VM found with the name: $2 ... $reset"
		echo -e "$textSetting the value to the DEFAULT_CONTROLLER $DEFAULT_CONTROLLER and continuing...$reset"
		PORTAL=$DEFAULT_CONTROLLER
	fi

	# if we checked the values better, we could just do 'virsh $1 $2'
	if [ $2 = $START ] ; then
		virsh start $PORTAL
	elif [ $2 = $RESTART ] ; then
		virsh restart $PORTAL
	elif [ $2 = $STOP ] ; then
		virsh stop $PORTAL
	elif [ $2 = $DESTROY ] ; then
		virsh destroy $PORTAL
	elif [ $2 = $EDIT ] ; then
		virsh edit $PORTAL
	else
		virsh start $PORTAL
	fi
	list_vms
	echo -e "$textNothing$reset$accent started$reset, $text name $strong'$2'$reset$text is not found. $reset"
}


# CHECK IF ARG PASSED OR DEFAULT TO `virsh list --all`
if [ -z $1 ] ; then
	echo -e "$fail[Error]$reset $textMissing portals command parameter:$reset"
	echo -e "$fail  Example: 'portals vm  {SUB-COMMAND} {OPTIONAL_FILTER}'$reset"
	echo -e "$fail      (OR) 'portals dev {SUB-COMMAND} {OPTIONAL_FILTER}'$reset"
	echo -e "$fail      (OR) 'portals net {SUB-COMMAND} {OPTIONAL_FILTER}'$reset"
	echo -e "$failtext$strongDefaulting to list command:$reset"
	echo -e "$success"
	list_vms
	echo -e "$reset"

	# [DEV] TODO: Below code is not currenmtly working

	#exit 1
	# TODO: Here, instead of exiting we would want to be more intiutive, and
	just prompt the user for the required information using specialzied
	input to ensure data what is expected, santizied and then used.

	# For now we will just put in good defaults
	1="start"
	2="default"
	echo -e "$textSetting $strongARG1$reset$text to: $reset $1"
	echo -e "$textSetting $strongARG2$reset$text to: $reset $2"
fi
#==============================================================================
# Script
#==============================================================================
## Print Banner
##==============
echo -e "$subheaderStarting ROUTER VMs...$reset"
echo -e "$headerMultiverse OS: Portals VM Manager$reset"
echo -e "$subtle============================================$reset"
echo -e "$textInitializing Multiverse networking, starting $strongROUTER VMs$reset$text...$reset"

## START VMs
##===========
echo -e "$subheaderStarting $strongROUTER VMs$reset$text...$reset"
start_router_vms
echo -e  "$successSuccessfully started $strongNETWORK 0$reset$success routers: [universe, galaxy, star-whonix]$reset"


## START CONTROLLER VM
##=====================
# IF ALREADY STARTED: probably restart, probably promtp user before doing it else just start or prenset user with menu
if [ -z $2 ] ; then
	echo -e "$subheaderStarting $strongDEFAULT_CONTROLLER$reset$subheader named: $DEFAULT_CONTROLLER $reset" 
	virsh start $DEFAULT_CONTROLLER 
	list_vms
else
 	# list should be in its own function and called here not in above fucn
	exec_action
	list_vms
fi
