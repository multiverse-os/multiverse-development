# Machine Types 
We are following the original four main types: 
    * Host Machine (Physical Bare-metal machine) - a completely locked machine,
      that the user never directly interacts with, the entire machine is hidden
      from the user, and the controller is automatically launched to give the
      user a "host" that is ephemeral, easily interchangable for different
      tasks, like gaming vs adobe photoshop. And ideally makes attacks on the
      host stall by having everything disabled, and every feature hooked to an
      alarm system to notify intrusion since any action in this machine is
      indication of an instruction. Upgrades happen via special rescue-mode type
      boot environment but still indirectly, the user never itneracts with this
      machine to ensure that this has the smallest possible attack surface. And
      the controller is also able to select only the features desired (like CPU
      features) to passthrough, allowing the controller to exclude any features
      or functionality of the host that increases the attack surface or
      introduces a security vulnerability.
    * Controller Machine  - a incredibly limited machine, that indirectly
      accesses the internet, and launches all its applications even file browser
      and terminal within application machines. It does however focus all logs,
      and all cluster data into a single location so that it can represent the
      entire cluster as a single machine. 
    * Application Machine - an encapsulated application machine with all
      necessary libraries and an ephemeral system to make tampering, viruses,
      and other attacks difficult by preventing any persistence and
      compartmentalizing logic. 
    * Service Machine - for long running tasks that do not require any UI, only
      rules for when they should be ran, and hooks to customize functionality.
