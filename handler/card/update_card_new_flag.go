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
	"github.com/tidwall/gjson"
)

func updateCardNewFlag(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateCardNewFlagRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, cardMasterId := range req.CardMasterIds.Slice {
		card := user_card.GetUserCard(session, cardMasterId)
		card.IsNew = false
		user_card.UpdateUserCard(session, card)
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UpdateCardNewFlagResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/card/updateCardNewFlag", updateCardNewFlag)
}
