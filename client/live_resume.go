package client

type LiveResume struct {
	LiveDifficultyId int32 `json:"live_difficulty_id"`
	DeckId           int32 `json:"deck_id"`
	ConsumedLp       int32 `json:"consumed_lp"`
}
