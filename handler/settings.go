package handler

import (
	"elichika/client/response"

	"github.com/gin-gonic/gin"
)

// TODO(push_notification): Support this once we figure push notifications out
func UpdatePushNotificationSettings(ctx *gin.Context) {
	// reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// req := request.PushNotificationSettingsRequest{}
	// err := json.Unmarshal([]byte(reqBody), &req)
	// utils.CheckErr(err)

	JsonResponse(ctx, response.EmptyResponse{})
}
