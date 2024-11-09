package repository_chat

import (
	"github.com/gin-gonic/gin"
	"math"
)

func RestChatV2() map[string]any {
	return gin.H{
		"instagramPostDTO1": instagramPostDTO1,
		"instagramPostDTO2": instagramPostDTO2,
		"instagramPostDTO3": instagramPostDTO3,
		"storyDTO":          storyDTO,
		"movieDTO":          movieDTO,
		"animationDTO":      animationDTO,

		"pdfDTO":           pdfDTO,
		"electronicDTO":    electronicDTO,
		"mapDTO":           mapDTO,
		"questionVoiceDTO": questionVoiceDTO,
		"cameraDTO":        cameraDTO,
	}
}

func InitModels() {
	instagramPostDTO1 = GetInstagram("/var/cloud/reynardlowell/", "far")
	instagramPostDTO2 = GetInstagram("/var/cloud/paris/", "narges2")
	instagramPostDTO3 = GetInstagram("/var/cloud/reynardlowell/", "01")

	storyDTO = GetStory("/var/cloud/fa/", "ma")

	movieDTO = GetMovies("/var/cloud/chat/movie/movie/")
	animationDTO = GetAnimation("/var/cloud/chat/movie/animation/")

	pdfDTO = GetPdfs("/var/cloud/chat/pdf/")
	electronicDTO = GetElectronic("/var/cloud/chat/electronic/")
	mapDTO = GetMaps("/var/cloud/chat/map/")
	questionVoiceDTO = GetQuestionVoices("/var/cloud/fa/")

	cameraDTO = GetCamera("/var/cloud/camera-sequrity/")
}

func dp(value float32) float32 {
	if value == 0 {
		return 0
	}
	return float32(math.Ceil(float64(2.625 * value)))
}
