package api

import (
	"bytes"
	"image/png"

	"github.com/ungerik/go-cairo"
)

type Logo struct {
	Image         []byte
	Height, Width float64
}

// NewLogo returns a Logo instance
func NewLogo(image []byte, height, width float64) *Logo {
	return &Logo{image, height, width}
}

// GetCairoSurface returns a cairo surface from the given attributes
func (this Logo) GetCairoSurface() *cairo.Surface {
	img0, _ := png.Decode(bytes.NewBuffer(this.Image))

	return cairo.NewSurfaceFromImage(img0)
}
