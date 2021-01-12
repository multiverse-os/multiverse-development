package vm

import (
	"os"
	"path/filepath"
	"strings"
)

// ImageGet downloads rawurl and prepares it for use as a backing disk image. If
// newName is not empty, the image is renamed to newName. If quiet is true, no
// progress is printed to stdout.
func ImageGet(rawurl string, newName string, quiet bool) error {
	var err error

	var filePath string
	if strings.HasPrefix(rawurl, "http") {
		filePath, err = download(rawurl, quiet)
		if err != nil {
			return err
		}
	} else {
		filePath, err = transfer(rawurl, quiet)
		if err != nil {
			return err
		}
	}

	finalFilePath, err := inspect(filePath, quiet)
	if err != nil {
		return err
	}

	if newName != "" {
		newpath := filepath.Join(filepath.Dir(finalFilePath), newName) + ".qcow2"
		if err := os.Rename(finalFilePath, newpath); err != nil {
			return err
		}
	}

	return nil
}
