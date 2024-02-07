package user_play_list

import (
	"elichika/userdata"
	"elichika/utils"
)

func userPlayListFinalizer(session *userdata.Session) {
	for userPlayListId, userPlayList := range session.UserModel.UserPlayListById.Map {
		if userPlayList == nil {
			_, err := session.Db.Table("u_play_list").
				Where("user_id = ? AND user_play_list_id = ?", session.UserId, userPlayListId).
				Delete()
			utils.CheckErr(err)
		} else {
			userdata.GenericDatabaseInsert(session, "u_play_list", *userPlayList)
		}
	}
}

func init() {
	userdata.AddFinalizer(userPlayListFinalizer)
}
