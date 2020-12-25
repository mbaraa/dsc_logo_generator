package api

import (
	"./Logo"
	"./LogoGenerator"
	"./RGB"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ungerik/go-cairo"
	"log"
	"net/http"
	"strconv"
)

type API struct {
	router *mux.Router
}

func NewAPI() *API {
	return &API{mux.NewRouter()}
}

func (this *API) Start() {
	this.router.HandleFunc("/api/uni_name/{uni_name}/img_color/{img_color}/bg_color/{bg_color}", this.getImage).Methods("GET")
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

func (_ API) getImageColor(imageColor string) string {
	switch imageColor {
	case "raw-logo-gray":
		return "676C72"
	case "raw-logo-color":
		return "676C72"
	case "raw-logo-white":
		return "FFFFFF"
	default:
		return ""
	}
}

func (this *API) getImage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)

	parameters := mux.Vars(r) // { "uni_name": "eg ASU", "bg_color": "eg #FFFFFF", "img_color": "eg color"}

	// plz use a map for colors blyat!
	text_color := this.getImageColor(parameters["img_color"])
	img_color := parameters["img_color"]
	uni_name := parameters["uni_name"]

	rawLogo, _ := cairo.NewSurfaceFromPNG(fmt.Sprintf("res/%s.png", img_color))
	bg_color, _ := strconv.ParseFloat(parameters["bg_color"], 16)
	gen := LogoGenerator.NewLogoGenerator(rawLogo, uni_name, RGB.NewFromHex(text_color), bg_color)

	content := gen.GetLogoWithTextWithPadding(200.0, 300.0*2.0, 300.0*2.0)

	gen.Logo.Finish()
	rawLogo.Finish()

	// dear future me or anyone reading this....
	// I parsed the array into json in a separate step "instead of json.NewEncoder(w).Encode(img)"
	// to not fuck up the response with a byte array instead of a regular string :)
	img := Logo.NewLogo(content, RGB.NewRGB(gen.TextColor.GetRGB()))

	base64.StdEncoding.EncodeToString(content)

	j, _ := json.Marshal(img)
	fmt.Fprintln(w, string(j))
}
