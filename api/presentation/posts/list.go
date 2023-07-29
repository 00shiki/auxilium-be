package posts

type ResponseListPosts struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	AvatarURL     string `json:"avatar_url"`
	Body          string `json:"body"`
	ImageURL      string `json:"image_url"`
	CommentsCount int    `json:"comments_count"`
	LikesCount    int    `json:"likes_count"`
}
