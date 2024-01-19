package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchTowerRankingResponse struct {
	TopRankingCells    generic.List[client.TowerRankingCell]       `json:"top_ranking_cells"`
	MyRankingCells     generic.List[client.TowerRankingCell]       `json:"my_ranking_cells"`
	FriendRankingCells generic.List[client.TowerRankingCell]       `json:"friend_ranking_cells"`
	RankingBorderInfo  generic.List[client.TowerRankingBorderInfo] `json:"ranking_border_info"`
	MyOrder            generic.Nullable[int32]                     `json:"my_order"`
}
