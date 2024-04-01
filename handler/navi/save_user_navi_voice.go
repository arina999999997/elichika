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

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, naviVoiceMasterId := range req.NaviVoiceMasterIds.Slice {
		user_voice.UpdateUserVoice(session, naviVoiceMasterId, false)
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/navi/saveUserNaviVoice", saveUserNaviVoice)
}
