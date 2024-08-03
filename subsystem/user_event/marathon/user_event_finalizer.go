package marathon

import (
	"elichika/userdata"
	"elichika/utils"
)

func userEventMarathonFinalizer(session *userdata.Session) {
	for _, userEvent := range session.UserModel.UserEventMarathonByEventMasterId.Map {
		affected, err := session.Db.Table("u_event_marathon").Where("user_id = ? AND event_master_id = ?",
			session.UserId, userEvent.EventMasterId).AllCols().Update(*userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_event_marathon", *userEvent)
		}
	}
}
func init() {
	userdata.AddFinalizer(userEventMarathonFinalizer)
}
