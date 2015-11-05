package models

type (
	MotionSensor struct {
		Device
		Sensor
		Temperature float64 `json: "temperature"`
		TempUnit    string  `json: "tempUnit"`
		Armed       bool    `json: "armed"`
	}
)
