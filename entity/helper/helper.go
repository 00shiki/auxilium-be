package helper

import (
	"gorm.io/gorm"
)

type Helper struct {
	gorm.Model
	UserID    uint `gorm:"user_id;unique"`
	Username  string
	AvatarURL string
	Lat       float64
	Lon       float64
}
