package live_deck

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"
	"elichika/subsystem/user_live_difficulty"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func fetchLiveDeckSelect(ctx *gin.Context) {
	// return last deck for this song
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchLiveDeckSelectRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
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
