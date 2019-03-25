# Rebuilding system to use Q35
When rebuilding every router on the host machine (turned out to be easiest 
approach to fixing the machines), I ran into an issue with spice.

The solution was defining the spice address as IPv6 instead of IPv4 (default). 

The essentially loopback in IPv6 format is `::1` and this is used for the
spice server address. After making this switch everything should work fine.
