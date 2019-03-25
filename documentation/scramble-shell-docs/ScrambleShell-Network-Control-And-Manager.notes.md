##
##  ScrambleShell: Networking
##
=================================================================
Review Selector and software like Littlesntich to get an idea of 
what we would like to provide the Multiverse OS user. 


Like with Selector we want to control what guard nodes and what
paths we are using, in an intelligent way, that ensures fast
connections and **most importantly, DOES ROUTING BASED ON**
**PATTERNS!**

=================================================================
## Tor Routing Via Patterns
-----------------------------------------------------------------
Pattern programming should be done using the Ruby programming 
language, in order to provide incredibly rich ability to filter
out information and route to different proxies. These should be
done in chunks or separate files (but able to be viewed in
a single file to make it easy to copy working filters/patterns).

Each PATTERN BLOCK upon saving is split into its own file, 
using the number system used by most linux config [5-pattern.rb,
10-patter.rb, 55-pattern.rb,.. and so on]. 

#### [__/!\__NOTE__: This should not just be for Tor] ####
	  Pattern Based Routing Using Ruby Based Patterns Is
	  THE FOUNDATION for Scramble Suit Identity System:
			(1) VPN (OpenVPN) traffic
			(2) SSH Proxy (SINGLE) and (MULTI-HOP)
			(3) and general traffic
		 This is a essentially THE FOUNDATION of the Scramble Shell
		 identity segregation system and without this everything else
		 falls apart 

=================================================================
## Networking Applet / Network Manager (GNOME)


[gnome-nettool](https://github.com/GNOME/gnome-nettool)
**gnome-nettool** GNOME interface for various networking tools. GNOME Net tool is a set of front-ends to various networking command-line
[GUIs:][tools, like ping, netstat, ifconfig, whois, traceroute, finger]




[network-manager-applet](https://github.com/GNOME/network-manager-applet)
network-manager-applet (mirrored from: http://projects.gnome.org/NetworkManager/)

http://www.linuxfromscratch.org/blfs/view/8.1/basicnet/glib-networking.html
 Introduction to GLib Networking
The GLib Networking package contains Network related gio modules for GLib.
This package is known to build and work properly using an LFS-8.1 platform. 
  **Optional**
		libproxy 




https://github.com/GNOME/network-manager-openvpn
OpenVPN support for NetworkManager

Added by Tim Niemueller http://www.niemueller.de

Assumes that you have a running OpenVPN X.509 setup as mentioned
in the OpenVPN 2.0 HOWTO on the OpenVPN homepage.

TODO: Support for static keys, support for password authentication,
only present auth-dialog if needed.





=================================================================
## Gnome WAN IP Address & Geographic Information
-----------------------------------------------------------------
A Gnome Extension UI drop down that provides succinct rich 
connection information. 

In Multiverse OS, Scramble Shell maintains more than one internet
identity, part of this is BRANCHING connections.

*Clicking should switch between branches maybe slide at the top*
*with the identity name, showing different end point.*

  The drop down should show the chain 
  [!] [>][LAN IP]
   |
   |__[!] [>][Gateway IP Address] [+][>][WAN IP Address]
      |
      |___[>][VPN IP Address]
			|	  |
			|	  |___[>][Entrance][>][Middle][>][Tor Exit] 
      |
  *[other branch]* *Clicking the other branch should let you*
  *navigate around your connection tree*
      
   [Open Street Maps]_Show_the_location_of_the_IP_Address
   |                                                    |
   |                                                    |
   |  ________________________________________________  |
   |  |                                       /       | |
   |  |            __          ___           /        | |
   |  |           (  )        ( o )   \     /         | |
   |  |            \/          \ /     x  --(         | |
   |  |                      /  V            \        | |
   |  |                             /                 | |
   |  |                                      /        | |
   |  |_____________________________________/_________| |
   |____________________________________________________|
      
    *Notes*: This concept of a ASCII map is a not terrible
    perhaps a version can be done in ASCII to serve over
    SSH or in terminal.


	[Multiverse OS:Connection Tree]
    In Multiverse there is a branching connection tree, or 
    a wormhole network (?).
    
    An important note, is EACH **Application VM** in Multiverse
    has its own branch, and so perhaps in the window relaying
    the application can provide access this drop-down IP
    information.
    
    Regardless, we need to break down connections:
  			 |                                       |
         |  <<     Scramble Suit ID Name     >>  |
  			 |                                       |
         |       <<   Application VM   >>        | (This one 
  			 |                                       | could be optional 
  			 |                                       |  dependent on
         |_______________________________________| context, as in
         								   ...	 									 is it topbar
         								*more not shown*					 dropdown?)
         																					 
         																					 
         	**Note: ID and Application**
         	So here we are providing a way to do
         	broad navigation through the connection
         	tree.
         	
         	This style very easy way to to view
         	the different connections and hopefully
         	navigate the associated branches.
					
					**Note: Possible Feature** 
      		Another possible feature is a visual node map
      		just dots and lines that is maybe light gray or
      		white and overlaid over the bottom right
      		corner like channels on a TV.
      		
      		Then the the path that you have selected is
      		the one that is darkened, to show the user
      		
      		
      		              o
      		            /   \
      		           o     o
      		          / \     \
      		         o   *     *
      		        / \
                 *   *
                 


	[Similar Sample/Example Projects]
	  [ip-finder](https://extensions.gnome.org/extension/1190/ip-finder/)
	  Ip finder is essentially what we are doing but not using 
	  open street maps and not nearly complex enough to support
	  the complexity of a Multiverse OS connection

=================================================================
## 
-----------------------------------------------------------------


=================================================================
## 
-----------------------------------------------------------------


=================================================================
## 
-----------------------------------------------------------------










