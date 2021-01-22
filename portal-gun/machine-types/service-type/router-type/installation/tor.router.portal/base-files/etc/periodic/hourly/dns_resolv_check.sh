#!/bin/ash                                                                   
###############################################################################



TUN_FILE=/sys/class/net/tun0/carrier
TUN_DIR=/sys/class/net/tun0

if [ ! -f "$TUN_FILE" -a ! -d "$TUN_DIR" ]; then
	echo "Tunnel does NOT exist!"
else
	echo "Tunnel tun0 exists..."

	exit 0
fi

ping -c 1 www.debian.org > /dev/null
STATUS=$?
if [ $STATUS -eq 0 ]; then
	echo "Failed to ping www.debian.org, rebuilding resolv.conf..."
	echo "Then restarting the openvpn service..."
	
	echo "nameserver 1.1.1.1" >  /etc/resolv.conf
	echo "nameserver 9.9.9.9" >> /etc/resolv.conf
	echo "nameserver 8.8.8.8" >> /etc/resolv.conf
	echo "nameserver 8.8.4.4" >> /etc/resolv.conf
	
	sleep 3
	
	service openvpn restart
fi
