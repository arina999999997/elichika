package game_settings

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

// TODO(push_notification): This only control the external push notification
// i.e. birthday and event.
// there are done using google's or apple's service, so it might not be possible to revive it at all
func updatePushNotificationSettings(ctx *gin.Context) {
	// req := request.PushNotificationSettingsRequest{}
	// err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	// utils.CheckErr(err)

	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	router.AddHandler("/", "POST", "/gameSettings/updatePushNotificationSettings", updatePushNotificationSettings)
}
