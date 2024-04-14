package daily_theater

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setLike(ctx *gin.Context) {
	req := request.DailyTheaterSetLikeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	session.UserModel.UserDailyTheaterByDailyTheaterId.Set(
		req.DailyTheaterId,
		client.UserDailyTheater{
			DailyTheaterId: req.DailyTheaterId,
			IsLiked:        req.IsLike,
		})

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/dailyTheater/setLike", setLike)
}
