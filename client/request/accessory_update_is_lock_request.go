package request

type AccessoryUpdateIsLockRequest struct {
	UserAccessoryId int64 `json:"user_accessory_id"`
	IsLock          bool  `json:"is_lock"`
}
