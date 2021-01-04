package LogoGenerator

import (
	"../Logo"
	"../Text"
	"github.com/golang/freetype/truetype"
	"github.com/ungerik/go-cairo"
	"golang.org/x/image/math/fixed"
	"image/color"
	"math"
)

type LogoGenerator struct {
	height, width float64
	text          *Text.Text
	logo          *Logo.Logo
	bgColor       color.RGBA64
	finalImage    *cairo.Surface
}

// NewLogoGenerator returns a LogoGenerator instance
func NewLogoGenerator(logo *Logo.Logo, text *Text.Text, backgroundColor color.RGBA64) *LogoGenerator {
	return &LogoGenerator{.0, .0, text, logo, backgroundColor, nil}
}

// getActualTextLength returns text length depending on text size and unicode size from the given font family
func (this *LogoGenerator) getActualTextLength(textSize float64) float64 {
	var finalLength fixed.Int26_6 = 0
	psansTTF, _ := truetype.Parse(this.text.TextFontFamily)
	for _, chr := range this.text.Text {
		finalLength += psansTTF.HMetric(fixed.Int26_6(textSize), psansTTF.Index(chr)).AdvanceWidth
	}

	return float64(finalLength)
}

// generateTextLength returns an appropriate text length ie 85% of the logo's width
// and new textSize after decreasing textLength
// I hate recursion, but there's no other way that I can think of :)
func (this *LogoGenerator) generateTextLength(textSize float64) (float64, float64) {
	length := this.getActualTextLength(textSize)
	if length <= this.logo.Width*0.85 {
		return length, textSize
	} else {
		return this.generateTextLength(textSize - 1)
	}
}

// initDimensions sets the dimensions of the final image(w/o padding)
// ie the longest dimension(since we want a square)
func (this *LogoGenerator) initDimensions() {
	shared := math.Max(this.logo.Width, this.logo.Height)
	this.width = shared
	this.height = shared
}

// getCenterStartOfElement return the coordinate of the first point of the child element
// hmm you want the math, fine!
// we have p as the parent's length and c as the child's length soooo
// we want the coordinate(x or y) that will make the child appear in the middle
// ie (p-(p-c))/2 the middle of the difference of the the difference between child and parent
// ok some magical math properties will get us to p-c/2
// (p-(p-c))/2 = ((p-p)-(p-c))/2
// = (-(p-c))/2 = |(p-c)/2| SINCE IT'S A LENGTH BLYAT!!
func (this LogoGenerator) getCenterStartOfElement(childLength float64, parentLength float64) float64 {
	return math.Abs((parentLength - childLength) / 2.0)
}

// GetLogoWithText returns a byte array logo
// with the provided text(instance attribute) with the given text size
// TODO
// adapt text size to logo dimensions ie remove textSize parameter!
func (this *LogoGenerator) GetLogoWithText(textSize float64) []byte {
	this.initDimensions()
	this.appendText(textSize)

	byteImg, _ := this.finalImage.WriteToPNGStream()
	return byteImg
}

// GetLogoWithTextWithPadding calls GetLogoWithText and adds padding to the final result
func (this *LogoGenerator) GetLogoWithTextWithPadding(textSize, paddX, paddY float64) []byte {
	this.GetLogoWithText(textSize)

	// tmp pointer
	tmpImg := this.finalImage

	this.finalImage = cairo.NewSurface(cairo.FORMAT_ARGB32, int(this.width+paddX), int(this.height+paddY))
	this.finalImage.SetSourceRGBA(
		this.getValidCairoRGBA(this.bgColor))
	this.finalImage.Paint()
	this.finalImage.SetSourceSurface(tmpImg, paddX/2, paddY/2)
	this.finalImage.Paint()

	byteImg, _ := this.finalImage.WriteToPNGStream()
	return byteImg
}

// appendText under the appended logo
func (this *LogoGenerator) appendText(textSize float64) {
	_, logoY := this.appendLogo()
	var modifiedTextSize float64
	this.text.XLength, modifiedTextSize = this.generateTextLength(textSize)
	// set font attributes
	this.finalImage.SelectFontFace("Product Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	this.finalImage.SetSourceRGB(this.text.GetColorRGB())
	this.finalImage.SetFontSize(modifiedTextSize)

	// pre-finally write the given text
	this.finalImage.MoveTo(
		this.getCenterStartOfElement(
			this.text.XLength, this.width), // in case of text length < logo width
		(logoY/1.69)+this.logo.Height)

	this.finalImage.ShowText(this.text.Text)
}

// appendLogo adds logo to the middle of the final image
func (this *LogoGenerator) appendLogo() (float64, float64) {
	logoX := this.getCenterStartOfElement(this.logo.Width, this.width)
	logoY := this.height / 2
	// create new empty transparent image
	this.finalImage = cairo.NewSurface(cairo.FORMAT_ARGB32, int(this.width), int(this.height))

	this.initBackground()
	// append given logo to the top center of the created image
	this.finalImage.SetSourceSurface(this.logo.GetCairoSurface(),
		logoX,
		logoY/2) // tp appear a bit above the text
	this.finalImage.Paint()

	return logoX, logoY
}

// initBackground sets the transparency level of the generated logo
func (this *LogoGenerator) initBackground() {
	this.finalImage.SetSourceRGBA(
		this.getValidCairoRGBA(this.bgColor))
	this.finalImage.Paint()
}

// getValidCairoRGBA returns a valid cairo rgba color
func (_ LogoGenerator) getValidCairoRGBA(rgba color.RGBA64) (float64, float64, float64, float64) {
	return float64(rgba.R) / 255.0,
		float64(rgba.G) / 255.0,
		float64(rgba.B) / 255.0,
		float64(rgba.A)
}
