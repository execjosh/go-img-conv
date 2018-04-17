package main

import (
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
	fmt.Println(inputFilePath)

	f, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open file:", inputFilePath)
		os.Exit(1)
	}
	defer f.Close()

	m, fmtName, err := image.Decode(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot decode image:", inputFilePath)
		os.Exit(1)
	}
	if fmtName != "jpeg" {
		fmt.Fprintln(os.Stderr, "Expected JPEG but got:", fmtName)
		os.Exit(1)
	}

	outputFilePath := strings.TrimSuffix(inputFilePath, filepath.Ext(inputFilePath)) + ".png"
	fmt.Println(outputFilePath)
	o, err := os.OpenFile(outputFilePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open for writing:", outputFilePath)
		os.Exit(1)
	}
	defer o.Close()

	if err := png.Encode(o, m); err != nil {
		fmt.Fprintln(os.Stderr, "Error converting to PNG:", err)
		os.Exit(1)
	}
}
