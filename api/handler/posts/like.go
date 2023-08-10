package posts

import (
	"auxilium-be/entity/responses"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"time"
)

func (handler *Controller) LikePost(w http.ResponseWriter, r *http.Request) {
	postIDParams := chi.URLParam(r, "postID")
	if postIDParams == "" {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "postID must not empty",
		})
		return
	}
	postID, errParse := strconv.Atoi(postIDParams)
	if errParse != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "postID must be integer",
		})
		return
	}

	_, claims, errClaims := jwtauth.FromContext(r.Context())
	if errClaims != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: errClaims.Error(),
		})
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

	errLike := handler.pr.LikePost(uint(postID))
	if errLike != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errLike.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusOK,
		Message: "like post success",
	})
}

func (handler *Controller) LikeComment(w http.ResponseWriter, r *http.Request) {
	commentIDParams := chi.URLParam(r, "commentID")
	if commentIDParams == "" {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "postID must not empty",
		})
		return
	}
	commentID, errParse := strconv.Atoi(commentIDParams)
	if errParse != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "postID must be integer",
		})
		return
	}

	_, claims, errClaims := jwtauth.FromContext(r.Context())
	if errClaims != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: errClaims.Error(),
		})
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

	errLike := handler.pr.LikeComment(uint(commentID))
	if errLike != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errLike.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusOK,
		Message: "like post success",
	})
}
