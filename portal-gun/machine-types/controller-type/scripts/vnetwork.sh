#!/bin/sh
###############################################################################
DRY_RUN=0

###############################################################################
## TASKS
#
## TODO: FIND GO LIBRARIES (RUBY?) 
##         * RTNETLINK Go library
##       RNLINK helps with some but we need device itartion, device creation 
##       ip tables control (nftables?), userspace packet routing.
#
## TODO: NETWORK BRIDGES
##       Add bridges here fuck using virsh net-edit, its not good and is
##       absolutely useless for our usecase in Multiverse OS.
#
## TODO: NETWORK DEVICES
##       Should be looping through but fuckt hat in shell get thkis work and fuck
##       off. && Do build variables based on system ASAP.
#
###############################################################################
## Functions
run(){ # 1=Command
  echo "$1"
  if [ $DRY_RUN -eq 1 ]; then 
    echo "[mvip] dry run, command was not executed."
  else
    eval "$1"
  fi
}
###############################################################################
echo " [_ Multiverse OS: Multiverse Routing Table(s) Builder _]"
echo "================================================================================"
echo "[mvip] Clearing default ip routing table..."
sudo ip route flush all

###############################################################################
# Building routing tables                                                     #
echo ".=========================================================."            #
echo "| Building routing tables for each conniection...         |"            #
echo ".=========================================================."            #
###############################################################################
#echo ""
#echo " Network 0"
#net0_controller_device="enp8s0"
#net0_controller_ip="10.1.1.2"
#net0_controller_subnet="10.1.1.0/24"
#net0_controller_gateway="10.1.1.1"
#
#net0_app_device="net0"
#net0_app_ip="10.101.101.1"
#net0_app_subnet="10.101.101.0/24"
#net0_rtable=$net0_app_device
## if [ -z $(grep net0 /etc/iprout2/rt_tables ]; then
##   cat "1 net0" >> /etc/iproute2/rt_tables
## fi
#echo "Setting up routing table '$net0_rtable'..." 
#echo "'---------------------------------------------------------'"
#run "sudo ip route flush table $net0_rtable"
#
##run "sudo ip route add $net0_app_subnet dev $net0_app_device proto kernel scope link src $net0_app_gateway"
#run "sudo ip route add $net0_controller_subnet dev $net0_controller_device proto kernel scope link src $net0_controller_ip dev $net0_controller_device metric 100"
#run "sudo ip route add default via $net0_controller_gateway dev $net0_controller_device uproto static metric 100"
#
## Tables Rules (Use table if...) 
#run "sudo ip rule add from $net0_app_gateway table $net0_rtable"
#run "sudo ip rule add to $net0_app_gateway table $net0_rtable"
#run "sudo ip rule add from $net0_app_subnet table $net0_rtable"
## Routing Table from app network to second 
##run "sudo ip route add default via $net1_controller_gateway dev $net1_controller_device proto static table $net0_rtable"
#run "sudo ip route add $net0_controller_subnet dev $net0_controller_device proto kernel scope link src $net0_controller_ip dev $net0_controller_device table $net0_rtable"
#run "sudo ip route add $net0_app_subnet dev $net0_app_device proto kernel scope link src $net0_app_gateway table $net0_rtable"



echo ".=========================================================."
echo " Network 1"
net1_controller_device="enp20s0"
net1_controller_ip="10.6.6.2"
net1_controller_subnet="10.6.6.0/24"
net1_controller_gateway="10.6.6.1"

net1_app_device="net1"
net1_app_gateway="10.106.106.1"
net1_app_subnet="10.106.106.0/24"
net1_rtable=$net1_app_device
# if [ -z $(grep net1 /etc/iprout2/rt_tables ]; then
#   cat "2 net1" >> /etc/iproute2/rt_tables
# fi
echo "Setting up routing table '$net1_table'..." 
echo "'---------------------------------------------------------'"
run "sudo ip route flush table $net1_rtable"

