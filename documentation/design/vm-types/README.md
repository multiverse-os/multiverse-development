##
## Multiverse OS: Virtual Machine (VM) Type Design Documents
## =========================================================

These design documents are the newest iteration of the classification (type or kind) for virtual machines within the context of Multiverse OS. These VM classes provide specific functionality to function as building blocks for building the default Multiverse OS cluster. 

Each class provides an `agent`, a default configuration, to provide the specific functionality specified in the design document. 

## Updates

  * Previous VM types incldued a unique **ROUTER VM** type, which will now be folded into **SERVICE VM**.

  * Previously, the **CONTROLLER VM** was only to encapsulate **APP and SERVICE VMs** while providing the interface for the user to interact with the Multiverse OS cluster. Now there are **TWO** (2) types of **CONTROLLER VMs**. Still they are containers to encapsulate and nest **APP and SERVICE VMs**, one for the networking infrastructure (to enable easily transportable complex network structure that can be dropped into Multiverse OS cluster), and the original one that encapsulated the **APP VMs and long-running SERVICE VMs**. 

  
