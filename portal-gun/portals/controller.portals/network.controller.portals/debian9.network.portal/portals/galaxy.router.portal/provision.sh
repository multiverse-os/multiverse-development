#!/bin/sh

##
## Galaxy Router Config Installer
##==========================================
# DHCPd, Shorewall and various other /etc/*
# configuration files need to be installed
# from the shared storage

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

galaxy_portal_mount=/mnt/multiverse-portal
base_files=$galaxy_portal_mount/base-files

local_base=/var/multiverse/base-files
mkdir -p local_base


# rc-update adding will not work here for at least shorewall
# because it needs the required configurations, should be added
# to the provision script and this should ONLY install packages

echo -e $header"Multiverse OS: Galaxy Router Package Installer"$reset
echo -e $accent"==============================================="$reset
echo -e $text"apk updating..."$reset
apk update
echo -e $text"apk installing dhcp..."$reset
apk add dhcp
echo -e $text"apk installing shorewall..."$reset
apk add shorewall
echo -e $text"apk installing openvpn..."$reset
apk add openvpn
echo -e $success"Package installation completed."$reset
echo -e ""
echo -e $success"Package setup completed successfully!"$reset

#==============================================================================




echo -e $header"Multiverse OS: Galaxy Router Config Installer"$reset
echo -e $accent"============================================="$reset
echo -e $text"Installing configuration files..."$reset
echo -e $subheader"dhcpd"$reset
echo -e $text"Deleting existing dhcpd configuration files..."$reset
rm -f /etc/dhcp/dhcp.conf
rm -f /etc/dhcp/dhcpd.conf
rm -f /etc/dhcp/dhcpd.conf.example
echo -e $text"Creating symbolic links from shared storage configuration files..."$reset
mkdir -p $local_base/etc/dhcp/
cp $base_files/etc/dhcp/dhcpd.conf $local_base/etc/dhcp/
ln -s $local_base/etc/dhcp/dhcpd.conf /etc/dhcp/
echo -e $subheader"shorewall"$reset
echo -e $text"Deleting existing shorewall configuration files..."$reset
rm -f /etc/shorewall/hosts
rm -f /etc/shorewall/interfaces
rm -f /etc/shorewall/masq
rm -f /etc/shorewall/policy
rm -f /etc/shorewall/rules
rm -f /etc/shorewall/shorewall.conf
rm -f /etc/shorewall/snat
rm -f /etc/shorewall/zones
echo -e $text"Copying 'base-files' from multiverse portal into local /var/multiverse/base-files/..."$reset
mkdir -p $local_base/etc/shorewall/
cp $base_files/etc/shorewall/hosts          $local_base/etc/shorewall/
cp $base_files/etc/shorewall/interfaces     $local_base/etc/shorewall/
cp $base_files/etc/shorewall/masq           $local_base/etc/shorewall/
cp $base_files/etc/shorewall/policy         $local_base/etc/shorewall/
cp $base_files/etc/shorewall/rules          $local_base/etc/shorewall/
cp $base_files/etc/shorewall/shorewall.conf $local_base/etc/shorewall/
cp $base_files/etc/shorewall/snat           $local_base/etc/shorewall/
cp $base_files/etc/shorewall/zones          $local_base/etc/shorewall/
echo -e $text"Creating symbolic links from shared storage configuration files..."$reset
ln -s $local_base/etc/shorewall/hosts          /etc/shorewall/
ln -s $local_base/etc/shorewall/interfaces     /etc/shorewall/
ln -s $local_base/etc/shorewall/masq           /etc/shorewall/
ln -s $local_base/etc/shorewall/policy         /etc/shorewall/
ln -s $local_base/etc/shorewall/rules          /etc/shorewall/
ln -s $local_base/etc/shorewall/shorewall.conf /etc/shorewall/
ln -s $local_base/etc/shorewall/snat           /etc/shorewall/
ln -s $local_base/etc/shorewall/zones          /etc/shorewall/
echo -e $subheader"/etc/network"$reset
echo -e $text"Removing interfaces file"$reset
rm -f /etc/network/interfaces
echo -e $text"Creating symbolic interfaces file link"$reset
mkdir -p $local_base/etc/network/
cp $base_files/etc/network/interfaces $local_base/etc/network/
ln -s $local_base/etc/network/interfaces /etc/network/

echo -e $subheader"sysctl.d"$reset
echo -e $text"Creating symbolic link for 05-multiverse.conf file"$reset
rm -f /etc/sysctl.d/05-multiverse.conf
cp $base_files/etc/sysctl.d/05-multiverse.conf $local_base/etc/sysctl.d/
ln -s $local_base/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/

echo -e $subheader"other /etc/* files"$reset
rm -f /etc/hosts
rm -f /etc/issue
rm -f /etc/motd
echo -e $text"Copying more /etc/* configuration files..."$reset
cp $base_files/etc/hosts $local_base/etc/
cp $base_files/etc/issue $local_base/etc/
cp $base_files/etc/motd  $local_base/etc/
echo -e $text"Creating more /etc/* configuration files symbolic links..."$reset
ln -s $local_base/etc/hosts /etc/
ln -s $local_base/etc/issue /etc/
ln -s $local_base/etc/motd /etc/

echo -e $text"Adding services to rc-update default"$reset
rc-update add shorewall default
rc-update add openvpn default
rc-update add dhcpd default

echo -e ""
echo -e $success"Configuration file installation completed!"$reset

