package helper

import "auxilium-be/entity/helper"

func (r *Repository) Create(helper *helper.Helper) error {
	result := r.db.Create(&helper)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) Remove(userID uint) error {
	result := r.db.Exec("DELETE helpers WHERE user_id = ?", userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
