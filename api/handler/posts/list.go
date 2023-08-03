package posts

import (
	POSTS_PRESENTATION "auxilium-be/api/presentation/posts"
	"auxilium-be/entity/responses"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"time"
)

func (handler *Controller) ListPosts(w http.ResponseWriter, r *http.Request) {
	pageQuery := r.URL.Query().Get("page")
	sizeQuery := r.URL.Query().Get("size")
	if pageQuery == "" {
		pageQuery = "0"
	}
	if sizeQuery == "" {
		sizeQuery = "0"
	}
	page, errPage := strconv.Atoi(pageQuery)
	if errPage != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errPage.Error(),
		})
		return
	}
	size, errSize := strconv.Atoi(sizeQuery)
	if errSize != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errSize.Error(),
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

	posts, errList := handler.pr.ListPosts(page, size)
	if errList != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errList.Error(),
		})
		return
	}
	list := []POSTS_PRESENTATION.ResponseListPosts{}
	for _, post := range posts {
		list = append(list, POSTS_PRESENTATION.ResponseListPosts{
			ID:            post.ID,
			Username:      post.Username,
			AvatarURL:     post.AvatarURL,
			Body:          post.Body,
			ImageURL:      post.ImageURL,
			CommentsCount: post.CommentsCount,
			LikesCount:    post.LikesCount,
		})
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusOK,
		Message: "list posts success",
		Data:    list,
	})
}
