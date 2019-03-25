# Controller Networking
There are now two operating modes for the controller which require
two different networking configurations. One requires fine grain
control over the routing table to support multiple default routes
depending on originating network, and the other hands off complexity
to the routers, in the same way it is done on the host. 

This part of the design is not finalized and is subject to change. 
