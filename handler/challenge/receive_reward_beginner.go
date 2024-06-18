package challenge

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_beginner_challenge"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func receiveRewardBeginner(ctx *gin.Context) {
	req := request.ChallengeBeginnerRewardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, user_beginner_challenge.ReceiveRewardBeginner(session, req.ChallengeId, req.ChallengeCellId))
}

func init() {
	router.AddHandler("/", "POST", "/challenge/receiveRewardBeginner", receiveRewardBeginner)
}
