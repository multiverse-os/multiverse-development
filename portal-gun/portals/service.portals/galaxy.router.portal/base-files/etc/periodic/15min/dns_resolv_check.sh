#!/bin/ash
ping -c 1 www.debian.org > /dev/null
STATUS=$?
if [ $STATUS -eq 0 ]; then
	exit 0
else
	echo "nameserver 1.1.1.1" > /etc/resolv.conf
	service openvpn restart
fi
