package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/icecreammatt/gopherwink/models"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

func ParseDevicesFromListData(devices []string) (lights []models.Light) {
	for i, device := range devices {
		if i < 2 {
			continue
		}
		pieces := strings.Fields(device)
		var light models.Light
		lightId, err := strconv.ParseInt(pieces[0], 10, 32)
		if err != nil {
			fmt.Println(err)
			light.Id = 0
		} else {
			light.Id = int(lightId)
		}
		light.Interconnect = pieces[2]
		light.Username = pieces[4]
		lights = append(lights, light)
	}

	lights = readDeviceAttributes(lights)
	return
}

func readDeviceAttributes(lights []models.Light) (lightStatus []models.Light) {
	for _, light := range lights {
		args := []string{"-m" + strconv.Itoa(light.Id), "-l"}
		response, err := exec.Command("/usr/sbin/aprontest", args...).Output()
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			// Put response into a buffer which can then
			// be split by lines
			reader := bytes.NewReader([]byte(response))
			scanner := bufio.NewScanner(reader)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			for i, line := range lines {
				switch i {
				case 14:
					pieces := strings.Fields(line)
					state := pieces[8]
					fmt.Println("On_Off: ", state)
					if state == "ON" {
						light.Active = true
					} else {
						light.Active = false
					}
				case 15:
					pieces := strings.Fields(line)
					level, err := strconv.ParseInt(pieces[8], 10, 32)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Level: ", level)
						light.Value = int(level)
					}
				}
			}
			lightStatus = append(lightStatus, light)
		}
	}
	return
}

func RunCommand(w http.ResponseWriter, command string, args []string) {
	out, err := exec.Command(command, args...).Output()
	var res struct {
		Status int
		Result string
		Error  string
	}
	res.Result = string(out)
	res.Status = 200
	if err != nil {
		res.Status = 500
		res.Error = err.Error()
	}
	resJson, _ := json.Marshal(res)
	fmt.Fprintf(w, "%s", resJson)
}

func LogError(err error) (isError bool) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		isError = true
	} else {
		isError = false
	}
	return
}
