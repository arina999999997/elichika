package marathon

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func FetchUserInfoTriggerEventMarathonShowResultRows(session *userdata.Session, result *generic.List[client.UserInfoTriggerEventMarathonShowResultRow]) {
	// this show the popup result sequence, only possible when the event is still available
	// the reward is already delivered when the event end to present box, and can't be missed
	activeEvent := session.Gamedata.EventActive.GetEventValue()
	if (activeEvent == nil) || (activeEvent.EventType != enum.EventType1Marathon) {
		return
	}
	eventMarathon := session.Gamedata.EventActive.GetEventMarathon()
	resultTriggers := []client.UserInfoTriggerBasic{}
	err := session.Db.Table("u_info_trigger_basic").Where("user_id = ? AND info_trigger_type = ? AND param_int = ?",
		session.UserId, enum.InfoTriggerTypeEventMarathonShowResult, eventMarathon.EventId).Find(&resultTriggers)
	utils.CheckErr(err)
	for _, trigger := range resultTriggers {
		result.Append(client.UserInfoTriggerEventMarathonShowResultRow{
			TriggerId:       trigger.TriggerId,
			EventMarathonId: eventMarathon.EventId,
			ResultAt:        activeEvent.ResultAt,
			EndAt:           activeEvent.EndAt,
		})
	}
}
