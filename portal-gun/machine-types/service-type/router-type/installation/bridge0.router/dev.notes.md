# Development Notes & Task List
#==============================================================================

A recent implementation of this router had started using the `/etc/local.d/*`
and `/etc/periodic/*` systems to add self-healing scripts to both fix the DNS
for galaxy router when the system reboots without a proper shutdown, which
leaves the DNS set to the VPN and pinging the hardcoded clients that have 
MAC address based leases.

Eventually it should just look at the leases assigned or do a very quick ping
scan in the subnet because due to some still unknown reason pining a client
will cause a client to instantly connect and not lag for some random amount of
time.

The pinging of other routers in the subnet and known client IP assignment
like `10.1.1.10` which is the standard primary controller IP typically caused
everything to setup much faster.

In addition, these scripts laid the groundwork for adding more complex self-
healing scripts in both startup and crontab **but these should be moved to a
daemon/agent as soon as possible anyways so any further efforts beyond quick
implementations to fix issues on any given machine should be avoided in favor
of implementimg the daemon.**

-------------------------------------------------------------------------------
