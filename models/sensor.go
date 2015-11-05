package models

type (
	Sensor struct {
		Device
		BatteryLevel int  `json: "batteryLevel"`
		OnOff        bool `json: "onOff"`
	}
)
