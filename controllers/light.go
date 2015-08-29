package controllers

import (
	"bufio"
	"bytes"
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
	LightController struct{}
)

func NewLightController() *LightController {
	return &LightController{}
}

func (lc LightController) LightPower(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lightId, err := strconv.Atoi(p.ByName("id"))
	if !utils.LogError(err) {

		var light models.Light
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println("state body:", string(body))
		err = json.Unmarshal(body, &light)
		if !utils.LogError(err) {
			active := "OFF"
			if light.Active {
				active = "ON"
			}

			args := []string{"-u", "-m" + strconv.Itoa(lightId), "-t1", "-v" + active}
			cmd := exec.Command("/usr/sbin/aprontest", args...)
			err = cmd.Run()
			if !utils.LogError(err) {
				w.WriteHeader(200)
			} else {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
		}
	}
}

func (lc LightController) LightBrightness(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lightId, err := strconv.Atoi(p.ByName("id"))
	if !utils.LogError(err) {

		var light models.Light
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println("value body:", string(body))
		err = json.Unmarshal(body, &light)
		if !utils.LogError(err) {
			args := []string{"-u", "-m" + strconv.Itoa(lightId), "-t2", "-v " + strconv.Itoa(light.Value)}
			cmd := exec.Command("/usr/sbin/aprontest", args...)
			err = cmd.Run()
			if !utils.LogError(err) {
				w.WriteHeader(200)
			} else {
				fmt.Fprintf(w, "Error: %s", err.Error())
			}
		}
	}
}

func (lc LightController) AddLight(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	args := []string{"-a", "-r", "zigbee"}
	utils.RunCommand(w, "/usr/sbin/aprontest", args)
}

func (lc LightController) LightsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	args := []string{"-l"}
	response, err := exec.Command("/usr/sbin/aprontest", args...).Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	} else {
		// Put response into a buffer which can then
		// be split by lines
		reader := bytes.NewReader([]byte(response))
		scanner := bufio.NewScanner(reader)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		// Extract all the devices
		var devices []string
		for _, line := range lines {
			// The empty line is the section divider before
			// the groups section
			if line != "" {
				devices = append(devices, line)
			} else {
				break
			}
		}
		devicesList := utils.ParseDevicesFromListData(devices)
		devicesJSON, _ := json.Marshal(devicesList)
		fmt.Fprintf(w, "%s", devicesJSON)
	}
}
