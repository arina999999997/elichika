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

func getOtherUserCard(ctx *gin.Context) {
	req := request.GetOtherUserCardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	// the name of request and response is not consistent for this one, for some reason
	common.JsonResponse(ctx, response.FetchOtherUserCardResponse{
		OtherUserCard: user_card.GetOtherUserCard(session, req.UserId, req.CardMasterId),
	})
}

func init() {
	router.AddHandler("/", "POST", "/card/getOtherUserCard", getOtherUserCard)
}
