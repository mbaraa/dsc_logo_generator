package logogen

// HorizontalWideLogo is a wide horizontal logo 🌚
type HorizontalWideLogo struct {
	image         []byte
	height, width int
}

// GetLogoPosition returns x, y coordinates on a 2D plane that will make the logo centered
// according to its parent image(background)
func (hl *HorizontalWideLogo) GetLogoPosition(parentWidth, parentHeight int) (x, y float64) {
	x = 0
	y = float64(parentHeight) * .22
	return
}

// GetHeight returns logo's height
func (hl *HorizontalWideLogo) GetHeight() int {
	return hl.height
}

// GetWidth returns logo's width
func (hl *HorizontalWideLogo) GetWidth() int {
	return hl.width
}

// GetImage returns image encoded in base64
func (hl *HorizontalWideLogo) GetImage() []byte {
	return hl.image
}

// GetHeightAndWidthForLogogen returns h, w the height and width that will be used
// when generating a logo
func (hl *HorizontalWideLogo) GetHeightAndWidthForLogogen() (h, w int) {
	return hl.height, hl.width
}
