package love_ranking

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_profile"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func fetch(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchLoveRankingRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// TODO(ranking): fetch from db instead
	resp := response.FetchLoveRankingResponse{}
	resp.LoveRankingData.Append(client.LoveRankingData{
		RankingUser: user_profile.GetRankingUser(session, session.UserId),
		Order:       1,
		LovePoint:   1000000,
	})
	resp.MyRankingOrder = generic.NewNullable(int32(1))
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/loveRanking/fetch", fetch)
}