# Tables Rules (Use table if...) 

echo "--[Table Filter Rules]-------------------------------------"
run "sudo ip rule add from $net1_app_subnet dev $net1_app_device table $net1_rtable"
run "sudo ip rule add  $net1_app_gateway table $net1_rtable"
run "sudo ip rule add from $net1_app_gateway table $net1_rtable"
#
#run "sudo ip rule add to $net1_app_gateway table $net1_rtable"
echo "-----------------------------------------------------------"


echo "..........................................................."
echo "DEFAULT --> TABLE"
run "sudo ip route add default via $net1_controller_gateway dev $net1_controller_device proto static table $net1_rtable"  # Error command not finsiehd
run "sudo ip route add $net1_controller_subnet dev $net1_controller_device proto kernel scope link src $net1_controller_ip dev $net1_controller_device table $net1_rtable"
#run "sudo ip route add $net1_app_subnet dev $net1_app_device proto kernel scope link src $net1_app_gateway table $net1_rtable"
echo "--RESULT---------------------------------------------------"
run "sudo ip route"
echo "..........................................................."

echo "..........................................................."
echo "--> $net1_rtable ---> TABLE --> "
#run "sudo ip route add default via $net1_controller_gateway dev $net1_controller_device proto static metric 102"
#run "sudo ip route add $net1_controller_subnet dev $net1_controller_device proto kernel scope link src $net1_controller_ip dev $net1_controller_device metric 102"
run "sudo ip route add $net1_app_subnet dev $net1_app_device proto kernel scope link src $net1_app_gateway" # Routing Table from app network to second 
echo "--RESULT---------------------------------------------------"
run "sudo ip route show table $net1_rtable"
echo "..........................................................."



#echo ".=========================================================."
#echo " Network 2"
#net2_controller_device="enp1s0"
#net2_controller_ip="10.11.11.2"
#net2_controller_subnet="10.11.11.0/24"
#net2_controller_gateway="10.11.11.1"
#
#net2_app_device="net2"
#net2_app_ip="10.111.111.1"
#net2_app_subnet="10.111.111.0/24"
#
#net2_rtable=$net2_app_device
## if [ -z $(grep $net2_rtable /etc/iprout2/rt_tables ]; then
##   cat "3 $net2_rtable" >> /etc/iproute2/rt_tables 
## fi
#echo "Setting up routing table '$net2_rtable'..."
#echo "'---------------------------------------------------------'"
#run "sudo ip route flush table $net2_rtable"
#
##run "sudo ip route add $net2_app_subnet dev $net2_app_device proto kernel scope link src $net2_app_gateway"
#run "sudo ip route add $net2_controller_subnet dev $net2_controller_device proto kernel scope link src $net2_controller_ip dev $net2_controller_device metric 104"
#run "sudo ip route add default via $net2_controller_gateway dev $net2_controller_device proto static metric 104"
#
#echo "..........................................................."
#echo "Default Routing Table"
#run "sudo ip route"
#echo "..........................................................."
#
#echo "--[Table Filter Rules]-------------------------------------"
## Tables Rules (Use table if...) 
##run "sudo ip rule add from $net2_app_gateway table $net2_rtable"
#run "sudo ip rule add to $net2_app_gateway table $net2_rtable"
#run "sudo ip rule add from $net2_app_subnet table $net2_rtable"
#echo "'---------------------------------------------------------'"
#
#
## Routing Table from app network to second 
##run "sudo ip route add default via $net1_controller_gateway dev $net1_controller_device proto static table $net0_rtable"
#run "sudo ip route add $net2_controller_subnet dev $net2_controller_device proto kernel scope link src $net2_controller_ip dev $net2_controller_device table $net2_rtable"
#run "sudo ip route add $net2_app_subnet dev $net2_app_device proto kernel scope link src $net2_app_gateway table $net2_rtable"
#
#echo "..........................................................."
#echo "[__ $net2_rtable __ Routing Table]"
#run "sudo ip route show table $net2_rtable"
#echo "..........................................................."
#
#echo ".=========================================================."
#echo " Network 3"
#net3_controller_device="enp2s0"
#net3_controller_ip="10.16.16.2"
#net3_controller_subnet="10.16.16.0/24"
#net3_controller_gateway="10.16.16.1"
#
#net3_app_device="net3"
#net3_app_gateway="10.116.116.1"
#net3_app_subnet="10.116.116.0/24"
#
#net3_rtable=$net3_app_device
## if [ -z $(grep net3 /etc/iprout2/rt_tables ]; then
##   cat "4 net3" >> /etc/iproute2/rt_tables
## fi
#echo "Setting up routing table '$net3_rtable'..." 
#echo "'========================================================='"
#run "sudo ip route flush table $net3_rtable"
#
##run "sudo ip route add $net3_app_subnet dev $net3_app_device proto kernel scope link src $net3_app_gateway"
#run "sudo ip route add $net3_controller_subnet dev $net3_controller_device proto kernel scope link src $net3_controller_ip dev $net3_controller_device metric 106"
#run "sudo ip route add default via $net3_controller_gateway dev $net3_controller_device proto static metric 106"
#
#echo "..........................................................."
#echo "Default Routing Table"
#run "sudo ip route"
#echo "..........................................................."
#
## Tables Rules (Use table if...) 
#run "sudo ip rule add from $net3_app_subnet table $net3_rtable"
#run "sudo ip rule add from $net3_controller_gateway $net3_rtable"
#run "sudo ip rule add to $net3_controller_gateway table $net3_rtable"
## Routing Table from app network to second 
#run "sudo ip route add $net3_app_subnet dev $net3_app_device proto kernel scope link src $net3_app_gateway table $net3_rtable"
#run "sudo ip route add $net3_controller_subnet dev $net3_controller_device proto kernel scope link src $net3_controller_ip dev $net3_controller_device table $net3_rtable"
#run "sudo ip route add default via $net3_controller_gateway dev $net3_controller_device proto static table $net3_rtable"
#
#echo "..........................................................."
#echo "[__ $net3_rtable __ Routing Table]"
#run "sudo ip route show table $net3_rtable"
#echo "..........................................................."

