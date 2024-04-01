package shop

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchShopSnsCoin(ctx *gin.Context) {
	// there is no request body
	// special behaviour to add 10000 gems if someone try to buy gem
	session := ctx.MustGet("session").(*userdata.Session)
	session.UserModel.UserStatus.FreeSnsCoin += 10000

	common.JsonResponse(ctx, &response.FetchShopSnsCoinResponse{})
}

func init() {
	router.AddHandler("/shop/fetchShopSnsCoin", fetchShopSnsCoin)
}
