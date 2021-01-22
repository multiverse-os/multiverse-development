# Dev Notes


Host controller combined for low resource boxes. (Different operating modes) 

* Update/generate shorewall configs basded on template and output of "route" or "ip route" commands. In addition to the ifconfig or ip a commands to ensure the devices and zones are based on actual machine information, configureation, instead of hard coding. Right now its really dumb. 

On "cannonical" operating mode, intel i9 eth0 and eth1 were flipped from the new host+controller operating mode. So the above is necessary soon, at least before the installer. 
