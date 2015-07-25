package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/api/v1/places.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "places.json")
	})
	log.Println("Listening on port 4000....")
	http.ListenAndServe(":4000", nil)
}
