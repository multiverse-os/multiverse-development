#!/bin/bash

sudo chown root:libvirt /etc/qemu
sudo chown root:libvirt /etc/qemu/bridge.conf 
sudo chmod 0640 /etc/qemu/bridge.conf 
