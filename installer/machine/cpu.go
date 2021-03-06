package machine

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// CPU information.
type CPU struct {
	Vendor  string `json:"vendor,omitempty"`
	Model   string `json:"model,omitempty"`
	Speed   uint   `json:"speed,omitempty"`   // CPU clock rate in MHz
	Cache   uint   `json:"cache,omitempty"`   // CPU cache size in KB
	Cpus    uint   `json:"cpus,omitempty"`    // number of physical CPUs
	Cores   uint   `json:"cores,omitempty"`   // number of physical CPU cores
	Threads uint   `json:"threads,omitempty"` // number of logical (HT) CPU cores
}

var (
	reTwoColumns = regexp.MustCompile("\t+: ")
	reExtraSpace = regexp.MustCompile(" +")
	reCacheSize  = regexp.MustCompile(`^(\d+) KB$`)
)

func (self *Machine) ParseCPU() {
	self.CPU.Threads = uint(runtime.NumCPU())

	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return
	}
	defer f.Close()

	cpu := make(map[string]bool)
	core := make(map[string]bool)

	var cpuID string

	s := bufio.NewScanner(f)
	for s.Scan() {
		if sl := reTwoColumns.Split(s.Text(), 2); sl != nil {
			switch sl[0] {
			case "physical id":
				cpuID = sl[1]
				cpu[cpuID] = true
			case "core id":
				coreID := fmt.Sprintf("%s/%s", cpuID, sl[1])
				core[coreID] = true
			case "vendor_id":
				if len(self.CPU.Vendor) == 0 {
					self.CPU.Vendor = sl[1]
				}
			case "model name":
				if len(self.CPU.Model) == 0 {
					// CPU model, as reported by /proc/cpuinfo, can be a bit ugly. Clean up...
					model := reExtraSpace.ReplaceAllLiteralString(sl[1], " ")
					self.CPU.Model = strings.Replace(model, "- ", "-", 1)
				}
			case "cache size":
				if self.CPU.Cache == 0 {
					if m := reCacheSize.FindStringSubmatch(sl[1]); m != nil {
						if cache, err := strconv.ParseUint(m[1], 10, 64); err == nil {
							self.CPU.Cache = uint(cache)
						}
					}
				}
			}
		}
	}
	if s.Err() != nil {
		return
	}

	// getNodeInfo() must have run first, to detect if we're dealing with a virtualized CPU! Detecting number of
	// physical processors and/or cores is totally unreliable in virtualized environments, so let's not do it.
	if len(self.Node.Hostname) == 0 || len(self.Node.Hypervisor) != 0 {
		return
	}

	self.CPU.Cpus = uint(len(cpu))
	self.CPU.Cores = uint(len(core))
}
