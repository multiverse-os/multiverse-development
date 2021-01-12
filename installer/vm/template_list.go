package vm

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

// TemplateList prints a list of available templates.
func TemplateList(sortBy string) error {
	index, err := newIndex()
	if err != nil {
		return err
	}

	sort.Slice(index.Templates, func(i, j int) bool {
		switch sortBy {
		case "arch":
			return index.Templates[i].Arch < index.Templates[j].Arch
		case "description", "desc":
			return index.Templates[i].Name < index.Templates[j].Name
		case "name":
			fallthrough
		default:
			return index.Templates[i].ININame < index.Templates[j].ININame
		}
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAME\tARCH\tDESCRIPTION\t")
	for _, template := range index.Templates {
		fmt.Fprintf(w, "%v\t%v\t%v\n", template.ININame, template.Arch, template.Name)
	}
	w.Flush()

	return nil
}
