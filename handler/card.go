package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func UpdateCardNewFlag(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateCardNewFlagRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, cardMasterId := range req.CardMasterIds.Slice {
		card := session.GetUserCard(int32(cardMasterId))
		card.IsNew = false
		session.UpdateUserCard(card)
	}

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, response.UpdateCardNewFlagResponse{
		UserModelDiff: &session.UserModel,
	})
}

func ChangeIsAwakeningImage(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ChangeIsAwakeningImageRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	userCard := session.GetUserCard(req.CardMasterId)
	userCard.IsAwakeningImage = req.IsAwakeningImage
	session.UpdateUserCard(userCard)

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, response.ChangeIsAwakeningImageResponse{
		UserModelDiff: &session.UserModel,
	})
}

func ChangeFavorite(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ChangeFavoriteRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	userCard := session.GetUserCard(req.CardMasterId)
	userCard.IsFavorite = req.IsFavorite
	session.UpdateUserCard(userCard)

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, &response.ChangeFavoriteResponse{
		UserModelDiff: &session.UserModel,
	})
}

func GetOtherUserCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.GetOtherUserCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// the name of request and response is not consistent for this one, for some reason
	JsonResponse(ctx, response.FetchOtherUserCardResponse{
		OtherUserCard: session.GetOtherUserCard(req.UserId, req.CardMasterId),
	})
}
