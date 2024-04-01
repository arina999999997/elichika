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

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	userCard := user_card.GetUserCard(session, req.CardMasterId)
	userCard.IsFavorite = req.IsFavorite
	user_card.UpdateUserCard(session, userCard)

	session.Finalize()
	common.JsonResponse(ctx, &response.ChangeFavoriteResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/card/changeFavorite", changeFavorite)
}
