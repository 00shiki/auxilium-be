package users

import (
	"auxilium-be/infrastructure/repository/users"
)

type Controller struct {
	repo *users.Repository
}

func ControllerHandler(repo *users.Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}
