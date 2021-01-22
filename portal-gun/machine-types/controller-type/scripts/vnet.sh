#!/bin/sh
###############################################################################
## TASKS

#  * After we solve this problem, we can generate *.conf files and install 
#    them in '/etc/iproute2/rt_tables.d/*.conf'

###############################################################################
DRY_RUN=1

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
    echo "[mvip] Dry run, command was not executed."
  else
    eval "$1"
  fi
}
###############################################################################
echo "==[ Multiverse OS: Multiverse Routing Table(s) Builder]================="
echo "------------------------------------------------------------------------"
echo "[mvip] Clearing default ip routing table..."
run "sudo ip route flush local"
run "sudo ip route flush all"
###############################################################################
# Building routing tables                                                     #
echo ".=========================================================."            #
echo "| Building routing tables for each conniection...         |"            #
echo ".=========================================================."
echo " Network 1"
echo "'---------------------------------------------------------'"
net1_controller_device="enp20s0"
net1_controller_ip="10.6.6.2"
net1_controller_subnet="10.6.6.0/24"
net1_controller_gateway="10.6.6.1"

net1_app_device="net1"
net1_app_gateway="10.106.106.1"
net1_app_subnet="10.106.106.0/24"
net1_rtable=$net1_app_device

# This an instance that really shows the problems with using sh for portability.
if [ -z $(grep $net1_rtable /etc/iproute2/rt_tables) ]; then
   echo "[vmip] '/etc/iproute2/rt_tables' file is missing entry for '$net1_o'"
   # TODO: Its not easy to just assign this because we need to read all the 
   #       assigned numbers for we can assign a number (really should not 
   #       have to do that, its unncessarily more difficult and priority 
   #       can override it anyways. ID should be determined by the database). 
   #cat "2 $net1_rtable" >> /etc/iproute2/rt_tables
fi
echo "[mvip] Setting up routing table '$net1_table'..." 
echo "'--[ BEFORE TABLE STATE ]---------------------------------'"
echo "[mvip] DEFAULT Routing Table BEFORE additions"
run "sudo ip route show"
echo "'---------------------------------------------------------'"
echo "[mvip] TABLE[$net1_rtable] Routing Table BEFORE additions"
run "sudo ip route show table $net1_rtable"
echo "'========================================================='"
echo ".........................................................."
echo "[mvip] Flushing IP Route TABLE[$net1_rtable]"
run "sudo ip route flush table $net1_rtable"
echo ".........................................................."
echo "[mvip] Building Routing TABLE[$net1_rtable], three (3) routes:"
run "sudo ip rule add from $net1_app_subnet dev $net1_app_device table $net1_rtable"
run "sudo ip rule add  $net1_app_gateway table $net1_rtable"
run "sudo ip rule add from $net1_app_gateway table $net1_rtable"
echo ".........................................................."
echo "[mvip] Building Routing TABLE[DEFAULT], three (3) routes:"
echo "Default route for network 1 -> 10.1.1.1"
run "sudo ip route add default via $net1_controller_gateway dev $net1_controller_device proto static table $net1_rtable"  # Error command not finsiehd
run "sudo ip route add $net1_controller_subnet dev $net1_controller_device proto kernel scope link src $net1_controller_ip dev $net1_controller_device table $net1_rtable"
run "sudo ip route add $net1_app_subnet dev $net1_app_device proto kernel scope link src $net1_app_gateway table $net1_rtable"
echo "..........................................................."
echo "'==[ RESULT(s) ]=========================================='"
echo "[mvip] Routuing TABLE[DEFAULT] AFTER additions"
run "sudo ip route show"
echo "[mvip] Routing TABLE[$net1_rtable] AFTER additions"
run "sudo ip route show"
echo "'---------------------------------------------------------'"



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
#echo "WE should be able to accomplish above task without the need"
#echo "for using IPtables."
#echo "Adding masquering iptables post route" 
#sudo iptables -t nat -A  POSTROUTING -o net2 -j MASQUERADE
#echo ".=========================================================."
#echo "| Sucessfully installed IP table rules!                   |"
#echo "'---------------------------------------------------------'"
