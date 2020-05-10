#!/bin/sh
#
###############################################################################
# 'v.sh' is an overlay command to put virsh more inline with the Multiverse OS
#        command to provide a stop gap solution until a more mature version of
#        portal gun is available. 
###############################################################################
# Controller Settings
###############################################################################
#export DEFAULT_CONTROLLER="ubuntu.controller"
DEFAULT_CONTROLLER="ubuntu.controller"
###############################################################################
# Overlay Settings
###############################################################################
_command_name="vsh"
_cmd="virsh"
_global_flags=""
#_global_flags="--connect=qemu:///session"
_command_flags="--all"
_command_version="0.1.0"

###############################################################################

help(){
	echo "'$_command_name' cli tool requires command, for example:"
	echo "    '$_command_name machine controller start'"
	echo "-OR-:" 
	echo "    '$_command_name machine network list'"
	echo "-OR-:" 
	echo "    '$_machine_name machine network (start|stop)'"
}

###############################################################################

if [ "machine" = "$1" -o "m" = "$1" -o "machines" = "$1" ]; then 
	if [ "controller" = "$2" -o "c" = "$2" ]; then
		if [ "start" = "$3" -o "s" = "$3" ]; then
			if [ $default_controller = "$4" -o "default" = "$4" -o "d" = "$4" ];  then
				sh "$_cmd $_global_flags start $DEFAULT_CONTROLLER"
			fi
		elif [ "stop" = "$3" -o "x" = "$3" ]; then
			if [ $DEFAULT_CONTROLLER = "$4" ]; then
				sh "$_cmd $_global_flags stop $DEFAULT_CONTROLLER"
			else 
				if [ -z "$4" ]; then
					sh "$_cmd $_global_flags stop $4"
				else 
					echo "A controller name was not specified, select one from the VM list below:"
					sh "sudo $_cmd $_global_flags list $_command_flags"
				fi
			fi
		elif [ "edit" = "$3" -o "e" = "$3" ]; then
			if [ $DEFAULT_CONTROLLER = "$4" ]; then
				sh "$_cmd $_global_flags edit $DEFAULT_CONTROLLER"
			else 
				if [ -z "$4" ]; then
					sh "$_cmd $_global_flags edit $4"
				else 
					echo "[$_command_name] Controller name not specified, select one from the VM list below:"
					sh "sudo $_cmd $_global_flags list $_command_flags"
				fi
			fi
		elif [ "list" = "$3" -o "l" = "$3" ]; then
			# TODO: Controller list not implemented
			echo "[$_command_name] Not implemented yet."
			sh "$_cmd $_global_flags list $_command_flags"
		else 
			echo "[$_command_name] Available network actions are:"
			echo "      [ list, stop, start, edit ]"
		fi
	elif [ "network" = "$2" -o "n" = "$2" ]; then
		if [ "start" = "$3" -o "s" = "$3" ]; then
			# Network 1
			echo "[$_command_name] Starting network 0, bridge 0, and bridge 1..."
			sh "$_cmd $_global_flags start network0.bridge0.router"
			sh "$_cmd $_global_flags start network0.bridge1.router"
			echo "[$_command_name] Starting network 1, bridge 0, and bridge 1..."
			sh "$_cmd $_global_flags start network1.bridge0.router"
			sh "$_cmd $_global_flags start network1.bridge1.router"
			echo "[$_command_name] Starting network 2, bridge 0, and bridge 1..."
			sh "$_cmd $_global_flags start network2.bridge0.router"
			sh "$_cmd $_global_flags start network2.bridge1.router"
			echo "[$_command_name] Starting network 3, bridge 0, and bridge 1..."
			sh "$_cmd $_global_flags start network3.bridge0.router"
			sh "$_cmd $_global_flags start network3.bridge1.router"
		elif [ "stop" = "$3" -o "x" = "$3" ]; then
			# Network 1
			echo "[$_command_name] Starting network 0, bridge 0, and bridge 1..."
			sh $"_cmd $_global_flags stop network0.bridge0.router"
			sh "$_cmd $_global_flags stop network0.bridge1.router"
			echo "[$_command_name] Starting network 1, bridge 0, and bridge 1..."
			sh "$_cmd $_global_flags stop network1.bridge0.router"
			sh "$_cmd $_global_flags stop network1.bridge1.router"
			echo "[$_command_name] Starting network 2, bridge 0, and bridge 1..."
			sh "$_cmd $_global_flags stop network2.bridge0.router"
			sh "$_cmd $_global_flags stop network2.bridge1.router"
			echo "[$_command_name] Starting network 3, bridge 0, and bridge 1..."
			sh "$_cmd $_global_flags stop network3.bridge0.router"
			sh "$_cmd $_global_flags stop network3.bridge1.router"
		elif [ "edit" = "$3" -o "e" = "$3" ]; then
			if [ -z "$3" ]; then
				echo "[$_command_name] Network router was not specified, select one from the VM list below:"
				sh "sudo $_cmd $_global_flags list $_command_flags"
			else
				sh "$_cmd $_global_flags edit $4"
			fi
		elif [ "list" = "$3" -o "l" = "$3" ]; then
			# TODO: Network list not implemented yet
			echo "[$_command_name] Not implemented yet."
			sh $_cmd $_global_flags list $_command_flags
		elif [ "bridges" = "$3" -o "b" = "$3" ]; then 
			sh "sudo $_cmd $_global_flags net-list $_command_flags"
		else 
			help
			echo ""
			echo "[$_command_name] Available network actions are:"
			echo "      [ list, stop, start, edit, bridges ]"
		fi
	elif [ "service" = "$2" -o "s" = "$2" ]; then
		if [ "start" = "$3" -o "s" = "$3" ]; then
			echo "[$_command_name] Service is not implemented yet."
		elif [ "stop" = "$3" -o "x" = "$3" ]; then
			echo "[$_command_name] Service is not implemented yet."
		elif [ "list" = "$3" -o "l" = "$3" ]; then
			echo "[$_command_name] Not implemented yet."
			sh "$_cmd $_global_flags list $_command_flags"
		else 
			help
			echo ""
			echo "[$_command_name] Available service actions are:"
			echo "      [ list, stop, start, edit ]"
		fi
	elif [ "app" = "$2" -o "a" = "$2" -o "application" = "$2" ]; then
		if [ "start" = "$3" -o "s" = "$3" ]; then
			echo "[$_command_name] service is not implemented yet."
		elif [ "stop" = "$3" -o "x" = "$3" ]; then
			echo "[$_command_name] service is not implemented yet."
		elif [ "list" = "$3" -o "l" = "$3" ]; then
			echo "[$_command_name] Not implemented yet."
			sh "$_cmd $_global_flags list $_command_flags"
		else 
			help
			echo ""
			echo "[$_command_name] Available app actions are:"
			echo "      [ list, stop, start, edit ]"
		fi
	else # "list"+help
		help
		echo ""
		sh "$_cmd $_global_flags list $_command_flags"
		echo ""
		echo "[$_command_name] Available machine types are:"
		echo "      [ controller, network, service, app ]"
	fi
