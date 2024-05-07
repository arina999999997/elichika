package user_member_guild

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateUserMemberGuild(session *userdata.Session, userMemberGuild client.UserMemberGuild) {
	session.UserModel.UserMemberGuildById.Set(userMemberGuild.MemberGuildId, userMemberGuild)
}
