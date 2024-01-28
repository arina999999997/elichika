package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func (session *Session) DeleteTriggerBasic(triggerId int64) {
	session.UserModel.UserInfoTriggerBasicByTriggerId.SetNull(triggerId)
}

func (session *Session) AddTriggerBasic(trigger client.UserInfoTriggerBasic) {
	if trigger.TriggerId == 0 {
		trigger.TriggerId = session.Time.UnixNano() + session.UniqueCount
		session.UniqueCount++
	}
	session.UserModel.UserInfoTriggerBasicByTriggerId.Set(trigger.TriggerId, trigger)
}

func triggerBasicFinalizer(session *Session) {
	for triggerId, trigger := range session.UserModel.UserInfoTriggerBasicByTriggerId.Map {
		if trigger != nil { // add
			GenericDatabaseInsert(session, "u_info_trigger_basic", *trigger)
		} else { // delete
			_, err := session.Db.Table("u_info_trigger_basic").
				Where("user_id = ? AND trigger_id = ?", session.UserId, triggerId).
				Delete(&client.UserInfoTriggerBasic{})
			utils.CheckErr(err)
		}
	}
}

func init() {
	AddContentFinalizer(triggerBasicFinalizer)
}
