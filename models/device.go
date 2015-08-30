package models

type (
	Device struct {
		Id           int    `json: "id"`
		Interconnect string `json: "interconnect"`
		Username     string `json: "username"`
	}
)
