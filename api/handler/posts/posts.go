package posts

import "auxilium-be/infrastructure/repository/posts"

type Controller struct {
	repo *posts.Repository
}

func ControllerHandler(repo *posts.Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}
