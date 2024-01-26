package request

type GetVoltageRankingDeckRequest struct {
	LiveDifficultyId int32 `json:"live_difficulty_id"`
	UserId           int32 `json:"user_id"`
}
