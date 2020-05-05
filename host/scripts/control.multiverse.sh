#!/bin/sh
###############################################################################

# Variables
# $1 - VM Action
# $2 - VM Class OR VM Name

# $vm - VM action will be performed on
# $action - Action to be performed on VM

# =====================================================

#
# Multiverse OS: Control Script
#

#
# A basic outline using bash to sketch the
# basic requirements of a shell based 
# command-line tool that is similar to virsh
# BUT SIMPLIFIED, less redudant.
#

#
# And most importantly more intuitive
# and concious of Multiverse VM classes (Router,
# App, Service, Controller) and the rules
# associated with them. 
#

# =====================================================

if [ -z "$1" ]; then
  echo "[!] First argument missing, must be 'VM action':"
  echo "    Action: [ 'Start', 'Stop', 'Restart', 'Shutdown' ]"

  # TODO: If elsif check to set available actions to virsh usable command

fi

# TODO: Check if valid VM Action

if [ -z "$2" ]; then
  echo "\n[!] Second argument missing, must be 'VM Name or VM Class':"
  echo "    Class: [ 'controller' (c), 'applicaiton' (a), 'router' (r), 'service' (s) ]"

  if [ "$2" -eq "controller" -o "$2" -eq "c" ]; then
	  $vm="controller.gravity.multiverse"
  else
	  echo "[Error] Invalid VM Class or Name, confirm it exists and try again,..."
	  exit 1
  fi

  echo "    Name: [ 'star.whonix-router', 'universe.router', 'controller.gravity', ... ]"
fi

# TODO: Check if valid class, check if valid name 
if [ -n "$1" -a -n "$2" ]; then 
	echo "Executing 'virsh' ..."
	echo "\$~> virsh $1 $2 "
else
	echo -e "\n[Error] Invalid arguaments." 
	exit 1
fi

