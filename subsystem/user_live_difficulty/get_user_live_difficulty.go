package user_live_difficulty

import (
	"elichika/client"
	"elichika/userdata"
)

func GetUserLiveDifficulty(session *userdata.Session, liveDifficultyId int32) client.UserLiveDifficulty {
	return GetOtherUserLiveDifficulty(session, session.UserId, liveDifficultyId)
}
