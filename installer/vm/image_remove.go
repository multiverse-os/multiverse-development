package vm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ImageRemove deletes image name from image storage directory. If force is
// true, the image is deleted without prompting for confirmation.
func ImageRemove(name string, force bool) error {
	imagesDir, err := getImagesDir()
	if err != nil {
		return err
	}

	filePath := filepath.Join(imagesDir, name+".qcow2")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	if !force {
		fmt.Printf("Are you sure you want to remove %v? (y/N) ", name+".qcow2")
		var response string
		if _, err := fmt.Scan(&response); err != nil {
			return err
		}
		if strings.ToLower(strings.TrimSpace(response)) != "y" {
			return nil
		}
	}

	if err := os.Remove(filePath); err != nil {
		return err
	}

	return nil
}
