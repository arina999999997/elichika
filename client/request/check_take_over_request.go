package request

type CheckTakeOverRequest struct {
	UserId     int32  `json:"user_id"`
	TakeOverId string `json:"take_over_id"`
	PassWord   string `json:"pass_word"`
	Mask       string `json:"mask"`
}
