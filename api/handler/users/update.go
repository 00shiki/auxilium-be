package users

import (
	"auxilium-be/entity/responses"
	USERS_ENTITY "auxilium-be/entity/users"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (handler *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	input := &USERS_ENTITY.Update{}
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

	if input.FirstName != "" {
		user.FirstName = input.FirstName
	}
	if input.LastName != "" {
		user.LastName = input.LastName
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.PhoneNumber != "" {
		user.PhoneNumber = input.PhoneNumber
	}
	if input.AvatarURL != "" {
		user.AvatarURL = input.AvatarURL
	}
	if input.Bio != "" {
		user.Bio = input.Bio
	}
	errUpdate := handler.ur.Update(&user)
	if errUpdate != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errUpdate.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusOK,
		Message: "update success",
	})
}

func (handler *Controller) ChangePassword(w http.ResponseWriter, r *http.Request) {
	input := &USERS_ENTITY.ChangePassword{}
	err := render.Bind(r, input)
	if err != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if input.Password != input.ConfirmPassword {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "password doesn't match",
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user, errDetail := handler.ur.DetailByEmail(input.Email)
	if errDetail != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errDetail.Error(),
		})
		return
	}

	user.Password = string(hashPassword)
	errUpdate := handler.ur.Update(&user)
	if errUpdate != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errUpdate.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusCreated,
		Message: "change password success",
	})
}
