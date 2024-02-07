package user_member

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_info_trigger"
	"elichika/userdata"
)

func UnlockNewLovePanel(session *userdata.Session, memberId, oldLoveLevel, newLoveLevel int32) {
	currentLovePanel := GetMemberLovePanel(session, memberId)
	unlockCount := currentLovePanel.MemberLovePanelCellIds.Size()
	if (unlockCount > 0) && (unlockCount%5 == 0) {
		// love panel is maxed out
		lastCell := currentLovePanel.MemberLovePanelCellIds.Slice[unlockCount-1]
		masterLovePanel := session.Gamedata.MemberLovePanelCell[lastCell].MemberLovePanel

		if (masterLovePanel.NextPanel != nil) &&
			(masterLovePanel.NextPanel.LoveLevelMasterLoveLevel <= newLoveLevel) &&
			(masterLovePanel.NextPanel.LoveLevelMasterLoveLevel > oldLoveLevel) {
			nextPanel := masterLovePanel.NextPanel
			user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
				ParamInt:        generic.NewNullable(nextPanel.Id)})
		}
	}
}
