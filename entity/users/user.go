package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string   `gorm:"size:255;not null;unique" json:"username"`
	FirstName    string   `gorm:"size:255;not null" json:"first_name"`
	LastName     string   `gorm:"size:255;not null" json:"last_name"`
	Email        string   `gorm:"size:255;not null;unique" json:"email"`
	Password     string   `gorm:"size:255;not null" json:"password"`
	PhoneNumber  string   `gorm:"size:255;not null" json:"phone_number"`
	AvatarUrl    string   `gorm:"size:255" json:"avatar_url"`
	Location     Location `gorm:"embedded"`
	AccessToken  string   `gorm:"size:255" json:"access_token"`
	RefreshToken string   `gorm:"size:255" json:"refresh_token"`
	Role         int
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
