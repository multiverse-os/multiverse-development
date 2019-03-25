#!/bin/sh

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

echo -e $header"Multiverse OS: Star Router Package Installer"$reset
echo -e $accent"==============================================="$reset
echo -e $text"apk updating..."$reset
apk update
echo -e $text"apk installing$accent dhcp$text..."$reset
apk add dhcp
echo -e $text"apk installing$accent shorewall$text..."$reset
apk add shorewall
echo -e $text"apk installing$accent tor$text..."$reset
apk add shorewall
echo -e $success"Package installation completed."$reset
echo -e $subheader"# Services"$reset
echo -e $strong"dhcpd"$reset
echo -e $text"Adding$accent dhcp$text to rc-update"$reset
rc-update add dhcp default
echo -e $strong"shorewall"$reset
echo -e $text"Starting the$accent shorewall$text service..."$reset
rc-update add shorewall default
echo -e $strong"tor"$reset
echo -e $text"Starting the$accent tor$text service..."$reset
rc-update add tor default
echo -e $success"Service configuration completed."$reset
echo -e ""
echo -e $success"Package setup completed successfully!"$reset

