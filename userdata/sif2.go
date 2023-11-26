package userdata

import (
	"elichika/utils"
)

func sif2DataLinkFinalizer(session *Session) {
	for _, userSif2DataLink := range session.UserModel.UserSif2DataLinkByID.Objects {
		affected, err := session.Db.Table("u_sif_2_data_link").
			Where("user_id = ? AND sif_2_id = ?",
				session.UserStatus.UserID, userSif2DataLink.Sif2ID).
			AllCols().Update(userSif2DataLink)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_sif_2_data_link").
				Insert(userSif2DataLink)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(sif2DataLinkFinalizer)
	addGenericTableFieldPopulator("u_sif_2_data_link", "UserSif2DataLinkByID")
}
