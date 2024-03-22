package user_mission

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
)

// mission initializer will be called by condition type, or default to the default behaviour
// - things that are tracked independent of mission progress like user_rank, number of costume, ... should use custom initializer
// - with these missions, maybe we can use the initializer to track progress too, although that might be slow
// - sometime we might even have to combine initializer with delta patching

type MissionInitializer = func(*userdata.Session, client.UserMission) client.UserMission

var missionInitializers = map[int32]MissionInitializer{}

func AddMissionInitializer(missionClearConditionType int32, initializer MissionInitializer) {
	_, exist := missionInitializers[missionClearConditionType]
	if exist {
		panic("mission initializer already exist")
	}
	missionInitializers[missionClearConditionType] = initializer
}

func defaultInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	if mission.TriggerType == enum.MissionTriggerClearMission {
		parent := session.Gamedata.Mission[mission.TriggerCondition1]
		keepProgress := (parent.MissionClearConditionType == mission.MissionClearConditionType)
		keepProgress = keepProgress && ((parent.MissionClearConditionParam1 == nil) == (mission.MissionClearConditionParam1 == nil))
		keepProgress = keepProgress && ((parent.MissionClearConditionParam2 == nil) == (mission.MissionClearConditionParam2 == nil))
		keepProgress = keepProgress && ((parent.MissionClearConditionParam1 == nil) ||
			(*parent.MissionClearConditionParam1 == *mission.MissionClearConditionParam1))
		keepProgress = keepProgress && ((parent.MissionClearConditionParam2 == nil) ||
			(*parent.MissionClearConditionParam2 == *mission.MissionClearConditionParam2))
		if keepProgress {
			userMission.MissionCount = getUserMission(session, mission.TriggerCondition1).MissionCount
			userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
		}
	}
	return userMission
}
