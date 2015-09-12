package main

import (
	"fmt"
	"github.com/icecreammatt/gopherwink/scheduler"
	"github.com/rs/cors"
	"net/http"
)

var port string = "5001"

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		OptionsPassthrough: false,
	})
	router := NewRouter()
	fmt.Println("Listening on port", port)
	scheduler.Start(60 * 1000)
	http.ListenAndServe(":"+port, c.Handler(router))
}
