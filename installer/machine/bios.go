package machine

// BIOS information.
type BIOS struct {
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Date    string `json:"date,omitempty"`
}

// TODO: This should be a method of BIOS 
func (self *BIOS) Parse() {
	self.Vendor = slurpFile("/sys/class/dmi/id/bios_vendor")
	self.Version = slurpFile("/sys/class/dmi/id/bios_version")
	self.Date = slurpFile("/sys/class/dmi/id/bios_date")
}
