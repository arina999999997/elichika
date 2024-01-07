package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) UpdateUserPlayList(item model.UserPlayListItem) {
	session.UserModel.UserPlayListById.PushBack(item)
}

func userPlayListFinalizer(session *Session) {
	for _, item := range session.UserModel.UserPlayListById.Objects {
		exist, err := session.Db.Table("u_play_list").
			Where("user_id = ? AND user_play_list_id = ?", session.UserId, item.UserPlayListId).
			AllCols().Update(&item)
		utils.CheckErr(err)
		if exist == 0 {
			genericDatabaseInsert(session, "u_play_list", item)
		}
	}
}

func init() {
	addFinalizer(userPlayListFinalizer)
	addGenericTableFieldPopulator("u_play_list", "UserPlayListById")
}
