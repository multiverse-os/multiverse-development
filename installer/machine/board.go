package machine

// Board information.
type Board struct {
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Version  string `json:"version,omitempty"`
	Serial   string `json:"serial,omitempty"`
	AssetTag string `json:"assettag,omitempty"`
}

func (self *Board) Parse() {
	self.Name = slurpFile("/sys/class/dmi/id/board_name")
	self.Vendor = slurpFile("/sys/class/dmi/id/board_vendor")
	self.Version = slurpFile("/sys/class/dmi/id/board_version")
	self.Serial = slurpFile("/sys/class/dmi/id/board_serial")
	self.AssetTag = slurpFile("/sys/class/dmi/id/board_asset_tag")
}
