package main

import (
	"github.com/mahdi-cpp/api-go-chat/cache"
	"github.com/mahdi-cpp/api-go-chat/redis_service"
	"github.com/mahdi-cpp/api-go-chat/repo"
	"github.com/mahdi-cpp/api-go-chat/repository"
	"github.com/mahdi-cpp/api-go-chat/websocket"
)

func main() {

	repo.InitChatRepository()

	//repo.CreateFakes()

	//repo.FakeChats()

	repository.InitModels()
	cache.ReadIcons()

	websocket.Start()

	redis_service.Start()
	//redis_service.Get()

	Run()
}
