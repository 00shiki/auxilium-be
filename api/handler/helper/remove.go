package helper

import (
	"auxilium-be/entity/responses"
	"auxilium-be/entity/users"
	"fmt"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func (handler *Controller) RemoveHelper(w http.ResponseWriter, r *http.Request) {
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
		Lat: 0,
		Lon: 0,
	}
	errUpdate := handler.ur.Update(&user)
	if errUpdate != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errUpdate.Error(),
		})
		return
	}

	errRemove := handler.hr.Remove(uint(userID))
	if errRemove != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errRemove.Error(),
		})
		return
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusOK,
		Message: "helper removed successfully",
	})
}
