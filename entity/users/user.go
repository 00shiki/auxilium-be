package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string   `gorm:"size:255;not null;unique"`
	FirstName    string   `gorm:"size:255;not null"`
	LastName     string   `gorm:"size:255;not null"`
	Email        string   `gorm:"size:255;not null;unique"`
	Password     string   `gorm:"size:255;not null"`
	PhoneNumber  string   `gorm:"size:255;not null"`
	AvatarURL    string   `gorm:"size:255"`
	Location     Location `gorm:"embedded"`
	AccessToken  string   `gorm:"size:255"`
	RefreshToken string   `gorm:"size:255"`
	Role         int
}

type Location struct {
	Lat float64
	Lon float64
}
