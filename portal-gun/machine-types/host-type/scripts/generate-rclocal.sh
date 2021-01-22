#!/bin/sh -e
###############################################################################
MULTIVERSE_DIR="/var/multiverse/scripts/pci-devices"
mkdir -p $MULTIVERSE_DIR
FILE="$MULTIVERSE_DIR/passthrough.sh"

echo "Processing lscpi to generate list of passthrough devices for this machine..."
echo "#!/bin/sh" > $FILE
echo "###############################################################################" >> $FILE
echo "\n" >> $FILE
echo "###############################################################################" >> $FILE
echo "## Default PCI device passthrough" >> $FILE
echo "## Generated by process-lspci.sh\n" >> $FILE
echo "## IMPORT vfio management module" >> $FILE
echo ". /home/user/multiverse/sh/modules/vfio-management.sh\n" >> $FILE
echo "###############################################################################" >> $FILE
echo "\n" >> $FILE

lspci -mnD | sed -E 's/"//g' | awk '{
  gsub(/^02..|^0d../, "network", $2)
  gsub(/^03../, "display", $2)
  gsub(/^09..|^0c../, "input", $2)
  gsub(/^040(1|3)/, "audio", $2)
  print "passthrough " $2 " " $3":"$4 " " $1
}' | grep -e "network\|display\|input\|audio" | sort >> $FILE

#### Interesting class prefixes ####
#### More details in /usr/share/misc/pci.ids
## Binding type "network"
#    02  Network controller
#    0d  Wireless controller
## Binding type "display"
#    03  Display controller
## Binding type "input"
#    09  Input device controller
#    0c  Serial bus controller
## Binding type "multimedia"
#    04  Multimedia controller
#      01  Multimedia audio controller
#      03  Audio device

chown user:libvirt $FILE
chmod 770 $FILE

echo "Installing systemctl service for vfio setup on boot..."
cp pci-passthrough.service $MULTIVERSE_DIR
chown user:libvirt $MULTIVERSE_DIR/pci-passthrough.service
chmod 770 $MULTIVERSE_DIR/pci-passthrough.service
ln -s $MULTIVERSE_DIR/pci-passthrough.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable pci-passthrough.service

echo "Passthrough configured.\n"
echo "Please review $FILE before next reboot."
echo "By default, all USB controllers are now owned by the controller VMS.\n"
echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
echo "!!!                                                                    !!!"
echo "     $FILE " 
echo "!!!                       must be edited manually                      !!!"
echo "!!! if you would like to reserve a USB controller for the host machine !!!"
echo "!!!                                                                    !!!"
echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
