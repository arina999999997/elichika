package userdata

import (
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func emblemFinalizer(session *Session) {
	for _, userEmblem := range session.UserModel.UserEmblemByEmblemId.Map {
		affected, err := session.Db.Table("u_emblem").Where("user_id = ? AND emblem_m_id = ?",
			session.UserId, userEmblem.EmblemMId).AllCols().Update(userEmblem)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_emblem", *userEmblem)
		}
	}
}
func init() {
	AddFinalizer(emblemFinalizer)
}
