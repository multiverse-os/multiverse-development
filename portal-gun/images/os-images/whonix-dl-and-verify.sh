#!/bin/sh

WHONIX_VERSION="13.0.0.1.4"

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

start_dir=$(pwd)

echo -e $text"Whonix Gateway $WHONIX_VERSION"$reset
echo -e $text"========================="$reset
echo -e $text"Downloading and verifying Whonix Gateway Libvirt KVM (VM) image..."$reset

mkdir -p /var/multiverse/images/os-images
cd /var/multiverse/images/os-images

# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more sophisticated, for example it should scan the folder, find the highest version number so it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate

wget https://download.whonix.org/linux/$WHONIX_VERSION/Whonix-Gateway-$WHONIX_VERSION.libvirt.xz
wget https://download.whonix.org/linux/$WHONIX_VERSION/Whonix-Gateway-$WHONIX_VERSION.libvirt.xz.asc

wget https://download.whonix.org/linux/$WHONIX_VERSION/Whonix-Gateway-$WHONIX_VERSION.sha512sums
wget https://download.whonix.org/linux/$WHONIX_VERSION/Whonix-Gateway-$WHONIX_VERSION.sha512sums.asc

wget https://www.whonix.org/patrick.asc

echo -e $subheader"##  DOWNLOAD"$reset
echo -e $text"Successfully downloaded (1) the Libvirt KVM image, (2) the Libvirt KVM image signature,"$reset
echo -e $text"(2) checksum file, and (3) checksum signature, and (4) Whonix Developer"$reset
echo -e $text"key"$reset

echo -e $subheader"##  CHECKSUM"$reset
echo -e $text"Listing the checksums for each file downloaded..."$reset
cat Whonix-Gateway-$WHONIX_VERSION.sha512sums


# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying on the user to manually compare.
echo -e $text"Executing 'sha512sum' on each file downloaded..."$reset
sha512sum Whonix-Gateway-$WHONIX_VERSION.libvirt.xz

echo -e $text"Manually compare the values, and verify the checksums match..."$reset

echo -e $subheader"## SIGNATURE"$reset
echo -e $text"Import the Whonix Developer key with gpg --import..."$reset
gpg --import patrick.asc

echo -e $text"Verifying the signature file with the Alpine developer release key..."$reset
gpg --verify Whonix-Gateway-$WHONIX_VERSION.libvirt.xz.asc


echo -e $header"    **NOTE** This script is simple and does not actually do comparisons, it simplifies"
echo -e "    the process by automating the steps, it is up to you to actually compare"
echo -e "    the checksums and read the output of the gpg --verify command."$reset

echo -e $subheader"##  Extracting from *.xz archive"$reset
tar -xvf Whonix-Gateway-$WHONIX_VERSION.libvirt.xz

echo -e $text"Remvoing *.xz archive to save space"$reset
rm Whonix-Gateway-$WHONIX_VERSION.libvirt.xz
 
echo -e $subheader"##  Shrink Whonix Gateway"$reset
echo -e $text"Shrinking *.qcow2 file to ~5 GB from 100 GB"$reset
qemu-img convert -O qcow2 Whonix-Gateway-13.0.0.1.4.qcow2 Whonix-Gateway-13.0.0.1.4.smaller.qcow2 
rm Whonix-Gateway-13.0.0.1.4.qcow2
mv Whonix-Gateway-13.0.0.1.4.smaller.qcow2 Whonix-Gateway-13.0.0.1.4.qcow2

echo -e "Removing verification files..."
rm -f Whonix-Gateway-13.0.0.1.4.libvirt.xz.asc
rm -f Whonix-Gateway-13.0.0.1.4.sha512sums
rm -f Whonix-Gateway-13.0.0.1.4.sha512sums.asc
rm -f Whonix-Gateway-13.0.0.1.4.xml
rm -f Whonix_network-13.0.0.1.4.xml
rm -f patrick.asc

echo -e ""

cd $start_dir
echo -e $success"\nComplete!"$reset
