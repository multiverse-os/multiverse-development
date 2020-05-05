# TODO: This should be a GO program so we can have this stuff done based on the system and not manual changes editing it for each system
#default via 10.1.1.1 dev enp6s0 proto dhcp metric 20103 
#10.1.1.0/24 dev enp6s0 proto kernel scope link src 10.1.1.31 metric 103 
#10.101.101.0/24 dev net0 proto kernel scope link src 10.101.101.1 linkdown 

## LOOK FOR RTNETLINK LIBRARY IN GO TO CONVERT THIS

sudo ip route flush all
sudo ip route flush table net0
sudo ip route flush table net1
sudo ip route flush table net2
sudo ip route flush table net3

# Network0 
#sudo ip route add 10.101.101.0/24 dev net0 proto kernel scope link src 10.101.101.1
#sudo ip route add 10.1.1.0/24 dev enp6s0 proto kernel scope link src 10.1.1.254 dev enp6s0 metric 100
#sudo ip route add default via 10.1.1.1 dev enp6s0 proto static metric 100                            
#
### Network1
#sudo ip route add 10.106.106.0/24 dev net1 proto kernel scope link src 10.106.106.1
#sudo ip route add 10.6.6.0/24 dev enp18s0 proto kernel scope link src 10.6.6.254 dev enp18s0 metric 104
#sudo ip route add default via 10.6.6.1 dev enp18s0 proto static metric 104                             
#
### Network2
#sudo ip route add 10.111.111.0/24 dev net2 proto kernel scope link src 10.111.111.1
#sudo ip route add 10.11.11.0/24 dev enp2s0 proto kernel scope link src 10.11.11.254 dev enp2s0 metric 102
#sudo ip route add default via 10.11.11.1 dev enp2s0 proto static metric 102                             
#
### Network3
#sudo ip route add 10.116.116.0/24 dev net3 proto kernel scope link src 10.116.116.1
#sudo ip route add 10.16.16.0/24 dev enp1s0 proto kernel scope link src 10.16.16.254 dev enp1s0 metric 103
#sudo ip route add default via 10.16.16.1 dev enp1s0 proto static metric 103                              



echo "Route Table Builder"
echo "==========================================================="
echo "Building routing tables for each of the four conniections"
echo "==========================================================="
# net0
#==============================================================================
#cat "1 net0" >> /etc/iproute2/rt_tables
# Establish the default route for the galaxy routing table
echo "Setting up network0 routing table..."
echo "==========================================================="
sudo ip rule add iif enp6s0 table net0
sudo ip rule add iif net0 table net0
sudo ip rule add from 10.101.101.0/24 table net0
sudo ip rule add from 10.1.1.0/24 table net0

sudo ip route add 10.101.101.0/24 dev net0 proto kernel scope link src 10.101.101.1 table net0
sudo ip route add 10.1.1.0/24 dev enp6s0 proto kernel scope link src 10.1.1.254 dev enp6s0 metric 100 table net0
sudo ip route add default via 10.1.1.1 dev enp6s0 proto static metric 100 table net0
echo "==========================================================="
# net1
#==============================================================================
##cat "2 net1" >> /etc/iproute2/rt_tables
## Establish the default route for the galaxy routing table
echo "Setting up network1 routing table..."
echo "==========================================================="
sudo ip rule add iif enp18s0 table net1
sudo ip rule add iif net1 table net1
sudo ip rule add from 10.106.106.0/24 table net1
sudo ip rule add from 10.6.6.0/24 table net1

sudo ip route add 10.106.106.0/24 dev net1 proto kernel scope link src 10.106.106.1 table net1
sudo ip route add 10.6.6.0/24 dev enp18s0 proto kernel scope link src 10.6.6.254 dev enp18s0 metric 100 table net1
sudo ip route add default via 10.6.6.1 dev enp18s0 proto static metric 100 table net1
echo "==========================================================="
# net2
#==============================================================================
##cat "3 net2" >> /etc/iproute2/rt_tables
## Establish the default route for the galaxy routing table
echo "Setting up network2 routing table..."
echo "==========================================================="
#sudo ip rule add iif enp2s0 table net2
#sudo ip rule add iif net2 table net2
sudo ip rule add from 10.111.111.0/24 table net2
sudo ip rule add to 10.11.11.0/24 table net2

sudo ip route add 10.111.111.0/24 dev net2 proto kernel scope link src 10.111.111.1 table net2
sudo ip route add 10.11.11.0/24 dev enp2s0 proto kernel scope link src 10.11.11.254 table net2
sudo ip route add default via 10.11.11.1 dev enp2s0 proto static metric 100 table net2
echo "==========================================================="
# net3
#==============================================================================
##cat "4 net3" >> /etc/iproute2/rt_tables
## Establish the default route for the galaxy routing table
echo "Setting up network3 routing table..."
echo "==========================================================="
sudo ip rule add iif enp1s0 table net3
sudo ip rule add iif net3 table net3
sudo ip rule add from 10.116.116.0/24 table net3
sudo ip rule add from 10.16.16.0/24 table net3


sudo ip route add 10.116.116.0/24 dev net3 proto kernel scope link src 10.116.116.1 table net3
sudo ip route add 10.16.16.0/24 dev enp1s0 proto kernel scope link src 10.16.16.254 table net3
sudo ip route add default via 10.16.16.1 dev enp1s0 proto static metric 100 table net3
echo "==========================================================="
echo "ROUTING TABLES UPDATED!"
echo "==========================================================="

#echo "Adding masquering iptables post route" 
sudo iptables -t nat -A  POSTROUTING -o net2 -j MASQUERADE
#sudo iptables -t nat -A  POSTROUTING -o enp2s0 -j MASQUERADE
#
