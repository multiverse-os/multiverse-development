#!/bin/sh

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"


echo -e $header"Multiverse OS: Universe Router Package Installer"$reset
echo -e $accent"==============================================="$reset
echo -e $subpackage"# Packages"$reset
echo -e $text"apk updating..."$reset
apk update
echo -e $text"apk installing dhcp..."$reset
apk add dhcp
echo -e $text"apk installing shorewall..."$reset
apk add shorewall
echo -e $success"Package installation completed."$reset
