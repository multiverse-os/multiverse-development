#!/bin/sh
###############################################################################
## KVM Nested VMs Check
###############################################################################
## Instead of defining functions, these shell files can be used as functions easily and more cleanly than typical variable usage. 
##
## TODO: Check if `Y` or `N`
##
###############################################################################

NESTED=$(cat /sys/module/kvm_intel/parameters/nested)

if [ "$NESTED" = "Y" ]; then
	echo "1"
	return 1
else
	echo "0"
	return 0
fi

