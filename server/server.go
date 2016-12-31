package server

import (
	"github.com/brown-csci1380/whatsup/whatsup"

	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// TODO: Implement handling messages from a client.
	// You will find whatsup.SendMsg and whatsup.RecvMsg methods useful for
	// serializing and deserializing messages.

}

func Start() {
	listen, port, err := whatsup.OpenListener()
	fmt.Printf("Listening on port %d\n", port)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept() // this blocks until connection or error
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
