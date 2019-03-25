## VM CPU-Passthrough And CPU Features
The ideal CPU configuration for our Controller VMs is using the mode `host-passthrough` and modified by going through each available CPU option; obtained by running the command `lscpu` on the host machine, taking the list of features and explicitly defining them in the libvirt XML either disabling or enabling each feature option. 

A reference for CPU features with security issues should be included with Multiverse OS so that all these features can be disabled by default.

The end result is described in other documentation; but in essence it provides a custom version of your CPU that only has access to secure and approved features.


