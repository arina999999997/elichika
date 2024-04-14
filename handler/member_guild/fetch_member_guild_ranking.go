package member_guild

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(member_guild): the logic of this part is wrong or missing

func fetchMemberGuildRanking(ctx *gin.Context) {
	// There is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.FetchMemberGuildRankingResponse{}
	resp.MemberGuildRanking.ViewYear = 2024
	// resp.MemberGuildRanking.NextYear = 2023
	// resp.MemberGuildRanking.PreviousYear = 2021
	oneTerm := client.MemberGuildRankingOneTerm{
		MemberGuildId: 1,
		StartAt:       1,
		EndAt:         1,
	}

	rank := int32(0)
	for _, member := range session.Gamedata.Member {
		rank++
		oneTerm.Channels.Append(client.MemberGuildRankingOneTermCell{
			Order:          rank,
			TotalPoint:     1000000,
			MemberMasterId: member.Id,
		})
	}

	resp.MemberGuildRanking.MemberGuildRankingList.Append(oneTerm)

	mgur := client.MemberGuildUserRanking{
		MemberGuildId: 1,
	}
	userData := client.MemberGuildUserRankingUserData{
		UserId:                 int32(session.UserId),
		UserName:               session.UserStatus.Name,
		UserRank:               session.UserStatus.Rank,
		CardMasterId:           session.UserStatus.RecommendCardMasterId,
		Level:                  80,
		IsAwakening:            true,
		IsAllTrainingActivated: true,
		EmblemMasterId:         session.UserStatus.EmblemId,
	}
	userRankingCell := client.MemberGuildUserRankingCell{
		Order:                          generic.NewNullable(int32(1)),
		TotalPoint:                     1000000,
		MemberGuildUserRankingUserData: userData,
	}
	mgur.TopRanking.Append(userRankingCell)
	mgur.MyRanking.Append(userRankingCell)
	rankingBorderInfo := client.MemberGuildUserRankingBorderInfo{
		RankingBorderPoint: 1,
		UpperRank:          1,
		// LowerRank:         1,
		DisplayOrder: 1,
	}
	mgur.RankingBorders.Append(rankingBorderInfo)
	resp.MemberGuildUserRankingList.Append(mgur)

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildRanking", fetchMemberGuildRanking)
}
