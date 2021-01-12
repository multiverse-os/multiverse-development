package cpu

type Type int

const(
	AMD64 Type = iota
	X86
	ARM
	ARM64
)

type Architecture int

const (
	Intel Architecture = iota
	AMD
	Other
)
