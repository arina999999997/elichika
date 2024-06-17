package user_member

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_info_trigger"
	"elichika/userdata"

	"sort"
)

func UnlockMemberLovePanel(session *userdata.Session, memberId, memberLovePanelId int32, lovePanelCellIds []int32) client.MemberLovePanel {
	panel := GetMemberLovePanel(session, memberId)
	for _, cellId := range lovePanelCellIds {
		panel.MemberLovePanelCellIds.Append(cellId)
	}
	sort.Slice(panel.MemberLovePanelCellIds.Slice, func(i, j int) bool {
		return panel.MemberLovePanelCellIds.Slice[i] < panel.MemberLovePanelCellIds.Slice[j]
	})
	// remove resource
	for _, cellId := range lovePanelCellIds {
		for _, resource := range session.Gamedata.MemberLovePanelCell[cellId].Resources {
			if config.Conf.ResourceConfig().ConsumePracticeItems {
				user_content.RemoveContent(session, resource)
			}
		}
	}

	// if is full panel, then we have to send a basic info trigger to actually open up the next panel
	unlockCount := panel.MemberLovePanelCellIds.Size()
	if unlockCount%5 == 0 {
		member := GetMember(session, panel.MemberId)
		masterLovePanel := session.Gamedata.MemberLovePanel[memberLovePanelId]
		if (masterLovePanel.NextPanel != nil) && (masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= member.LoveLevel) {
			user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
				ParamInt:        generic.NewNullable(masterLovePanel.NextPanel.Id)})
		}
	}
	UpdateMemberLovePanel(session, panel)
	return panel
}
