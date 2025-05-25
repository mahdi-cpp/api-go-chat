package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
)

var musicDTO MusicDTO

type MusicDTO struct {
	Caption string  `json:"caption"`
	Musics  []Music `json:"musics"`
}

type Music struct {
	Artist string        `json:"artist"`
	Track  string        `json:"track"`
	Cover  model.UIImage `json:"cover"`
}

func GetMusics(folder string) MusicDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto MusicDTO

	var index = 0
	var nameIndex = 0

	if count > 10 {
		count = 10
	}

	for i := 0; i < count; i++ {
		var music Music
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		music.Artist = utils.MovieNames[nameIndex]

		music.Cover = photos[index]
		dto.Musics = append(dto.Musics, music)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
