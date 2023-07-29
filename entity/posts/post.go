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
	CommentsCount int
	LikesCount    int
}

type Comment struct {
	gorm.Model
	UserID     uint
	User       users.User
	PostID     uint
	Body       string
	LikesCount int
}
