package handler

import (
	"elichika/client/response"

	"github.com/gin-gonic/gin"
)

func FetchStill(ctx *gin.Context) {
	// there is no request body

	JsonResponse(ctx, &response.FetchStillResponse{})
}
