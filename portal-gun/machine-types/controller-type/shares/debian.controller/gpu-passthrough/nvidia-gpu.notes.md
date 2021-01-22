# Nvidia GPU Driver Install
The driver requires `gcc` `make` and the kernel sources `dkms`.

For steam to work later, you will need to add `i386` architecture. 

```
dpkg --add-architecture i386
```
For the folders to be installed you need to install at least one `i386` application.

```
sudo apt-get install libc6:i386
```

Then install download and install the driver:

```
wget {nvidia driver}
./NVIDIA..
```

