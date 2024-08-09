package usecase

import (
	"errors"
	"log"

	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	repo        repository.UserRepository
	accessToken AccessTokenUseCase
}

func NewAuthUseCase(r repository.UserRepository, uc AccessTokenUseCase) *AuthUseCase {
	return &AuthUseCase{repo: r, accessToken: uc}
}

func (uc AuthUseCase) Login(email, password string) (*entity.AccessToken, error) {
	user, _ := uc.repo.Get(repository.FindOneParam{Email: email})

	if user == nil {
		return nil, errors.New("invalid user email or password")
	}

	err := validateUserPassword(password, user.Password)
	if err != nil {
		return nil, errors.New("invalid user email or password")
	}

	return uc.accessToken.CreateAccessToken(*user)
}

func (uc AuthUseCase) Logout(token string) {
	uc.accessToken.Remove(token)
}

func validateUserPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
