package userdata

import (
	"elichika/utils"
)

func gpsPresentReceivedFinalizer(session *Session) {
	for _, userGpsPresentReceived := range session.UserModel.UserGpsPresentReceivedByID.Objects {
		affected, err := session.Db.Table("u_gps_present_received").
			Where("user_id = ? AND campaign_id = ?",
				session.UserStatus.UserID, userGpsPresentReceived.CampaignID).
			AllCols().Update(userGpsPresentReceived)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_gps_present_received").
				Insert(userGpsPresentReceived)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(gpsPresentReceivedFinalizer)
	addGenericTableFieldPopulator("u_gps_present_received", "UserGpsPresentReceivedByID")
}
