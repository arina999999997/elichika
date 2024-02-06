package user_live_difficulty

import (
	"elichika/userdata"
)

func UnlockLive(session*userdata.Session, liveId int32) {
	// insert empty record for relevant items
	for _, masterLiveDifficulty := range session.Gamedata.Live[liveId].LiveDifficulties {
		userLiveDifficulty := GetUserLiveDifficulty(session, masterLiveDifficulty.LiveDifficultyId)
		UpdateLiveDifficulty(session, userLiveDifficulty)
	}
}