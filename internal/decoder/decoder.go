package decoder

import (
	"errors"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
)

// An ImageDecoder can decode an image from a file
type ImageDecoder interface {
	Decode(f *os.File) (image.Image, error)
}

// NewJpegDecoder returns a new JPEG image decoder
func NewJpegDecoder() ImageDecoder {
	return newDecoderContext("image/jpeg", "jpeg")
}

// NewPngDecoder returns a PNG image decoder
func NewPngDecoder() ImageDecoder {
	return newDecoderContext("image/png", "png")
}

type decoderContext struct {
	desiredFormatName string
	desiredMimeType   string
}

func newDecoderContext(mimeType, formatName string) *decoderContext {
	return &decoderContext{
		desiredFormatName: formatName,
		desiredMimeType:   mimeType,
	}
}

func (c *decoderContext) Decode(f *os.File) (image.Image, error) {
	if detectContentType(f) != c.desiredMimeType {
		return nil, errors.New(fmt.Sprint("MIME type was not ", c.desiredMimeType))
	}

	m, fmtName, err := image.Decode(f)
	if err != nil {
		return nil, errors.New("Cannot decode image")
	}

	if fmtName != c.desiredFormatName {
		return nil, errors.New(fmt.Sprint("Image was not ", c.desiredFormatName))
	}

	return m, nil
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
