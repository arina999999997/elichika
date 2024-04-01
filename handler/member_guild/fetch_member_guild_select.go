package member_guild

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	// "elichika/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(member_guild): the logic of this part is wrong or missing

func fetchMemberGuildSelect(ctx *gin.Context) {
	// There is no request body
	// session := ctx.MustGet("session").(*userdata.Session)

	// this just work
	resp := response.FetchMemberGuildSelectResponse{}

	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/memberGuild/fetchMemberGuildSelect", fetchMemberGuildSelect)
}
