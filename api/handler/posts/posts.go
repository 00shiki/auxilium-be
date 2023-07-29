package posts

import (
	"auxilium-be/infrastructure/repository/posts"
	"auxilium-be/infrastructure/repository/users"
)

type Controller struct {
	pr *posts.Repository
	ur *users.Repository
}

func ControllerHandler(pr *posts.Repository, ur *users.Repository) *Controller {
	return &Controller{
		pr: pr,
		ur: ur,
	}
}
