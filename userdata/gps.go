package userdata

import (
	"elichika/utils"
)

func gpsPresentReceivedFinalizer(session *Session) {
	for _, userGpsPresentReceived := range session.UserModel.UserGpsPresentReceivedById.Map {
		affected, err := session.Db.Table("u_gps_present_received").
			Where("user_id = ? AND campaign_id = ?",
				session.UserId, userGpsPresentReceived.CampaignId).
			AllCols().Update(*userGpsPresentReceived)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_gps_present_received", *userGpsPresentReceived)
		}
	}
}

func init() {
	addFinalizer(gpsPresentReceivedFinalizer)
	addGenericTableFieldPopulator("u_gps_present_received", "UserGpsPresentReceivedById")
}
