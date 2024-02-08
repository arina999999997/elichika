package user_story_side

import (
	"elichika/userdata"
	"elichika/utils"
)

func userStorySideFinalizer(session *userdata.Session) {
	for _, userStorySide := range session.UserModel.UserStorySideById.Map {
		affected, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
			session.UserId, userStorySide.StorySideMasterId).AllCols().Update(userStorySide)
		utils.CheckErr(err)
		if affected == 0 { // need to insert
			userdata.GenericDatabaseInsert(session, "u_story_side", *userStorySide)
		}
	}
}

func init() {
	userdata.AddFinalizer(userStorySideFinalizer)
}
