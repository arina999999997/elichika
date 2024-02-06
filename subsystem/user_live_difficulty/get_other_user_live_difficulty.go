package user_live_difficulty

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserLiveDifficulty(session *userdata.Session, otherUserId int32, liveDifficultyId int32) client.UserLiveDifficulty {
	userLiveDifficulty := client.UserLiveDifficulty{}
	exist, err := session.Db.Table("u_live_difficulty").
		Where("user_id = ? AND live_difficulty_id = ?", otherUserId, liveDifficultyId).
		Get(&userLiveDifficulty)
	utils.CheckErr(err)
	if !exist {
		userLiveDifficulty.LiveDifficultyId = liveDifficultyId
		userLiveDifficulty.EnableAutoplay = true
		userLiveDifficulty.IsNew = true
	}
	return userLiveDifficulty
}
