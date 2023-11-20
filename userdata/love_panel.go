package userdata

import (
	"elichika/model"
	// "fmt"
)

func (session *Session) GetAllMemberLovePanels() []model.UserMemberLovePanel {
	lovePanels := []model.UserMemberLovePanel{}
	err := session.Db.Table("u_member").
		Where("user_id = ?", session.UserStatus.UserID).Find(&lovePanels)
	if err != nil {
		panic(err)
	}
	for i := range lovePanels {
		lovePanels[i].Fill()
	}
	return lovePanels
}

func (session *Session) GetMemberLovePanel(memberMasterID int) model.UserMemberLovePanel {
	panel, exists := session.UserMemberLovePanelDiffs[memberMasterID]
	if exists {
		return panel
	}
	exists, err := session.Db.Table("u_member").
		Where("user_id = ? AND member_master_id = ?", session.UserStatus.UserID, memberMasterID).
		Get(&panel)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("doesn't exist")
	}
	panel.Fill()
	return panel
}

func (session *Session) GetLovePanelCellIDs(memberID int) []int {
	userMemberLovePanel := session.GetMemberLovePanel(memberID)
	userMemberLovePanel.Fill()
	return userMemberLovePanel.MemberLovePanelCellIDs
}

func (session *Session) UpdateMemberLovePanel(panel model.UserMemberLovePanel) {
	session.UserMemberLovePanelDiffs[panel.MemberID] = panel
}

func (session *Session) FinalizeMemberLovePanelDiffs() []model.UserMemberLovePanel {
	panels := []model.UserMemberLovePanel{}
	for _, panel := range session.UserMemberLovePanelDiffs {
		_, err := session.Db.Table("u_member").
			Where("user_id = ? AND member_master_id = ?", panel.UserID, panel.MemberID).
			AllCols().Update(panel)
		if err != nil {
			panic(err)
		}
		panel.Fill()
		panels = append(panels, panel)
	}
	return panels
}
