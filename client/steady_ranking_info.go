package client

import (
	"elichika/generic"
)

type SteadyRankingInfo struct {
	SteadyRankingId              int32                `json:"steady_ranking_id"`
	StartAt                      int64                `json:"start_at"`
	FinishAt                     int64                `json:"finish_at"`
	OpenAt                       int64                `json:"open_at"`
	CloseAt                      int64                `json:"close_at"`
	LiveDifficultyMasterIdList   generic.Array[int32] `json:"live_difficulty_master_id_list"`
	SelectLiveDifficultyMasterId int32                `json:"select_live_difficulty_master_id"`
}
