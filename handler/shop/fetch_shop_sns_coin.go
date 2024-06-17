package shop

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func fetchShopSnsCoin(ctx *gin.Context) {
	// there is no request body
	common.JsonResponse(ctx, &response.FetchShopSnsCoinResponse{})
}

func init() {
	router.AddHandler("/", "POST", "/shop/fetchShopSnsCoin", fetchShopSnsCoin)
}
