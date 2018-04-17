// Package imgconv implements image conversion convenience methods.
package imgconv

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "image/jpeg" // Register JPEG codec
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

// JpegToPng traverses RootDir and attempts to convert JPEGs to PNGs
func (c *ImgConv) JpegToPng() error {
	return filepath.Walk(c.rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, path, ":", err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if err := convertFromJpegToPng(path); err != nil {
			fmt.Fprintln(os.Stderr, path, ":", err)
		}

		return nil
	})
}

func detectContentType(f *os.File) string {
	// Make sure we return to top of file
	defer f.Seek(0, 0)

	buf := make([]byte, 128)

	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return "unknown"
	}

	return http.DetectContentType(buf[:n])
}

func convertFromJpegToPng(inputFilePath string) error {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return errors.New(fmt.Sprint("Cannot open file"))
	}
	defer f.Close()

	if detectContentType(f) != "image/jpeg" {
		return errors.New("Not a JPEG")
	}

	m, fmtName, err := image.Decode(f)
	if err != nil {
		return errors.New(fmt.Sprint("Cannot decode image"))
	}
	if fmtName != "jpeg" {
		return errors.New("Not a JPEG")
	}

	outputFilePath := strings.TrimSuffix(inputFilePath, filepath.Ext(inputFilePath)) + ".png"
	o, err := os.OpenFile(outputFilePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New(fmt.Sprint("Cannot open output for writing: ", outputFilePath))
	}
	defer o.Close()

	if err := png.Encode(o, m); err != nil {
		return errors.New(fmt.Sprint("Error converting to PNG: ", err))
	}

	return nil
}
