package user_member

import (
	"elichika/userdata"
)

func memberLovePanelPopulator(session *userdata.Session) {
	for _, member := range session.Gamedata.Member {
		session.MemberLovePanels = append(session.MemberLovePanels, GetMemberLovePanel(session, member.Id))
	}
}

func init() {
	userdata.AddContentPopulator(memberLovePanelPopulator)
}
