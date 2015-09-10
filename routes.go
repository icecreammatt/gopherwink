package main

import (
	"github.com/icecreammatt/gopherwink/controllers"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	lc := controllers.NewLightController()
	led := controllers.NewLEDController()

	// Manage Zigbee Lights
	router.POST("/lights", lc.AddLight)
	router.GET("/lights", lc.LightsList)
	router.PUT("/lights/:id/power", lc.LightPower)

	router.POST("/lights/:id/timer", lc.Timer)
	router.DELETE("/lights/:id/timer", lc.CancelTimer)

	router.PUT("/lights/:id/name", lc.SetName)
	router.PUT("/lights/:id/brightness", lc.LightBrightness)
	router.DELETE("/lights/:id", lc.RemoveLight)

	// Modify LED on Wink Hub
	router.PUT("/led", led.HandleLED)

	return router
}
