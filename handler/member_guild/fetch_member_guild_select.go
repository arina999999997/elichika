package member_guild

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member_guild"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchMemberGuildSelect(ctx *gin.Context) {
	// There is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, user_member_guild.FetchMemberGuildSelect(session))
}

func init() {
	router.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildSelect", fetchMemberGuildSelect)
}
