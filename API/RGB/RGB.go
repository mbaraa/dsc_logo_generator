package RGB

import "strconv"

type RGB struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

// constructors defaults to transparent bg

func NewRGB(red, green, blue float64) *RGB {
	return &RGB{red, green, blue, 1.0}
}

func NewFromHex(hexColor string) *RGB {
	r, _ := strconv.ParseInt(hexColor[0:2], 16, 16)
	g, _ := strconv.ParseInt(hexColor[2:4], 16, 16)
	b, _ := strconv.ParseInt(hexColor[4:6], 16, 16)

	return &RGB{float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0, 1.0}
}

func (this RGB) GetRGB() (float64, float64, float64) {
	return this.Red, this.Green, this.Blue
}

func (this RGB) GetRGBA() (float64, float64, float64, float64) {
	return this.Red, this.Green, this.Blue, this.Alpha
}
