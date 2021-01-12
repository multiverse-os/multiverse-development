package vm

import (
	"fmt"
	"os/exec"
)

func init() {
	for _, c := range []string{"qemu-img"} {
		if _, err := exec.LookPath(c); err != nil {
			panic(fmt.Errorf("%v", err))
		}
	}
}
