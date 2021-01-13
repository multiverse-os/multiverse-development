package machine

type Type int

const (
	Host Type = iota
	Controller
	Service
	Router // TODO: Traditionally router was a service type
	App
)

func (self Type) String() string {
	switch self {
	case Host:
		return "host"
	case Controller:
		return "controller"
	case Service:
		return "service"
	case Router:
		return "router"
	case App:
		return "app"
	default:
		return "undefined"
	}
}


type Machine struct {
	Node    Node			`json:"node"`
	OS      OperatingSystem         `json:"os"`
	Kernel  Kernel			`json:"kernel"`
	Product Product			`json:"product"`
	Board   Board			`json:"board"`
	Chassis Chassis			`json:"chassis"`
	BIOS    BIOS			`json:"bios"`
	CPU     CPU			`json:"cpu"`
	Memory  Memory			`json:"memory"`
	Storage []StorageDevice		`json:"storage,omitempty"`
	Network []NetworkDevice		`json:"network,omitempty"`
}

// GetSysInfo gathers all available system information.
func (self *Machine) Parse() *Machine {
	// DMI info
	self.Product.Parse()
	self.Board.Parse()
	self.Chassis.Parse()
	self.BIOS.Parse()
	// SMBIOS info
	self.ParseMemory()
	// Node info
	self.Node.Parse() // depends on BIOS info
	// Hardware info
	self.ParseCPU() // depends on Node info
	self.ParseStorage()
	self.ParseNetwork()
	// Software info
	self.OS.Parse()
	self.Kernel.Parse()
	return self
}

