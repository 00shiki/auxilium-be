package users

import "auxilium-be/api/presentation/posts"

type ResponseDetailUser struct {
	Username    string                    `json:"username"`
	FirstName   string                    `json:"first_name"`
	LastName    string                    `json:"last_name"`
	Email       string                    `json:"email"`
	PhoneNumber string                    `json:"phone_number"`
	AvatarURL   string                    `json:"avatar_url"`
	Bio         string                    `json:"bio"`
	Posts       []posts.ResponseListPosts `json:"posts"`
}
