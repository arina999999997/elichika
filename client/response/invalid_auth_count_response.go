package response

type InvalidAuthCountResponse struct {
	AuthorizationCount int32 `json:"authorization_count"`
}
