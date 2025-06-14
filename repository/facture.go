package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
)

var factureDTO PhotoListDTO

type PhotoListDTO struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Photos      []model.UIImage `json:"photos"`
}

func GetPhotoListDTO(folder string, title string, description string) PhotoListDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto PhotoListDTO

	var index = 0

	dto.Title = title
	dto.Description = description

	for i := 0; i < count; i++ {
		var photo model.UIImage
		photo = photos[index]
		dto.Photos = append(dto.Photos, photo)
		index++
	}

	index = 0

	return dto
}
