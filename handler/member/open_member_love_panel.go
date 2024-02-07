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
	"github.com/tidwall/gjson"
)

func openMemberLovePanel(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.OpenMemberLovePanelRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.OpenMemberLovePanelResponse{
		UserModel: &session.UserModel,
	}
	resp.MemberLovePanels.Append(user_member.UnlockMemberLovePanel(
		session, req.MemberId, req.MemberLovePanelId, req.MemberLovePanelCellIds.Slice))

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/member/openMemberLovePanel", openMemberLovePanel)
}
