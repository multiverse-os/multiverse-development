#!/bin/sh
###############################################################################
## DEVELOPMENT NOTE: 
# For now, during development, we want to leave the files in our multiverse
# portal share and symbolically link from the plan9 share. This way we can
# easily make changes and updates from our controller by simply editing the 
# files under portal-gun and have our changes immediately reflected in the
# router portals without needing to copy files around and easily be able to
# push our changes to the development repo so that changes can be shared
# with other developers without any extra work. 

###############################################################################
## Multiverse OS Script Color Palette
# TODO: With how often we are using this as a throw in to jazz up our dev 
# scripts, we should definitely be opting for a method that calls in a 
# library that is insalled into each machines multiverse folder. Because right
# now any changes we want to make to this would be a giant headache and that
# method of doing this is bad practice. 

header="\e[0;95m"
accent="\e[2m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"


###############################################################################
## Global Variables
# TODO: Ideally we want to move all unique values in the script to variables
# so we can begin merging provision scripts and eventually when we move to
# replace these with Go we will be able to easily determine what values should
# go into YAML config files so that we can have a single provisioning script
# that loads unique data for a given portal from a YAML config. 

dev_mode=1 # Currently all scripts are in development mode

multiverse_mount="/mnt/multiverse"
base_files="$multiverse_mount/base-files"

router_type="universe"
package_manager="apk"
# When converted to Go, the package manager type and associated logic
# will be stored in a struct so that we will know for example, what
# subcommand is used to install, uninstall, purge, force remove, update
# where the repository data is located, what repository options will be
# used and all other important information needed to utilize the package
# manager in a script. Because currently this would only work for alpine
# and just switching out apk to apt-get would not be sufficient, we would
# need additional variables for install and remove subcommand, and so
# this should just be a note and ignored until we start upgrading, which
# is a high priority item.
 
#packages=[] # Should be an array of packages stored here to simplify script
# when its moved to go, we can make a struct for packages, so we can easily
# indicate what files need to be deleted for configuration and what files
# need to be either copied for symbolically linked.



###############################################################################
## TODO: Ideally the below argument checking functionality should be moved into
##       a function then it can be put into provisioning script shell library
##       to condense down the logic to what is unique, making them more readable
##       and easily modifiable. 
##
##       This is mostly to help outline some of the options we want to have end
##       up in the final provisoning system. Ideally you can just run a portion
##       of the script, it will be broken into sections similar to the setup-alpine
##       script which is made up of a lot of sub scritps that can be ran individually
##       or run together in order when running the primary setup-alpine. 
#==============================================================================
echo -e $header"Multiverse OS:$reset$text Universe Router Provisioning"$reset
echo -e $accent"==========================================="$reset
if [ -z "$1" ]
then
	echo -e $strong"[WARNING]$reset$subheader No operating mode specified, current default setting is $text\`--dev\`$reset$subheader mode."$reset

else
	if [ "$1" == "-d" -o "$1" == "--dev" ]
	then
		echo -e $success"DEVELOPMENT MODE$reset$accent specified, using symbolic linking instead of copying..."$reset
	elif [ "$1" == "-c" -o "$1" == "--copy" -o "$1" == "-s" -o "$1" == "--std" ]
	then
		dev_mode=0
		echo -e $success"STANDARD MODE$reset$accent specified, copying files into the filesystem instead of using symbolic linking..."$reset
	elif [ "$1" == "-h" -o "$1" == "--help" ]
	then
		echo -e $subheader"The default is development mode, which instead of copying the"
		echo -e "configuration files during the provisioning process, the files are"
		echo -e "symbollically linked from the multiverse portal p9 share.\n"
		echo -e "Multiverse OS development provisioning scripts have two operating"
		echo -e "Modes: $accent[$text development$accent,$text standard$accent ]$reset\n"
		echo -e $text"Development Mode$reset        $subheader-d, --dev$reset"
		echo -e $text"Standard Mode   $reset        $subheader-c, --copy, -s, --std$reset\n"
		exit 0
	fi
fi


#==============================================================================
echo -e $accent"------------------------------------------------"$reset
	echo -e $header"Provisoning$reset$accent Installing Packages"$reset
echo -e $accent"------------------------------------------------"$reset
echo -e $subheader"# Install $accent'apk'$reset$subheader packages for router portal: [$accent dhcp$reset$subheader,$accent shorewall$reset$subheader ]"$reset
echo -e $accent"\`apk update\`$reset$text: updating installed packages"$reset
apk update &>/dev/null
echo -e $accent"\`apk add dhcp\`$reset$text: installing $reset$accent dhcp"$reset
apk add dhcp &>/dev/null
echo -e $accent"\`apk add shorewall\`$reset$text: installing $reset$accent shorewall"$reset
apk add shorewall &>/dev/null
## Remove development files and other unncessary packages 
# TODO: Determine what is installed by default and minimize the installation
# by removing unnecessary packages and locking down the router server as much
# as possible. Similar to the host lockdown phase.
#echo -e "Removing development packages: vim"
#apk del vim


#==============================================================================
## Configuration File Installation
# TODO: For now we will be creating symbolic links directly from the 
# multiverse portal p9 share. Ideally we will just comment out the
# copying option, and eventually we will create a switch that can be
# turned on with a command-line argument that will let us either use
# development mode which would create symbolic links or normal mode
# which would copy the files to the local macchine. 
echo -e $accent"------------------------------------------------"$reset
if [ $dev_mode == 0 ]
then
	echo -e $subheader"[Standard]$reset$header  Provisoning$reset$accent Configuration Installer"$reset
else
	echo -e $subheader"[Developer]$reset$header Provisoning$reset$accent Configuration Installer"$reset
