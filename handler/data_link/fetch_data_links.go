package data_link

import (
	"elichika/client"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func fetchDataLinks(ctx *gin.Context) {
	// there's no request body
	common.JsonResponse(ctx, client.LinkedInfo{
		IsTakeOverIdLinked: true,
	})
}

func init() {
	router.AddHandler("/dataLink/fetchDataLinks", fetchDataLinks)
}
