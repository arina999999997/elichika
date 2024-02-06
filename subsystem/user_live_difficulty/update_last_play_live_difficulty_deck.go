package user_live_difficulty

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func UpdateLastPlayLiveDifficultyDeck(session*userdata.Session, deck client.LastPlayLiveDifficultyDeck) {
	affected, err := session.Db.Table("u_last_play_live_difficulty_deck").Where("user_id = ? and live_difficulty_id = ?", session.UserId, deck.LiveDifficultyId).
		AllCols().Update(&deck)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_last_play_live_difficulty_deck", deck)
	}
}
