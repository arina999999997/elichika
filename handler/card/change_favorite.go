package card

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func changeFavorite(ctx *gin.Context) {
	req := request.ChangeFavoriteRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	userCard := user_card.GetUserCard(session, req.CardMasterId)
	userCard.IsFavorite = req.IsFavorite
	user_card.UpdateUserCard(session, userCard)

	common.JsonResponse(ctx, &response.ChangeFavoriteResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/card/changeFavorite", changeFavorite)
}
