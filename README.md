# vTraceRoute - A simple traceroute visualizer.

vTraceRoute is a simple tool that allows visualization of IP address's Geolocation points on an OpenStreetMap (Mapbox) map.

#Getting started
First install `npm` and `go` programming language SDK. Once you've got it setup perform the following in sequence:
1. `go get github.com/aishraj/vtracerroute`
2. `cd $GOPATH/src/github.com/aishraj/vtracerroute`
3. `npm install`
4. `gulp build`
5. `go install`
6. `$GOPATH/bin/vtracerroute <server you'd like to ping>`
7. Navigate to localhost:4000/ to view the paths that your packet took to reach the host specified in step 6 above.

#Disclaimer
This is my first react.js project and hence the quality of the JS (and JSX) code might be mediocre at best. In case you feel that things can be improved, go ahead and send a pull request.

#TODO
1. Minify the JS files
2. Make paths traceable
3. Figure out if the React DOM is having any benefit (since leaflet.js renders the map component)
