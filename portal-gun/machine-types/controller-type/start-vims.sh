#!/bin/bash
###############################################################################

## Multiverse Controller Startup Process
###############################################################################
## Start Controller's default VMs after setting up networking. This will 
## prevent less attack surface on our controller. The terminal that is 
## launched should be our base desktop style VM.

#### TASKS 

## * Build network devices, routing tables, and ensure routes work for each 
##   device (testing with four)

## * Startup default desktop, firefox, and other basic virtual machines that 
##   isolate and compartmentalize various aspects of our machine, which will 
##   be more clear as we drop libvirt support. 

## * Have a basic application check status and shape of the network graph
