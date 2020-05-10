#!/bin/sh

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

start_dir=$(pwd)

YEAR="2020"
VERSION="1b"
TYPE="live"
# Options: 'live', 'installer', 'installer-netinst'
ARCH="amd64"

echo -e $header"Kali $YEAR.$VERSION $ARCH"$reset
echo -e $accent"=================="$reset
echo -e $text"Downloading and verifying$header Kali $YEAR.$VERSION $TYPE$text ISO image..."$reset

mkdir -p /var/multiverse/images/os-images
cd /var/multiverse/images/os-images

# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more
# sophisticated, for example it should scan the folder, find the highest version number so
# it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate
wget -c https://cdimage.kali.org/kali-$YEAR.$VERSION/kali-linux-$YEAR.$VERSION-$TYPE-$ARCH.iso
wget https://cdimage.kali.org/kali-$YEAR.$VERSION/SHA256SUMS
wget https://cdimage.kali.org/kali-$YEAR.$VERSION/SHA256SUMS.gpg


echo -e $subheader"##  DOWNLOAD"$reset
echo -e $text"Successfully downloaded (1) the ISO image, (2) checksum files, and"$reset
echo -e $text"(3) checksum signature..."$reset


echo -e $subheader"##  CHECKSUM"$reset
echo -e $text"Listing the checksums for each file downloaded..."$reset
# TODO: MAKE SURE THIS IS THE ACTUAL SUMS AND NOT DEBIAN ONES. SINCE THEY HAVE THE SAME NAME
cat SHA256SUMS

# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying
# on the user to manually compare.
echo -e $text"Executing 'sha256sum' on each file downloaded..."$reset
sha256sum kali-linux-$YEAR.$VERSION-$TYPE-$ARCH.iso

echo -e $text"Manually compare the values, and verify the checksums match..."$reset

echo -e $subheader"##  SIGNATURES"$reset
echo -e $text"Downloading the Kali developer release key to verify the signature..."$reset
#wget -q -O - https://www.kali.org/archive-key.asc | gpg --import
gpg --keyserver hkp://keys.gnupg.net --recv-key 7D8D0BF6


gpg --fingerprint 7D8D0BF6



echo -e $text"Verifying the signature file with the Kali developer release key..."$reset
gpg --verify SHA256SUMS.gpg SHA256SUMS

echo -e $header"\n    **NOTE** This script is simple and does not actually do comparisons, it simplifies"
echo -e "    the process by automating the steps, it is up to you to actually compare"
echo -e "    the checksums and read the output of the gpg --verify command."$reset

echo -e $text"Removing verification files..."$reset
rm -f SHA256SUMS
rm -f SHA256SUMS.sign

echo -e ""

cd $start_dir
echo -e $success"\nComplete!"$reset
