package userdata

import (
	"elichika/utils"
)

func memberGuildFinalizer(session *Session) {
	for _, userMemberGuild := range session.UserModel.UserMemberGuildById.Objects {
		affected, err := session.Db.Table("u_member_guild").Where("user_id = ? AND member_guild_id = ?",
			session.UserStatus.UserId, userMemberGuild.MemberGuildId).AllCols().Update(userMemberGuild)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_member_guild").Insert(userMemberGuild)
			utils.CheckErr(err)
		}
	}
	for _, userMemberGuildSupportItem := range session.UserModel.UserMemberGuildSupportItemById.Objects {
		affected, err := session.Db.Table("u_member_guild_support_item").Where("user_id = ? AND support_item_id = ?",
			session.UserStatus.UserId, userMemberGuildSupportItem.SupportItemId).AllCols().Update(userMemberGuildSupportItem)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_member_guild_support_item").Insert(userMemberGuildSupportItem)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(memberGuildFinalizer)
	addGenericTableFieldPopulator("u_member_guild", "UserMemberGuildById")
	addGenericTableFieldPopulator("u_member_guild_support_item", "UserMemberGuildSupportItemById")
}
