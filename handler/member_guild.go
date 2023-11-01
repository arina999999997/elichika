package handler

import (
	"elichika/config"
	"elichika/model"
	"elichika/response"
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
func FetchMemberGuildTop(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	// this does't work
	signBody := session.Finalize(GetData("fetchMemberGuildTop.json"), "user_model_diff")
	resp := SignResp(ctx, signBody, config.SessionKey)
	// fmt.Println(resp)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchMemberGuildSelect(ctx *gin.Context) {
	resp := SignResp(ctx, GetData("fetchMemberGuildSelect.json"), config.SessionKey)
	// fmt.Println(resp)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// doesn't work (kinda expected)
// probably need to read the code more carefully
func FetchMemberGuildRanking(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	respObj := response.FetchMemberGuildRankingResponse{}
	respObj.MemberGuildRanking.ViewYear = 2022
	respObj.MemberGuildRanking.NextYear = 2023
	respObj.MemberGuildRanking.PreviousYear = 2021
	oneTerm := response.MemberGuildRankingOneTerm{
		MemberGuildID: 1,
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
				memberID := group*100 + id
				oneTerm.Channels = append(oneTerm.Channels, response.MemberGuildRankingOneTermCell{
					Order:          order,
					TotalPoint:     1000000,
					MemberMasterID: memberID,
				})
				order++
			}
		}
	}

	respObj.MemberGuildRanking.MemberGuildRankingList = append(respObj.MemberGuildRanking.MemberGuildRankingList, oneTerm)

	mgur := response.MemberGuildUserRanking{
		MemberGuildID: 1,
	}
	userData := response.MemberGuildUserRankingUserData{
		UserID:                 session.UserStatus.UserID,
		UserRank:               session.UserStatus.Rank,
		CardMasterID:           session.UserStatus.RecommendCardMasterID,
		Level:                  80,
		IsAwakening:            true,
		IsAllTrainingActivated: true,
		EmblemMasterID:         session.UserStatus.EmblemID,
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
	// fmt.Println(resp)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func CheerMemberGuild(ctx *gin.Context) {
	// this is extracted from Serialization_DeserializeCheerMemberGuildResponse
	type CheerMemberGuildResp struct {
		Rewards              []model.Content `json:"rewards"`
		MemberGuildTopStatus []any           `json:"member_guild_top_status"`
		UserModelDiff        []any           `json:"user_model_diff"`
	}

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	signBody := session.Finalize(GetData("fetchMemberGuildTop.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "rewards", []model.Content{})

	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func JoinMemberGuild(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// fmt.Println(reqBody)
	type JoinMemberGuildReq struct {
		MemberMasterID int `json:"member_master_id"`
	}
	req := JoinMemberGuildReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	session.UserStatus.MemberGuildMemberMasterID = req.MemberMasterID
	signBody := session.Finalize(GetData("userModel.json"), "user_model")

	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
