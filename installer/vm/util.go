package vm

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/ulikunitz/xz"
)

func getDataDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, ".local", "share", "vm")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
	}

	return dir, nil
}

func getImagesDir() (string, error) {
	dir, err := getDataDir()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, "images")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
	}

	return dir, nil
}

func getInstancesDir() (string, error) {
	dir, err := getDataDir()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, "instances")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
	}

	return dir, nil
}

// inspect begins a state-machine of sorts that will convert the file at filePath
// to a qcow2 image. If quiet is true, no progress is printed to stdout.
func inspect(filePath string, quiet bool) (string, error) {
	switch filepath.Ext(filePath) {
	case ".gz", ".xz":
		return decompress(filePath, quiet)
	case ".raw", ".img", ".vdi":
		return convert(filePath, quiet)
	case ".tar":
		return unarchive(filePath, quiet)
	case ".box":
		newFilePath := strings.TrimSuffix(filePath, ".box") + ".tar.gz"
		if err := os.Rename(filePath, newFilePath); err != nil {
			return "", err
		}
		return inspect(newFilePath, quiet)
	case ".qcow2":
		return filePath, nil
	}

	return "", fmt.Errorf("unsupported file type: %v", filePath)
}

// transfer copies filePath into the images directory and returns the path to the
// new image. If quiet is true, no progress is printed to stdout.
func transfer(filePath string, quiet bool) (string, error) {
	var err error

	imagesDir, err := getImagesDir()
	if err != nil {
		return "", err
	}

	r, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer r.Close()

	destFilePath := filepath.Join(imagesDir, filepath.Base(filePath))
	w, err := os.Create(destFilePath + ".tmp")
	if err != nil {
		return "", err
	}
	defer w.Close()

	var bytesWritten uint64
	err = copyBytes(w, r, func(buf []byte) {
		if !quiet {
			bytesWritten += uint64(len(buf))
			fmt.Printf("\r%s", strings.Repeat(" ", 40))
			fmt.Printf("\rcopying... %s", humanize.Bytes(bytesWritten))
		}
	})
	if !quiet {
		fmt.Println()
	}
	if err != nil {
		return "", err
	}

	err = os.Rename(destFilePath+".tmp", destFilePath)
	if err != nil {
		return "", err
	}

	return destFilePath, nil
}

// download streams the body of rawurl into a file in the images directory and
// returns a path to the new image. If quiet is true, no progress is printed to
// stdout.
func download(rawurl string, quiet bool) (string, error) {
	URL, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(URL.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	imagesDir, err := getImagesDir()
	if err != nil {
		return "", err
	}

	destFilePath := filepath.Join(imagesDir, filepath.Base(URL.Path))
	w, err := os.Create(destFilePath + ".tmp")
	if err != nil {
		return "", err
	}
	defer w.Close()

	var bytesWritten uint64
	err = copyBytes(w, resp.Body, func(buf []byte) {
		if !quiet {
			bytesWritten += uint64(len(buf))
			fmt.Printf("\r%s", strings.Repeat(" ", 40))
			fmt.Printf("\rdownloading... %s", humanize.Bytes(bytesWritten))
		}
	})
	if !quiet {
		fmt.Println()
	}
	if err != nil {
		return "", err
	}

	err = os.Rename(destFilePath+".tmp", destFilePath)
	if err != nil {
		return "", err
	}

	return destFilePath, nil
}

// decompress inspects the file extension of filePath and decompresses the file
// at filePath. Only gz and xz compression formats are supported. If quiet is
// true, no progress is printed to stdout.
func decompress(filePath string, quiet bool) (string, error) {
	var err error

	r, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer r.Close()

	var g io.Reader
	switch filepath.Ext(filePath) {
	case ".gz":
		g, err = gzip.NewReader(r)
	case ".xz":
		g, err = xz.NewReader(r)
	}
	if err != nil {
		return "", err
	}

	destFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath))
	w, err := os.Create(destFilePath + ".tmp")
	if err != nil {
		return "", err
	}
	defer w.Close()

	var bytesWritten uint64
	err = copyBytes(w, g, func(buf []byte) {
		if !quiet {
			bytesWritten += uint64(len(buf))
			fmt.Printf("\r%s", strings.Repeat(" ", 40))
			fmt.Printf("\rdecompressing... %s", humanize.Bytes(bytesWritten))
		}
	})
	if !quiet {
		fmt.Println()
	}
	if err != nil {
		return "", err
	}

	if err := os.Rename(destFilePath+".tmp", destFilePath); err != nil {
		return "", err
	}

	if err := os.Remove(filePath); err != nil {
		return "", err
	}

	return inspect(destFilePath, quiet)
}

// convert calls qemu-img to convert the image at filePath to a qcow2 image. If
// quiet is true, no progress is printed to stdout.
func convert(filePath string, quiet bool) (string, error) {
	var err error

	destFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".qcow2"
	var format string
	switch filepath.Ext(filePath) {
	case ".vdi":
		format = "vdi"
	default:
		format = "raw"
	}
	cmd := exec.Command("qemu-img", "convert",
		"-f", format,
		"-O", "qcow2",
		filePath, destFilePath)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	if err := os.Remove(filePath); err != nil {
		return "", err
	}

	return inspect(destFilePath, quiet)
}

// unarchive extracts a file named "box.img" from the tar archive at filePath. If
// quiet is true, no progress is printed to stdout.
func unarchive(filePath string, quiet bool) (string, error) {
	var err error

	r, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer r.Close()

	destFilePath := strings.TrimSuffix(filePath, ".tar") + ".qcow2"

	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if hdr.Name == "box.img" {
			w, err := os.Create(destFilePath)
			if err != nil {
				return "", err
			}
			defer w.Close()

			var bytesWritten uint64
			err = copyBytes(w, tr, func(buf []byte) {
				if !quiet {
					bytesWritten += uint64(len(buf))
					fmt.Printf("\r%s", strings.Repeat(" ", 40))
					fmt.Printf("\rextracting... %s", humanize.Bytes(bytesWritten))
				}
			})
			if !quiet {
				fmt.Println()
			}
			if err != nil {
				return "", err
			}
		}
	}

	if err := os.Remove(filePath); err != nil {
		return "", err
	}

	return inspect(destFilePath, quiet)
}

// copyBytes transfers bytes from src to dest, piping through a TeeReader, calling
// writeFunc on each read from src.
func copyBytes(dest io.Writer, src io.Reader, writeFunc func(buf []byte)) error {
	w := printWriter{
		print: writeFunc,
	}

	_, err := io.Copy(dest, io.TeeReader(src, w))
	if err != nil {
		return err
	}
	return nil
}

type printWriter struct {
	print func(buf []byte)
}

func (p printWriter) Write(buf []byte) (int, error) {
	p.print(buf)
	return len(buf), nil
}
