package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar;not null"`
	Email    string `gorm:"type:varchar;not null"`
	Password string `gorm:"type:varchar;not null"`
}
