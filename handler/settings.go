package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO(refactor): Change to use request and response types
func UpdatePushNotificationSettings(ctx *gin.Context) {
	resp := SignResp(ctx, "{}", config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
