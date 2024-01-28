package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func (session *Session) GetOtherUserMemberLovePanel(userId, memberId int32) client.MemberLovePanel {
	result := client.MemberLovePanel{}
	exist, err := session.Db.Table("u_member_love_panel").
		Where("user_id = ? AND member_id = ?", session.UserId, memberId).
		Get(&result)
	utils.CheckErr(err)
	if !exist {
		return client.MemberLovePanel{
			MemberId: memberId,
		}
	}
	return result
}
func (session *Session) GetMemberLovePanel(memberId int32) client.MemberLovePanel {
	panel, exist := session.MemberLovePanelDiffs[memberId]
	if exist {
		return panel
	}
	return session.GetOtherUserMemberLovePanel(int32(session.UserId), memberId)
}

func (session *Session) UpdateMemberLovePanel(panel client.MemberLovePanel) {
	session.MemberLovePanelDiffs[panel.MemberId] = panel
}

func finalizeMemberLovePanelDiffs(session *Session) {
	for _, panel := range session.MemberLovePanelDiffs {
		session.MemberLovePanels = append(session.MemberLovePanels, panel)
	}
	for i := range session.MemberLovePanels {
		affected, err := session.Db.Table("u_member_love_panel").
			Where("user_id = ? AND member_id = ?", session.UserId,
				session.MemberLovePanels[i].MemberId).AllCols().Update(session.MemberLovePanels[i])
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_member_love_panel", session.MemberLovePanels[i])
		}
	}
}

func memberLovePanelPopulator(session *Session) {
	for _, member := range session.Gamedata.Member {
		session.MemberLovePanels = append(session.MemberLovePanels, session.GetMemberLovePanel(member.Id))
	}
}

func init() {
	AddContentPopulator(memberLovePanelPopulator)
	// TODO: separate the database so we can use this finalizer instead of calling it manually
	// AddContentFinalizer(finalizeMemberLovePanelDiffs)
}
