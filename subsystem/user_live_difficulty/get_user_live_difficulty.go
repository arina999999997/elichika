package user_live_difficulty

import (
	"elichika/userdata"
	"elichika/client"
)

func GetUserLiveDifficulty(session *userdata.Session, liveDifficultyId int32) client.UserLiveDifficulty {
	return GetOtherUserLiveDifficulty(session, session.UserId, liveDifficultyId)
}