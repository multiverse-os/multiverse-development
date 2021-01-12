package vm

import (
	"fmt"
)

// TemplateInfo prints information about template with name and arch.
func TemplateInfo(name, arch string) error {
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

	fmt.Println(template)

	return nil
}
