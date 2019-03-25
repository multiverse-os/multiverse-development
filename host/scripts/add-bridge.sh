#!/bin/sh

if [ -z "$1" ]; then
	echo "[error] Bridge XML undefined. Must specify path to bridge XML."
fi

# TODO: When converting this to Go lets scan the XML and remove lines that
# are not important or desired, like UUID assignment, and ensure the XML
# is sane. As in does not use an existing subnet for the network, does not
# assign the HOST to *.*.*.1 (what the gateway will be in most cases). 
# Ensure that DHCP is not enabled. Assign MAC to 00:00:network subnet:fe
# and matching the ip which ideally will end in 254 (fe in hex). 
# Ideally since we are moving away from libvirt we should make the Go program
# more general if possible and easy to migrate away from libvirt and ensuring
# most of the code will be kept. 

sudo virsh net-define $1 2>/dev/null
