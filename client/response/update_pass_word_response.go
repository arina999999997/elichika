package response

type UpdatePassWordResponse struct {
	TakeOverId string `json:"take_over_id"` // this field is actually named _takeOverId for some reason
}
