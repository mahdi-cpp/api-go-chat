package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
)

var questionSoundDTO QuestionSoundDTO

type QuestionSoundDTO struct {
	Caption        string          `json:"name"`
	QuestionSounds []QuestionSound `json:"questionSounds"`
}

type QuestionSound struct {
	Name  string        `json:"name"`
	Photo model.UIImage `json:"photo"`
}

func GetQuestionSounds(folder string) QuestionSoundDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto QuestionSoundDTO

	if count > 10 {
		count = 10
	}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var sound QuestionSound
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		sound.Name = utils.MovieNames[nameIndex]

		sound.Photo = photos[index]

		dto.QuestionSounds = append(dto.QuestionSounds, sound)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
