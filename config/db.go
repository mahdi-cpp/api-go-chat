package config

import (
	"github.com/mahdi-cpp/api-go-chat/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var DB *gorm.DB
var err error

// docker run -d --name web-portainer --restart always -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce

// docker run -d --name postgres --restart always -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin@123456 -e POSTGRES_DB=mqtt -p 5432:5432 postgres

// docker run -d --name redis_service --restart always -p 6379:6379 redis_service
// docker run -d --name nginx --restart=always -p 8081:80 -v /var/cloud:/usr/share/nginx/html nginx

func DatabaseInit() *gorm.DB {
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=PostgreSQL user=admin password=admin@123456 dbname=chat port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "api_v1.", // schema name
			SingularTable: false,
		}})

	if err != nil {
		println("Failed to connect database gallery\"")
		os.Exit(1)
	}

	err = DB.AutoMigrate(&model.Message{}, &model.User{}, &model.Chat{})
	if err != nil {
		return nil
	}

	return DB
}
