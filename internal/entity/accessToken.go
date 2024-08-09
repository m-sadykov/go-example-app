package entity

import (
	"time"

	"gorm.io/gorm"
)

type AccessToken struct {
	gorm.Model
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    uint
	User      User
}
