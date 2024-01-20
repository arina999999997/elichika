package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
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

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	JsonResponse(ctx, session.FetchProfile(req.UserId))
}

func SetProfile(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetUserProfileRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
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
	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func SetProfileBirthday(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetUserProfileBirthDayRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// birthdate is probably calculated using gplay or apple id
	session.UserStatus.BirthDay = generic.NewNullable(req.Day)
	session.UserStatus.BirthMonth = generic.NewNullable(req.Month)

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseNameInput {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseCorePlayable
	}

	session.Finalize()
	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func SetRecommendCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetRecommendCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.RecommendCardMasterId = req.CardMasterId

	session.Finalize()
	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func SetLivePartner(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetLivePartnerCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// set the bit on the correct card
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	newCard := session.GetUserCard(req.CardMasterId)
	newCard.LivePartnerCategories |= (1 << req.LivePartnerCategoryId)
	session.UpdateUserCard(newCard)

	// TODO(refactor): Get rid of this bit hacking stuff
	// remove the bit on the other cards
	partnerCards := userdata.FetchPartnerCards(userId)
	for _, card := range partnerCards {
		if card.CardMasterId == req.CardMasterId {
			continue
		}
		if (card.LivePartnerCategories & (1 << req.LivePartnerCategoryId)) != 0 {
			card.LivePartnerCategories ^= (1 << req.LivePartnerCategoryId)
			session.UpdateUserCard(card)
		}
	}

	session.Finalize()
	JsonResponse(ctx, response.EmptyResponse{})
}

func SetScoreOrComboLive(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetLiveRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
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
	JsonResponse(ctx, response.SetLiveResponse{
		LiveDifficultyMasterId: req.LiveDifficultyMasterId,
	})
}
