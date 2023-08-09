package helper

import (
	HELPER_PRESENTATION "auxilium-be/api/presentation/helper"
	"auxilium-be/entity/responses"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"time"
)

func (handler *Controller) List(w http.ResponseWriter, r *http.Request) {
	radiusQuery := r.URL.Query().Get("radius")
	latQuery := r.URL.Query().Get("lat")
	lonQuery := r.URL.Query().Get("lon")
	if radiusQuery == "" {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "params radius must not empty",
		})
		return
	}
	if latQuery == "" {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "params lat must not empty",
		})
		return
	}
	if lonQuery == "" {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: "params lon must not empty",
		})
		return
	}
	radius, errRadius := strconv.ParseFloat(radiusQuery, 64)
	if errRadius != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errRadius.Error(),
		})
		return
	}
	lat, errLat := strconv.ParseFloat(latQuery, 64)
	if errLat != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errLat.Error(),
		})
		return
	}
	lon, errLon := strconv.ParseFloat(lonQuery, 64)
	if errLon != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusBadRequest,
			Message: errLon.Error(),
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

	helpers, errList := handler.hr.List(lat, lon, radius)
	if errList != nil {
		render.Render(w, r, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: errList.Error(),
		})
		return
	}
	list := []HELPER_PRESENTATION.ResponseListHelpers{}
	for _, helper := range helpers {
		list = append(list, HELPER_PRESENTATION.ResponseListHelpers{
			ID:        helper.ID,
			Username:  helper.Username,
			AvatarURL: helper.AvatarURL,
			Lat:       helper.Lat,
			Lon:       helper.Lon,
		})
	}

	render.Render(w, r, &responses.Response{
		Code:    http.StatusOK,
		Message: "list helper success",
		Data:    list,
	})
}
