package LogoGenerator

import (
	"../RGB"
	"../Resources"
	"github.com/golang/freetype/truetype"
	"github.com/ungerik/go-cairo"
	"golang.org/x/image/math/fixed"
	"math"
)

type LogoGenerator struct {
	Logo             *cairo.Surface
	finalImage       *cairo.Surface
	height           float64
	width            float64
	Text             string
	actualTextLength float64
	TextColor        *RGB.RGB
	bgAlpha          float64
}

// the text will appear in Google Sans font
// I think I'll add a font attribute in the future but...
func NewLogoGenerator(logo *cairo.Surface, text string, textColor *RGB.RGB, backgroundAlpha float64) *LogoGenerator {
	return &LogoGenerator{
		logo, nil,
		0.0, 0.0, text, 0.0, textColor, backgroundAlpha}
}

func (this *LogoGenerator) getActualTextLength(textSize float64) float64 {
	// go to previous commits to see the epic stupid calcs :')
	var finalLength fixed.Int26_6 = 0
	psansByte := Resources.GetProductSansFont()
	psansTTF, _ := truetype.Parse(psansByte)
	for _, chr := range this.Text {
		finalLength += psansTTF.HMetric(fixed.Int26_6(textSize), psansTTF.Index(chr)).AdvanceWidth
	}

	return float64(finalLength)
}

// I hate recursion, but there's no other way that I can think of :)
func (this *LogoGenerator) generateTextLength(textSize float64) (float64, float64) {
	length := this.getActualTextLength(textSize)
	if length <= float64(this.Logo.GetWidth())*0.85 {
		return length, textSize
	} else {
		return this.generateTextLength(textSize - 1)
	}
}

func (this *LogoGenerator) initDimensions(textSize float64) {
	this.actualTextLength, _ = this.generateTextLength(textSize)
	// pretty self explanatory huh ?!
	shared := math.Max(float64(this.Logo.GetWidth()), float64(this.Logo.GetHeight()))
	this.width = shared
	this.height = shared
}

// return the coordinate of the first point of the child element
func (this LogoGenerator) getCenterStartOfElement(childLength float64, parentLength float64) float64 {
	return math.Abs((parentLength - childLength) / 2.0)
	// hmm you want the math, fine
	// we have p as the parent's length and c as the child's length soooo
	// we want the coordinate(x or y) that will make the child appear in the middle
	// ie (p-(p-c))/2 the middle of the difference of the the difference between child and parent
	// ok some magical math properties will get us to p-c/2
	// (p-(p-c))/2 = ((p-p)-(p-c))/2
	// = (-(p-c))/2 = |(p-c)/2| SINCE IT'S A LENGTH BLYAT!!
}

// TODO
// adapt to logo dimensions
func (this *LogoGenerator) GetLogoWithText(textSize float64) []byte {
	this.initDimensions(textSize)
	this.appendText(textSize)

	byteImg, _ := this.finalImage.WriteToPNGStream()
	return byteImg
}

func (this *LogoGenerator) GetLogoWithTextWithPadding(textSize, paddX, paddY float64) []byte {
	this.GetLogoWithText(textSize)

	newImg := cairo.NewSurface(cairo.FORMAT_ARGB32, int(this.width+paddX), int(this.height+paddY))
	newImg.SetSourceRGBA(1, 1, 1, this.bgAlpha)
	newImg.Paint()
	newImg.SetSourceSurface(this.finalImage, paddX/2, paddY/2)
	newImg.Paint()
	this.finalImage = newImg

	byteImg, _ := this.finalImage.WriteToPNGStream()
	return byteImg
}

func (this *LogoGenerator) appendText(textSize float64) {
	_, logoY := this.appendLogo()
	_, modifiedTextSize := this.generateTextLength(textSize)
	// set font attributes
	this.finalImage.SelectFontFace("Product Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	this.finalImage.SetSourceRGB(this.TextColor.GetRGB())

	this.finalImage.SetFontSize(modifiedTextSize)
	// pre-finally write the given text
	this.finalImage.MoveTo(
		this.getCenterStartOfElement(
			this.actualTextLength, this.width), // in case of text length < logo width
		(logoY/1.69)+float64(this.Logo.GetHeight()))

	this.finalImage.ShowText(this.Text)
}

func (this *LogoGenerator) appendLogo() (float64, float64) {
	logoX := this.getCenterStartOfElement(float64(this.Logo.GetWidth()), this.width)
	logoY := this.height / 2
	// create new empty transparent image
	this.finalImage = cairo.NewSurface(cairo.FORMAT_ARGB32, int(this.width), int(this.height))

	this.initBackground()
	// append given logo to the top center of the created image
	this.finalImage.SetSourceSurface(this.Logo,
		logoX,
		logoY/2) // tp appear a bit above the text
	this.finalImage.Paint()

	return logoX, logoY
}

func (this *LogoGenerator) initBackground() {
	// the optional bg blyat
	this.finalImage.SetSourceRGBA(1, 1, 1, this.bgAlpha)
	this.finalImage.Paint()
}
