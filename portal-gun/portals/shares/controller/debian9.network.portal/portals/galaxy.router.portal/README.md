# Galaxy Router VM
The primary function of the **Galaxy Router VM** is transparently proxy all 
Multiverse OS cluster traffic over a VPN. This prevents LAN attacks, packet 
sniffing and other basic attacks. Eventaully the Galaxy Router VM will also
 function to connect various Multiverse OS hosts together to combine several 
physical bare-metal host machines and pool their resources in way that is 
presented to the end user as interacting with a single computer.

The Galaxy Router VM requires you to provide an OpenVPN configuration file 
(\*.ovpn) to transparently route all traffic over.

In the future, the Galaxy Router VM will support a wide variety of other 
options including but not limited to: SSH tunnels, ICMP tunnels, DNS tunnels,
and many more off-site proxying options. 

By default all routers are ephemeral, and changes will require updating the 
template VM.

#### Provisioning
Currently provisioning does a link from the portal, but the portal should be
used to setup multiple galaxy routers, and to scale up automatically based off
a generic portal setttings files. Then we need to `cp` instead of just `link -s`.

Need to also work out the alpine answers file, custom grub to allow sumbission
of keys, text, and so on.

The current sollution is to cp the files into a local 'base-files' then create
symbolic links from there. The advantage of this is that we can make the base-files
a git and if changes are made that are meant to be merged so that the change is 
reflected in all other iterations of this VM i.e. galaxy0, galaxy1, galaxy2, etc


#### Ephemerality
Some of the most important VMs to get to be ephemeral are the router VMs, and
to do this, we should jsut leverage the snapshots we can create of these. 
So we should snapshot and load from snapshot every boot.


#### Feature Brainstorming
Below are a list of potential features that can be implemented to improve and 
extend the functionality of the Galaxy Router VM.

  * Simple scripts to simplify the process of deploying OpenVPN, or SSH proxies 
    on VPS hosting

  * SSH based proxying

  * Support for multiple proxy options active and spreading traffic across the 
    multiple proxies to improve psuedo-anonymity

  * Scriptable filtering, and routing over multiple proxy endpoints
    *(For example, if connecting to a remote server and remote server is in x 
     country, use VPN nearest proxy end-point, so if the remote server is in 
     NYC, use the NYC VPN to connect. This would allow transparent and 
     automatic bypassing of geographic based restrictions. Such as, 
     restrictions to watch BBC streaming shows.)* 


#### TODO
In the future, we want to change the name of `config` to base-files to match 
other similar systems. Whenever possible, lets use existing naming conventions 
to make it easier to understand and apply existing knowledge  to this system.


  * Need a script to ensure that on boot the DNS in `/etc/resolv.conf` is back to 8.8.4.4 or something else otherwise it will requiire manual 
  
