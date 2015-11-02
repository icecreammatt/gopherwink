package apron

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/icecreammatt/gopherwink/utils"
	"os/exec"
)

type Apron struct{}

func (apron Apron) List() []byte {
	args := []string{"-l"}
	response, err := exec.Command("/usr/sbin/aprontest", args...).Output()
	if err != nil {
		fmt.Println("Error", err.Error())
		return []byte("Error executing list command")
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
		fmt.Printf("%s", devicesJSON)
		return devicesJSON
	}
}
