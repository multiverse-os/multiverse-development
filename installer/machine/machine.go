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
func Init() *Machine {
	machine := &Machine{
		Node: Node{},
		OS: OperatingSystem{},
		Kernel: Kernel{},
		Product: Product{},
		Board: Board{},
		Chassis: Chassis{},
		BIOS: BIOS{},
		CPU: CPU{},
		Memory: Memory{},
		Network: []NetworkDevice{},
		Storage: []StorageDevice{},
	}
	// DMI info
	machine.Product.Parse()
	machine.Board.Parse()
	machine.Chassis.Parse()
	machine.BIOS.Parse()
	// SMBIOS info
	machine.ParseMemory()
	// Node info
	machine.Node.Parse() // depends on BIOS info
	// Hardware info
	machine.ParseCPU() // depends on Node info
	machine.ParseStorage()
	machine.ParseNetwork()
	// Software info
	machine.OS.Parse()
	machine.Kernel.Parse()
	return machine
}

