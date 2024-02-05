package encoderdecoder


import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	
)

func InitClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	// Create a new Message to send to the server
	message := Message{Text: "Hello from the client!"}

	// Create a new JSON encoder to write to the connection
	encoder := json.NewEncoder(conn)

	// Encode and send the Message to the server
	if err := encoder.Encode(message); err != nil {
		log.Println("Client: Error encoding message:", err)
		return
	}

	fmt.Println("Client: Sent message to server")

	// Create a new JSON decoder to read from the connection
	decoder := json.NewDecoder(conn)

	// Decode the received JSON data into a Message struct
	var receivedMessage Message
	if err := decoder.Decode(&receivedMessage); err != nil {
		log.Println("Client: Error decoding message:", err)
		return
	}

	fmt.Println("Client: Received response from server:", receivedMessage.Text)
}
