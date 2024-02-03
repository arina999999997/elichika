package userdata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/utils"
)

func (session *Session) GetOtherUserLiveDifficulty(otherUserId int32, liveDifficultyId int32) client.UserLiveDifficulty {
	userLiveDifficulty := client.UserLiveDifficulty{}
	exist, err := session.Db.Table("u_live_difficulty").
		Where("user_id = ? AND live_difficulty_id = ?", otherUserId, liveDifficultyId).
		Get(&userLiveDifficulty)
	if err != nil {
		panic(err)
	}
	if !exist {
		// userLiveDifficulty.UserId = otherUserId
		userLiveDifficulty.LiveDifficultyId = liveDifficultyId
		userLiveDifficulty.EnableAutoplay = true
		userLiveDifficulty.IsNew = true
	}
	return userLiveDifficulty
}

func (session *Session) GetUserLiveDifficulty(liveDifficultyId int32) client.UserLiveDifficulty {
	return session.GetOtherUserLiveDifficulty(session.UserId, liveDifficultyId)
}

func (session *Session) GetAllLiveDifficulties() []client.UserLiveDifficulty {
	records := []client.UserLiveDifficulty{}
	err := session.Db.Table("u_live_difficulty").Where("user_id = ?", session.UserId).
		Find(&records)
	utils.CheckErr(err)
	return records
}

func (session *Session) UpdateLiveDifficulty(userLiveDifficulty client.UserLiveDifficulty) {
	session.UserModel.UserLiveDifficultyByDifficultyId.Set(userLiveDifficulty.LiveDifficultyId, userLiveDifficulty)
}

func liveDifficultyFinalizer(session *Session) {
	for _, userLiveDifficulty := range session.UserModel.UserLiveDifficultyByDifficultyId.Map {
		updated, err := session.Db.Table("u_live_difficulty").
			Where("user_id = ? AND live_difficulty_id = ?", session.UserId, userLiveDifficulty.LiveDifficultyId).
			AllCols().Update(*userLiveDifficulty)
		utils.CheckErr(err)
		if updated == 0 {
			GenericDatabaseInsert(session, "u_live_difficulty", *userLiveDifficulty)
		}
	}

}

func (session *Session) GetLastPlayLiveDifficultyDeck(liveDifficultyId int32) generic.Nullable[client.LastPlayLiveDifficultyDeck] {
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

func (session *Session) UpdateLastPlayLiveDifficultyDeck(deck client.LastPlayLiveDifficultyDeck) {
	affected, err := session.Db.Table("u_last_play_live_difficulty_deck").Where("user_id = ? and live_difficulty_id = ?", session.UserId, deck.LiveDifficultyId).
		AllCols().Update(&deck)
	utils.CheckErr(err)
	if affected == 0 {
		GenericDatabaseInsert(session, "u_last_play_live_difficulty_deck", deck)
	}
}

func init() {
	AddFinalizer(liveDifficultyFinalizer)
}
