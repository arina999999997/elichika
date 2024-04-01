package live_deck

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_difficulty"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchLiveDeckSelect(ctx *gin.Context) {
	// return last deck for this song
	req := request.FetchLiveDeckSelectRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, response.FetchLiveDeckSelectResponse{
		LastPlayLiveDifficultyDeck: user_live_difficulty.GetLastPlayLiveDifficultyDeck(session, req.LiveDifficultyId),
	})
}

func init() {
	router.AddHandler("/liveDeck/fetchLiveDeckSelect", fetchLiveDeckSelect)
}
