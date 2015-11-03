package models

type BaseDevice interface {
	Id() int
	Interconnect() string
	Username() string
}

type Device struct {
	Id           int    `json: "id"`
	Interconnect string `json: "interconnect"`
	Username     string `json: "username"`
}
