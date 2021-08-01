package api

import (
	"image/color"
	"net/http"
	"strconv"

	"github.com/mbaraa/dsc_logo_generator/files"
	"github.com/mbaraa/dsc_logo_generator/logogen"
)

// LogoProps , well it's written on the box :)
//
type LogoProps struct {
	ImgColor           string
	UniName            string
	Opacity            float64
	LogoType           logogen.LogoType
	LogoColor          string
	XPadding, YPadding int
	LogoSth            files.DecodedB64File
}

// LogoPropsJuicer gets the appropriate logo and text color based on the request query
//
type LogoPropsJuicer struct {
	props *LogoProps
}

// NewLogoPropsJuicer returns a new LogoPropsJuicer instance
//
func NewLogoPropsJuicer(req *http.Request) (lp *LogoPropsJuicer) {
	opacity, _ := strconv.ParseFloat(req.URL.Query()["opacity"][0], 32)
	logoType, _ := strconv.ParseInt(req.URL.Query()["logo_type"][0], 10, 16)

	return (&LogoPropsJuicer{
		props: &LogoProps{
			ImgColor: req.URL.Query()["img_color"][0],
			UniName:  req.URL.Query()["uni_name"][0],
			Opacity:  opacity,
			LogoType: logogen.LogoType(logoType - 1),
		},
	}).setLogoColorAndPadding().
		setLogoProps()
}

func (lp *LogoPropsJuicer) setLogoColorAndPadding() *LogoPropsJuicer {
	switch lp.props.LogoType {
	case logogen.VerticalLogo:
		lp.props.XPadding, lp.props.YPadding = 200*2, 200*2
		lp.props.LogoColor = "v-"

	case logogen.HorizontalLogo:
		lp.props.XPadding, lp.props.YPadding = 75*2, 0
		lp.props.LogoColor = "h-"
	}
	lp.props.LogoColor += lp.props.ImgColor

	return lp
}

// GetPadding returns the appropriate padding for the logo depending on its orientation
//
func (lp *LogoPropsJuicer) GetPadding() (x, y int) {
	return lp.props.XPadding, lp.props.YPadding
}

// GetTextColor returns a color.RGBA64 that represents the text color based on
// the required logo color-style.
// if color-style is not recognised it returns a colored logo text :)
//
func (lp *LogoPropsJuicer) GetTextColor() color.RGBA64 {
	switch lp.props.LogoColor {
	case "v-white", "h-white":
		return color.RGBA64{R: 255, G: 255, B: 255}
	case "v-color", "h-color":
		fallthrough
	default:
		return color.RGBA64{R: 103, G: 108, B: 114}
	}
}

// GetRawLogo returns a byte array of the required logo color-style.
// if color-style is not recognised it returns a colored logo :)
//
func (lp *LogoPropsJuicer) GetRawLogo() logogen.Logo {
	switch lp.props.LogoType {
	case logogen.VerticalLogo:
		return logogen.NewLogo(lp.props.LogoSth.Data, logogen.VerticalLogo)
	case logogen.HorizontalLogo:
		fallthrough
	default:
		return logogen.NewLogo(lp.props.LogoSth.Data, logogen.HorizontalLogo)
	}
}

func (lp *LogoPropsJuicer) setLogoProps() *LogoPropsJuicer {
	switch lp.props.LogoColor {
	case "v-white":
		lp.props.LogoSth = files.GetFilesInstance().GetVerticalLogoWhite()
	case "h-white":
		lp.props.LogoSth = files.GetFilesInstance().GetHorizontalLogoWhite()
	case "h-color":
		lp.props.LogoSth = files.GetFilesInstance().GetHorizontalLogoColored()
	case "v-color":
		fallthrough
	default:
		lp.props.LogoSth = files.GetFilesInstance().GetVerticalLogoColored()
	}

	return lp
}

// GetImageBackground returns a color color.RGBA64 that represents the logo background
// color based on the logo color-type, if color-style is not recognised
// it returns a white background
//
func (lp *LogoPropsJuicer) GetImageBackground() color.RGBA64 {
	switch lp.props.LogoColor {
	case "v-color", "v-gray", "h-color", "h-gray":
		return color.RGBA64{R: 255, G: 255, B: 255, A: uint16(lp.props.Opacity)}
	case "v-white", "h-white":
		return color.RGBA64{R: 45, G: 45, B: 45, A: uint16(lp.props.Opacity)}
	default:
		return color.RGBA64{R: 255, G: 255, B: 255, A: uint16(lp.props.Opacity)}
	}
}

// GetText returns a text for the logo with the appropriate color and size 200
//
func (lp *LogoPropsJuicer) GetText() *logogen.Text {
	text, _ := logogen.NewText(
		lp.props.UniName,
		lp.GetTextColor(),
		lp.props.LogoSth.TextSize,
		files.GetFilesInstance().GetProductSansFont(),
	)
	return text
}
