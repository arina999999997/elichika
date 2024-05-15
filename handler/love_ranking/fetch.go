package love_ranking

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_social"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	req := request.FetchLoveRankingRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	// TODO(ranking): fetch from db instead
	resp := response.FetchLoveRankingResponse{}
	resp.LoveRankingData.Append(client.LoveRankingData{
		RankingUser: user_social.GetRankingUser(session, session.UserId),
		Order:       1,
		LovePoint:   1000000,
	})
	resp.MyRankingOrder = generic.NewNullable(int32(1))
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/loveRanking/fetch", fetch)
}
