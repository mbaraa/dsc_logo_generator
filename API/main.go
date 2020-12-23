package main

import (
	"./LogoGenerator"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ungerik/go-cairo"
	"io/ioutil"
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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/image/{uni_name}", getImage).Methods("GET")
	// cors for the fucking bitch javascript ie throwing shit like crazy :)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":6969", handler))

}

func testGenerator() {
	f, _ := cairo.NewSurfaceFromPNG("res/raw-logo-color.png")
	k := LogoGenerator.NewLogoGenerator(f, "Applied Science University")
	img := k.GetLogoWithText(250.0)
	img.WriteToPNG("final.png")
	img.Finish()
	f.Finish()
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// for json formatting
type image struct {
	Image string `json:"image"`
}

func getImage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)

	parameters := mux.Vars(r) // { "uni_name": "eg ASU" }
	s, _ := cairo.NewSurfaceFromPNG("res/raw-logo-color.png")
	gen := LogoGenerator.NewLogoGenerator(s, parameters["uni_name"])
	gen.Text = parameters["uni_name"]

	gen.GetLogoWithText(260.0)
	gen.Logo.WriteToPNG("new-logo.png")
	cont, _ := ioutil.ReadFile("new-logo.png")
	gen.Logo.Finish()
	s.Finish()

	img := image{base64.StdEncoding.EncodeToString(cont)}
	json.NewEncoder(w).Encode(img)
}
