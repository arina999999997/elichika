package bootstrap

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func getClearedPlatformAchievement(ctx *gin.Context) {
	common.JsonResponse(ctx, &response.GetClearedPlatformAchievementResponse{})
}

func init() {
	router.AddHandler("/bootstrap/getClearedPlatformAchievement", getClearedPlatformAchievement)
}
