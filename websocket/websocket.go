package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"

	"github.com/mahdi-cpp/api-go-chat/model"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (for development purposes)
	},
}

type Client struct {
	conn *websocket.Conn
	id   string // Unique identifier for the client
}

var (
	clients   = make(map[string]*Client) // Map to hold connected clients
	clientsMu sync.Mutex                 // Mutex to protect access to the clients map
)

func Start() {
	http.HandleFunc("/ws", handleConnections)
	//go handleMessages() // Start the message handling goroutine

	// Start the WebSocket server in a separate goroutine
	go func() {
		log.Println("Websocket Server started on :8097")
		err := http.ListenAndServe(":8097", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	clientID := r.URL.Query().Get("id") // Get client ID from query parameters
	client := &Client{conn: conn, id: clientID}

	clientsMu.Lock()
	clients[clientID] = client // Store the client in the map
	clientsMu.Unlock()

	for {
		messageType, msg, err := conn.ReadMessage() // Keep the connection alive
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		Received(messageType, msg)
	}

	clientsMu.Lock()
	delete(clients, clientID) // Remove the client from the map when done
	clientsMu.Unlock()
}

//func handleMessages() {
//	for {
//		// Here you would implement your logic to send messages to the Android client.
//		// For example, you could use a channel to receive messages to send.
//
//		// Example: Sending a message to a specific client
//		clientsMu.Lock()
//		if androidClient, ok := clients["android_mahdi_galaxy_a51_client_id"]; ok { // Replace with actual ID
//			err := androidClient.conn.WriteMessage(websocket.TextMessage, []byte("Hello Android Client!"))
//			if err != nil {
//				log.Println("Error sending message to Android client:", err)
//			}
//		}
//		clientsMu.Unlock()
//
//		// Sleep or wait for an event before sending the next message
//		// time.Sleep(time.Second) // Uncomment if you want to send messages periodically
//	}
//}

//---------------------------------------------------------

func SendToClient(clientId string, text []byte) {

	fmt.Println("send websocket:")
	fmt.Println(string(text))

	clientsMu.Lock()
	if androidClient, ok := clients[clientId]; ok {
		err := androidClient.conn.WriteMessage(websocket.TextMessage, text)
		if err != nil {
			log.Println("Error sending message to Android client:", err)
		}
	}
	clientsMu.Unlock()
}

func SendText(text string) {

	var object model.Object
	object.Type = "text"
	object.JsonString = text

	//jsonBytes, err := utils.ConvertObjectToBytes(object)
	//if err != nil {
	//	fmt.Println("Error converting object to bytes:", err)
	//	return
	//}

	//if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
	//	fmt.Println("Error while writing message:", err)
	//}
}
