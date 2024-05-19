package user_love_ranking

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/generic"
	"elichika/subsystem/user_social"
	"elichika/userdata"

	"sort"
)

func friendRanking(session *userdata.Session, condition int32, rankingOrder generic.Nullable[int32]) response.FetchLoveRankingResponse {
	resp := response.FetchLoveRankingResponse{}
	userIds := user_social.GetFriendUserIds(session)
	userIds = append(userIds, session.UserId)

	for _, userId := range userIds {
		resp.LoveRankingData.Append(client.LoveRankingData{
			RankingUser: user_social.GetRankingUser(session, userId),
			LovePoint:   GetUserTotalLovePoint(session, userId, condition),
		})
	}
	sort.Slice(resp.LoveRankingData.Slice, func(i, j int) bool {
		return resp.LoveRankingData.Slice[i].LovePoint < resp.LoveRankingData.Slice[j].LovePoint
	})
	for i := range resp.LoveRankingData.Slice {
		if (i == 0) || (resp.LoveRankingData.Slice[i].LovePoint != resp.LoveRankingData.Slice[i-1].LovePoint) {
			resp.LoveRankingData.Slice[i].Order = int32(i) + 1
		} else {
			resp.LoveRankingData.Slice[i].Order = resp.LoveRankingData.Slice[i-1].Order
		}
		if resp.LoveRankingData.Slice[i].RankingUser.UserId == session.UserId {
			resp.MyRankingOrder = generic.NewNullable(resp.LoveRankingData.Slice[i].Order)
		}
	}
	return resp
}
