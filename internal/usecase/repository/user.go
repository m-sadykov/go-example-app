package repository

import (
	"github.com/m-sadykov/go-example-app/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Store(u *entity.User) (*entity.User, error) {
	res := r.db.Create(&u)

	if res.Error != nil {
		return nil, res.Error
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
		return nil, err
	}

	return &u, nil
}
