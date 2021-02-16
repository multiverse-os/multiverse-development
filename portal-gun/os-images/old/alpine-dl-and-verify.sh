#!/bin/sh
###############################################################################
# Multiverse OS Script Color Palette ##########################################
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"


###############################################################################
###############################################################################
## Go Conversion & Merging into a single program
#
# While this software is a sub-program of portal-gun, it should remain
# a standalone program that can be used to lookup os versions, download
# any image, automatically doing checksum validation. Automatically
# checking for updates, and providing access to multiple versions, so
# instead of just tracking what is the newest version, it should also
# track at least 3 versions back, and what download options are available
# like for example Alpine linux: standard, vm, minimal, etc 
#
# This information is provided by the new c library, libos, which is 
# what is used by gnome boxes for similar functionality. Another method
# would be using wikipedia to get the list of recent versions or some
# other linux OS API provided by some community site. I don't like the
# later two options as much. Libos relies on a embedded XML file that
# is regularly updates and included with the library. **So it may be
# best to just use that XML file, possibly locate the git it is
# maintained on and simply grab the XML file directly and use it. In
# the end all these methods rely on a third party and internet access
# so the best may end up being a combination of methods and crossp
# referencing.**
#
# With the example below, all the OS image logic is moved from the script
# logic and into the variables. This provides a frame to build each 
# iteration for each OS image, so the logic is generic enough that it
# can be used to download any OS installation or live image without
# changing the logic at all, just simply loading a different config
# file for each OS. This is important for the conversion to Go. 
#
#
# There are two other important changes that will be necessary to 
# introduce:
#
#     1) Tracking of several recent versions, ideally next version
#        and release date. Ability to download new or older versions.
#
#     2) Tracking of different "flavors"; for example, so one can
#        download server, or desktop when downloading Ubuntu. 
#
#

###############################################################################
## Variables
## Ideally these should allow these scripts to be essentially merged beyond
## these specific variables which could be loaded via YAML or similar. 
##
OS_NAME="alpine"
VERSION="standard"
ARCHITECTURE="x86_64"
MAJOR="3"
MINOR="11"
PATCH="2"
BASE_URL="http://dl-cdn.alpinelinux.org/" # Do we have https?
DL_URL="$BASE_URL/$OS_NAME/v$MAJOR.$MINOR/releases/$ARCHITECTURE/"

FULL_VERSION="$MAJOR.$MINOR.$PATCH"

FILENAME="$OS_NAME-$VERSION-$FULL_VERSION-$ARCHITECTURE.iso"
CHECKSUM_FILENAME="$FILENAME.sha256"
SIGNATURE_FILENAME="$FILENAME.asc"

DL_COMMAND="wget" # Later we may want to write our own, or use alternate
CHECKSUM_COMMAND="sha256sum"
GPG_COMMAND="gpg"

DEV_KEY_FILENAME="ncopa.asc"
REMOTE_DEV_KEY="https://alpinelinux.org/keys/"$DEV_KEY_FILENAME

DL_PATH="$(pwd)" # Can we just use wd? or $(pwd)
wd=$(pwd)

## TODO ########################################################################
# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more
# sophisticated, for example it should scan the folder, find the highest version number so
# it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate

# TODO: The improved script should allow one to download any of the alpine images, with architecture
# selection using a menu.
# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying 
# on the user to manually compare. (in reference to checksum comparing)


###############################################################################
echo -e $header"$OS_NAME $VERSION v$FULLVERSION"$reset
echo -e $accent"================================================================================"$reset
echo -e $text"  Downloading and verifying Alpine Standard 3.7.0 x86_64 install ISO image..."$reset
mkdir -p $DL_PATH
cd $DL_PATH


echo -e $accent"================================================================================"$reset
echo -e $subheader"##  DOWNLOAD"
echo -e $accent"================================================================================"$reset
echo -e $text"  Successfully downloaded (1) the ISO image, (2) checksum files, and"$reset
echo -e $text" (3) checksum signature..."$reset
$DL_COMMAND $DL_URL$FILENAME
$DL_COMMAND $DL_URL$CHECKSUM_FILENAME
$DL_COMMAND $DL_URL$SIGNATURE_FILENAME


echo -e $accent"================================================================================"$reset
echo -e $subheader"##  CHECKSUM"
echo -e $accent"================================================================================"$reset
echo -e $text"  Listing the checksums for each file downloaded..."$reset
cat $CHECKSUM_FILENAME
echo -e $text"  Executing '$CHECKSUM_COMMAND' on each file downloaded..."$reset
$CHECKSUM_COMMAND $FILENAME
echo -e $text"  Manually compare the values, and verify the checksums match..."$reset


echo -e $accent"================================================================================"$reset
echo -e $subheader"##  SIGNATURE VALIDATION"
echo -e $accent"================================================================================"$reset
echo -e $text"Downloading the $OS_NAME Developer release key to verify the signature..."$reset
echo -e $text"The $OS_NAME Developer key is located at: "$REMOTE_DEV_KEY$reset
$DL_COMMAND $REMOTE_DEV_KEY
echo -e $text"Import the $OS_NAME Developer key with '$GPG_COMMAND' --import..."$reset
$GPG_COMMAND --import $DEV_KEY_FILENAME
echo -e $text"Verifying the signature file with the $OS_NAME developer release key..."$reset
$GPG_COMMAND --verify $SIGNATURE_FILENAME
#echo -e $header"    **NOTE** This script is simple and does not actually do comparisons, it simplifies"
#echo -e "    the process by automating the steps, it is up to you to actually compare"
#echo -e "    the checksums and read the output of the gpg --verify command."$reset
#
#echo -e $text"Removing verification files..."$reset
#rm -f $SIGNATURE_FILENAME
#rm -f $CHECKSUM_FILENAME
#rm -r $DEV_KEY_FILENAME


echo -e $accent"================================================================================"$reset
cd $wd
echo -e $success"\nComplete!"$reset
