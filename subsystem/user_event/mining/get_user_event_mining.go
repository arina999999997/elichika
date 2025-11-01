package mining

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserEventMining(session *userdata.Session) client.UserEventMining {
	event := session.Gamedata.EventActive.GetEventMining()
	ptr, exist := session.UserModel.UserEventMiningByEventMasterId.Get(event.EventId)
	if exist {
		return *ptr
	}
	userEventMining := client.UserEventMining{}
	exist, err := session.Db.Table("u_event_mining").Where("user_id = ? AND event_master_id = ?", session.UserId, event.EventId).
		Get(&userEventMining)
	utils.CheckErr(err)
	if !exist {
		userEventMining = client.UserEventMining{
			EventMasterId:     event.EventId,
			EventPoint:        0,
			EventVoltagePoint: 0,
			OpenedStoryNumber: 1,
			ReadStoryNumber:   0,
		}
	}
	session.UserModel.UserEventMiningByEventMasterId.Set(event.EventId, userEventMining)
	return userEventMining
}
