package helper

import (
	HELPER_ENTITY "auxilium-be/entity/helper"
	"auxilium-be/entity/responses"
	"auxilium-be/entity/users"
	"fmt"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func (handler *Controller) CreateHelper(w http.ResponseWriter, r *http.Request) {
	input := HELPER_ENTITY.Create{}
	err := render.Bind(r, &input)
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

	user.Location = users.Location{
		Lat: input.Lat,
		Lon: input.Lon,
	}
	errUpdate := handler.ur.Update(&user)
	if errUpdate != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errUpdate.Error(),
		})
		return
	}

	helper := HELPER_ENTITY.Helper{
		UserID:    uint(userID),
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
		Lat:       input.Lat,
		Lon:       input.Lon,
	}
	errCreate := handler.hr.Create(&helper)
	if errCreate != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errCreate.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusCreated,
		Message: "helper created successfully",
	})
}
