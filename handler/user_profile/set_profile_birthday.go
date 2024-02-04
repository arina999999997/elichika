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
	"github.com/tidwall/gjson"
)

func setProfileBirthday(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetUserProfileBirthDayRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// birthdate is probably calculated using gplay or apple id
	session.UserStatus.BirthDay = generic.NewNullable(req.Day)
	session.UserStatus.BirthMonth = generic.NewNullable(req.Month)

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseNameInput {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseCorePlayable
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/userProfile/setProfileBirthday", setProfileBirthday)
}
