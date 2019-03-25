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
NET0BR0="universe0"
NET0BR1="galaxy0"
NET0BR2="star-whonix"
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
	echo "$subheaderStarting NETWORK 0$reset"
	# Network 0
	echo -e "$textBooting up Universe0 and Universe1 routers and starting the corresponding networks...$reset"
	# Network 0
	virsh start universe0.router.portal
	# Network 1
	#virsh start universe1.router.portal
	sleep 10
	echo -e "$textBooting galaxy0 for Universe0 and galaxy0 for Universe1...$reset"
	#=================================================================================
	echo "$subheaderStarting NETWORK 1$reset"
	# Network 1
	virsh start galaxy0.router.universe0
	# Network 1
	#virsh start galaxy0.router.universe1
	sleep 5
	echo -e "$textBooting star0 for Universe0 and star0 for Universe1...$reset"
	# Network 0
	virsh start star0.whonix-router.universe0
	# Network 1
	#virsh start star0.whonix-router.universe1
	echo -e "$textAll routers booted, ready to start the user interface controller portal...$reset"
	echo -e "$textChecking status of routers and networks...$reset"
	echo -e "$green"
	list_vms
	echo -e "$reset"
	echo -e "$subheader======[^] ROUTER VMs Successfully Luanched [^]=====$reset"
	echo -e "Options for User Interface Controllers:"
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
