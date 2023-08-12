package users

type ResponseLogin struct {
	Username     string `json:"username"`
	AvatarURL    string `json:"avatar_url"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
