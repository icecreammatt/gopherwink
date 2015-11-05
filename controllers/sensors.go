package controllers

import (
	"fmt"
	"github.com/icecreammatt/gopherwink/apron"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	SensorController struct{}
)

func NewSensorController() *SensorController {
	return &SensorController{}
}

func (lc SensorController) SensorsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	aprontest := apron.Apron{}
	devicesJSON := aprontest.ListSensors()
	fmt.Fprintf(w, "%s", devicesJSON)
}
