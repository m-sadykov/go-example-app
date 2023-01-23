package service

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/m-sadykov/go-example-app/internal/config"
	dto "github.com/m-sadykov/go-example-app/internal/controller/dto/user"
	"github.com/m-sadykov/go-example-app/internal/models"
	"github.com/m-sadykov/go-example-app/internal/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	Save(dto.UserCreateDto) (*models.User, error)
	Get(email string) (*models.User, error)
}

func NewUserService(repo repository.UserRepository) UserService {
	return userService{userRepository: repo}
}

func (u userService) Save(d dto.UserCreateDto) (*models.User, error) {
	existingUser, err := u.userRepository.Get(d.Email)

	if existingUser != nil && existingUser.Email == d.Email {
		err = errors.New("user with email already exists")
		return nil, err
	}

	user := &models.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: hashPassword(d.Password),
	}

	return u.userRepository.Save(user)
}

func (u userService) Get(email string) (*models.User, error) {
	return u.userRepository.Get(email)
}

func hashPassword(password string) string {
	cfg := config.InitConfig()

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(cfg.SALT_ROUNDS))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	return password
}
