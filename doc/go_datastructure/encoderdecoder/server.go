package encoderdecoder

import (
	"encoding/json"
	"fmt"
	
	"log"
	"net"
)

// Message represents the data structure we want to send between client and server
type Message struct {
	Text string
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a new JSON decoder to read from the connection
	decoder := json.NewDecoder(conn)

	// Decode the received JSON data into a Message struct
	var receivedMessage Message
	if err := decoder.Decode(&receivedMessage); err != nil {
		log.Println("Server: Error decoding message:", err)
		return
	}

	fmt.Println("Server: Received message from client:", receivedMessage.Text)

	// Create a new Message to send back to the client
	responseMessage := Message{Text: "Hello from the server!"}

	// Create a new JSON encoder to write to the connection
	encoder := json.NewEncoder(conn)

	// Encode and send the response Message to the client
	if err := encoder.Encode(responseMessage); err != nil {
		log.Println("Server: Error encoding message:", err)
		return
	}

	fmt.Println("Server: Sent response to client")
}

