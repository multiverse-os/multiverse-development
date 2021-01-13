package operatingsystem

type OperatingSystem struct {
	Name Name
	Version Version
}

type Name int

const (
	Debian Name = iota
	Alpine
)

type Version struct {
	Name string
	Version int
	Previous *Version
	Next *Version
}
