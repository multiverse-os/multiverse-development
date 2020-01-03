#!/bin/sh

dpkg --add-architecture i386
sudo apt-get update
sudo apt-get upgrade

sudo apt-get install libgl1-mesa-dri:i386, libgl1-mesa-glx:i386, libc6:i386 git vim 
sudo apt-get remove nano

sudo dpkg -i steam.deb 
sudo apt-get -f install

# The driver for the video card needs to be identified and installed
# ideally, we could just check the users card and download it for them, checking# against known checksums.

## GeForce 970
## ===========

## arm32 / 32-bit
# http://us.download.nvidia.com/XFree86/Linux-x86-ARM/390.48/NVIDIA-Linux-armv7l-gnueabihf-390.48.run

## i386  / 32-bit
# http://us.download.nvidia.com/XFree86/Linux-x86/390.48/NVIDIA-Linux-x86-390.48.run

## amd64 / 64-bit
# http://us.download.nvidia.com/XFree86/Linux-x86_64/390.48/NVIDIA-Linux-x86_64-390.48.run

# The following instructions are for the card:
chmod +x NVIDIA-Linux-x86_64-390.48.run

sudo systemctl stop gdm3

sudo ./NVIDIA-Linux-x86_64-390.48.run

