package posts

import "auxilium-be/entity/posts"

func (r *Repository) Create(post *posts.Post) error {
	result := r.db.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
