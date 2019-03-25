# Star Router VM
The primary function of the **Star Router VM** is transparently proxy all Multiverse OS cluster traffic over a Tor after going over the Galaxy Router's VPN. The result is all traffic being sent over both the combination of VPN and Tor, concealing usage of Tor from LAN and ISP snooping.

Eventually the Star Router will provide complex and rich control over the Tor connection very similar to `Selektor` functionality but with an API consistent with other Multiverse OS APIs and simplified scripting system in Ruby.

By default all routers are ephemeral, and changes will require updating the template VM.  

#### Feature Brainstorming
Below are a list of potential features that can be implemented to improve and extend the functionality of the Galaxy Router VM.

  * Establish and maintain several different Tor connections, tracking location of exit node and speed. Allowing traffic to be routed over different tunnels.

  * Scriptable filtering, and routing over multiple proxy endpoints
    *(For example, if connecting to a remote server and remote server is in x country, use Tor exit node nearest remote server, so if the remote server is in NYC, use a NYC Tor exit node to connect or at least one geographically close).*

