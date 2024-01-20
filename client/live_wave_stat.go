package client

type LiveWaveStat struct {
	WaveId       int32 `json:"wave_id"`
	IsClear      bool  `json:"is_clear"`
	MissionCount int64 `json:"mission_count"`
}
