package love_ranking

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_love_ranking"
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

	common.JsonResponse(ctx, user_love_ranking.FetchLoveRanking(session, req.LoveRankingType, req.Condition, req.RankingOrder))
}

func init() {
	router.AddHandler("/", "POST", "/loveRanking/fetch", fetch)
}
