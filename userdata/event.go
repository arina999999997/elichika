package userdata

import (
	"elichika/utils"
)

func userEventFinalizer(session *Session) {
	for _, userEvent := range session.UserModel.UserEventMarathonByEventMasterID.Objects {
		affected, err := session.Db.Table("u_event_marathon").Where("user_id = ? AND event_master_id = ?",
			session.UserStatus.UserID, userEvent.EventMasterID).AllCols().Update(userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_event_marathon").Insert(userEvent)
			utils.CheckErr(err)
		}
	}
	for _, userEvent := range session.UserModel.UserEventMiningByEventMasterID.Objects {
		affected, err := session.Db.Table("u_event_mining").Where("user_id = ? AND event_master_id = ?",
			session.UserStatus.UserID, userEvent.EventMasterID).AllCols().Update(userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_event_mining").Insert(userEvent)
			utils.CheckErr(err)
		}
	}
	for _, userEvent := range session.UserModel.UserEventCoopByEventMasterID.Objects {
		affected, err := session.Db.Table("u_event_coop").Where("user_id = ? AND event_master_id = ?",
			session.UserStatus.UserID, userEvent.EventMasterID).AllCols().Update(userEvent)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_event_coop").Insert(userEvent)
			utils.CheckErr(err)
		}
	}
}
func init() {
	addFinalizer(userEventFinalizer)

	addGenericTableFieldPopulator("u_event_marathon", "UserEventMarathonByEventMasterID")
	addGenericTableFieldPopulator("u_event_mining", "UserEventMiningByEventMasterID")
	addGenericTableFieldPopulator("u_event_coop", "UserEventCoopByEventMasterID")
}
