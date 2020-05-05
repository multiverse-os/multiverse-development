#!/bin/sh
###############################################################################
# Multiverse OS Script Color Palette                                
###############################################################################
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"
###############################################################################
##
## TODO: This is only the frame of a basic Multiverse OS 
##       development stage provisioning file. It does not 
##       yet contain more logic than running 
## 
##         `apt-get update && apt-get upgrade` 
## 
##       Which post-install will have already been done by 
##       the netinstall. 
##
##       There is currently a provisioning file for the host
##       machine but it does not follow the development stage
##       provisioning style structure and so this was created
##       to migrate and merge the several attempts at a host
##       installation from a base of vanilla Debian 9 
##       net installation. Once completed, this can be 
##       used to build a very basic pre-alpha installation
##       disk using 'debootstrap' command. 
##
##       TL/DR: This is an incomplete stop-gap to intended to
##       provide a frame for merging all the attempts at a 
##       host installation script in a single shell script
##       for provisioning the host, that thematically matches
##       the other development stage portal provision scripts.
##
## TODO: Refer to the newly upgraded universe router provision 
##       script, which contains newly added code for two 
##       operating modes: standard which copies the base files 
##       directly to the system.
##
###############################################################################
## Host Machine (bare-metal) Config Installer                                ##
###############################################################################
echo -e $header"Multiverse OS: Host Machine Provisioning/Installer"$reset
echo -e $accent"=================================================="$reset
###############################################################################
##                                                                           ##
## TODO: Put in information here about operating mode; for example,          ##
##       standard vs development and other global information or             ##
##       prompt input/questions for the user to guide the                    ##
##       installation if a completed YAML configuration is not               ##
##       supplied.                                                           ##
##                                                                           ##
###############################################################################
echo -e $accent"------------------------------------------------------"$reset
echo -e $header"Provision: Host Machine Package Management"$reset
echo -e $accent"------------------------------------------------------"$reset
echo -e $subheader"  # Updating, installing and removing packages..."$reset
echo -e $text"    'apt-get update' package repository data..."$reset
apt-get update
echo -e $text"    'apt-get upgrade' out of date packages..."$reset
apt-get upgrade
echo -e $accent"-------------------------------------------------------"$reset
echo -e $header"Provision: Host Machine Config Installer"$reset
echo -e $accent"-------------------------------------------------------"$reset
echo -e $text"    Installing Multiverse OS configuration files..."$reset
echo -e $accent"-------------------------------------------------------"$reset
echo -e $header"Provision: Host Machine Services/Daemons"$reset
echo -e $accent"-------------------------------------------------------"$reset
echo -e $success"\n    Configuration file installation completed!"$reset



