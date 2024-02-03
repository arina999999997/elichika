package user_info_trigger

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func DeleteTriggerBasic(session *userdata.Session, triggerId int64) {
	session.UserModel.UserInfoTriggerBasicByTriggerId.SetNull(triggerId)
}

func AddTriggerBasic(session *userdata.Session, trigger client.UserInfoTriggerBasic) {
	if trigger.TriggerId == 0 {
		trigger.TriggerId = session.NextUniqueId()
	}
	session.UserModel.UserInfoTriggerBasicByTriggerId.Set(trigger.TriggerId, trigger)
}

func triggerBasicFinalizer(session *userdata.Session) {
	for triggerId, trigger := range session.UserModel.UserInfoTriggerBasicByTriggerId.Map {
		if trigger != nil { // add
			userdata.GenericDatabaseInsert(session, "u_info_trigger_basic", *trigger)
		} else { // delete
			_, err := session.Db.Table("u_info_trigger_basic").
				Where("user_id = ? AND trigger_id = ?", session.UserId, triggerId).
				Delete(&client.UserInfoTriggerBasic{})
			utils.CheckErr(err)
		}
	}
}

func init() {
	userdata.AddFinalizer(triggerBasicFinalizer)
}
