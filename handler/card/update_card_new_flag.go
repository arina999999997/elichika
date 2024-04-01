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

func updateCardNewFlag(ctx *gin.Context) {
	req := request.UpdateCardNewFlagRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	for _, cardMasterId := range req.CardMasterIds.Slice {
		card := user_card.GetUserCard(session, cardMasterId)
		card.IsNew = false
		user_card.UpdateUserCard(session, card)
	}

	common.JsonResponse(ctx, response.UpdateCardNewFlagResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/card/updateCardNewFlag", updateCardNewFlag)
}
