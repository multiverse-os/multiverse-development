package machine

// Product information.
type Product struct {
	Name    string `json:"name,omitempty"`
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Serial  string `json:"serial,omitempty"`
}

func (self *Product) Parse() {
	self.Name = slurpFile("/sys/class/dmi/id/product_name")
	self.Vendor = slurpFile("/sys/class/dmi/id/sys_vendor")
	self.Version = slurpFile("/sys/class/dmi/id/product_version")
	self.Serial = slurpFile("/sys/class/dmi/id/product_serial")
}
