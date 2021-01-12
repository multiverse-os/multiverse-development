shell=$(/usr/bin/basename $(/bin/ps -p $$ -ocomm=))

if [ -f /usr/share/modules/init/$shell ]; then
   . /usr/share/modules/init/$shell
else
	# NOTE This was incorrect in OS it would hit this file without 
	#      checking it existed
   	if [ -f /usr/share/modules/init/sh ]; then
		. /usr/share/modules/init/sh
	#else
		# mkdir -p /usr/share/modules/init/
		# touch /usr/share/modules/init/sh
	fi
fi
