package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Response struct {
		Status int    `json: "status"`
		Result string `json: "result"`
		Error  string `json: "active"`
	}
)

func (r *Response) Respond(w http.ResponseWriter) {
	json, _ := json.Marshal(r)
	fmt.Fprintf(w, "%s", json)
}

func (r *Response) Log() {
	fmt.Println(r)
}
