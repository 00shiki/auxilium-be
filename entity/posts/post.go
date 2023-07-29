package posts

import (
	"auxilium-be/entity/users"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID        uint
	User          users.User
	Body          string
	ImageURL      string
	Comments      []Comment
	CommentsCount int64
	LikesCount    int64
}

type Comment struct {
	gorm.Model
	UserID     uint
	User       users.User
	Body       string
	LikesCount int64
}
