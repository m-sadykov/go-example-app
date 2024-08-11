package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/internal/util"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc UserUseCase) Create(d entity.User) (*entity.User, error) {
	password, err := util.HashPassword(d.Password)
	if err != nil {
		return nil, errors.New("failed to create password")
	}

	input := &entity.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: password,
	}

	newUser, err := uc.repo.Store(input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return newUser, nil
}

func (uc UserUseCase) GetOneById(id uint) (*entity.User, error) {
	return uc.repo.Get(repository.FindOneParam{ID: id})
}

func (uc UserUseCase) Update(id uint, param repository.UserUpdateParam) (*entity.User, error) {
	existingUser, err := uc.repo.Get(repository.FindOneParam{ID: id})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if existingUser == nil {
		return nil, fmt.Errorf("user with given id: %d not found", id)
	}

	return uc.repo.Update(id, param)
}

func (uc UserUseCase) Delete(id uint) {
	uc.repo.Delete(id)
}
