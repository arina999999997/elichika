package user_mission

import (
	"elichika/client/response"
	"elichika/userdata"
)

func FetchMission(session *userdata.Session) response.FetchMissionResponse {
	session.SendMissionDetail = true
	resp := response.FetchMissionResponse{
		UserModel: &session.UserModel,
	}

	// we have to go through all the mission in the database, then get the user mission, if we want things to actually work correctly
	// getting the mission from the database will result in
	populateMissions(session)
	populateDailyMissions(session)
	populateWeeklyMissions(session)

	for _, mission := range session.UserModel.UserMissionByMissionId.Map {
		resp.MissionMasterIdList.Append(mission.MissionMId)
	}
	for _, mission := range session.UserModel.UserDailyMissionByMissionId.Map {
		resp.MissionMasterIdList.Append(mission.MissionMId)
	}
	for _, mission := range session.UserModel.UserWeeklyMissionByMissionId.Map {
		resp.MissionMasterIdList.Append(mission.MissionMId)
	}
	return resp
}
