package live_deck

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_deck"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// request: ChangeNameLiveDeckRequest
// response: UserModelResponse
// error response: RecoverableExceptionResponse
func changeDeckNameLiveDeck(ctx *gin.Context) {
	req := request.ChangeNameLiveDeckRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureResponse := user_live_deck.SetLiveDeckName(session, req.DeckId, req.DeckName)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	router.AddHandler("/", "POST", "/liveDeck/changeDeckNameLiveDeck", changeDeckNameLiveDeck)
}
