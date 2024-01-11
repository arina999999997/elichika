package userdata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/utils"
)

func (session *Session) RemoveTriggerMemberLoveLevelUp(triggerId int64) {
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.SetZero(triggerId)
}

func (session *Session) AddTriggerMemberLoveLevelUp(trigger client.UserInfoTriggerMemberLoveLevelUp) {
	if trigger.TriggerId == 0 {
		trigger.TriggerId = session.Time.UnixNano() + session.UniqueCount
		session.UniqueCount++
	}
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Set(trigger.TriggerId, generic.NewNullable(trigger))
}

func triggerMemberLoveLevelUpFinalizer(session *Session) {
	for triggerId, trigger := range session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Map {
		if trigger.HasValue {
			genericDatabaseInsert(session, "u_info_trigger_member_love_level_up", trigger.Value)
		} else {
			_, err := session.Db.Table("u_info_trigger_member_love_level_up").
				Where("trigger_id = ?", triggerId).
				Delete(&client.UserInfoTriggerMemberLoveLevelUp{})
			utils.CheckErr(err)
		}
	}
}

func (session *Session) ReadAllMemberLoveLevelUpTriggers() {
	session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.LoadFromDb(
		session.Db, session.UserId, "u_info_trigger_member_love_level_up", "trigger_id")
	for key := range session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.Map {
		session.UserModel.UserInfoTriggerMemberLoveLevelUpByTriggerId.SetZero(key)
	}
}

func init() {
	addFinalizer(triggerMemberLoveLevelUpFinalizer)
}
