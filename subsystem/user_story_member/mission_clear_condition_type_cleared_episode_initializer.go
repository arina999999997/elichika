package user_story_member

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
	"elichika/utils"
)

func missionClearConditionTypeClearedEpisodeInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	count, err := session.Db.Table("u_story_member").Where("user_id = ?", session.UserId).Count()
	utils.CheckErr(err)
	userMission.MissionCount = int32(count)
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}
func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeClearedEpisode, missionClearConditionTypeClearedEpisodeInitializer)
}
