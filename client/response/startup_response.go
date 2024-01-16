package response

type StartupResponse struct {
	UserId           int32  `json:"user_id"`
	AuthorizationKey string `json:"authorization_key"`
}
