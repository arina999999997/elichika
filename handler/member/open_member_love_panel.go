package member

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func openMemberLovePanel(ctx *gin.Context) {
	req := request.OpenMemberLovePanelRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.OpenMemberLovePanelResponse{
		UserModel: &session.UserModel,
	}
	resp.MemberLovePanels.Append(user_member.UnlockMemberLovePanel(
		session, req.MemberId, req.MemberLovePanelId, req.MemberLovePanelCellIds.Slice))

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/member/openMemberLovePanel", openMemberLovePanel)
}
