package LogoGenerator

import (
	"../RGB"
	"github.com/ungerik/go-cairo"
	"math"
	"unicode"
)

type LogoGenerator struct {
	Logo             *cairo.Surface
	finalImage       *cairo.Surface
	height           float64
	width            float64
	Text             string
	actualTextLength float64
	TextColor        *RGB.RGB
}

func NewLogoGenerator(logo *cairo.Surface, text string, textColor *RGB.RGB) *LogoGenerator {
	return &LogoGenerator{
		logo, nil,
		0.0, 0.0, text, 0.0, textColor}
}

func (this *LogoGenerator) generateActualTextLength(textSize float64) {
	// IDK, I noticed that each char is taking size/2, given the fact that each char has different size :) sooooo
	// well the above is pure horseshit :(, well not all of it
	// here's the thing blyat
	// charsSizes = {lower: size/2, upper: size/1.5, digit: size/2, space: 0}
	finalSize := 0.0
	for _, chr := range this.Text {
		if unicode.IsLower(chr) || unicode.IsDigit(chr) {
			finalSize += textSize / 2
		} else if unicode.IsUpper(chr) {
			finalSize += textSize / 1.5
		} // if you expected a third else, I feel sorry for you :)
	}
	// you can feel sorry for my too since this shitty calculation
	// will be fucked up if the given string is all lower cases
	// so whoever you are reading this stick to the rules
	// OR if you know a rational way to fix it tell me, or fork it and fix it on your own :)
	this.actualTextLength = finalSize
}

func (this *LogoGenerator) initDimensions(textSize float64) {
	this.generateActualTextLength(textSize)
	// pretty self explanatory huh ?!
	if int(this.actualTextLength) < this.Logo.GetWidth() {
		shared := math.Max(float64(this.Logo.GetWidth()), float64(this.Logo.GetHeight()))
		this.width = shared + textSize
		this.height = shared + textSize
	} else {
		this.width = this.actualTextLength
		this.height = this.actualTextLength
	}
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
func (this *LogoGenerator) GetLogoWithText(textSize float64) *cairo.Surface {
	this.initDimensions(textSize)
	logoX := this.getCenterStartOfElement(float64(this.Logo.GetWidth()), this.width)
	logoY := this.height / 2
	// create new empty transparent image
	this.finalImage = cairo.NewSurface(cairo.FORMAT_ARGB32, int(this.width), int(this.height))
	// append given logo to the top center of the created image
	this.finalImage.SetSourceSurface(this.Logo,
		logoX,
		logoY/2) // tp appear a bit above the text
	this.finalImage.Paint()
	// set font attributes
	this.finalImage.SelectFontFace("Product Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	this.finalImage.SetSourceRGB(this.TextColor.GetRGB())
	this.finalImage.SetFontSize(textSize)
	// pre-finally write the given text
	this.finalImage.MoveTo(
		this.getCenterStartOfElement(
			this.actualTextLength, this.width), // in case of text length < logo width
		(logoY/1.75)+float64(this.Logo.GetHeight()))

	this.finalImage.ShowText(this.Text)

	return this.finalImage
}
