package userdata

import (
	"elichika/utils"
)

func emblemFinalizer(session *Session) {
	for _, userEmblem := range session.UserModel.UserEmblemByEmblemId.Objects {
		affected, err := session.Db.Table("u_emblem").Where("user_id = ? AND emblem_m_id = ?",
			session.UserId, userEmblem.EmblemMId).AllCols().Update(userEmblem)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_emblem", userEmblem)
		}
	}
}
func init() {
	addFinalizer(emblemFinalizer)
	addGenericTableFieldPopulator("u_emblem", "UserEmblemByEmblemId")
}