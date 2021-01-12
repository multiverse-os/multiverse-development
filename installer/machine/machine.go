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


type Machine interface {
	Install() (bool, error)
}
