package Text

import (
	"image/color"
)

type Text struct {
	Text           string
	Color          color.RGBA64
	XLength        float64
	textFontFamily []byte
}

func NewText(text string, color color.RGBA64, xLength float64, fontFamily []byte) *Text {
	return &Text{text, color, xLength, fontFamily}
}

func (this Text) GetColorRGBA() (float64, float64, float64, float64) {
	return float64(this.Color.R) / 255.0,
		float64(this.Color.G) / 255.0,
		float64(this.Color.B) / 255.0,
		float64(this.Color.A) / 255.0
}

func (this Text) GetColorRGB() (float64, float64, float64) {
	return float64(this.Color.R) / 255.0,
		float64(this.Color.G) / 255.0,
		float64(this.Color.B) / 255.0
}
