#!/bin/sh

##
## Star Router Config Installer
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

echo -e $header"Multiverse OS: Star Router Config Installer"$reset
echo -e $accent"============================================="$reset

echo -e $text"Running ./packages.sh script..."$reset
echo -e $text"Installing configuration files..."$reset
echo -e $subheader"dhcpd"$reset
echo -e $text"Deleting existing dhcpd configuration files..."$reset
rm -f /etc/dhcp/dhcp.conf
rm -f /etc/dhcp/dhcpd.conf
rm -f /etc/dhcp/dhcpd.conf.example
echo -e $text"Creating symbolic links from shared storage configuration files..."$reset
ln -s /mnt/multiverse/config/etc/dhcp/dhcpd.conf /etc/dhcp/
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
echo -e $text"Creating symbolic links from shared storage configuration files..."$reset
ln -s /mnt/multiverse/config/etc/shorewall/hosts /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/interfaces /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/masq /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/policy /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/rules /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/shorewall.conf /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/snat /etc/shorewall/
ln -s /mnt/multiverse/config/etc/shorewall/zones /etc/shorewall/
echo -e $subheader"/etc/network"$reset
echo -e $text"Removing interfaces file"$reset
rm -f /etc/network/interfaces
echo -e $text"Creating symbolic interfaces file link"$reset
ln -s /mnt/multiverse/config/etc/network/interfaces /etc/network/
echo -e $subheader"sysctl.d"$reset
echo -e $text"Creating symbolic link for 05-multiverse.conf file"$reset
rm -f /etc/sysctl.d/05-multiverse.conf
ln -s /mnt/multiverse/config/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/
echo -e $subheader"sysctl.d"$reset
echo -e $text"Creating symbolic link for 05-multiverse.conf file"$reset
rm -f /etc/sysctl.d/05-multiverse.conf
ln -s /mnt/multiverse/config/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/
echo -e $subheader"other /etc/* files"$reset
rm -f /etc/hosts
rm -f /etc/hostname
rm -f /etc/issue
rm -f /etc/motd
echo -e $text"Deleting existing /etc/* configuration files..."$reset
ln -s /mnt/multiverse/config/etc/hosts /etc/
ln -s /mnt/multiverse/config/etc/hostname /etc/
ln -s /mnt/multiverse/config/etc/issue /etc/
ln -s /mnt/multiverse/config/etc/motd /etc/

echo -e ""
echo -e $success"Configuration file installation completed!"$reset


