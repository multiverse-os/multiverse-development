
###############################################################################
# Add virtual PCI subsystem & slots
###############################################################################
# This concept needs to be translated to portalgun, this is the
# device on the host machine:
#
# <device>
#   <name>pci_0000_00_00_0</name>
#
#   *(This is one the virtual functions specified by path)*
#   <path>/sys/devices/pci0000:00/0000:00:01.0/0000:00:00.0</path>
#
#   <parent>computer</parent>
#   <driver>
#     <name>e1000e</name>
#   </driver>
#   <capability type='pci'>
#     <domain>0</domain>
#     <bus>0</bus>
#     <slot>25</slot>
#     <function>0</function>
#     <product id='0x1502'>82579LM Gigabit Network Connection</product>
#     <vendor id='0x8086'>Intel Corporation</vendor>
#     <iommuGroup number='7'>
#       <address domain='0x0000' bus='0x00' slot='0x00' function='0x0' />
#     </iommuGroup>
#   </capability>
# </device>
#
#
# Translates to a PCI passthrough:
#
# <hostdev>
#   <source> 
#     <address domain='0' bus='0' slot='25' function='0' />
#   </source>
#   <rom bar='off' />  (lspci -v command to check the device for expansion rom)
#   (the rom can sometimes be bypassed for speed, or other issues related can 
#   prevent booting)
# </hostdev>
#
###############################################################################
# **Virtual function without the device**
# SR-IOV PCI Device, Virtual Functionality 
# Virtual functions provide data protection between virtual machines and 
# the physical server as data managed is controlled within the hardware. 
#
#
# <interface type="hostdev">
