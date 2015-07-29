package main

import (
	"github.com/aeden/traceroute"
	"log"
)

func routetrace(hostname string, nodes chan string, end chan bool) {
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

	log.Printf("trying to reach %v servers\n", hostname)
	out, err := traceroute.Traceroute(hostname, new(traceroute.TracerouteOptions), c)
	if err == nil {
		if len(out.Hops) == 0 {
			log.Println("Expected at least one hop")
		}
	} else {
		log.Printf("Traceroute failed due to an error: %v", err)
	}
}
