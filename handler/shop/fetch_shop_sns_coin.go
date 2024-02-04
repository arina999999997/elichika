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
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.UserModel.UserStatus.FreeSnsCoin += 10000

	session.Finalize()
	common.JsonResponse(ctx, &response.FetchShopSnsCoinResponse{})
}

func init() {
	router.AddHandler("/shop/fetchShopSnsCoin", fetchShopSnsCoin)
}
