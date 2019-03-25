##
##
## reserved values
##
#255	local
#254	main
#253	default
#0	unspec
##
## local
##
##1	inr.ruhep
#1 net0br0
#2 net0br1
#3 net0br2
#4 net1br0
#5 net1br1
#6 net1br2


# net0br0 Routing Table
#==============================================================================
# Add galaxy routing table
#cat "1 net0br0" >> /etc/iproute2/rt_tables

# Establish the default route for the galaxy routing table
sudo ip route add default via 10.0.0.1 dev ens4 table net0br0

# Select which subnets follow the rules within the galaxy routing table
sudo ip rule add from 10.0.0.0/24 table net0br0
sudo ip rule add from 10.255.0.0/24 table net0br0

# Copied from the original routing table
sudo ip route add 10.0.0.0/24 dev ens4 proto kernel scope link src 10.0.0.10 table net0br0
sudo ip route add 10.255.0.0/24 dev net0br0 proto kernel scope link src 10.255.0.1 table net0br0
sudo ip route add 10.255.0.0/24 via 10.255.0.1 dev net0br0 metric 1 table net0br0

# net0br1 Routing Table
#==============================================================================
# Add galaxy routing table
#cat "2 net0br1" >> /etc/iproute2/rt_tables

# Establish the default route for the galaxy routing table
sudo ip route add default via 10.1.1.1 dev ens4 table net0br1

# Select which subnets follow the rules within the galaxy routing table
sudo ip rule add from 10.1.1.0/24 table net0br1
sudo ip rule add from 10.255.1.0/24 table net0br1

# Copied from the original routing table
sudo ip route add 10.1.1.0/24 dev ens4 proto kernel scope link src 10.1.1.10 table net0br1
sudo ip route add 10.255.1.0/24 dev net0br1 proto kernel scope link src 10.255.1.1 table net0br1
sudo ip route add 10.255.1.0/24 via 10.255.1.1 dev net0br1 metric 1 table net0br1


# net0br2 Routing Table
#==============================================================================
#cat "3 net0br2" >> /etc/iproute2/rt_tables


# net1br0 Routing Table
#==============================================================================
# Add galaxy routing table
#cat "4 net1br0" >> /etc/iproute2/rt_tables

# Establish the default route for the galaxy routing table
sudo ip route add default via 10.100.100.1 dev ens4 table net1br0

# Select which subnets follow the rules within the galaxy routing table
sudo ip rule add from 10.100.100.0/24 table net1br0
sudo ip rule add from 10.255.100.0/24 table net1br0

# Copied from the original routing table
sudo ip route add 10.100.100.0/24 dev ens4 proto kernel scope link src 10.100.100.10 table net1br0
sudo ip route add 10.255.100.0/24 dev net1br0 proto kernel scope link src 10.255.100.1 table net1br0
sudo ip route add 10.255.100.0/24 via 10.255.100.1 dev net1br0 metric 1 table net1br0

# net0br1 Routing Table
#==============================================================================
# Add galaxy routing table
#cat "5 net1br1" >> /etc/iproute2/rt_tables

# Establish the default route for the galaxy routing table
sudo ip route add default via 10.110.110.1 dev ens4 table net1br1

# Select which subnets follow the rules within the galaxy routing table
sudo ip rule add from 10.110.110.0/24 table net1br1
sudo ip rule add from 10.255.110.0/24 table net1br1

# Copied from the original routing table
sudo ip route add 10.110.110.0/24 dev ens4 proto kernel scope link src 10.110.110.10 table net1br1
sudo ip route add 10.255.110.0/24 dev net1br1 proto kernel scope link src 10.255.110.1 table net1br1
sudo ip route add 10.255.110.0/24 via 10.255.110.1 dev net1br1 metric 1 table net1br1


# net1br2 Routing Table
#==============================================================================
#cat "6 net1br2" >> /etc/iproute2/rt_tables


# Then cleanup by disabling default routing table or at least set it to default to a net0br1 so cant dns leak actual location
