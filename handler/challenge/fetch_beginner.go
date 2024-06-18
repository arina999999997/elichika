package challenge

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_beginner_challenge"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchBeginner(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, user_beginner_challenge.FetchChallengeBeginner(session))
}

func init() {
	router.AddHandler("/", "POST", "/challenge/fetchBeginner", fetchBeginner)
}
