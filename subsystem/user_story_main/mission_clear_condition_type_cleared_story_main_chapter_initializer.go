package user_story_main

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
)

func missionClearConditionTypeClearedStoryMainChapterInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	chapterId := mission.MissionClearConditionCount
	requiredCell := session.Gamedata.StoryMainChapter[chapterId].LastCellId
	if hasStoryMainCell(session, requiredCell) {
		userMission.MissionCount = chapterId
		userMission.IsCleared = true
	} else {
		userMission.MissionCount = 0
		userMission.IsCleared = false
	}
	return userMission
}
func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeClearedStoryMainChapter, missionClearConditionTypeClearedStoryMainChapterInitializer)
}
