package vm

import (
	"fmt"
	"os"
	"strings"
)

// ImageList prints a list of downloaded base images.
func ImageList() error {
	imagesDir, err := getImagesDir()
	if err != nil {
		return err
	}

	dir, err := os.Open(imagesDir)
	if err != nil {
		return err
	}
	defer dir.Close()

	names, err := dir.Readdirnames(0)
	if err != nil {
		return err
	}

	for _, name := range names {
		if strings.HasSuffix(name, ".qcow2") {
			fmt.Println(strings.TrimSuffix(name, ".qcow2"))
		}
	}

	return nil
}
