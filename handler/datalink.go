package handler

import (
	"elichika/client"

	"github.com/gin-gonic/gin"
)

func FetchDataLinks(ctx *gin.Context) {
	// there's no request body
	JsonResponse(ctx, client.LinkedInfo{
		IsTakeOverIdLinked: true,
	})
}
