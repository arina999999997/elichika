package user_card

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
	"elichika/utils"
)

func missionClearConditionTypeCountPassiveSkillInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]

	userMission.MissionCount = 0
	cards := []client.UserCard{}
	err := session.Db.Table("u_card").Where("user_id = ?", session.UserId).Find(&cards)
	utils.CheckErr(err)

	for _, card := range cards {
		if card.AdditionalPassiveSkill1Id != 0 {
			userMission.MissionCount++
		}
		if card.AdditionalPassiveSkill2Id != 0 {
			userMission.MissionCount++
		}
		if card.AdditionalPassiveSkill3Id != 0 {
			userMission.MissionCount++
		}
		if card.AdditionalPassiveSkill4Id != 0 {
			userMission.MissionCount++
		}
	}
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeCountPassiveSkill, missionClearConditionTypeCountPassiveSkillInitializer)
}