elif [ "version" = "$1"  -o "v" = "$1" -o "-v" = "$1" -o "--version" = "$1" ]; then
	echo "[$_command_name] Command overlay version is: v$_command_version"
else # "help" "--help" "-h" 
	help
	echo ""
	sh "$_cmd $_global_flags list $_command_flags"
	echo ""
	echo "[$_command_name] Available network actions are:"
	echo "      [ list, stop, start, edit, bridges ]"

fi 


###############################################################################
# Lower layer of command overlay
###############################################################################
#	virsh [options]... [<command_string>]
#	virsh [options]... <command> [args...]
#
#	options:
#	-c | --connect=URI      hypervisor connection URI
#	-d | --debug=NUM        debug level [0-4]
#	-e | --escape <char>    set escape sequence for console
#	-h | --help             this help
#	-k | --keepalive-interval=NUM
#	keepalive interval in seconds, 0 for disable
#	-K | --keepalive-count=NUM
#	number of possible missed keepalive messages
#	-l | --log=FILE         output logging to file
#	-q | --quiet            quiet mode
#	-r | --readonly         connect readonly
#	-t | --timing           print timing information
#	-v                      short version
#	-V                      long version
#	--version[=TYPE]   version, TYPE is short or long (default short)
#	commands (non interactive mode):
#
#	Domain Management (help keyword 'domain')
#	attach-device                  attach device from an XML file
#	attach-disk                    attach disk device
#	attach-interface               attach network interface
#	autostart                      autostart a domain
#	blkdeviotune                   Set or query a block device I/O tuning parameters.
#	blkiotune                      Get or set blkio parameters
#	blockcommit                    Start a block commit operation.
#	blockcopy                      Start a block copy operation.
#	blockjob                       Manage active block operations
#	blockpull                      Populate a disk from its backing image.
#	blockresize                    Resize block device of domain.
#	change-media                   Change media of CD or floppy drive
#	console                        connect to the guest console
#	cpu-stats                      show domain cpu statistics
#	create                         create a domain from an XML file
#	define                         define (but don't start) a domain from an XML file
#	desc                           show or set domain's description or title
#	destroy                        destroy (stop) a domain
#	detach-device                  detach device from an XML file
#	detach-device-alias            detach device from an alias
#	detach-disk                    detach disk device
#	detach-interface               detach network interface
#	domdisplay                     domain display connection URI
#	domfsfreeze                    Freeze domain's mounted filesystems.
#	domfsthaw                      Thaw domain's mounted filesystems.
#	domfsinfo                      Get information of domain's mounted filesystems.
#	domfstrim                      Invoke fstrim on domain's mounted filesystems.
#	domhostname                    print the domain's hostname
#	domid                          convert a domain name or UUID to domain id
#	domif-setlink                  set link state of a virtual interface
#	domiftune                      get/set parameters of a virtual interface
#	domjobabort                    abort active domain job
#	domjobinfo                     domain job information
#	domname                        convert a domain id or UUID to domain name
#	domrename                      rename a domain
#	dompmsuspend                   suspend a domain gracefully using power management functions
#	dompmwakeup                    wakeup a domain from pmsuspended state
#	domuuid                        convert a domain name or id to domain UUID
#	domxml-from-native             Convert native config to domain XML
#	domxml-to-native               Convert domain XML to native config
#	dump                           dump the core of a domain to a file for analysis
#	dumpxml                        domain information in XML
#	edit                           edit XML configuration for a domain
#	event                          Domain Events
#	inject-nmi                     Inject NMI to the guest
#	iothreadinfo                   view domain IOThreads
#	iothreadpin                    control domain IOThread affinity
#	iothreadadd                    add an IOThread to the guest domain
#	iothreadset                    modifies an existing IOThread of the guest domain
#	iothreaddel                    delete an IOThread from the guest domain
#	send-key                       Send keycodes to the guest
#	send-process-signal            Send signals to processes
#	lxc-enter-namespace            LXC Guest Enter Namespace
#	managedsave                    managed save of a domain state
#	managedsave-remove             Remove managed save of a domain
#	managedsave-edit               edit XML for a domain's managed save state file
#	managedsave-dumpxml            Domain information of managed save state file in XML
#	managedsave-define             redefine the XML for a domain's managed save state file
#	memtune                        Get or set memory parameters
#	perf                           Get or set perf event
#	metadata                       show or set domain's custom XML metadata
#	migrate                        migrate domain to another host
#	migrate-setmaxdowntime         set maximum tolerable downtime
#	migrate-getmaxdowntime         get maximum tolerable downtime
#	migrate-compcache              get/set compression cache size
#	migrate-setspeed               Set the maximum migration bandwidth
#	migrate-getspeed               Get the maximum migration bandwidth
#	migrate-postcopy               Switch running migration from pre-copy to post-copy
#	numatune                       Get or set numa parameters
#	qemu-attach                    QEMU Attach
#	qemu-monitor-command           QEMU Monitor Command
#	qemu-monitor-event             QEMU Monitor Events
#	qemu-agent-command             QEMU Guest Agent Command
#	guest-agent-timeout            Set the guest agent timeout
#	reboot                         reboot a domain
#	reset                          reset a domain
#	restore                        restore a domain from a saved state in a file
#	resume                         resume a domain
#	save                           save a domain state to a file
#	save-image-define              redefine the XML for a domain's saved state file
#	save-image-dumpxml             saved state domain information in XML
#	save-image-edit                edit XML for a domain's saved state file
#	schedinfo                      show/set scheduler parameters
#	screenshot                     take a screenshot of a current domain console and store it into a file
#	set-lifecycle-action           change lifecycle actions
#	set-user-password              set the user password inside the domain
#	setmaxmem                      change maximum memory limit
#	setmem                         change memory allocation
#	setvcpus                       change number of virtual CPUs
#	shutdown                       gracefully shutdown a domain
#	start                          start a (previously defined) inactive domain
#	suspend                        suspend a domain
#	ttyconsole                     tty console
#	undefine                       undefine a domain
#	update-device                  update device from an XML file
#	vcpucount                      domain vcpu counts
#	vcpuinfo                       detailed domain vcpu information
#	vcpupin                        control or query domain vcpu affinity
#	emulatorpin                    control or query domain emulator affinity
#	vncdisplay                     vnc display
#	guestvcpus                     query or modify state of vcpu in the guest (via agent)
#	setvcpu                        attach/detach vcpu or groups of threads
#	domblkthreshold                set the threshold for block-threshold event for a given block device or it's backing chain element
#	guestinfo                      query information about the guest (via agent)
#
#	Domain Monitoring (help keyword 'monitor')
#	domblkerror                    Show errors on block devices
#	domblkinfo                     domain block device size information
#	domblklist                     list all domain blocks
#	domblkstat                     get device block stats for a domain
#	domcontrol                     domain control interface state
#	domif-getlink                  get link state of a virtual interface
#	domifaddr                      Get network interfaces' addresses for a running domain
#	domiflist                      list all domain virtual interfaces
#	domifstat                      get network interface stats for a domain
#	dominfo                        domain information
#	dommemstat                     get memory statistics for a domain
#	domstate                       domain state
#	domstats                       get statistics about one or multiple domains
#	domtime                        domain time
#	list                           list domains
#
#	Host and Hypervisor (help keyword 'host')
#	allocpages                     Manipulate pages pool size
#	capabilities                   capabilities
#	cpu-baseline                   compute baseline CPU
#	cpu-compare                    compare host CPU with a CPU described by an XML file
#	cpu-models                     CPU models
#	domcapabilities                domain capabilities
#	freecell                       NUMA free memory
#	freepages                      NUMA free pages
#	hostname                       print the hypervisor hostname
#	hypervisor-cpu-baseline        compute baseline CPU usable by a specific hypervisor
#	hypervisor-cpu-compare         compare a CPU with the CPU created by a hypervisor on the host
#	maxvcpus                       connection vcpu maximum
#	node-memory-tune               Get or set node memory parameters
#	nodecpumap                     node cpu map
#	nodecpustats                   Prints cpu stats of the node.
#	nodeinfo                       node information
#	nodememstats                   Prints memory stats of the node.
#	nodesuspend                    suspend the host node for a given time duration
#	sysinfo                        print the hypervisor sysinfo
#	uri                            print the hypervisor canonical URI
#	version                        show version
#
#	Checkpoint (help keyword 'checkpoint')
#	checkpoint-create              Create a checkpoint from XML
#	checkpoint-create-as           Create a checkpoint from a set of args
#	checkpoint-delete              Delete a domain checkpoint
#	checkpoint-dumpxml             Dump XML for a domain checkpoint
#	checkpoint-edit                edit XML for a checkpoint
#	checkpoint-info                checkpoint information
#	checkpoint-list                List checkpoints for a domain
#	checkpoint-parent              Get the name of the parent of a checkpoint
#
#	Interface (help keyword 'interface')
#	iface-begin                    create a snapshot of current interfaces settings, which can be later committed (iface-commit) or restored (iface-rollback)
#	iface-bridge                   create a bridge device and attach an existing network device to it
#	iface-commit                   commit changes made since iface-begin and free restore point
#	iface-define                   define an inactive persistent physical host interface or modify an existing persistent one from an XML file
#	iface-destroy                  destroy a physical host interface (disable it / "if-down")
#	iface-dumpxml                  interface information in XML
#	iface-edit                     edit XML configuration for a physical host interface
#	iface-list                     list physical host interfaces
#	iface-mac                      convert an interface name to interface MAC address
#	iface-name                     convert an interface MAC address to interface name
#	iface-rollback                 rollback to previous saved configuration created via iface-begin
#	iface-start                    start a physical host interface (enable it / "if-up")
#	iface-unbridge                 undefine a bridge device after detaching its slave device
#	iface-undefine                 undefine a physical host interface (remove it from configuration)
#
#	Network Filter (help keyword 'filter')
#	nwfilter-define                define or update a network filter from an XML file
#	nwfilter-dumpxml               network filter information in XML
#	nwfilter-edit                  edit XML configuration for a network filter
#	nwfilter-list                  list network filters
#	nwfilter-undefine              undefine a network filter
#	nwfilter-binding-create        create a network filter binding from an XML file
#	nwfilter-binding-delete        delete a network filter binding
#	nwfilter-binding-dumpxml       network filter information in XML
#	nwfilter-binding-list          list network filter bindings
#
#	Networking (help keyword 'network')
#	net-autostart                  autostart a network
#	net-create                     create a network from an XML file
#	net-define                     define an inactive persistent virtual network or modify an existing persistent one from an XML file
#	net-destroy                    destroy (stop) a network
#	net-dhcp-leases                print lease info for a given network
#	net-dumpxml                    network information in XML
#	net-edit                       edit XML configuration for a network
#	net-event                      Network Events
#	net-info                       network information
#	net-list                       list networks
#	net-name                       convert a network UUID to network name
#	net-start                      start a (previously defined) inactive network
#	net-undefine                   undefine a persistent network
#	net-update                     update parts of an existing network's configuration
#	net-uuid                       convert a network name to network UUID
#	net-port-list                  list network ports
#	net-port-create                create a network port from an XML file
#	net-port-dumpxml               network port information in XML
#	net-port-delete                delete the specified network port
#
#	Node Device (help keyword 'nodedev')
#	nodedev-create                 create a device defined by an XML file on the node
#	nodedev-destroy                destroy (stop) a device on the node
#	nodedev-detach                 detach node device from its device driver
#	nodedev-dumpxml                node device details in XML
#	nodedev-list                   enumerate devices on this host
#	nodedev-reattach               reattach node device to its device driver
#	nodedev-reset                  reset node device
#	nodedev-event                  Node Device Events
#
#	Secret (help keyword 'secret')
#	secret-define                  define or modify a secret from an XML file
#	secret-dumpxml                 secret attributes in XML
#	secret-event                   Secret Events
#	secret-get-value               Output a secret value
#	secret-list                    list secrets
#	secret-set-value               set a secret value
#	secret-undefine                undefine a secret
#
#	Snapshot (help keyword 'snapshot')
#	snapshot-create                Create a snapshot from XML
#	snapshot-create-as             Create a snapshot from a set of args
#	snapshot-current               Get or set the current snapshot
#	snapshot-delete                Delete a domain snapshot
#	snapshot-dumpxml               Dump XML for a domain snapshot
#	snapshot-edit                  edit XML for a snapshot
#	snapshot-info                  snapshot information
#	snapshot-list                  List snapshots for a domain
#	snapshot-parent                Get the name of the parent of a snapshot
#	snapshot-revert                Revert a domain to a snapshot
#
#	Backup (help keyword 'backup')
#	backup-begin                   Start a disk backup of a live domain
#	backup-dumpxml                 Dump XML for an ongoing domain block backup job
#
#	Storage Pool (help keyword 'pool')
#	find-storage-pool-sources-as   find potential storage pool sources
#	find-storage-pool-sources      discover potential storage pool sources
#	pool-autostart                 autostart a pool
#	pool-build                     build a pool
#	pool-create-as                 create a pool from a set of args
#	pool-create                    create a pool from an XML file
#	pool-define-as                 define a pool from a set of args
#	pool-define                    define an inactive persistent storage pool or modify an existing persistent one from an XML file
#	pool-delete                    delete a pool
#	pool-destroy                   destroy (stop) a pool
#	pool-dumpxml                   pool information in XML
#	pool-edit                      edit XML configuration for a storage pool
#	pool-info                      storage pool information
#	pool-list                      list pools
#	pool-name                      convert a pool UUID to pool name
#	pool-refresh                   refresh a pool
#	pool-start                     start a (previously defined) inactive pool
#	pool-undefine                  undefine an inactive pool
#	pool-uuid                      convert a pool name to pool UUID
#	pool-event                     Storage Pool Events
#	pool-capabilities              storage pool capabilities
#
#	Storage Volume (help keyword 'volume')
#	vol-clone                      clone a volume.
#	vol-create-as                  create a volume from a set of args
#	vol-create                     create a vol from an XML file
#	vol-create-from                create a vol, using another volume as input
#	vol-delete                     delete a vol
#	vol-download                   download volume contents to a file
#	vol-dumpxml                    vol information in XML
#	vol-info                       storage vol information
#	vol-key                        returns the volume key for a given volume name or path
#	vol-list                       list vols
#	vol-name                       returns the volume name for a given volume key or path
#	vol-path                       returns the volume path for a given volume name or key
#	vol-pool                       returns the storage pool for a given volume key or path
#	vol-resize                     resize a vol
#	vol-upload                     upload file contents to a volume
#	vol-wipe                       wipe a vol
#
#	Virsh itself (help keyword 'virsh')
#	cd                             change the current directory
#	echo                           echo arguments
#	exit                           quit this interactive terminal
#	help                           print help
#	pwd                            print the current directory
#	quit                           quit this interactive terminal
#	connect                        (re)connect to hypervisor
#
#
#	(specify help <group> for details about the commands in the group)
#
#	(specify help <command> for details about the command)
#
