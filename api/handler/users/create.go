package users

import (
	"auxilium-be/entity/responses"
	USERS_ENTITY "auxilium-be/entity/users"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (handler *Controller) CreateUsers(w http.ResponseWriter, r *http.Request) {
	input := &USERS_ENTITY.Create{}
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

	data := &USERS_ENTITY.User{
		Username:    input.Username,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Password:    string(hashPassword),
		PhoneNumber: input.PhoneNumber,
	}

	errCreate := handler.repo.Create(data)
	if errCreate != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: "error save to database",
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusCreated,
		Message: "register success",
	})
}
