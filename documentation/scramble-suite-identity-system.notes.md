##
##  Multiverse OS: Scramble Suite Identity System
##
##  Onion-routed P2P idenity system, supporting multiple nested isolated identiteis, peer based recovery and backup, all relevant
##  key types (GPG, SSH,. ECDSA, Bitcoin, Ethereum Classic, Tor v3, x509, etc), using epeheraml tree system for temporal session
##  keys, decentralized authentication and revocation of entire branches.
======================================================================================================================================



**UI**

  [*][Window] 

    Sunburst graph to navigate the keys. Perhas in addition to some other tree/mycelia map, so you can visualize yourself in two
    different dimesions, which provides the depthx. 

   Use visual flashes, window flashing. for indicators.





## Legacy systems that need to be replaced
_The keys should be stored in a new database. It should be easy to explort the key you need for any CLI task, for example you try to SSH, if contenxt does not already select the key (i.e. server, account, remote server, project scope, etc)

If not, maybe key menu drops down. 

Key system should come with pluigins for major browsers for seamless no-password intergration WITHOUT JS!!!



ssh-keygen -C user@host -t ed25519


SHA256:Q5CoHpWeyPJzm1SRjLKCCQan0y5eBNeKrK1L28JjOL0 user@host

+--[ED25519 256]--+
|o....=.o         |
|.=+ =.=.         |
|BooO.. ..        |
|==B.o ..         |
|o*.o .  S        |
|o.* o    .       |
|o= + o           |
|==+ o            |
|o+Eo             |
+----[SHA256]-----+


-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACDg+DcAzvBczGR+ZyliMKjKYl15/RHXIQhI4KihzsfyUAAAAJCTY8iTk2PI
kwAAAAtzc2gtZWQyNTUxOQAAACDg+DcAzvBczGR+ZyliMKjKYl15/RHXIQhI4KihzsfyUA
AAAECgeP7CSd5pItkRyah9742KzJFGWwYx7ufUjsB9UtGVSeD4NwDO8FzMZH5nKWIwqMpi
XXn9EdchCEjgqKHOx/JQAAAACXVzZXJAaG9zdAECAwQ=
-----END OPENSSH PRIVATE KEY-----



ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOD4NwDO8FzMZH5nKWIwqMpiXXn9EdchCEjgqKHOx/JQ user@host






