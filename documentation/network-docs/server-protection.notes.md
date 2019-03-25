# Server Portection and Instrusion Detection

## Instrusion Detection Methods
Modern motherboards often include instrusion detection systems, but they are known about so inadequate to defend important servers. A new collection of techniques can be applied in random combination to provide unparalleled physical security, ensuring that even servers stored in a data center are safe from instrusion and possibly provide information about the intruder.

*Proximity*
Proximity based sensors will rely on components being in vicinity of one another. Removal of a device from other devices (removal from the rack) will trigger an alarm, computer shutdown, and possibly cause components to be damaged.

* Wifi based
* RFID Based
* Magnet/Magnometer Based

*Distance*
Distance sensors can be used to determine if the case as been removed and if an intruder has illegally entered the case to tamper with or steal components.

* IR
* Laser
* Ultrasonic

*Survellience*
Survellience systems to gather more details about an intruder so that when an alarm is triggered, data points around the time of triggering the alarm can be streamed offsite.

* Sound (Microphone)
* Visual (Camera)

*Other*

* Light sensors to determine if the case is open

## Honeypots and Alarms
The motherboard of the machine should not be installed against the back end of the case, instead a custom 3D printed backplate should be installed and available ports should be routed to a intermediary proxy computer, using either a Rapsbery Pi or similar SOC (system-on-chip) device that can act as a router and proxy between the datacenter network and your server. This will allow greater physical security and better network design.

* Use 2x SOC to wire all motherboard into 1 and all outside ports into another and a line between the two for connections.

* 3D printed Custom backplate exposing SOC devices

* Wire front USB drives to SOC

* Allow SOC to programatically route hot swap HDs to the motherboard of the main system. 

* Wire hotswap HDs alarm to SOC

* Provide the SOC with a backup battery, possibly two, one for trigger detection low power mode that turns on SOC when required if no power is being provided by the datacenter.

## Self-desctruction
If required by clients, self-destruct mechanisms may be installed to prevent components from being stolen or confiscated by third parties. 

* Wire PSU voltage to SOC

* Include an extra battery that can be used in combination with a set of capacitors to fry specific components

## Locking Mechanisms

## Environmental Status

* Use temperature, humidity, pressure sensors to detect temperature or other abnormalities

* 

