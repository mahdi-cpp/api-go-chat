package repository

import (
	"bytes"
	"fmt"
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var movieDTO MovieDTO

type MovieDTO struct {
	Caption string  `json:"name"`
	Movies  []Movie `json:"movies"`
}

type Movie struct {
	Name  string        `json:"name"`
	Photo model.UIImage `json:"photo"`
}

func GetMovies(folder string) MovieDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto MovieDTO

	dir := "/home/mahdi/files/"
	videoFormats, err := ListVideoFormatsInDirectory(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return MovieDTO{}
	}

	var index = 0
	var nameIndex = 0

	var hackMovie Movie

	for i := 0; i < count; i++ {
		var movie Movie
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		movie.Name = utils.MovieNames[nameIndex]
		movie.Photo = photos[index]
		movie.Photo.VideoFormat = videoFormats[movie.Photo.Name]
		movie.Photo.IsVideo = true
		movie.Photo.HasVideoControl = false

		//fmt.Println(videoFormats[movie.Photo.Name])

		//duration, err := getVideoDuration("/home/mahdi/files/" + movie.Photo.Name + "." + movie.Photo.VideoFormat)
		//if err != nil {
		//	fmt.Println("Error getting duration:"+"/home/mahdi/files/"+movie.Photo.Name+"."+movie.Photo.VideoFormat, err)
		//	continue
		//}
		//movie.Photo.VideoDuration = int(duration)

		if strings.Contains(movie.Photo.Name, "hack") {
			movie.Photo.HasSubtitle = true
			movie.Photo.HasVideoControl = true
			hackMovie = movie
		} else {
			dto.Movies = append(dto.Movies, movie)
		}

		nameIndex++
		index++
	}

	dto.Movies = append(dto.Movies, hackMovie)

	index = 0

	return dto
}

type VideoInfo struct {
	Format   string
	Duration int
}

// ListVideoFormatsInDirectory reads a directory and returns a map of video formats.
func ListVideoFormatsInDirectory(dir string) (map[string]string, error) {
	videoInfoArray := make(map[string]string)

	// Read the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Check each file
	for _, file := range files {
		if !file.IsDir() { // Skip directories
			filename := file.Name()
			format, err := GetVideoFormat(filename)
			if err == nil {
				//// Retrieve video duration
				//duration, err := getVideoDuration(filepath.Join(dir, filename))
				//if err != nil {
				//	fmt.Println("Error getting duration:", err)
				//	continue
				//}

				filename = strings.ReplaceAll(filename, ".mp4", "")
				filename = strings.ReplaceAll(filename, ".mkv", "")
				videoInfoArray[filename] = format
				//VideoInfo{Format: format}
				//videoInfoArray. = format // Store the filename and format
				//VideoInfo[filename] = map[string]interface{}{
				//	"format":   format,
				//	"duration": duration,
				//}
				fmt.Println(filename)
			} else {
				// Optionally log the unsupported format
				fmt.Println(err)
			}
		}
	}

	return videoInfoArray, nil
}

// GetVideoFormat returns the video format based on the file extension.
func GetVideoFormat(filename string) (string, error) {
	// Get the file extension
	ext := strings.ToLower(filepath.Ext(filename))

	// Check for video formats
	switch ext {
	case ".mp4":
		return "mp4", nil
	case ".mkv":
		return "mkv", nil
	default:
		return "", fmt.Errorf("unsupported video format: %s", ext)
	}
}

// getVideoDuration retrieves the duration of a video file in seconds using FFmpeg.
func getVideoDuration(filename string) (float64, error) {
	// Prepare the FFmpeg command to retrieve video info
	cmd := exec.Command("ffmpeg", "-i", filename)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // FFmpeg sends info to stderr

	err := cmd.Run()
	if err != nil {
		return 0, err
	}

	// Parse the output to find the duration
	output := out.String()
	duration := parseDuration(output)
	if duration < 0 {
		return 0, fmt.Errorf("duration not found in output")
	}

	return duration, nil
}

// parseDuration extracts the duration from the FFmpeg output
func parseDuration(output string) float64 {
	// Find the duration line in the output
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Duration:") {
			// Example line: Duration: 00:01:30.00, start: 0.000000, bitrate: 128 kb/s
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				// The duration is in the second part
				durationStr := parts[1] // e.g., "00:01:30.00"
				return convertDurationToSeconds(durationStr)
			}
		}
	}
	return -1 // Indicate that duration was not found
}

// convertDurationToSeconds converts a duration string (HH:MM:SS.sss) to seconds
func convertDurationToSeconds(durationStr string) float64 {
	var hours, minutes, seconds float64
	fmt.Sscanf(durationStr, "%f:%f:%f", &hours, &minutes, &seconds)
	return hours*3600 + minutes*60 + seconds
}
