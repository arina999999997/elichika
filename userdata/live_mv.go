package userdata

import (
	"elichika/utils"
)

func liveMvFinalizer(session *Session) {
	for _, userLiveMv := range session.UserModel.UserLiveMvDeckById.Objects {
		affected, err := session.Db.Table("u_live_mv_deck").Where("user_id = ? AND live_master_id = ?",
			session.UserId, userLiveMv.LiveMasterId).AllCols().Update(userLiveMv)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_live_mv_deck", userLiveMv)
		}
	}
	for _, userLiveMvCustom := range session.UserModel.UserLiveMvDeckCustomById.Objects {
		affected, err := session.Db.Table("u_live_mv_deck_custom").Where("user_id = ? AND live_master_id = ?",
			session.UserId, userLiveMvCustom.LiveMasterId).AllCols().Update(userLiveMvCustom)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_live_mv_deck_custom", userLiveMvCustom)
		}
	}
}
func init() {
	addFinalizer(liveMvFinalizer)
	addGenericTableFieldPopulator("u_live_mv_deck", "UserLiveMvDeckById")
	addGenericTableFieldPopulator("u_live_mv_deck_custom", "UserLiveMvDeckCustomById")
}