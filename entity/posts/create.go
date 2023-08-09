package posts

import (
	"errors"
	"net/http"
)

type Create struct {
	Anonymous bool   `json:"anonymous"`
	Body      string `json:"body"`
	ImageURL  string `json:"image_url"`
}

func (c Create) Bind(r *http.Request) error {
	if c.Body == "" {
		return errors.New("body must not empty")
	}
	return nil
}

type CreateComment struct {
	Anonymous bool   `json:"anonymous"`
	Body      string `form:"body"`
}

func (c CreateComment) Bind(r *http.Request) error {
	if c.Body == "" {
		return errors.New("body must not empty")
	}
	return nil
}
