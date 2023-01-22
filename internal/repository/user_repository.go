package repository

import (
	"log"

	"github.com/m-sadykov/go-example-app/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Save(*models.User) (*models.User, error)
	Get(email string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (repo userRepository) Save(u *models.User) (*models.User, error) {
	err := repo.db.Create(&u)

	if err != nil {
		log.Fatal(err)
		return nil, err.Error
	}

	return repo.Get(u.Email)
}

func (repo userRepository) Get(email string) (*models.User, error) {
	var user models.User

	err := repo.db.Where(&models.User{Email: email}).First(&user).Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		log.Fatal(err)
		return nil, err
	}

	return &user, nil
}
