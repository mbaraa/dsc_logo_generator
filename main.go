package main

import (
	"./api"
	"log"
	"net/http"
)

func main() {
	go startAPI()
	startFrontend()
}

func startFrontend() {
	http.Handle("/", http.FileServer(http.Dir("./ui/")))
	log.Fatal(http.ListenAndServe(":1105", nil))
}

func startAPI() {
	generatorAPI := api.NewAPI()
	generatorAPI.Start()
}
