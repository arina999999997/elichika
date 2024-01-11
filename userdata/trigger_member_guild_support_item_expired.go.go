package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func triggerMemberGuildSupportItemExpiredFinalizer(session *Session) {
	for triggerId, trigger := range session.UserModel.UserInfoTriggerMemberGuildSupportItemExpiredByTriggerId.Map {
		if trigger != nil { // add
			genericDatabaseInsert(session, "u_info_trigger_member_guild_support_item_expired", *trigger)
		} else { // delete
			_, err := session.Db.Table("u_info_trigger_member_guild_support_item_expired").
				Where("user_id = ? AND trigger_id = ?", session.UserId, triggerId).
				Delete(&client.UserInfoTriggerMemberGuildSupportItemExpired{})
			utils.CheckErr(err)
		}
	}
}

func (session *Session) ReadMemberGuildSupportItemExpired() {
	session.UserModel.UserInfoTriggerMemberGuildSupportItemExpiredByTriggerId.
		LoadFromDb(session.Db, session.UserId, "u_info_trigger_member_guild_support_item_expired", "trigger_id")

	for key := range session.UserModel.UserInfoTriggerMemberGuildSupportItemExpiredByTriggerId.Map {
		session.UserModel.UserInfoTriggerMemberGuildSupportItemExpiredByTriggerId.SetNull(key)
	}
	// already marked as removed, the finalizer will take care of things
	// there's also no need to remove the item, the client won't show them if they're expired
}

// TODO: Trigger member love level up isn't really that persistent, so it's probably better to only keep it in ram
// This could be done by keeping a full user model in ram too.

func init() {
	addFinalizer(triggerMemberGuildSupportItemExpiredFinalizer)
}
