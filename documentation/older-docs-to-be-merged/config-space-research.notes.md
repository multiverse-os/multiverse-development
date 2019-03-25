# Config Space Issue

Under Debian Buster an error has been introduced that causes failure to do pci-passthrough for NIC and GPU devices due to failuer to open the config space file within the PCI device /sys/bus folder. However this issue is not present in the previous Debian version Stretch. Below is research relating to this topic and hopefully the eventual patch that fixes the problem.

### Software Differences
The seemingly best way to approach this issue is to track the differences in software versions then step through the git repositories for the projects and try to determine where the bug was introduced so a patch can be written to fix the issue. Below is a table of relevant software versions used by each version of Debian.


|    Software             |    Debian Stretch      |    Debian Buster       |
|-------------------------|------------------------|------------------------|
| libvirt0                | 3.0.0-4+deb9           | 4.1.0-2                |
| libvirt-clients         | 3.0.0-4+deb9           | 4.1.0-2                |
| libvirt-daemon          | 3.0.0-4+deb9           | 4.1.0-2                |
| libvirt-daemon-system   | 3.0.0-4+deb9           | 4.1.0-2                |
| libvirt-glib            | 1.0-                   | 1.0.0-1                |
| qemu                    | 1:2.8                  | 1:2.11                 |
| qemu-kvm                | 1:2.8                  | 1:2.11                 |
| qemu-system-x86         | 1:2.8                  | 1:2.11                 |

