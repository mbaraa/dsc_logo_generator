package api

import (
	"./Logo"
	"./LogoGenerator"
	"./Resources"
	"./Text"
	"encoding/base64"
	"encoding/json"
	"github.com/rs/cors"
	"image/color"
	"log"
	"net/http"
	"strconv"
)

type API struct {
	router *http.ServeMux
}

func NewAPI() *API {
	return &API{http.NewServeMux()}
}

func (this *API) Start() {
	this.router.HandleFunc("/logo-gen/api/gen", this.getImage) // url/?uni_name=someName&img_color=colored&opacity=1or0
	// cors for the fucking bitch javascript ie throwing shit like crazy :)
	handler := cors.Default().Handler(this.router)
	log.Fatal(http.ListenAndServe(":6969", handler))
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (_ API) getTextColor(imageColor string) color.RGBA64 {
	switch imageColor {
	case "colored":
		return color.RGBA64{103, 108, 114, 0}
	case "gray":
		return color.RGBA64{103, 108, 114, 0}
	case "white":
		return color.RGBA64{255, 255, 255, 0}
	default:
		return color.RGBA64{}
	}
}

func (_ API) getRawLogo(logoColor string) []byte {
	switch logoColor {
	case "colored":
		return Resources.GetColoredLogo()
	case "gray":
		return Resources.GetGrayLogo()
	case "white":
		return Resources.GetWhiteLogo()
	default:
		return nil
	}
}

func (this *API) getImage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)

	parameters := r.URL.Query()

	imgColor := parameters["img_color"][0]
	uniName := parameters["uni_name"][0]
	opacity, _ := strconv.ParseFloat(parameters["opacity"][0], 16)

	rawLogo := this.getRawLogo(imgColor)

	generator := LogoGenerator.NewLogoGenerator(
		Logo.NewLogo(rawLogo, 1276, 3390),
		Text.NewText(
			uniName, this.getTextColor(imgColor), 0, Resources.GetProductSansFont()),
		opacity)

	newLogoBytes := generator.GetLogoWithTextWithPadding(200.0, 300.0*2.0, 300.0*2.0)

	responseJSON := make(map[string]string, 2)
	responseJSON["image"] = base64.StdEncoding.EncodeToString(newLogoBytes)
	json.NewEncoder(w).Encode(responseJSON)

}
