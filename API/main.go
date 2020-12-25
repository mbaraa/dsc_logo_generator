package main

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
)

// this thing is still alpha! so don't expect much out of it :)
// also the front-end is just sad :(
func main() {
	testServer()
	//testGenerator()
}

func testServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/uni_name/{uni_name}/img_color/{img_color}/bg_color/{bg_color}", getImage).Methods("GET")
	// cors for the fucking bitch javascript ie throwing shit like crazy :)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":6969", handler))

}

func testGenerator() {
	/*f, _ := cairo.NewSurfaceFromPNG("res/raw-logo-color.png")
	k := LogoGenerator.NewLogoGenerator(f, "Applied Science University")
	img := k.GetLogoWithText(200.0)
	img.WriteToPNG("final.png")
	img.Finish()
	f.Finish()
	*/
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// oi blyat a shitty function :(
func getImage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)

	parameters := mux.Vars(r) // { "uni_name": "eg ASU", "bg_color": "eg #FFFFFF", "img_color": "eg color"}
	uni_name := parameters["uni_name"]
	img_color := parameters["img_color"] // color, gray, or white
	bg_color := parameters["bg_color"]

	if bg_color == "-16" {
		// lol
	}

	// plz use a map for colors blyat!
	var text_color string
	switch img_color {
	case "raw-logo-gray":
		text_color = "676C72"
		break
	case "raw-logo-color":
		text_color = "676C72"
		break
	case "raw-logo-white":
		text_color = "FFFFFF"
		break
	}

	rawLogo, _ := cairo.NewSurfaceFromPNG(fmt.Sprintf("res/%s.png", img_color))
	gen := LogoGenerator.NewLogoGenerator(rawLogo, uni_name, RGB.NewFromHex(text_color))

	// TODO
	// replace with byte array pass around! well for speed issues
	content, _ := gen.GetLogoWithText(200.0).WriteToPNGStream()
	gen.Logo.Finish()
	rawLogo.Finish()

	// dear future me or anyone reading this....
	// I parsed the array into json in a separate step "instead of json.NewEncoder(w).Encode(img)"
	// to not fuck up the response with a byte array instead of a regular string :)
	img := Logo.NewLogo(content, RGB.NewRGB(gen.TextColor.GetRGB()))
	{
		base64.StdEncoding.EncodeToString(content)
	}
	j, _ := json.Marshal(img)
	fmt.Fprintln(w, string(j))
}
