package users

import (
	"errors"
	"net/http"
)

type Update struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	AvatarURL   string `json:"avatar_url"`
	Bio         string `json:"bio"`
}

func (u Update) Bind(r *http.Request) error {
	return nil
}

type ChangePassword struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (c ChangePassword) Bind(r *http.Request) error {
	if c.Email == "" {
		return errors.New("email must not empty")
	}
	if c.Password == "" {
		return errors.New("password must not empty")
	}
	if c.ConfirmPassword == "" {
		return errors.New("confirm_password must not empty")
	}
	return nil
}
