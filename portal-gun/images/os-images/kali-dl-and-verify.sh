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

echo -e $header"Kali 2018.3 amd64"$reset
echo -e $accent"=================="$reset
echo -e $text"Downloading and verifying$header Kali 2018.3$text Install ISO image..."$reset

mkdir -p /var/multiverse/images/os-images
cd /var/multiverse/images/os-images

# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more
# sophisticated, for example it should scan the folder, find the highest version number so
# it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate
wget https://cdimage.kali.org/kali-2018.3/kali-linux-2018.3-amd64.iso
wget http:///mirror.pwnieexpress.com/kali-images/kali-weekly/SHA256SUMS
wget http:///mirror.pwnieexpress.com/kali-images/kali-weekly/SHA256SUMS.gpg

echo -e $subheader"##  DOWNLOAD"$reset
echo -e $text"Successfully downloaded (1) the ISO image, (2) checksum files, and"$reset
echo -e $text"(3) checksum signature..."$reset


echo -e $subheader"##  CHECKSUM"$reset
echo -e $text"Listing the checksums for each file downloaded..."$reset
cat SHA256SUMS

# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying
# on the user to manually compare.
echo -e $text"Executing 'sha256sum' on each file downloaded..."$reset
sha256sum kali-linux-2018.3-amd64.iso

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
