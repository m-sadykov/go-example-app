package user

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/m-sadykov/go-example-app/internal/config"
	dto "github.com/m-sadykov/go-example-app/internal/user/dto"
	"github.com/m-sadykov/go-example-app/internal/user/models"
)

type userService struct {
	userRepository UserRepository
}

type UserService interface {
	Save(dto.UserCreateDto) (*models.User, error)
	Get(email string) (*models.User, error)
}

func NewUserService(repo UserRepository) UserService {
	return userService{userRepository: repo}
}

func (self userService) Save(d dto.UserCreateDto) (*models.User, error) {
	existingUser, err := self.userRepository.Get(d.Email)

	if existingUser != nil && existingUser.Email == d.Email {
		err = errors.New("user with email already exists")
		return nil, err
	}

	u := &models.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: hashPassword(d.Password),
	}

	return self.userRepository.Save(u)
}

func (self userService) Get(email string) (*models.User, error) {
	return self.userRepository.Get(email)
}

func hashPassword(password string) string {
	cfg := config.InitConfig()

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(cfg.SALT_ROUNDS))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	return password
}
