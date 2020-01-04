#
# Multiverse OS: Basic VM Startup Script
##########################################################################

DEFAULT_CONTROLLER="ubuntu18.controller"

#
# virsh list --all
##########################################################################
#  Id    Name                           State
# ----------------------------------------------------
######## CONTROLLER OPTIONS ##############################################
#  -     debian10.controller            shut off
#  -     games.ubuntu18.controller      shut off
#  -     ubuntu18.controller            shut off
######## CLUSTER NETWORKING ##############################################
#  -     universe0.bridge0.router       shut off
#  -     universe0.bridge1.router       shut off
#  -     universe1.bridge0.router       shut off
#  -     universe1.bridge1.router       shut off
#  -     universe2.bridge0.router       shut off
#  -     universe2.bridge1.router       shut off
##########################################################################

echo "Multiverse OS: Starting Host Cluster"
echo "===================================="
echo "Booting [ 3 ] universeX bridge0 routers..."

echo "Starting universe0.bridge0.router..."
virsh start universe0.bridge0.router
echo "Starting universe1.bridge0.router..."
virsh start universe1.bridge0.router
echo "Starting universe2.bridge0.router..."
virsh start universe2.bridge0.router

echo "Waiting 5 seconds to allow for boot processes to complete..."
sleep 5


echo "Booting [ 3 ] universeX bridge1 (GALAXY) routers..."
echo "Starting universe0.bridge1.router..."
virsh start universe0.bridge1.router
echo "Starting universe1.bridge1.router..."
virsh start universe1.bridge1.router
echo "Starting universe2.bridge1.router..."
virsh start universe2.bridge1.router

echo "Waiting 5 seconds to allow for boot processes to complete..."
sleep 5

echo "-----------------------------"
echo "Multiverse Networking Online!"
echo "-----------------------------"

echo "Starting the default Multiverse Controller VM..."
virsh start $DEFAULT_CONTROLLER

echo "$DEFAULT_CONTROLLER has been started..."
echo "---------------------------------------"
echo "Multiverse OS cluster startup complete!"
echo "---------------------------------------"


