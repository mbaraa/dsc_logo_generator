package main

import (
	"./api"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	println("Server started @ http://localhost:1105")
	startServer()
}

// startServer sets end-points and sets CORS for JS :)
// then start the server
func startServer() {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./ui/")))
	router.HandleFunc("/logo-gen/api/gen", api.GetLogo) // url?uni_name=someName&img_color=colored|gray|white&opacity=1|0&logo_type=1|2
	// cors for the fucking bitch javascript ie throwing shit like crazy :)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":1105", handler))
}
