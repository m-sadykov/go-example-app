package repository

import (
	"github.com/m-sadykov/go-example-app/internal/entity"
	"gorm.io/gorm"
)

type AccessTokenRepository struct {
	db *gorm.DB
}

func NewAccessTokenRepository(db *gorm.DB) *AccessTokenRepository {
	return &AccessTokenRepository{db}
}

func (r *AccessTokenRepository) Create(tokenString string, user entity.User) (*entity.AccessToken, error) {

	res := r.db.Create(&entity.AccessToken{
		Token:  tokenString,
		UserID: user.ID,
		User:   user,
	})

	if res.Error != nil {
		return nil, res.Error
	}

	return r.Get(tokenString)
}

func (r *AccessTokenRepository) Get(token string) (*entity.AccessToken, error) {
	var t entity.AccessToken

	res := r.db.Where(&entity.AccessToken{Token: token}).First(&t)
	if res.Error != nil {
		return nil, res.Error
	}

	return &t, nil
}

func (r *AccessTokenRepository) Delete(token string) {
	var t entity.AccessToken

	r.db.Where(&entity.AccessToken{Token: token}).Delete(&t)
}
