package user_live_difficulty

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateLiveDifficulty(session*userdata.Session, userLiveDifficulty client.UserLiveDifficulty) {
	session.UserModel.UserLiveDifficultyByDifficultyId.Set(userLiveDifficulty.LiveDifficultyId, userLiveDifficulty)
}