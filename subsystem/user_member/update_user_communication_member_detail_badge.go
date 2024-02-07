package user_member

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateUserCommunicationMemberDetailBadge(session *userdata.Session, badge client.UserCommunicationMemberDetailBadge) {
	session.UserModel.UserCommunicationMemberDetailBadgeById.Set(badge.MemberMasterId, badge)
}
