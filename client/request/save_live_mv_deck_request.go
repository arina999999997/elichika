package request

import (
	"elichika/generic"
)

type SaveLiveMvDeckRequest struct {
	LiveMasterId        int32                                              `json:"live_master_id"`
	LiveMvDeckType      int32                                              `json:"live_mv_deck_type" enum:"LiveMvDeckType"`
	MemberMasterIdByPos generic.Dictionary[int32, int32]                   `json:"member_master_id_by_pos"`
	SuitMasterIdByPos   generic.Dictionary[int32, generic.Nullable[int32]] `json:"suit_master_id_by_pos"`
	ViewStatusByPos     generic.Dictionary[int32, int32]                   `json:"view_status_by_pos" enum:"MemberViewStatus"`
}
