# Galaxy Router VM
The primary function of the **Galaxy Router VM** is transparently proxy all Multiverse OS cluster traffic over a VPN. This prevents LAN attacks, packet sniffing and other basic attacks. Eventaully the Galaxy Router VM will also function to connect various Multiverse OS hosts together to combine several physical bare-metal host machines and pool their resources in way that is presented to the end user as interacting with a single computer.  

The Galaxy Router VM requires you to provide an OpenVPN configuration file (\*.ovpn) to transparently route all traffic over.

In the future, the Galaxy Router VM will support a wide variety of other options including but not limited to: SSH tunnels, ICMP tunnels, DNS tunnels, and many more off-site proxying options. 

By default all routers are ephemeral, and changes will require updating the template VM.  

#### Feature Brainstorming
Below are a list of potential features that can be implemented to improve and extend the functionality of the Galaxy Router VM.

  * Simple scripts to simplify the process of deploying OpenVPN, or SSH proxies on VPS hosting

  * SSH based proxying

  * Support for multiple proxy options active and spreading traffic across the multiple proxies to improve psuedo-anonymity

  * Scriptable filtering, and routing over multiple proxy endpoints
    *(For example, if connecting to a remote server and remote server is in x country, use VPN nearest proxy end-point, so if the remote server is in NYC, use the NYC VPN to connect. This would allow transparent and automatic bypassing of geographic based restrictions. Such as, restrictions to watch BBC streaming shows.)* 

