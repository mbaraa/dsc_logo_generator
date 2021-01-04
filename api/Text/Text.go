package Text

import (
	"image/color"
)

type Text struct {
	Text           string
	Color          color.RGBA64
	XLength        float64
	TextFontFamily []byte
}

// NewText returns a Text instance
func NewText(text string, color color.RGBA64, xLength float64, fontFamily []byte) *Text {
	return &Text{text, color, xLength, fontFamily}
}

// GetColorRGBA returns text's color in RGBA form
func (this Text) GetColorRGBA() (float64, float64, float64, float64) {
	return float64(this.Color.R) / 255.0,
		float64(this.Color.G) / 255.0,
		float64(this.Color.B) / 255.0,
		float64(this.Color.A) / 255.0
}

// GetColorRGB returns text's color in RGB form
func (this Text) GetColorRGB() (float64, float64, float64) {
	return float64(this.Color.R) / 255.0,
		float64(this.Color.G) / 255.0,
		float64(this.Color.B) / 255.0
}
