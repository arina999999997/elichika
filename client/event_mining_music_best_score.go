package client

type EventMiningMusicBestScore struct {
	LiveId           int32 `json:"live_id"`
	FrameNo          int32 `json:"frame_no"`
	LiveDifficultyId int32 `json:"live_difficulty_id"`
	Score            int32 `json:"score"`
}
