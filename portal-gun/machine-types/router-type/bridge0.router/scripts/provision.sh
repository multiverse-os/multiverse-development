#!/bin/sh
###############################################################################
## Multiverse OS Script Color Palette
###############################################################################
header="\e[0;95m"
accent="\e[2m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"
###############################################################################
## Global Variables
###############################################################################
multiverse_mount="/mnt/multiverse"
base_files="$multiverse_mount/base-files"
router_type="universe"
package_manager="apk"
#==============================================================================
echo -e $header"Multiverse OS:$reset$text Universe Router Provisioning"$reset
echo -e $accent"==========================================="$reset
#==============================================================================
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
echo -e $accent"------------------------------------------------"$reset
echo -e $subheader"[Bridge]$reset$header Provisoning$reset$accent Configuration Installer"$reset
echo -e $accent"------------------------------------------------"$reset
echo -e $text"Installing configuration files..."$reset
#==============================================================================
echo -e $subheader"Configuring: [$text dhcpd$reset $subheader]"$reset
echo -e $text"Deleting existing dhcpd configuration files..."$reset
rm -f /etc/dhcp/dhcp.conf          2>/dev/null
rm -f /etc/dhcp/dhcpd.conf         2>/dev/null
rm -f /etc/dhcp/dhcpd.conf.example 2>/dev/null
echo -e $subheader"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
cp $base_files/etc/dhcp/dhcpd.conf /etc/dhcp/    2>/dev/null
#=======================================================================
echo -e $subheader"Configuring: [$text shorewall$reset $subheader]"$reset
echo -e $text"Deleting existing shorewall configuration files..."$reset
rm -f /etc/shorewall/hosts           2>/dev/null
rm -f /etc/shorewall/interfaces      2>/dev/null
rm -f /etc/shorewall/masq            2>/dev/null
rm -f /etc/shorewall/policy          2>/dev/null
rm -f /etc/shorewall/rules           2>/dev/null
rm -f /etc/shorewall/shorewall.conf  2>/dev/null
rm -f /etc/shorewall/snat            2>/dev/null
rm -f /etc/shorewall/zones           2>/dev/null
echo -e $subheader"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
cp $base_files/etc/shorewall/hosts          /etc/shorewall/  2>/dev/null
cp $base_files/etc/shorewall/interfaces     /etc/shorewall/  2>/dev/null   
cp $base_files/etc/shorewall/masq           /etc/shorewall/  2>/dev/null
cp $base_files/etc/shorewall/policy         /etc/shorewall/  2>/dev/null
cp $base_files/etc/shorewall/rules          /etc/shorewall/  2>/dev/null
cp $base_files/etc/shorewall/shorewall.conf /etc/shorewall/  2>/dev/null
cp $base_files/etc/shorewall/snat           /etc/shorewall/  2>/dev/null
cp $base_files/etc/shorewall/zones          /etc/shorewall/  2>/dev/null
#=======================================================================
echo -e $subheader"Network Interfaces: [$text /etc/network/interfaces$reset $subheader]"$reset
echo -e $text"Removing interfaces file"$reset
rm -f /etc/network/interfaces 2>/dev/null
echo -e $subheader"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
cp $base_files/etc/network/interfaces /etc/network/    2>/dev/null
#=======================================================================
echo -e $subheader"System Configuration: [$text /etc/sysctl.d/multiverse.conf$reset $subheader]"$reset
echo -e $text"Creating symbolic link for 05-multiverse.conf file"$reset
rm -f /etc/sysctl.d/05-multiverse.conf 2>/dev/null
echo -e $t"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
cp $base_files/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/    2>/dev/null
#=======================================================================
echo -e $subheader"General Linux Configuration: [$text /etc/*$reset $subheader]"$reset
rm -f /etc/hosts 2>/dev/null
rm -f /etc/issue 2>/dev/null
rm -f /etc/motd  2>/dev/null
echo -e $subheader"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
cp $base_files/etc/hosts /etc/ 2>/dev/null
cp $base_files/etc/issue /etc/ 2>/dev/null
cp $base_files/etc/motd  /etc/ 2>/dev/null
#=======================================================================
echo -e $subheader"Adding router services to startup via$reset$text rc-update"$reset
echo -e $text"Adding [$reset$subheader shorewall$reset,$subheader dhcpd$reset$text ]"$reset
rc-update add shorewall default 2>/dev/null
rc-update add dhcpd default     2>/dev/null
echo -e $success"Configuration file installation completed!"$reset
