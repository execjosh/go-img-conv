package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	_ "image/jpeg"
)

func main() {
	inputFilePath := os.Args[1]

	err := convertFromJpegToPng(inputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func convertFromJpegToPng(inputFilePath string) error {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return errors.New(fmt.Sprint("Cannot open file:", inputFilePath))
	}
	defer f.Close()

	m, fmtName, err := image.Decode(f)
	if err != nil {
		return errors.New(fmt.Sprint("Cannot decode image:", inputFilePath))
	}
	if fmtName != "jpeg" {
		return errors.New(fmt.Sprint("Expected JPEG but got:", fmtName))
	}

	outputFilePath := strings.TrimSuffix(inputFilePath, filepath.Ext(inputFilePath)) + ".png"
	o, err := os.OpenFile(outputFilePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New(fmt.Sprint("Cannot open for writing:", outputFilePath))
	}
	defer o.Close()

	if err := png.Encode(o, m); err != nil {
		return errors.New(fmt.Sprint("Error converting to PNG:", err))
	}

	return nil
}
