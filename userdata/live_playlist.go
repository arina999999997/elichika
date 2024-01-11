package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func (session *Session) AddUserPlayList(userPlayList client.UserPlayList) {
	session.UserModel.UserPlayListById.Set(userPlayList.UserPlayListId, userPlayList)
}
func (session *Session) DeleteUserPlayList(userPlayListId int32) {
	session.UserModel.UserPlayListById.SetNull(userPlayListId)
}

func userPlayListFinalizer(session *Session) {
	for userPlayListId, userPlayList := range session.UserModel.UserPlayListById.Map {
		if userPlayList == nil {
			_, err := session.Db.Table("u_play_list").
				Where("user_id = ? AND user_play_list_id = ?", session.UserId, userPlayListId).
				Delete(userPlayList)
			utils.CheckErr(err)
		} else {
			genericDatabaseInsert(session, "u_play_list", *userPlayList)
		}

	}
}

func init() {
	addFinalizer(userPlayListFinalizer)
	addGenericTableFieldPopulator("u_play_list", "UserPlayListById")
}
