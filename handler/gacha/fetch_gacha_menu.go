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
	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, &response.FetchGachaMenuResponse{
		GachaList:     user_gacha.GetGachaList(session),
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/gacha/fetchGachaMenu", fetchGachaMenu)
}
