#!/bin/sh

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

echo -e $header"Multiverse OS: Voyager Service VM Package Installer"$reset
echo -e $accent"=================================================="$reset
echo -e $subheader"# Packages"$reset
echo -e $text"apk updating..."$reset
apk update
echo -e $strong"Transmission Torrent Client"$reset
echo -e $text"apk installing$accent transmission daemon$reset..."$reset
apk add transmission-daemon
echo -e $text"apk installing$accent transmission cli$reset..."$reset
apk add transmission-cli
echo -e $strong"Git"$reset
echo -e $text"apk installing$accent git$reset..."$reset
apk add git
echo -e $success"Package installation completed."$reset
echo -e $subheader"# Services"$reset
echo -e $strong"transmission-daemon"$reset
echo -e $text"Adding$accent transmission-daemon$reset to rc-update"$reset
rc-update add transmission-daemon default
echo -e $text"Starting the$accent transmission-daemon$text service..."$reset
service transmission-daemon start
echo -e $success"Service configuration completed."$reset
echo -e ""
echo -e $success"Package setup completed successfully!"$reset

