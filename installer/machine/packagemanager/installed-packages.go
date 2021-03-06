package packagemanager

func InstalledPackages() []string {
	return []string{"adduser",
	"apparmor",
	"apt-listchanges",
	"apt-utils",
	"apt",
	"arj",
	"at",
	"augeas-lenses",
	"autoconf",
	"automake",
	"autopoint",
	"autotools-dev",
	"avahi-daemon",
	"base-files",
	"base-passwd",
	"bash-completion",
	"bash",
	"bind9-host",
	"bind9-libs",
	"binfmt-support",
	"binstats",
	"binutils-common",
	"binutils-dev",
	"binutils-x86-64-linux-gnu",
	"binutils",
	"binwalk",
	"bison",
	"blt",
	"boot-info-script",
	"bootpc",
	"bsdmainutils",
	"bsdutils",
	"build-essential",
	"busybox",
	"bzip2-doc",
	"bzip2",
	"ca-certificates-java",
	"ca-certificates",
	"chkrootkit",
	"console-setup-linux",
	"console-setup",
	"coreutils",
	"cpio",
	"cpp-8",
	"cpp-9",
	"cpp",
	"cramfsswap",
	"cron",
	"cryptsetup-bin",
	"cryptsetup-initramfs",
	"cryptsetup-run",
	"cryptsetup",
	"cups-pk-helper",
	"curl",
	"dash",
	"dbus-user-session",
	"dbus-x11",
	"dbus",
	"dconf-cli",
	"dconf-gsettings-backend",
	"dconf-service",
	"dctrl-tools",
	"debconf-i18n",
	"debconf",
	"debhelper",
	"debian-archive-keyring",
	"debian-faq",
	"debian-keyring",
	"debianutils",
	"default-jre-headless",
	"desktop-base",
	"desktop-file-utils",
	"device-tree-compiler",
	"devio",
	"devrplay3",
	"devscripts",
	"devtodo",
	"dh-autoreconf",
	"dh-strip-nondeterminism",
	"dictionaries-common",
	"diffstat",
	"diffutils",
	"dirmngr",
	"discover-data",
	"discover",
	"distro-info-data",
	"dmeventd",
	"dmidecode",
	"dmsetup",
	"dns-root-data",
	"dnsmasq-base",
	"doc-debian",
	"dosfstools",
	"dpkg-dev",
	"dpkg",
	"dput",
	"dwz",
	"e2fslibs",
	"e2fsprogs-l10n",
	"e2fsprogs",
	"ed",
	"efibootmgr",
	"efingerd",
	"efitools",
	"efivar",
	"elfrc",
	"elfutils",
	"emacsen-common",
	"equivs",
	"exim4-base",
	"exim4-config",
	"exim4-daemon-light",
	"fakeroot",
	"fdisk",
	"file",
	"findutils",
	"firmware-linux-free",
	"flashbench",
	"flashrom",
	"flex",
	"fontconfig-config",
	"fontconfig",
	"fonts-cantarell",
	"fonts-dejavu-core",
	"fonts-droid-fallback",
	"fonts-font-awesome",
	"fonts-lato",
	"fonts-liberation",
	"fonts-lyx",
	"fonts-noto-mono",
	"fonts-quicksand",
	"fonts-urw-base35",
	"freeglut3",
	"fuse",
	"g++-9",
	"g++",
	"gawk",
	"gcc-10-base",
	"gcc-6-base",
	"gcc-8-base",
	"gcc-8",
	"gcc-9-base",
	"gcc-9",
	"gcc",
	"gettext-base",
	"gettext",
	"gfortran-9",
	"gfortran",
	"ghostscript",
	"gir1.2-glib-2.0",
	"git-man",
	"git",
	"glib-networking-common",
	"glib-networking-services",
	"glib-networking",
	"gnome-accessibility-themes",
	"gnome-backgrounds",
	"gnome-menus",
	"gnome-themes-extra-data",
	"gnome-themes-extra",
	"gnupg-agent",
	"gnupg-l10n",
	"gnupg-utils",
	"gnupg2",
	"gnupg",
	"gnustep-base-common",
	"gnustep-base-runtime",
	"gnustep-common",
	"golang-1.14-doc",
	"golang-1.14-go",
	"golang-1.14-src",
	"golang-1.14",
	"golang-any",
	"golang-doc",
	"golang-go",
	"golang-src",
	"golang",
	"gpg-agent",
	"gpg-wks-client",
	"gpg-wks-server",
	"gpg",
	"gpgconf",
	"gpgsm",
	"gpgv",
	"grep",
	"groff-base",
	"grub-common",
	"grub-pc-bin",
	"grub-pc",
	"grub2-common",
	"grub2",
	"gsettings-desktop-schemas",
	"gstreamer1.0-libav",
	"gstreamer1.0-plugins-base",
	"gstreamer1.0-plugins-good",
	"gstreamer1.0-plugins-ugly",
	"gstreamer1.0-x",
	"gtk2-engines-pixbuf",
	"guile-2.2-libs",
	"gvfs-bin",
	"gvfs-common",
	"gzip",
	"hdparm",
	"hostname",
	"i965-va-driver",
	"iamerican",
	"ibritish",
	"ibverbs-providers",
	"ienglish-common",
	"ifupdown",
	"init-system-helpers",
	"init",
	"initramfs-tools-core",
	"initramfs-tools",
	"installation-report",
	"intel-media-va-driver",
	"intltool-debian",
	"iproute2",
	"iptables",
	"iputils-ping",
	"ipxe-qemu",
	"isc-dhcp-client",
	"isc-dhcp-common",
	"iso-codes",
	"ispell",
	"java-common",
	"javascript-common",
	"kbd",
	"keyboard-configuration",
	"klibc-utils",
	"kmod",
	"krb5-locales",
	"kvmtool",
	"laptop-detect",
	"less",
	"liba52-0.7.4",
	"libaa1",
	"libaacs0",
	"libacl1",
	"libafflib0v5",
	"libaio1",
	"libalgorithm-diff-perl",
	"libalgorithm-diff-xs-perl",
	"libalgorithm-merge-perl",
	"libaom0",
	"libapparmor1",
	"libapt-pkg-perl",
	"libapt-pkg6.0",
	"libarchive-cpio-perl",
	"libarchive-zip-perl",
	"libargon2-1",
	"libarray-intspan-perl",
	"libasan5",
	"libasm1",
	"libasound2-data",
	"libasound2",
	"libass9",
	"libassuan0",
	"libasync-mergepoint-perl",
	"libasyncns0",
	"libatk-adaptor",
	"libatk-bridge2.0-0",
	"libatk1.0-0",
	"libatk1.0-data",
	"libatomic1",
	"libatspi2.0-0",
	"libattr1",
	"libaudit-common",
	"libaudit1",
	"libaugeas0",
	"libauthen-sasl-perl",
	"libavahi-client3",
	"libavahi-common-data",
	"libavahi-common3",
	"libavahi-core7",
	"libavc1394-0",
	"libavcodec58",
	"libavfilter7",
	"libavformat58",
	"libavutil56",
	"libb-hooks-endofscope-perl",
	"libb-hooks-op-check-perl",
	"libbdplus0",
	"libbinutils",
	"libbison-dev",
	"libblas3",
	"libblkid1",
	"libbluray2",
	"libboost-filesystem1.67.0",
	"libboost-iostreams1.67.0",
	"libboost-system1.67.0",
	"libboost-thread1.67.0",
	"libbrlapi0.7",
	"libbrotli1",
	"libbs2b0",
	"libbsd0",
	"libbz2-1.0",
	"libc-bin",
	"libc-dev-bin",
	"libc-l10n",
	"libc6-dev",
	"libc6",
	"libcaca0",
	"libcacard0",
	"libcairo-gobject2",
	"libcairo2",
	"libcanberra0",
	"libcap-ng0",
	"libcap2-bin",
	"libcap2",
	"libcapstone3",
	"libcapture-tiny-perl",
	"libcaribou-gtk-module",
	"libcbor0",
	"libcc1-0",
	"libcdio18",
	"libcdparanoia0",
	"libchromaprint1",
	"libclass-inspector-perl",
	"libclass-isa-perl",
	"libclass-method-modifiers-perl",
	"libclass-xsaccessor-perl",
	"libclone-perl",
	"libcodec2-0.9",
	"libcom-err2",
	"libcomerr2",
	"libcommon-sense-perl",
	"libconst-fast-perl",
	"libcontextual-return-perl",
	"libconvert-binhex-perl",
	"libcpanel-json-xs-perl",
	"libcroco3",
	"libcrypt-dev",
	"libcrypt1",
	"libcryptsetup12",
	"libcryptsetup4",
	"libctf-nobfd0",
	"libctf0",
	"libcups2",
	"libcurl3-gnutls",
	"libcurl4",
	"libdaemon0",
	"libdata-dump-perl",
	"libdata-optlist-perl",
	"libdate-manip-perl",
	"libdatrie1",
	"libdb5.3",
	"libdbus-1-3",
	"libdconf1",
	"libdebconfclient0",
	"libdebhelper-perl",
	"libdevel-callchecker-perl",
	"libdevel-size-perl",
	"libdevmapper-dev",
	"libdevmapper-event1.02.1",
	"libdevmapper1.02.1",
	"libdigest-bubblebabble-perl",
	"libdigest-hmac-perl",
	"libdiscover2",
	"libdistro-info-perl",
	"libdns-export1110",
	"libdns-export162",
	"libdouble-conversion3",
	"libdpkg-perl",
	"libdrm-amdgpu1",
	"libdrm-common",
	"libdrm-intel1",
	"libdrm-nouveau2",
	"libdrm-radeon1",
	"libdrm2",
	"libdv4",
	"libdvdread8",
	"libdw1",
	"libdynaloader-functions-perl",
	"libebml4v5",
	"libedit2",
	"libefiboot1",
	"libefivar1",
	"libegl-mesa0",
	"libegl1",
	"libelf1",
	"libemail-valid-perl",
	"libencode-locale-perl",
	"libepoxy0",
	"liberror-perl",
	"libestr0",
	"libevdev2",
	"libevent-2.1-7",
	"libewf2",
	"libexpat1",
	"libexporter-tiny-perl",
	"libext2fs2",
	"libfakeroot",
	"libfastjson4",
	"libfastutil-java",
	"libfcgi-perl",
	"libfdisk1",
	"libfdt1",
	"libffi6",
	"libffi7",
	"libfftw3-double3",
	"libfido2-1",
	"libfile-basedir-perl",
	"libfile-chdir-perl",
	"libfile-desktopentry-perl",
	"libfile-fcntllock-perl",
	"libfile-find-rule-perl",
	"libfile-homedir-perl",
	"libfile-listing-perl",
	"libfile-stripnondeterminism-perl",
	"libfile-which-perl",
	"libfl-dev",
	"libfl2",
	"libflac8",
	"libflashrom-dev",
	"libflashrom1",
	"libflite1",
	"libfont-afm-perl",
	"libfont-ttf-perl",
	"libfontconfig1",
	"libfreetype6",
	"libfribidi0",
	"libfstrm0",
	"libftdi1-2",
	"libfuse2",
	"libfuture-perl",
	"libgbm1",
	"libgc1c2",
	"libgcc-8-dev",
	"libgcc-9-dev",
	"libgcc-s1",
	"libgcc1",
	"libgcrypt20",
	"libgdbm-compat4",
	"libgdbm3",
	"libgdbm6",
	"libgdk-pixbuf2.0-0",
	"libgdk-pixbuf2.0-bin",
	"libgdk-pixbuf2.0-common",
	"libgetopt-long-descriptive-perl",
	"libgfapi0",
	"libgfortran-9-dev",
	"libgfortran5",
	"libgfrpc0",
	"libgfxdr0",
	"libgirepository-1.0-1",
	"libgit-wrapper-perl",
	"libgitlab-api-v4-perl",
	"libgl1-mesa-dri",
	"libgl1-mesa-glx",
	"libgl1",
	"libglapi-mesa",
	"libglib2.0-0",
	"libglib2.0-bin",
	"libglib2.0-data",
	"libglu1-mesa",
	"libglusterfs0",
	"libglvnd0",
	"libglx-mesa0",
	"libglx0",
	"libgme0",
	"libgmp10",
	"libgnustep-base1.26",
	"libgnutls-dane0",
	"libgnutls30",
	"libgomp1",
	"libgpg-error-l10n",
	"libgpg-error0",
	"libgpgme11",
	"libgpm2",
	"libgraphite2-3",
	"libgs9-common",
	"libgs9",
	"libgsasl7",
	"libgsf-1-114",
	"libgsf-1-common",
	"libgsf-bin",
	"libgsm1",
	"libgssapi-krb5-2",
	"libgstreamer-plugins-base1.0-0",
	"libgstreamer1.0-0",
	"libgudev-1.0-0",
	"libharfbuzz0b",
	"libhash-fieldhash-perl",
	"libhogweed4",
	"libhogweed5",
	"libhtml-form-perl",
	"libhtml-format-perl",
	"libhtml-parser-perl",
	"libhtml-tagset-perl",
	"libhtml-tree-perl",
	"libhttp-cookies-perl",
	"libhttp-daemon-perl",
	"libhttp-date-perl",
	"libhttp-message-perl",
	"libhttp-negotiate-perl",
	"libhttp-tiny-multipart-perl",
	"libibverbs1",
	"libice6",
	"libicu63",
	"libident",
	"libidn11",
	"libidn2-0",
	"libiec61883-0",
	"libigdgmm11",
	"libijs-0.35",
	"libimagequant0",
	"libimobiledevice6",
	"libimport-into-perl",
	"libinput-bin",
	"libinput10",
	"libio-async-loop-epoll-perl",
	"libio-async-perl",
	"libio-html-perl",
	"libio-prompter-perl",
	"libio-pty-perl",
	"libio-sessiondata-perl",
	"libio-socket-ssl-perl",
	"libio-string-perl",
	"libio-stringy-perl",
	"libip4tc0",
	"libip4tc2",
	"libip6tc0",
	"libip6tc2",
	"libipc-run-perl",
	"libipc-system-simple-perl",
	"libiptc0",
	"libisc-export1105",
	"libisc-export160",
	"libiscsi7",
	"libisl22",
	"libitm1",
	"libjack-jackd2-0",
	"libjansson4",
	"libjbig0",
	"libjbig2dec0",
	"libjim0.79",
	"libjpeg62-turbo",
	"libjs-jquery-ui",
	"libjs-jquery",
	"libjson-c4",
	"libjson-maybexs-perl",
	"libjson-perl",
	"libjson-xs-perl",
	"libk5crypto3",
	"libkeyutils1",
	"libklibc",
	"libkmod2",
	"libkrb5-3",
	"libkrb5support0",
	"libksba8",
	"libkyotocabinet16v5",
	"liblapack3",
	"liblbfgsb0",
	"liblcms2-2",
	"libldap-2.4-2",
	"libldap-common",
	"liblilv-0-0",
	"liblinux-epoll-perl",
	"liblist-compare-perl",
	"liblist-moreutils-perl",
	"liblist-someutils-perl",
	"liblist-someutils-xs-perl",
	"libllvm9",
	"liblmdb0",
	"liblocale-gettext-perl",
	"liblockfile-bin",
	"liblog-any-adapter-screen-perl",
	"liblog-any-perl",
	"liblogging-stdlog0",
	"liblognorm5",
	"liblsan0",
	"libltdl-dev",
	"libltdl7",
	"libluajit-5.1-2",
	"libluajit-5.1-common",
	"liblvm2cmd2.03",
	"liblwp-mediatypes-perl",
	"liblwp-protocol-https-perl",
	"liblz4-1",
	"liblzma5",
	"liblzo2-2",
	"libmagic-mgc",
	"libmagic1",
	"libmail-sendmail-perl",
	"libmailtools-perl",
	"libmailutils6",
	"libmariadb3",
	"libmatroska6v5",
	"libmaxminddb0",
	"libmime-tools-perl",
	"libmnl0",
	"libmodbus-dev",
	"libmodbus5",
	"libmodplug-dev",
	"libmodplug1",
	"libmodule-implementation-perl",
	"libmodule-runtime-perl",
	"libmoo-perl",
	"libmoox-aliases-perl",
	"libmoox-struct-perl",
	"libmount1",
	"libmp3lame0",
	"libmpc3",
	"libmpdec2",
	"libmpeg2-4",
	"libmpfr6",
	"libmpg123-0",
	"libmpx2",
	"libmsgpackc2",
	"libmtdev1",
	"libmysofa1",
	"libnamespace-autoclean-perl",
	"libnamespace-clean-perl",
	"libncurses5",
	"libncurses6",
	"libncursesw5",
	"libncursesw6",
	"libnet-dns-perl",
	"libnet-dns-sec-perl",
	"libnet-domain-tld-perl",
	"libnet-http-perl",
	"libnet-ip-perl",
	"libnet-libidn-perl",
	"libnet-smtp-ssl-perl",
	"libnet-ssleay-perl",
	"libnetcf1",
	"libnetfilter-conntrack3",
	"libnettle6",
	"libnettle7",
	"libnewt0.52",
	"libnfnetlink-dev",
	"libnfnetlink0",
	"libnfs13",
	"libnftables1",
	"libnftnl-dev",
	"libnftnl11",
	"libnghttp2-14",
	"libnl-3-200",
	"libnl-route-3-200",
	"libnorm1",
	"libnotify4",
	"libnpth0",
	"libnspr4",
	"libnss-mdns",
	"libnss-mymachines",
	"libnss-systemd",
	"libnss3",
	"libntlm0",
	"libnuma1",
	"libnumber-compare-perl",
	"libnumber-range-perl",
	"libobjc4",
	"libobject-id-perl",
	"libogg0",
	"libopencore-amrnb0",
	"libopencore-amrwb0",
	"libopenjp2-7",
	"libopenmpt0",
	"libopus0",
	"liborc-0.4-0",
	"libosmpbf-java",
	"libossp-uuid-perl",
	"libossp-uuid16",
	"libp11-kit0",
	"libpackage-stash-perl",
	"libpackage-stash-xs-perl",
	"libpam-cap",
	"libpam-gnome-keyring",
	"libpam-modules-bin",
	"libpam-modules",
	"libpam-runtime",
	"libpam-systemd",
	"libpam0g",
	"libpango-1.0-0",
	"libpangocairo-1.0-0",
	"libpangoft2-1.0-0",
	"libpaper-utils",
	"libpaper1",
	"libparams-classify-perl",
	"libparams-util-perl",
	"libparams-validate-perl",
	"libparted2",
	"libpath-iterator-rule-perl",
	"libpath-tiny-perl",
	"libpcap0.8",
	"libpci3",
	"libpciaccess0",
	"libpcre2-16-0",
	"libpcre2-32-0",
	"libpcre2-8-0",
	"libpcre2-dev",
	"libpcre2-posix2",
	"libpcre3",
	"libpcsclite1",
	"libperl4-corelibs-perl",
	"libperl5.30",
	"libperlio-gzip-perl",
	"libpgm-5.2-0",
	"libpipeline1",
	"libpixman-1-0",
	"libplist3",
	"libplymouth4",
	"libpmem1",
	"libpng16-16",
	"libpod-constants-perl",
	"libpolkit-agent-1-0",
	"libpolkit-gobject-1-0",
	"libpopt0",
	"libpostproc55",
	"libprocps6",
	"libprocps8",
	"libprotobuf-c1",
	"libprotobuf-java",
	"libproxy1v5",
	"libpsl5",
	"libpugixml1v5",
	"libpulse0",
	"libpython-stdlib",
	"libpython2-stdlib",
	"libpython2.7-minimal",
	"libpython2.7-stdlib",
	"libpython3-stdlib",
	"libpython3.7-minimal",
	"libpython3.7-stdlib",
	"libpython3.8-minimal",
	"libpython3.8-stdlib",
	"libpython3.8",
	"libqt5core5a",
	"libqt5dbus5",
	"libqt5designer5",
	"libqt5gui5",
	"libqt5help5",
	"libqt5network5",
	"libqt5opengl5",
	"libqt5printsupport5",
	"libqt5sql5-sqlite",
	"libqt5sql5",
	"libqt5svg5",
	"libqt5test5",
	"libqt5widgets5",
	"libqt5xml5",
	"libquadmath0",
	"librados2",
	"libraw1394-11",
	"librbd1",
	"librdmacm1",
	"libre-engine-re2-perl",
	"libre2-6",
	"libreadline5",
	"libreadline7",
	"libreadline8",
	"libreadonly-perl",
	"libref-util-perl",
	"libref-util-xs-perl",
	"libregexp-pattern-license-perl",
	"libregexp-pattern-perl",
	"librole-tiny-perl",
	"librplay3",
	"librsvg2-2",
	"librsvg2-common",
	"librtmp1",
	"librubberband2",
	"libruby2.7",
	"libruby",
	"libsamplerate0",
	"libsasl2-2",
	"libsasl2-modules-db",
	"libsasl2-modules",
	"libsdl1.2debian",
	"libseccomp2",
	"libselinux1-dev",
	"libselinux1",
	"libsemanage-common",
	"libsemanage1",
	"libsensors-config",
	"libsensors5",
	"libsepol1-dev",
	"libsepol1",
	"libserd-0-0",
	"libsereal-decoder-perl",
	"libsereal-encoder-perl",
	"libsereal-perl",
	"libshine3",
	"libshout3",
	"libsidplay1v5",
	"libsigsegv2",
	"libslang2",
	"libslirp0",
	"libsm6",
	"libsmartcols1",
	"libsnappy1v5",
	"libsndfile1",
	"libsoap-lite-perl",
	"libsodium23",
	"libsord-0-0",
	"libsort-key-perl",
	"libsort-versions-perl",
	"libsoup2.4-1",
	"libsoxr0",
	"libspeex1",
	"libspice-server1",
	"libsqlite3-0",
	"libsratom-0-0",
	"libss2",
	"libssh-4",
	"libssh-gcrypt-4",
	"libssh2-1",
	"libssl1.0.2",
	"libssl1.1",
	"libstdc++-9-dev",
	"libstdc++6",
	"libstrictures-perl",
	"libstring-copyright-perl",
	"libstring-escape-perl",
	"libstring-shellquote-perl",
	"libstruct-dumb-perl",
	"libsub-exporter-perl",
	"libsub-exporter-progressive-perl",
	"libsub-identify-perl",
	"libsub-install-perl",
	"libsub-name-perl",
	"libsub-override-perl",
	"libsub-quote-perl",
	"libswitch-perl",
	"libswresample3",
	"libswscale5",
	"libsys-hostname-long-perl",
	"libsystemd0",
	"libtag1v5-vanilla",
	"libtag1v5",
	"libtask-weaken-perl",
	"libtasn1-6",
	"libtcl8.6",
	"libtdb1",
	"libterm-readkey-perl",
	"libtermkey1",
	"libtest-fatal-perl",
	"libtest-refcount-perl",
	"libtext-charwidth-perl",
	"libtext-glob-perl",
	"libtext-iconv-perl",
	"libtext-levenshtein-perl",
	"libtext-wrapi18n-perl",
	"libthai-data",
	"libthai0",
	"libtheora0",
	"libtiff5",
	"libtimedate-perl",
	"libtinfo5",
	"libtinfo6",
	"libtirpc-common",
	"libtirpc3",
	"libtk8.6",
	"libtool",
	"libtry-tiny-perl",
	"libtsan0",
	"libtsk13",
	"libtwolame0",
	"libtype-tiny-perl",
	"libtype-tiny-xs-perl",
	"libtypes-serialiser-perl",
	"libubsan1",
	"libuchardet0",
	"libudev-dev",
	"libudev1",
	"libunbound8",
	"libunibilium4",
	"libunicode-utf8-perl",
	"libunistring0",
	"libunistring2",
	"libunwind8",
	"liburi-perl",
	"libusb-0.1-4",
	"libusb-1.0-0",
	"libusbauth-configparser1",
	"libusbmuxd6",
	"libusbredirparser1",
	"libustr-1.0-1",
	"libuuid1",
	"libuv1",
	"libv4l-0",
	"libv4lconvert0",
	"libva-drm2",
	"libva-x11-2",
	"libva2",
	"libvariable-magic-perl",
	"libvdeplug2",
	"libvdpau-va-gl1",
	"libvdpau1",
	"libvidstab1.1",
	"libvirglrenderer1",
	"libvirt-clients",
	"libvirt-daemon-driver-lxc",
	"libvirt-daemon-driver-qemu",
	"libvirt-daemon-driver-vbox",
	"libvirt-daemon-driver-xen",
	"libvirt-daemon-system-systemd",
	"libvirt-daemon-system",
	"libvirt-daemon",
	"libvirt-dbus",
	"libvirt-dev",
	"libvirt-glib-1.0-0",
	"libvirt0",
	"libvisual-0.4-0",
	"libvncserver1",
	"libvorbis0a",
	"libvorbisenc2",
	"libvorbisfile3",
	"libvpx6",
	"libvterm0",
	"libvulkan1",
	"libwacom-bin",
	"libwacom-common",
	"libwacom2",
	"libwant-perl",
	"libwavpack1",
	"libwayland-client0",
	"libwayland-server0",
	"libwebp6",
	"libwebpdemux2",
	"libwebpmux3",
	"libwrap0",
	"libwww-perl",
	"libwww-robotrules-perl",
	"libx11-6",
	"libx11-data",
	"libx11-xcb1",
	"libx264-155",
	"libx265-179",
	"libxapian30",
	"libxau6",
	"libxcb-dri2-0",
	"libxcb-dri3-0",
	"libxcb-glx0",
	"libxcb-icccm4",
	"libxcb-image0",
	"libxcb-keysyms1",
	"libxcb-present0",
	"libxcb-randr0",
	"libxcb-render-util0",
	"libxcb-render0",
	"libxcb-shape0",
	"libxcb-shm0",
	"libxcb-sync1",
	"libxcb-util0",
	"libxcb-xfixes0",
	"libxcb-xinerama0",
	"libxcb-xinput0",
	"libxcb-xkb1",
	"libxcb1",
	"libxdamage1",
	"libxdmcp6",
	"libxen-dev",
	"libxencall1",
	"libxendevicemodel1",
	"libxenevtchn1",
	"libxenforeignmemory1",
	"libxengnttab1",
	"libxenmisc4.11",
	"libxenstore3.0",
	"libxentoolcore1",
	"libxentoollog1",
	"libxext6",
	"libxfixes3",
	"libxft2",
	"libxi6",
	"libxkbcommon-x11-0",
	"libxkbcommon0",
	"libxml-libxml-perl",
	"libxml-namespacesupport-perl",
	"libxml-parser-perl",
	"libxml-sax-base-perl",
	"libxml-sax-expat-perl",
	"libxml-sax-perl",
	"libxml-writer-perl",
	"libxml2-utils",
	"libxml2",
	"libxmlrpc-lite-perl",
	"libxmu6",
	"libxmuu1",
	"libxrender1",
	"libxshmfence1",
	"libxslt1.1",
	"libxss1",
	"libxt6",
	"libxtables12",
	"libxtst6",
	"libxv1",
	"libxvidcore4",
	"libxxf86vm1",
	"libyajl2",
	"libyaml-0-2",
	"libyaml-libyaml-perl",
	"libz3-4",
	"libzc-dev",
	"libzc4",
	"libzmq5",
	"libzstd1",
	"libzvbi-common",
	"libzvbi0",
	"licensecheck",
	"lintian",
	"linux-base",
	"linux-compiler-gcc-8-x86",
	"linux-headers-4.19.0-6-all-amd64",
	"linux-headers-4.19.0-6-all",
	"linux-headers-4.19.0-6-amd64",
	"linux-headers-4.19.0-6-cloud-amd64",
	"linux-headers-4.19.0-6-common-rt",
	"linux-headers-4.19.0-6-common",
	"linux-headers-4.19.0-6-rt-amd64",
	"linux-headers-4.19.0-8-all-amd64",
	"linux-headers-4.19.0-8-amd64",
	"linux-headers-4.19.0-8-cloud-amd64",
	"linux-headers-4.19.0-8-common-rt",
	"linux-headers-4.19.0-8-common",
	"linux-headers-4.19.0-8-rt-amd64",
	"linux-image-4.19.0-8-amd64",
	"linux-image-4.9.0-8-amd64",
	"linux-image-5.5.0-2-amd64",
	"linux-kbuild-4.19",
	"linux-libc-dev",
	"locales",
	"login",
	"logrotate",
	"logsave",
	"lsb-base",
	"lsb-release",
	"lsof",
	"lua-luv",
	"lvm2",
	"m4",
	"mailutils-common",
	"mailutils",
	"make",
	"man-db",
	"manpages-dev",
	"manpages",
	"mariadb-common",
	"mawk",
	"mesa-va-drivers",
	"mesa-vdpau-drivers",
	"mesa-vulkan-drivers",
	"mime-support",
	"mitmproxy",
	"mkelfimage",
	"mkgmap",
	"mknfonts.tool",
	"mksh",
	"mktorrent",
	"mkvtoolnix",
	"mount",
	"mtd-utils",
	"multiarch-support",
	"mysql-common",
	"ncompress",
	"ncurses-base",
	"ncurses-bin",
	"ncurses-term",
	"neovim-runtime",
	"neovim",
	"net-tools",
	"netbase",
	"netcat-openbsd",
	"netcat-traditional",
	"nftables",
	"ocl-icd-libopencl1",
	"openhackware",
	"openjdk-11-jre-headless",
	"openssh-client",
	"openssl",
	"os-prober",
	"ovmf",
	"p7zip-full",
	"p7zip",
	"parted",
	"pass",
	"passwd",
	"patch",
	"patchutils",
	"pci.ids",
	"pciutils",
	"perl-base",
	"perl-modules-5.24",
	"perl-modules-5.30",
	"perl-openssl-defaults",
	"perl",
	"pigz",
	"pinentry-curses",
	"pkg-config",
	"plymouth-label",
	"plymouth",
	"po-debconf",
	"policykit-1",
	"poppler-data",
	"powermgmt-base",
	"ppp",
	"procps",
	"psmisc",
	"python-apt-common",
	"python-matplotlib-data",
	"python-minimal",
	"python2-minimal",
	"python2.7-minimal",
	"python2.7",
	"python2",
	"python3-apt",
	"python3-binwalk",
	"python3-blinker",
	"python3-brotli",
	"python3-certifi",
	"python3-cffi-backend",
	"python3-chardet",
	"python3-click",
	"python3-colorama",
	"python3-cryptography",
	"python3-cups",
	"python3-cupshelpers",
	"python3-cycler",
	"python3-dateutil",
	"python3-dbus",
	"python3-debconf",
	"python3-debian",
	"python3-debianbts",
	"python3-decorator",
	"python3-gi",
	"python3-gpg",
	"python3-greenlet",
	"python3-h11",
	"python3-h2",
	"python3-hpack",
	"python3-httplib2",
	"python3-hyperframe",
	"python3-idna",
	"python3-kaitaistruct",
	"python3-kiwisolver",
	"python3-ldap3",
	"python3-magic",
	"python3-matplotlib",
	"python3-minimal",
	"python3-msgpack",
	"python3-neovim",
	"python3-numpy",
	"python3-olefile",
	"python3-opengl",
	"python3-openssl",
	"python3-passlib",
	"python3-pil",
	"python3-pkg-resources",
	"python3-pyasn1",
	"python3-pycurl",
	"python3-pynvim",
	"python3-pyparsing",
	"python3-pyperclip",
	"python3-pyqt5.qtopengl",
	"python3-pyqt5",
	"python3-pyqtgraph",
	"python3-pysimplesoap",
	"python3-reportbug",
	"python3-requests",
	"python3-ruamel.yaml",
	"python3-scipy",
	"python3-sip",
	"python3-six",
	"python3-sortedcontainers",
	"python3-tk",
	"python3-tornado",
	"python3-unidiff",
	"python3-urllib3",
	"python3-urwid",
	"python3-wsproto",
	"python3-xdg",
	"python3.7-minimal",
	"python3.7",
	"python3.8-minimal",
	"python3.8",
	"python3",
	"python",
	"qemu-block-extra",
	"qemu-efi-aarch64",
	"qemu-efi-arm",
	"qemu-efi",
	"qemu-kvm",
	"qemu-system-arm",
	"qemu-system-common",
	"qemu-system-data",
	"qemu-system-mips",
	"qemu-system-misc",
	"qemu-system-ppc",
	"qemu-system-sparc",
	"qemu-system-x86",
	"qemu-system",
	"qemu-user-static",
	"qemu-user",
	"qemu-utils",
	"qemu",
	"rake",
	"readline-common",
	"rkflashtool",
	"rkhunter",
	"rsync",
	"rsyslog",
	"ruby-minitest",
	"ruby-net-telnet",
	"ruby-power-assert",
	"ruby-test-unit",
	"ruby-xmlrpc",
	"ruby2.7",
	"ruby",
	"rubygems-integration",
	"sbsigntool",
	"seabios",
	"sed",
	"sensible-utils",
	"shared-mime-info",
	"sleuthkit",
	"sound-theme-freedesktop",
	"squashfs-tools",
	"strace",
	"sudo",
	"system-config-printer-udev",
	"systemd-container",
	"systemd-sysv",
	"systemd-timesyncd",
	"systemd",
	"sysuser-helper",
	"sysvinit-utils",
	"t1utils",
	"tar",
	"task-english",
	"tasksel-data",
	"tasksel",
	"tcl8.6",
	"thin-provisioning-tools",
	"tk8.6-blt2.5",
	"tree",
	"ttf-bitstream-vera",
	"tzdata",
	"ucf",
	"udev",
	"unhide.rb",
	"unhide",
	"unzip",
	"update-inetd",
	"usb-modeswitch-data",
	"usb-modeswitch",
	"usb.ids",
	"usbauth-notifier",
	"usbauth",
	"usbip",
	"usbmuxd",
	"usbtop",
	"usbutils",
	"util-linux-locales",
	"util-linux",
	"va-driver-all",
	"vdpau-driver-all",
	"vim-common",
	"vim-runtime",
	"vim-tiny",
	"vim",
	"wdiff",
	"wget",
	"whiptail",
	"x11-common",
	"xauth",
	"xclip",
	"xdg-user-dirs",
	"xkb-data",
	"xxd",
	"xz-utils",
	"zip",
	"zlib1g"}
}
