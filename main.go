package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
)

type coordindate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

const locationFile = "places.json"

var locationMutex = new(sync.Mutex)

func main() {
	hostnameIpPntr := flag.String("hostIp", "www.google.com", "A hostname or an IP address that you want to traceroute to.")
	fi, err := os.Stat(locationFile)
	if err != nil {
		log.Panic("Unable to stat the file. Aborting....")
	}
	var dummy []int
	emptyData, err := json.Marshal(dummy)
	if err != nil {
		log.Panic("Unable to marshalle empty json array. Aborting....")
	}
	_, err = writeTofile(fi, emptyData)
	if err != nil {
		log.Panic("Unable to write empty json array to file. Aborting....")
	}
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/api/v1/places.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, locationFile)
	})

	log.Println("Starting traceroute....")
	ipHop := make(chan string)
	traceEnd := make(chan bool)
	go routetrace(*hostnameIpPntr, ipHop, traceEnd)
	go func() {
		for {
			select {
			case <-traceEnd:
				log.Println("Done with lookup of the ip address")
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
	if ((location.X - 0.0) < 0.00001) && ((location.Y - 0.0) < 0.00001) {
		return true, nil //we didn't get valid location (yes, i'm assuming nobody lives near 0,0 or no server is there)
	}
	locationMutex.Lock()
	defer locationMutex.Unlock()
	// Stat the file, so we can find its current permissions
	fi, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}

	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return false, err
	}

	var coordindates []coordindate
	if err := json.Unmarshal(fileData, &coordindates); err != nil {
		return false, err
	}

	coordindates = append(coordindates, location)

	locationData, err := json.MarshalIndent(coordindates, "", "    ")
	if err != nil {
		return false, err
	}
	_, err = writeTofile(fi, locationData)
	if err != nil {
		return false, err
	}
	return true, nil
}

func writeTofile(fi os.FileInfo, data []byte) (bool, error) {
	err := ioutil.WriteFile(locationFile, data, fi.Mode())
	if err != nil {
		return false, err
	}
	return true, nil
}
