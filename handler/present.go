package handler

import (
	"elichika/client/response"

	"github.com/gin-gonic/gin"
)

// TODO(present): Implement presents
func FetchPresent(ctx *gin.Context) {
	// there is no request body

	JsonResponse(ctx, &response.FetchPresentResponse{})
}
