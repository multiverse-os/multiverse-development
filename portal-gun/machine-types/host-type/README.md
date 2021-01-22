# Host Machine
The **Host Machine** is the bare-metal machine that runs all the virtual
machines in the Multiverse OS cluster. During normal Multiverse OS operation
the host machine is both completely hidden and completely inaccessible from the
user in order to protect the system and limit the attack surface. Instead of
using the Host Machine like every other operating system, all activity occurs
within the **Controller VM** an ephemeral virtual machine that resets all
changes every reboot. 


#### Feature Brainstorming
Below are a list of potential features that can be implemented to improve and
extend the functionality of the Host Machine.

  * File integrity checking and freezing system that allows for ephemeral
functionality in Multiverse OS VMs. Every reboot files are checked for
integrity using a merkel tree and any changed files are revered to the version
in the last saved merkel tree. The tree can only be saved with the signature of
registered Multiverse OS developers. All Multiverse OS host machines should be
identical and their untampered state can be verified. This design is intended
to allow future resource sharing functionality.
