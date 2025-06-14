package repository

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/model"
	"github.com/mahdi-cpp/api-go-chat/utils"
)

var pdfDTO PdfDTO

type PdfDTO struct {
	Caption string `json:"name"`
	Pdfs    []Pdf  `json:"pdfs"`
}

type Pdf struct {
	Name  string        `json:"name"`
	Photo model.UIImage `json:"photo"`
}

func GetPdfs(folder string) PdfDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var pdfDTO PdfDTO

	//if count > 50 {
	//	count = 50
	//}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var pdf Pdf
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		pdf.Name = utils.MovieNames[nameIndex]

		pdf.Photo = photos[index]

		pdfDTO.Pdfs = append(pdfDTO.Pdfs, pdf)
		nameIndex++
		index++
	}

	index = 0

	return pdfDTO
}
