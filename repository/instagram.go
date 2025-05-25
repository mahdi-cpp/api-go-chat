package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
)

type InstagramPostDTO struct {
	Caption string          `json:"caption"`
	Avatar  model.UIImage   `json:"avatar"`
	Photos  []model.UIImage `json:"photos"`
}

//curl -fsSL https://ollama.com/install.sh | sh

var instagramPostDTO1 InstagramPostDTO
var instagramPostDTO2 InstagramPostDTO
var instagramPostDTO3 InstagramPostDTO

func GetInstagram(folder string, avatar string) InstagramPostDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var localInstagramPostDTO InstagramPostDTO

	if count > 11 {
		count = 11
	}

	var index = 0
	var nameIndex = 0

	var photo = model.UIImage{}
	photo.Name = avatar
	photo.FileType = ".jpg"
	photo.Size.Width = 50
	photo.Size.Height = 50
	photo.IsVideo = false
	localInstagramPostDTO.Avatar = photo

	localInstagramPostDTO.Caption = "Mahdi"

	for i := 0; i < count; i++ {
		if nameIndex >= len(utils.FackNames) {
			nameIndex = 0
		}

		var photo = model.UIImage{}
		photo = photos[index]

		localInstagramPostDTO.Photos = append(localInstagramPostDTO.Photos, photo)

		nameIndex++
		index++
	}

	index = 0

	return localInstagramPostDTO
}
