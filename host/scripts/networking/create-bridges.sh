# Proto-Multiverse OS Networking 
# ==============================
# TODO:
# The current solution is a stop-gap solution that fails to satisfy the basic
# functional requirements specified in the Multiverse OS design documents. 
# 
# Until we can build our desired solution using VSOCK and virtaul-pci-net devices
##=============================================================================
## CRITICAL
##=============================================================================
## TODO: [CRITICAL: Must migrate to using userspace isolated packet routing that
## entirely avoids host kernel, root level permissions and host networking k
## interfaces to meet minimum requirements of the Multiverse OS design 
## specification]
##-----------------------------------------------------------------------------
# That utilize shared memory, netmap, and pf_ring. We will leverage the existing
# userspace system of network bridges. Use of libvirt style network bridges are
# completely completely insecure, do not provide adequate isolation and fail to
# limit host control over the Multiverse OS cluster. Both libvirt and use of
# this style of bridge must be deprecated as soon as possible. 
##=============================================================================
##=============================================================================
## TODO: If not root, sudo su
#==============================================================================
## Initialize Network Bridges
##=============================================================================



# NETWORK 0
##=============================================================================

# Initialize net0br0
ip link add name net0br0 type bridge
ip link set dev net0br0 up 


# Initialize net0br1
ip link add name net0br1 type bridge
ip link set dev net0br1 up 

#=================================================
# TODO:
# Whonix bridge is being deprecated, we MUST build
# our own tor routing solution to satisfy security
# requirements, limit resource usage and meet
# basic functional requirements of Multiverse OS
# Initialize net0br2
#ip link add name net0br2 type bridge
#ip link set dev net0br2 up 
#================================================

# NETWORK 1
#==============================================================================

# Initialize net1br0
ip link add name net1br0 type bridge
ip link set dev net1br0 up 


# Initialize net1br1
ip link add name net1br1 type bridge
ip link set dev net1br1 up 


#=================================================
# TODO:
# Whonix bridge is being deprecated, we MUST build
# our own tor routing solution to satisfy security
# requirements, limit resource usage and meet
# basic functional requirements of Multiverse OS
# Initialize net0br2
# Initialize net1br2
#ip link add name net1br2 type bridge
#ip link set dev net1br2 up 
#================================================


