package marathon

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserEventMarathon(session *userdata.Session) client.UserEventMarathon {
	event := session.Gamedata.EventActive.GetEventMarathon()
	ptr, exist := session.UserModel.UserEventMarathonByEventMasterId.Get(event.EventId)
	if exist {
		return *ptr
	}
	userEventMarathon := client.UserEventMarathon{}
	exist, err := session.Db.Table("u_event_marathon").Where("user_id = ? AND event_master_id = ?", session.UserId, event.EventId).
		Get(&userEventMarathon)
	utils.CheckErr(err)
	if !exist {
		userEventMarathon = client.UserEventMarathon{
			EventMasterId:     event.EventId,
			EventPoint:        0,
			OpenedStoryNumber: 1,
			ReadStoryNumber:   0,
		}
	}
	session.UserModel.UserEventMarathonByEventMasterId.Set(event.EventId, userEventMarathon)
	return userEventMarathon
}
