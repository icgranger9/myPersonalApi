package main

import ("testGoProject/cmd/server")

func main() {
	// should read env variables
	port := ":3000"

	// start server
	server.CreateAndListen(port)
}