package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Listening on port 4000....")
	http.ListenAndServe(":4000", nil)
}
