package helper

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewHelperRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
