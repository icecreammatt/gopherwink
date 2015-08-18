package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
)

type Light struct {
	Id     int  `json: "id"`
	Value  int  `json: "value"`
	Active bool `json: "active"`
}

type RGB struct {
	Red   int `json: "red"`
	Green int `json: "green"`
	Blue  int `json: "blue"`
}

var ServerName = "*"
var accessControlHeaders = "Origin, X-Requested-With, Content-Type, Accept"
var accessControlMethods = "GET, POST, PUT"

func HandleLightState(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", ServerName)
	w.Header().Add("Access-Control-Allow-Headers", accessControlHeaders)
	w.Header().Add("Access-Control-Allow-Methods", accessControlMethods)

	var light Light
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("state body:", string(body))
	err = json.Unmarshal(body, &light)
	if !logError(err) {
		active := "OFF"
		if light.Active {
			active = "ON"
		}

		args := []string{"-u", "-m" + strconv.Itoa(light.Id), "-t1", "-v" + active}
		cmd := exec.Command("/usr/sbin/aprontest", args...)
		err = cmd.Run()
		if !logError(err) {
			status := 200
			fmt.Fprintf(w, "Success: %d", status)
		} else {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	}
}

func HandleLightValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", ServerName)
	w.Header().Add("Access-Control-Allow-Headers", accessControlHeaders)
	w.Header().Add("Access-Control-Allow-Methods", accessControlMethods)

	var light Light
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("value body:", string(body))
	err = json.Unmarshal(body, &light)
	if !logError(err) {
		args := []string{"-u", "-m" + strconv.Itoa(light.Id), "-t2", "-v " + strconv.Itoa(light.Value)}
		cmd := exec.Command("/usr/sbin/aprontest", args...)
		err = cmd.Run()
		if !logError(err) {
			status := 200
			fmt.Fprintf(w, "Success: %d", status)
		} else {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	}
}

func HandleLED(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", ServerName)
	w.Header().Add("Access-Control-Allow-Headers", accessControlHeaders)
	w.Header().Add("Access-Control-Allow-Methods", accessControlMethods)
	var colors RGB
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &colors)
	if !logError(err) {
		args := []string{strconv.Itoa(colors.Red), strconv.Itoa(colors.Green), strconv.Itoa(colors.Blue)}
		cmd := exec.Command("/usr/sbin/set_rgb", args...)
		err = cmd.Run()
		if !logError(err) {
			status := 200
			fmt.Fprintf(w, "Success: %d", status)
		} else {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	}
}

func HandleSearchForLight(w http.ResponseWriter, r *http.Request) {
	args := []string{"-a", "-r", "zigbee"}
	runCommand(w, "/usr/sbin/aprontest", args)
}

func HandleListLights(w http.ResponseWriter, r *http.Request) {
	args := []string{"-l"}
	runCommand(w, "/usr/sbin/aprontest", args)
}

func runCommand(w http.ResponseWriter, command string, args []string) {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	} else {
		fmt.Fprintf(w, "Success: %s", out)
	}
}

func main() {
	http.HandleFunc("/lights", HandleListLights)
	http.HandleFunc("/light/search", HandleSearchForLight)
	http.HandleFunc("/light/state", HandleLightState)
	http.HandleFunc("/light/value", HandleLightValue)
	http.HandleFunc("/led", HandleLED)
	fmt.Println("Listening on port 5000")
	http.ListenAndServe(":5000", nil)
}

func logError(err error) (isError bool) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		isError = true
	} else {
		isError = false
	}
	return
}
