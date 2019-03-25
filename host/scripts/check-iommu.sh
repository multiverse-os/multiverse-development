#!/bin/sh
## NOTE: VMs will say intel regardless if the passthrough CPU is intel 

sudo dmesg | grep -e DMAR -e IOMMU


###############################################################################
#### TODO
####
##
## TODO: Intel will report: "Intel-IOMMU: Enabled" if successful.
##
## TODO: Check if driver is loaded by using "(AMD|INTEL) IOMMUv2 driver"
## TODO: Check if "functionality not available on this system" is present,
## because if so, the machine will not support passthrough.
##
