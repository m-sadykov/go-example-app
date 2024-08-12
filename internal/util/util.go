package util

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	var err error

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
