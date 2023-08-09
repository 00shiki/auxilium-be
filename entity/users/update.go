package users

import "net/http"

type Update struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	AvatarURL   string `json:"avatar_url"`
}

func (u Update) Bind(r *http.Request) error {
	return nil
}
