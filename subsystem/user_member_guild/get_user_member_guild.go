package user_member_guild

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserMemberGuild(session *userdata.Session, memberGuildId int32) client.UserMemberGuild {
	if memberGuildId == GetCurrentMemberGuildId(session) {
		return GetCurrentUserMemberGuild(session)
	}
	userMemberGuild := client.UserMemberGuild{}
	exist, err := session.Db.Table("u_member_guild").Where("user_id = ? AND member_guild_id = ?",
		session.UserId, memberGuildId).Get(&userMemberGuild)
	utils.CheckErr(err)
	if !exist {
		userMemberGuild = client.UserMemberGuild{}
	}
	return userMemberGuild
}
