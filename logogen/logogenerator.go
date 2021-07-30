package logogen

import (
	"bytes"
	"image/color"
	"image/png"

	"github.com/ungerik/go-cairo"
)

type LogoGenerator struct {
	height, width int
	text          *Text
	logo          Logo
	bgColor       color.RGBA64
	finalImage    *cairo.Surface
}

// NewLogoGenerator returns a LogoGenerator instance
func NewLogoGenerator(logo Logo, text *Text, backgroundColor color.RGBA64) *LogoGenerator {
	return &LogoGenerator{.0, .0, text, logo, backgroundColor, nil}
}

// GenerateLogo returns a byte array logo
// with the provided text(instance attribute) with the given text size
func (lg *LogoGenerator) GenerateLogo() []byte {
	lg.initDimensions()
	lg.appendText()

	byteImg, _ := lg.finalImage.WriteToPNGStream()
	return byteImg
}

// GenerateLogoWithPadding calls GenerateLogo and adds padding to the final result
func (lg *LogoGenerator) GenerateLogoWithPadding(xPadding, yPadding int) []byte {
	lg.initDimensions()
	lg.appendText()

	// tmp pointer
	tmpImg := lg.finalImage

	lg.finalImage = cairo.NewSurface(cairo.FORMAT_ARGB32, lg.width+xPadding, lg.height+yPadding)
	lg.finalImage.SetSourceRGBA(
		lg.normalizeRGBA(lg.bgColor))
	lg.finalImage.Paint()
	lg.finalImage.SetSourceSurface(tmpImg, float64(xPadding/2), float64(yPadding/2))
	lg.finalImage.Paint()

	byteImg, _ := lg.finalImage.WriteToPNGStream()
	lg.finalImage.Finish()
	lg.finalImage.Flush()
	tmpImg.Finish()
	tmpImg.Destroy()
	tmpImg.Flush()

	return byteImg
}

// initDimensions sets the dimensions of the final image(w/o padding)
// ie the longest dimension(since we want a square)
func (lg *LogoGenerator) initDimensions() {
	lg.height, lg.width = lg.logo.GetHeightAndWidthForLogogen()
}

// appendText under the appended logo
func (lg *LogoGenerator) appendText() {
	_, logoY := lg.appendLogo()
	textXLength, modifiedTextSize := lg.text.GetXLengthUsingParent(float64(lg.width), 0.85)
	// set font attributes
	lg.finalImage.SelectFontFace("Product Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	lg.finalImage.SetSourceRGB(lg.text.GetColorRGB())
	lg.finalImage.SetFontSize(modifiedTextSize)

	// pre-finally write the given text
	// yeah I know this breaks the single shit pattern, but I'm too lazy to fix it (:
	// maybe I'll fix it later, no one knows (:
	switch lg.logo.(type) {
	case *HorizontalWideLogo:
		lg.finalImage.MoveTo((18.9*float64(lg.width))/100, logoY+float64(lg.logo.GetHeight()/2)+lg.text.GetFontSize()/2)

	default: // Vertical logos share the same shit :)
		lg.finalImage.MoveTo(
			getCenterStartOfElement(
				textXLength, float64(lg.width)), // in case of text length < logo width
			logoY+float64(lg.logo.GetHeight())+modifiedTextSize/1.11) // add the text strictly under the logo
	}

	lg.finalImage.ShowText(lg.text.GetContent())
}

// appendLogo adds logo to the middle of the final image
func (lg *LogoGenerator) appendLogo() (float64, float64) {
	// the magical GetLogoPosition method that made handling different logo shapes possible
	logoX, logoY := lg.logo.GetLogoPosition(lg.width, lg.height)

	if _, isWideRec := lg.logo.(*WideRectangleLogo); isWideRec {
		// this is only necessary when the logo is a wide rectangle
		logoY -= lg.text.GetFontSize() // to make the logo and the text appear in the middle
	}

	// create new empty transparent image
	lg.finalImage = cairo.NewSurface(cairo.FORMAT_ARGB32, int(lg.width), int(lg.height))

	lg.initBackground()

	img0, _ := png.Decode(bytes.NewReader(lg.logo.GetImage()))
	cairoLogo := cairo.NewSurfaceFromImage(img0)

	// append given logo to the top center of the created image
	lg.finalImage.SetSourceSurface(cairoLogo, logoX, logoY)
	lg.finalImage.Paint()

	cairoLogo.Finish()
	cairoLogo.Destroy()
	cairoLogo.Flush()

	return logoX, logoY
}

// initBackground sets the transparency level of the generated logo
func (lg *LogoGenerator) initBackground() {
	lg.finalImage.SetSourceRGBA(
		lg.normalizeRGBA(lg.bgColor))
	lg.finalImage.Paint()
}

// normalizeRGBA returns a valid cairo rgba color
func (*LogoGenerator) normalizeRGBA(rgba color.RGBA64) (float64, float64, float64, float64) {
	return float64(rgba.R) / 255.0,
		float64(rgba.G) / 255.0,
		float64(rgba.B) / 255.0,
		float64(rgba.A)
}
