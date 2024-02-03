package user_card

import (
	"elichika/userdata"
	"elichika/utils"
)

func cardContentFinalizer(session *userdata.Session) {
	for _, card := range session.UserModel.UserCardByCardId.Map {
		affected, err := session.Db.Table("u_card").
			Where("user_id = ? AND card_master_id = ?", session.UserId, card.CardMasterId).AllCols().Update(*card)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_card", *card)
		}
	}
}

func init() {
	userdata.AddFinalizer(cardContentFinalizer)
}
