package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/icecreammatt/gopherwink/models"
	"github.com/icecreammatt/gopherwink/utils"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
)

type (
	LEDController struct{}
)

func NewLEDController() *LEDController {
	return &LEDController{}
}

func (lc LEDController) HandleLED(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var colors models.LED
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &colors)
	if !utils.LogError(err) {
		args := []string{strconv.Itoa(colors.Red), strconv.Itoa(colors.Green), strconv.Itoa(colors.Blue)}
		cmd := exec.Command("/usr/sbin/set_rgb", args...)
		err = cmd.Run()
		if !utils.LogError(err) {
			w.WriteHeader(200)
		} else {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	}
}
