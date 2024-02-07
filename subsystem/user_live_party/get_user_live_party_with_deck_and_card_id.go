package user_live_party

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserLivePartyWithDeckAndCardId(session *userdata.Session, deckId, cardId int32) client.UserLiveParty {
	liveParty := client.UserLiveParty{}
	exist, err := session.Db.Table("u_live_party").
		Where("user_id = ? AND user_live_deck_id = ? AND (card_master_id_1 = ? OR card_master_id_2 = ? OR card_master_id_3 = ?)",
			session.UserId, deckId, cardId, cardId, cardId).
		Get(&liveParty)
	utils.CheckErr(err)
	if !exist {
		panic("Party doesn't exist")
	}
	return liveParty
}
