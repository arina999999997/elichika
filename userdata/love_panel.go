package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetMemberLovePanel(memberMasterID int) model.UserMemberLovePanel {
	panel, exist := session.UserMemberLovePanelDiffs[memberMasterID]
	if exist {
		return panel
	}
	exist, err := session.Db.Table("u_member").
		Where("user_id = ? AND member_master_id = ?", session.UserStatus.UserID, memberMasterID).
		Get(&panel)
	utils.CheckErr(err)
	if !exist {
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

func finalizeMemberLovePanelDiffs(session *Session) {
	for _, panel := range session.UserMemberLovePanelDiffs {
		session.UserMemberLovePanels = append(session.UserMemberLovePanels, panel)
	}
	for i := range session.UserMemberLovePanels {
		// TODO: this is not necessary after we split the database
		session.UserMemberLovePanels[i].Normalize()
		affected, err := session.Db.Table("u_member").
			Where("user_id = ? AND member_master_id = ?", session.UserMemberLovePanels[i].UserID,
				session.UserMemberLovePanels[i].MemberID).AllCols().Update(session.UserMemberLovePanels[i])
		utils.CheckErr(err)
		if affected != 1 {
			panic("wrong number of member affected!")
		}
		session.UserMemberLovePanels[i].Fill()
	}
}

func memberLovePanelPopulator(session *Session) {
	err := session.Db.Table("u_member").
		Where("user_id = ?", session.UserStatus.UserID).Find(&session.UserMemberLovePanels)
	utils.CheckErr(err)
	for i := range session.UserMemberLovePanels {
		session.UserMemberLovePanels[i].Fill()
	}
}

func init() {
	addPopulator(memberLovePanelPopulator)
	// TODO: separate the database so we can use this finalizer instead of calling it manually
	// addFinalizer(finalizeMemberLovePanelDiffs)
}
