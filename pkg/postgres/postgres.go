package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return db
}
