package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
	"os"
	"path/filepath"
	"strings"
)

var movieDTO MovieDTO

type MovieDTO struct {
	Caption string  `json:"name"`
	Movies  []Movie `json:"movies"`
}

type Movie struct {
	Name  string          `json:"name"`
	Photo model.PhotoBase `json:"photo"`
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

	for i := 0; i < count; i++ {
		var movie Movie
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		movie.Name = utils.MovieNames[nameIndex]
		movie.Photo = photos[index]
		movie.Photo.Key = -1
		movie.Photo.ThumbSize = 540
		movie.Photo.VideoFormat = videoFormats[movie.Photo.Name]
		//fmt.Println(movie.Photo.VideoFormat)
		//movie.Photo.PaintWidth = float32(photos[index].Width / 4)
		//movie.Photo.PaintHeight = float32(photos[index].Height / 4)
		//movie.Photo.Crop = 1
		//movie.Photo.Round = 10
		movie.Photo.IsVideo = true
		movie.Photo.PaintWidth = float32(movie.Photo.Width)
		movie.Photo.PaintHeight = float32(movie.Photo.Height)
		dto.Movies = append(dto.Movies, movie)
		nameIndex++
		index++
	}

	index = 0

	return dto
}

// ListVideoFormatsInDirectory reads a directory and returns a map of video formats.
func ListVideoFormatsInDirectory(dir string) (map[string]string, error) {
	videoFormats := make(map[string]string)

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
				filename = strings.ReplaceAll(filename, ".mp4", "")
				filename = strings.ReplaceAll(filename, ".mkv", "")
				videoFormats[filename] = format // Store the filename and format
				fmt.Println(filename)
			} else {
				// Optionally log the unsupported format
				fmt.Println(err)
			}
		}
	}

	return videoFormats, nil
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
