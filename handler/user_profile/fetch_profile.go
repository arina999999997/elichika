package user_profile

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_profile"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchProfile(ctx *gin.Context) {
	req := request.UserProfileRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, user_profile.GetOtherUserProfileResponse(session, req.UserId))
}

func init() {
	router.AddHandler("/userProfile/fetchProfile", fetchProfile)
}
