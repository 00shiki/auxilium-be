package posts

import (
	"auxilium-be/api/presentation/posts"
	"auxilium-be/entity/responses"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"time"
)

func (handler *Controller) DetailPost(w http.ResponseWriter, r *http.Request) {
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

	postIDString := chi.URLParam(r, "postID")
	postID, errPostID := strconv.Atoi(postIDString)
	if errPostID != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errPostID.Error(),
		})
		return
	}

	post, errPost := handler.pr.DetailByID(uint(postID))
	if errPost != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusNotFound,
			Message: errPost.Error(),
		})
		return
	}

	comments, errComments := handler.pr.ListComments(uint(postID))
	if errComments != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errComments.Error(),
		})
		return
	}

	detailComments := []posts.ResponseListComments{}
	for _, comment := range comments {
		detailComments = append(detailComments, posts.ResponseListComments{
			ID:         comment.ID,
			Username:   comment.User.Username,
			AvatarURL:  comment.User.AvatarURL,
			Body:       comment.Body,
			LikesCount: comment.LikesCount,
		})
	}

	detail := posts.ResponseDetailPost{
		ID:            uint(postID),
		Username:      post.User.Username,
		AvatarURL:     post.User.AvatarURL,
		Body:          post.Body,
		ImageURL:      post.ImageURL,
		CommentsCount: post.CommentsCount,
		LikesCount:    post.LikesCount,
		Comments:      detailComments,
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusOK,
		Message: "detail get success",
		Data:    detail,
	})
}
