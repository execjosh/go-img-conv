package encoder

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// An ImageEncoder has the ability to encode an image
type ImageEncoder interface {
	Encode(w io.Writer, m image.Image) error
}

// NewJpegEncoder returns a new JPEG image encoder
func NewJpegEncoder() ImageEncoder {
	return &jpegEncoder{}
}

// NewPngEncoder returns a new PNG image encoder
func NewPngEncoder() ImageEncoder {
	return &pngEncoder{}
}

type pngEncoder struct {
}

func (e *pngEncoder) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

type jpegEncoder struct {
}

func (e *jpegEncoder) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}
