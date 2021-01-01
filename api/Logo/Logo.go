package Logo

import (
	"bytes"
	"github.com/ungerik/go-cairo"
	"image/png"
)

type Logo struct {
	Image         []byte `json:"image"`
	Height, Width float64
}

func NewLogo(image []byte, height, width float64) *Logo {
	return &Logo{image, height, width}
}

func (this Logo) GetCairoSurface() *cairo.Surface {
	img0, _ := png.Decode(bytes.NewBuffer(this.Image))

	return cairo.NewSurfaceFromImage(img0)
}
