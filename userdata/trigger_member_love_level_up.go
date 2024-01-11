package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func (session *Session) RemoveTriggerMemberLoveLevelUp(triggerId int64) {
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.SetNull(triggerId)
}

func (session *Session) AddTriggerMemberLoveLevelUp(trigger client.UserInfoTriggerMemberLoveLevelUp) {
	if trigger.TriggerId == 0 {
		trigger.TriggerId = session.Time.UnixNano() + session.UniqueCount
		session.UniqueCount++
	}
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Set(trigger.TriggerId, trigger)
}

func triggerMemberLoveLevelUpFinalizer(session *Session) {
	for triggerId, trigger := range session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Map {
		if trigger != nil {
			genericDatabaseInsert(session, "u_info_trigger_member_love_level_up", *trigger)
		} else {
			_, err := session.Db.Table("u_info_trigger_member_love_level_up").
				Where("user_id = ? AND trigger_id = ?", session.UserId, triggerId).
				Delete(&client.UserInfoTriggerMemberLoveLevelUp{})
			utils.CheckErr(err)
		}
	}
}

func (session *Session) ReadAllMemberLoveLevelUpTriggers() {
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.LoadFromDb(
		session.Db, session.UserId, "u_info_trigger_member_love_level_up", "trigger_id")
	for key := range session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Map {
		session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.SetNull(key)
	}
}

func init() {
	addFinalizer(triggerMemberLoveLevelUpFinalizer)
}
