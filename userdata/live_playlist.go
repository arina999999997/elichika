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
			Where("user_id = ? AND user_play_list_id = ?", session.UserStatus.UserId, item.UserPlayListId).
			AllCols().Update(&item)
		utils.CheckErr(err)
		if exist == 0 {
			exist, err = session.Db.Table("u_play_list").Insert(item)
			utils.CheckErrMustExist(err, exist != 0)
		}
	}
}

func init() {
	addFinalizer(userPlayListFinalizer)
	addGenericTableFieldPopulator("u_play_list", "UserPlayListById")
}
