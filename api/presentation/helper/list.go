package helper

type ResponseListHelpers struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	AvatarURL string  `json:"avatar_url"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
}
