package posts

import POSTS_ENTITY "auxilium-be/entity/posts"

func (r *Repository) ListPosts(page int, size int) ([]POSTS_ENTITY.Post, error) {
	var posts []POSTS_ENTITY.Post
	if page <= 0 {
		page = 1
	}
	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}
	offset := (page - 1) * size
	pagination := r.db.Offset(offset).Limit(size).Find(&posts)
	if pagination.Error != nil {
		return []POSTS_ENTITY.Post{}, pagination.Error
	}
	return posts, nil
}

func (r *Repository) ListComments(postID uint) ([]POSTS_ENTITY.Comment, error) {
	var comments []POSTS_ENTITY.Comment
	result := r.db.Find(&comments, "post_id = ?", postID)
	if result.Error != nil {
		return []POSTS_ENTITY.Comment{}, result.Error
	}
	return comments, nil
}

func (r *Repository) DetailByID(postID uint) (POSTS_ENTITY.Post, error) {
	var post POSTS_ENTITY.Post
	result := r.db.First(&post, postID)
	if result.Error != nil {
		return POSTS_ENTITY.Post{}, result.Error
	}
	return post, nil
}
