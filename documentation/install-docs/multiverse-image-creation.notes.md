# Multiverse Image Creation
One of the main design principles around Multiverse OS is shifting the provisioning process/step into the image creation step, making fully provisioned live CDs for immutable ephemeral VMs. 

This was a major missing feature in both Docker and Vagrant. This step in these programs causes insecure system where one relies far too much on unknown people. In constrast, Multiverse builds every disk locally, including your install disk. This is the foundation of a trustless system

Packer has overlap with this, so packer should be looked at for a code examples, and look at the issues for pain points and lessions to prevent reinventing the wheel and reninventing the same issues to learn from their mistakes.

---

__OPTIONS__
Unviersal spinner for Linux Medium
https://github.com/solus-project/USpin
