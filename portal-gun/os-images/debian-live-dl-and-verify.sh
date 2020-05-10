#!/bin/sh
###############################################################################
# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

start_dir=$(pwd)

MAJOR="10"
MINOR="3"
PATCH="0"

DL_PATH="$(pwd)"

ISO_TYPE="gnome"
# Options: cinnamon, gnome, kde, lxde, lxqt, mate, standard, xfce

ARCH="amd64"

DEPS="dirmngr"

sudo apt-get install $DEPS


echo -e $header"Debian $MAJOR.$MINOR.$PATCH $ARCH"$reset
echo -e $accent"=================="$reset
echo -e $text"Downloading and verifying$header Debian $MAJOR.$MINOR.$PATCH$text Net Install ISO image..."$reset

mkdir -p /var/multiverse/images/os-images
cd /var/multiverse/images/os-images

# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more
# sophisticated, for example it should scan the folder, find the highest version number so
# it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate
wget https://cdimage.debian.org/cdimage/release/current-live/amd64/iso-hybrid/debian-live-$MAJOR.$MINOR.$PATCH-$ARCH-$ISO_TYPE.iso
wget https://cdimage.debian.org/cdimage/release/current-live/amd64/iso-hybrid/SHA256SUMS
wget https://cdimage.debian.org/cdimage/release/current-live/amd64/iso-hybrid/SHA256SUMS.sign

echo -e $subheader"##  DOWNLOAD"$reset
echo -e $text"Successfully downloaded (1) the ISO image, (2) checksum files, and"$reset
echo -e $text"(3) checksum signature..."$reset


echo -e $subheader"##  CHECKSUM"$reset
echo -e $text"Listing the checksums for each file downloaded..."$reset
cat SHA256SUMS

# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying
# on the user to manually compare.
echo -e $text"Executing 'sha256sum' on each file downloaded..."$reset
sha256sum debian-live-$MAJOR.$MINOR.$PATCH-$ARCH-$ISO_TYPE.iso

echo -e $text"Manually compare the values, and verify the checksums match..."$reset

echo -e $subheader"##  SIGNATURES"$reset
echo -e $text"Downloading the Debian developer release key to verify the signature..."$reset
gpg --keyserver keyring.debian.org --recv 6294BE9B


echo -e $text"Verifying the signature file with the Debian developer release key..."$reset
gpg --verify SHA256SUMS.sign

echo -e $header"\n    **NOTE** This script is simple and does not actually do comparisons, it simplifies"
echo -e "    the process by automating the steps, it is up to you to actually compare"
echo -e "    the checksums and read the output of the gpg --verify command."$reset

echo -e $text"Removing verification files..."$reset
#rm -f SHA256SUMS
#rm -f SHA256SUMS.sign

echo -e ""

cd $start_dir
echo -e $success"\nComplete!"$reset
