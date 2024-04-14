package user_profile

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setProfileBirthday(ctx *gin.Context) {
	req := request.SetUserProfileBirthDayRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	// birthdate is probably calculated using gplay or apple id
	session.UserStatus.BirthDay = generic.NewNullable(req.Day)
	session.UserStatus.BirthMonth = generic.NewNullable(req.Month)

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseNameInput {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseCorePlayable
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/userProfile/setProfileBirthday", setProfileBirthday)
}
