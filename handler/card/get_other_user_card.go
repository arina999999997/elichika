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

func getOtherUserCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.GetOtherUserCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// the name of request and response is not consistent for this one, for some reason
	common.JsonResponse(ctx, response.FetchOtherUserCardResponse{
		OtherUserCard: user_card.GetOtherUserCard(session, req.UserId, req.CardMasterId),
	})
}

func init() {
	router.AddHandler("/card/getOtherUserCard", getOtherUserCard)
}
