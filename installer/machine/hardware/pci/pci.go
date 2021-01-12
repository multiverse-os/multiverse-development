package pci

type Type int

const (
	Video Type = iota
	Network
	USB
	FPGA
)
	


type PCI struct {
	Type Type
}
