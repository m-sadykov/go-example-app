package usecase

import (
	"crypto/sha1"
	"fmt"

	"github.com/m-sadykov/go-example-app/config"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/usecase/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc UserUseCase) Save(d entity.User) (*entity.User, error) {
	newUser := &entity.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: hashPassword(d.Password),
	}

	u, err := uc.repo.Save(newUser)
	if err != nil {
		return newUser, fmt.Errorf("UserUseCase - Save - uc.repo.Save: %w", err)
	}

	return u, nil
}

func (us UserUseCase) Get(email string) (*entity.User, error) {
	return us.repo.Get(email)
}

func hashPassword(password string) string {
	cfg := config.InitConfig()

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(cfg.SALT_ROUNDS))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	return password
}
