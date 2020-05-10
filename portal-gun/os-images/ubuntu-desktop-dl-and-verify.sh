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

MAJOR="18"
MINOR="04"
PATCH="4"

ARCH="amd64"

ISO_TYPE="live-server"

echo -e $header"Ubuntu $ISO_TYPE $MAJOR.$MINOR.$PATCH $ARCH"$reset
echo -e $accent"============================"$reset
echo -e $text"Downloading and verifying$header Ubuntu $ISO_TYPE $MAJOR.$MINOR.$PATCH $ARCH$text Install ISO image..."$reset

mkdir -p /var/multiverse/images/os-images
cd /var/multiverse/images/os-images

# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more
# sophisticated, for example it should scan the folder, find the highest version number so
# it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate
wget http://releases.ubuntu.com/$MAJOR.$MINOR.$PATCH/ubuntu-$MAJOR.$MINOR.$PATCH-$ISO_TYPE-$ARCH.iso
wget http://releases.ubuntu.com/$MAJOR.$MINOR.$PATCH/SHA256SUMS
wget http://releases.ubuntu.com/$MAJOR.$MINOR.$PATCH/SHA256SUMS.gpg
echo -e $subheader"##  DOWNLOAD"$reset
echo -e $text"Successfully downloaded (1) the ISO image, (2) checksum files, and"$reset
echo -e $text"(3) checksum signature..."$reset


echo -e $subheader"##  CHECKSUM"$reset
echo -e $text"Listing the checksums for each file downloaded..."$reset
cat SHA256SUMS

# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying
# on the user to manually compare.
echo -e $text"Executing 'sha256sum' on each file downloaded..."$reset
sha256sum ubuntu-$MAJOR.$MINOR.$PATCH-$ISO_TYPE-$ARCH.iso

echo -e $text"Manually compare the values, and verify the checksums match..."$reset

echo -e $subheader"##  SIGNATURES"$reset
echo -e $text"Downloading the Ubuntu developer release key to verify the signature..."$reset
gpg --keyserver hkp://keyserver.ubuntu.com --recv-keys "8439 38DF 228D 22F7 B374 2BC0 D94A A3F0 EFE2 1092" "C598 64BF 1257 FFA8 6632 CBA7 4618 1433 FBB7 5451"


echo -e $text"Verifying the signature file with the Ubuntu developer release key..."$reset
gpg --verify SHA256SUMS.gpg SHA256SUMS

echo -e $header"\n    **NOTE** This script is simple and does not actually do comparisons, it simplifies"
echo -e "    the process by automating the steps, it is up to you to actually compare"
echo -e "    the checksums and read the output of the gpg --verify command."$reset

echo -e $text"Removing verification files..."$reset
rm -f SHA256SUMS
rm -f SHA256SUMS.gpg

echo -e ""

cd $start_dir
echo -e $success"\nComplete!"$reset
