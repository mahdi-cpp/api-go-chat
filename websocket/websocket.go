package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mahdi-cpp/api-go-chat/repo"
	"log"
	"math/rand"
	"net/http"
	"time"

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

	go func() {
		SendData()
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
	//	return               -097
	//}

	//if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
	//	fmt.Println("Error while writing message:", err)
	//}
}

//adb shell am start -W -a android.intent.action.VIEW -d "https://www.ali.com/mahdi_path?message=HelloWorld" com.helium.chat

// SendData periodically sends data to the client.
func SendData() {

	ticker := time.NewTicker(50 * time.Second) // Change the duration as needed
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var i = rand.Intn(4)
			switch i {
			case 0:
				SendNewMessageToClients()
				break
			case 1:
				SendSound()
				break
			case 2:
				SendSound()
				break
			case 3:
				SendLocation()
				break
			case 4:
				break
			}
		}
	}
}
func SendSound() {

	message := model.Message{
		ID:        rand.Intn(1000) + 1,
		ChatID:    rand.Intn(70) + 1,
		UserID:    1,
		Type:      "sound",
		Content:   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	SendByWebsocket2("new_message", message)
}

func SendLocation() {

	message := model.Message{
		ID:        rand.Intn(1000) + 1,
		ChatID:    rand.Intn(70) + 1,
		UserID:    1,
		Type:      "location",
		Content:   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	SendByWebsocket2("new_message", message)
}

func SendNewMessageToClients() {

	message := model.Message{
		ID:        rand.Intn(1000) + 1,
		ChatID:    rand.Intn(70) + 1,
		UserID:    1,
		Content:   repo.EnglishMessages[rand.Intn(len(repo.EnglishMessages))],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	SendByWebsocket2("new_message", message)
}

func processFetchChats() {

	chatList, err := repo.FetchChatsForUser(1)
	if err != nil {
		return
	}

	SendByWebsocket("chatList", chatList)
}
