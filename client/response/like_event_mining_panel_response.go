package response

import (
	"elichika/client"
	"elichika/generic"
)

type LikeEventMiningPanelResponse struct {
	EventMiningTopCellStateList generic.List[client.EventMiningTopCellState] `json:"event_mining_top_cell_state_list"`
}
