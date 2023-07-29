package responses

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	Code    int         `json:"-"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (res Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, res.Code)
	return nil
}
