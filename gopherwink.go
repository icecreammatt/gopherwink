package main

import (
	"fmt"
	"github.com/icecreammatt/gopherwink/scheduler"
	"github.com/rs/cors"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
)

var port string = "5001"

func main() {

	db, err := leveldb.OpenFile("devices.db", nil)
	defer db.Close()

	data, err := db.Get([]byte("key"), nil)
	if err != nil {
		fmt.Println("error fetching data", err.Error())
	}
	fmt.Println(string(data))
	err = db.Put([]byte("key"), []byte("here is another sample"), nil)

	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		OptionsPassthrough: false,
	})
	router := NewRouter()
	fmt.Println("Listening on port", port)
	scheduler.Start(60 * 1000)
	http.ListenAndServe(":"+port, c.Handler(router))
}
