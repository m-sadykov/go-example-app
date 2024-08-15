package util

import (
	faker "github.com/go-faker/faker/v4"
	"github.com/m-sadykov/go-example-app/internal/entity"
	"github.com/m-sadykov/go-example-app/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ClearDatabase(db *gorm.DB) {
	db.Exec("delete from public.users")
	db.Exec("delete from public.access_tokens")
}

func CreateUser(userRepo repository.UserRepository) (*entity.User, error) {
	password, _ := HashPassword("123")

	return userRepo.Store(&entity.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: password,
	})
}
