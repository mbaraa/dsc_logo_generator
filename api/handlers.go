package api

import (
	"encoding/base64"
	"encoding/json"
	"image/color"
	"net/http"
	"strconv"
)

// setupResponse sets required response headers.
func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// getTextColor returns a color.RGBA64 that represents the text color based on
// the required logo color-style.
// if color-style is not recognised it returns a colored logo text :)
func getTextColor(logoColor string) color.RGBA64 {
	switch logoColor {
	case "v-colored", "v-gray", "h-colored", "h-gray":
		return color.RGBA64{R: 103, G: 108, B: 114}
	case "v-white", "h-white":
		return color.RGBA64{R: 255, G: 255, B: 255}
	default:
		return color.RGBA64{R: 103, G: 108, B: 114}
	}
}

// getRawLogo returns a byte array of the required logo color-style.
// if color-style is not recognised it returns a colored logo :)
func getRawLogo(logoColor string) *Logo {
	switch logoColor {
	case "v-color":
		return GetColoredLogo()
	case "v-gray":
		return GetGrayLogo()
	case "v-white":
		return GetWhiteLogo()
	case "h-color":
		return GetColoredHorizontalLogo()
	case "h-gray":
		return GetGrayHorizontalLogo()
	case "h-white":
		return GetWhiteHorizontalLogo()
	default:
		return GetColoredLogo()
	}
}

// getImageBackground returns a color color.RGBA64 that represents the logo background
// color based on the logo color-type, if color-style is not recognised
// it returns a white background
func getImageBackground(logoColor string, bgTransparency float64) color.RGBA64 {
	switch logoColor {
	case "v-color", "v-gray", "h-color", "h-gray":
		return color.RGBA64{R: 255, G: 255, B: 255, A: uint16(bgTransparency)}
	case "v-white", "h-white":
		return color.RGBA64{R: 45, G: 45, B: 45, A: uint16(bgTransparency)}
	default:
		return color.RGBA64{R: 255, G: 255, B: 255, A: uint16(bgTransparency)}
	}
}

// GetLogo generates a dsc logo based on the given request body it uses
// GetLogoWithTextWithPadding from the LogoGenerator package to append university
// name with padding. it works on the very basic 4 steps: 1. get logo properties
// from the request body, 2. get a proper raw-logo and font color based on image
// color type 3. pass data to LogoGenerator, and generate a logo :) 4. send the
// generated b64 image to the response. no error handling is available yet!,
// since this api is ONLY called from the provided front end.
func GetLogo(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)

	parameters := r.URL.Query()
	imgColor := parameters["img_color"][0]
	uniName := parameters["uni_name"][0]
	opacity, _ := strconv.ParseFloat(parameters["opacity"][0], 16)
	logoType, _ := strconv.ParseInt(parameters["logo_type"][0], 10, 16)

	// logo type management
	var xPadding, yPadding float64

	switch logoType {
	case 1:
		xPadding, yPadding = 300.0*2, 300.0*2
		imgColor = "v-" + imgColor
		break
	case 2:
		xPadding, yPadding = 75.0*2, .0
		imgColor = "h-" + imgColor
		break
	}

	rawLogo := getRawLogo(imgColor)

	generator := NewLogoGenerator(
		rawLogo,
		NewText(
			uniName, getTextColor(imgColor), 0, GetProductSansFont()),
		getImageBackground(imgColor, opacity))

	newLogoBytes := generator.GetLogoWithTextWithPadding(200.0, xPadding, yPadding, int(logoType))

	responseJSON := make(map[string]string)
	responseJSON["image"] = base64.StdEncoding.EncodeToString(newLogoBytes)
	_ = json.NewEncoder(w).Encode(responseJSON)
}
