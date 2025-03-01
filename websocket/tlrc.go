package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/repo"
)

func processFetchChats() {

	chatList, err := repo.FetchChatsForUser(1)
	if err != nil {
		return
	}

	sendB("chatList", chatList)
}

func sendB(commandType string, data interface{}) {

	//switch v := data.(type) {
	//case string:
	//	fmt.Println("Data is a string:", v)
	//default:
	//	fmt.Println("Data is not a string, it is of type:", fmt.Sprintf("%T", v))
	//}

	// Convert the structs to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	// Calculate size in bytes
	byteSize := len(string(jsonData))

	// Convert to kilobytes
	kbSize := float64(byteSize) / 1024.0

	fmt.Printf("The size of the string is: %.2f KB\n", kbSize)

	var object model.Object
	object.Type = commandType
	object.JsonString = "{" + commandType + ":" + string(jsonData) + "}"

	// Convert the Object to a byte slice
	jsonBytes, err := ConvertObjectToBytes(object)
	if err != nil {
		fmt.Println("Error converting object to bytes:", err)
		return
	}

	SendToClient(clientID, jsonBytes)
}
