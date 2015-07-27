package main

import (
	"github.com/aeden/traceroute"
	"log"
)

func routetrace(nodes chan string, end chan bool) {
	c := make(chan traceroute.TracerouteHop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				end <- true
				return
			}
			log.Println("hop is there")
			log.Println(hop.AddressString())
			nodes <- hop.AddressString()
		}
	}()

	log.Println("trying to reach google servers")
	out, err := traceroute.Traceroute("google.com", new(traceroute.TracerouteOptions), c)
	if err == nil {
		if len(out.Hops) == 0 {
			log.Println("Expected at least one hop")
		}
	} else {
		log.Printf("Traceroute failed due to an error: %v", err)
	}
}
