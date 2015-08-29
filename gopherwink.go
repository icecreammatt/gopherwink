package main

import (
	"fmt"
	"github.com/icecreammatt/gopherwink/controllers"
	"github.com/julienschmidt/httprouter"
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

	router := httprouter.New()
	lc := controllers.NewLightController()
	led := controllers.NewLEDController()

	router.GET("/lights", lc.LightsList)
	router.POST("/lights", lc.AddLight)
	router.PUT("/lights/:id/power", lc.LightPower)
	router.PUT("/lights/:id/brightness", lc.LightBrightness)
	router.PUT("/led", led.HandleLED)

	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, c.Handler(router))
}
