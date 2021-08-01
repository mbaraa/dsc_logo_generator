package logogen

import "math"

// SquareLogo is a square logo 🌚
type SquareLogo struct {
	image         []byte
	height, width int
}

// GetLogoPosition returns x, y coordinates on a 2D plane that will make the logo centered
// according to its parent image(background)
func (sl *SquareLogo) GetLogoPosition(parentWidth, parentHeight int) (x, y float64) {
	x = getCenterStartOfElement(float64(sl.width), float64(parentWidth))
	y = getCenterStartOfElement(float64(sl.height), float64(parentHeight))
	return
}

// GetHeight returns logo's height
func (sl *SquareLogo) GetHeight() int {
	return sl.height
}

// GetWidth returns logo's width
func (sl *SquareLogo) GetWidth() int {
	return sl.width
}

// GetImage returns image encoded in base64
func (sl *SquareLogo) GetImage() []byte {
	return sl.image
}

func (sl *SquareLogo) GetHeightAndWidthForLogogen() (h, w int) {
	shared := int(math.Max(float64(sl.width), float64(sl.height)))
	return shared, shared
}
