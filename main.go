package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type coordindate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

const locationFile = "places.json"

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/api/v1/places.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, locationFile)
	})

	log.Println("Starting traceroute....")
	ipHop := make(chan string)
	traceEnd := make(chan bool)
	go routetrace(ipHop, traceEnd)
	go func() {
		for {
			select {
			case <-traceEnd:
				log.Println("Done with lookup up the ip address")
				return
			case nodeString := <-ipHop:
				log.Println("Got IP as ", nodeString)
				x, y := lookupIP(nodeString)
				location := coordindate{X: x, Y: y}
				updateJSON(location, locationFile)
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
	return coordindate{X: float64(randomIntx), Y: float64(randomInty)}
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
