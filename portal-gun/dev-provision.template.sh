#!/bin/sh
#========================================================
# [TODO] This is an incomplete stop-gap to intended to
# provide a frame for merging all the attempts at a 
# provision script for a given VM/portal.
#
# This script thematically matches the other development
# stage portal provision scripts to provide a common
# UI between these simple development stage scripts
# for the various components of the Multiverse OS system.


#########################################################
## Host Machine (bare-metal) Config Installer
##=======================================================
# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"
#========================================================
echo -e $header"Multiverse OS: Host Machine Package Installer"$reset
echo -e $accent"=================================================="$reset
echo -e $subheader"# Updating, installing and removing packages..."$reset
echo -e $text"Running 'apt-get update' to obtain package info from sources..."$reset
apt-get update
echo -e $text"Running 'apt-get upgrade' to update out of date packages..."$reset
apt-get upgrade
echo -e $success"Package installation completed."$reset


#========================================================
#echo -e $subheader"# systemctl Services"$reset
#echo -e $success"Service configuration completed."$reset
#============================================================
echo -e $header"Multiverse OS: Host Machine Config Installer"$reset
echo -e $accent"============================================="$reset
echo -e $text"Installing Multiverse OS configuration files..."$reset
echo -e $success"\nConfiguration file installation completed!"$reset

