package user_event

import (
	"elichika/userdata"
	"elichika/utils"
)

func userEventFinalizer(session *userdata.Session) {
	for _, userEvent := range session.UserModel.UserEventMarathonByEventMasterId.Map {
		affected, err := session.Db.Table("u_event_marathon").Where("user_id = ? AND event_master_id = ?",
			session.UserId, userEvent.EventMasterId).AllCols().Update(*userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_event_marathon", *userEvent)
		}
	}
	for _, userEvent := range session.UserModel.UserEventMiningByEventMasterId.Map {
		affected, err := session.Db.Table("u_event_mining").Where("user_id = ? AND event_master_id = ?",
			session.UserId, userEvent.EventMasterId).AllCols().Update(*userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_event_mining", *userEvent)
		}
	}
	for _, userEvent := range session.UserModel.UserEventCoopByEventMasterId.Map {
		affected, err := session.Db.Table("u_event_coop").Where("user_id = ? AND event_master_id = ?",
			session.UserId, userEvent.EventMasterId).AllCols().Update(*userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_event_coop", *userEvent)
		}
	}
}
func init() {
	userdata.AddFinalizer(userEventFinalizer)
}
