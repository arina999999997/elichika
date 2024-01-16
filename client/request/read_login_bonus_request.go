package request

type ReadLoginBonusRequest struct {
	LoginBonusType int32 `json:"login_bonus_type" enum:"LoginBonusType"`
	LoginBonusId   int32 `json:"login_bonus_id"`
}
