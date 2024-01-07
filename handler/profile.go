package handler

import (
	"elichika/config"
	"elichika/enum"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchProfile(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type FetchProfileReq struct {
		UserId int `json:"user_id"`
	}
	req := FetchProfileReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	profile := session.FetchProfile(req.UserId)

	signBody, err := json.Marshal(profile)
	utils.CheckErr(err)

	resp := SignResp(ctx, string(signBody), config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SetProfile(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	req := gjson.Parse(reqBody).Array()[0]
	if req.Get("name").String() != "" {
		session.UserStatus.Name.DotUnderText = gjson.Parse(reqBody).Array()[0].Get("name").String()
	} else if req.Get("nickname").String() != "" {
		session.UserStatus.Nickname.DotUnderText = gjson.Parse(reqBody).Array()[0].Get("nickname").String()
	} else if req.Get("message").String() != "" {
		session.UserStatus.Message.DotUnderText = gjson.Parse(reqBody).Array()[0].Get("message").String()
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
func SetProfileBirthday(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetProfileBirthdayRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// birthdate is probably calculated using gplay or apple id
	session.UserStatus.BirthDay = req.Day
	session.UserStatus.BirthMonth = req.Month

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseNameInput {
		// session.UserStatus.TutorialPhase = enum.TutorialPhaseStory1
		session.UserStatus.TutorialPhase = enum.TutorialPhaseCorePlayable
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SetRecommendCard(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	cardMasterId := int(gjson.Parse(reqBody).Array()[0].Get("card_master_id").Int())
	session.UserStatus.RecommendCardMasterId = cardMasterId

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SetLivePartner(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SetLivePartnerReq struct {
		LivePartnerCategoryId int `json:"live_partner_category_id"`
		CardMasterId          int `json:"card_master_id"`
	}
	req := SetLivePartnerReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	// set the bit on the correct card
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	newCard := session.GetUserCard(req.CardMasterId)
	newCard.LivePartnerCategories |= (1 << req.LivePartnerCategoryId)
	session.UpdateUserCard(newCard)

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

	session.Finalize("{}", "")
	// this is correct, the server send {}
	//
	resp := SignResp(ctx, "{}", config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SetScoreOrComboLive(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SetScoreOrComboReq struct {
		LiveDifficultyMasterId int `json:"live_difficulty_master_id"`
	}
	req := SetScoreOrComboReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	setProfile := session.GetUserSetProfile()
	if ctx.Request.URL.Path == "/userProfile/setScoreLive" {
		setProfile.VoltageLiveDifficultyId = req.LiveDifficultyMasterId
	} else {
		setProfile.ComboLiveDifficultyId = req.LiveDifficultyMasterId
	}
	session.SetUserSetProfile(setProfile)
	session.Finalize("{}", "dummy")
	resp := SignResp(ctx, reqBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
