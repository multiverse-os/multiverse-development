 #xpra start ssh:app.firefox.debian:100 --start=gnome-terminal

xpra start ssh:app.firefox.debian:100 --start-child=xterm --start-via-proxy=no --xvfb="/usr/bin/Xwayland -rootless -noreset" 
