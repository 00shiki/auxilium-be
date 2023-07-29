package users

import "auxilium-be/entity/users"

func (r *Repository) DetailById(userId int) (users.User, error) {
	var user users.User
	result := r.db.First(&user, userId)
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
