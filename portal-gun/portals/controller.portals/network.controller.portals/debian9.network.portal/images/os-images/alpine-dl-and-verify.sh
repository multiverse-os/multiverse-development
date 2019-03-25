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

echo -e $header"Alpine Standard 3.7.0"$reset
echo -e $accent"====================="$reset
echo -e $text"Downloading and verifying Alpine Standard 3.7.0 x86_64 install ISO image..."$reset

mkdir -p /var/multiverse/images/os-images
cd /var/multiverse/images/os-images

# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more
# sophisticated, for example it should scan the folder, find the highest version number so
# it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate

# TODO: The improved script should allow one to download any of the alpine images, with architecture
# selection using a menu.
wget http://dl-cdn.alpinelinux.org/alpine/v3.7/releases/x86_64/alpine-standard-3.7.0-x86_64.iso
wget http://dl-cdn.alpinelinux.org/alpine/v3.7/releases/x86_64/alpine-standard-3.7.0-x86_64.iso.sha256
wget http://dl-cdn.alpinelinux.org/alpine/v3.7/releases/x86_64/alpine-standard-3.7.0-x86_64.iso.asc

echo -e $subheader"##  DOWNLOAD"
echo -e $text"Successfully downloaded (1) the ISO image, (2) checksum files, and"$reset
echo -e $text"(3) checksum signature..."$reset

echo -e $subheader"##  CHECKSUM"
echo -e $text"Listing the checksums for each file downloaded..."$reset
cat alpine-standard-3.7.0-x86_64.iso.sha256


# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying
# on the user to manually compare.
echo -e $text"Executing 'sha256sum' on each file downloaded..."$reset
sha256sum alpine-standard-3.7.0-x86_64.iso

echo -e $text"Manually compare the values, and verify the checksums match..."$reset

echo -e $subheader"##  SIGNATURE"
echo -e $text"Downloading the Alpine Developer release key to verify the signature..."$reset
echo -e $text"The Alpine Developer key is located at: https://alpinelinux.org/"$reset
wget https://alpinelinux.org/keys/ncopa.asc

echo -e $text"Import the Alpine Developer key with gpg --import..."$reset
gpg --import ncopa.asc

echo -e $text"Verifying the signature file with the Alpine developer release key..."$reset
gpg --verify alpine-standard-3.7.0-x86_64.iso.asc

echo -e $header"    **NOTE** This script is simple and does not actually do comparisons, it simplifies"
echo -e "    the process by automating the steps, it is up to you to actually compare"
echo -e "    the checksums and read the output of the gpg --verify command."$reset

echo -e $text"Removing verification files..."$reset
rm -f alpine-standard-3.7.0-x86_64.iso.asc
rm -f alpine-standard-3.7.0-x86_64.iso.sha256
rm -r ncopa.asc

echo -e ""

cd $start_dir
echo -e $success"\nComplete!"$reset
