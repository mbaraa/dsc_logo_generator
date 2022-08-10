package main

import (
	"log"
	"net/http"

	"github.com/GDSC-ASU/logo_generator/api"
)

func main() {
	println("Server started @ http://localhost:1105")
	startServer()
}

func startServer() {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./client/dist/")))
	router.HandleFunc("/api/genlogo/", api.GetLogo) // url?uni_name=someName&img_color=colored|gray|white&opacity=1|0&logo_type=1|2

	log.Fatal(http.ListenAndServe(":1105", router))
}
