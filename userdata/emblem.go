package userdata

import (
	"elichika/utils"
)

func emblemFinalizer(session *Session) {
	for _, userEmblem := range session.UserModel.UserEmblemByEmblemID.Objects {
		affected, err := session.Db.Table("u_emblem").Where("user_id = ? AND emblem_m_id = ?",
			session.UserStatus.UserID, userEmblem.EmblemMID).AllCols().Update(userEmblem)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_emblem").Insert(userEmblem)
			utils.CheckErr(err)
		}
	}
}
func init() {
	addFinalizer(emblemFinalizer)
	addGenericTableFieldPopulator("u_emblem", "UserEmblemByEmblemID")
}
