package request

type LoginRequest struct {
	UserId         int32  `json:"user_id"`
	AuthCount      int32  `json:"auth_count"`
	Mask           string `json:"mask"`
	AssetState     string `json:"asset_state"`
	RecaptchaToken string `json:"recaptcha_token"`
}
