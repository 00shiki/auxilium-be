package posts

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID        uint
	Username      string
	AvatarURL     string
	Body          string
	ImageURL      string
	Comments      []Comment
	CommentsCount int
	LikesCount    int
}

type Comment struct {
	gorm.Model
	UserID     uint
	Username   string
	AvatarURL  string
	PostID     uint
	Body       string
	LikesCount int
}
