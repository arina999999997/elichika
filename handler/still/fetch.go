package still

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	// there is no request body

	common.JsonResponse(ctx, &response.FetchStillResponse{})
}

func init() {
	router.AddHandler("/still/fetch", fetch)
}
