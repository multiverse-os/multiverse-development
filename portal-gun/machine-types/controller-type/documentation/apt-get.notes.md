

# Install virt-manager, git, pass, tor, goalng, ruby, 

# Create bridges and allow  them with /etc/qemu/bridge.conf
# Add storage pools

# Setup passthrough disks and decryption, then symbolic link from Vault to home folder

# Remove firefox (in ubuntu this works)

# edit /etc/security/limits.conf
# edit /etc/sysctl.conf
# edit /etc/hosts
# edit /etc/udev/rules.d/10-multiverse.rules
# edit /etc/iproute2/rt_tables

# Add add-routing-tables.sh script (needs to be modified to be more flexible)

# Add /var/multiver/ directory 

# Add to ~/.bashrc:
```
alias vim="nvim"
alias ssh="ssh -X"
alias apt-install="sudo apt-get install --no-install-recommends --no-install-suggests"
alias apt-remove="sudo apt-get --purge remove"
```

sudo usermod -a -G libvirt user
