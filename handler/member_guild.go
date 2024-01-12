package handler

import (
	"elichika/client"
	"elichika/config"
	"elichika/generic"
	"elichika/protocol/response"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// TODO: the logic of this part is wrong or missing
// the request and response are sound for the most part
// TODO(refactor): Change to use request and response types
func FetchMemberGuildTop(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	// this does't work
	signBody := session.Finalize(GetData("fetchMemberGuildTop.json"), "user_model_diff")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func FetchMemberGuildSelect(ctx *gin.Context) {
	resp := SignResp(ctx, GetData("fetchMemberGuildSelect.json"), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func FetchMemberGuildRanking(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	respObj := response.FetchMemberGuildRankingResponse{}
	respObj.MemberGuildRanking.ViewYear = 2022
	respObj.MemberGuildRanking.NextYear = 2023
	respObj.MemberGuildRanking.PreviousYear = 2021
	oneTerm := response.MemberGuildRankingOneTerm{
		MemberGuildId: 1,
		StartAt:       1,
		EndAt:         1,
	}
	{
		order := 1
		for group := 0; group <= 2; group++ {
			for id := 1; id <= 12; id++ {
				if (id > 9) && (group != 2) {
					break
				}
				memberId := group*100 + id
				oneTerm.Channels = append(oneTerm.Channels, response.MemberGuildRankingOneTermCell{
					Order:          order,
					TotalPoint:     1000000,
					MemberMasterId: memberId,
				})
				order++
			}
		}
	}

	respObj.MemberGuildRanking.MemberGuildRankingList = append(respObj.MemberGuildRanking.MemberGuildRankingList, oneTerm)

	mgur := response.MemberGuildUserRanking{
		MemberGuildId: 1,
	}
	userData := response.MemberGuildUserRankingUserData{
		UserId:                 session.UserId,
		UserRank:               int(session.UserStatus.Rank),
		CardMasterId:           int(session.UserStatus.RecommendCardMasterId),
		Level:                  80,
		IsAwakening:            true,
		IsAllTrainingActivated: true,
		EmblemMasterId:         int(session.UserStatus.EmblemId),
	}
	userData.UserName.DotUnderText = session.UserStatus.Name.DotUnderText
	userRankingCell := response.MemberGuildUserRankingCell{
		Order:                          1,
		TotalPoint:                     1000000,
		MemberGuildUserRankingUserData: userData,
	}
	mgur.TopRanking = append(mgur.TopRanking, userRankingCell)
	mgur.MyRanking = append(mgur.MyRanking, userRankingCell)
	rankingBorderInfo := response.MemberGuildUserRankingBorderInfo{
		RankingOrderPoint: 1,
		UpperRank:         1,
		LowerRank:         1,
		DisplayOrder:      1,
	}
	mgur.RankingBorders = append(mgur.RankingBorders, rankingBorderInfo)
	respObj.MemberGuildUserRankingList = append(respObj.MemberGuildUserRankingList, mgur)

	respBytes, err := json.Marshal(respObj)
	utils.CheckErr(err)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func CheerMemberGuild(ctx *gin.Context) {
	// this is extracted from Serialization_DeserializeCheerMemberGuildResponse
	// type CheerMemberGuildResp struct {
	// 	Rewards              []client.Content `json:"rewards"`
	// 	MemberGuildTopStatus []any           `json:"member_guild_top_status"`
	// 	UserModelDiff        []any           `json:"user_model_diff"`
	// }

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	signBody := session.Finalize(GetData("fetchMemberGuildTop.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "rewards", []client.Content{})

	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func JoinMemberGuild(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type JoinMemberGuildReq struct {
		MemberMasterId int32 `json:"member_master_id"`
	}
	req := JoinMemberGuildReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.UserStatus.MemberGuildMemberMasterId = generic.NewNullable(req.MemberMasterId)
	signBody := session.Finalize("{}", "user_model")

	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
