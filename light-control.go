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
	Id     uint `json: "id"`
	Value  uint `json: "value"`
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

func HandleLight(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", ServerName)
	w.Header().Add("Access-Control-Allow-Headers", accessControlHeaders)
	w.Header().Add("Access-Control-Allow-Methods", accessControlMethods)
	status := r.FormValue("status")
	if status != "" {
		fmt.Fprintf(w, "Handling Light Switch\n")
		statusFlagString := "-v" + status
		cmd := exec.Command("/usr/sbin/aprontest", "-u", "-m1", "-t1", statusFlagString)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Fatal" + err.Error())
			return
		}
		fmt.Println("Light status", status)
	}
	value := r.FormValue("value")
	if value != "" {
		valueFlagString := "-v" + value
		cmd := exec.Command("/usr/sbin/aprontest", "-u", "-m1", "-t2", valueFlagString)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Fatal" + err.Error())
			return
		}
		fmt.Println("Light value ", value)
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

func main() {
	http.HandleFunc("/light", HandleLight)
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
