package imgconv

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/execjosh/go-img-conv/internal/decoder"
	"github.com/execjosh/go-img-conv/internal/encoder"
)

// A Work is some runnable thing
type Work interface {
	Run() error
}

// WorkFactories holds implemented conversions
var WorkFactories = map[string]func(string) Work{
	"jpeg-png": newJpegToPngConversion,
	"png-jpeg": newPngToJpegConversion,
}

type conversionWork struct {
	decoder decoder.ImageDecoder
	encoder encoder.ImageEncoder

	inpath  string
	outpath string

	m image.Image
}

func newConversionWork(inpath string, decoder decoder.ImageDecoder, encoder encoder.ImageEncoder, ext string) *conversionWork {
	return &conversionWork{
		decoder: decoder,
		encoder: encoder,
		inpath:  inpath,
		outpath: getOutputPath(inpath, ext),
	}
}

func getOutputPath(inpath, newExt string) string {
	return strings.TrimSuffix(inpath, filepath.Ext(inpath)) + newExt
}

func newJpegToPngConversion(inpath string) Work {
	return newConversionWork(inpath, decoder.NewJpegDecoder(), encoder.NewPngEncoder(), ".png")
}

func newPngToJpegConversion(inpath string) Work {
	return newConversionWork(inpath, decoder.NewPngDecoder(), encoder.NewJpegEncoder(), ".jpg")
}

// Run executes the conversion
func (c *conversionWork) Run() error {
	if err := c.decode(); err != nil {
		return err
	}

	return c.encode()
}

func (c *conversionWork) decode() error {
	f, err := os.Open(c.inpath)
	if err != nil {
		return errors.New("Cannot open file")
	}
	defer f.Close()

	m, err := c.decoder.Decode(f)
	if err != nil {
		return err
	}

	// Save the decoded image
	c.m = m

	return nil
}

func (c *conversionWork) encode() error {
	outfile, err := os.OpenFile(c.outpath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New(fmt.Sprint("Cannot open output for writing: ", c.outpath))
	}
	defer outfile.Close()

	if err := c.encoder.Encode(outfile, c.m); err != nil {
		return errors.New(fmt.Sprint("Error converting: ", err))
	}

	return nil
}
