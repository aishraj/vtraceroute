package main

import "github.com/fiorix/freegeoip"
import "net"
import "log"

type customQuery struct {
	Country struct {
		ISOCode string            `maxminddb:"iso_code"`
		Names   map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	Location struct {
		Latitude  float64 `maxminddb:"latitude"`
		Longitude float64 `maxminddb:"longitude"`
		TimeZone  string  `maxminddb:"time_zone"`
	} `maxminddb:"location"`
}

func lookupIP(ip string) (float64, float64) {
	db, err := freegeoip.Open("/tmp/GeoLite2-City.mmdb.gz")
	if err != nil {
		log.Fatal(err)
	}
	var result customQuery
	err = db.Lookup(net.ParseIP(ip), &result)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return result.Location.Latitude, result.Location.Longitude
}
