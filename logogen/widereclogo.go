package logogen

import "math"

// WideRectangleLogo is a wide rectangle logo 🌚
type WideRectangleLogo struct {
	image         []byte
	height, width int
}

// GetLogoPosition returns x, y coordinates on a 2D plane that will make the logo centered
// according to its parent image(background)
func (wl *WideRectangleLogo) GetLogoPosition(parentWidth, parentHeight int) (x, y float64) {
	x = getCenterStartOfElement(float64(wl.width), float64(parentWidth))
	y = float64((parentWidth / 2) - (wl.height / 2))
	return
}

// GetHeight returns logo's height
func (wl *WideRectangleLogo) GetHeight() int {
	return wl.height
}

// GetWidth returns logo's width
func (wl *WideRectangleLogo) GetWidth() int {
	return wl.width
}

// GetImage returns image encoded in base64
func (wl *WideRectangleLogo) GetImage() []byte {
	return wl.image
}

func (wl *WideRectangleLogo) GetHeightAndWidthForLogogen() (h, w int) {
	shared := int(math.Max(float64(wl.height), float64(wl.width)))
	return shared, shared
}
