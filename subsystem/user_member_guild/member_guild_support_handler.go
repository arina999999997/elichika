package user_member_guild

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/userdata"
	"elichika/utils"
)

func userMemberGuildContentHandler(session *userdata.Session, content *client.Content) any {
	item, exist := session.UserModel.UserMemberGuildSupportItemById.Get(content.ContentId)
	if !exist {
		item = &client.UserMemberGuildSupportItem{}
		exist, err := session.Db.Table("u_member_guild_support_item").Where("user_id = ?", session.UserId).Get(item)
		utils.CheckErr(err)
		if !exist {
			item.SupportItemId = content.ContentId
		}
	}

	if int64(item.SupportItemResetAt) <= session.Time.Unix() {
		item.Amount = 0
		item.SupportItemResetAt = int32(GetCurrentMemberGuildEnd(session))
	}

	item.Amount += content.ContentAmount
	content.ContentAmount = 0
	session.UserModel.UserMemberGuildSupportItemById.Set(item.SupportItemId, *item)
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeMemberGuildSupport, userMemberGuildContentHandler)
}
