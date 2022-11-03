package main

import (
	server "example.com/triage-grpc/triage_server"
)

func main() {
	server.OnMessage() //pass in a processor function 
	server.StartServer()

	// server.onMessage -> accepts a fxn - this callback does the processing 
	// to figure out whether it shoudl be sent to filter or reaper
	// start server
	// accept connection
}
