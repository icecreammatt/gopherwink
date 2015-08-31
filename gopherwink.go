package main

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
)

var port string = "5000"

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		OptionsPassthrough: false,
	})
	router := NewRouter()
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, c.Handler(router))
}
