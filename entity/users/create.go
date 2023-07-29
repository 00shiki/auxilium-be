package users

import (
	"errors"
	"net/http"
)

type Create struct {
	Username        string `json:"username"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	PhoneNumber     string `json:"phone_number"`
}

func (c Create) Bind(r *http.Request) error {
	if c.Username == "" {
		return errors.New("username must not empty")
	}
	if c.FirstName == "" {
		return errors.New("first_name must not empty")
	}
	if c.LastName == "" {
		return errors.New("last_name must not empty")
	}
	if c.Email == "" {
		return errors.New("email must not empty")
	}
	if c.Password == "" {
		return errors.New("password must not empty")
	}
	if c.ConfirmPassword == "" {
		return errors.New("confirm_password must not empty")
	}
	if c.PhoneNumber == "" {
		return errors.New("phone_number must not empty")
	}
	return nil
}
