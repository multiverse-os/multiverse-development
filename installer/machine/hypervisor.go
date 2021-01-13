package machine

import (
	"strings"
	"unsafe"

	"./cpuid"
)

//TODO: Should definitely be using enumerators, and not doing string comparoissons

// https://en.wikipedia.org/wiki/CPUID#EAX.3D0:_Get_vendor_ID
var hvmap = map[string]string{
	"bhyve bhyve ": "bhyve",
	"KVMKVMKVM":    "kvm",
	"Microsoft Hv": "hyperv",
	" lrpepyh vr":  "parallels",
	"VMwareVMware": "vmware",
	"XenVMMXenVMM": "xenhvm",
}

func isHypervisorActive() bool {
	var info [4]uint32
	cpuid.CPUID(&info, 0x1)
	return info[2]&(1<<31) != 0
}

func getHypervisorCpuid(ax uint32) string {
	var info [4]uint32
	cpuid.CPUID(&info, ax)
	return hvmap[strings.TrimRight(string((*[12]byte)(unsafe.Pointer(&info[1]))[:]), "\000")]
}

func (self *Machine) ParseHypervisor() {
	if !isHypervisorActive() {
		if hypervisorType := slurpFile("/sys/hypervisor/type"); len(hypervisorType) != 0 {
			if hypervisorType == "xen" {
				self.Node.Hypervisor = "xenpv"
			}
		}
		return
	}

	// KVM has been caught to move its real signature to this leaf, and put something completely different in the
	// standard location. So this leaf must be checked first.
	if hv := getHypervisorCpuid(0x40000100); len(hv) != 0 {
		self.Node.Hypervisor = hv
		return
	}

	if hv := getHypervisorCpuid(0x40000000); len(hv) != 0 {
		self.Node.Hypervisor = hv
		return
	}

	// getBIOSInfo() must have run first, to detect BIOS vendor
	if self.BIOS.Vendor == "Bochs" {
		self.Node.Hypervisor = "bochs"
		return
	}

	self.Node.Hypervisor = "unknown"
}
