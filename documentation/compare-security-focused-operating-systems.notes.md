## Existing Secure OS Options
With no intention to downplay or ignore the long list of acheivements and projects created by developers working on hardended, secure and psuedoanonymous versions of linux, including the projects that directly led to creation Tails OS and Whonix, to simplify analysis of existing solutions our list will beegin with Tails OS and Whonix.

*Despite these operating systems may be older, they are stable, tested and trusted by security researchers, activists and journalists. If you are looking for the most secure OS solutions available, Tails OS and Whonix OS should considered first.*

New is not always better, Multiverse OS suffers from lack of testing, lack of peer review and expert analysis. Existing software attempting to solve the same problems must be evaluated so successful features can be incorporated into Multiverse OS.

### Tails OS

Tails OS or amnesiac linux, the operating system reasonably deserves at least some credit for popularizing enhancing the security using emphemeral root file system to accomplished, with an optional encrypted persistent storage. Tails OS running the operating system in ram never saves any changes to the root system and resets after each reboot.

#### Positive


The primary security advantages of using an emphemeral system include:
*  Chat history, browsing history, logs, and other local records are purged at shutdown.
*  Rootkits, torjans and remote administration tools can only remain installed for a single session because they are purged at shutdown.
*  An empty chat history, browsing history, logs and other local records can not be used to create a unique identifier for tracking between sessions.
*  A USB kill script can be automaticaly run, to wipe the RAM and shutdown, if the USB key with Tails OS is removed.

Combined with a firewall that routes **all** traffic through what in essence is an isolating Tor proxy makes Tais OS capable of providing a secure work environment, that is psuedoanonymous and capable of preventing some forms of tracking.

Tails OS comes with secure defaults, configurations and packages to increase anonymity and security.

Easy to use.

#### Issues

The emphemeral linux is by design difficult to customize, any additional software must be installed every time the OS is booted. In complex development environments the booting process slows down.

The firewall method of routing traffic through Tor is not as secure as using a separate router for isolating and proxying traffic through Tor.

Without compartmentalization built into the operating system, a user must reboot to change activities. Otherwise any leak or unauthorized access can link all active activity together.

#### Whonix OS

#### Positive Aspects

Isolating proxy using a separate dedicated router VM preventing even a user with root access on the workstation from leaking connection details. 

Can be run entirely in VMs making it possible to compartmentalize activities. Each piece can be run on bare-metal hosts for increased flexibility.

Can be incorporated into clusters easily.

Whonix comes with secure defaults, configurations and packages to increase anonymity and security.

Easy to use.

#### Issues

### Subgraph OS
