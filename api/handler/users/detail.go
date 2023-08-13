package users

import (
	POSTS_PRESENTATION "auxilium-be/api/presentation/posts"
	"auxilium-be/api/presentation/users"
	"auxilium-be/entity/responses"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func (handler *Controller) DetailUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "username must not empty",
		})
		return
	}
	user, errDetail := handler.ur.DetailByUsername(username)
	if errDetail != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errDetail.Error(),
		})
		return
	}

	posts, errPosts := handler.pr.ListPostsByUserID(user.ID)
	if errPosts != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errPosts.Error(),
		})
		return
	}

	userPosts := []POSTS_PRESENTATION.ResponseListPosts{}
	for _, post := range posts {
		if post.Username == "" {
			continue
		}
		userPosts = append(userPosts, POSTS_PRESENTATION.ResponseListPosts{
			ID:            post.ID,
			Username:      user.Username,
			AvatarURL:     user.AvatarURL,
			Body:          post.Body,
			ImageURL:      post.ImageURL,
			CommentsCount: post.CommentsCount,
			LikesCount:    post.LikesCount,
		})
	}

	detail := users.ResponseDetailUser{
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		AvatarURL:   user.AvatarURL,
		Posts:       userPosts,
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusOK,
		Message: "detail get success",
		Data:    detail,
	})
}
