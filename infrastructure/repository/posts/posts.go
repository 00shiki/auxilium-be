package posts

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
