package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
)

var animationDTO AnimationDTO

type AnimationDTO struct {
	Caption    string      `json:"name"`
	Animations []Animation `json:"animations"`
}

type Animation struct {
	Name  string        `json:"name"`
	Photo model.UIImage `json:"photo"`
}

func GetAnimation(folder string) AnimationDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto AnimationDTO

	if count > 20 {
		count = 20
	}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var animation Animation
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		animation.Name = utils.MovieNames[nameIndex]

		animation.Photo = photos[index]

		dto.Animations = append(dto.Animations, animation)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
