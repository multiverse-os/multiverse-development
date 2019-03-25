## Multiverse OS: Potentially The Most Secure Hypervior 

A daemon, background process, that unifies Multivesre OS components
to provide a unified, and consistent user experience. Specifically 
the _USER CONTROLLER VM_ is the primary (essentially only) way for
the user to interface with the Multiverse OS oerpating system. 

Unlike most of the technology utilizing the bleeding edge features of virtual machines (cloud computing), Multiverse OS is desgined to be
used by normal (end) users. 

From the beginning it was designed to be a secure hypervisor capable
of segregating multiple online identiets within the same system. It is designed to be able to torrent immediately, or play games on steam, in between programming in Go or Rust in a highly secure environment, capable
of producting reproducible builds. 

No other security oreitend operating system currently offers the level of security provided by Multiverse OS while being a computer that can meet the demands of everyone between gamers and hackers, and everything in bewteen. For example, everything is ran unpriviledged, it was designed that way from Multiverse OS v0.1.0. Complex randomly generated passwords are used and stored behind a locally hosted encrypted password store. 

Multiverse OS is desgined using the security in depth model, every attack vector is assumed to be comprimised including the hardware. That is why the user NEVER interacts with the HOST, ALL activity happens ENTIRELY within nested immutable and/or ephemeral (in regards to disk images) virtual machines. This model has not yet been implemented and offers unparelled security, and usability. Even the newest FPS games can be played within these nested VMs. 

And everything is open source: welcome to the open source Multiverse.

[Designs of comparable security oreinted hypervisors]
Unlike hypervisors, even dedicated to security like Qubes, the user
interacts with the HOST in some way. In Qubes, it is where VMs are launched and managed, it is isolatd by disconnecting the intenret. (which is not enough when one considers the ease of building a purely userspace network stack that could bypass the restrictions),




