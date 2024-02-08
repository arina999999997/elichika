package user_sif_2_data_link

import (
	"elichika/userdata"
	"elichika/utils"
)

func userSif2DataLinkFinalizer(session *userdata.Session) {
	for _, userSif2DataLink := range session.UserModel.UserSif2DataLinkById.Map {
		affected, err := session.Db.Table("u_sif_2_data_link").
			Where("user_id = ? AND sif_2_id = ?",
				session.UserId, userSif2DataLink.Sif2Id).
			AllCols().Update(*userSif2DataLink)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_sif_2_data_link", *userSif2DataLink)
		}
	}
}

func init() {
	userdata.AddFinalizer(userSif2DataLinkFinalizer)
}
