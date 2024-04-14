package member_guild

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member_guild"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func cheerMemberGuild(ctx *gin.Context) {
	req := request.CheerMemberGuildRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_member_guild.Cheer(session, req.CheerItemAmount)

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/memberGuild/cheerMemberGuild", cheerMemberGuild)
}
