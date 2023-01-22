package repository

import (
	"github.com/m-sadykov/go-example-app/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Save(models.User) (models.User, error)
	Get(email string) (models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (repo userRepository) Save(user models.User) (models.User, error) {
	err := repo.db.Create(&user)
	return user, err.Error
}

func (repo userRepository) Get(email string) (models.User, error) {
	err := repo.db.First(&models.User{}, email)
	return models.User{}, err.Error
}
