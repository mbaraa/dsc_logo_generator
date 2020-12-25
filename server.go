package main

import (
	"./api"
	"log"
	"net/http"
)

// this thing is still alpha! so don't expect much out of it :)
// also the front-end is just sad :(
func main() {
	go startAPI()
	startFrontend()
	//testServer()
	//testGenerator()
}

func startFrontend() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Fatal(http.ListenAndServe(":1105", nil))
}

func startAPI() {
	generatorAPI := api.NewAPI()
	generatorAPI.Start()
}
