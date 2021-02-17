#!/bin/bash
###############################################################################
# Network 0                                                                   #
###############################################################################
$net0_app_device="net0"
#=============================================================================#
$net0_app_ip="10.101.101.1"
$net0_app_subnet="10.101.101.0/24"
#=============================================================================#
$net0_controller_device="enp3s0"
#=============================================================================#
$net0_controller_ip="10.1.1.2"
$net0_controller_subnet="10.1.1.0/24"
$net0_controller_gateway="10.1.1.1"
#=============================================================================#
$net0_rtable=$net0_app_device
#####################################################################################################################################
## Clear Old Settings If Exists
#sudo ip route flush all
sudo ip route flush table $net0_rtable
#####################################################################################################################################
# Network0 
sudo ip route add $net0_app_subnet dev $net0_app_device proto kernel scope link src $net0_app_gateway
sudo ip route add $net0_controller_subnet dev $net0_controller_device proto kernel scope link src $net0_controller_ip dev $net0_controller_device metric 100
sudo ip route add default via $net0_controller_gateway dev $net0_controller_device uproto static metric 100
#####################################################################################################################################
echo ".=========================================================."
echo "| Network 0                                               |"
echo "[=========================================================]"
echo "| Building routing tables for each conniection...         |"
echo "|                                                         |"
echo "| Setting up routing table 'net0'...                      |" 
echo "|                                                         |"
echo "|                                                         |"
echo "'---------------------------------------------------------'"
if [ -z $(grep net0 /etc/iprout2/rt_tables ]; then
  cat "1 net0" >> /etc/iproute2/rt_tables
fi


echo ".=========================================================."
echo "| IP ROUTE TABLE 'net0' Rules'                            |"
echo "'---------------------------------------------------------'\n"
sudo ip rule add from $net0_app_subnet table $net0_rtable
sudo ip rule add to $net0_app_subnet table $net0_rtable

sudo ip rule add from $net0_controller_subnet table $net0_rtable
sudo ip rule add to $net0_controller_subnet table $net0_rtable

sudo ip route add $net0_app_subnet dev $net0_app_device proto kernel scope link src $net0_app_gateway table $net0_rtable
sudo ip route add $net0_controller_subnet dev $net0_controller_device proto kernel scope link src $net0_controller_ip dev $net0_controller_device metric 100 table $net0_rtable
sudo ip route add default via $net0_controller_gateway dev $net0_controller_device proto static metric 100 table $net0_rtable
echo ".=========================================================."
echo "| Sucessfully installed routing tables!                   |"
echo "'---------------------------------------------------------'\n"
echo ""
echo ".=========================================================."
echo "| IP Tables                                               |"
echo "| Setting up ip tables for Controller VM...               |"
echo "'---------------------------------------------------------'\n"
#echo "Adding masquering iptables post route" 
sudo iptables -t nat -A  POSTROUTING -o $net0_app_device -j MASQUERADE
#echo "Adding masquering iptables post route" 
sudo iptables -t nat -A  POSTROUTING -o $net0_controller_device -j MASQUERADE
echo ".=========================================================."
echo "| Sucessfully installed IP table rules!                   |"
echo "'---------------------------------------------------------'\n"

