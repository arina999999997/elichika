package user_mission

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

// note that this function doesn't handle carrying over / initial progress.
// whoever unlock the mission is responsible for filling in initial / transfering the progress.
// TODO(mission): maybe move this docs
// - this will depend on the MissionClearConditionType
// - the default behavior is as follow:
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
	}
	return *ptr
}
