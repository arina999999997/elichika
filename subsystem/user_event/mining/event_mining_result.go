package mining

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/log"
	"elichika/scheduled_task"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_present"
	"elichika/userdata"

	"fmt"
	"strconv"
	"time"

	"xorm.io/xorm"
)

// finish the event and pay out the reward for everyone who participated
func resultEventScheduledHandler(userdata_db *xorm.Session, task scheduled_task.ScheduledTask) {
	activeEvent := gamedata.Instance.EventActive.GetActiveEventUnix(task.Time)
	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)
	if (activeEvent == nil) || (activeEvent.EventId != eventId) ||
		(activeEvent.EventType != enum.EventType1Mining) || (activeEvent.ResultAt != task.Time) {
		log.Println("Warning: Failed to result event: ", task)
		return
	}

	eventMining := gamedata.Instance.EventActive.GetEventMining()
	eventName := fmt.Sprintf("m.event_mining_title_%d", eventId)
	timePoint := time.Unix(task.Time, 0)
	// point ranking is awarded first, then voltage ranking
	// we will iterate over the point ranking, and get the voltage ranking that way
	pointRank := int32(0)
	pointRankingResults := GetPointRanking(userdata_db, eventId).GetRange(1, 1<<31-1)
	voltageRanking := GetVoltageRanking(userdata_db, eventId)
	for i, result := range pointRankingResults {
		if (i == 0) || (result.Score != pointRankingResults[i-1].Score) {
			pointRank = int32(i + 1)
		}
		session := userdata.GetBasicSession(userdata_db, timePoint, result.Id)
		rewardGroupId := eventMining.GetPointRankingReward(pointRank)
		for _, content := range gamedata.Instance.EventMiningReward[rewardGroupId] {
			user_present.AddPresent(session, client.PresentItem{
				Content:          *content,
				PresentRouteType: enum.PresentRouteTypeEventMiningPointRankingReward,
				PresentRouteId:   generic.NewNullable(eventMining.EventId),
				ParamClient:      generic.NewNullable(strconv.Itoa(int(pointRank))),
				ParamServer: generic.NewNullable(client.LocalizedText{
					DotUnderText: eventName, // this is resolved everytime user fetch present
				}),
			})
		}
		voltageRank, isVoltageRanked := voltageRanking.TiedRankOf(result.Id)
		if isVoltageRanked {
			rewardGroupId := eventMining.GetVoltageRankingReward(int32(voltageRank))
			for _, content := range gamedata.Instance.EventMiningReward[rewardGroupId] {
				user_present.AddPresent(session, client.PresentItem{
					Content:          *content,
					PresentRouteType: enum.PresentRouteTypeEventMiningVoltageRankingReward,
					PresentRouteId:   generic.NewNullable(eventMining.EventId),
					ParamClient:      generic.NewNullable(strconv.Itoa(int(pointRank))),
					ParamServer: generic.NewNullable(client.LocalizedText{
						DotUnderText: eventName, // this is resolved everytime user fetch present
					}),
				})
			}
		}

		user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
			InfoTriggerType: enum.InfoTriggerTypeEventMiningShowResult,
			ParamInt:        generic.NewNullable(eventId),
		})
		session.Finalize()
	}

	// schedule the event actual end
	scheduled_task.AddScheduledTask(scheduled_task.ScheduledTask{
		Time:     activeEvent.EndAt,
		TaskName: "event_mining_end",
		Params:   task.Params,
	})

}

func init() {
	scheduled_task.AddScheduledTaskHandler("event_mining_result", resultEventScheduledHandler)
}
