package request

type SetTakeOverRequest struct {
	TakeOverId string `json:"take_over_id"`
	PassWord   string `json:"pass_word"`
	Mask       string `json:"mask"`
}
