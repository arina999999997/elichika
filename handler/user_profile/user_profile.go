package user_profile

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_profile"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchProfile(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UserProfileRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, user_profile.GetOtherUserProfileResponse(session, req.UserId))
}

func SetProfile(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetUserProfileRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
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

func SetProfileBirthday(ctx *gin.Context) {
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

func SetRecommendCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetRecommendCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.RecommendCardMasterId = req.CardMasterId

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func SetScoreOrComboLive(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetLiveRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	setProfile := session.GetUserSetProfile()
	if ctx.Request.URL.Path == "/userProfile/setScoreLive" {
		setProfile.VoltageLiveDifficultyId = req.LiveDifficultyMasterId
	} else {
		setProfile.CommboLiveDifficultyId = req.LiveDifficultyMasterId
	}
	session.SetUserSetProfile(setProfile)

	session.Finalize()
	common.JsonResponse(ctx, response.SetLiveResponse{
		LiveDifficultyMasterId: req.LiveDifficultyMasterId,
	})
}

func init() {
	// TODO(refactor): move to individual files. 
	router.AddHandler("/userProfile/fetchProfile", FetchProfile)
	router.AddHandler("/userProfile/setCommboLive", SetScoreOrComboLive)
	router.AddHandler("/userProfile/setProfile", SetProfile)
	router.AddHandler("/userProfile/setProfileBirthday", SetProfileBirthday)
	router.AddHandler("/userProfile/setRecommendCard", SetRecommendCard)
	router.AddHandler("/userProfile/setScoreLive", SetScoreOrComboLive)
}
