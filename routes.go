package main

import (
	"github.com/icecreammatt/gopherwink/controllers"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	lc := controllers.NewLightController()
	led := controllers.NewLEDController()

	router.POST("/lights", lc.AddLight)
	router.GET("/lights", lc.LightsList)
	router.PUT("/lights/:id/power", lc.LightPower)
	router.PUT("/lights/:id/name", lc.SetName)
	router.PUT("/lights/:id/brightness", lc.LightBrightness)
	router.PUT("/led", led.HandleLED)
	router.DELETE("/lights/:id", lc.RemoveLight)

	return router
}
