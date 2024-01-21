package present

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

// TODO(present): Implement presents
func FetchPresent(ctx *gin.Context) {
	// there is no request body

	common.JsonResponse(ctx, &response.FetchPresentResponse{})
}

func init() {
	router.AddHandler("/present/fetch", FetchPresent)
}
