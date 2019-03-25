# Multiverse OS: Virtual Machine (VM) Install/Live Images
#=========================================================
This folder contains the install images and the live images used by Multiverse OS for configuring Multiverse OS VMs.

Currently Multiverse is using very basic bash scripts to download images, checksums and developer signatures to simplify the process of downloading the newest installer images for all the major end-user/desktop linux operating systems. 

	*We would like to include OSX and Windows but we are not able to because of copyright restrictions.*

### Development
The bash scripts need to be moved to Go, Rust, or Ruby scripts. Using the bash scripts we learned that we will need to supply the following variables: 
	(1) operating system name
	(2) primary/official operating system URL
	(3) version of the operating system newest stable version
            (A) Reliable third party or location on the primary/official website that can tell us what the newest stable version is
            (B) **Ther is are some great resources for this, and one I found recently** [LOOK THIS UP]
	(4) download URL (hopefully it will enable us to just switch out the version to download each verison of the OS)
	(5) location/URL of both the installer image and the live image (if available)
	(6) location of the checksums for each image type (typically named SHA256CHECKSUMS
	(7) location of the developer signatures and preferably a way to verify these are actually the devleopers signatures from a reliable third party

Most of this information can be obtained beforehand for all supported VM operating systems. The main changing info will be the OS version, which can be obtained from reliable third party sources or possibly the primary/official website for the given operating system.  

## Notes
* Updated the alpine script to use a global set of variables that should leave the script generic enough to work for any OS download. This could be the basis for an eventual piece of software built in Go or C using libos. 

