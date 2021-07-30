package logogen

import (
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
)

// Text holds a fancy text that has size, style color and more
type Text struct {
	content    string
	color      color.RGBA64
	fontSize   float64
	font       *truetype.Font
	fontFamily []byte

	// isOriginalStateChanged indicated whether the original state of the Text has changed or not
	// state change means change in Content, fontFamily, or FontSize
	//
	// it is used to reduce XLength calculations I mean when the arguments are the same why re-calculate?
	isOriginalStateChanged bool

	// xLength is an already calculated XLength, as mentioned above
	xLength float64
}

// NewText returns a Text instance
func NewText(text string, color color.RGBA64, fontSize float64, fontFamily []byte) (*Text, error) {
	font, err := getFontFromB64(fontFamily)
	if err != nil {
		return nil, err
	}

	return &Text{text, color, fontSize, font, fontFamily, false, .0}, nil
}

// getFontFromB64 returns a ttf font from a base64 encoded byte slice that represents the font
func getFontFromB64(font []byte) (*truetype.Font, error) {
	f, err := truetype.Parse(font)
	if err != nil {
		return nil, err
	}

	// happily ever after
	return f, nil
}

// GetContent returns text's content string
func (t *Text) GetContent() string {
	return t.content
}

// SetContent sets content's value to the given value
func (t *Text) SetContent(content string) *Text {
	t.isOriginalStateChanged = true // hmm

	t.content = content
	return t
}

// GetFontSize returns text's font size
func (t *Text) GetFontSize() float64 {
	return t.fontSize
}

// SetFontSize sets 's value to the given value
func (t *Text) SetFontSize(fontSize float64) *Text {
	t.isOriginalStateChanged = true // hmm

	t.fontSize = fontSize
	return t
}

// GetFontColor returns text's color
func (t *Text) GetFontColor() color.RGBA64 {
	return t.color
}

// SetFontColor sets 's value to the given value
func (t *Text) SetFontColor(fontColor color.RGBA64) *Text {
	t.color = fontColor
	return t
}

// GetFontFamily returns font's family's byte slice
func (t *Text) GetFontFamily() []byte {
	return t.fontFamily
}

// SetFontFamily sets a new font family for the text
//
// the reason of using setter and getter methods for fontFamily, to avoid arbitrary font updating
// since updating the font requires updating the font parser too :)
func (t *Text) SetFontFamily(fontFamily []byte) *Text {
	t.isOriginalStateChanged = true // hmm

	t.fontFamily = fontFamily
	f, err := getFontFromB64(fontFamily)
	if err != nil {
		panic(err)
		return nil
	}

	// happily ever after
	t.font = f
	return t
}

// GetXLength calculates and returns text's horizontal XLength as in its font family
// if the text hasn't changed the calculation is skipped :)
func (t *Text) GetXLength() float64 {
	if !t.isOriginalStateChanged && t.xLength != .0 { // same arguments & XLength is already calculated
		return t.xLength
	}

	t.isOriginalStateChanged = false // hmm

	// store XLength's value to avoid multi calculations of the same value
	t.xLength = t.getXLength(t.fontSize)
	return t.xLength
}

// getXLength returns Text's horizontal XLength as in its font family and respecting the given fontSize
func (t *Text) getXLength(fontSize float64) float64 {
	var xLength fixed.Int26_6 = 0
	for _, v := range t.content {
		xLength +=
			t.font.HMetric(fixed.Int26_6(fontSize), t.font.Index(v)).AdvanceWidth
	}

	return float64(xLength)
}

// GetXLengthUsingParent returns text's XLength and new fontSize when placed in a container with a specific ratio
//
// the original FontSize is not affected by this method!
func (t *Text) GetXLengthUsingParent(parentWidth, child2ParentRatio float64) (float64, float64) {
	return t.getXLengthUsingParent(parentWidth, child2ParentRatio, t.fontSize)
}

// getXLengthUsingParent is the BTS of GetXLengthUsingParent to avoid giving a dozen of parameters :)
func (t *Text) getXLengthUsingParent(parentWidth, child2ParentRatio, fontSize float64) (float64, float64) {
	length := t.getXLength(fontSize)
	if length <= child2ParentRatio*parentWidth {
		return length, fontSize
	}

	return t.getXLengthUsingParent(parentWidth, child2ParentRatio, fontSize-1)
}

// GetColorRGBA returns a normalized text's color in RGBA form
func (t *Text) GetColorRGBA() (float64, float64, float64, float64) {
	return float64(t.color.R) / 255.0,
		float64(t.color.G) / 255.0,
		float64(t.color.B) / 255.0,
		float64(t.color.A) / 255.0
}

// GetColorRGB returns a normalized text's color in RGB form
func (t *Text) GetColorRGB() (float64, float64, float64) {
	return float64(t.color.R) / 255.0,
		float64(t.color.G) / 255.0,
		float64(t.color.B) / 255.0
}
