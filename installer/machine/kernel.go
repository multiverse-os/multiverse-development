package machine

import (
	"strings"
	"syscall"
	"unsafe"
)

type Kernel struct {
	Release      string `json:"release,omitempty"`
	Version      string `json:"version,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

func (self *Kernel) Parse() {
	self.Release = slurpFile("/proc/sys/kernel/osrelease")
	self.Version = slurpFile("/proc/sys/kernel/version")

	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		return
	}

	self.Architecture = strings.TrimRight(string((*[65]byte)(unsafe.Pointer(&uname.Machine))[:]), "\000")
}
