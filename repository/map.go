package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/repo"
	"github.com/mahdi-cpp/api-go-chat/utils"
	"strconv"
)

var mapDTO MapDTO

type MapDTO struct {
	Caption   string          `json:"name"`
	Maps      []Map           `json:"maps"`
	Photos    []model.UIImage `json:"photos"`
	MapPhotos []model.UIImage `json:"mapPhotos"`
	Users     []string        `json:"users"`

	Cafes []model.UIImage `json:"cafes"`
}

type Map struct {
	Name  string        `json:"name"`
	Photo model.UIImage `json:"photo"`
}

func GetMaps(folder string) MapDTO {

	var dto MapDTO

	for i := 14; i < 34; i++ {
		var image = model.UIImage{}
		image.Name = "chat_" + strconv.Itoa(i)
		image.Size.Width = 400
		image.Size.Height = 400

		dto.Photos = append(dto.Photos, image)
		dto.Users = append(dto.Users, repo.Usernames[i])
	}

	for i := 14; i < 34; i++ {
		var image = model.UIImage{}
		image.Name = "chat_" + strconv.Itoa(i)
		image.Size.Width = 400
		image.Size.Height = 400

		dto.MapPhotos = append(dto.MapPhotos, image)
	}

	for i := 14; i < 34; i++ {
		var image = model.UIImage{}
		image.Name = "chat_" + strconv.Itoa(i)
		image.Size.Width = 400
		image.Size.Height = 400

		dto.Cafes = append(dto.Cafes, image)
	}

	var index = 0
	var nameIndex = 0
	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)

	for i := 0; i < count; i++ {
		var mapData Map
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		mapData.Name = utils.MovieNames[nameIndex]

		mapData.Photo = photos[index]

		dto.Maps = append(dto.Maps, mapData)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
