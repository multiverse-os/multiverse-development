#!/bin/sh
###############################################################################
# Proto-Multiverse OS Networking 
###############################################################################

# Until we can build our desired solution using VSOCK and virtaul-pci-net devices 
# that utilize shared memory, netmap, and pf_ring. We will leverage the existing
# userspace system of network bridges. 
###############################################################################
# Initialize Network Bridges
###############################################################################
## NETWORK 0
###############################################################################
#### Initialize net0br0
ip link add name net0br0 type bridge
ip link set dev net0br0 up 

#### Initialize net0br1
ip link add name net0br1 type bridge
ip link set dev net0br1 up 

###############################################################################
###############################################################################
## NETWORK 1
###############################################################################
#### Initialize net1br0
ip link add name net1br0 type bridge
ip link set dev net1br0 up 

#### Initialize net1br1
ip link add name net1br1 type bridge
ip link set dev net1br1 up 

