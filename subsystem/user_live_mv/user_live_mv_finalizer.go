package user_live_mv

import (
	"elichika/userdata"
	"elichika/utils"
)

func userLiveMvFinalizer(session *userdata.Session) {
	for _, userLiveMv := range session.UserModel.UserLiveMvDeckById.Map {
		affected, err := session.Db.Table("u_live_mv_deck").Where("user_id = ? AND live_master_id = ?",
			session.UserId, userLiveMv.LiveMasterId).AllCols().Update(*userLiveMv)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_live_mv_deck", *userLiveMv)
		}
	}
	for _, userLiveMvCustom := range session.UserModel.UserLiveMvDeckCustomById.Map {
		affected, err := session.Db.Table("u_live_mv_deck_custom").Where("user_id = ? AND live_master_id = ?",
			session.UserId, userLiveMvCustom.LiveMasterId).AllCols().Update(*userLiveMvCustom)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_live_mv_deck_custom", *userLiveMvCustom)
		}
	}
}
func init() {
	userdata.AddFinalizer(userLiveMvFinalizer)
}
