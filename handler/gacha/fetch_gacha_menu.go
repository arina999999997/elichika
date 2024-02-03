package gacha

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_gacha"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchGachaMenu(ctx *gin.Context) {
	// there is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, &response.FetchGachaMenuResponse{
		GachaList:     user_gacha.GetGachaList(session),
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/gacha/fetchGachaMenu", fetchGachaMenu)
}
