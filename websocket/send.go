package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/repository"
)

var clientID = "android_mahdi_galaxy_a51_client_id"

func processLamps() error {

	var lamps []model.Lamp

	var a = model.Lamp{Name: "Living Room", Icon: "lap-4-100.png", Value: 1}
	var b = model.Lamp{Name: "Ali Room", Icon: "icons8-lights-100.png", Value: 1}
	var c = model.Lamp{Name: "Mahdi Room", Icon: "icons8-table-lights-100.png", Value: 0}
	var d = model.Lamp{Name: "Kitchen 3", Icon: "lap-4-100.png", Value: 0}
	var a1 = model.Lamp{Name: "Room Lamp", Icon: "lap-4-100.png", Value: 1}
	var b1 = model.Lamp{Name: "Lamp 21", Icon: "lap-4-100.png", Value: 1}
	var c1 = model.Lamp{Name: "Lamp 31", Icon: "lap-4-100.png", Value: 0}
	var d1 = model.Lamp{Name: "Lamp 13", Icon: "lap-4-100.png", Value: 1}

	lamps = append(lamps, a)
	lamps = append(lamps, b)
	lamps = append(lamps, c)
	lamps = append(lamps, d)
	lamps = append(lamps, a1)
	lamps = append(lamps, b1)
	lamps = append(lamps, c1)
	lamps = append(lamps, d1)

	// Convert the slice of Temperature structs to JSON
	jsonData, err := json.Marshal(lamps)
	if err != nil {
		return err
	}

	var object model.Object
	object.Type = "lamps"
	object.JsonString = "{lamps:" + string(jsonData) + "}"

	// Convert the Object to a byte slice
	jsonBytes, err := ConvertObjectToBytes(object)
	if err != nil {
		fmt.Println("Error converting object to bytes:", err)
		return err
	}

	SendToClient(clientID, jsonBytes)

	return nil
}

func processTemperatureChart() {

	//Read temperatures from the last hour and get the JSON string
	jsonResult, err := repository.ReadTemperatures()
	if err != nil {
		fmt.Println("Error reading recent temperatures:", err)
		return
	}

	var object model.Object
	object.Type = "temperature_chart"
	object.JsonString = "{temperatures:" + jsonResult + "}"

	// Convert the Object to a byte slice
	jsonBytes, err := ConvertObjectToBytes(object)
	if err != nil {
		fmt.Println("Error converting object to bytes:", err)
		return
	}

	SendToClient(clientID, jsonBytes)
}

func processSingleTemperature() {

	var temp = model.Lamp{Name: "Living Room", Icon: "icons8-temperature-100.png", Value: 13}

	// Convert the slice of Temperature structs to JSON
	jsonData, err := json.Marshal(temp)
	if err != nil {
		return
	}

	var object model.Object
	object.Type = "temperature"
	object.JsonString = string(jsonData)

	// Convert the Object to temp byte slice
	jsonBytes, err := ConvertObjectToBytes(object)
	if err != nil {
		fmt.Println("Error converting object to bytes:", err)
		return
	}

	SendToClient(clientID, jsonBytes)

}

func processText(text string) {

	var temp = model.Lamp{Name: "Living Room", Icon: "icons8-temperature-100.png", Value: 13}

	// Convert the slice of Temperature structs to JSON
	jsonData, err := json.Marshal(temp)
	if err != nil {
		return
	}

	var object model.Object
	object.Type = "temperature"
	object.JsonString = string(jsonData)

	// Convert the Object to temp byte slice
	jsonBytes, err := ConvertObjectToBytes(object)
	if err != nil {
		fmt.Println("Error converting object to bytes:", err)
		return
	}

	SendToClient(clientID, jsonBytes)

}
