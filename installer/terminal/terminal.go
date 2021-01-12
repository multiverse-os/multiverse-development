package terminal

import (
	"bytes"
)

type Terminal struct {
	Stdout bytes.Buffer
	Stderr bytes.Buffer
	

	History []string
}
