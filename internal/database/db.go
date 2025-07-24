package database

import (
	"log"

	"gorm.io/driver/postgres"

	"github.com/sancheschris/shorty/internal/model"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(databaseURL string) {
	var err error
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = DB.AutoMigrate(&model.URL{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
