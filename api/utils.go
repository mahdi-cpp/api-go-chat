package api

import (
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-chat/repository"
	"os"
	"os/exec"
)

func Convert3GPToWAV(inputFile string, outputFile string) error {
	// Prepare the ffmpeg command with additional audio options
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-acodec", "pcm_s16le", "-ac", "1", "-ar", "16000", outputFile)

	// Run the command and capture any error
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to convert file: %w", err)
	}
	return nil
}

func SaveRecordToFile(sound repository.Sound) error {

	// Create or open the file
	file, err := os.Create("/home/mahdi/files/sounds/" + sound.FileName + ".txt")
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Marshal the struct into JSON
	jsonData, err := json.MarshalIndent(sound, "", "    ") // Indent for readability
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Write the JSON string to the file
	_, err = file.WriteString(string(jsonData))
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func SaveJSONStringToFile(jsonString, filename string) error {
	// Create or open the file
	file, err := os.Create("/home/mahdi/files/sounds/" + filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write the JSON string to the file
	_, err = file.WriteString(jsonString)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
