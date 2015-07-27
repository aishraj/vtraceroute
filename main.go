package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type coordindate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

const locationFile = "places.json"

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/api/v1/places.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, locationFile)
	})
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Trying to update JSON in file....")
				location := generateCoordinates()
				log.Println("New location is :", location)
				_, err := updateJSON(location, locationFile)
				if err != nil {
					log.Panic("Unable to write data to the file")
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	log.Println("Listening on port 4000....")
	http.ListenAndServe(":4000", nil)

}

//generateCoordinates generates two numbers x and y between -90,+90
func generateCoordinates() coordindate {
	randomIntx := rand.Intn(91)
	randomInty := rand.Intn(91)
	xSign := rand.Intn(2)
	ySign := rand.Intn(2)
	if xSign < 1 {
		randomIntx = randomIntx * (-1)
	}
	if ySign < 1 {
		randomInty = randomInty * (-1)
	}
	return coordindate{X: randomIntx, Y: randomInty}
}

func updateJSON(location coordindate, filepath string) (bool, error) {
	// Stat the file, so we can find its current permissions
	fi, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	locationData, err := json.MarshalIndent(location, "", "    ")
	if err != nil {
		return false, err
	}
	err = ioutil.WriteFile(filepath, locationData, fi.Mode())
	if err != nil {
		return false, err
	}
	return true, nil
}
