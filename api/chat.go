package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-chat/repository_chat"
	"net/http"
)

func AddChatRoute(rg *gin.RouterGroup) {

	route := rg.Group("/chat")

	route.GET("/chatV2", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository_chat.RestChatV2())
	})

}
