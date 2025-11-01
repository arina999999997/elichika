package mining

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func FetchUserInfoTriggerEventMiningShowResultRows(session *userdata.Session, result *generic.List[client.UserInfoTriggerEventMiningShowResultRow]) {
	// this show the popup result sequence, only possible when the event is still available
	// the reward is already delivered when the event end to present box, and can't be missed
	activeEvent := session.Gamedata.EventActive.GetEventValue()
	if (activeEvent == nil) || (activeEvent.EventType != enum.EventType1Mining) {
		return
	}
	eventMining := session.Gamedata.EventActive.GetEventMining()
	resultTriggers := []client.UserInfoTriggerBasic{}
	err := session.Db.Table("u_info_trigger_basic").Where("user_id = ? AND info_trigger_type = ? AND param_int = ?",
		session.UserId, enum.InfoTriggerTypeEventMiningShowResult, eventMining.EventId).Find(&resultTriggers)
	utils.CheckErr(err)
	for _, trigger := range resultTriggers {
		result.Append(client.UserInfoTriggerEventMiningShowResultRow{
			TriggerId:     trigger.TriggerId,
			EventMiningId: eventMining.EventId,
			ResultAt:      activeEvent.ResultAt,
			EndAt:         activeEvent.EndAt,
		})
	}
}
