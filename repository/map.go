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
	Caption   string            `json:"name"`
	Maps      []Map             `json:"maps"`
	Photos    []model.PhotoBase `json:"photos"`
	MapPhotos []model.PhotoBase `json:"mapPhotos"`
	Users     []string          `json:"users"`

	Cafes []model.PhotoBase `json:"cafes"`
}

type Map struct {
	Name  string          `json:"name"`
	Photo model.PhotoBase `json:"photo"`
}

func GetMaps(folder string) MapDTO {

	var dto MapDTO

	for i := 1; i < 14; i++ {
		var photoBase = model.PhotoBase{}
		photoBase.Name = "chat_" + strconv.Itoa(i)
		photoBase.Key = -1
		photoBase.ThumbSize = 135
		photoBase.Circle = true
		photoBase.PaintWidth = dp(130)
		photoBase.PaintHeight = dp(130)

		dto.Photos = append(dto.Photos, photoBase)
		dto.Users = append(dto.Users, repo.Usernames[i])
	}

	for i := 1; i < 14; i++ {
		var photo = model.PhotoBase{}
		photo.Name = "chat_" + strconv.Itoa(i)
		photo.Key = -1
		photo.ThumbSize = 135
		photo.Circle = true
		photo.PaintWidth = dp(100)
		photo.PaintHeight = dp(100)

		dto.MapPhotos = append(dto.MapPhotos, photo)
	}

	for i := 1; i < 14; i++ {
		var photo = model.PhotoBase{}
		photo.Name = "chat_" + strconv.Itoa(i)
		photo.Key = -1
		photo.ThumbSize = 135
		photo.Circle = true
		photo.PaintWidth = dp(100)
		photo.PaintHeight = dp(100)

		dto.Cafes = append(dto.Cafes, photo)
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
		mapData.Photo.Key = -1
		mapData.Photo.ThumbSize = 540
		mapData.Photo.PaintWidth = dp(70)
		mapData.Photo.PaintHeight = dp(120)

		dto.Maps = append(dto.Maps, mapData)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
