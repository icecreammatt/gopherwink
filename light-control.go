package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

type Light struct {
	Id     uint
	Value  uint
	Active bool
}

var ServerName = "http://192.168.1.11:5000"

func HandleLight(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", ServerName)
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

func main() {
	http.HandleFunc("/light", HandleLight)
	http.ListenAndServe(":5000", nil)
}
