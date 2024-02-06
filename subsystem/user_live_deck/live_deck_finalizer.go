package user_live_deck

import (
	"elichika/userdata"
	"elichika/utils"
)

func liveDeckFinalizer(session *userdata.Session) {
	for _, deck := range session.UserModel.UserLiveDeckById.Map {
		affected, err := session.Db.Table("u_live_deck").
			Where("user_id = ? AND user_live_deck_id = ?", session.UserId, deck.UserLiveDeckId).AllCols().
			Update(*deck)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_live_deck", *deck)
		}
	}
}

func init() {
	userdata.AddFinalizer(liveDeckFinalizer)
}
