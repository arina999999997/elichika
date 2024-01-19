package client

type CurrentUserData struct {
	UserId      int32         `json:"user_id"`
	Name        LocalizedText `json:"name"`
	LastLoginAt int64         `json:"last_login_at"`
	SnsCoin     int32         `json:"sns_coin"`
}
