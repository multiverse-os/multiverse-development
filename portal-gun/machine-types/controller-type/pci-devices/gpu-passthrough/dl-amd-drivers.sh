#!/bin/sh
##
## Controller VM Config Installer
##==========================================

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

echo -e $header"Multiverse OS: Controller VM GPU Setup"$reset
echo -e $accent"======================================"$reset

wget --referer http://support.amd.com https://www2.ati.com/drivers/linux/ubuntu/amdgpu-pro-17.50-511655.tar.xz
