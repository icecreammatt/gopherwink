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
	"time"
)

type (
	LightController struct{}
)

func NewLightController() *LightController {
	return &LightController{}
}

var timers map[int]*time.Timer

func timer(lightIndex int, args []string, countDown time.Duration) {
	fmt.Printf("starting timer for %d\n", lightIndex)
	if timers == nil {
		timers = make(map[int]*time.Timer)
	}
	timers[lightIndex] = time.NewTimer(time.Second * countDown)
	<-timers[lightIndex].C
	fmt.Printf("timer is up for %d\n", lightIndex)
	utils.RunSilentCommand("/usr/sbin/aprontest", args)
}

func (lc LightController) CancelTimer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var res models.Response
	if lightId, err := strconv.Atoi(p.ByName("id")); err == nil {
		res = models.Response{Status: 200, Result: "Canceling Timer for " + p.ByName("id")}
		fmt.Printf("Canceling timer for %d", lightId)
		go timers[lightId].Stop()
		res.Respond(w)
	} else {
		res = models.Response{Status: 400, Error: err.Error()}
		res.Respond(w)
	}
}

func (lc LightController) Timer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if lightId, err := strconv.Atoi(p.ByName("id")); err == nil {
		var light models.Light
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			res := models.Response{Status: 404, Error: err.Error()}
			res.Respond(w)
		}
		fmt.Println("state body:", string(body))
		err = json.Unmarshal(body, &light)
		if err != nil {
			res := models.Response{Status: 400, Error: err.Error()}
			res.Respond(w)
			return
		}

		var state string
		if light.Active == true {
			state = "-vON"
		} else if light.Active == false {
			state = "-vOFF"
		}

		args := []string{"-u", "-m" + strconv.Itoa(lightId), "-t1", state}
		countDown := light.CountDownInSeconds
		if countDown > 0 {
			fmt.Printf("Setting %d to %s in %d seconds\n", lightId, state, countDown)
			go timer(lightId, args, time.Duration(countDown))
			res := models.Response{Status: 200}
			res.Respond(w)
			return
		} else {
			res := models.Response{Status: 400, Error: "Time must be > 0"}
			res.Respond(w)
			return
		}
	} else {
		res := models.Response{Status: 400, Error: err.Error()}
		res.Respond(w)
	}
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

func (lc LightController) SetName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lightId, err := strconv.Atoi(p.ByName("id"))
	if !utils.LogError(err) {
		var light models.Light
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println("state body:", string(body))
		err = json.Unmarshal(body, &light)
		if !utils.LogError(err) {
			args := []string{"-m", strconv.Itoa(lightId), "--set-name", light.Username}
			utils.RunCommand(w, "/usr/sbin/aprontest", args)
		}
	}
}

func (lc LightController) AddLight(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	args := []string{"-a", "-r", "zigbee"}
	utils.RunCommand(w, "/usr/sbin/aprontest", args)
}

func (lc LightController) RemoveLight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lightId, err := strconv.Atoi(p.ByName("id"))
	if !utils.LogError(err) {
		args := []string{"-d", "-m", strconv.Itoa(lightId)}
		utils.RunCommand(w, "/usr/sbin/aprontest", args)
	}
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
