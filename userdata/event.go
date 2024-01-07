package userdata

import (
	"elichika/utils"
)

func userEventFinalizer(session *Session) {
	for _, userEvent := range session.UserModel.UserEventMarathonByEventMasterId.Objects {
		affected, err := session.Db.Table("u_event_marathon").Where("user_id = ? AND event_master_id = ?",
			session.UserStatus.UserId, userEvent.EventMasterId).AllCols().Update(userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_event_marathon").Insert(userEvent)
			utils.CheckErr(err)
		}
	}
	for _, userEvent := range session.UserModel.UserEventMiningByEventMasterId.Objects {
		affected, err := session.Db.Table("u_event_mining").Where("user_id = ? AND event_master_id = ?",
			session.UserStatus.UserId, userEvent.EventMasterId).AllCols().Update(userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_event_mining").Insert(userEvent)
			utils.CheckErr(err)
		}
	}
	for _, userEvent := range session.UserModel.UserEventCoopByEventMasterId.Objects {
		affected, err := session.Db.Table("u_event_coop").Where("user_id = ? AND event_master_id = ?",
			session.UserStatus.UserId, userEvent.EventMasterId).AllCols().Update(userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_event_coop").Insert(userEvent)
			utils.CheckErr(err)
		}
	}
}
func init() {
	addFinalizer(userEventFinalizer)

	addGenericTableFieldPopulator("u_event_marathon", "UserEventMarathonByEventMasterId")
	addGenericTableFieldPopulator("u_event_mining", "UserEventMiningByEventMasterId")
	addGenericTableFieldPopulator("u_event_coop", "UserEventCoopByEventMasterId")
}
