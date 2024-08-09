package entity

import (
	"gorm.io/gorm"
)

type AccessToken struct {
	gorm.Model
	Token  string `gorm:"not null"`
	UserID uint
	User   User
}
