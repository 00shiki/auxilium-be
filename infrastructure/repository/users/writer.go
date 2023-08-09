package users

import "auxilium-be/entity/users"

func (r *Repository) Create(user *users.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) StoreToken(userId uint, accessToken string, refreshToken string) error {
	var user users.User
	result := r.db.First(&user, userId)
	if result.Error != nil {
		return result.Error
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	result = r.db.Save(user)
	return nil
}

func (r *Repository) Update(user *users.User) error {
	result := r.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
