#!/bin/sh
###############################################################################
## DEVELOPMENT NOTE ###########################################################
# For now, during development, we want to leave the files in our multiverse
# portal share and symbolically link from the plan9 share. This way we can
# easily make changes and updates from our controller by simply editing the 
# files under portal-gun and have our changes immediately reflected in the
# router portals without needing to copy files around and easily be able to
# push our changes to the development repo so that changes can be shared
# with other developers without any extra work. 
#
# As we do each router, the preferred style is to include the custom features
# using if statements to each provisioning script starting from universe.
###############################################################################
#= Multiverse OS Script Color Palette #=======================================#
# TODO: With how often we are using this as a throw in to jazz up our dev 
# scripts, we should definitely be opting for a method that calls in a 
# library that is insalled into each machines multiverse folder. Because right
# now any changes we want to make to this would be a giant headache and that
# method of doing this is bad practice. 
###############################################################################
# TODO: Need warning, info, and other standard log levels
header="\e[0;95m"
accent="\e[2m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"
###############################################################################
## Global Variables ###########################################################

router_type="bridge1"
router_instance="0" 
package_manager="apk" # Alpine Linux

multiverse_mount="/mnt/multiverse"
base_files="$multiverse_mount/base-files"

if [ $router_type = "bridge1" ]; then
	mkdir /etc/multiverse/vpn-configs
fi


## Alpine `apk` package based directories

## Packages to be installed for Galaxy Router:
##    [openvpn, shorewall, dhcp]
mkdir -p /etc/openvpn
mkdir -p /etc/shorewall
mkdir -p /etc/dhcp
###############################################################################
echo -e $header"Multiverse OS:$reset$text $router_type router provisioning"$reset
echo -e $accent"==============================================================================="$reset
#===============================================================================
echo -e $accent"------------------------------------------------"$reset
	echo -e $header"Provisoning$reset$accent Installing Packages"$reset
echo -e $accent"------------------------------------------------"$reset
echo -e $subheader"# Install $accent'apk'$reset$subheader packages for router portal: [$accent dhcp$reset$subheader,$accent shorewall$reset$subheader ]"$reset
echo -e $accent"\`apk update\`$reset$text: updating installed packages"$reset
apk update &>/dev/null
echo -e $accent"\`apk add dhcp\`$reset$text: installing $reset$accent dhcp"$reset
apk add dhcp &>/dev/null

echo -e $accent"\`apk add shorewall\`$reset$text: installing $reset$accent shorewall"$reset
apk add shorewall &>/dev/null

if [ $router_type = "bridge1" ]; then
	echo -e $accent"\`apk add openvpn\`$reset$text: installing $reset$accent openvpn"$reset
	apk add openvpn &>/dev/null
fi
echo -e $accent"-------------------------------------------------------------------------------"$reset
	echo -e $subheader"[Development]$reset$header Provisoning$reset$accent Configuration Installer"$reset
echo -e $accent"-------------------------------------------------------------------------------"$reset

echo -e $text"Installing configuration files..."$reset

#===============================================================================

echo -e $subheader"Configuring: [$text dhcpd$reset $subheader]"$reset
	cp $base_files/etc/dhcp/dhcpd.conf /etc/dhcp/    2>/dev/null
#===============================================================================
echo -e $subheader"Configuring: [$text shorewall$reset $subheader]"$reset
	cp $base_files/etc/shorewall/hosts          /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/interfaces     /etc/shorewall/  2>/dev/null   
	cp $base_files/etc/shorewall/masq           /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/policy         /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/rules          /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/shorewall.conf /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/snat           /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/zones          /etc/shorewall/  2>/dev/null
#===============================================================================
echo -e $subheader"Network Interfaces: [$text /etc/network/interfaces$reset $subheader]"$reset
echo -e $text"Removing interfaces file"$reset
rm -f /etc/network/interfaces 2>/dev/null
	cp $base_files/etc/network/interfaces /etc/network/    2>/dev/null
#===============================================================================

if [ $router_type = "bridge1" ]; then


	cp $base_files/etc/openvpn/up.sh /etc/openvpn/    2>/dev/null
	cp $base_files/etc/openvpn/down.sh /etc/openvpn/    2>/dev/null
# TODO: Move openvpn files over

fi
#==============================================================================
echo -e $subheader"System Configuration: [$text /etc/sysctl.d/multiverse.conf$reset $subheader]"$reset
echo -e $text"Creating symbolic link for 05-multiverse.conf file"$reset
rm -f /etc/sysctl.d/05-multiverse.conf 2>/dev/null
	cp $base_files/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/    2>/dev/null

#===============================================================================
echo -e $subheader"General Linux Configuration: [$text /etc/*$reset $subheader]"$reset
rm -f /etc/hosts 2>/dev/null
rm -f /etc/issue 2>/dev/null
rm -f /etc/motd  2>/dev/null
	cp $base_files/etc/hosts /etc/ 2>/dev/null
	cp $base_files/etc/issue /etc/ 2>/dev/null
	cp $base_files/etc/motd  /etc/ 2>/dev/null
#===============================================================================
echo -e $subheader"Adding router services to startup via$reset$text rc-update"$reset
echo -e $text"Adding [$reset$subheader shorewall$reset,$subheader dhcpd$reset,$subheader openvpn$reset$text ]"$reset
rc-update add shorewall default 2>/dev/null
rc-update add dhcpd default     2>/dev/null
if [ $router_type = "bridge1" ]; then
	rc-update add openvpn default     2>/dev/null
fi

#===============================================================================

echo -e $success"Configuration file installation completed!"$reset
