package user_gps_present

import (
	"elichika/userdata"
	"elichika/utils"
)

func gpsPresentReceivedFinalizer(session *userdata.Session) {
	for _, userGpsPresentReceived := range session.UserModel.UserGpsPresentReceivedById.Map {
		affected, err := session.Db.Table("u_gps_present_received").
			Where("user_id = ? AND campaign_id = ?",
				session.UserId, userGpsPresentReceived.CampaignId).
			AllCols().Update(*userGpsPresentReceived)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_gps_present_received", *userGpsPresentReceived)
		}
	}
}

func init() {
	userdata.AddFinalizer(gpsPresentReceivedFinalizer)
}
