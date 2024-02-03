package user_info_trigger

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func RemoveTriggerMemberLoveLevelUp(session *userdata.Session, triggerId int64) {
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.SetNull(triggerId)
}

func AddTriggerMemberLoveLevelUp(session *userdata.Session, trigger client.UserInfoTriggerMemberLoveLevelUp) {
	if trigger.TriggerId == 0 {
		trigger.TriggerId = session.NextUniqueId()
	}
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Set(trigger.TriggerId, trigger)
}

func triggerMemberLoveLevelUpFinalizer(session *userdata.Session) {
	for triggerId, trigger := range session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Map {
		if trigger != nil {
			userdata.GenericDatabaseInsert(session, "u_info_trigger_member_love_level_up", *trigger)
		} else {
			_, err := session.Db.Table("u_info_trigger_member_love_level_up").
				Where("user_id = ? AND trigger_id = ?", session.UserId, triggerId).
				Delete(&client.UserInfoTriggerMemberLoveLevelUp{})
			utils.CheckErr(err)
		}
	}
}

func ReadAllMemberLoveLevelUpTriggers(session *userdata.Session) {
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.LoadFromDb(
		session.Db, session.UserId, "u_info_trigger_member_love_level_up", "trigger_id")
	for key := range session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Map {
		session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.SetNull(key)
	}
}

func init() {
	userdata.AddFinalizer(triggerMemberLoveLevelUpFinalizer)
}
