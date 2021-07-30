package logogen

import "math"

// TallRectangleLogo is a tall rectangle logo 🌚
type TallRectangleLogo struct {
	image         []byte
	height, width int
}

// GetLogoPosition returns x, y coordinates on a 2D plane that will make the logo centered
// according to its parent image(background)
func (tl *TallRectangleLogo) GetLogoPosition(parentWidth, parentHeight int) (x, y float64) {
	x = float64((parentHeight / 2) - (tl.width / 2))
	y = getCenterStartOfElement(float64(tl.height), float64(parentHeight))
	return
}

// GetHeight returns logo's height
func (tl *TallRectangleLogo) GetHeight() int {
	return tl.height
}

// GetWidth returns logo's width
func (tl *TallRectangleLogo) GetWidth() int {
	return tl.width
}

// GetImage returns image encoded in base64
func (tl *TallRectangleLogo) GetImage() []byte {
	return tl.image
}

func (tl *TallRectangleLogo) GetHeightAndWidthForLogogen() (h, w int) {
	shared := int(math.Max(float64(tl.height), float64(tl.width)))
	return shared, shared
}
