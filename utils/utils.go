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

type Object interface {
}

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

func readProductType(lines []string, regex *regexp.Regexp) int64 {
	for _, line := range lines {
		if regex.MatchString(line) {
			result := regex.FindAllString(line, 1)
			if len(result) > 0 {
				productString := result[0]
				fmt.Println("Regex Result:", result)
				productId := productString[len(productString)-6:]
				fmt.Println("Parsed product Id", productId)
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

func readDeviceAttributes(devices []models.Light) (lightStatus []models.Light) {
	// sample := "Manufacturer ID: 0x10dc, Product Type: 0x2001 Product Number: 0xce3d"
	// r, err := regexp.Compile(`Product Number: [x0-9a-f]*`)
	// if err != nil {
	// 	fmt.Println("Issue with regex")
	// 	return
	// }
	// if r.MatchString(sample) == true {
	// 	fmt.Println("Match ")
	// 	fmt.Println(r.FindAllString(sample, 1))
	// } else {
	// 	fmt.Println("No match ")
	// }
	regexString := "Product Number: [x0-9a-f]*"
	r := regexp.MustCompile(regexString)

	for _, light := range devices {
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
			productType := readProductType(lines, r)
			switch productType {
			case GELinkBulb:
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
				break
			case GoControlSiren:
				break
			case GoControlSwitch:
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
