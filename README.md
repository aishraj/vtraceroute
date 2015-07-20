#Traceroute server

A Websocket server that accepts an incoming client request and visualizes it on a uniquely created map for that connection.


#How it should work for the client
1. A client sends an array of IP Addresses.
2. Client excepts a JSON value containing an error code or/and a URL.

#How it should work for the server
1. The server should listen for incoming Websocket connection.
2. Once the data on the Websocket is verfied, it returns back a Unique URL for the client.
3. In case of failure to validate, it returns a status code.
4. Now, the server starts to serve the web page at /visuals/<unique-client-id>
5. The server in batches (or timeout), looks up the ip and visualizes the data in the map.
6. As the client starts to populate more data, it asynchronously populates data into the map. (This would need react.js or backbone et al)
7. Once the end message is sent. The Websocket server closes. The client id is cleared once cache invalidation happens.
