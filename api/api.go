package api

import (
	"./Logo"
	"./LogoGenerator"
	"./Resources"
	"./Text"
	"encoding/base64"
	"encoding/json"
	"image/color"
	"net/http"
	"strconv"
)

// setupResponse sets required response headers.
func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// getTextColor returns a color.RGBA64 that represents the text color based on
// the required logo color-style.
// if color-style is not recognised it returns a colored logo text :)
func getTextColor(logoColor string) color.RGBA64 {
	switch logoColor {
	case "colored":
		return color.RGBA64{103, 108, 114, 0}
	case "gray":
		return color.RGBA64{103, 108, 114, 0}
	case "white":
		return color.RGBA64{255, 255, 255, 0}
	default:
		return color.RGBA64{103, 108, 114, 0}
	}
}

// getRawLogo returns a byte array of the required logo color-style.
// if color-style is not recognised it returns a colored logo :)
func getRawLogo(logoColor string) []byte {
	switch logoColor {
	case "colored":
		return Resources.GetColoredLogo()
	case "gray":
		return Resources.GetGrayLogo()
	case "white":
		return Resources.GetWhiteLogo()
	default:
		return Resources.GetColoredLogo()
	}
}

// getImageBackground returns a color color.RGBA64 that represents the logo background
// color based on the logo color-type, if color-style is not recognised
// it returns a white background
func getImageBackground(logoColor string, bgTransparency float64) color.RGBA64 {
	switch logoColor {
	case "colored":
		return color.RGBA64{255, 255, 255, uint16(bgTransparency)}
	case "gray":
		return color.RGBA64{255, 255, 255, uint16(bgTransparency)}
	case "white":
		return color.RGBA64{45, 45, 45, uint16(bgTransparency)}
	default:
		return color.RGBA64{0, 0, 0, uint16(bgTransparency)}
	}
}

// GetImage generates a dsc logo based on the given request body it uses
// GetLogoWithTextWithPadding from the LogoGenerator package to append university
// name with padding. it works on the very basic 4 steps: 1. get logo properties
// from the request body, 2. get a proper raw-logo and font color based on image
// color type 3. pass data to LogoGenerator, and generate a logo :) 4. send the
// generated b64 image to the response. no error handling is available yet!,
// since this api is ONLY called from the provided front end.
func GetImage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)

	parameters := r.URL.Query()
	imgColor := parameters["img_color"][0]
	uniName := parameters["uni_name"][0]
	opacity, _ := strconv.ParseFloat(parameters["opacity"][0], 16)
	rawLogo := getRawLogo(imgColor)

	generator := LogoGenerator.NewLogoGenerator(
		Logo.NewLogo(rawLogo, 1276, 3390),
		Text.NewText(
			uniName, getTextColor(imgColor), 0, Resources.GetProductSansFont()),
		getImageBackground(imgColor, opacity))

	newLogoBytes := generator.GetLogoWithTextWithPadding(200.0, 300.0*2.0, 300.0*2.0)

	responseJSON := make(map[string]string)
	responseJSON["image"] = base64.StdEncoding.EncodeToString(newLogoBytes)
	_ = json.NewEncoder(w).Encode(responseJSON)
}
