package request

import (
	"elichika/client"
	"elichika/generic"
)

type StartLiveRequest struct {
	LiveDifficultyId        int32                                            `json:"live_difficulty_id"`
	DeckId                  int32                                            `json:"deck_id"`
	CellId                  generic.Nullable[int32]                          `json:"cell_id"`
	PartnerUserId           int32                                            `json:"partner_user_id"`
	PartnerCardMasterId     int32                                            `json:"partner_card_master_id"`
	LpMagnification         int32                                            `json:"lp_magnification"`
	IsAutoPlay              bool                                             `json:"is_auto_play"`
	LiveEventMarathonStatus generic.Nullable[client.LiveEventMarathonStatus] `json:"live_event_marathon_status"` // pointer
	LiveTowerStatus         generic.Nullable[client.LiveTowerStatus]         `json:"live_tower_status"`          // pointer
	IsReferenceBook         bool                                             `json:"is_reference_book"`
}
