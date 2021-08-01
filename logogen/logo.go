package logogen

import (
	"bytes"
	"image/png"
	"math"
)

// LogoType defines logo orientation type
type LogoType uint

// LogoType constants, which are vertical or horizontal :)
const (
	VerticalLogo LogoType = iota
	HorizontalLogo
)

// Logo interface represents a general logo despite its shape
type Logo interface {
	// GetLogoPosition returns x, y coordinates on a 2D plane that will make the logo centered
	// according to its parent image(background)
	GetLogoPosition(parentWidth, parentHeight int) (x, y float64)

	// GetHeight returns logo's height
	GetHeight() int

	// GetWidth returns logo's width
	GetWidth() int

	// GetImage returns a byte slice that has the image
	GetImage() []byte

	// GetHeightAndWidthForLogogen returns h, w the height and width that will be used
	// when generating a logo
	GetHeightAndWidthForLogogen() (h, w int)
}

// NewLogo returns a Logo instance depending on its shit
// wow much factory!
func NewLogo(image []byte, logoType LogoType) Logo {
	// wow auto dimensions detector
	height, width := getDimensionsOfByteSliceImage(image)

	switch logoType {
	case VerticalLogo:

		if width > height {
			return &WideRectangleLogo{image, height, width}
		} else if width < height {
			return &TallRectangleLogo{image, height, width}
		}
		return &SquareLogo{image, height, width}

	case HorizontalLogo:
		return &HorizontalWideLogo{image, height, width}
	}

	return nil
}

func getDimensionsOfByteSliceImage(image []byte) (height, width int) {
	imgReader := bytes.NewReader(image)
	img, err := png.Decode(imgReader)
	if err != nil {
		panic(err)
	}

	return img.Bounds().Max.Y, img.Bounds().Max.X
}

// getCenterStartOfElement return the coordinate of the first point of the child element
// that will allow it to appear in the middle of the parent
// hmm you want the math, fine!
// we have p as the parent's length and c as the child's length soooo
// we want the coordinate(x or y) that will make the child appear in the middle
// ie (p-(p-c))/2 the middle of the difference of the the difference between child and parent
// ok some magical math properties will get us to p-c/2
// (p-(p-c))/2 = ((p-p)-(p-c))/2
// = (-(p-c))/2 = |(p-c)/2| SINCE IT'S A LENGTH BLYAT!!
func getCenterStartOfElement(childLength float64, parentLength float64) float64 {
	return math.Abs((parentLength - childLength) / 2.0)
}
