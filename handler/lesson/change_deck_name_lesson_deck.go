package handler

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_lesson_deck"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// request: ChangeNameLessonDeckRequest
// response: UserModelResponse
// error response: RecoverableExceptionResponse
func changeDeckNameLessonDeck(ctx *gin.Context) {
	req := request.ChangeNameLessonDeckRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureResponse := user_lesson_deck.SetLessonDeckName(session, req.DeckId, req.DeckName)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	router.AddHandler("/", "POST", "/lesson/changeDeckNameLessonDeck", changeDeckNameLessonDeck)
}
