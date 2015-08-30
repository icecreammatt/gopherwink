package models

type (
	Light struct {
		Device
		Value  int  `json: "value"`
		Active bool `json: "active"`
	}
)
