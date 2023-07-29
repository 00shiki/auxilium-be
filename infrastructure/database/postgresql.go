package database

import (
	"auxilium-be/entity/posts"
	"auxilium-be/entity/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&users.User{}, &posts.Post{}, &posts.Comment{})
	return db, nil
}
