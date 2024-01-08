package handler

import (
	"elichika/config"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func UpdateCardNewFlag(ctx *gin.Context) {
	// mark the cards as read (is_new = false)
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	type UpdateCardNewFlagReq struct {
		CardMasterIds []int `json:"card_master_ids"`
	}
	req := UpdateCardNewFlagReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	for _, cardMasterId := range req.CardMasterIds {
		card := session.GetUserCard(int32(cardMasterId))
		card.IsNew = false
		session.UpdateUserCard(card)
	}

	signBody := session.Finalize("{}", "user_model_diff")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeIsAwakeningImage(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()

	req := model.CardAwakeningReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	userCard := session.GetUserCard(int32(req.CardMasterId))
	userCard.IsAwakeningImage = req.IsAwakeningImage
	session.UpdateUserCard(userCard)

	cardResp := session.Finalize("{}", "user_model_diff")
	resp := SignResp(ctx, cardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeFavorite(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()

	req := model.CardFavoriteReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	userCard := session.GetUserCard(int32(req.CardMasterId))
	userCard.IsFavorite = req.IsFavorite
	session.UpdateUserCard(userCard)

	cardResp := session.Finalize("{}", "user_model_diff")
	resp := SignResp(ctx, cardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GetOtherUserCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type OtherUserCardReq struct {
		UserId       int `json:"user_id"`
		CardMasterId int `json:"card_master_id"`
	}
	req := OtherUserCardReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	partnerCard := session.GetPartnerCardFromUserCard(userdata.GetOtherUserCard(req.UserId, req.CardMasterId))
	userCardResp, _ := sjson.Set("{}", "other_user_card", partnerCard)
	resp := SignResp(ctx, userCardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
