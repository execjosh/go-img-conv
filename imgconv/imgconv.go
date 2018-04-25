// Package imgconv implements image conversion convenience methods.
package imgconv

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ImgConv main type
type ImgConv struct {
	rootDir string
}

// New creates an ImgConv
func New(rootDir string) *ImgConv {
	return &ImgConv{
		rootDir: rootDir,
	}
}

// Convert traverses rootDir and attempts to convert images from `from` to `to`
func (c *ImgConv) Convert(from, to string) error {
	if from == to {
		return errors.New("from and to are the same")
	}

	workFactory := WorkFactories[from+"-"+to]
	if workFactory == nil {
		return errors.New("Conversion not yet implemented")
	}

	return filepath.Walk(c.rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, path, ":", err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if err := workFactory(path).Run(); err != nil {
			fmt.Fprintln(os.Stderr, path, ":", err)
		}

		return nil
	})
}
