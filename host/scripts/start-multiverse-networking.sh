
echo "Multiverse OS"
echo "============="
echo "Initializing Multiverse networking..."


echo "(1) Creating bridges to connect routers and VMs..."
#==============================================================================
## Initialize Network Bridges
##=============================================================================
# NETWORK 0
##=============================================================================
# Initialize net0br0
sudo ip link add name net0br0 type bridge
sudo ip link set dev net0br0 up 

# Initialize net0br1
sudo ip link add name net0br1 type bridge
sudo ip link set dev net0br1 up 

# Initialize net0br2
sudo ip link add name net0br2 type bridge
sudo ip link set dev net0br2 up 

# NETWORK 1
#==============================================================================
# Initialize net1br0
sudo ip link add name net1br0 type bridge
sudo ip link set dev net1br0 up 

# Initialize net1br1
sudo ip link add name net1br1 type bridge
sudo ip link set dev net1br1 up 

# Initialize net1br2
sudo ip link add name net1br2 type bridge
sudo ip link set dev net1br2 up 

#==============================================================================
## Boot Routers
##=============================================================================
echo "(2) Booting up Universe0 and Universe1 routers and starting the corresponding networks..."
# Network 0
#=========================================
virsh start universe0.router.portal

# Network 1
#=========================================
#virsh start universe1.router.portal
sleep 10
echo "Booting galaxy0 for Universe0 and galaxy0 for Universe1..."

# Network 0
#=========================================
virsh start galaxy0.router.universe0

# Network 1
#=========================================
#virsh start galaxy0.router.universe1
sleep 5
echo "Booting star0 for Universe0 and star0 for Universe1..."

# Network 0
#=========================================
virsh start star0.whonix-router.universe0

# Network 1
#=========================================
#virsh start star0.whonix-router.universe1

echo "All routers booted, ready to start the user interface controller portal..."


echo "Checking status of routers and networks..."

virsh list
echo "============"
echo "Options for User Interface Controllers:"
virsh list --all | grep controller
