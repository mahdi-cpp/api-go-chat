package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
)

var storyDTO StoryDTO

type StoryDTO struct {
	Caption string        `json:"name"`
	Avatar  model.UIImage `json:"avatar"`
	Stories []Story       `json:"stories"`
}

type Story struct {
	Name  string        `json:"name"`
	Photo model.UIImage `json:"photo"`
}

func GetStory(folder string, avatar string) StoryDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var storyDTO StoryDTO

	if count > 50 {
		count = 50
	}

	var index = 0
	var nameIndex = 0

	var photo = model.UIImage{}
	photo.Name = avatar
	photo.FileType = ".jpg"
	photo.Size.Width = 200
	photo.Size.Height = 200
	photo.AspectRatio = 1

	storyDTO.Avatar = photo

	for i := 0; i < count; i++ {
		var story Story
		if nameIndex >= len(utils.FackNames) {
			nameIndex = 0
		}

		story.Name = utils.FackNames[nameIndex]
		story.Photo = photos[index]

		storyDTO.Stories = append(storyDTO.Stories, story)
		nameIndex++
		index++
	}

	index = 0

	return storyDTO
}
