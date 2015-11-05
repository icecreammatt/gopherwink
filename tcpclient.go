package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/icecreammatt/gopherwink/apron"
	"strings"
	"time"
)

type Message struct {
	ClientId int    `json:"clientId"`
	Message  string `json:"message"`
}

func startSocketClient(service string) {
	// read default x509 certificate
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		panic("Error loading X509 key pair")
	}

	config := tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", service, &config)
	if err != nil {
		fmt.Println("Error aborting connection", err.Error())
		return
	} else {
		go func() {
			for {
				buf := make([]byte, 512)
				_, err := conn.Read(buf)
				if err != nil {
					fmt.Println("\nError reading from socket connection")
					conn.Close()
					fmt.Println("\nConnection closed, will attempt to reconnect in 1 minute...")
					time.Sleep(1 * time.Minute)
					fmt.Println("\nAttempting to reconnect...")
					conn, err = tls.Dial("tcp", service, &config)
					if err != nil {
						fmt.Println("Reconnect failed...")
					} else {
						_, err = conn.Write([]byte("test_auth_key\n"))
						if err != nil {
							fmt.Println("Error", err.Error())
						}
					}
				}
				stringCleaned := bytes.Trim(buf, "\x00")
				var str string = fmt.Sprintf("%s", stringCleaned)
				var message Message
				err = json.Unmarshal([]byte(stringCleaned), &message)
				if err != nil {
					fmt.Println("\nIssue parsing JSON")
				} else {
					fmt.Println(str)
					fmt.Printf("%d: %s", message.ClientId, message.Message)
					fmt.Print("\n> ")

					// Strip \n for interactive client testing
					messageStrippped := strings.Replace(message.Message, "\n", "", -1)
					switch messageStrippped {
					case "test":
						conn.Write([]byte(`{"test": "success"}`))
					case "list-sensors":
						aprontest := apron.Apron{}
						devices := aprontest.ListSensors()
						fmt.Println("Devices", string(devices))
						// conn.Write([]byte(`{"devices": "` + string(devices) + `"}`))
						conn.Write(devices)
					case "list-lights":
						aprontest := apron.Apron{}
						devices := aprontest.ListLights()
						fmt.Println("Devices", string(devices))
						// conn.Write([]byte(`{"devices": "` + string(devices) + `"}`))
						conn.Write(devices)
					case "list-all":
						aprontest := apron.Apron{}
						devices := aprontest.ListAll()
						fmt.Println("Devices", string(devices))
						// conn.Write([]byte(`{"devices": "` + string(devices) + `"}`))
						conn.Write(devices)
					default:
						fmt.Printf("Unexpected command %T\n", message.Message)
					}
				}
			}
		}()

		// Login
		_, err = conn.Write([]byte("test_auth_key\n"))
		if err != nil {
			fmt.Println("Error", err.Error())
		}
	}
}
