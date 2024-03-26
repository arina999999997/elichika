package user_live_difficulty

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
)

func missionClearConditionTypeClearedSRankInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	liveDifficultyId := *mission.MissionClearConditionParam1
	liveDifficulty := session.Gamedata.LiveDifficulty[liveDifficultyId]
	userLiveDifficulty := GetUserLiveDifficulty(session, liveDifficultyId)
	if userLiveDifficulty.MaxScore >= liveDifficulty.EvaluationSScore {
		userMission.MissionCount = 1
	} else {
		userMission.MissionCount = 0
	}
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeClearedSRank, missionClearConditionTypeClearedSRankInitializer)
}
