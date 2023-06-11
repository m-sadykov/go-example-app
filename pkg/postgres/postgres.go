package postgres

import (
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const MAX_ATTEMPTS_COUNT = 7

func InitPostgres(dsn string) (*gorm.DB, error) {
	for i := 1; i <= MAX_ATTEMPTS_COUNT; i++ {
		db, err := connectToDb(dsn)

		if err != nil {
			log.Println(err)
			connectToDb(dsn)
		} else {
			return db, nil
		}
	}

	return nil, errors.New("Database connection failed. Exceeded maximum attempts count")
}

func connectToDb(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
