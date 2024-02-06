package user_live_deck

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserLiveDeck(session *userdata.Session, userLiveDeckId int32) client.UserLiveDeck {
	liveDeck := client.UserLiveDeck{}
	exist, err := session.Db.Table("u_live_deck").
		Where("user_id = ? AND user_live_deck_id = ?", session.UserId, userLiveDeckId).
		Get(&liveDeck)
	utils.CheckErr(err)
	if !exist {
		panic("Deck doesn't exist")
	}
	return liveDeck
}
