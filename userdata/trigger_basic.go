package userdata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/utils"
)

func (session *Session) DeleteTriggerBasic(triggerId int64) {
	session.UserModel.UserInfoTriggerBasicByTriggerId.SetZero(triggerId)
}

func (session *Session) AddTriggerBasic(trigger client.UserInfoTriggerBasic) {
	if trigger.TriggerId == 0 {
		trigger.TriggerId = session.Time.UnixNano() + session.UniqueCount
		session.UniqueCount++
	}
	session.UserModel.UserInfoTriggerBasicByTriggerId.Set(trigger.TriggerId, generic.NewNullable(trigger))
}

func triggerBasicFinalizer(session *Session) {
	for triggerId, trigger := range session.UserModel.UserInfoTriggerBasicByTriggerId.Map {
		if trigger.HasValue { // add
			genericDatabaseInsert(session, "u_info_trigger_basic", trigger.Value)
		} else { // delete
			_, err := session.Db.Table("u_info_trigger_basic").Where("trigger_id = ?", triggerId).Delete(
				&client.UserInfoTriggerBasic{})
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(triggerBasicFinalizer)
}
