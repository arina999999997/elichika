package user_member

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateMember(session *userdata.Session, member client.UserMember) {
	session.UserModel.UserMemberByMemberId.Set(member.MemberMasterId, member)
}
