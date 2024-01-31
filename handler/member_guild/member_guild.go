package member_guild

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
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

// TODO: the logic of this part is wrong or missing

func FetchMemberGuildTop(ctx *gin.Context) {
	// There is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

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

func FetchMemberGuildSelect(ctx *gin.Context) {
	// There is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// this just work
	resp := response.FetchMemberGuildSelectResponse{}

	common.JsonResponse(ctx, resp)
}

func FetchMemberGuildRanking(ctx *gin.Context) {
	// There is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

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

func CheerMemberGuild(ctx *gin.Context) {
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

func JoinMemberGuild(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.JoinMemberGuildRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.MemberGuildMemberMasterId = generic.NewNullable(req.MemberMasterId)
	session.UserStatus.MemberGuildLastUpdatedAt = session.Time.Unix()

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	// TODO(refactor): move to individual files. 
	router.AddHandler("/memberGuild/cheerMemberGuild", CheerMemberGuild)
	router.AddHandler("/memberGuild/fetchMemberGuildRanking", FetchMemberGuildRanking)
	router.AddHandler("/memberGuild/fetchMemberGuildRankingYear", FetchMemberGuildRanking)
	router.AddHandler("/memberGuild/fetchMemberGuildSelect", FetchMemberGuildSelect)
	router.AddHandler("/memberGuild/fetchMemberGuildTop", FetchMemberGuildTop)
	router.AddHandler("/memberGuild/joinMemberGuild", JoinMemberGuild)
}
