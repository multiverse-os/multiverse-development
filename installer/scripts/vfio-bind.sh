#/bin/sh
###############################################################################
## Global Input Variables / Arguments
##=============================================================================
argument=""
argument_format=""
argument_length=0
## Derived Data
device_id=""
device_id_length=0
long_pci_address=""
long_pci_address_length=0
short_pci_address=""
short_pci_address_length=0
config_space_path=""
sysfs_path=""
kernel_module=""
kernel_driver=""
###############################################################################
## Shell Helpers
##=============================================================================
command_name="vfio-bind"
mv_user="user"
mv_group="kvm"
ignore_errors="2>/dev/null"
ignore_output="&>/dev/null"
###############################################################################
## Multiverse OS Script Color Palette
##=============================================================================
header="\e[0;95m"   # PURPLE
accent="\e[37m"     # WHITE
subheader="\e[98m"  # GRAY 
strong="\e[96m"     # CYAN
text="\e[94m"       # BLUE
success="\e[92m"    # GREEN
warning="\e[93m"    # YELLOW
fail="\e[91m"       # RED
reset="\e[0m"       # Terminal Default
###############################################################################
## Global Variables
##=============================================================================
short_address_format="00:00.0"
device_id_format="0000:0000"
long_address_format="0000:00:00.0"
short_address_format_length=7
device_id_format_length=9
long_address_format_length=12
###############################################################################
## Errors
##=============================================================================
error_must_be_root="Must be logged in as root. Run 'su' and try again."
error_device_not_found="No device found; verify your device ID or PCI address and try again."
error_invalid_input="No Device Id or PCI Address of device argument provided."
error_invalid_short_id="No device ID found, invalid short PCI address."
error_invalid_long_id="No device ID found, invalid long PCI address."
error_invalid_device_found="Device found is invalid, incorrect length."
error_invalid_format="Invalid input; must have character length of:\n$reset$accent    $short_address_length$reset$subheader characters$reset$strong  ....$reset$accent  $short_address_format$reset$header [Short PCI Address]\n$reset$accent    $device_id_length$reset$subheader characters$reset$strong  ....$reset$accent  $device_id_format  $reset$header [Device ID]\n$reset$accent    $long_address_length$reset$subheader characters$reset$strong  ....$reset$accent  $long_address_length$reset$header [Full Address] to be valid."

in_brackets(){
	text=$1
	color=$2
	echo "$accent[$reset$color$text$accent]$reset"
}

