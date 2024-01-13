package handler

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(refactor): Change to use request and response types
func OpenMemberLovePanel(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.OpenMemberLovePanelRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	panel := session.GetMemberLovePanel(int32(req.MemberId))
	for _, cellId := range req.MemberLovePanelCellIds {
		panel.MemberLovePanelCellIds.Append(cellId)
	}
	sort.Slice(panel.MemberLovePanelCellIds.Slice, func(i, j int) bool {
		return panel.MemberLovePanelCellIds.Slice[i] < panel.MemberLovePanelCellIds.Slice[j]
	})
	// remove resource
	for _, cellId := range req.MemberLovePanelCellIds {
		for _, resource := range gamedata.MemberLovePanelCell[cellId].Resources {
			session.RemoveResource(resource)
		}
	}

	// if is full panel, then we have to send a basic info trigger to actually open up the next panel
	unlockCount := panel.MemberLovePanelCellIds.Size()
	if unlockCount%5 == 0 {
		member := session.GetMember(panel.MemberId)
		masterLovePanel := gamedata.MemberLovePanel[int32(req.MemberLovePanelId)]
		if (masterLovePanel.NextPanel != nil) && (masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= member.LoveLevel) {
			session.AddTriggerBasic(client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
				ParamInt:        generic.NewNullable(masterLovePanel.NextPanel.Id)})
		}
	}
	session.UpdateMemberLovePanel(panel)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
