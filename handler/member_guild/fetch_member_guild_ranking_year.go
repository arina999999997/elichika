package member_guild

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member_guild"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchMemberGuildRankingYear(ctx *gin.Context) {
	req := request.FetchMemberGuildRankingYearRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, response.FetchMemberGuildRankingYearResponse{
		user_member_guild.FetchMemberGuildRankingYear(session, req.Year),
	})
}

func init() {
	router.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildRankingYear", fetchMemberGuildRankingYear)
}
