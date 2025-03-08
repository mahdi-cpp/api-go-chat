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
