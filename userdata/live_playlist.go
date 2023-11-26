package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) UpdateUserPlayList(item model.UserPlayListItem) {
	session.UserModel.UserPlayListByID.PushBack(item)
}

func userPlayListFinalizer(session *Session) {
	for _, item := range session.UserModel.UserPlayListByID.Objects {
		exist, err := session.Db.Table("u_play_list").
			Where("user_id = ? AND user_play_list_id = ?", session.UserStatus.UserID, item.UserPlayListID).
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
	addGenericTableFieldPopulator("u_play_list", "UserPlayListByID")
}