log_prefix=$(in_brackets "$command_name" $strong)
debug_prefix=$(in_brackets "DEBUG" $success)
error_prefix=$(in_brackets "ERROR" $fail)
fatal_error_prefix=$(in_brackets "FATAL ERROR" $fail)
warning_prefix=$(in_brackets "WARNING" $fail)
info_prefix=$(in_brackets "INFO" $fail)
success_prefix=$(in_brackets "SUCCESS" $success) 
###############################################################################
##  COMMAND-LINE INTERFACE shell functions
##=============================================================================
length(){
	echo ${#1}
}

exit_if_not_root(){
	if [ $(whoami) = $mv_user ]; then
		print_fatal "must_be_root"
	fi
}

color_text(){
	color=$1
	text=$2
	if [ $color -eq "red" -o $color -eq "fail" ]; then
		echo "$fail$text$reset"
	elif [ $color -eq "green" -o $color -eq "success" ]; then
		echo "$success$text$reset"
	elif [ $color -eq "purple" -o $color -eq "header" ]; then
		echo "$header$text$reset"
	elif [ $color -eq "gray" -o $color -eq "subheader" ]; then
		echo "$subheader$text$reset"
	elif [ $color -eq "white" -o $color -eq "accent" ]; then
		echo "$accent$text$reset"
	elif [ $color -eq "blue" -o $color -eq "text" ]; then
		echo "$text$text$reset"
	elif [ $color -eq "yellow" -o $color -eq "warning" ]; then
		echo "$warning$text$reset"
	elif [ $color -eq "cyan" -o $color -eq "strong" ]; then
		echo "$strong$text$reset"
	else
		echo "$reset$text$reset"
	fi
}
###############################################################################
##  COMMAND-LINE INTERFACE shell functions
###############################################################################
print_error(){
	if [ $1 = "must_be_root" ]; then
		echo "$error_prefix $accent$error_must_be_root$reset"
	elif [ $1 = "invalid_input" ]; then
		echo "$error_prefix $accent$error_invalid_input$reset"
	elif [ $1 = "invalid_short_id" ]; then
		echo "$error_prefix $accent$error_invalid_short_id$reset"
	elif [ $1 = "invalid_long_id" ]; then
		echo "$error_prefix $accent$error_invalid_long_id$reset"
	elif [ $1 = "invalid_device_found" ]; then
		echo "$error_prefix $accent$error_invalid_device_found$reset"
	elif [ $1 = "invalid_format" ]; then
		echo "$error_prefix $accent$error_invalid_format$reset"
	elif [ $1 = "device_not_found" ]; then
		echo "$error_prefix $accent$error_device_not_found$reset"
	else
		echo "$error_prefix $accent$1$reset"
	fi
}

# 1=error_name
print_fatal(){
	print_error $1
	exit 1
}

# no parameters
print_banner(){
	echo $header"Multiverse OS:$reset$strong PCI Device$reset$accent '$command_name'$reset$strong Bind Tool$reset"
	echo $accent"================================================================================$reset"
}

# no parameters
print_usage(){
	echo $strong"usage:$reset$accent \`$command_name $device_id_format\`$reset$text or$reset$accent \`$command_name $short_address_format\`$reset$text or$reset$accent \`$command_name $long_address_format\`$reset"
}

# no parameters
print_help(){ 
	print_usage
	echo "$strong  $short_id_length$reset$subheader...$reset$success $short_address_format   $reset$subheader  .......$reset$strong Short PCI Address format.$reset"
	echo "$stron  $device_id_length$reset$subheader...$reset$success $device_id_format    $reset$subheader .....$reset$strong Device ID format.$reset"
	echo "$strong  $long_id_length$reset$subheader..$reset$success $long_address_format   $reset$subheader ...$reset$strong Long PCI Address format.$reset"
}
###############################################################################
##  PCI Devices Shell Functions; Lookup 
##=============================================================================
device_id_lookup(){
	device_id=$(lspci -n | grep $short_pci_address | cut -c 15- | cut -c -$device_id_format_length)
	device_id_length=$(length $device_id)
	print_argument_status
	if [ $(length $device_id) -eq $device_id_format_length ]; then 
		echo $device_id
	else
		print_error "invalid_device_found"
	fi
}

pci_address_lookup(){
	echo "$log_prefix$reset$text Searching for all devices with matching device ID$reset$text[$reset$accent $device_id $reset$text]$reset"
	line_count=0
	lspci -nn | grep $device_id | while read -r line; do 
	line_count=$((line_count+1))
		echo "$log_prefix $reset$accent===========[$reset$header LINE$reset$subheader $line_count$reset$accent ]==================$reset"
		short_pci_address=$(echo "$line" | cut -c -$short_address_format_length)
		long_pci_address="0000:$short_pci_address"
		short_pci_address_length=$(length $short_pci_address)
		long_pci_address_length=$(length $long_pci_address)
		print_argument_status
		#bind_device $device_id $short_pci_address
		bind_device
	done
}

#kernel_module_lookup(){
#	echo $(lspci -knn | grep $short_pci_address -A 2 | grep modules | cut -d : -f 2 | cut -c 2- )
#}
#
#kernel_driver_lookup(){
#	echo $(lspci -knn | grep $short_pci_address -A 3 | grep driver | cut -d : -f 2 | cut -c 2- )
#}

kernel_module_lookup(){
  echo $(lspci -knn | grep $short_pci_address -A 3 | grep modules | cut -c 18-)
}

kernel_driver_lookup(){
  echo $(lspci -knn | grep $short_pci_address -A 2 | grep driver | cut -c 24-)
}


bind_device_to_vfio(){

	bind_device_id=$(echo $device_id | sed 's/:/ /')
	echo "$log_prefix$reset$text bind device id:$reset$accent [$bind_device_id]$reset"
  	echo "$bind_device_id" > "/sys/bus/pci/drivers/vfio-pci/new_id"
}

assign_device_config_space_permissions(){
	chown root:$mv_group $config_space_path
	chmod 0660 $config_space_path
}
###############################################################################
## Bind Device with 'device_id' to vfio-pci for passthrough
##=============================================================================
is_bound_to_vfio(){
	kernel_driver=$(kernel_driver_lookup $device_id)
	if [ $kernel_driver = "vfio-pci" ]; then 
		echo $true
	else
		echo $false
	fi
}

bind_device(){ 
	kernel_module=$(kernel_module_lookup $device_id)
	kernel_driver=$(kernel_driver_lookup $device_id)
	print_argument_status
	sysfs_path="/sys/bus/pci/devices/$long_pci_address"
	config_space_path="/sys/bus/pci/devices/$long_pci_address/config"
	unbind_fd_path="/sys/bus/pci/devices/$long_pci_address/driver/unbind"
	echo $header"$log_prefix Attempting to configure$strong PCI device"$reset
	echo ""
	echo $text"    Device ID  ............"$reset$accent $device_id         $reset
	echo $text"    PCI Address  .........."$reset$accent $short_pci_address $reset
	echo $text"    Full PCI Address  ....."$reset$accent $long_pci_address  $reset
	echo $text"    SysFS Path  ..........."$reset$accent $sysfs_path        $reset
	echo $text"    Unbind FD Path  ......."$reset$accent $unbind_fd_path    $reset
	echo $text"    Config Space Path  ...."$reset$accent $config_space_path $reset
	echo $text"    Kernel Module  ........"$reset$accent $kernel_module     $reset
	echo $text"    Kernel Driver  ........"$reset$accent $kernel_driver     $reset
	echo ""
	echo $text"$log_prefix$reset$text Checking if device ID$reset$strong [$reset$header$device_id$reset$strong]$reset$text is already unbound by checking if unbound is present ...$reset"

	echo "$log_prefix If [$kernel_driver] is 'vfio-pci' then the device is ready to be passed through to a VM."

	## TODO: Does this need to be -a -z $unbind_fd_path?
	if [ $kernel_driver = "vfio-pci" ]; then 
		echo $accent"[$reset$success Success $reset$accent] The specified device with ID$reset$text [$reset$strong$device_id$reset$strong]$reset$accent was already unbound and assignable to virtual machines (VMs).$reset"
		exit 0
	elif [ -z $unbind_fd_path ]; then
		echo "$success_prefix$reset$text Device already unbound! Now attempting to bind to$reset$accent 'vfio-pci'$reset$text...$reset"

	else
		echo "LONG PCI ADDRESS: $long_pci_address"	

		echo "$log_prefix$reset$text Attempting to unbind by passing PCI Address [$long_pci_address] to:$reset$accent $unbind_fd_path$reset"
		echo $long_pci_address > "$unbind_fd_path"
		if [ ! -z $unbind_fd_path ]; then
			echo "$success_prefix$text Device with ID$reset$accent [$reset$header$device_id$reset$accent]$reset$text unbound and ready to be bound to$reset$accent 'vfio-pci'.$reset"
		fi
		echo $text"$log_prefix$reset$text Attempting to bind device to $reset$accent'vfio-pci'$reset$text to make it assignable using device ID [$device_id].$reset"
		bind_device_to_vfio
		if [ is_bound_to_vfio -eq $true ]; then
			echo $text"$success_prefix$reset$text PCI Device$reset$accent can now be directly assigned to a virtual machine (VM).$reset"
		else
			kernel_driver=$(kernel_driver_lookup $device_id)
			if [ $kernel_driver = "vfio-pci" ]; then 
				echo "IS BOUND!"
			else
				echo "IS NOT BOUND!"
			fi
		fi
	fi
	echo "$log_prefix$reset$text Assinging config space permissions (this is fixed in newer configs).$reset"
	assign_device_config_space_permissions $config_space_path
}

###############################################################################
## Tests
#==============================================================================
# no parameters
test_color_text(){
	color_text "purple" "Purple text color"
	color_text "white" "White text color"
	color_text "gray" "Gray text color"
	color_text "cyan" "Cyan text color"
	color_text "blue" "Blue text color"
	color_text "green" "Green text color"
	color_text "yellow" "Yellow text color"
	color_text "red" "Red text color"
	color_text "reset" "Default text color"
}

###############################################################################
## Argument / User Input
#==============================================================================
print_argument_status(){
	if [ ! $argument_format = "" ]; then
		echo "$log_prefix$reset$text Input is type$reset$accent [$reset$success $argument_format$reset$accent:$reset$success$argument$reset$accent ]$reset$text, length$reset$accent [$reset$success $argument_length $reset$accent]$reset"
	else 
		echo "$log_prefix User input is: [$argument] with a length of $argument_length"
	fi
	if [ ! $short_pci_address = "" ]; then 
		echo "$log_prefix  |->[$reset$success SHORT$reset$accent ]$reset$text PCI Address format is$reset$accent [$reset$success $short_pci_address$reset$accent ]$reset$text with a length of$reset$accent [$reset$success $short_pci_address_length $reset$accent]$reset"
	fi
	if [ ! $long_pci_address = "" ]; then 
		echo "$log_prefix  '->[$reset$success LONG$reset$accent ]$reset$text PCI Address format is$reset$accent [$reset$success $long_pci_address$reset$accent ]$reset$text with a length of$reset$accent [$reset$success $long_pci_address_length $reset$accent]$reset"
	fi
	if [ ! $device_id = "" ]; then 
		echo "$log_prefix    [$reset$success DEVICE ID$reset$accent ]$reset$text Device ID format is$reset$accent [$reset$success $device_id $reset$accent ]$reset$text with a length of$reset$accent [$reset$success $device_id_length $reset$accent]$reset"
	fi
}

validate_input(){
	argument=$1
	argument_length=$(length $argument)
	if [ $argument_length -eq $device_id_format_length ]; then
		argument_format="device_id"
		device_id=$argument
		device_id_length=$argument_length
		print_argument_status
	elif [ $argument_length -eq $short_address_format_length ]; then
		argument_format="short_address"
		short_pci_address=$argument
		short_pci_address_length=$argument_length
		long_pci_address="0000:"$short_pci_address
		long_pci_address_length=$(length $long_pci_address)
		print_argument_status
		if [ $long_pci_address_length -eq $long_address_format_length -a $short_pci_address_length -eq $short_address_format_length ]; then
			echo "$log_prefix$reset$text Successfully parsed all forms of PCI Address$reset"
		else
			print_fatal "invalid_input"
		fi
	elif [ $argument_length -eq $long_address_format_length ]; then
		argument_format="long_address"
		long_pci_address=$argument
		long_pci_address_length=$argument_length
		short_pci_address="$(echo $1 | cut -c 6-)"
		short_pci_address_length=$(length $short_pci_address)
		print_argument_status
		if [ $long_pci_address_length -eq $long_address_format_length -a $short_pci_address_length -eq $short_address_format_length ]; then
			echo "$log_prefix$reset$text Successfully parsed all forms of PCI Address$reset"
		else
			print_error "invalid_input"
		fi
	else
		print_error "invalid_input"
	fi
}
##==============================================================================
## MAIN_LOOP
##==============================================================================
# 1=argument
main(){
	print_banner
	exit_if_not_root
	validate_input $1
	if [ $argument_format = "device_id" ]; then
		pci_address_lookup
	else
		device_id_lookup
		device_id_length=$(length $device_id)
		echo "$success_prefix Successfully looked up device id: $device_id"
		echo "$log_prefix Device lookup by$reset$accent [$reset$success PCI Address$reset$accent ]$rest$text; maximum devices that can be bound is 1."
		bind_device $bind_device $short_pci_address
	fi
}

#bind_device $device_id $pci_address
###############################################################################
## Main Loop Execution
main $1
