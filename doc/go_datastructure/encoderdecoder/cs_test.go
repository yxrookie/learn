package encoderdecoder

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestCs(t *testing.T) {
	// start a server
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on localhost:8080")

	// start a client
	go InitClient()

	// should keep listen until have the problem
	// for {
	// now only listen 1, just for test...
	for i:= 0; i < 1; i++ {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
        
		go HandleConnection(conn)
	}
	
}