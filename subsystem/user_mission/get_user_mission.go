package user_mission

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

// note that this function doesn't handle carrying over / initial progress.
// whoever unlock the mission is responsible for filling in initial / transfering the progress.
// - this will depend on the MissionClearConditionType
// - the default behavior should work for a bunch of case, but not every:
//   - if the unlocked mission has the same MissionClearConditionType and MissionClearConditionParam1/2 with the parent
//     mission, then the progress is carried over
//   - otherwise the new mission progress is left at 0

func getUserMission(session *userdata.Session, missionId int32) client.UserMission {
	// check if the mission is valid and update it
	if !hasMission(session, missionId) {
		return client.UserMission{}
	}

	ptr, exist := session.UserModel.UserMissionByMissionId.Get(missionId)
	if exist {
		return *ptr
	}
	ptr = new(client.UserMission)
	exist, err := session.Db.Table("u_mission").Where("user_id = ? AND mission_m_id = ?",
		session.UserId, missionId).Get(ptr)
	utils.CheckErr(err)
	if !exist { // create an empty mission
		*ptr = client.UserMission{
			MissionMId:       missionId,
			IsNew:            true,
			MissionCount:     0,
			IsCleared:        false,
			IsReceivedReward: false,
			NewExpiredAt:     0,
		}
		// carry over old progress
		mission := session.Gamedata.Mission[missionId]
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
				ptr.MissionCount = getUserMission(session, mission.TriggerCondition1).MissionCount
				ptr.IsCleared = ptr.MissionCount >= parent.MissionClearConditionCount
			}
		}
	}
	return *ptr
}