echo ".=========================================================."
echo "| Sucessfully installed routing tables!                   |"
echo "'---------------------------------------------------------'"
###############################################################################
# Adding IP Tables Rules                                                      #
###############################################################################
#echo ""
#echo ".=========================================================."
#echo "| _IP Tables_                                             |"
#echo "| Setting up ip tables for Controller VM...               |"
#echo "'---------------------------------------------------------'"
#echo ""
#echo "Adding masquering iptables post route" 
#sudo iptables -t nat -A  POSTROUTING -o net2 -j MASQUERADE
#eval "sudo iptables -t nat -A POSTROUTING -o $net0_app_device -j MASQUERADE"
#eval "sudo iptables -t nat -A POSTROUTING -o $net1_app_device -j MASQUERADE"
#eval "sudo iptables -t nat -A POSTROUTING -o $net2_app_device -j MASQUERADE"
#eval "sudo iptables -t nat -A POSTROUTING -o $net3_app_device -j MASQUERADE"
#echo "Adding masquering iptables post route" 
#sudo iptables -t nat -A  POSTROUTING -o enp2s0 -j MASQUERADE
#eval "sudo iptables -t nat -A POSTROUTING -o $net0_controller_device -j MASQUERADE"
#eval "sudo iptables -t nat -A POSTROUTING -o $net1_controller_device -j MASQUERADE"
#eval "sudo iptables -t nat -A POSTROUTING -o $net2_controller_device -j MASQUERADE"
#eval "sudo iptables -t nat -A POSTROUTING -o $net3_controller_device -j MASQUERADE"
#echo ".=========================================================."
#echo "| Sucessfully installed IP table rules!                   |"
#echo "'---------------------------------------------------------'"
