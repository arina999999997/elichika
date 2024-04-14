package member_guild

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(member_guild): the logic of this part is wrong or missing

func fetchMemberGuildTop(ctx *gin.Context) {
	// There is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.FetchMemberGuildTopResponse{
		UserModelDiff: &session.UserModel,
	}
	rank := int32(0)
	for _, member := range session.Gamedata.Member {
		rank++
		resp.MemberGuildTopStatus.MemberGuildRankingAnimationInfo.Append(
			client.MemberGuildRankingAnimationInfo{
				MemberMasterId:          member.Id,
				MemberGuildRankingOrder: rank,
				MemberGuildRankingPoint: 100000 - rank*1000,
			})
	}

	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildTop", fetchMemberGuildTop)
}
