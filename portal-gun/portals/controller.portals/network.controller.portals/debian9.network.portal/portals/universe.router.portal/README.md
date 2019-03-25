# Universe Router
The **Universe Router VM** is the bridge between a Multiverse OS cluster and the LAN it is connected to. It is the primary line of defense from local attacks and is the VM that has all network devices passed to it using PCI passthrough and functions to air-gap the **Host Machine**.

#### Feature Brainstorming
Below are a list of potential features that can be implemented to improve and extend the functionality of the Controller VM.

  * Scanning detection to track port scans, and other incomming connection attempts.

  * Honey pot functionality to catch login attempt and determine what attacks are being attempted.

  * Stealth functionality to avoid all attempts at scanning or incomming connections from unauthorized sources.

