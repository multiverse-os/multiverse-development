package machine

import (
	"strconv"
)

type Chassis struct {
	Type     uint   `json:"type,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Version  string `json:"version,omitempty"`
	Serial   string `json:"serial,omitempty"`
	AssetTag string `json:"assettag,omitempty"`
}

func (self *Chassis) Parse() {
	if chtype, err := strconv.ParseUint(slurpFile("/sys/class/dmi/id/chassis_type"), 10, 64); err == nil {
		self.Type = uint(chtype)
	}
	self.Vendor = slurpFile("/sys/class/dmi/id/chassis_vendor")
	self.Version = slurpFile("/sys/class/dmi/id/chassis_version")
	self.Serial = slurpFile("/sys/class/dmi/id/chassis_serial")
	self.AssetTag = slurpFile("/sys/class/dmi/id/chassis_asset_tag")
}
