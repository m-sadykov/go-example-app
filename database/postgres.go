package database

import (
	"log"

	"github.com/m-sadykov/go-example-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	DB = db
}
