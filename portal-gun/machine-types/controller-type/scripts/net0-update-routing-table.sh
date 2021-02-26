#!/bin/bash
###############################################################################
app_device="net0"
controller_device="enp3s0"
#=============================================================================#
app_subnet="10.101.101.0/24"
app_gateway="10.101.101.1"
app_ip="10.101.101.1"
#=============================================================================#
controller_subnet="10.1.1.0/24"
controller_gateway="10.1.1.1"
controller_ip="10.1.1.2"
#=============================================================================#
rtable=$app_device
###############################################################################
shell(){
	echo "$1"
	$1
}
###############################################################################
#sudo ip route flush all
#shell "sudo ip route add $app_subnet dev $app_device proto kernel scope link src $app_gateway"
#shell "sudo ip route add $controller_subnet dev $controller_device proto kernel scope link src $controller_ip dev $controller_device metric 100"
#shell "sudo ip route add default via $controller_gateway dev $controller_device uproto static metric 100"
###############################################################################
echo ".=========================================================."
echo "| Network 0                                               |"
echo "[=========================================================]"
echo "| Building routing tables for each conniection...         |"
echo "| Setting up routing table '$app_device'...                      |" 
echo "|                                                         |"
echo "'---------------------------------------------------------'"
if [[ -z $(grep "$app_device" /etc/iproute2/rt_tables) ]]; then
  sudo cat "1 $app_device" >> /etc/iproute2/rt_tables
fi
echo ".=========================================================."
echo "| IP ROUTE TABLE 'net0' Rules'                            |"
echo "[=========================================================]"
shell "sudo ip route flush table $rtable"
###############################################################################
shell "sudo ip rule add from $app_subnet table $rtable"
#shell "sudo ip rule add to $app_subnet table $rtable"
###############################################################################
shell "sudo ip rule add from $controller_subnet table $rtable"
#shell "sudo ip rule add to $controller_subnet table $rtable"
###############################################################################
shell "sudo ip route add $app_subnet dev $app_device proto kernel scope link src $app_gateway table $rtable"
shell "sudo ip route add $controller_subnet dev $controller_device proto kernel scope link src $controller_ip dev $controller_device metric 100 table $rtable"
shell "sudo ip route add default via $controller_gateway dev $controller_device proto static metric 100 table $rtable"
###############################################################################
echo ".=========================================================."
echo "| Sucessfully installed routing tables!                   |"
echo "'---------------------------------------------------------'"
echo ""
echo ".=========================================================."
echo "| IP Tables                                               |"
echo "[=========================================================]"
echo "| Setting up ip tables for Controller VM...               |"
echo "|                                                         |"
echo "'---------------------------------------------------------'"
shell "sudo iptables -t nat -A  POSTROUTING -o $app_device -j MASQUERADE"
shell "sudo iptables -t nat -A  POSTROUTING -o $controller_device -j MASQUERADE"
echo ".=========================================================."
echo "| Sucessfully installed IP table rules!                   |"
echo "'---------------------------------------------------------'"

