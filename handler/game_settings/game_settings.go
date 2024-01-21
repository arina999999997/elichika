package game_settings

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

// TODO(push_notification): Support this once we figure push notifications out
func UpdatePushNotificationSettings(ctx *gin.Context) {
	// reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// req := request.PushNotificationSettingsRequest{}
	// err := json.Unmarshal([]byte(reqBody), &req)
	// utils.CheckErr(err)

	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	router.AddHandler("/gameSettings/updatePushNotificationSettings", UpdatePushNotificationSettings)

}
