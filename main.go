package main

import (
	"fmt"
	"os"

	"github.com/execjosh/go-img-conv/imgconv"
)

func main() {
	inputDir := os.Args[1]

	if err := imgconv.ConvertAllJpegToPngIn(inputDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
