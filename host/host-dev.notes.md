# Host Development Notes & Research
===============================================================================
The next steps of Multiverse OS host development involve compiling the the 
most recently developed shell (`/bin/sh`) scripts and finalized manual install
guides into a finalized shell script that has removed all `libvirt` references
and dependencies in favor of a direct QEMU/NEMU+KVM based system that uses 
Multiverse OS custom YAML (eventually a Ruby based configuration) that defines
all the QEMU command-line options in a simpler, much easier to understand 
configuration with secure defaults and options that are between insecure and
secure are removed to always favor the more secure, more preformant options.

Leaving only the configuration options which make sense for the user to decide.
In addition, automating important aspects that are tedious to manage such as
CPU pinning based on the Multiverse OS class and the passed PCI devices. So
that the user achieves maximum fully virtualized preformance and more of the
configuration can be dedicated the extremely low level provisioning, that 
defines the base level of the ephemeral Multiverse OS virtual machines. 

## Multiverse OS Alpha Installer
In addition to removing the `libvirt` depeendency and directly using QEMU/NEMU
the alpha installer provide a CLI/TUI based installation, and a GUI based
installation utilziing webframe in combination with webview/webcomponent based
UI; providing GUI based installation for both live media and dedicated 
installation media. 


