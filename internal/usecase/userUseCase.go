package usecase

import (
	"crypto/sha1"
	"fmt"
	"log"

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

func (uc UserUseCase) Create(d entity.User) (*entity.User, error) {
	newUser := &entity.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: hashPassword(d.Password),
	}

	u, err := uc.repo.Store(newUser)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return u, nil
}

func (us UserUseCase) GetOneById(id uint) (*entity.User, error) {
	return us.repo.Get(repository.FindOneParam{ID: id})
}

func hashPassword(password string) string {
	cfg := config.InitConfig()

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(cfg.SALT_ROUNDS))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	return password
}
