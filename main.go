package main

import (
	"fmt"
	"os"

	"github.com/execjosh/go-img-conv/imgconv"
)

func main() {
	inputDir := os.Args[1]

	ic := imgconv.New(inputDir)

	if err := ic.JpegToPng(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
