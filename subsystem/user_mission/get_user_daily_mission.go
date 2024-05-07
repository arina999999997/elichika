package user_mission

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func getUserDailyMission(session *userdata.Session, missionId int32) client.UserDailyMission {
	// check if the mission is valid and update it
	if !hasMission(session, missionId) {
		return client.UserDailyMission{}
	}

	ptr, exist := session.UserModel.UserDailyMissionByMissionId.Get(missionId)
	if exist {
		return *ptr
	}
	ptr = new(client.UserDailyMission)
	exist, err := session.Db.Table("u_daily_mission").Where("user_id = ? AND mission_m_id = ?",
		session.UserId, missionId).Get(ptr)
	utils.CheckErr(err)
	if !exist { // create an empty mission
		*ptr = client.UserDailyMission{
			MissionMId:        missionId,
			IsNew:             true,
			MissionStartCount: 0,
			MissionCount:      0,
			IsCleared:         false,
			IsReceivedReward:  false,
			ClearedExpiredAt:  generic.NewNullable(int64(0)),
		}
	}
	if session.Time.Unix() >= ptr.ClearedExpiredAt.Value { // expired, reset the progress
		ptr.IsNew = true
		ptr.MissionStartCount = ptr.MissionCount
		ptr.IsCleared = false
		ptr.IsReceivedReward = false
		ptr.ClearedExpiredAt = generic.NewNullable(utils.BeginOfNextDay(session.Time).Unix())
	}
	return *ptr
}
