package repository

import (
	"github.com/gin-gonic/gin"
	"math"
)

func RestChatV2() map[string]any {
	soundDto2 = GetSounds()
	return gin.H{
		"instagramPostDTO1": instagramPostDTO1,
		"instagramPostDTO2": instagramPostDTO2,
		"instagramPostDTO3": instagramPostDTO3,
		"storyDTO":          storyDTO,
		"movieDTO":          movieDTO,
		"video2DTO":         video2DTO,
		"musicDTO":          musicDTO,
		"animationDTO":      animationDTO,
		"pdfDTO":            pdfDTO,
		"electronicDTO":     electronicDTO,
		"mapDTO":            mapDTO,
		"questionSoundDTO":  questionSoundDTO,
		"cameraDTO":         cameraDTO,
		"soundDTO":          soundDto2,
		"factureDTO":        factureDTO,
	}
}
func RestSounds() map[string]any {
	soundDto = GetSounds()
	return gin.H{
		"sounds": soundDto.Sounds,
	}
}

func RestMusic() map[string]any {
	return gin.H{
		"caption": musicDTO.Caption,
		"musics":  musicDTO.Musics,
	}
}
func RestSubtitle() map[string]any {
	return gin.H{
		"name":          newSubTitle.Name,
		"subtitleItems": newSubTitle.Subtitles,
	}
}

var newSubTitle *SubtitleDTO

func InitModels() {

	instagramPostDTO1 = GetInstagram("/var/cloud/id/messi/", "chat_12")
	instagramPostDTO2 = GetInstagram("/var/cloud/fa/", "chat_25")
	instagramPostDTO3 = GetInstagram("/var/cloud/id/ali/", "chat_18")

	storyDTO = GetStory("/var/cloud/fa/", "ma")

	video2DTO = GetMovies("/var/cloud/id/ali/")
	movieDTO = GetVideo2("/var/cloud/id/ali/")

	animationDTO = GetAnimation("/var/cloud/chat/movie/animation/")

	pdfDTO = GetPdfs("/var/cloud/chat/pdf/")

	electronicDTO = GetElectronic("/var/cloud/behance/ali/")

	mapDTO = GetMaps("/var/cloud/chat/map/")
	questionSoundDTO = GetQuestionSounds("/var/cloud/chat_users/")

	cameraDTO = GetCamera("/var/cloud/tinyhome/06/")

	musicDTO = GetMusics("/var/cloud/music/2/")

	factureDTO = GetPhotoListDTO("/var/cloud/facture/", "لیست فاکتور های آبان 1403", "در این مقاله از بخش راهنمای آنلاین محک به یکی دیگر از آموزش های نرم افزار حسابداری محک می‌پردازیم. این آموزش در مورد فاکتور فروش است و فیلم آموزش ثبت و صدور فاکتور فروش در نرم افزار حسابداری محک را می‌توانید در ادامه مشاهده نمایید.")

}

func ReloadSubtitle() {
	newSubTitle, _ = GetSubtitle()
}

func dp(value float32) float32 {
	if value == 0 {
		return 0
	}
	return float32(math.Ceil(float64(2.625 * value)))
}
