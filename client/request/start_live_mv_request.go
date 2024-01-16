package request

import (
	"elichika/generic"
)

type StartLiveMvRequest struct {
	LiveMasterId   int32                   `json:"live_master_id"`
	LiveMvDeckType generic.Nullable[int32] `json:"live_mv_deck_type" enum:"LiveMvDeckType"`
	StageMasterId  generic.Nullable[int32] `json:"stage_master_id"`
}
