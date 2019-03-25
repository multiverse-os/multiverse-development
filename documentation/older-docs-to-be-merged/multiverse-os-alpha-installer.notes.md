# Building The First Multiverse OS Alpha Installer 

"When building linux, going down deep enough will always lead to circular dependencies." And this is circular dependencies all the way down.

There are four (4) basic components of a linux system capable of building itself under itself from source coude. (1) A kernel, (2) A C library (or Rust), (3) a set of command-line tool libraries, and (4) a set of build utilities (a compiler, linker, etc). 

A system doesnt have to build itself under itself, but for security reasons and with how powerful the security implemented in Multiverse is, it is incredibly desirable. 

## LIVE CD

Live CD based installer makes a lot of sense. It would provide a borked host + working controller and some basic applicaiton VMs. 
____________________________
[LIVE BOOT METHOD]
Install the following package:

`sudo apt-get install deboostrap squashfs-tools xorriso grub-pc-bin grub-efi-amd64-bin mtools`


mkdir $HOME/LIVE_BOOT

### Bootstrap and Configure eiban

Use `debootstrap` to configure the ARCH, variant base, version, and mirror. 

**NOTE** [For the Multiverse OS installer VM + software, we will iterate through all reasonable combinations of: [arch, version, etc]

`sudo debootstrap --arch=amd64 --variant=minbase stretch $HOME/LIVE_BOOT/chroot http://ftp.us.debian.org/debian/`


See: https://willhaley.com/blog/custom-debian-live-environment





