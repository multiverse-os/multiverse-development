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

echo -e $header"Multiverse OS: Tor Router XML Backup"$reset
echo -e $accent"==============================================="$reset
echo -e $text"Backing up Tor Router VM active XML files on the host machine..."$reset
echo -e $text"Copying $accent template.tor.router.multiverse.xml $text ..."$reset
mv /home/user/.config/libvirt/qemu/template.tor.router.multiverse.xml /home/user/multiverse/machines/tor.router.multiverse/xml/
echo -e "Creating symbolic link from the original to the newly copied file..."

ln -s /home/user/multiverse/machines/tor.router.multiverse/xml/ /home/user/.config/libvirt/qemu/template.tor.router.multiverse.xml

echo -e $text"Copying $accent tor.router.multiverse.xml $text ..."$reset
cp /home/user/.config/libvirt/qemu/tor.router.multiverse.xml /home/user/multiverse/machines/tor.router.multiverse/xml/
 
echo -e $success"Successfully copied the *.xml files!"$reset




