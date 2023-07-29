package users

import (
	"auxilium-be/api/presentation/users"
	"auxilium-be/entity/responses"
	USERS_ENTITY "auxilium-be/entity/users"
	"auxilium-be/pkg/utils"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (handler *Controller) Login(w http.ResponseWriter, r *http.Request) {
	input := &USERS_ENTITY.Login{}
	err := render.Bind(r, input)
	if err != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user, err := handler.repo.DetailByEmail(input.Email)
	if err != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if errPass != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: errPass.Error(),
		})
		return
	}

	exp := time.Hour * 24 * 7
	expRef := time.Hour * 24 * 7 * 4
	accessToken, errToken := utils.CreateToken(user.ID, user.Role, exp)
	if errToken != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errToken.Error(),
		})
		return
	}
	refreshToken, errRefToken := utils.CreateToken(user.ID, user.Role, expRef)
	if errRefToken != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errRefToken.Error(),
		})
		return
	}
	errStore := handler.repo.StoreToken(user.ID, accessToken, refreshToken)
	if errStore != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errStore.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Message: "login success",
		Code:    http.StatusOK,
		Data: users.ResponseLogin{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}
