# Multiverse OS Notes
A running log of notes, thoughts, ideas, that will be refined and broken up into more narrow topics for later use as documentation and development guides. 

* Gnome/WM feature to see WAN IP of a given ID or browser, including region information and other information about the IP address (IE flagged as tor exit node, flagged as suspicious, spammer, VPN, etc). Also the ability to see multiple hops if using Tor or multi-hop VPN.

# SSH Server 

[Auth]

https://github.com/simmel/oath-ssh-command 

[Limited Jail]
A limited jail server for remote control with limtied control over the machine

# Key Server / Password Manager

SSH https://github.com/deamwork/otp-ssh

https://github.com/cromega/keyguard - served ssh keys over authenticated endpoint
* Password Manager 
https://github.com/Bridouille/go-otp-manager - has a limited webui
  - OTP Support
    https://github.com/tschuy/gotp
https://github.com/adminfromhell/otp
    https://github.com/pquerna/otp
    https://github.com/zesik/otp
    https://github.com/fonglh/go-totp/blob/master/totp.go
  - Identity Support
  - Epehemeral Keys

https://github.com/skyjia/go-otpserver - OTP JSON API server

# Software Needed

* VM Manager
  - WebUI
  - Full VM, NoVM, Clear Containers
  - Custom BiOS
  - Epehemeral Machines
  - Custom network stack
  - Router
  - Networking
  - Versioned Checkpointing VM harddrives 
  - VM Types:
    - Controller VM
      * CPU Pinning
      * Fallback VM (if VM fails to boot, then go to always working simple VM with controls to get shit working
    - Application VM
    - Utility VM





* OHT

# Multiverse OS

* Create a Onion device, any packets sent to this will be torified, one can recieve packets from a packet queue over this device. If one has root access, one can create a device, otherwise a bridge can be used with a socket to provide similar functionaltiy in userspace (preferred).

* Use NoVM or Libvirt/KVM to isolate and use immutable drives for applications. Build each machine from scratch, do not rely on others like docker. Use full virtualization for actual security and real flexibility (most of the speed gains are during boot). Pin CPUs and autoconfig to best settings.

* Use the onion network to create a call home service, to access your device if it connects to the internet even if behind a NAT.

* Finish the chat example

* Use the chat system to manage a cluster. Status reporting, etc

* Shamirs secret sharing, with contacts, breaking up and backuping up all important keys, offsite and in the cojoined/layered-custody of friends and family.

## Setting up OpenVPN and IPtables to ensure all connections are sent over VPN 
https://www.linode.com/docs/networking/vpn/set-up-a-hardened-openvpn-server
