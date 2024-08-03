package event

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func GetUserEventStatus(session *userdata.Session, eventId int32) database.UserEventStatus {
	status := database.UserEventStatus{}
	exist, err := session.Db.Table("u_event_status").Where("user_id = ? AND event_id = ?", session.UserId, eventId).Get(&status)
	utils.CheckErr(err)
	if !exist {
		status = database.UserEventStatus{
			EventId:          eventId,
			IsFirstAccess:    true,
			IsNew:            false,
			IsRewardReceived: false,
		}
	}
	return status
}
