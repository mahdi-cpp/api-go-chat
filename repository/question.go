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
	Name  string          `json:"name"`
	Photo model.PhotoBase `json:"photo"`
}

func GetQuestionSounds(folder string) QuestionSoundDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto QuestionSoundDTO

	//if count > 50 {
	//	count = 50
	//}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var sound QuestionSound
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		sound.Name = utils.MovieNames[nameIndex]

		sound.Photo = photos[index]
		sound.Photo.Key = -1
		sound.Photo.ThumbSize = 270
		sound.Photo.Circle = true
		sound.Photo.Round = int(dp(10))
		sound.Photo.PaintWidth = dp(95)
		sound.Photo.PaintHeight = dp(95)

		dto.QuestionSounds = append(dto.QuestionSounds, sound)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
