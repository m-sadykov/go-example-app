package repository

import (
	"github.com/m-sadykov/go-example-app/internal/entity"
	"gorm.io/gorm"
)

type FindOneParam struct {
	ID    uint
	Name  string
	Email string
}

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

	return r.Get(FindOneParam{Email: u.Email})
}

func (r *UserRepository) Get(param FindOneParam) (*entity.User, error) {
	var u entity.User

	res := r.db.Where(param).First(&u)

	if res.Error != nil {
		err := res.Error
		if err.Error() == "record not found" {
			return nil, nil
		}

		return nil, err
	}

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return &u, nil
}
