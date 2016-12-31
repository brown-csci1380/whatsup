package client

import (
	"github.com/brown-csci1380/whatsup/whatsup"

	"fmt"
)

func Start(user string, serverPort string, serverAddr string) {
	// Connect to chat server
	_, err := whatsup.ServerConnect(user, serverAddr, serverPort)
	if err != nil {
		fmt.Printf("unable to connect to server: %s\n", err)
		return
	}
	// TODO: Receive input from the user and use the first return value of whatsup.ServerConnect
	// (currently ignored so the stencil will compile) to talk to the server.
}
