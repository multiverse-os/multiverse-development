#!/bin/sh

# ======================================
#
# Multiverse OS Script Color Palette
# --------------------------------------
header="\e[0;95m"
accent="\e[37m"
subheader="\e[98m"
strong="\e[96m"
text="\e[94m"
success="\e[92m"
warning="\e[93m"
fail="\e[91m"
reset="\e[0m"

# Global Variables
# Below are variables relating to Multiverse

MV_USER="user"
MV_GROUP="libvirt"
MV_CONFIG_PATH="/var/www/multiverse/"
MV_PATH="/home/user/multiverse-os/"

## ======================================
##
##  PCI Devices Shell Functions
## ======================================
device_id_lookup(){
  echo $(lspci -n | grep $1 | cut -c 15- | cut -c -9)
}

pci_address_lookup(){
  echo $(lspci -nn | grep $1 | cut -c -7)
}

kernel_module_lookup(){
  echo $(lspci -knn | grep $1 -A 3 | grep modules | cut -c 18-)
}

kernel_driver_lookup(){
  echo $(lspci -knn | grep $1 -A 2 | grep driver | cut -c 24-)
}

bind_device_to_vfio(){
  echo $( echo "$1" | sed 's/:/ /') > /sys/bus/pci/drivers/vfio-pci/new_id
}

# TODO: unbind device from host machine kernel

assign_device_qemu_bridge_permissions(){
  # After QEMU is updated, these permissions will need to be updated
  chown -R root:libvirt /usr/lib/qemu/
  chmod 4750 /usr/lib/qemu/qemu-bridge-helper
}

assign_device_config_space_permissions(){
  # Debian Buster introduced an error with config space permissions
  # that is no present in Debian Stretch. Below is a hack solution
  # to correct the permissions to allow forward movement on
  # Multiverse OS development.
  chown root:libvirt $DEVICE_SYSFS_PATH/config
  chmod 0660 $DEVICE_SYSFS_PATH/config
}

blacklist_kernel_module(){
  touch /etc/modprobe.d/multiverse.conf
  echo "blacklist $1" >> /etc/modprobe.d/multiverse.conf
}

length(){
	echo ${#1}
}

## ==========================================================
##
##   Bind Device with DEVICE_ID to vfio-pci for passthrough
## ----------------------------------------------------------
CURRENT_USER=$(whoami)
if [ $CURRENT_USER = "user" ]; then
  echo $fail"[Error] Must be logged in as root. Run 'su' and try again."$reset
  exit 0
fi

bind_device_id_to_vfio(){
  # TODO: Input Validation, validate format
  DEVICE_ID=$1
  echo $header"Multiverse OS: PCI Device 'vfio-pci' Bind Tool"$reset
  echo $accent"=============================================="$reset
  ## PCI Device Details
  PCI_ADDRESS=$(pci_address_lookup $DEVICE_ID)
  KERNEL_MODULE=$(kernel_module_lookup $DEVICE_ID)
  KERNEL_DRIVER=$(kernel_driver_lookup $DEVICE_ID)
  FULL_PCI_ADDRESS="0000:$PCI_ADDRESS"
  DEVICE_SYSFS_PATH="/sys/bus/pci/devices/$FULL_PCI_ADDRESS"
  UNBIND_FD_PATH="$DEVICE_SYSFS_PATH/driver/unbind"
  echo $subheader"Attempting to configure PCI Device"$reset
  echo $text"    Device ID:        "$reset $DEVICE_ID
  echo $text"    PCI Address:      "$reset $PCI_ADDRESS
  echo $text"    Full PCI Address: "$reset $FULL_PCI_ADDRESS
  echo $text"    SysFS Path:       "$reset $DEVICE_SYSFS_PATH
  echo $text"    Unbind FD Path:   "$reset $UNBIND_FD_PATH
  echo $text"    Kernel Module:    "$reset $KERNEL_MODULE
  echo $text"    Kernel Driver:    "$reset $KERNEL_DRIVER
  echo $text""$reset
  if [ "$KERNEL_DRIVER" = "vfio-pci" ]; then 
    echo $success"[Sucess] The specified device was already assignable directly to virtual machines."$reset
    exit 0
  fi
  ## PCI Device Unbinding
  echo $text"Checking if the specified PCI device is already bound..."$reset
  if [ ! -z $UNBIND_FD_PATH ] ; then
    echo "Passing Device PCI Address to '$UNBIND_FD_PATH'."
    #echo "Running 'ls' on UNBIND Folder"
    #echo $(ls $UNBIND_FD_PATH)
    echo $FULL_PCI_ADDRESS > $UNBIND_FD_PATH
  else
    #echo "Running 'ls' on UNBIND Folder"
    #echo $(ls $UNBIND_FD_PATH)
    echo $text"Device is already unbound, binding to 'vfio-pci' to enable PCI passthrough."$reset
  fi
  echo "Attempting to bind device to 'vfio-pci' to make it assignable..."
  bind_device_to_vfio $DEVICE_ID
  echo $success"[Success]$reset PCI Device can now be directly assigned to a virtual machine."$reset
}

# Executable Requirements
if [ -z "$1" ]; then
  echo $fail"[Error]$reset$text No$strong Device Id$reset$text or$strong PCI Address$reset$text of device to be passed through."$reset
  echo $text"Usage:$reset$header vfio-bind 8086:15b8$reset$text or$reset$header vfio-bind 02:00.0"$reset
  exit 0
else
  length_result=$(length "$1")
  if [ $length_result = "9" ]; then
    bind_device_id_to_vfio $1
  fi

  if [ $length_result = "7" ]; then
    echo "Looking for device with PCI Address $1..."
    dev_id=$(device_id_lookup $1)
    echo "Found device with ID [$dev_id] at address $1."
    bind_device_id_to_vfio $dev_id
  fi
  if [ $length_result = "12" ]; then
    lookup_address=$(echo $1 | cut -c 6-)
    echo "Looking for device with PCI Address $lookup_address..."
    echo "Using cut address $lookup_address"
    dev_id=$(device_id_lookup $lookup_address)
    echo "Found device with ID [$dev_id] at address $lookup_address."
    bind_device_id_to_vfio $dev_id
  fi
fi
