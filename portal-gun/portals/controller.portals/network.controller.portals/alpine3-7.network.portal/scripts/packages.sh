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

# TODO: Need to add changing the package repository
# options under /etc/apk/repositories


echo -e $text"apk updating & upgrading..."$reset
apk update
apk upgrade

# Add if dev environment checks (probably after we migrate
# to a real language)
echo -e $text"[DEV] apk installing vim..."$reset
apk add vim

# Add if check for hypervisor
# 
echo -e $text"[HYPERVISOR] apk installing libvirt..."$reset
apk add libvirt

echo -e $text"[HYPERVISOR] apk installing libvirt-daemon..."$reset
apk add libvirt-daemon

echo -e $text"[HYPERVISOR] apk installing qemu..."$reset
apk add qemu

echo -e $text"[HYPERVISOR] apk installing qemu-system-x86_64..."$reset
apk add qemu-system-x86_64
echo -e $text"[HYPERVISOR] apk installing qemu-img..."$reset
apk add qemu-img
echo -e $text"[HYPERVISOR] apk installing ovmf..."$reset
apk add ovmf
# Add if dev environment checks (probably after we migrate
# to a real language)
echo -e $text"[HYPERVISOR][DEV] apk installing virt-manager..."$reset
apk add virt-manager

echo -e $text"apk installing dhcp..."$reset
apk add dhcp
echo -e $text"apk installing shorewall..."$reset
apk add shorewall
echo -e $success"Package installation completed."$reset
