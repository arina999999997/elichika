package user_love_ranking

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

// by network capture from official server:
// - first, if ranking order is not given, it's set to this player's own rank
// - then the begining position is set to max(1, ranking order - 124)
// - the end position is set to beginning + 250 - 1
// - the 250 relevant users are the returned
func FetchLoveRanking(session *userdata.Session, rankingType, condition int32, rankingOrder generic.Nullable[int32]) response.FetchLoveRankingResponse {
	if rankingType == enum.LoveRankingTypeAll {
		return commonRanking(session, condition, rankingOrder)
	} else {
		return friendRanking(session, condition, rankingOrder)
	}
}
