package response

import (
	"elichika/client"
	"elichika/generic"
)

type GetVoltageRankingResponse struct {
	VoltageRankingCells generic.List[client.VoltageRankingCell] `json:"voltage_ranking_cells"`
}
