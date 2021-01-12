package vm

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/subpop/go-ini"
)

const baseURL string = "http://builder.libguestfs.org/"

type template struct {
	ININame        string
	Name           string            `ini:"name"`
	OSInfo         string            `ini:"osinfo,omitempty"`
	Arch           string            `ini:"arch"`
	File           string            `ini:"file"`
	Revision       int               `ini:"revision,omitempty"`
	Checksum       map[string]string `ini:"checksum"`
	Format         string            `ini:"format"`
	Size           uint64            `ini:"size"`
	CompressedSize uint64            `ini:"compressed_size"`
	Expand         string            `ini:"expand"`
	Notes          string            `ini:"notes"`
}

type index struct {
	Templates []template `ini:"*"`
}

func newIndex() (i index, err error) {
	imagesDir, err := getImagesDir()
	if err != nil {
		return
	}

	f, err := os.Open(filepath.Join(imagesDir, "index"))
	if err != nil {
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	err = ini.UnmarshalWithOptions(data, &i, ini.Options{AllowMultilineValues: true})
	if err != nil {
		return
	}

	return
}

func (t template) String() string {
	var b strings.Builder
	w := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "NAME\t%v\n", t.Name)
	fmt.Fprintf(w, "OSINFO\t%v\n", t.OSInfo)
	fmt.Fprintf(w, "ARCH\t%v\n", t.Arch)
	fmt.Fprintf(w, "FILE\t%v\n", t.File)
	fmt.Fprintf(w, "REVISION\t%v\n", t.Revision)
	fmt.Fprintf(w, "CHECKSUM\t%v\n", t.Checksum)
	fmt.Fprintf(w, "FORMAT\t%v\n", t.Format)
	fmt.Fprintf(w, "SIZE\t%v\n", t.Size)
	fmt.Fprintf(w, "NOTES\n%v\n", t.Notes)
	w.Flush()
	return b.String()
}
