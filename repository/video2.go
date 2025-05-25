package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-chat/cache"
)

var video2DTO MovieDTO

func GetVideo2(folder string) MovieDTO {

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

		movie.Name = "Mahdi"
		movie.Photo = photos[index]
		movie.Photo.VideoFormat = videoFormats[movie.Photo.Name]
		movie.Photo.IsVideo = true
		movie.Photo.HasVideoControl = false

		dto.Movies = append(dto.Movies, movie)

		nameIndex++
		index++
	}

	dto.Movies = append(dto.Movies, hackMovie)

	index = 0

	return dto
}
