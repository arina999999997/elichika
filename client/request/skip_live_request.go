package request

import (
	"elichika/client"
	"elichika/generic"
)

type SkipLiveRequest struct {
	LiveDifficultyMasterId  int32                                            `json:"live_difficulty_master_id"`
	DeckId                  int32                                            `json:"deck_id"`
	TicketUseCount          int32                                            `json:"ticket_use_count"`
	CellId                  generic.Nullable[int32]                          `json:"cell_id"`
	LiveEventMarathonStatus generic.Nullable[client.LiveEventMarathonStatus] `json:"live_event_marathon_status"`
}
