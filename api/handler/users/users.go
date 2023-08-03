package users

import (
	"auxilium-be/infrastructure/repository/posts"
	"auxilium-be/infrastructure/repository/users"
)

type Controller struct {
	ur *users.Repository
	pr *posts.Repository
}

func ControllerHandler(ur *users.Repository, pr *posts.Repository) *Controller {
	return &Controller{
		ur: ur,
		pr: pr,
	}
}
