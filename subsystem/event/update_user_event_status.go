package event

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func UpdateUserEventStatus(session *userdata.Session, status database.UserEventStatus) {
	affected, err := session.Db.Table("u_event_status").Where("user_id = ? AND event_id = ?", session.UserId, status.EventId).AllCols().
		Update(&status)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_event_status", status)
	}
}
