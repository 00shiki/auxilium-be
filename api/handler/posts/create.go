package posts

import (
	POSTS_ENTITY "auxilium-be/entity/posts"
	"auxilium-be/entity/responses"
	"fmt"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func (handler *Controller) CreatePost(w http.ResponseWriter, r *http.Request) {
	input := POSTS_ENTITY.Create{}
	err := render.Bind(r, &input)
	if err != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	_, claims, errClaims := jwtauth.FromContext(r.Context())
	if errClaims != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: fmt.Sprintf("claims: %v", errClaims.Error()),
		})
		return
	}

	now := time.Now()
	exp := claims["exp"].(time.Time)
	if exp.Unix() < now.Unix() {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "token expired",
		})
		return
	}

	userID := claims["id"].(float64)
	user, errDetail := handler.ur.DetailByID(uint(userID))
	if errDetail != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusNotFound,
			Message: errDetail.Error(),
		})
		return
	}

	if input.Anonymous {
		user.ID = 1
		user.Username = ""
		user.AvatarURL = ""
	}
	post := &POSTS_ENTITY.Post{
		UserID:    user.ID,
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
		Body:      input.Body,
		ImageURL:  input.ImageURL,
	}
	errCreate := handler.pr.Create(post)
	if errCreate != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errCreate.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusCreated,
		Message: "post created successfully",
	})
}
