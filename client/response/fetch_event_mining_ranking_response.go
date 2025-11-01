package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchEventMiningRankingResponse struct {
	PointTopRankingCells      generic.List[client.EventMiningRankingCell]       `json:"point_top_ranking_cells"`
	PointMyRankingCells       generic.List[client.EventMiningRankingCell]       `json:"point_my_ranking_cells"`
	PointFriendRankingCells   generic.List[client.EventMiningRankingCell]       `json:"point_friend_ranking_cells"`
	VoltageTopRankingCells    generic.List[client.EventMiningRankingCell]       `json:"voltage_top_ranking_cells"`
	VoltageMyRankingCells     generic.List[client.EventMiningRankingCell]       `json:"voltage_my_ranking_cells"`
	VoltageFriendRankingCells generic.List[client.EventMiningRankingCell]       `json:"voltage_friend_ranking_cells"`
	PointRankingBorderInfo    generic.List[client.EventMiningRankingBorderInfo] `json:"point_ranking_border_info"`
	VoltageRankingBorderInfo  generic.List[client.EventMiningRankingBorderInfo] `json:"voltage_ranking_border_info"`
}
