package shop

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func fetchShopPack(ctx *gin.Context) {
	common.JsonResponse(ctx, &response.FetchShopPackResponse{})
}

func init() {
	router.AddHandler("/", "POST", "/shop/fetchShopPack", fetchShopPack)
}
