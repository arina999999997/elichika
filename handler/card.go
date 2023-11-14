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
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	type UpdateCardNewFlagReq struct {
		CardMasterIDs []int `json:"card_master_ids"`
	}
	req := UpdateCardNewFlagReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	for _, cardMasterID := range req.CardMasterIDs {
		card := session.GetUserCard(cardMasterID)
		card.IsNew = false
		session.UpdateUserCard(card)
	}

	signBody := session.Finalize(GetData("userModelDiff.json"), "user_model_diff")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeIsAwakeningImage(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// fmt.Println(reqBody)

	req := model.CardAwakeningReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	userCard := session.GetUserCard(req.CardMasterID)
	userCard.IsAwakeningImage = req.IsAwakeningImage
	session.UpdateUserCard(userCard)

	cardResp := session.Finalize(GetData("userModelDiff.json"), "user_model_diff")
	resp := SignResp(ctx, cardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeFavorite(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// fmt.Println(reqBody)

	req := model.CardFavoriteReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	userCard := session.GetUserCard(req.CardMasterID)
	userCard.IsFavorite = req.IsFavorite
	session.UpdateUserCard(userCard)

	cardResp := session.Finalize(GetData("userModelDiff.json"), "user_model_diff")
	resp := SignResp(ctx, cardResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GetOtherUserCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// fmt.Println(reqBody)
	type OtherUserCardReq struct {
		UserID       int `json:"user_id"`
		CardMasterID int `json:"card_master_id"`
	}
	req := OtherUserCardReq{}
	// userCardReq := model.UserCardReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	partnerCard := session.GetPartnerCardFromUserCard(userdata.GetOtherUserCard(req.UserID, req.CardMasterID))
	userCardResp, _ := sjson.Set("{}", "other_user_card", partnerCard)
	resp := SignResp(ctx, userCardResp, config.SessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
