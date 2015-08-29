package models

type (
	Light struct {
		Id           int    `json: "id"`
		Value        int    `json: "value"`
		Active       bool   `json: "active"`
		Interconnect string `json: "interconnect"`
		Username     string `json: "username"`
	}
)
