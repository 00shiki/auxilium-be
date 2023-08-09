package helper

import (
	"errors"
	"net/http"
)

type Create struct {
	Lat float64
	Lon float64
}

func (c Create) Bind(r *http.Request) error {
	if c.Lat == 0 {
		return errors.New("lat must no empty")
	}
	if c.Lon == 0 {
		return errors.New("lon must no empty")
	}
	return nil
}
