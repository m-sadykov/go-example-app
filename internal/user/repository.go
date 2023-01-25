package user

import (
	"log"

	"github.com/m-sadykov/go-example-app/internal/user/models"
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

func (self userRepository) Save(u *models.User) (*models.User, error) {
	err := self.db.Create(&u)

	if err != nil {
		log.Fatal(err)
		return nil, err.Error
	}

	return self.Get(u.Email)
}

func (self userRepository) Get(email string) (*models.User, error) {
	var u models.User

	err := self.db.Where(&models.User{Email: email}).First(&u).Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		log.Fatal(err)
		return nil, err
	}

	return &u, nil
}
