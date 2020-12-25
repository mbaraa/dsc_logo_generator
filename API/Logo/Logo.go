package Logo

import "../RGB"

type Logo struct {
	Image           []byte   `json:"image"`
	BackgroundColor *RGB.RGB `json:"bg_color"`
	TextColor       *RGB.RGB `json:"text_color"`
}

func NewLogo(image []byte, textColor *RGB.RGB) *Logo {
	return &Logo{image,
		&RGB.RGB{0.0, 0.0, 0.0, 0.0}, textColor}
}
