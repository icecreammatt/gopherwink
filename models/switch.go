package models

type (
	Switch struct {
		Device
		Sensor
		Armed bool `json: "armed"`
	}
)
