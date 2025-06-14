package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
)

var cameraDTO CameraDTO

type CameraDTO struct {
	Caption string   `json:"caption"`
	Cameras []Camera `json:"cameras"`
}

type Camera struct {
	Name  string        `json:"name"`
	Photo model.UIImage `json:"photo"`
}

func GetCamera(folder string) CameraDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto CameraDTO

	//if count > 50 {
	//	count = 50
	//}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var camera Camera
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		camera.Name = utils.MovieNames[nameIndex]

		camera.Photo = photos[index]

		dto.Cameras = append(dto.Cameras, camera)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
