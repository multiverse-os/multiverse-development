#!/bin/sh

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

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

