package request

type FetchGameServiceDataBeforeLoginRequest struct {
	UserId    int32  `json:"user_id"`
	ServiceId string `json:"service_id"`
	Mask      string `json:"mask"`
}
