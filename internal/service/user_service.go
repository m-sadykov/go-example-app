package service

import (
	"errors"
	"log"

	"github.com/m-sadykov/go-example-app/internal/repository"
	"github.com/m-sadykov/go-example-app/models"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	Save(models.User) (models.User, error)
}

func NewUserService(repo repository.UserRepository) UserService {
	return userService{userRepository: repo}
}

func (u userService) Save(user models.User) (models.User, error) {
	existingUser, err := u.userRepository.Get(*user.Email)
	if err != nil {
		log.Fatal(err)
	}

	if existingUser.Email == user.Email {
		log.Fatal(err)
		err = errors.New("user with email already exists")
	}

	return u.userRepository.Save(user)
}
