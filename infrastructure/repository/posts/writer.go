package posts

import (
	"auxilium-be/entity/posts"
	"gorm.io/gorm"
)

func (r *Repository) Create(post *posts.Post) error {
	result := r.db.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) Comment(comment *posts.Comment) error {
	result := r.db.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) Update(post *posts.Post) error {
	result := r.db.Save(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) LikePost(postID uint) error {
	result := r.db.Exec("UPDATE posts SET likes_count = ? WHERE id = ?", gorm.Expr("likes_count + ?", 1), postID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) LikeComment(commentID uint) error {
	result := r.db.Exec("UPDATE comments SET likes_count = ? WHERE id = ?", gorm.Expr("likes_count + ?", 1), commentID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) DislikePost(postID uint) error {
	result := r.db.Exec("UPDATE posts SET likes_count = ? WHERE id = ?", gorm.Expr("likes_count - ?", 1), postID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) DislikeComment(commentID uint) error {
	result := r.db.Exec("UPDATE comments SET likes_count = ? WHERE id = ?", gorm.Expr("likes_count - ?", 1), commentID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
