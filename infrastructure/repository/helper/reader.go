package helper

import (
	"auxilium-be/entity/helper"
)

func (r *Repository) List(lat float64, lon float64, radius float64) ([]helper.Helper, error) {
	var helpers []helper.Helper
	result := r.db.Raw("SELECT * FROM helpers WHERE lat BETWEEN ? AND ? AND lon BETWEEN ? AND ?", lat+radius, lat-radius, lon+radius, lon-radius).Scan(&helpers)
	if result.Error != nil {
		return []helper.Helper{}, result.Error
	}
	return helpers, nil
}
