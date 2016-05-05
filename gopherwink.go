package main

import (
	"fmt"
	"github.com/icecreammatt/gopherwink/scheduler"
	"github.com/rs/cors"
	"github.com/sasbury/mini"
	"net/http"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		OptionsPassthrough: false,
	})
	router := NewRouter()

	var latitude float64
	var longitude float64
	var connectionString string
	var authkey string
	var autoNightLights []int64
	config, err := mini.LoadConfiguration("settings.ini")
	if err != nil {
		latitude = 0
		longitude = 0
		fmt.Println("Configuration Missing")
	} else {
		latitude = config.FloatFromSection("geo-location", "latitude", 0)
		longitude = config.FloatFromSection("geo-location", "longitude", 0)
		fmt.Println(latitude)
		fmt.Println(longitude)

		autoNightLights = config.IntegersFromSection("auto-night-lights", "deviceId")
		fmt.Println(autoNightLights)

		connectionString = config.StringFromSection("remote-client", "connectionString", "")
		fmt.Println(connectionString)

		authkey = config.StringFromSection("remote-client", "authkey", "")
		fmt.Println(authkey)
	}
	scheduler.Start(60*1000, latitude, longitude, autoNightLights)

	// TODO: Check for config to see if service is provided
	// else ignore this part and do not enable the remote client
	// if authkey != "" && connectionString != "" {
	// 	fmt.Println("Starting remote socket connection")
	// 	startSocketClient(connectionString, authkey)
	// }

	appPort := config.StringFromSection("port", "app", "5000")
	fmt.Println("App Listening on port", appPort)
	go http.ListenAndServe(":"+appPort, c.Handler(router))

	staticPort := config.StringFromSection("port", "static", "80")
	fmt.Println("Static Listening on port", staticPort)
	panic(http.ListenAndServe(":"+staticPort, http.FileServer(http.Dir("/var/www/"))))
}
