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
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	panel := session.GetMemberLovePanel(int32(req.MemberId))
	panel.LovePanelLastLevelCellIds = append(panel.LovePanelLastLevelCellIds, req.MemberLovePanelCellIds...)
	// remove resource
	for _, cellId := range req.MemberLovePanelCellIds {
		for _, resource := range gamedata.MemberLovePanelCell[cellId].Resources {
			session.RemoveResource(resource)
		}
	}

	// if is full panel, then we have to send a basic info trigger to actually open up the next panel
	if len(panel.LovePanelLastLevelCellIds) == 5 {
		member := session.GetMember(panel.MemberId)
		masterLovePanel := gamedata.MemberLovePanel[int32(req.MemberLovePanelId)]
		if (masterLovePanel.NextPanel != nil) && (masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= member.LoveLevel) {
			// TODO: remove magic id from love panel system
			panel.LevelUp()
			session.AddTriggerBasic(client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
				ParamInt:        generic.NewNullable(int32(panel.LovePanelLevel*1000 + panel.MemberId))})
		}
	}
	session.UpdateMemberLovePanel(panel)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
