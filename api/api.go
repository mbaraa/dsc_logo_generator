package api

import (
	"./Logo"
	"./LogoGenerator"
	"./RGB"
	"./Resources"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"github.com/ungerik/go-cairo"
	"image/png"
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

func (_ API) getTextColor(imageColor string) string {
	switch imageColor {
	case "gray":
		return "676C72"
	case "colored":
		return "676C72"
	case "white":
		return "FFFFFF"
	default:
		return ""
	}
}

func getRawLogo(logoColor string) []byte {
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

	textColor := this.getTextColor(parameters["img_color"][0])
	imgColor := parameters["img_color"][0]
	uniName := parameters["uni_name"][0]
	opacity, _ := strconv.ParseFloat(parameters["opacity"][0], 16)

	rawLogo0, _ := png.Decode(bytes.NewReader(getRawLogo(imgColor)))
	rawLogo := cairo.NewSurfaceFromImage(rawLogo0)

	generator := LogoGenerator.NewLogoGenerator(rawLogo, uniName, RGB.NewFromHex(textColor), opacity)

	content := generator.GetLogoWithTextWithPadding(200.0, 300.0*2.0, 300.0*2.0)

	generator.Logo.Finish()
	rawLogo.Finish()

	// dear future me or anyone reading this....
	// I parsed the array into json in a separate step "instead of json.NewEncoder(w).Encode(img)"
	// to not fuck up the response with a byte array instead of a regular string :)
	img := Logo.NewLogo(content, RGB.NewRGB(generator.TextColor.GetRGB()))

	base64.StdEncoding.EncodeToString(content)

	j, _ := json.Marshal(img)
	fmt.Fprintln(w, string(j))
}
