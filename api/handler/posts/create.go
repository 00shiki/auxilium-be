package posts

import (
	"auxilium-be/entity/jwt"
	POSTS_ENTITY "auxilium-be/entity/posts"
	"auxilium-be/entity/responses"
	"auxilium-be/entity/users"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func (handler *Controller) CreatePost(w http.ResponseWriter, r *http.Request) {
	input := &POSTS_ENTITY.Create{}
	err := render.Bind(r, input)
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
			Message: errClaims.Error(),
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

	payload := claims["sub"].(jwt.Payload)
	user, errDetail := handler.ur.DetailById(payload.UserID)
	if errDetail != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusNotFound,
			Message: errDetail.Error(),
		})
		return
	}

	imageURL := ""
	//tunggu aws nya
	//if input.Image != nil {
	//
	//}
	if input.Anonymous {
		user = users.User{}
	}
	post := &POSTS_ENTITY.Post{
		UserID:   0,
		User:     user,
		Body:     input.Body,
		ImageURL: imageURL,
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
