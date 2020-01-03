#!/bin/sh
##
## Voyager Service VM Config Installer
##==========================================

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

echo -e $header"Multiverse OS: Voyager Service VM Config Installer"$reset
echo -e $accent"=================================================="$reset
echo -e $text"Installing configuration files..."$reset
echo -e $subheader"/etc/network"$reset
echo -e $text"Removing interfaces file"$reset
rm -f /etc/network/interfaces
echo -e $text"Creating symbolic interfaces file link"$reset
ln -s /mnt/multiverse/config/etc/network/interfaces /etc/network/
echo -e $subheader"mother /etc/* files"$reset
rm -f /etc/hosts
rm -f /etc/hostname
rm -f /etc/issue
rm -f /etc/motd
echo -e $text"Deleting existing /etc/* configuration files..."$reset
ln -s /mnt/multiverse/config/etc/hosts /etc/
ln -s /mnt/multiverse/config/etc/hostname /etc/
ln -s /mnt/multiverse/config/etc/issue /etc/
ln -s /mnt/multiverse/config/etc/motd /etc/
echo -e $accent"===================================="$reset
echo -e $subheader"ECDSA Key Generation"$reset
ssh-keygen 

echo -e $accent"===================================="$reset
echo -e $subheader"Unpriviledged User Creation"$reset
echo -e $text"Creating 'user' account"$reset
adduser user

echo -e $success"Configuration file installation completed!"$reset

