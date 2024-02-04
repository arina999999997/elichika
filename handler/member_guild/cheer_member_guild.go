package member_guild

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(member_guild): the logic of this part is wrong or missing

func cheerMemberGuild(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.CheerMemberGuildRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.CheerMemberGuildResponse{
		UserModelDiff: &session.UserModel,
	}
	// this is the same with FetchMemberGuildTop
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

	// for now award 100 stargems, but we do have the drop table available
	// it's not clear whether the chance are the same though
	if !req.CheerItemAmount.HasValue {
		req.CheerItemAmount.Value = 1
		// mark the free chance as used
	}

	// remove the items or update the free cheer

	for i := int32(0); i < req.CheerItemAmount.Value; i++ {
		user_present.AddPresent(session, client.PresentItem{
			Content:          item.StarGem.Amount(100),
			PresentRouteType: enum.PresentRouteTypeMemberGuildSupportReward,
			PresentRouteId:   session.UserStatus.MemberGuildMemberMasterId, // this doesn't seems to be doing anything but the official server kept it
		})
		resp.Rewards.Append(item.StarGem.Amount(100))
	}

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/memberGuild/cheerMemberGuild", cheerMemberGuild)
}
