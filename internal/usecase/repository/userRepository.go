package repository

import (
	"fmt"

	"github.com/m-sadykov/go-example-app/internal/entity"
	"gorm.io/gorm"
)

type FindOneParam struct {
	ID    uint
	Name  string
	Email string
}

type UserUpdateParam struct {
	Email    string
	Name     string
	Password string
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

func (r *UserRepository) Update(id uint, param UserUpdateParam) (*entity.User, error) {
	var u entity.User

	db := r.db.Model(&u).Where(u.ID, id)
	if param.Email != "" {
		db.Update("Email", param.Email)
	}

	if param.Name != "" {
		db.Update("Name", param.Name)
	}

	if param.Password != "" {
		db.Update("Password", param.Password)
	}

	if db.Error != nil {
		return nil, fmt.Errorf("failed to update user %v", db.Error)
	}

	return r.Get(FindOneParam{ID: id})
}

func (r *UserRepository) Delete(id uint) {
	r.db.Delete(&entity.User{}, id)
}
