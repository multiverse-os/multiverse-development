package machine

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"strings"
	"time"
)

// Node information.
type Node struct {
	Hostname   string `json:"hostname,omitempty"`
	MachineID  string `json:"machineid,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
}

func (self *Node) ParseHostname() {
	self.Hostname = slurpFile("/proc/sys/kernel/hostname")
}

func (self *Node) ParseMachineID() {
	const pathSystemdMachineID = "/etc/machine-id"
	const pathDbusMachineID = "/var/lib/dbus/machine-id"

	systemdMachineID := slurpFile(pathSystemdMachineID)
	dbusMachineID := slurpFile(pathDbusMachineID)

	if len(systemdMachineID) != 0 && len(dbusMachineID) != 0 {
		// All OK, just return the machine id.
		if systemdMachineID == dbusMachineID {
			self.MachineID = systemdMachineID
			return
		}

		// They both exist, but they don't match! Copy systemd machine id to DBUS machine id.
		spewFile(pathDbusMachineID, systemdMachineID, 0444)
		self.MachineID = systemdMachineID
		return
	}

	// Copy DBUS machine id to non-existent systemd machine id.
	if len(systemdMachineID) == 0 && len(dbusMachineID) != 0 {
		spewFile(pathSystemdMachineID, dbusMachineID, 0444)
		self.MachineID = dbusMachineID
		return
	}

	// Copy systemd machine id to non-existent DBUS machine id.
	if len(systemdMachineID) != 0 && len(dbusMachineID) == 0 {
		spewFile(pathDbusMachineID, systemdMachineID, 0444)
		self.MachineID = systemdMachineID
		return
	}

	// Generate and write fresh new machine ID to both locations, conforming to the DBUS specification:
	// https://dbus.freedesktop.org/doc/dbus-specification.html#uuids

	random := make([]byte, 12)
	if _, err := rand.Read(random); err != nil {
		return
	}
	newMachineID := fmt.Sprintf("%x%x", random, time.Now().Unix())

	spewFile(pathSystemdMachineID, newMachineID, 0444)
	spewFile(pathDbusMachineID, newMachineID, 0444)
	self.MachineID = newMachineID
}

func (self *Node) ParseTimezone() {
	if fi, err := os.Lstat("/etc/localtime"); err == nil {
		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			if tzfile, err := os.Readlink("/etc/localtime"); err == nil {
				if strings.HasPrefix(tzfile, "/usr/share/zoneinfo/") {
					self.Timezone = strings.TrimPrefix(tzfile, "/usr/share/zoneinfo/")
					return
				}
			}
		}
	}

	if timezone := slurpFile("/etc/timezone"); len(timezone) != 0 {
		self.Timezone = timezone
		return
	}

	if f, err := os.Open("/etc/sysconfig/clock"); err == nil {
		defer f.Close()
		s := bufio.NewScanner(f)
		for s.Scan() {
			if sl := strings.Split(s.Text(), "="); len(sl) == 2 {
				if sl[0] == "ZONE" {
					self.Timezone = strings.Trim(sl[1], `"`)
					return
				}
			}
		}
	}
}

func (self *Node) Parse() {
	self.ParseHostname()
	self.ParseMachineID()
	self.ParseTimezone()
}
