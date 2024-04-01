package game_settings

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

// TODO(push_notification): Support this once we figure push notifications out
func updatePushNotificationSettings(ctx *gin.Context) {
	// req := request.PushNotificationSettingsRequest{}
	// err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	// utils.CheckErr(err)

	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	router.AddHandler("/gameSettings/updatePushNotificationSettings", updatePushNotificationSettings)
}
