package marathon

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
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
func resultEventScheduledHandler(serverdata_db *xorm.Session, userdata_db *xorm.Session, task scheduled_task.ScheduledTask) {
	activeEvent := gamedata.Instance.EventActive.GetActiveEventUnix(task.Time)
	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)
	if (activeEvent == nil) || (activeEvent.EventId != eventId) ||
		(activeEvent.EventType != enum.EventType1Marathon) || (activeEvent.ResultAt != task.Time) {
		fmt.Println("Warning: Failed to result event: ", task)
		return
	}

	results := GetRanking(userdata_db, eventId).GetRange(1, 1<<31-1)
	eventMarathon := gamedata.Instance.EventActive.GetEventMarathon()
	rank := int32(0)
	timePoint := time.Unix(task.Time, 0)
	for i, result := range results {
		if (i == 0) || (result.Score != results[i-1].Score) {
			rank = int32(i + 1)
		}
		session := userdata.GetBasicSession(userdata_db, serverdata_db, timePoint, result.Id)
		rewardGroupId := eventMarathon.GetRankingReward(rank)
		for _, content := range gamedata.Instance.EventMarathonReward[rewardGroupId] {
			user_present.AddPresent(session, client.PresentItem{
				Content:          *content,
				PresentRouteType: enum.PresentRouteTypeEventMarathonRankingReward,
				PresentRouteId:   generic.NewNullable(eventMarathon.EventId),
				// TODO(localization): Fill in param_client and param_server to show event and ranking
				ParamClient: generic.NewNullable(strconv.Itoa(int(rank))),
				ParamServer: generic.NewNullable(client.LocalizedText{
					DotUnderText: "Secret Party!", // client doesn't resolve this automatically from number
				}),
			})
		}

		user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
			InfoTriggerType: enum.InfoTriggerTypeEventMarathonShowResult,
			ParamInt:        generic.NewNullable(eventId),
		})
		session.Finalize()
	}

	// schedule the event actual end
	scheduled_task.AddScheduledTask(serverdata_db, scheduled_task.ScheduledTask{
		Time:     activeEvent.EndAt,
		TaskName: "event_marathon_end",
		Params:   task.Params,
	})

}

func init() {
	scheduled_task.AddScheduledTaskHandler("event_marathon_result", resultEventScheduledHandler)
}
