package request

// TODO(refactor)
type AccessoryMeltRequest struct {
	UserAccessoryIds []int64 `json:"user_accessory_ids"`
}
