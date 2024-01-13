package userdata

import (
	"elichika/client"
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetOtherUserMemberLovePanel(userId, memberId int32) client.MemberLovePanel {
	result := client.MemberLovePanel{
		MemberId: memberId,
	}
	_, err := session.Db.Table("u_member_love_panel").
		Where("user_id = ? AND member_id = ?", session.UserId, memberId).
		Get(&result)
	utils.CheckErr(err)
	return result
}

func (session *Session) GetMemberLovePanel(memberMasterId int32) model.UserMemberLovePanel {
	panel, exist := session.UserMemberLovePanelDiffs[memberMasterId]
	if exist {
		return panel
	}
	exist, err := session.Db.Table("u_member").
		Where("user_id = ? AND member_master_id = ?", session.UserId, memberMasterId).
		Get(&panel)
	utils.CheckErr(err)
	if !exist {
		panic("doesn't exist")
	}
	panel.Fill()
	return panel
}

func (session *Session) GetLovePanelCellIds(memberId int32) []int32 {
	userMemberLovePanel := session.GetMemberLovePanel(memberId)
	userMemberLovePanel.Fill()
	return userMemberLovePanel.MemberLovePanelCellIds
}

func (session *Session) UpdateMemberLovePanel(panel model.UserMemberLovePanel) {
	session.UserMemberLovePanelDiffs[panel.MemberId] = panel
}

func finalizeMemberLovePanelDiffs(session *Session) {
	for _, panel := range session.UserMemberLovePanelDiffs {
		session.UserMemberLovePanels = append(session.UserMemberLovePanels, panel)
	}
	for i := range session.UserMemberLovePanels {
		// TODO: this is not necessary after we split the database
		session.UserMemberLovePanels[i].Normalize()
		affected, err := session.Db.Table("u_member").
			Where("user_id = ? AND member_master_id = ?", session.UserId,
				session.UserMemberLovePanels[i].MemberId).AllCols().Update(session.UserMemberLovePanels[i])
		utils.CheckErr(err)
		if affected != 1 {
			panic("wrong number of member affected!")
		}
		session.UserMemberLovePanels[i].Fill()
	}
}

func memberLovePanelPopulator(session *Session) {
	err := session.Db.Table("u_member").
		Where("user_id = ?", session.UserId).Find(&session.UserMemberLovePanels)
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
