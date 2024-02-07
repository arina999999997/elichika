package user_member

import (
	"elichika/userdata"
	"elichika/utils"
)

func memberLovePanelsFinalizer(session *userdata.Session) {
	for _, panel := range session.MemberLovePanelDiffs {
		session.MemberLovePanels = append(session.MemberLovePanels, panel)
	}
	for i := range session.MemberLovePanels {
		affected, err := session.Db.Table("u_member_love_panel").
			Where("user_id = ? AND member_id = ?", session.UserId,
				session.MemberLovePanels[i].MemberId).AllCols().Update(session.MemberLovePanels[i])
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_member_love_panel", session.MemberLovePanels[i])
		}
	}
}

func init() {
	userdata.AddFinalizer(memberLovePanelsFinalizer)
}
