package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchEventMarathonRankingResponse struct {
	TopRankingCells    generic.List[client.EventMarathonRankingCell]       `json:"top_ranking_cells"`
	MyRankingCells     generic.List[client.EventMarathonRankingCell]       `json:"my_ranking_cells"`
	FriendRankingCells generic.List[client.EventMarathonRankingCell]       `json:"friend_ranking_cells"`
	RankingBorderInfo  generic.List[client.EventMarathonRankingBorderInfo] `json:"ranking_border_info"`
}
