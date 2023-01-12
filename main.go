package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./html"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}