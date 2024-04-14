package member_guild

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// TODO(member_guild): the logic of this part is wrong or missing

func joinMemberGuild(ctx *gin.Context) {
	req := request.JoinMemberGuildRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	session.UserStatus.MemberGuildMemberMasterId = generic.NewNullable(req.MemberMasterId)
	session.UserStatus.MemberGuildLastUpdatedAt = session.Time.Unix()

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/memberGuild/joinMemberGuild", joinMemberGuild)
}
