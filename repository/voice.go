package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type SoundDTO struct {
	Sounds []Sound `json:"sounds"`
}

var soundDto SoundDTO
var soundDto2 SoundDTO

type Sound struct {
	FileName    string  `json:"fileName"`
	Duration    int     `json:"duration"`
	Description string  `json:"description"`
	Timestamp   int64   `json:"timestamp"`
	Signals     []int16 `json:"signals"`
}

func GetSounds() SoundDTO {

	dir := "/home/mahdi/files/sounds"
	var soundDTO SoundDTO

	sounds, err := readTxtFiles(dir)
	log.Println("sounds count: ", len(sounds))

	if err != nil {
		log.Fatalf("Error reading files: %v", err)
	}

	// Print the data for demonstration
	for _, sound := range sounds {
		soundJSON, _ := json.MarshalIndent(sound, "", "  ")
		fmt.Printf("Sound: %s\n", soundJSON)
	}
	soundDTO.Sounds = sounds
	return soundDTO
}

// Function to read .txt files from a directory and convert their content to Sound structs
func readTxtFiles(dir string) ([]Sound, error) {
	var sounds []Sound

	// Read all files in the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Iterate through the files
	for _, file := range files {

		// Check if the file is a .txt file

		if filepath.Ext(file.Name()) == ".txt" {

			// Read the content of the file
			data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				log.Printf("Error reading file %s: %v", file.Name(), err)
				continue
			}

			var sound Sound
			err = json.Unmarshal(data, &sound)
			if err != nil {
				log.Fatalf("Error unmarshaling JSON: %v", err)
			}

			sounds = append(sounds, sound)
		}
	}

	return sounds, nil
}

// Function to parse signal data from a string
func parseSignals(signalStr string) []int16 {
	var signals []int16
	signalValues := strings.Fields(signalStr) // Split by whitespace

	for _, value := range signalValues {
		if signal, err := strconv.ParseInt(value, 10, 16); err == nil {
			signals = append(signals, int16(signal))
		} else {
			log.Printf("Invalid signal value: %s, error: %v", value, err)
		}
	}

	return signals
}
