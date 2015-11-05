package apron

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/icecreammatt/gopherwink/models"
	"github.com/icecreammatt/gopherwink/utils"
	"os/exec"
)

type Apron struct{}

func (apron Apron) ListAll() []byte {
	devicesList, err := apron.List()
	if err != nil {
		fmt.Println("Error", err.Error())
		return []byte("Error executing list command")
	}
	devicesJSON, _ := json.Marshal(devicesList)
	return devicesJSON
}

func (apron Apron) ListLights() []byte {
	devicesList, err := apron.List()
	if err != nil {
		fmt.Println("Error", err.Error())
		return []byte("Error executing list command")
	}
	devicesJSON, _ := json.Marshal(devicesList.Lights)
	return devicesJSON
}

func (apron Apron) ListSensors() []byte {
	devicesList, err := apron.List()
	if err != nil {
		fmt.Println("Error", err.Error())
		return []byte("Error executing list command")
	}

	devicesJSON, _ := json.Marshal(devicesList.Switches)
	return devicesJSON
}

func (apron Apron) List() (products models.Products, err error) {
	args := []string{"-l"}
	response, err := exec.Command("/usr/sbin/aprontest", args...).Output()
	if err != nil {
		return products, err
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
		return utils.ParseDevicesFromListData(devices), nil
	}
}

func filter(devices []models.Light, f func(models.Light) bool) []models.Light {
	lights := make([]models.Light, 0)
	for _, light := range devices {
		if f(light) {
			lights = append(lights, light)
		}
	}
	return lights
}
