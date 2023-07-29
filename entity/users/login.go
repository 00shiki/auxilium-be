package users

import (
	"errors"
	"net/http"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l Login) Bind(r *http.Request) error {
	if l.Email == "" {
		return errors.New("email must not empty")
	}
	if l.Password == "" {
		return errors.New("password must not empty")
	}
	return nil
}
