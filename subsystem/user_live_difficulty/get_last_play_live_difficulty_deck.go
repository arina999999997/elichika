package user_live_difficulty

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func GetLastPlayLiveDifficultyDeck(session*userdata.Session, liveDifficultyId int32) generic.Nullable[client.LastPlayLiveDifficultyDeck] {
	lastPlayDeck := client.LastPlayLiveDifficultyDeck{}

	exist, err := session.Db.Table("u_last_play_live_difficulty_deck").
		Where("user_id = ? AND live_difficulty_id = ?", session.UserId, liveDifficultyId).
		Get(&lastPlayDeck)
	utils.CheckErr(err)
	if !exist {
		return generic.Nullable[client.LastPlayLiveDifficultyDeck]{}
	} else {
		return generic.NewNullable(lastPlayDeck)
	}
}