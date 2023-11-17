package handler

import (
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/model"
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
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	panel := session.GetMemberLovePanel(req.MemberID)
	panel.LovePanelLastLevelCellIDs = append(panel.LovePanelLastLevelCellIDs, req.MemberLovePanelCellIDs...)
	// remove resource
	for _, cellID := range req.MemberLovePanelCellIDs {
		for _, resource := range gamedata.MemberLovePanelCell[cellID].Resources {
			session.RemoveResource(resource)
		}
	}

	// if is full panel, then we have to send a basic info trigger to actually open up the next panel
	if len(panel.LovePanelLastLevelCellIDs) == 5 {
		member := session.GetMember(panel.MemberID)
		masterLovePanel := gamedata.MemberLovePanel[req.MemberLovePanelID]
		if (masterLovePanel.NextPanel != nil) && (masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= member.LoveLevel) {
			// TODO: remove magic id from love panel system
			panel.LevelUp()
			session.AddTriggerBasic(model.TriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeUnlockBondBoard,
				ParamInt:        panel.LovePanelLevel*1000 + panel.MemberID})
		}
	}
	session.UpdateMemberLovePanel(panel)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