fi
echo -e $accent"------------------------------------------------"$reset
echo -e $text"Installing configuration files..."$reset
#==============================================================================
echo -e $subheader"Configuring: [$text dhcpd$reset $subheader]"$reset
echo -e $text"Deleting existing dhcpd configuration files..."$reset
rm -f /etc/dhcp/dhcp.conf          2>/dev/null
rm -f /etc/dhcp/dhcpd.conf         2>/dev/null
rm -f /etc/dhcp/dhcpd.conf.example 2>/dev/null
if [ $dev_mode == 0 ]
then
	echo -e $subheader"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
	cp $base_files/etc/dhcp/dhcpd.conf /etc/dhcp/    2>/dev/null
else
	echo -e $subheader"[Development Mode]$reset$text Creating symbolic links from p9 shared storage..."$reset
	ln -s $base_files/etc/dhcp/dhcpd.conf /etc/dhcp/ 2>/dev/null
fi
#=======================================================================
echo -e $subheader"Configuring: [$text shorewall$reset $subheader]"$reset
echo -e $text"Deleting existing shorewall configuration files..."$reset
rm -f /etc/shorewall/hosts           2>/dev/null
rm -f /etc/shorewall/interfaces      2>/dev/null
rm -f /etc/shorewall/masq            2>/dev/null
rm -f /etc/shorewall/policy          2>/dev/null
rm -f /etc/shorewall/rules           2>/dev/null
rm -f /etc/shorewall/shorewall.conf  2>/dev/null
rm -f /etc/shorewall/snat            2>/dev/null
rm -f /etc/shorewall/zones           2>/dev/null
if [ $dev_mode == 0 ]
then
	echo -e $subheader"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
	cp $base_files/etc/shorewall/hosts          /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/interfaces     /etc/shorewall/  2>/dev/null   
	cp $base_files/etc/shorewall/masq           /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/policy         /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/rules          /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/shorewall.conf /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/snat           /etc/shorewall/  2>/dev/null
	cp $base_files/etc/shorewall/zones          /etc/shorewall/  2>/dev/null
else
	echo -e $subheader"[Development Mode]$reset$text Creating symbolic links from p9 shared storage..."$reset
	ln -s $base_files/etc/shorewall/hosts          /etc/shorewall/  2>/dev/null
	ln -s $base_files/etc/shorewall/interfaces     /etc/shorewall/  2>/dev/null
	ln -s $base_files/etc/shorewall/masq           /etc/shorewall/  2>/dev/null  
	ln -s $base_files/etc/shorewall/policy         /etc/shorewall/  2>/dev/null 
	ln -s $base_files/etc/shorewall/rules          /etc/shorewall/  2>/dev/null 
	ln -s $base_files/etc/shorewall/shorewall.conf /etc/shorewall/  2>/dev/null 
	ln -s $base_files/etc/shorewall/snat           /etc/shorewall/  2>/dev/null 
	ln -s $base_files/etc/shorewall/zones          /etc/shorewall/  2>/dev/null 
fi
#=======================================================================
echo -e $subheader"Network Interfaces: [$text /etc/network/interfaces$reset $subheader]"$reset
echo -e $text"Removing interfaces file"$reset
rm -f /etc/network/interfaces 2>/dev/null
if [ $dev_mode == 0 ]
then
	echo -e $subheader"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
	cp $base_files/etc/network/interfaces /etc/network/    2>/dev/null
else
	echo -e $subheader"[Development Mode]$reset$text Creating symbolic links from p9 shared storage..."$reset
	ln -s $base_files/etc/network/interfaces /etc/network/ 2>/dev/null
fi
#=======================================================================
echo -e $subheader"System Configuration: [$text /etc/sysctl.d/multiverse.conf$reset $subheader]"$reset
echo -e $text"Creating symbolic link for 05-multiverse.conf file"$reset
rm -f /etc/sysctl.d/05-multiverse.conf 2>/dev/null
if [ $dev_mode == 0 ]
then
	echo -e $t"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
	cp $base_files/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/    2>/dev/null

else
	echo -e $text"[Development Mode] Creating symbolic links from p9 shared storage..."$reset
	ln -s $base_files/etc/sysctl.d/05-multiverse.conf /etc/sysctl.d/ 2>/dev/null
fi
#=======================================================================
echo -e $subheader"General Linux Configuration: [$text /etc/*$reset $subheader]"$reset
rm -f /etc/hosts 2>/dev/null
rm -f /etc/issue 2>/dev/null
rm -f /etc/motd  2>/dev/null
if [ $dev_mode == 0 ]
then
	echo -e $subheader"[Standard Mode]$reset$text Copying configuration files from p9 shared storage..."$reset
	cp $base_files/etc/hosts /etc/ 2>/dev/null
	cp $base_files/etc/issue /etc/ 2>/dev/null
	cp $base_files/etc/motd  /etc/ 2>/dev/null
else
	echo -e $subheader"[Development Mode]$reset$text Creating symbolic links from p9 shared storage..."$reset
	ln -s $base_files/etc/hosts /etc/ 2>/dev/null
	ln -s $base_files/etc/issue /etc/ 2>/dev/null
	ln -s $base_files/etc/motd  /etc/ 2>/dev/null
fi
#=======================================================================
echo -e $subheader"Adding router services to startup via$reset$text rc-update"$reset
echo -e $text"Adding [$reset$subheader shorewall$reset,$subheader dhcpd$reset$text ]"$reset
rc-update add shorewall default 2>/dev/null
rc-update add dhcpd default     2>/dev/null
echo -e $success"Configuration file installation completed!"$reset
