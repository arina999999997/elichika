package userdata

import (
	"elichika/utils"
)

func sif2DataLinkFinalizer(session *Session) {
	for _, userSif2DataLink := range session.UserModel.UserSif2DataLinkById.Map {
		affected, err := session.Db.Table("u_sif_2_data_link").
			Where("user_id = ? AND sif_2_id = ?",
				session.UserId, userSif2DataLink.Sif2Id).
			AllCols().Update(*userSif2DataLink)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_sif_2_data_link", *userSif2DataLink)
		}
	}
}

func init() {
	addFinalizer(sif2DataLinkFinalizer)
	addGenericTableFieldPopulator("u_sif_2_data_link", "UserSif2DataLinkById")
}
