package handler

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func OpenMemberLovePanel(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.OpenMemberLovePanelRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	panel := session.GetMemberLovePanel(req.MemberId)
	for _, cellId := range req.MemberLovePanelCellIds.Slice {
		panel.MemberLovePanelCellIds.Append(cellId)
	}
	sort.Slice(panel.MemberLovePanelCellIds.Slice, func(i, j int) bool {
		return panel.MemberLovePanelCellIds.Slice[i] < panel.MemberLovePanelCellIds.Slice[j]
	})
	// remove resource
	for _, cellId := range req.MemberLovePanelCellIds.Slice {
		for _, resource := range session.Gamedata.MemberLovePanelCell[cellId].Resources {
			session.RemoveResource(resource)
		}
	}

	// if is full panel, then we have to send a basic info trigger to actually open up the next panel
	unlockCount := panel.MemberLovePanelCellIds.Size()
	if unlockCount%5 == 0 {
		member := session.GetMember(panel.MemberId)
		masterLovePanel := session.Gamedata.MemberLovePanel[req.MemberLovePanelId]
		if (masterLovePanel.NextPanel != nil) && (masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= member.LoveLevel) {
			session.AddTriggerBasic(client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
				ParamInt:        generic.NewNullable(masterLovePanel.NextPanel.Id)})
		}
	}
	session.UpdateMemberLovePanel(panel)

	resp := response.OpenMemberLovePanelResponse{
		UserModel: &session.UserModel,
	}
	resp.MemberLovePanels.Append(panel)

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, &resp)
}
