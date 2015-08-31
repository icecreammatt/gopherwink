package main

import (
	"github.com/icecreammatt/gopherwink/controllers"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	lc := controllers.NewLightController()
	led := controllers.NewLEDController()

	router.GET("/lights", lc.LightsList)
	router.POST("/lights", lc.AddLight)
	router.PUT("/lights/:id/power", lc.LightPower)
	router.DELETE("/lights/:id", lc.RemoveLight)
	router.PUT("/lights/:id/brightness", lc.LightBrightness)
	router.PUT("/led", led.HandleLED)

	return router
}
