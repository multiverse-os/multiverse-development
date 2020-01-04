#!/bin/sh

# Multiverse OS Script Color Palette
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
reset="\e[0m"

iso_filename="Fedora-Workstation-netinst-x86_64-27-1.6.iso"
checksum_filename="Fedora-Workstation-27-1.6-x86_64-CHECKSUM"
start_dir=$(pwd)

echo -e $header"Fedora Workstation 27-1.6 x86_64"$reset
echo -e $accent"=================="$reset
echo -e $text"Downloading and verifying$header Fedora Workstation 27-1.6$text Net Install ISO image..."$reset

mkdir -p /var/multiverse/images/os-images
cd /var/multiverse/images/os-images

# TODO: This script should eventually be replaced with a Go, Ruby or Rust script that is more
# sophisticated, for example it should scan the folder, find the highest version number so
# it will not need to be manually updated too often.

# TODO: Check if the file is already downloaded, if it is, skip, and just validate
wget https://download.fedoraproject.org/pub/fedora/linux/releases/27/Workstation/x86_64/iso/$iso_filename
wget https://download.fedoraproject.org/pub/fedora/linux/releases/27/Workstation/x86_64/iso/$checksum_filename

echo -e $subheader"##  DOWNLOAD"$reset
echo -e $text"Successfully downloaded (1) the ISO image (2) and signed checksum files"$reset


echo -e $subheader"##  SIGNATURES"$reset
echo -e $text"Downloading the Fedora developer release key to verify the signature..."$reset
wget https://getfedora.org/static/fedora.gpg -O - | gpg --import

echo -e $text"Verifying the signature file with the Fedora developer release key..."$reset
echo -e $text"The CHECKSUM file should have a good signature from the key F5282EE4"$reset
#  NOTE: The valid signature changes between releases:
#    429476B4 - Fedora 29
#    9DB62FB1 - Fedora 28
#    F5282EE4 - Fedora 27
#    64DAB85D - Fedora 26
#    3B921D09 - Fedora 26 secondary arches (AArch64, PPC64, PPC64le, s390 and s390x)

gpg --verify $checksum_filename


echo -e $header"\n    **NOTE** This script is simple and does not actually do comparisons, it simplifies"
echo -e "    the process by automating the steps, it is up to you to actually compare"
echo -e "    the checksums and read the output of the gpg --verify command."$reset

echo -e $subheader"##  CHECKSUM"$reset
echo -e $text"Manually compare the values, and verify the checksums match..."$reset
# TODO: In the improved version of this, we will actually do a _deep_ equals instead of relying
# on the user to manually compare.

echo -e $text"Listing the checksums for each file downloaded..."$reset
cat $checksum_filename | grep $iso_filename

echo -e $text"Executing 'sha256sum' on downloaded iso file..."$reset
sha256sum $iso_filename

echo -e $text"Removing checksum file..."$reset
rm -f $checksum_filename

cd $start_dir
echo -e $success"\nComplete!"$reset
