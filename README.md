# vTraceRoute - A simple traceroute visualizer.

vTraceRoute is a simple tool that allows visualization of IP address's Geolocation points on an OpenStreetMap (Mapbox) map.
Tools used:
1. Go programming language
2. Reactjs
3. Mapbox.js (leaflet)

#Getting started
First install `npm` and `go` programming language SDK. Once you've got it setup perform the following in sequence:

1. `go get github.com/aishraj/vtracerroute`
2. `cd $GOPATH/src/github.com/aishraj/vtracerroute`
3. `npm install`
4. `gulp build`
5. `go install`
6. `$GOPATH/bin/vtracerroute -host <server you'd like to ping>`
7. Navigate to localhost:4000/ to view the paths that your packet took to reach the host specified in step 6 above.

#Disclaimer
This is my first react.js project and hence the quality of the JS (and JSX) code might be mediocre at best. In case you feel that things can be improved, go ahead and send a pull request.

#Sample
![Example on running traceroute with facebook.com from my EC2 machine in North Virigina](http://i.imgur.com/Jx2kChl.png?1)

#TODO
1. Make paths traceable in terms of direction.
2. Figure out if the React DOM is having any benefit (since leaflet.js renders the map component)
3. Stop redrawing once you hit a certain hops / certain number of fetch calls.
