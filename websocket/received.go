package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/repo"
	"log"
	"strings"
)

func Received(messageType int, msg []byte) {

	var jsonString = string(msg)

	fmt.Println(jsonString)
	// Create a variable to hold the unmarshalled data
	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Now you can access the data dynamically

	fmt.Println("-------------------------")

	for key, value := range result {
		if strings.Compare(key, "json") == 0 {

			// Attempt to convert the interface{} to string
			str, ok := value.(string)
			if !ok {
				fmt.Println("Value is not a string")
			} else {
				//fmt.Println("Converted string:", str)
				processJsonString(str, jsonString)
			}
			break
		}
		//fmt.Printf("%s: %v\n", key, value)
	}

	fmt.Println("-------------------------")

	//if strings.Compare(text, "lamp on") == 0 {
	//	sendMQTT(text)
	//} else if strings.Compare(text, "lamp off") == 0 {
	//	sendMQTT(text)
	//	//sendText("lamp is off")
	//} else if strings.Compare(text, "lamps") == 0 {
	//	err := processLamps()
	//	if err != nil {
	//		return
	//	}
	//} else if strings.Compare(text, "temperature") == 0 || strings.Compare(text, "tmp") == 0 {
	//	processSingleTemperature()
	//} else if strings.Compare(text, "chart") == 0 {
	//	processTemperatureChart()
	//} else {
	//	SendText("command is wrong!")
	//}
}

func processJsonString(class string, jsonString string) {
	switch class {
	case "temp":
		processLamps()
		break
	case "FetchChats":
		processFetchChats()
		break
	case "SetChatPin":

		var command model.CommandChatPin
		err := UnmarshalJSON(jsonString, &command)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
		err = repo.UpdateLastPinChange(command.ChatID, command.IsPin)
		if err != nil {
			return
		}
		break
	case "FetchChatId":

		break
	case "Message":
		var message model.Message
		err := json.Unmarshal([]byte(jsonString), &message)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}

		fmt.Printf("Content: %s\n", message.Content)

		id, err := repo.DatabaseSaveMessage(message)
		// Check for errors
		if err != nil {
			fmt.Println("Error saving temperature:", err)
		} else {
			fmt.Println("Message saved successfully with ID:", id)
		}

		if strings.Compare(message.Content, "tmp") == 0 {
			processTemperatureChart()
		}
		break
	default:
		break
	}
}

func UnmarshalJSON(jsonString string, dest interface{}) error {
	err := json.Unmarshal([]byte(jsonString), dest)
	return err
}
