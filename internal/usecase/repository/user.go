package repository

import (
	"log"

	"github.com/m-sadykov/go-example-app/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// type UserRepository interface {
// 	Save(*entity.User) (*entity.User, error)
// 	Get(email string) (*entity.User, error)
// }

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(u *entity.User) (*entity.User, error) {
	err := r.db.Create(&u)

	if err != nil {
		log.Fatal(err)
		return nil, err.Error
	}

	return r.Get(u.Email)
}

func (r *UserRepository) Get(email string) (*entity.User, error) {
	var u entity.User

	err := r.db.Where(&entity.User{Email: email}).First(&u).Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		log.Fatal(err)
		return nil, err
	}

	return &u, nil
}
