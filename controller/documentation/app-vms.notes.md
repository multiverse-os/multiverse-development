## Application VMs




#### Nesting
Ideally use 1 core per Application VM. 

#### Communication
Many solutions opt to use `QMP` protocol and perhaps we should implement that, but ideally, and as soon as possible; we want to implement a virtual keyboard device that we can use to send commands through without the need for internet, sockets, or anything other than PS2 or more likely USB. 

The code for virtual keyboard has been partially written. It just needs to be implemented in a low level way into Multiverse OS `portal-gun` VM manager. 

One important aspect of this tool will be confirming successful key sending; so that for example we are not relying on a timer that if off by some number of seconds fails an entire install. We want a functionality that can verify if a step is successful, and if not it can be re-attempted.

In addition to virtual keyboard for low level input with as few dependencies as possible, even though all navigation should be doable through a, it may be useful to provide a virtual mouse or tablet so that clicks, right clicks and so on can be done at a given x,y. If anything this would be useful for later simple automation functionality in Multiverse OS analogous to automater in OSX.


#### Monitoring
Currently the simplest way to preform monitoring, determining if boot is successful and so on; we will be using a serial connection and log output to a file using the following libvirt XML configuration:

````
<console type='file'>
  <source path='/var/multiverse/portals/controller/user/status.out'/>
  <target type='serial' port='0'/>
</console>
````


