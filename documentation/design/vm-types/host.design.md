##
##  Multiverse OS: Host Design Document
##======================================
## Introduction
Multiverse OS alpha release will require an installer and preferablly a live 
image that can be put on an USB device or disk for trying out Multiverse OS or 
using your cluster on a computer temporarily. 

### Installer


### Live Image


_______________________________________________________________________________
## Two Controller Host Design
The newest iteration of the Host design is two controllers, one containing
all the routers and one containning all the VMs that the user interfaces with
to use their cluster.

The process of implementing these two controllers will start by implementing
a `qcow2` disk for each one that will eventually be used by `portalgun` to
automate configuration of each controller. But in the meantime we will be 
manually building these and putting all the notes and changed configuration 
files in the corresponding `qcow2` imge.


###
