package postgres

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var MAX_ATTEMPTS_COUNT = 7

func InitPostgres(dsn string) (*gorm.DB, error) {
	db, err := connectToDb(dsn)
	for err != nil {
		if err != nil {
			log.Println(err)
		}

		if MAX_ATTEMPTS_COUNT > 1 {
			MAX_ATTEMPTS_COUNT--

			time.Sleep(5 * time.Second)
			connectToDb(dsn)

			continue
		}

		return nil, errors.New("Database connection failed. Exceeded maximum attempts count")
	}

	return db, nil
}

func connectToDb(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
