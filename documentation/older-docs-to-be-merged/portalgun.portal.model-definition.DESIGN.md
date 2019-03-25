# Multiverse OS: Portal Gun's 'Portal' Model 
#============================================
Portal Gun is Multierse OS hypervisor or virtual machine (VM) manager. Each VM or `portal` is a complex virtual machine encapsualted into a binary and very easy to move around carrying essentially all of its dependencies. Including the USER CONTROLLER VM, a portal that is itself a hypervisor that runs nested inside the HOST, and the ROUTER CONTROLLER VM which encapsualtes at least three routers inside to allow portals to be transparently routed over at least three possible endpoints: [local area network (LAN), virtual private network (VPN), onion network (TOR)]

Below is the model for a portal, by defining the database model, we can provide a shared library which is used across all Multiverse OS software to access and interact with databases holding portal data. 

One important aspect of Multiverse portal design is that they are ephemeral
meaning the operating system components are revereted back to match a merkle
tree checksum defined after the last install or update. The OS portion of
the OS is immutable or read-only. Additionally the VMs need to be modified
in such a way that is it difficult to impossible to verify that the VM is
a virtualized environment.

## Notes

  * Each VM should have at least two (2) serial connections: 
	(1) for virtual console connection for maintenance and development 
	(2) 1-way connection to file to log status named `status.out`

  * Each VM needs to supply a `unix socket` that we can use the UDP like 
    protocl over: 
      [**SOCK_SOCKDGRAM**][UDP] UDP lacks the extra logic on the
      protocol for ensuring reliable transmission and ordering of
      packets. Since we are moving through the same machine on
      virtual connections, this is the best option.

      [or]

      [**UDP network console**][UDP] Since UDP is just sending packets and
      not checking anything unnessarily, which without any packet loss
      is likely the most ideal way to send data. **And because of this**
      **lack of unnecessary checking added for issues with physical**
      **networking issues, UDP is a wise choice.**

  * Need other connections for VM-to-VM connections

  * Increase the number of serialVirtIO to match the number of vims and
    serial connections so we have more bandwidth

  * `copy_on_read` should be used on the base image, as well as `readonly`

  * 


## Scripting the optimization of XML
This will be critical the development of portalgun and will make it possible to get rid of virt-manager immediately, and scale back the use of libvirt until we eventually abandon its use. 

  * We need to immediately script the optimizaiton of XML
    * CPU pinning
    * Correct clock setup
    * Adding disks for specific VMs
    * optimizing disk settings, memory settings, etc.

    

## Portal Model
The format is: '[attribute_name][data_type] description'




[*][portal_name][string] a unique name that should be in the format of a
 | domain, so that it can be resolved by internal DNS servers within the
 | Multiverse OS custer.
 |
 |-[i.e.] [user.controller.portal]
 |        [subdomain.domain.tld]
 o


[*][id][string] a uuid based on scramble id that is compatible with onion
 | addresses.
 |
 |-[i.e.]
 |
 o

[*][id][string] a uuid based on scramble id that is compatible with onion
 | addresses.
 |
 |-[i.e.]
 |
 o




## Template for adding entries

[*][attribute_name][data_type] description
 |
 |-[i.e.]
 |
 |-[*]
 |
 |
 [1]>
 |
 [2]>
  |
  |-[A]
  |
  |-[B]
  |
  |-[C]
  |
  o













