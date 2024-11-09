package main

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/config"
	"github.com/mahdi-cpp/api-go-chat/repository_chat"
)

func main() {

	config.LayoutInit()
	repository_chat.InitModels()
	cache.ReadIcons()
	Run()
}
