package user_profile

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setProfile(ctx *gin.Context) {
	req := request.SetUserProfileRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if req.Name.HasValue {
		session.UserStatus.Name.DotUnderText = req.Name.Value
	}
	if req.Nickname.HasValue {
		session.UserStatus.Nickname.DotUnderText = req.Nickname.Value
	}
	if req.Message.HasValue {
		session.UserStatus.Message.DotUnderText = req.Message.Value
	}
	if req.DeviceToken.HasValue {
		session.UserStatus.DeviceToken = req.DeviceToken.Value
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/userProfile/setProfile", setProfile)
}
