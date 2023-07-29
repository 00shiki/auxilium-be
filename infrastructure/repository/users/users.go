package users

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
