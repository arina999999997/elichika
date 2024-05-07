package member_guild

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member_guild"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchMemberGuildTop(ctx *gin.Context) {
	// There is no request body
	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, user_member_guild.FetchMemberGuildTop(session))
}

func init() {
	router.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildTop", fetchMemberGuildTop)
}
