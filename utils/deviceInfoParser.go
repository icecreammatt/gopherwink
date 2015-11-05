package utils

import (
	"fmt"
	"github.com/icecreammatt/gopherwink/models"
	"strconv"
	"strings"
)

func ParseLightInfo(lines []string) (light models.Light) {
	for i, line := range lines {
		switch i {
		case 14: // Powered
			pieces := strings.Fields(line)
			state := pieces[8]
			if state == "ON" {
				light.Active = true
			} else {
				light.Active = false
			}
		case 15: // Brightness 0-255
			pieces := strings.Fields(line)
			level, err := strconv.ParseInt(pieces[8], 10, 32)
			if err != nil {
				fmt.Println(err)
			} else {
				light.Value = int(level)
			}
		}
	}
	return
}

func ParseSwitchInfo(lines []string) (switchSensor models.Switch) {
	for i, line := range lines {
		switch i {
		case 6: // Open
			pieces := strings.Fields(line)
			state := pieces[8]
			if state == "TRUE" {
				switchSensor.OnOff = true
			} else {
				switchSensor.OnOff = false
			}
		case 7: // Battery Level
			pieces := strings.Fields(line)
			battery := pieces[8]
			batteryLvl, err := strconv.Atoi(battery)
			if err != nil {
				fmt.Println(err)
				batteryLvl = -1
			}
			switchSensor.BatteryLevel = batteryLvl
		case 8: // Armed
			pieces := strings.Fields(line)
			state := pieces[8]
			if state == "TRUE" {
				switchSensor.Armed = true
			} else {
				switchSensor.Armed = false
			}
		}
	}
	return
}
