package helper

import (
	"auxilium-be/infrastructure/repository/helper"
	"auxilium-be/infrastructure/repository/users"
)

type Controller struct {
	hr *helper.Repository
	ur *users.Repository
}

func ControllerHandler(hr *helper.Repository, ur *users.Repository) *Controller {
	return &Controller{
		hr: hr,
		ur: ur,
	}
}
