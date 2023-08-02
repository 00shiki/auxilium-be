package posts

import (
	POSTS_ENTITY "auxilium-be/entity/posts"
	"auxilium-be/entity/responses"
	"auxilium-be/entity/users"
	"fmt"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func (handler *Controller) CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("parse: %v", err.Error()),
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

	imageURL := ""
	//tunggu aws nya
	//if input.Image != nil {
	//
	//}
	//image, header, errImage := r.FormFile("image")

	anonymous := r.FormValue("anonymous")
	if anonymous == "1" {
		user = users.User{}
	}
	body := r.FormValue("body")
	post := &POSTS_ENTITY.Post{
		UserID:   user.ID,
		User:     user,
		Body:     body,
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
