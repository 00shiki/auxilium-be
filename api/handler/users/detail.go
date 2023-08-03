package users

import (
	POSTS_PRESENTATION "auxilium-be/api/presentation/posts"
	"auxilium-be/api/presentation/users"
	"auxilium-be/entity/responses"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func (handler *Controller) DetailUser(w http.ResponseWriter, r *http.Request) {
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

	userID := claims["id"].(float64)
	user, errDetail := handler.ur.DetailByID(uint(userID))
	if errDetail != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errDetail.Error(),
		})
		return
	}

	posts, errPosts := handler.pr.ListPostsByUserID(uint(userID))
	if errPosts != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errPosts.Error(),
		})
		return
	}

	userPosts := []POSTS_PRESENTATION.ResponseListPosts{}
	for _, post := range posts {
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
