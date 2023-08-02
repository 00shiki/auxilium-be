package users

import "auxilium-be/entity/users"

func (r *Repository) DetailByID(userID uint) (users.User, error) {
	var user users.User
	result := r.db.First(&user, userID)
	if result.Error != nil {
		return users.User{}, result.Error
	}
	return user, nil
}

func (r *Repository) DetailByEmail(email string) (users.User, error) {
	var user users.User
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return users.User{}, result.Error
	}
	return user, nil
}
