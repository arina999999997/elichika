package user_member

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateMemberLovePanel(session *userdata.Session, panel client.MemberLovePanel) {
	session.MemberLovePanelDiffs[panel.MemberId] = panel
}
