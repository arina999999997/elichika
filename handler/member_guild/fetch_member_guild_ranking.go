package member_guild

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member_guild"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchMemberGuildRanking(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, user_member_guild.FetchMemberGuildRanking(session))
}

func init() {
	router.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildRanking", fetchMemberGuildRanking)
}
