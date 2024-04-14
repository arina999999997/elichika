package communication_member

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_custom_background"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setTheme(ctx *gin.Context) {
	req := request.SetThemeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	member := user_member.GetMember(session, req.MemberMasterId)
	member.SuitMasterId = req.SuitMasterId
	member.CustomBackgroundMasterId = req.CustomBackgroundMasterId
	user_member.UpdateMember(session, member)
	user_custom_background.ReadCustomBackground(session, req.CustomBackgroundMasterId)

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/communicationMember/setTheme", setTheme)
}
