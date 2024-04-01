package navi

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_voice"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func saveUserNaviVoice(ctx *gin.Context) {
	req := request.SaveUserNaviVoiceRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	for _, naviVoiceMasterId := range req.NaviVoiceMasterIds.Slice {
		user_voice.UpdateUserVoice(session, naviVoiceMasterId, false)
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/navi/saveUserNaviVoice", saveUserNaviVoice)
}
