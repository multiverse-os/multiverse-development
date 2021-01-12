package vm

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

// TemplateGet downloads and prepares a disk template for use as a backing disk
// image. If quiet is true, no progress is printed to stdout.
func TemplateGet(name, arch string, quiet bool) error {
	index, err := newIndex()
	if err != nil {
		return err
	}

	var template template
	for _, t := range index.Templates {
		if t.ININame == name && t.Arch == arch {
			template = t
			break
		}
	}

	URL, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	URL.Path += template.File

	filePath, err := download(URL.String(), quiet)
	if err != nil {
		return err
	}

	checksum, ok := template.Checksum["sha512"]
	if !ok {
		checksum = template.Checksum[""]
	}
	if err := verify(filePath, checksum); err != nil {
		return err
	}

	destFilePath := strings.TrimSuffix(filePath, ".xz")
	destFilePath += "." + template.Format + ".xz"

	err = os.Rename(filePath, destFilePath)
	if err != nil {
		return err
	}

	if _, err := inspect(destFilePath, quiet); err != nil {
		return err
	}

	return nil
}

func verify(filePath, checksum string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	hash := sha512.New()
	_, err = hash.Write(data)
	if err != nil {
		return err
	}
	computed := fmt.Sprintf("%x", hash.Sum(nil))
	if checksum != computed {
		return fmt.Errorf("invalid checksum: %v != %v", checksum, computed)
	}

	return nil
}
