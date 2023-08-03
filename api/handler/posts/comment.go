package posts

import (
	POSTS_ENTITY "auxilium-be/entity/posts"
	"auxilium-be/entity/responses"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"time"
)

func (handler *Controller) CreateComment(w http.ResponseWriter, r *http.Request) {
	input := &POSTS_ENTITY.CreateComment{}
	err := render.Bind(r, input)
	if err != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	postIDString := chi.URLParam(r, "postID")
	postID, errPostID := strconv.Atoi(postIDString)
	if errPostID != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errPostID.Error(),
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

	userID := claims["id"].(float64)
	user, errDetail := handler.ur.DetailByID(uint(userID))
	if errDetail != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusNotFound,
			Message: errDetail.Error(),
		})
		return
	}

	comment := POSTS_ENTITY.Comment{
		UserID:    user.ID,
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
		PostID:    uint(postID),
		Body:      input.Body,
	}

	errComment := handler.pr.Comment(&comment)
	if errComment != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errComment.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusCreated,
		Message: "comment created successfully",
	})
}
