package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/icecreammatt/gopherwink/models"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type ProductType int64

const (
	Unknown               ProductType = 0x0000
	GELinkBulb                        = 0xce3d
	GoControlSwitch                   = 0x0102
	GoControlMotionSensor             = 0x0203
	GoControlSiren                    = 0x0503
)

func ParseDevicesFromListData(deviceList []string) (products models.Products) {
	var devices []models.Device
	for i, deviceString := range deviceList {
		if i < 2 {
			continue
		}
		pieces := strings.Fields(deviceString)
		var device models.Device
		deviceId, err := strconv.ParseInt(pieces[0], 10, 32)
		if err != nil {
			fmt.Println(err)
			device.Id = 0
		} else {
			device.Id = int(deviceId)
		}
		device.Interconnect = pieces[2]
		device.Username = pieces[4]
		devices = append(devices, device)
	}

	products = readDeviceAttributes(devices)
	return
}

func readProductType(lines []string, regex *regexp.Regexp) int64 {
	for _, line := range lines {
		if regex.MatchString(line) {
			result := regex.FindAllString(line, 1)
			if len(result) > 0 {
				productString := result[0]
				productId := productString[len(productString)-6:]
				product, err := strconv.ParseInt(productId, 0, 32)
				if err != nil {
					fmt.Println("Error parsing product id", err)
					return 0
				}
				return product
			}
			return 0
		}
	}
	return 0
}

func readDeviceAttributes(devices []models.Device) (products models.Products) {
	// "Manufacturer ID: 0x10dc, Product Type: 0x2001 Product Number: 0xce3d"
	regexString := "Product Number: [x0-9a-f]*"
	r := regexp.MustCompile(regexString)

	for _, device := range devices {
		args := []string{"-m" + strconv.Itoa(device.Id), "-l"}
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
			productType := readProductType(lines, r)
			switch productType {
			case GELinkBulb:
				light := ParseLightInfo(lines)
				light.Id = device.Id
				light.Interconnect = device.Interconnect
				light.Username = device.Username
				products.Lights = append(products.Lights, light)
				break
			case GoControlSiren:
				break
			case GoControlSwitch:
				switches := ParseSwitchInfo(lines)
				switches.Id = device.Id
				switches.Interconnect = device.Interconnect
				switches.Username = device.Username
				products.Switches = append(products.Switches, switches)
				break
			case GoControlMotionSensor:
				break
			default:
				break
			}

		}
	}
	return
}

func RunSilentCommand(command string, args []string) {
	out, err := exec.Command(command, args...).Output()
	res := models.Response{Result: string(out), Status: 200}
	if err != nil {
		res.Error = err.Error()
	}
	res.Log()
}

func RunCommand(w http.ResponseWriter, command string, args []string) {
	out, err := exec.Command(command, args...).Output()
	res := models.Response{Result: string(out), Status: 200}
	if err != nil {
		res.Status = 500
		res.Error = err.Error()
	}
	res.Respond(w)
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
